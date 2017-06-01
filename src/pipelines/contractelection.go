package pipelines

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
)

import (
	beegoLog "github.com/astaxie/beego/logs"
)

import (
	"unicontract/src/common"
	"unicontract/src/common/monitor"
	"unicontract/src/config"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
)

//---------------------------------------------------------------------------
const (
	_NEW_VAL          = "new_val"
	_HTTP_OK          = 200
	_VERSION          = 2
	_OPERATION        = "CONTRACT"
	_CONSENSUS_TYPE   = "contract"
	_CONSENSUS_REASON = "contract illegal"
	_THREAD_NUM       = 10
	_INPUT_LEN        = 20
	_OUTPUT_LEN       = 20
	_FULFILLMENT      = "cf:4:RtTtCxNf1Bq7MFeIToEosMAa3v_jKtZUtqiWAXyFz1ejPMv-t7vT6DANcrYvKFHAsZblmZ1Xk03HQdJbGiMyb5CmQqGPHwlgKusNu9N_IDtPn7y16veJ1RBrUP-up4YD"
)

var (
	gslPublicKeys   []string
	gnPublicKeysNum int
	gstrPublicKey   string

	gPool     *ThreadPool
	gchInput  chan string
	gchOutput chan string
)

//---------------------------------------------------------------------------
func startContractElection() {
	defer gPool.Stop()
	defer close(gchInput)

	gslPublicKeys = config.GetAllPublicKey()
	sort.Strings(gslPublicKeys)
	gnPublicKeysNum = len(gslPublicKeys)
	gstrPublicKey = config.Config.Keypair.PublicKey

	gchInput = make(chan string, _INPUT_LEN)
	gchOutput = make(chan string, _OUTPUT_LEN)

	gPool = new(ThreadPool)
	gPool.Init(_THREAD_NUM)
	for i := 0; i < _THREAD_NUM; i++ {
		gPool.AddTask(func() error {
			return ceHeadFilter()
		})
	}
	go gPool.Start()

	go ceChangefeed()

	for {
		if out, ok := <-gchOutput; ok {
			f, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
			if err != nil {
				beegoLog.Error(err.Error())
			}
			f.Write([]byte(out))
		} else {
			break
		}
	}
}

//---------------------------------------------------------------------------
func ceChangefeed() {
	beegoLog.Debug("1.进入ceChangefeed")
	var value interface{}
	res := rethinkdb.Changefeed(rethinkdb.DBNAME, rethinkdb.TABLE_VOTES)
	for res.Next(&value) {
		votes_changefeed_time := monitor.Monitor.NewTiming()
		beegoLog.Debug("1.1 ceChangefeed get new_val")
		mValue := value.(map[string]interface{})
		// 提取new_val的值
		slVote, err := json.Marshal(mValue[_NEW_VAL])
		if err != nil {
			beegoLog.Error(err.Error())
			continue
		}
		if bytes.Equal(slVote, []byte("null")) {
			beegoLog.Error("is null")
			continue
		}

		beegoLog.Debug("1.2 ceChangefeed --->")
		gchInput <- string(slVote)
		votes_changefeed_time.Send("votes_changefeed")
	}
	beegoLog.Error("--------------------------------------------------------")
	beegoLog.Error("changfeed exit")
	beegoLog.Error("--------------------------------------------------------")
}

