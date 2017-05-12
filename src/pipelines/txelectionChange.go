package pipelines

import (
	"encoding/json"

	"unicontract/src/common/monitor"
	"unicontract/src/config"
	"unicontract/src/core/model"
	engineCommon "unicontract/src/core/engine/common"
	"errors"
	"github.com/astaxie/beego/logs"
	"unicontract/src/common"
	"unicontract/src/chain"
)

const (
	_TXTHREAD = 10
)

var (
	txPool   *ThreadPool
	txInput  chan string
	txOutput chan string
)

func txHeadFilter(args ... interface{}) {
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

func pip1(i int) interface{} {
	var p1 = "p1"
	logs.Info("p1", i)
	return p1
}
func pip2(f float32, a string) interface{} {

	var p2 = "p2"
	logs.Info(p2)
	return p2
}

func pip3(args ... interface{}) interface{} {
	var p3 = "p3"
	logs.Info(p3)
	return p3
}

func pip4(args ... interface{}) interface{} {
	var p4 = "p4"
	logs.Info(p4)
	return p4
}

func getChangefeed() Node {
	//todo  fix
	change := ChangeFeed{
		node:Node{target: pip3, processes: 1,name:"pip3"},
		table:"ContractOutputs",
		operation:[]string{"insert"},
	}
	change.runChangeFeed()
	return 	change.node
}

func createTxPip() (txPip Pipeline) {
	logs.Info("createTxPip")
	txNodeSlice := make([]Node, 0, 5)
	//txNodeSlice = append(txNodeSlice, Node{target: pip1, processes: 1})
	//txNodeSlice = append(txNodeSlice, Node{target: pip2, processes: 2})
	txNodeSlice = append(txNodeSlice, Node{target: pip3, processes: 1,name:"pip3"})
	txNodeSlice = append(txNodeSlice, Node{target: pip4, processes: 2,name:"pip4"})
	txPip = Pipeline{
		nodes: txNodeSlice,
	}
	return txPip
}

func start() {
	logs.Info("start1")
	txPip := createTxPip()
	logs.Info("start2")
	txPip.setup(getChangefeed())
	logs.Info("start3")
	txPip.start()
	logs.Info("start4")
}
