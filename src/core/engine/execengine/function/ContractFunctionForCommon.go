package function

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"
	"unicontract/src/config"
	"unicontract/src/core/engine/common"
	"unicontract/src/transaction"
)

//++++++++++++++++++++++++++++++++++++++++++++++++++++++
//+++++++++++++++合约机公用方法集+++++++++++++++++++++++
//++++++++++++++++++++++++++++++++++++++++++++++++++++++

//测试方法
func FuncTestMethod(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	var v_map_args map[string]interface{} = nil
	if len(args) != 0 {
		v_map_args = make(map[string]interface{}, 0)
	}
	//识别可变参数
	for v_idx, v_args := range args {
		tmp_arg := "v_arg_" + strconv.Itoa(v_idx)
		v_map_args[tmp_arg] = v_args
	}
	//调用参数
	for v_name, v_value := range v_map_args {
		fmt.Println(v_name, ":", v_value)
	}
	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//资产转移方法
/*
  Desc:transfer asset, only generate the output,Not insert into the db
  Args:
  	0: ownerbefore(string):	the pubkey who transfer assets
  	1: recipients([][2]interface{}): A list of keys that represent the receivers of this transfer.
	2: contractStr(string):the contract str which this task execute
	3: contractHashId(string): contractHashId
	4: contractId(string): contractId
	5: taskId(string): taskId
	6: TaskExecuteIdx(int): TaskExecuteIdx
	7: mainPubkey(string): the node pubkey which will freeze asset
*/
func FuncTransferAsset(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult = common.OperateResult{}
	var v_err error = nil

	//var v_map_args map[string]interface{} = nil
	if len(args) != 8 {
		v_err = errors.New("param num error")
		return v_result, v_err
	}

	//user provide
	var ownerBefore string = args[0].(string)
	var recipients [][2]interface{} = args[1].([][2]interface{})
	//executer provide
	var contractStr string = args[2].(string)
	var contractHashId string = args[3].(string)
	var contractId string = args[4].(string)
	var taskId string = args[5].(string)
	var taskIndex int = args[6].(int)
	var mainPubkey string = args[7].(string)
	var metadataStr string = ""
	var relationStr string = transaction.GenerateRelation(contractHashId, contractId, taskId, taskIndex)

	var outputStr string
	/*
		do freeze
	*/
	mykey := config.Config.Keypair.PublicKey
	//check main pubkey
	if mainPubkey == mykey {
		//if mainNode, do freeze;
		var reciForFre [][2]interface{} = [][2]interface{}{
			[2]interface{}{ownerBefore, 100},
		}
		outputStr, v_err = transaction.ExecuteFreeze("FREEZE", ownerBefore, reciForFre, metadataStr, relationStr, contractStr)
		//if v_err != nil {
		//	logs.Error(v_err)
		//	v_result.SetCode(400)
		//	v_result.SetMessage(v_err.Error())
		//	return v_result, v_err
		//}
		//wait for the freeze asset write into the unichain
		time.Sleep(time.Second * 3)
	} else {
		// not mainNode, wait for the main node write the freeze-asset into the unchain
		time.Sleep(time.Second * 5)
	}
	/*
		do transfer
	*/
	logs.Info("for ")
	for i := 0; i <= 3; i++ {
		//transfer asset
		outputStr, v_err = transaction.ExecuteTransfer("TRANSFER", ownerBefore, recipients, metadataStr, relationStr, contractStr)
		if v_err != nil && i == 3 {
			logs.Error(v_err)
			v_result.SetCode(400)
			v_result.SetMessage(v_err.Error())
			return v_result, v_err
		}
		if v_err == nil {
			break
		}
		if i != 3 {
			time.Sleep(time.Second * 5)
		}
	}
	//构建返回值
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(outputStr)
	return v_result, v_err
}

/*
  Desc:transfer asset, update the output and insert into the db
  Args:
	0: contractStr(string): the return from the func `FuncTransferAsset`
	1: taskStatus(string): the taskStatus need to update
*/
func FuncTransferAssetComplete(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult = common.OperateResult{}
	var v_err error = nil
	if len(args) != 2 {
		v_err = errors.New("param num error")
		return v_result, v_err
	}
	var contractOutPut string = args[0].(string)
	var taskStatus string = args[1].(string)
	outputStr, v_err := transaction.ExecuteTransferComplete(contractOutPut, taskStatus)

	//构建返回值
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(outputStr)
	return v_result, v_err
}

//create asset
/*
  Desc:create asset, generate the output and insert into db
  Args:
  	0: ownerbefore(string):	the pubkey who transfer assets
  	1: recipients([][2]interface{}): A list of keys that represent the receivers of this transfer.
	2: contractStr(string):the contract str which this task execute
	3: contractHashId(string): contractHashId
	4: contractId(string): contractId
	5: taskId(string): taskId
	6: TaskExecuteIdx(int): TaskExecuteIdx
*/
func FuncCreateAsset(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult = common.OperateResult{}
	var v_err error = nil
	if len(args) != 7 {
		v_err = errors.New("param num error")
		return v_result, v_err
	}
	//user provide
	var ownerBefore string = args[0].(string)
	var recipients [][2]interface{} = args[1].([][2]interface{})
	//executer provide
	var contractStr string = args[2].(string)
	var contractHashId string = args[3].(string)
	var contractId string = args[4].(string)
	var taskId string = args[5].(string)
	var taskIndex int = args[6].(int)
	//var mainPubkey string = args[7].(string)
	var metadataStr string = ""
	var relationStr string = transaction.GenerateRelation(contractHashId, contractId, taskId, taskIndex)
	//tx_signers []string, recipients [][2]interface{}, metadataStr string,
	//relationStr string, contractStr string
	outputStr, v_err := transaction.ExecuteCreate(ownerBefore, recipients, metadataStr, relationStr, contractStr)

	if v_err != nil {
		return v_result, v_err
	}
	//构建返回值
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(outputStr)
	return v_result, v_err
}

/*
  Desc: generate an empty output as a interim, Not insert into the db
  Args:
  	0: ownerbefore(string):	the pubkey who transfer assets
  	1: recipients([][2]interface{}): A list of keys that represent the receivers of this transfer. it should be nil
	2: contractStr(string):the contract str which this task execute
	3: contractHashId(string): contractHashId
	4: contractId(string): contractId
	5: taskId(string): taskId
	6: TaskExecuteIdx(int): TaskExecuteIdx
*/
func FuncInterim(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult = common.OperateResult{}
	var v_err error = nil
	if len(args) != 5 {
		v_err = errors.New("param num error")
		return v_result, v_err
	}
	var contractStr string = args[0].(string)
	var contractHashId string = args[1].(string)
	var contractId string = args[2].(string)
	var taskId string = args[3].(string)
	var taskIndex int = args[4].(int)
	var metadataStr string = ""
	var relationStr string = transaction.GenerateRelation(contractHashId, contractId, taskId, taskIndex)
	outputStr, v_err := transaction.ExecuteInterim(metadataStr, relationStr, contractStr)
	if v_err != nil {
		return v_result, v_err
	}
	//构建返回值
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(outputStr)
	return v_result, v_err
}

/*
  Desc:update tasksStatus in the empty output and insert into the db
  Args:
	0: contractStr(string): the return from the func `FuncInterim`
	1: taskStatus(string): the taskStatus need to update
*/
func FuncInterimComplete(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult = common.OperateResult{}
	var v_err error = nil
	if len(args) != 2 {
		v_err = errors.New("param num error")
		return v_result, v_err
	}
	var contractOutPut string = args[0].(string)
	var taskStatus string = args[1].(string)
	outputStr, v_err := transaction.ExecuteInterimComplete(contractOutPut, taskStatus)
	if v_err != nil {
		return v_result, v_err
	}
	//构建返回值
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(outputStr)
	return v_result, v_err
}

//解冻资产方法
/*
  Desc: unfreeze the asset
  Args:
  	0: ownerbefore(string):	the pubkey who need unfereeze assets
  	1: recipients([][2]interface{}): it should be nil here
	2: contractStr(string):the contract str which the freeze task execute
	3: contractHashId(string): contractHashId
	4: contractId(string): contractId
	5: taskId(string): taskId
	6: TaskExecuteIdx(int): TaskExecuteIdx
*/
func FuncUnfreezeAsset(args ...interface{}) (common.OperateResult, error) {
	//userPubKey string, contractId string, taskId string, taskNum int
	var v_result common.OperateResult = common.OperateResult{}
	var v_err error = nil
	if len(args) != 7 {
		v_err = errors.New("param num error")
		return v_result, v_err
	}
	//user provide
	var ownerBefore string = args[0].(string)
	var recipients [][2]interface{} = [][2]interface{}{}
	//executer provide
	var contractStr string = args[2].(string)
	var contractHashId string = args[3].(string)
	var contractId string = args[4].(string)
	var taskId string = args[5].(string)
	var taskIndex int = args[6].(int)
	//var mainPubkey string = args[7].(string)
	var metadataStr string = ""
	var relationStr string = transaction.GenerateRelation(contractHashId, contractId, taskId, taskIndex)

	outputStr, v_err := transaction.ExecuteUnfreeze("UNFREEZE", ownerBefore, recipients,
		metadataStr, relationStr, contractStr)

	if v_err != nil {
		return v_result, v_err
	}
	//构建返回值
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(outputStr)
	return v_result, v_err
}

/*
Desc:根据合约ContractID查找合约
Args:
	0: contract id
return: only the operation of the output is "CONTRACT" ,get its contract model
*/
func FuncGetContracOutputtById(args ...interface{}) (common.OperateResult, error) {
	//contractId string
	var v_err error = nil
	v_result := common.OperateResult{}
	if len(args) != 1 {
		v_err = errors.New("param num error")
		return v_result, v_err
	}
	var conId string = args[0].(string)
	conStr, v_err := transaction.ExecuteGetContract(conId)
	if v_err != nil {
		v_result.SetCode(400)
		v_result.SetMessage("get error")
		v_result.SetData(conStr)
		return v_result, v_err
	}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(conStr)
	logs.Info(conStr)
	return v_result, v_err
}

/*
Desc: check out is the output in unichain with the contracthashid
Args:
	0:contracthashid
return:
	true/false
*/
func FuncIsConPutInUnichian(args ...interface{}) (common.OperateResult, error) {

	var v_result common.OperateResult = common.OperateResult{}
	var v_err error = nil
	if len(args) != 1 {
		v_err = errors.New("param num error")
		return v_result, v_err
	}
	contractHashId := args[0].(string)

	flag, v_err := transaction.IsOutputInUnichain(contractHashId)
	if v_err != nil {
		v_result.SetCode(400)
		v_result.SetMessage("query error!")
		v_result.SetData(strconv.FormatBool(false))
		return v_result, v_err
	}
	//构建返回值
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(strconv.FormatBool(flag))
	return v_result, v_err
}
