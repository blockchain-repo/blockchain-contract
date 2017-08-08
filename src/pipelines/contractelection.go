package pipelines

import (
	"encoding/json"
	"fmt"
	"runtime"
	"sort"
	"sync"
)

import (
	"unicontract/src/common"
	"unicontract/src/common/monitor"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/config"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
)

//---------------------------------------------------------------------------
const (
	_VERSION          = 2
	_OPERATION        = "CONTRACT"
	_CONSENSUS_TYPE   = "contract"
	_CONSENSUS_REASON = "contract illegal"
	_FULFILLMENT      = "cf:4:RtTtCxNf1Bq7MFeIToEosMAa3v_jKtZUtqiWAXyFz1ejPMv-t7vT6DANcrYvKFHAsZblmZ1Xk03HQdJbGiMyb5CmQqGPHwlgKusNu9N_IDtPn7y16veJ1RBrUP-up4YD"
)

var (
	gslPublicKeys   []string
	gnPublicKeysNum int
	gstrPublicKey   string
)

//---------------------------------------------------------------------------
func getCEChangefeed() *ChangeFeed {
	change := &ChangeFeed{
		db:        rethinkdb.DBNAME,
		table:     rethinkdb.TABLE_VOTES,
		operation: INSERT | UPDATE,
	}
	go change.runForever()
	return change
}

//---------------------------------------------------------------------------
func createCEPip() (cePip Pipeline) {
	ceNodeSlice := make([]*Node, 0)
	ceNodeSlice = append(ceNodeSlice, &Node{target: ceHeadFilter, routineNum: 1, name: "ceHeadFilter"})
	ceNodeSlice = append(ceNodeSlice, &Node{target: ceQueryExists, routineNum: 1, name: "ceQueryExists"})
	cePip = Pipeline{
		nodes: ceNodeSlice,
	}
	return cePip
}

//---------------------------------------------------------------------------
func startContractElection() {
	gslPublicKeys = config.GetAllPublicKey()
	sort.Strings(gslPublicKeys)
	gnPublicKeysNum = len(gslPublicKeys)
	gstrPublicKey = config.Config.Keypair.PublicKey

	cePip := createCEPip()
	ceChangefeed := getCEChangefeed()
	cePip.setup(&ceChangefeed.node)
	cePip.start()

	waitRoutine := sync.WaitGroup{}
	waitRoutine.Add(1)
	waitRoutine.Wait()
}

//---------------------------------------------------------------------------
func ceHeadFilter(arg interface{}) interface{} {
	uniledgerlog.Debug("1.changefeed -> ceHeadFilter")

	// 读取new_val的值
	contract_election_time := monitor.Monitor.NewTiming()
	slVote, err := json.Marshal(arg)
	if err != nil {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.SERIALIZE_ERROR, err.Error()))
		return nil
	}

	// 将new_val转化为vote
	vote := model.Vote{}
	err = json.Unmarshal(slVote, &vote)
	if err != nil {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.DESERIALIZE_ERROR, err.Error()))
		return nil
	}

	// 验证是否为头节点
	uniledgerlog.Debug("1.2 get publickey and verify")
	uniledgerlog.Debug("1.2.1 get publickey")
	mainPubkey, err := rethinkdb.GetContractMainPubkeyByContract(vote.VoteBody.VoteFor)
	if err == nil {
		uniledgerlog.Debug("1.2.2 verify head node")
		if isHead, _ := _verifyHeadNode(mainPubkey); isHead {
			// 将contract更新为正在共识阶段
			rethinkdb.SetContractConsensusResultById(vote.VoteBody.VoteFor,
				common.GenTimestamp(), 1)
			uniledgerlog.Debug("1.2.3 verify vote")
			vote_validate_time := monitor.Monitor.NewTiming()
			pContractOutput, pass, err := _verifyVotes(vote.VoteBody.VoteFor)
			vote_validate_time.Send("vote_validate")
			if err != nil {
				if !pass {
					uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, err.Error()))
				} else {
					// vote not enough
					uniledgerlog.Info(err.Error())
				}
				return nil
			}

			// 将contract更新为共识结束
			rethinkdb.SetContractConsensusResultById(vote.VoteBody.VoteFor,
				common.GenTimestamp(), 2)

			if pass { // 合约合法
				uniledgerlog.Debug("1.3 ceHeadFilter --->")
				ceQueryExists(*pContractOutput)
				contract_election_time.Send("contract_election")
			} else { // 合约不合法
				uniledgerlog.Debug("1.3 contract invalid and insert consensusFailure")
				var consensusFailure model.ConsensusFailure
				consensusFailure.Id = common.GenerateUUID()
				consensusFailure.ConsensusType = _CONSENSUS_TYPE
				consensusFailure.ConsensusId = vote.VoteBody.VoteFor
				consensusFailure.ConsensusReason = _CONSENSUS_REASON
				consensusFailure.Timestamp = common.GenTimestamp()

				slConsensusFailure, err := json.Marshal(consensusFailure)
				if err != nil {
					uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.SERIALIZE_ERROR, err.Error()))
					return nil
				}
				if !rethinkdb.InsertConsensusFailure(string(slConsensusFailure)) {
					uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, "rethinkdb.InsertConsensusFailure error"))
				}
				//consensus_failure_count, err := rethinkdb.GetConsensusFailuresCount()
				//if err != nil {
				//	uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, err.Error()))
				//	continue
				//}
			}
		} else {
			uniledgerlog.Info("I am not head node.")
		}
	} else {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, err.Error()))
	}

	return nil
}

