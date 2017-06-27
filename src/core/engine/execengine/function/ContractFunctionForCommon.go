package function

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"unicontract/src/config"
	"unicontract/src/core/engine/common"
	"unicontract/src/transaction"

	"github.com/astaxie/beego/logs"
	"reflect"
	"strings"
)

//++++++++++++++++++++++++++++++++++++++++++++++++++++++
//+++++++++++++++合约机公用方法集+++++++++++++++++++++++
//++++++++++++++++++++++++++++++++++++++++++++++++++++++
//获取当前日期FuncGetNowDay()
//获取当前时间FuncGetNowDate()
//休眠指定时间FuncSleepTime(sleeptime)
//资产转移FuncTransferAsset(user_A, '{user_B, amount}')
//资产创建FuncCreateAsset(user_A, '{user_B, amount}')
//获取合约产出FuncGetContracOutputtById(contract_id)
//合约产出是否入链FuncIsConPutInUnichian(id)

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

//获取当前日期的Day: int
func FuncGetNowDay(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	day := time.Now().Day()

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(day)
	return v_result, v_err
}

//获取当前时间 Date: 2017-06-20 17:00:00
func FuncGetNowDate(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	str_date := tm.Format("2006-01-02 03:04:05")

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(str_date)
	return v_result, v_err
}

//获取时间戳
func FuncGetNowDateTimestamp(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(common.GenTimestamp())
	return v_result, v_err
}

