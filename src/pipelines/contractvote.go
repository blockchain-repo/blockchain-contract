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
	r "unicontract/src/core/db/rethinkdb"
)

func cvChangefeed(in io.Reader, out io.Writer) {
	var value interface{}
	res := r.Changefeed("Unicontract", "Contract")
	for res.Next(&value) {
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
	}
}

func cvValidateContract(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {
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
		v := model.Votes{}
		if mod.Validate() {
			//vote true
			v.Vote.IsValid = true
		} else {
			//vote flase
			v.Vote.IsValid = false
		}
		v.Vote.VoteForContract = mod.Id
		out.Write([]byte(v.ToString()))
	}
}

func cvVote(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {
		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		v :=model.Votes{}
		err := json.Unmarshal(t,&v)
		if err != nil {
			log.Fatalf(err.Error())
			continue
		}
		//TODO make vote(NodePubkey)
		v.NodePubkey = "EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet"
		v.Signature = v.SignVote(v)
		v.Vote.Timestamp = common.GenTimestamp()
		out.Write(t)
	}
}

func cvWriteVote(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {
		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		r.Insert("Unicontract", "Votes", string(t))
		out.Write(t)
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