//---------------------------------------------------------------------------
func ceHeadFilter() error {
	beegoLog.Debug("2.进入ceHeadFilter")
	defer close(gchOutput)
	for {
		// 读取new_val的值
		contract_election_time := monitor.Monitor.NewTiming()
		readData, ok := <-gchInput
		if !ok {
			break
		}

		// 将new_val转化为vote
		vote := model.Vote{}
		err := json.Unmarshal([]byte(readData), &vote)
		if err != nil {
			beegoLog.Error(err.Error())
			continue
		}

		// 验证是否为头节点
		beegoLog.Debug("2.2 get publickey and verify")
		beegoLog.Debug("2.2.1 get publickey")
		mainPubkey, err := rethinkdb.GetContractMainPubkeyByContract(vote.VoteBody.VoteFor)
		if err == nil {
			beegoLog.Debug("2.2.2 verify head node")
			if isHead, _ := _verifyHeadNode(mainPubkey); isHead {
				beegoLog.Debug("2.2.3 verify vote")
				vote_validate_time := monitor.Monitor.NewTiming()
				pContractOutput, pass, err := _verifyVotes(vote.VoteBody.VoteFor)
				vote_validate_time.Send("vote_validate")
				if err != nil {
					if !pass {
						beegoLog.Error(err.Error())
					} else {
						// vote not enough
						beegoLog.Info(err.Error())
					}
					continue
				}

				if pass { // 合约合法
					beegoLog.Debug("2.3 ceHeadFilter --->")
					ceQueryEists(*pContractOutput)
					contract_election_time.Send("contract_election")
				} else { // 合约不合法
					beegoLog.Debug("2.3 contract invalid and insert consensusFailure")
					var consensusFailure model.ConsensusFailure
					consensusFailure.Id = common.GenerateUUID()
					consensusFailure.ConsensusType = _CONSENSUS_TYPE
					consensusFailure.ConsensusId = vote.VoteBody.VoteFor
					consensusFailure.ConsensusReason = _CONSENSUS_REASON
					consensusFailure.Timestamp = common.GenTimestamp()

					slConsensusFailure, err := json.Marshal(consensusFailure)
					if err != nil {
						beegoLog.Error(err.Error())
						continue
					}
					if !rethinkdb.InsertConsensusFailure(string(slConsensusFailure)) {
						beegoLog.Error(err.Error())
					}
					//consensus_failure_count, err := rethinkdb.GetConsensusFailuresCount()
					//if err != nil {
					//	beegoLog.Error(err.Error())
					//	continue
					//}
					//monitor.Monitor.Gauge("consensus_failure", consensus_failure_count)
					monitor.Monitor.Count("consensus_failure", 1)
				}
			}
		} else {
			beegoLog.Error(err.Error())
		}
	}
	return nil
}

//---------------------------------------------------------------------------
func ceQueryEists(contractOutput model.ContractOutput) {
	beegoLog.Debug("3.进入ceQueryEists")

	beegoLog.Debug("3.2 query contractoutput table")
	output, err := rethinkdb.GetContractOutputById(contractOutput.Id)
	if err != nil {
		beegoLog.Error(err.Error())
		return
	}

	slContractOutput, _ := json.Marshal(contractOutput)

	if len(output) == 0 {
		beegoLog.Debug("3.3 insert contractoutput table")
		rethinkdb.InsertContractOutput(string(slContractOutput))
	} else {
		beegoLog.Debug("3.3 contractoutput exist")
	}
	beegoLog.Debug("3.4 ceQueryEists ---> /dev/null")
	gchOutput <- string(slContractOutput)
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

	beegoLog.Debug("2.2.3.1 verify vote signature")
	eligible_votes := make(map[string]model.Vote)
	for _, tmpVote := range slVote {
		if tmpVote.VerifyVoteSignature() {
			if isExist, _ := _verifyPublicKey(tmpVote.NodePubkey); isExist {
				eligible_votes[tmpVote.NodePubkey] = tmpVote
			}
		}
	}

	beegoLog.Debug("2.2.3.2 count votes")
	if len(eligible_votes)*2 < gnPublicKeysNum { // vote没有达到节点数的一半时
		return nil, true, fmt.Errorf("vote not enough")
	}

	beegoLog.Debug("2.2.3.3 valid votes")
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

	beegoLog.Debug("contractOutput : %+v", common.StructSerialize(contractOutput))
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
		log.Println("===========================================")
		log.Printf("Panic !!!!, err is %+v", err)
		if outputStack {
			var slStack [4096]byte
			nReadNum := runtime.Stack(slStack[:], true)
			log.Printf(string(slStack[:nReadNum]))
		}
		log.Println("===========================================")
	}
}

//---------------------------------------------------------------------------
