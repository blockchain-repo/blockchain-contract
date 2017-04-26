package pipelines

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
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
	gslPublicKeys      []string
	gnPublicKeysNum    int
	gstrPublicKey      string
	gmContractOutputId map[string]string
)

//---------------------------------------------------------------------------
func init() {
	gslPublicKeys = config.GetAllPublicKey()
	gnPublicKeysNum = len(gslPublicKeys)
	gstrPublicKey = config.Config.Keypair.PublicKey
	gmContractOutputId = make(map[string]string)
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

		strOutputId := gmContractOutputId[string(slRealData)]
		delete(gmContractOutputId, string(slRealData))

		beegoLog.Debug("3.2 query contractoutput table")
		output, err := rethinkdb.GetContractOutputById(strOutputId)
		if err != nil {
			beegoLog.Error(err.Error())
			continue
		}

		if len(output) == 0 {
			beegoLog.Debug("3.3 insert contractoutput table")
			rethinkdb.InsertContractOutput(string(slRealData))
		} else {
			beegoLog.Debug("3.3 contractoutput exist")
		}
		beegoLog.Debug("3.4 ceQueryEists ---> /dev/null")
		out.Write(slRealData)
		time.Send("ce_query_contract")
	}
}

//---------------------------------------------------------------------------
func startContractElection() {
	p := Pipe(
		ceChangefeed,
		ceHeadFilter,
		ceQueryEists)

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

	beegoLog.Debug("2.2.3.1 verify vote signature")
	eligible_votes := make(map[string]model.Vote)
	for _, tmpVote := range slVote {
		//do not forget fix it !!!!!
		if /*true*/ tmpVote.VerifyVoteSignature() {
			if isExist, _ := _verifyPublicKey(tmpVote.NodePubkey); isExist {
				eligible_votes[tmpVote.NodePubkey] = tmpVote
			}
		}
	}

	beegoLog.Debug("2.2.3.2 count votes")
	// do not forget fix debug!!!!
	if len(eligible_votes)*2 < gnPublicKeysNum { // vote没有达到节点数的一半时
		log.Println("vote not enough")
		return nil, true, errors.New("vote not enough")
	}

	beegoLog.Debug("2.2.3.3 valid votes")
	bValid := _verifyValid(eligible_votes)

	if bValid {
		contractOutput, err := _produceContractOutput(contractId, slVote)
		if err != nil {
			return nil, bValid, err
		}
		slMyContract, err := json.Marshal(contractOutput)
		gmContractOutputId[string(slMyContract)] = contractOutput.Id
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
	contractOutput.Id = common.HashData(common.Serialize(contractOutput))

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
