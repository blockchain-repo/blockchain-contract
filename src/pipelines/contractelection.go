package pipelines

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
)

import (
	"unicontract/src/chain"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
)

//---------------------------------------------------------------------------
type _MyContract struct {
	model.ContractModel `json:'contract'`
	SLVotes             []model.Vote `json:'votes'`
}

type _GetContractInput struct {
	Id string `json:'id'`
}

//---------------------------------------------------------------------------
const (
	_DBName         = "Unicontract"
	_TableNameVotes = "vote"
	_NewVal         = "new_val"
	_HTTPOK         = 200
)

//---------------------------------------------------------------------------
func ceChangefeed(in io.Reader, out io.Writer) {
	var value interface{}
	res := rethinkdb.Changefeed(_DBName, _TableNameVotes)
	for res.Next(&value) {
		mValue := value.(map[string]interface{})
		// 提取new_val的值
		slVote, err := json.Marshal(mValue[_NewVal])
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		if bytes.Equal(slVote, []byte("null")) {
			log.Fatalln("is null")
			continue
		}
		out.Write(slVote)
	}
}

//---------------------------------------------------------------------------
func ceHeadFilter(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	slVote := make([]byte, MaxSizeTX)
	for {
		// 读取new_val的值
		nReadNum, err := rd.Read(slVote)
		if err != nil {
			log.Fatalln(err.Error())
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
			log.Fatalln(err.Error())
			continue
		}

		// 验证是否为头节点
		mainPubkey, err := rethinkdb.GetContractMainPubkeyById(
			vote.VoteBody.VoteForContract)
		if err == nil {
			isHead, err := _verifyHeadNode(mainPubkey)
			if err == nil {
				if isHead {
					slMyContract, pass, err := _verifyVotes(mainPubkey,
						vote.VoteBody.VoteForContract)
					if err != nil {
						log.Fatalln(err.Error())
						continue
					}

					if pass {
						out.Write(slMyContract)
					} else {
						//TODO InValid情况
					}
				}
			} else {
				log.Fatalln(err.Error())
				continue
			}
		} else {
			log.Fatalln(err.Error())
			continue
		}
	}
}

//---------------------------------------------------------------------------
func ceQueryEists(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	slMyContract := make([]byte, MaxSizeTX)
	for {
		nReadNum, err := rd.Read(slMyContract)
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		if nReadNum == 0 {
			continue
		}
		slReadData := slMyContract[:nReadNum]

		myContract := _MyContract{}
		err = json.Unmarshal(slReadData, &myContract)
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}

		input := _GetContractInput{Id: myContract.Id}
		slInput, err := json.Marshal(input)
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}

		responseResult := chain.GetContract(string(slInput))
		if responseResult.Code == _HTTPOK {
			if responseResult.Data == nil {
				out.Write(slReadData)
			}
		}
	}
}

//---------------------------------------------------------------------------
func ceSend(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	slReadData := make([]byte, MaxSizeTX)
	for {
		nReadNum, err := rd.Read(slReadData)
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		if nReadNum == 0 {
			continue
		}
		slRealData := slReadData[:nReadNum]

		responseResult := chain.CreateContract(slRealData)
		if responseResult.Code != _HTTPOK {
			log.Fatalln(responseResult.Message)
		}

		out.Write(slRealData)
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
		log.Fatalf(err.Error())
	}
	w := bufio.NewWriter(f)
	p(nil, w)
}

//---------------------------------------------------------------------------

//---------------------------------------------------------------------------
// 私有工具函数
//---------------------------------------------------------------------------
func _verifyHeadNode(mainPubkey string) (bool, error) {
	//TODO 验证是否是头节点
	return true, nil
}

//---------------------------------------------------------------------------
func _verifyVotes(mainPubkey, contractId string) ([]byte, bool, error) {
	// 查询所有的vote
	var myContract _MyContract
	strVotes, err := rethinkdb.GetVotesByContractId(mainPubkey)
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
		if tmpVote.VerifyVoteSignature() {
			if isExist, _ := _verifyPublicKey(tmpVote.NodePubkey); isExist {
				eligible_votes[tmpVote.NodePubkey] = tmpVote
			}
		}
	}

	// 统计vote并判断valid
	bValid := _verifyValid(eligible_votes)

	if bValid {
		strContract, err := rethinkdb.GetContractById(contractId)
		err = json.Unmarshal([]byte(strContract), &myContract.ContractModel)
		if err != nil {
			return nil, false, err
		}
		copy(myContract.SLVotes, slVote)
	} else {
		//TODO InValid情况
	}

	slMyContract, err := json.Marshal(myContract)
	if err != nil {
		return nil, bValid, err
	}

	return slMyContract, bValid, nil
}

//---------------------------------------------------------------------------
func _verifyPublicKey(NodePubkey string) (bool, error) {
	//TODO 验证公钥是否在公钥环
	return true, nil
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
