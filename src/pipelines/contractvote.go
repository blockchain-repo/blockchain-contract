package pipelines

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"

	"unicontract/src/core/model"
	"unicontract/src/common"
	"unicontract/src/config"
	r "unicontract/src/core/db/rethinkdb"
	"unicontract/src/common/monitor"
)

func cvChangefeed(in io.Reader, out io.Writer) {
	var value interface{}
	res := r.Changefeed("Unicontract", "Contracts")
	for res.Next(&value) {

		time := monitor.Monitor.NewTiming()

		m := value.(map[string]interface{})
		v, err := json.Marshal(m["new_val"])
		if err != nil {
			log.Fatalf(err.Error())
			continue
		}
		if bytes.Equal(v, []byte("null")) {
			continue
		}
		out.Write(v)

		time.Send("contrant_changefeed")
	}
}

func cvValidateContract(in io.Reader, out io.Writer) {

	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {

		time := monitor.Monitor.NewTiming()

		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		mod := model.ContractModel{}
		err := json.Unmarshal(t,&mod)
		if err != nil {
			log.Fatalf(err.Error())
			continue
		}
		v := model.Vote{}
		if mod.Validate() {
			//vote true
			v.VoteBody.IsValid = true
		} else {
			//vote flase
			v.VoteBody.IsValid = false
		}
		v.VoteBody.VoteForContract = mod.Id
		out.Write([]byte(v.ToString()))

		time.Send("cv_validate_contract")
	}
}

func cvVote(in io.Reader, out io.Writer) {

	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {

		time := monitor.Monitor.NewTiming()

		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		v :=model.Vote{}
		err := json.Unmarshal(t,&v)
		if err != nil {
			log.Fatalf(err.Error())
			continue
		}
		v.NodePubkey = config.Config.Keypair.PublicKey
		v.Id = v.GenerateId()
		v.Signature = v.SignVote()
		v.VoteBody.Timestamp = common.GenTimestamp()
		v.VoteBody.VoteType = "node"
		out.Write([]byte(v.ToString()))

		time.Send("cv_validate_contract")
	}
}

func cvWriteVote(in io.Reader, out io.Writer) {

	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {
		time := monitor.Monitor.NewTiming()

		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		r.Insert("Unicontract", "Votes", string(t))
		out.Write(t)

		time.Send("cv_vote_contract")
	}
}

func startContractVote() {

	p := Pipe(
		cvChangefeed,
		cvValidateContract,
		cvVote,
		cvWriteVote)

	f, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
	if err != nil {
		log.Fatalf(err.Error())
	}
	w := bufio.NewWriter(f)
	p(nil, w)
}
