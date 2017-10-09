package pipelines

import (
	"encoding/json"
	"errors"
	"sync"

	"unicontract/src/chain"
	"unicontract/src/common"
	"unicontract/src/common/monitor"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/config"
	engineCommon "unicontract/src/core/engine/common"
	"unicontract/src/core/engine/common/db"
	"unicontract/src/core/model"

	"github.com/astaxie/beego"
)

func txHeadFilter(arg interface{}) interface{} {
	uniledgerlog.Info(" txElection step2 : head filter")
	bs, err := json.Marshal(arg)
	if err != nil {
		uniledgerlog.Error(err.Error())
		return nil
	}
	conout := model.ContractOutput{}
	err = json.Unmarshal(bs, &conout)
	if err != nil {
		uniledgerlog.Error(err.Error())
		return nil
	}
	//main node filter
	contractHead := conout.Transaction.ContractModel.ContractHead
	mainNodeKey := ""
	if contractHead != nil {
		mainNodeKey = contractHead.MainPubkey
	}
	myNodeKey := config.Config.Keypair.PublicKey
	if mainNodeKey != myNodeKey {
		uniledgerlog.Info("I am not the mainnode of the C-output %s", conout.Id)
		return nil
	}
	return conout
}

func txValidate(arg interface{}) interface{} {

	uniledgerlog.Info(" txElection step3 : Validate", arg)
	contractOutput_validate_time := monitor.Monitor.NewTiming()

	coModel := arg.(model.ContractOutput)

	if !coModel.HasEnoughVotes() {
		//TODO always has enough votes
		//not enough votes
		return nil
	}
	if !coModel.ValidateHash() {
		//invalid hash
		uniledgerlog.Error(errors.New("invalid hash"))
		return nil
	}
	/*	if coModel.Transaction.Operation == "TRANSFER" {
		isFreeze := false
		conditions := coModel.Transaction.Conditions

		for _, value := range conditions {
			isFreeze = value.Isfreeze
			if isFreeze {
				return coModel
			}
		}

	}*/
	if coModel.Transaction.Operation == "CONTRACT" {
		//TODO ValidateVote  no-need-tood
		//if !coModel.ValidateVote(){
		//	uniledgerlog.Error(errors.New("invalid vote"))
		//	continue
		//}
	}
	//uniledgerlog.Debug("Validate Hash")
	if !coModel.ValidateContractOutput() {
		//invalid signature
		uniledgerlog.Error(errors.New("invalid signature"))
		return nil
	}
	//uniledgerlog.Debug("Validate sign")
	contractOutput_validate_time.Send("contractOutput_validate")
	return coModel
}

func txQueryEists(arg interface{}) interface{} {
	uniledgerlog.Info("txElection step4 : query eists:", arg)
	coModel := arg.(model.ContractOutput)
	chainType := coModel.Chaintype
	//check whether already exist
	id := coModel.Id
	result, err := chain.GetContractTx(`{"tx_id":"`+id+`"}`, chainType)
	if err != nil {
		uniledgerlog.Error(err.Error())
		return coModel
	} else {
		if result.Code != 200 {
			uniledgerlog.Error(errors.New("request send failed"))
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
	uniledgerlog.Info("txElection step5 : send contractoutput")
	//write the contract to the taskschedule
	coModel := arg.(model.ContractOutput)
	contractOwner := coModel.Transaction.ContractModel.ContractBody.ContractOwners
	contractSign := coModel.Transaction.ContractModel.ContractBody.ContractSignatures
	contractState := coModel.Transaction.ContractModel.ContractBody.ContractState
	if ("Contract_Signature" == contractState) && (len(contractSign) == len(contractOwner)) {
		var taskSchedule db.TaskSchedule
		taskSchedule.ContractHashId = coModel.Transaction.ContractModel.Id
		taskSchedule.ContractId = coModel.Transaction.ContractModel.ContractBody.ContractId
		taskSchedule.StartTime = coModel.Transaction.ContractModel.ContractBody.StartTime
		taskSchedule.EndTime = coModel.Transaction.ContractModel.ContractBody.EndTime

		taskSchedule_write_time := monitor.Monitor.NewTiming()
		err := engineCommon.InsertTaskSchedules(taskSchedule)
		taskSchedule_write_time.Send("taskSchedule_write")
		if err != nil {
			uniledgerlog.Error("err is \" %s \"\n", err.Error())
		}
	}
	//TODO qianming num  not equal owner num
	//if ("Contract_Discarded" == contractState) && (len(contractSign) == len(contractOwner)) {
	if "Contract_Discarded" == contractState {
		contractionId := coModel.Transaction.ContractModel.ContractBody.ContractId
		err := engineCommon.TerminateContractBatch(contractionId)
		if err != nil {
			uniledgerlog.Error("err is \" %s \"\n", err.Error())
		}
	}
	//var chainType = ""
	//write the contractoutput to unichain.
	chainType := coModel.Chaintype
	uniledgerlog.Info(chainType)
	result, err := chain.CreateContractTx(common.StructSerialize(coModel), chainType)
	if err != nil {
		uniledgerlog.Error(err.Error())
		SaveOutputErrorData(_TableNameSendFailingRecords, coModel)
		//count, err := rethinkdb.GetSendFailingRecordsCount()
		//if err != nil {
		//	uniledgerlog.Error(err.Error())
		//}
		return nil
	}
	if result.Code != 200 {
		uniledgerlog.Error(result.Message)
		uniledgerlog.Error(errors.New("request send failed"))
		SaveOutputErrorData(_TableNameSendFailingRecords, coModel)
		//count, err := rethinkdb.GetSendFailingRecordsCount()
		//if err != nil {
		//	uniledgerlog.Error(err.Error())
		//}
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
	NodeGoroutineNum, err := beego.AppConfig.Int("PipelineNodeGoroutineNum")
	if err != nil {
		uniledgerlog.Error(err)
		NodeGoroutineNum = 1
	}
	txNodeSlice = append(txNodeSlice, &Node{target: txHeadFilter, routineNum: NodeGoroutineNum, name: "txHeadFilter"})
	txNodeSlice = append(txNodeSlice, &Node{target: txValidate, routineNum: NodeGoroutineNum, name: "txValidate"})
	txNodeSlice = append(txNodeSlice, &Node{target: txQueryEists, routineNum: NodeGoroutineNum, name: "txQueryEists"})
	txNodeSlice = append(txNodeSlice, &Node{target: txSend, routineNum: NodeGoroutineNum, name: "txSends", timeout: 10})
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
