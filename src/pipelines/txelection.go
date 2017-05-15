package pipelines

import (
	"encoding/json"

	"errors"
	"github.com/astaxie/beego/logs"
	"sync"
	engineCommon "unicontract/src/core/engine/common"

	"unicontract/src/chain"
	"unicontract/src/common"
	"unicontract/src/common/monitor"
	"unicontract/src/config"
	"unicontract/src/core/model"
)

func txHeadFilter(arg interface{}) interface{} {
	logs.Info(" txElection step2 : head filter")
	time := monitor.Monitor.NewTiming()
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
	time.Send("txe_validate_head")
	return conout
}
func txValidate(arg interface{}) interface{} {

	logs.Info(" txElection step3 : Validate", arg)
	time := monitor.Monitor.NewTiming()

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
	logs.Debug("Validate Hash")
	if !coModel.ValidateContractOutput() {
		//invalid signature
		logs.Error(errors.New("invalid signature"))
		return nil
	}
	logs.Debug("Validate sign")
	time.Send("txe_contractOutput_validate")
	return coModel
}

func txQueryEists(arg interface{}) interface{} {
	logs.Info("txElection step4 : query eists:", arg)
	time := monitor.Monitor.NewTiming()
	coModel := arg.(model.ContractOutput)
	//check whether already exist
	id := coModel.Id
	result, err := chain.GetContractTx("{'id':" + id + "}")
	if err != nil {
		logs.Error(err.Error())
	} else {
		if result.Code != 200 {
			logs.Error(errors.New("request send failed"))
		}
		//if the unichain already has the contractoutput ,do nothing
		if result.Data == nil {
			return coModel
		}
	}
	time.Send("txe_query_contractOutput")
	return coModel
}
func txSend(arg interface{}) interface{} {
	logs.Info("txElection step5 : send contractoutput")
	time := monitor.Monitor.NewTiming()
	//write the contract to the taskschedule
	coModel := arg.(model.ContractOutput)
	var taskSchedule model.TaskSchedule
	taskSchedule.ContractId = coModel.Transaction.ContractModel.ContractBody.ContractId
	taskSchedule.StartTime = coModel.Transaction.ContractModel.ContractBody.StartTime
	taskSchedule.EndTime = coModel.Transaction.ContractModel.ContractBody.EndTime

	err := engineCommon.InsertTaskSchedules(taskSchedule)
	if err != nil {
		logs.Error("err is \" %s \"\n", err.Error())
	}

	//write the contractoutput to unichain.
	result, err := chain.CreateContractTx(common.StructSerialize(coModel))
	if err != nil {
		logs.Error(err.Error())
		SaveOutputErrorData(_TableNameSendFailingRecords, coModel)
		return nil
	}
	if result.Code != 200 {
		logs.Error(errors.New("request send failed"))
		SaveOutputErrorData(_TableNameSendFailingRecords, coModel)
	}
	time.Send("txe_send_contractOutput")
	return coModel
}

func getChangefeed() *ChangeFeed {
	change := &ChangeFeed{
		db:        "Unicontract",
		table:     "ContractOutputs",
		operation: []string{"insert"},
	}
	go change.runChangeFeed()
	return change
}

func createTxPip() (txPip Pipeline) {
	txNodeSlice := make([]*Node, 0)
	txNodeSlice = append(txNodeSlice, &Node{target: txHeadFilter, routineNum: 1, name: "txHeadFilter"})
	txNodeSlice = append(txNodeSlice, &Node{target: txValidate, routineNum: 1, name: "txValidate"})
	txNodeSlice = append(txNodeSlice, &Node{target: txQueryEists, routineNum: 1, name: "txQueryEists"})
	txNodeSlice = append(txNodeSlice, &Node{target: txSend, routineNum: 1, name: "txSends"})
	txPip = Pipeline{
		nodes: txNodeSlice,
	}
	return txPip
}

func start() {
	waitRoutine := sync.WaitGroup{}
	waitRoutine.Add(1)

	txPip := createTxPip()
	changefeed := getChangefeed()
	txPip.setup(&changefeed.node)
	txPip.start()
	//TODO how to run in main.go
	waitRoutine.Wait()
}

//runtime.NumGoroutine()
