package pipelines

import (
	"bytes"
	"encoding/json"
	"os"

	"unicontract/src/common"
	"unicontract/src/common/monitor"
	"unicontract/src/config"
	r "unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"

	"github.com/astaxie/beego/logs"
	"time"
)

const (
	_CVTHREADNUM = 10
)

var (
	cvPool   *ThreadPool
	cvInput  chan string
	cvOutput chan string
)

func cvChangefeed() {
	var value interface{}
	res := r.Changefeed(r.DBNAME, r.TABLE_CONTRACTS)
	for res.Next(&value) {
		contrant_changefeed_time := monitor.Monitor.NewTiming()
		m := value.(map[string]interface{})
		v, err := json.Marshal(m["new_val"])
		if err != nil {
			logs.Error(err.Error())
			continue
		}
		if bytes.Equal(v, []byte("null")) {
			continue
		}
		logs.Debug("-------cvChangefeed:", common.Serialize(m))
		cvInput <- string(v)

		contrant_changefeed_time.Send("contrant_changefeed")
	}
}

func cvValidateContract() error {
	defer close(cvOutput)
	for {
		contract_vote_time := monitor.Monitor.NewTiming()
		t, ok := <-cvInput
		if !ok {
			break
		}
		mod := model.ContractModel{}
		err := json.Unmarshal([]byte(t), &mod)
		if err != nil {
			logs.Error(err.Error())
			continue
		}
		v := model.Vote{}
		contract_validate_time := monitor.Monitor.NewTiming()
		if mod.Validate() {
			//vote true
			v.VoteBody.IsValid = true
		} else {
			//vote flase
			v.VoteBody.IsValid = false
		}
		contract_validate_time.Send("contract_validate")
		v.VoteBody.VoteFor = mod.Id
		logs.Debug("-------cvValidateContract:", common.Serialize(v))

		v.NodePubkey = config.Config.Keypair.PublicKey
		v.VoteBody.Timestamp = common.GenTimestamp()
		v.VoteBody.VoteType = "Contract"
		v.Id = v.GenerateId()
		v.Signature = v.SignVote()

		logs.Debug("-------cvWriteVote:", common.Serialize(v))
		vote_write_time := monitor.Monitor.NewTiming()
		res := r.Insert("Unicontract", "Votes", v.ToString())
		vote_write_time.Send("vote_write")

		cvOutput <- common.Serialize(res)

		contract_vote_time.Send("contract_vote")
	}
	return nil
}

func startContractVote() {
	defer cvPool.Stop()
	defer close(cvInput)

	cvInput = make(chan string, _INPUT_LEN)
	cvOutput = make(chan string, _OUTPUT_LEN)

	cvPool = new(ThreadPool)
	cvPool.Init(_CVTHREADNUM)

	for i := 0; i < _CVTHREADNUM; i++ {
		cvPool.AddTask(func() error {
			return cvValidateContract()
		})
	}

	go cvPool.Start()
	go cvChangefeed()

	for {
		if out, ok := <-cvOutput; ok {
			f, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
			if err != nil {
				logs.Error(err.Error())
			}
			f.Write([]byte(out))
		} else {
			break
		}
	}
}

//func cvVote(in io.Reader, out io.Writer) {
//
//	rd := bufio.NewReader(in)
//	p := make([]byte, MaxSizeTX)
//	for {
//		n, _ := rd.Read(p)
//		time := monitor.Monitor.NewTiming()
//		if n == 0 {
//			continue
//		}
//		t := p[:n]
//		v :=model.Vote{}
//		err := json.Unmarshal(t,&v)
//		if err != nil {
//			logs.Error(err.Error())
//			continue
//		}
//		v.NodePubkey = config.Config.Keypair.PublicKey
//		v.VoteBody.Timestamp = common.GenTimestamp()
//		v.VoteBody.VoteType = "Contract"
//		v.Id = v.GenerateId()
//		v.Signature = v.SignVote()
//
//		logs.Debug("-------cvVote:",common.Serialize(v))
//		out.Write([]byte(v.ToString()))
//
//		time.Send("cv_validate_contract")
//	}
//}
//
//func cvWriteVote(in io.Reader, out io.Writer) {
//
//	rd := bufio.NewReader(in)
//	p := make([]byte, MaxSizeTX)
//	for {
//		n, _ := rd.Read(p)
//		time := monitor.Monitor.NewTiming()
//		if n == 0 {
//			continue
//		}
//		t := p[:n]
//		res :=r.Insert("Unicontract", "Votes", string(t))
//
//		logs.Debug("-------cvWriteVote:",common.Serialize(res))
//		out.Write(nil)
//
//		time.Send("cv_vote_contract")
//	}
//}