//---------------------------------------------------------------------------
func ceQueryExists(arg interface{}) interface{} {
	uniledgerlog.Debug("2.ceHeadFilter -> ceQueryExists")

	uniledgerlog.Debug("2.2 query contractoutput table")
	contractOutput, ok := arg.(model.ContractOutput)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, "ceQueryExists assert error"))
		return ""
	}
	output, err := rethinkdb.GetContractOutputById(contractOutput.Id)
	if err != nil {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, err.Error()))
		return ""
	}

	slContractOutput, _ := json.Marshal(contractOutput)

	if len(output) == 0 {
		uniledgerlog.Debug("2.3 insert contractoutput table")
		contractOutput_write_time := monitor.Monitor.NewTiming()
		rethinkdb.InsertContractOutput(string(slContractOutput))
		contractOutput_write_time.Send("contractOutput_write")
	} else {
		uniledgerlog.Debug("2.3 contractoutput exist")
	}
	uniledgerlog.Debug("2.4 ceQueryExists ---> /dev/null")
	return ""
}

//---------------------------------------------------------------------------

//---------------------------------------------------------------------------
// 私有工具函数
//---------------------------------------------------------------------------
func _verifyHeadNode(mainPubkey string) (bool, error) {
	ok := false
	if mainPubkey == gstrPublicKey {
		ok = true
	}
	return ok, nil
}

//---------------------------------------------------------------------------
func _verifyVotes(contractId string) (*model.ContractOutput, bool, error) {
	strVotes, err := rethinkdb.GetVotesByContractId(contractId)
	if err != nil {
		return nil, false, err
	}

	var slVote []model.Vote
	err = json.Unmarshal([]byte(strVotes), &slVote)
	if err != nil {
		return nil, false, err
	}

	nTotalVoteNum := len(slVote)
	if nTotalVoteNum == 0 {
		return nil, false, fmt.Errorf("no vote")
	}

	uniledgerlog.Debug("2.2.3.1 verify vote signature")
	eligible_votes := make(map[string]model.Vote)
	for _, tmpVote := range slVote {
		if tmpVote.VerifyVoteSignature() {
			if isExist, _ := _verifyPublicKey(tmpVote.NodePubkey); isExist {
				eligible_votes[tmpVote.NodePubkey] = tmpVote
			}
		}
	}

	uniledgerlog.Debug("2.2.3.2 count votes")
	if len(eligible_votes)*2 < gnPublicKeysNum { // vote没有达到节点数的一半时
		return nil, true, fmt.Errorf("vote not enough")
	}

	uniledgerlog.Debug("2.2.3.3 valid votes")
	nValid := _verifyValid(eligible_votes)

	if nValid == 0 {
		return nil, true, fmt.Errorf("vote not enough")
	} else if nValid == 1 {
		contractOutput, err := _produceContractOutput(contractId, slVote)
		if err != nil {
			return nil, false, err
		}
		return &contractOutput, true, nil
	} else if nValid == 2 {
		return nil, false, nil
	}
	return nil, false, fmt.Errorf("unknow error")
}

