package pipelines

import (
	"bytes"
	"encoding/json"

	"unicontract/src/common/monitor"
	"unicontract/src/config"
	r "unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"

	"bufio"
	"errors"
	"github.com/astaxie/beego/logs"
	"io"
	"unicontract/src/common"
	"time"
)

var (
	txPool   *ThreadPool
	txInput  chan string
	txOutput chan string
)

func txChangefeed() {
	var value interface{}
	res := r.Changefeed("Unicontract", "ContractOutputs")
	for res.Next(&value) {
		time := monitor.Monitor.NewTiming()

		logs.Info(" txElection step1 : txeChangefeed ")
		m := value.(map[string]interface{})
		v, err := json.Marshal(m["new_val"])
		if err != nil {
			logs.Error(err.Error())
			continue
		}
		if bytes.Equal(v, []byte("null")) {
			continue
		}
		logs.Info("txeChangefeed result : %s", v)
		txInput <- string(v)
		time.Send("contractOutputs_changefeed")
	}
}

func txHeadFilter() error {
	defer close(txOutput)
	for {
		t, ok := <-txInput
		if !ok {
			//TODO break?
			break
		}
		logs.Info(" txElection step2 : head filter")
		time := monitor.Monitor.NewTiming()
		conout := model.ContractOutput{}
		err := json.Unmarshal([]byte(t), &conout)
		if err != nil {
			logs.Error(err.Error())
			continue
		}
		//main node filter
		mainNodeKey := conout.Transaction.ContractModel.ContractHead.MainPubkey
		myNodeKey := config.Config.Keypair.PublicKey
		if mainNodeKey != myNodeKey {
			logs.Info("I am not the mainnode of the C-output %s", conout.Id)
			continue
		}
		txOutput <- common.Serialize(conout)
		time.Send("txe_validate_head")
	}
	return nil
}

func txValidate() error {
	defer close(txOutput)
	for {

		logs.Info(" txElection step3 : Validate")
		time := monitor.Monitor.NewTiming()
		t, ok := <-cvInput
		if !ok {
			break
		}
		coModel := model.ContractOutput{}
		err := json.Unmarshal([]byte(t), &coModel)
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
		if coModel.Transaction.Operation == "CONTRACT" {
			//TODO ValidateVote  no-need-tood
			//if !coModel.ValidateVote(){
			//	logs.Error(errors.New("invalid vote"))
			//	continue
			//}
		}
		logs.Debug("Validate Hash")
		if !coModel.ValidateContractOutput() {
			//invalid signature
			logs.Error(errors.New("invalid signature"))
			continue
		}
		logs.Debug("Validate sign")
		txOutput <- common.Serialize(coModel)
		time.Send("txe_contractOutput_validate")
	}
}
