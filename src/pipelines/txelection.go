package pipelines

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"

	"unicontract/src/chain"
	"unicontract/src/common/monitor"
	"unicontract/src/config"
	r "unicontract/src/core/db/rethinkdb"
	engineCommon "unicontract/src/core/engine/common"
	"unicontract/src/core/model"

	//	"unicontract/src/common"

	"github.com/astaxie/beego/logs"
	"unicontract/src/core/db/rethinkdb"
)

func txeChangefeed(in io.Reader, out io.Writer) {
	var value interface{}
	res := r.Changefeed("Unicontract", "ContractOutputs")
	for res.Next(&value) {
		contractOutputs_changefeed_time := monitor.Monitor.NewTiming()

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
		out.Write(v)
		contractOutputs_changefeed_time.Send("contractOutputs_changefeed")
	}
}

func txeHeadFilter(in io.Reader, out io.Writer) {

	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {

		n, _ := rd.Read(p)
		logs.Info(" txElection step2 : head filter")
		if n == 0 {
			continue
		}
		t := p[:n]
		conout := model.ContractOutput{}
		err := json.Unmarshal(t, &conout)
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
		out.Write(t)
	}
}

func txeValidate(in io.Reader, out io.Writer) {

	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {

		n, _ := rd.Read(p)
		logs.Info(" txElection step3 : Validate")
		contractOutput_validate_time := monitor.Monitor.NewTiming()
		if n == 0 {
			continue
		}
		t := p[:n]
		coModel := model.ContractOutput{}
		err := json.Unmarshal(t, &coModel)
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
		out.Write(t)

		contractOutput_validate_time.Send("contractOutput_validate")
	}
}

func txeQueryEists(in io.Reader, out io.Writer) {

	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {

		n, _ := rd.Read(p)
		logs.Info("txElection step4 : query eists")
		if n == 0 {
			continue
		}
		t := p[:n]
		coModel := model.ContractOutput{}
		err := json.Unmarshal(t, &coModel)
		if err != nil {
			logs.Error(err.Error())
			continue
		}
		//check whether already exist
		id := coModel.Id
		result, err := chain.GetContractTx("{'id':" + id + "}")
		if err != nil {
			logs.Error(err.Error())
		} else {
			if result.Code != 200 {
				logs.Error(errors.New("request send failed"))
			}
			logs.Info(result.Data)
			//if the unichain already has the contractoutput ,do nothing
			if result.Data == nil {
				continue
			}
		}

		out.Write(t)

	}
}

func txeSend(in io.Reader, out io.Writer) {

	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {

		n, _ := rd.Read(p)
		logs.Info("txElection step5 : send contractoutput")
		time := monitor.Monitor.NewTiming()
		if n == 0 {
			continue
		}
		t := p[:n]
		//write the contract to the taskschedule
		coModel := model.ContractOutput{}
		err := json.Unmarshal(t, &coModel)
		if err != nil {
			logs.Error(err.Error())
			continue
		}
		var taskSchedule model.TaskSchedule
		taskSchedule.ContractId = coModel.Transaction.ContractModel.ContractBody.ContractId
		taskSchedule.StartTime = coModel.Transaction.ContractModel.ContractBody.StartTime
		taskSchedule.EndTime = coModel.Transaction.ContractModel.ContractBody.EndTime

		taskSchedule_write_time := monitor.Monitor.NewTiming()
		err = engineCommon.InsertTaskSchedules(taskSchedule)
		taskSchedule_write_time.Send("taskSchedule_write")
		if err != nil {
			logs.Error("err is \" %s \"\n", err.Error())
		}

		//write the contractoutput to unichain.
		result, err := chain.CreateContractTx(t)
		if err != nil {
			logs.Error(err.Error())
			SaveOutputErrorData(_TableNameSendFailingRecords, t)
			sendFailRecording_count, err := rethinkdb.GetSendFailingRecordsCount()
			if err != nil {
				logs.Error(err.Error())
				continue
			}
			monitor.Monitor.Gauge("sendFailRecording_count", sendFailRecording_count)
			continue
		}
		if result.Code != 200 {
			logs.Error(errors.New("request send failed"))
			SaveOutputErrorData(_TableNameSendFailingRecords, t)
			sendFailRecording_count, err := rethinkdb.GetSendFailingRecordsCount()
			if err != nil {
				logs.Error(err.Error())
				continue
			}
			monitor.Monitor.Gauge("sendFailRecording_count", sendFailRecording_count)
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
		txeSend,
	)

	f, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
	if err != nil {
		logs.Error(err.Error())
	}
	w := bufio.NewWriter(f)
	p(nil, w)
}
