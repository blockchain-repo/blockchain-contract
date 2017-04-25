package pipelines

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"os"
	"fmt"
	"errors"

	r "unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
	"unicontract/src/chain"
	"unicontract/src/config"
	"unicontract/src/common/monitor"

	"github.com/astaxie/beego/logs"
)

func txeChangefeed(in io.Reader, out io.Writer) {
	var value interface{}
	res := r.Changefeed("Unicontract", "ContractOutputs")
	fmt.Printf("changefeed result : %s",res)
	for res.Next(&value) {
		time := monitor.Monitor.NewTiming()

		m := value.(map[string]interface{})
		v, err := json.Marshal(m["new_val"])
		if err != nil {
			logs.Error(err.Error())
			continue
		}
		if bytes.Equal(v, []byte("null")) {
			continue
		}
		fmt.Printf("change result : %s",v)
		out.Write(v)

		time.Send("contractOutputs_changefeed")
	}
}

func txeHeadFilter(in io.Reader, out io.Writer){

	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {
		time := monitor.Monitor.NewTiming()

		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		conout := model.ContractOutput{}
		err := json.Unmarshal(t,&conout)
		if err != nil {
			logs.Error(err.Error())
			continue
		}
		//main node filter
		mainNodeKey  := conout.Transaction.ContractModel.ContractHead.MainPubkey
		myNodeKey := config.Config.Keypair.PublicKey
		if mainNodeKey != myNodeKey {
			logs.Info("I am not the mainnode of the C-output %s",conout.Id)
			continue
		}
		out.Write(t)

		time.Send("txe_validate_head")
	}
}

func txeValidate(in io.Reader, out io.Writer) {

	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {
		time := monitor.Monitor.NewTiming()

		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		coModel := model.ContractOutput{}
		err := json.Unmarshal(t,&coModel)
		if err != nil {
			logs.Error(err.Error())
			continue
		}

		if !coModel.HasEnoughVotes() {
			//not enough votes
			continue
		}
		if !coModel.ValidateHash() {
			//invalid hash
			logs.Error(errors.New("invalid hash"))
			continue
		}
		if !coModel.ValidateContractOutput() {
			//invalid signature
			logs.Error(errors.New("invalid signature"))
			continue
		}
		out.Write(t)

		time.Send("txe_contractOutput_validate")
	}
}

func txeQueryEists(in io.Reader, out io.Writer) {

	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {
		time := monitor.Monitor.NewTiming()

		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		coModel := model.ContractOutput{}
		err := json.Unmarshal(t,&coModel)
		if err != nil {
			logs.Error(err.Error())
			continue
		}
		//check whether already exist
		id := coModel.Id
		result,err := chain.GetContractTx("{'id':"+id+"}")
		if err != nil{
			logs.Error(err.Error())
		}

		if result.Code != 200 {
			errors.New("request send failed")
		}

		fmt.Print(result.Data)
		//if the unichain already has the contractoutput ,do nothing
		if result.Data == nil {
			continue
		}
		out.Write(t)

		time.Send("txe_query_contractOutput")
	}
}

func txeSend(in io.Reader, out io.Writer) {

	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {
		time := monitor.Monitor.NewTiming()

		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		//write the contractoutput to unichain.
		result,err:= chain.CreateContractTx(t)
		if err != nil{
			logs.Error(err.Error())
		}
		fmt.Print(result.Data)
		if result.Code != 200 {
			errors.New("request send failed")
			SaveOutputErrorData(_TableNameSendFailingRecords,t)
		}
		out.Write(t)

		time.Send("txe_send_contractOutput")
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
		logs.Error(err.Error())
	}
	w := bufio.NewWriter(f)
	p(nil, w)
}
