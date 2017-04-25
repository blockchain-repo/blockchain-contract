package pipelines

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

import (
	beegoLog "github.com/astaxie/beego/logs"
)

import (
	"unicontract/src/chain"
	"unicontract/src/common"
	"unicontract/src/common/monitor"
	"unicontract/src/config"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
)

//---------------------------------------------------------------------------
const (
	_DBName          = "Unicontract"
	_TableNameVotes  = "Votes"
	_NewVal          = "new_val"
	_HTTPOK          = 200
	_VERSION         = 2
	_OPERATION       = "CONTRACT"
	_CONSENSUSTYPE   = "contract"
	_CONSENSUSREASON = "contract illegal"
)

var (
	gslPublicKeys   []string
	gnPublicKeysNum int
	gstrPublicKey   string
)

//---------------------------------------------------------------------------
func init() {
	gslPublicKeys = config.GetAllPublicKey()
	gnPublicKeysNum = len(gslPublicKeys)
	gstrPublicKey = config.Config.Keypair.PublicKey
}

//---------------------------------------------------------------------------
func ceChangefeed(in io.Reader, out io.Writer) {
	beegoLog.Debug("1.进入ceChangefeed")
	var value interface{}
	res := rethinkdb.Changefeed(_DBName, _TableNameVotes)
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
		out.Write(slVote)
		time.Send("votes_changefeed")
	}
}

//---------------------------------------------------------------------------
func ceHeadFilter(in io.Reader, out io.Writer) {
	beegoLog.Debug("2.进入ceHeadFilter")
	rd := bufio.NewReader(in)
	slVote := make([]byte, MaxSizeTX)
	for {
		// 读取new_val的值
		nReadNum, err := rd.Read(slVote)
		time := monitor.Monitor.NewTiming()
		beegoLog.Debug("2.1 ceHeadFilter get data")
		if err != nil {
			beegoLog.Error(err.Error())
			continue
		}
		if nReadNum == 0 {
			continue
		}
		slReadData := slVote[:nReadNum]

		// 将new_val转化为vote
		vote := model.Vote{}
		err = json.Unmarshal(slReadData, &vote)
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
				slMyContract, pass, err := _verifyVotes(vote.VoteBody.VoteFor)
				if err != nil {
					if !pass { // 只有产生错误时才记录日志，当vote数量不够节点数量的一半时直接进入下次等待，不记录错误
						beegoLog.Error(err.Error())
					}
					continue
				}

				if pass { // 合约合法
					beegoLog.Debug("2.3 ceHeadFilter --->")
					out.Write(slMyContract)
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

		time.Send("ce_validate_head")
	}
}

//---------------------------------------------------------------------------
func ceQueryEists(in io.Reader, out io.Writer) {
	beegoLog.Debug("3.进入ceQueryEists")
	rd := bufio.NewReader(in)
	slContractOutput := make([]byte, MaxSizeTX)
	for {
		nReadNum, err := rd.Read(slContractOutput)
		time := monitor.Monitor.NewTiming()
		beegoLog.Debug("3.1 ceQueryEists get data")
		if err != nil {
			beegoLog.Error(err.Error())
			continue
		}
		if nReadNum == 0 {
			continue
		}
		slRealData := slContractOutput[:nReadNum]

		myContract := model.ContractOutput{}
		err = json.Unmarshal(slRealData, &myContract)
		if err != nil {
			beegoLog.Error(err.Error())
			continue
		}

		beegoLog.Debug("3.2 query chain")
		strInput := fmt.Sprintf("{\"id\":\"%s\"}", myContract.Transaction.ContractModel.Id)
		responseResult, err := chain.GetContract(strInput)
		if err != nil {
			beegoLog.Error(err.Error())
			beegoLog.Debug("3.3 query failed and SaveOutputErrorData")
			SaveOutputErrorData(_TableNameSendFailingRecords, slRealData)
		} else {
			if responseResult.Code == _HTTPOK {
				if responseResult.Data == nil { // not exist
					beegoLog.Debug("3.3 ceQueryEists --->")
					out.Write(slRealData)
				} else { // exist
					beegoLog.Debug("3.3 exist")
				}
			}
		}

		time.Send("ce_query_contract")
	}
}

//---------------------------------------------------------------------------
func ceSend(in io.Reader, out io.Writer) {
	beegoLog.Debug("4.进入ceSend")
	rd := bufio.NewReader(in)
	slContractOutput := make([]byte, MaxSizeTX)
	for {
		nReadNum, err := rd.Read(slContractOutput)
		time := monitor.Monitor.NewTiming()
		beegoLog.Debug("4.1 ceSend get data")
		if err != nil {
			beegoLog.Error(err.Error())
			continue
		}
		if nReadNum == 0 {
			continue
		}
		slRealData := slContractOutput[:nReadNum]

		beegoLog.Debug("4.2 CreateContract")
		responseResult, err := chain.CreateContract(slRealData)
		if err != nil {
			beegoLog.Error(err.Error())
			beegoLog.Debug("4.3 CreateContract failed and SaveOutputErrorData")
			SaveOutputErrorData(_TableNameSendFailingRecords, slRealData)
			continue
		}
		if responseResult.Code != _HTTPOK {
			beegoLog.Error(responseResult.Message)
			beegoLog.Debug("4.3 CreateContract failed and SaveOutputErrorData")
			SaveOutputErrorData(_TableNameSendFailingRecords, slRealData)
		}

		beegoLog.Debug("4.4 ceSend --->")
		out.Write(slRealData)
		time.Send("ce_send_contract")
	}
}

//---------------------------------------------------------------------------
func startContractElection() {
	p := Pipe(
		ceChangefeed,
		ceHeadFilter,
		ceQueryEists,
		ceSend)

	f, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
	if err != nil {
		beegoLog.Error(err.Error())
	}
	w := bufio.NewWriter(f)
	p(nil, w)
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
func _verifyVotes(contractId string) ([]byte, bool, error) {
	// 查询所有的vote
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
		return nil, false, errors.New("no vote")
	}

	// 验证vote
	eligible_votes := make(map[string]model.Vote)
	for _, tmpVote := range slVote {
		//do not forget fix it !!!!!
		if true /*tmpVote.VerifyVoteSignature()*/ {
			if isExist, _ := _verifyPublicKey(tmpVote.NodePubkey); isExist {
				eligible_votes[tmpVote.NodePubkey] = tmpVote
			}
		}
	}

	// 统计vote并判断valid
	// do not forget fix debug!!!!
	//	if len(eligible_votes)*2 < gnPublicKeysNum { // vote没有达到节点数的一半时
	//		log.Println("vote not enough")
	//		return nil, true, errors.New("vote not enough")
	//	}
	bValid := _verifyValid(eligible_votes)

	if bValid {
		contractOutput, err := _produceContractOutput(contractId, slVote)
		if err != nil {
			return nil, bValid, err
		}
		slMyContract, err := json.Marshal(contractOutput)
		return slMyContract, bValid, err
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

	err = json.Unmarshal([]byte(strContract),
		&contractOutput.Transaction.ContractModel)
	if err != nil {
		return contractOutput, err
	}

	contractOutput.Transaction.Relaction = new(model.Relaction)
	for key, value := range slVote {
		contractOutput.Transaction.Relaction.Voters =
			append(contractOutput.Transaction.Relaction.Voters, value.NodePubkey)
		contractOutput.Transaction.Relaction.Votes =
			append(contractOutput.Transaction.Relaction.Votes, &slVote[key])
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
