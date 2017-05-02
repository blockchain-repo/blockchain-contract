package pipelines

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
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
	_NewVal          = "new_val"
	_HTTPOK          = 200
	_VERSION         = 2
	_OPERATION       = "CONTRACT"
	_CONSENSUSTYPE   = "contract"
	_CONSENSUSREASON = "contract illegal"
	_THREADNUM       = 10
	_INPUTLEN        = 20
	_OUTPUTLEN       = 20
	_Fulfillment     = "cf:4:RtTtCxNf1Bq7MFeIToEosMAa3v_jKtZUtqiWAXyFz1ejPMv-t7vT6DANcrYvKFHAsZblmZ1Xk03HQdJbGiMyb5CmQqGPHwlgKusNu9N_IDtPn7y16veJ1RBrUP-up4YD"
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
func init() {
	gslPublicKeys = config.GetAllPublicKey()
	gnPublicKeysNum = len(gslPublicKeys)
	gstrPublicKey = config.Config.Keypair.PublicKey

	gchInput = make(chan string, _INPUTLEN)
	gchOutput = make(chan string, _OUTPUTLEN)

	gPool = new(ThreadPool)
}

//---------------------------------------------------------------------------
func ceChangefeed() {
	beegoLog.Debug("1.进入ceChangefeed")
	var value interface{}
	res := rethinkdb.Changefeed(rethinkdb.DBNAME, rethinkdb.TABLE_VOTES)
	for res.Next(&value) {
		time := monitor.Monitor.NewTiming()
		beegoLog.Debug("1.1 ceChangefeed get new_val")
		mValue := value.(map[string]interface{})
		// 提取new_val的值
		slVote, err := json.Marshal(mValue[_NewVal])
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
		time.Send("votes_changefeed")
	}
}

//---------------------------------------------------------------------------
func ceHeadFilter() error {
	beegoLog.Debug("2.进入ceHeadFilter")
	defer close(gchOutput)
	for {
		// 读取new_val的值
		time := monitor.Monitor.NewTiming()
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
				pContractOutput, pass, err := _verifyVotes(vote.VoteBody.VoteFor)
				if err != nil {
					if !pass {
						beegoLog.Error(err.Error())
					} else {
						// vote not enough
						beegoLog.Debug(err.Error())
					}
					continue
				}

				if pass { // 合约合法
					beegoLog.Debug("2.3 ceHeadFilter --->")
					ceQueryEists(*pContractOutput)
					time.Send("ce_validate_head")
				} else { // 合约不合法
					beegoLog.Debug("2.3 contract invalid and insert consensusFailure")
					var consensusFailure model.ConsensusFailure
					consensusFailure.Id = common.GenerateUUID()
					consensusFailure.ConsensusType = _CONSENSUSTYPE
					consensusFailure.ConsensusId = vote.VoteBody.VoteFor
					consensusFailure.ConsensusReason = _CONSENSUSREASON
					consensusFailure.Timestamp = common.GenTimestamp()

					slConsensusFailure, err := json.Marshal(consensusFailure)
					if err != nil {
						beegoLog.Error(err.Error())
						continue
					}
					if !rethinkdb.InsertConsensusFailure(string(slConsensusFailure)) {
						beegoLog.Error(err.Error())
					}
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
	time := monitor.Monitor.NewTiming()

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
	time.Send("ce_query_contract")
}

//---------------------------------------------------------------------------
func startContractElection() {
	defer gPool.Stop()
	defer close(gchInput)

	gPool.Init(_THREADNUM)
	for i := 0; i < _THREADNUM; i++ {
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
	bValid := _verifyValid(eligible_votes)

	if bValid {
		contractOutput, err := _produceContractOutput(contractId, slVote)
		if err != nil {
			return nil, bValid, err
		}
		return &contractOutput, bValid, err
	} else {
		return nil, bValid, nil
	}
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
func _verifyValid(mVotes map[string]model.Vote) bool {
	var nValidNum, nInValidNum int
	for _, value := range mVotes {
		if value.VoteBody.IsValid {
			nValidNum++
		} else {
			nInValidNum++
		}
	}

	bValid := false
	nTotalVoteNum := len(mVotes)
	if nValidNum*2 > nTotalVoteNum {
		bValid = true
	}
	if nInValidNum*2 > nTotalVoteNum {
		bValid = false
	}

	return bValid
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
	contractOutput.Transaction.Relation.ContractId = contractModel.Id
	for _, value := range gslPublicKeys {
		contractOutput.Transaction.Relation.Voters =
			append(contractOutput.Transaction.Relation.Voters, value)
	}

	beegoLog.Debug("contractOutput : %+v", contractOutput)

	contractOutput.Id = common.HashData(common.Serialize(contractOutput))

	fulfillment := new(model.Fulfillment)
	fulfillment.Fulfillment = _Fulfillment
	contractOutput.Transaction.Fulfillments =
		append(contractOutput.Transaction.Fulfillments, fulfillment)

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
