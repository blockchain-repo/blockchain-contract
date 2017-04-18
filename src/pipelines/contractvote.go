package pipelines

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"

	"unicontract/src/core/model"
	r "unicontract/src/core/db/rethinkdb"
)

func cvChangefeed(in io.Reader, out io.Writer) {
	var value interface{}
	//TODO table name
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
		//TODO validate
		mod := model.ContractModel{}
		err := json.Unmarshal(t,&mod)
		if err != nil {
			log.Fatalf(err.Error())
			continue
		}
		mod.Validate()
		out.Write(t)
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
		//TODO make vote
		t := p[:n]
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
		//TODO write vote
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
