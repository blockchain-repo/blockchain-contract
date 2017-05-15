package pipelines

import (
	"encoding/json"

	"errors"
	"github.com/astaxie/beego/logs"
	"sync"
	"unicontract/src/chain"
	"unicontract/src/common"
	"unicontract/src/common/monitor"
	"unicontract/src/config"
	engineCommon "unicontract/src/core/engine/common"
	"unicontract/src/core/model"
)

const (
	_TXTHREAD = 10
)

var (
	txPool   *ThreadPool
	txInput  chan string
	txOutput chan string
)

func txHeadFilter(args ...interface{}) {
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
}

func txValidate(coModel model.ContractOutput) {
	defer close(txOutput)
	for {

		logs.Info(" txElection step3 : Validate")
		time := monitor.Monitor.NewTiming()
		t, ok := <-txInput
		if !ok {
			break
		}
		//coModel := model.ContractOutput{}
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
func txQueryEists() {
	defer close(txOutput)
	for {
		logs.Info("txElection step4 : query eists")
		time := monitor.Monitor.NewTiming()

		t, ok := <-txInput
		if !ok {
			break
		}

		coModel := model.ContractOutput{}
		err := json.Unmarshal([]byte(t), &coModel)
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
		txOutput <- common.Serialize(coModel)
		time.Send("txe_query_contractOutput")
	}
}

func txSend() {
	defer close(txOutput)
	for {
		logs.Info("txElection step5 : send contractoutput")
		time := monitor.Monitor.NewTiming()
		t, ok := <-txInput
		if !ok {
			break
		}
		//write the contract to the taskschedule
		coModel := model.ContractOutput{}
		err := json.Unmarshal([]byte(t), &coModel)
		if err != nil {
			logs.Error(err.Error())
			continue
		}
		var taskSchedule model.TaskSchedule
		taskSchedule.ContractId = coModel.Transaction.ContractModel.ContractBody.ContractId
		taskSchedule.StartTime = coModel.Transaction.ContractModel.ContractBody.StartTime
		taskSchedule.EndTime = coModel.Transaction.ContractModel.ContractBody.EndTime

		err = engineCommon.InsertTaskSchedules(taskSchedule)
		if err != nil {
			logs.Error("err is \" %s \"\n", err.Error())
		}

		//write the contractoutput to unichain.
		result, err := chain.CreateContractTx(t)
		if err != nil {
			logs.Error(err.Error())
			SaveOutputErrorData(_TableNameSendFailingRecords, []byte(t))
			continue
		}
		if result.Code != 200 {
			logs.Error(errors.New("request send failed"))
			SaveOutputErrorData(_TableNameSendFailingRecords, []byte(t))
		}

		time.Send("txe_send_contractOutput")
	}
}

func pip3(arg interface{}) interface{} {
	logs.Info("P3 param:", arg)
	s := common.Serialize(arg)
	logs.Info("P3 return:===", s)
	return s
}

func pip4(arg interface{}) interface{} {
	logs.Info("P4 param:===", arg)
	return arg
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
	txNodeSlice = append(txNodeSlice, &Node{target: pip3, routineNum: 1, name: "pip3"})
	txNodeSlice = append(txNodeSlice, &Node{target: pip4, routineNum: 1, name: "pip4"})
	txPip = Pipeline{
		nodes: txNodeSlice,
	}
	return txPip
}

func start() {
	txPip := createTxPip()
	changefeed := getChangefeed()
	txPip.setup(&changefeed.node)
	waitRoutine := sync.WaitGroup{}
	waitRoutine.Add(1)
	txPip.start()
	waitRoutine.Wait()
}

//runtime.NumGoroutine()
