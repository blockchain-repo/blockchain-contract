package pipelines

import (
	"encoding/json"
	"errors"
	"sync"

	"unicontract/src/chain"
	"unicontract/src/common"
	"unicontract/src/common/monitor"
	"unicontract/src/config"
	engineCommon "unicontract/src/core/engine/common"
	"unicontract/src/core/model"

	"github.com/astaxie/beego/logs"
)

func txHeadFilter(arg interface{}) interface{} {
	logs.Info(" txElection step2 : head filter")
	bs, err := json.Marshal(arg)
	if err != nil {
		logs.Error(err.Error())
		return nil
	}
	conout := model.ContractOutput{}
	err = json.Unmarshal(bs, &conout)
	if err != nil {
		logs.Error(err.Error())
		return nil
	}
	//main node filter
	mainNodeKey := conout.Transaction.ContractModel.ContractHead.MainPubkey
	myNodeKey := config.Config.Keypair.PublicKey
	if mainNodeKey != myNodeKey {
		logs.Info("I am not the mainnode of the C-output %s", conout.Id)
		return nil
	}
	return conout
}

func txValidate(arg interface{}) interface{} {

	logs.Info(" txElection step3 : Validate", arg)
	contractOutput_validate_time := monitor.Monitor.NewTiming()

	coModel := arg.(model.ContractOutput)

	if !coModel.HasEnoughVotes() {
		//not enough votes
		return nil
	}
	if !coModel.ValidateHash() {
		//invalid hash
		logs.Error(errors.New("invalid hash"))
		return nil
	}
	if coModel.Transaction.Operation == "CONTRACT" {
		//TODO ValidateVote  no-need-tood
		//if !coModel.ValidateVote(){
		//	logs.Error(errors.New("invalid vote"))
		//	continue
		//}
	}
	//logs.Debug("Validate Hash")
	if !coModel.ValidateContractOutput() {
		//invalid signature
		logs.Error(errors.New("invalid signature"))
		return nil
	}
	//logs.Debug("Validate sign")
	contractOutput_validate_time.Send("contractOutput_validate")
	return coModel
}

func txQueryEists(arg interface{}) interface{} {
	logs.Info("txElection step4 : query eists:", arg)
	coModel := arg.(model.ContractOutput)
	//check whether already exist
	id := coModel.Id
	result, err := chain.GetContractTx(`{"tx_id":"` + id + `"}`)
	if err != nil {
		logs.Error(err.Error())
		return coModel
	} else {
		if result.Code != 200 {
			logs.Error(errors.New("request send failed"))
			return coModel
		}
	}
	res, ok := result.Data.([]interface{})
	if !ok {
		return coModel
	}
	if len(res) != 0 {
		return nil
	}
	return coModel
}
func txSend(arg interface{}) interface{} {
	logs.Info("txElection step5 : send contractoutput")
	//write the contract to the taskschedule
	coModel := arg.(model.ContractOutput)
	var taskSchedule model.TaskSchedule
	taskSchedule.ContractHashId = coModel.Transaction.ContractModel.Id
	taskSchedule.ContractId = coModel.Transaction.ContractModel.ContractBody.ContractId
	taskSchedule.StartTime = coModel.Transaction.ContractModel.ContractBody.StartTime
	taskSchedule.EndTime = coModel.Transaction.ContractModel.ContractBody.EndTime

	taskSchedule_write_time := monitor.Monitor.NewTiming()
	err := engineCommon.InsertTaskSchedules(taskSchedule)
	taskSchedule_write_time.Send("taskSchedule_write")
	if err != nil {
		logs.Error("err is \" %s \"\n", err.Error())
	}

	//write the contractoutput to unichain.
	result, err := chain.CreateContractTx(common.StructSerialize(coModel))
	if err != nil {
		logs.Error(err.Error())
		SaveOutputErrorData(_TableNameSendFailingRecords, coModel)
		//count, err := rethinkdb.GetSendFailingRecordsCount()
		//if err != nil {
		//	logs.Error(err.Error())
		//}
		//monitor.Monitor.Gauge("sendFailingRecords_count", count)
		monitor.Monitor.Count("sendFailingRecords_count", 1)
		return nil
	}
	if result.Code != 200 {
		logs.Error(errors.New("request send failed"))
		SaveOutputErrorData(_TableNameSendFailingRecords, coModel)
		//count, err := rethinkdb.GetSendFailingRecordsCount()
		//if err != nil {
		//	logs.Error(err.Error())
		//}
		//monitor.Monitor.Gauge("sendFailingRecords_count", count)
		monitor.Monitor.Count("sendFailingRecords_count", 1)
	}
	return coModel
}

func getTxChangefeed() *ChangeFeed {
	change := &ChangeFeed{
		db:        "Unicontract",
		table:     "ContractOutputs",
		operation: INSERT | UPDATE,
	}
	go change.runForever()
	return change
}

func createTxPip() (txPip Pipeline) {
	txNodeSlice := make([]*Node, 0)
	txNodeSlice = append(txNodeSlice, &Node{target: txHeadFilter, routineNum: 1, name: "txHeadFilter"})
	txNodeSlice = append(txNodeSlice, &Node{target: txValidate, routineNum: 1, name: "txValidate"})
	txNodeSlice = append(txNodeSlice, &Node{target: txQueryEists, routineNum: 1, name: "txQueryEists"})
	txNodeSlice = append(txNodeSlice, &Node{target: txSend, routineNum: 1, name: "txSends",timeout:10})
	txPip = Pipeline{
		nodes: txNodeSlice,
	}
	return txPip
}

func startTxElection() {
	txPip := createTxPip()
	changefeed := getTxChangefeed()
	txPip.setup(&changefeed.node)
	txPip.start()

	waitRoutine := sync.WaitGroup{}
	waitRoutine.Add(1)
	waitRoutine.Wait()
}

//runtime.NumGoroutine()
