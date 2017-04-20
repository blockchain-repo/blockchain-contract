package pipelines

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"fmt"
	"errors"

	r "unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
	"unicontract/src/chain"
)

func txeChangefeed(in io.Reader, out io.Writer) {
	var value interface{}
	res := r.Changefeed("Unicontract", "ContractOutputs")
	fmt.Printf("changefeed result : %s",res)
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
		fmt.Printf("change result : %s",v)
		out.Write(v)
	}
}

func txeHeadFilter(in io.Reader, out io.Writer){
	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {
		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		conout := model.ContractOutput{}
		err := json.Unmarshal(t,&conout)
		if err != nil {
			log.Fatalf(err.Error())
			continue
		}
		//main node filter
		mainNodeKey  := conout.Transaction.Contracts.MainPubkey
		//TODO get node pubkey
		myNodeKey := ""
		if mainNodeKey != myNodeKey {
			continue
		}
		out.Write(t)
	}
}

func txeValidate(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {
		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		coModel := model.ContractOutput{}
		err := json.Unmarshal(t,&coModel)
		if err != nil {
			log.Fatalf(err.Error())
			continue
		}

		if !coModel.HasEnoughVotes() {
			//not enough votes
			continue
		}
		if !coModel.ValidateHash() {
			//invalid hash
			log.Fatal(errors.New("invalid hash"))
			continue
		}
		if !coModel.ValidateContractOutput() {
			//invalid signature
			log.Fatal(errors.New("invalid signature"))
			continue
		}
		out.Write(t)
	}
}

func txeQueryEists(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {
		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		coModel := model.ContractOutput{}
		err := json.Unmarshal(t,&coModel)
		if err != nil {
			log.Fatalf(err.Error())
			continue
		}
		//check whether already exist
		id := coModel.Id
		result := chain.GetContractTx("{'id':"+id+"}")

		if result.Code != 200 {
			//TODO error handling
			errors.New("")
		}

		fmt.Print(result.Data)
		//if the unichain already has the contractoutput ,do nothing
		if result.Data == nil {
			continue
		}
		out.Write(t)
	}
}

func txeSend(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {
		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		//write the contractoutput to unichain.
		result := chain.CreateContractTx(t)
		fmt.Print(result.Data)
		if result.Code != 200 {
			//TODO error handling
			errors.New("")
		}
		out.Write(t)
	}
}

func starttxElection() {

	p := Pipe(
		txeChangefeed,
		txeHeadFilter,
		txeValidate,
		txeQueryEists,
		txeSend)

	f, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
	if err != nil {
		log.Fatalf(err.Error())
	}
	w := bufio.NewWriter(f)
	p(nil, w)
}