//休眠指定时间
//Args: sleeptime int
func FuncSleepTime(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	if len(args) != 1 {
		v_err = errors.New("param num error")
		return v_result, v_err
	}
	//user provide
	var sleeptime string = args[0].(string)
	int_sleeptime, _ := strconv.Atoi(sleeptime)
	time.Sleep(time.Second * time.Duration(int_sleeptime))

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("")
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
	logs.Error("======In FuncTransferAsset==========")
	var v_result common.OperateResult = common.OperateResult{}
	var v_err error = nil

	//var v_map_args map[string]interface{} = nil
	if len(args) != 8 {
		v_err = errors.New("param num error")
		return v_result, v_err
	}
	logs.Error("======after param check FuncTransferAsset==========")
	//user provide
	var ownerBefore string = ""
	switch args[0].(type) {
	case reflect.Value:
		ownerBefore = args[0].(reflect.Value).String()
	case string:
		ownerBefore = args[0].(string)
	}
	ownerBefore = strings.Replace(ownerBefore, "\"", "", -1)
	ownerBefore = strings.Trim(ownerBefore, " ")
	var recipientsStr string = ""
	switch args[1].(type) {
	case reflect.Value:
		recipientsStr = args[1].(reflect.Value).String()
	case string:
		recipientsStr = args[1].(string)
	}
	recipientsStr = strings.Replace(recipientsStr, "\"", "", -1)
	recipientsStr = strings.Trim(recipientsStr, " ")
	var money_amount string = ""
	switch args[2].(type) {
	case reflect.Value:
		money_amount = args[2].(reflect.Value).String()
	case string:
		money_amount = args[2].(string)
	}
	money_amount = strings.Replace(money_amount, "\"", "", -1)
	money_amount = strings.Trim(money_amount, " ")
	var recipients [][2]interface{} = [][2]interface{}{}
	//s := strings.Split(recipientsStr, "&")
	//fmt.Println(s, len(s))
	//for i := 0; i < len(s); i++ {
	//	ss := strings.Split(s[i], "#")
	//	ownAfter := ss[0]
	//	amount, _ := strconv.ParseFloat(ss[1], 64)
	//	recipients = append(recipients, [2]interface{}{ownAfter, amount})
	//}
	amount, _ := strconv.ParseFloat(money_amount, 64)
	recipients = append(recipients, [2]interface{}{recipientsStr, amount})
	//executer provide
	var contractStr string = args[3].(string)
	contractStr = strings.Replace(contractStr, "\\\\", "", -1)
	contractStr = strings.TrimLeft(contractStr, " ")
	contractStr = strings.TrimRight(contractStr, " ")
	contractStr = strings.TrimLeft(contractStr, "\"")
	contractStr = strings.TrimRight(contractStr, "\"")
	contractStr = strings.TrimLeft(contractStr, "\"")

	var contractId string = args[4].(string)
	contractId = strings.Replace(contractId, "\"", "", -1)
	contractId = strings.Trim(contractId, " ")
	var taskId string = args[5].(string)
	taskId = strings.Replace(taskId, "\"", "", -1)
	taskId = strings.Trim(taskId, " ")
	var str_taskIndex string = args[6].(string)
	str_taskIndex = strings.Replace(str_taskIndex, "\"", "", -1)
	str_taskIndex = strings.Trim(str_taskIndex, " ")
	taskIndex, _ := strconv.Atoi(str_taskIndex)
	var mainPubkey string = args[7].(string)
	mainPubkey = strings.Replace(mainPubkey, "\"", "", -1)
	mainPubkey = strings.Trim(mainPubkey, " ")
	logs.Error("======Param:=========")
	logs.Error(ownerBefore)
	logs.Error(recipientsStr)
	logs.Error(money_amount)
	logs.Error(contractStr)
	logs.Error(contractId)
	logs.Error(taskId)
	logs.Error(taskIndex)
	logs.Error(mainPubkey)
	var contractHashId string = ""
	var metadataStr string = ""
	var relationStr string = transaction.GenerateRelation(contractHashId, contractId, taskId, taskIndex)

	var outputStr string
	/*
		do freeze
	*/
	mykey := config.Config.Keypair.PublicKey
	logs.Info("==MinPubkey: ", mainPubkey)
	logs.Info("==mykey: ", mykey)
	logs.Info("==equals: ", mainPubkey == mykey)
	//check main pubkey
	if mainPubkey == mykey {
		//if mainNode, do freeze;
		logs.Info("mainPubkey ")
		var reciForFre [][2]interface{} = [][2]interface{}{
			[2]interface{}{ownerBefore, amount},
		}
		logs.Info("contractStr: ", contractStr)
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
		logs.Info("not mainPubkey ")
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
	if len(args) != 6 {
		v_err = errors.New("param num error")
		return v_result, v_err
	}
	//user provide
	var ownerBefore string = args[0].(string)
	//var recipients [][2]interface{} = args[1].([][2]interface{})
	var recipientsStr string = args[1].(string)
	var recipients [][2]interface{} = [][2]interface{}{}
	s := strings.Split(recipientsStr, "&")
	//fmt.Println(s, len(s))
	for i := 0; i < len(s); i++ {
		ss := strings.Split(s[i], "#")
		ownAfter := ss[0]
		amount, _ := strconv.ParseFloat(ss[1], 64)
		recipients = append(recipients, [2]interface{}{ownAfter, amount})
	}

	//executer provide
	var contractStr string = args[2].(string)
	var contractId string = args[3].(string)
	var taskId string = args[4].(string)
	var taskIndex int = args[5].(int)
	var contractHashId string = ""
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
	if len(args) != 4 {
		v_err = errors.New("param num error")
		return v_result, v_err
	}
	var contractStr string = args[0].(string)
	var contractId string = args[1].(string)
	var taskId string = args[2].(string)
	var taskIndex int = args[3].(int)
	var contractHashId string = ""
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
	v_result.SetOutput(outputStr)
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
	if len(args) != 3 {
		v_err = errors.New("param num error")
		return v_result, v_err
	}
	var contractOutPut string = args[0].(string)
	var taskStatus string = args[1].(string)
	var contractState = args[2].(string)
	outputStr, v_err := transaction.ExecuteInterimComplete(contractOutPut, taskStatus, contractState)
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
	if len(args) != 5 {
		v_err = errors.New("param num error")
		return v_result, v_err
	}
	//user provide
	var ownerBefore string = args[0].(string)
	//executer provide
	var contractStr string = args[1].(string)
	var contractId string = args[2].(string)
	var taskId string = args[3].(string)
	var taskIndex int = args[4].(int)
	//var mainPubkey string = args[7].(string)
	var recipients [][2]interface{} = [][2]interface{}{}
	var contractHashId string = ""
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