//---------------------------------------------------------------------------
func _verifyPublicKey(NodePubkey string) (bool, error) {
	isExist := false
	for _, value := range gslPublicKeys {
		if value == NodePubkey {
			isExist = true
			break
		}
	}
	return isExist, nil
}

//---------------------------------------------------------------------------
// 0 投票不够（例如三个节点，已投两个，一个true，一个false）
// 1 达成共识，合约有效
// 2 达成共识，合约无效
func _verifyValid(mVotes map[string]model.Vote) int {
	var nValidNum, nInValidNum int
	for _, value := range mVotes {
		if value.VoteBody.IsValid {
			nValidNum++
		} else {
			nInValidNum++
		}
	}

	nValid := 0
	if nValidNum*2 > gnPublicKeysNum {
		nValid = 1
	}
	if nInValidNum*2 >= gnPublicKeysNum {
		nValid = 2
	}

	return nValid
}

//---------------------------------------------------------------------------
func _produceContractOutput(contractId string, slVote []model.Vote) (model.ContractOutput, error) {
	var contractOutput model.ContractOutput

	contractOutput.Version = _VERSION
	contractOutput.Transaction.Operation = _OPERATION

	strContract, err := rethinkdb.GetContractById(contractId)
	if err != nil {
		return contractOutput, err
	}
	var contractModel model.ContractModel
	err = json.Unmarshal([]byte(strContract), &contractModel)
	if err != nil {
		return contractOutput, err
	}
	contractOutput.Transaction.ContractModel = contractModel
	contractOutput.Transaction.ContractModel.ContractHead = nil

	contractOutput.Transaction.Relation = new(model.Relation)
	contractOutput.Transaction.Relation.ContractHashId = contractModel.Id
	contractOutput.Transaction.Relation.ContractId = contractModel.ContractBody.ContractId
	for _, value := range gslPublicKeys {
		contractOutput.Transaction.Relation.Voters =
			append(contractOutput.Transaction.Relation.Voters, value)
	}

	//fulfillment := &model.Fulfillment{
	//	Fid:          0,
	//	OwnersBefore: []string{config.Config.Keypair.PublicKey},
	//}
	//contractOutput.Transaction.Fulfillments = append(contractOutput.Transaction.Fulfillments, fulfillment)
	//contractOutput.Transaction.Conditions = []*model.ConditionsItem{}

	contractOutput.Transaction.Conditions = []*model.ConditionsItem{} //todo
	contractOutput.Transaction.Fulfillments = []*model.Fulfillment{}
	contractOutput.Transaction.Asset = &model.Asset{}

	//uniledgerlog.Debug("contractOutput : %+v", common.StructSerialize(contractOutput))
	contractOutput.Id = common.HashData(common.StructSerialize(contractOutput))

	//fulfillment.Fulfillment = _FULFILLMENT
	contractOutput.Transaction.Timestamp = common.GenTimestamp()
	contractOutput.Transaction.ContractModel.ContractHead = contractModel.ContractHead
	contractOutput.Transaction.Relation.Votes = make([]*model.Vote, gnPublicKeysNum)
	for key, _ := range slVote {
		for index, _ := range gslPublicKeys {
			if gslPublicKeys[index] == slVote[key].NodePubkey {
				contractOutput.Transaction.Relation.Votes[index] = &slVote[key]
			}
		}
	}

	return contractOutput, err
}

//---------------------------------------------------------------------------

//---------------------------------------------------------------------------
// 公有工具函数
//---------------------------------------------------------------------------
func PanicRecoverAndOutputStack(outputStack bool) {
	if err := recover(); err != nil {
		fmt.Println("===========================================")
		fmt.Printf("Panic !!!!, err is %+v", err)
		if outputStack {
			var slStack [4096]byte
			nReadNum := runtime.Stack(slStack[:], true)
			fmt.Printf(string(slStack[:nReadNum]))
		}
		fmt.Println("===========================================")
	}
}

//---------------------------------------------------------------------------
