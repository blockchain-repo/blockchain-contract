package function

import "time"
import "unicontract/src/core/engine/common"

//++++++++++++++++++++++++++++++++++++++++++++++++++++++
//++++++++++合约机【自动缴纳房租】专用扩展方法++++++++++
//++++++++++++++++++++++++++++++++++++++++++++++++++++++
//判定下月是否需要缴纳房租FuncIfContinueToPayNextMonth(contract_demo_2.EndTime)
//到期合约退出FuncContractExitForComplete()
//查询账户余额FuncQueryUserBalance(User_A)
//转房租给房东FuncTransferMoney(UserA, UserB, 5000)
//打印凭条收据FuncPrintReceipt(UserA, UserB, 5000)
//提示给房东打钱FuncRemindAccount(UserA, UserB, 5000)
//空动作FuncNoAction()

//查询租房合同终止日期，判定下月是否需要缴纳房租
//Args:　contract_demo_2.EndTime　　合约终止日期
func FuncIfContinueToPayNextMonth(args ...interface{}) (common.OperateResult, error) {
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

//下月合同到期，不需继续缴纳房租，合约退出
func FuncContractExitForComplete(args ...interface{}) (common.OperateResult, error) {
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

//查询账户余额
//Args:　User_A
func FuncQueryUserBalance(args ...interface{}) (common.OperateResult, error) {
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

//转账给房东5000元
//Args: UserA
//      UserB
//      5000
func FuncTransferMoney(args ...interface{}) (common.OperateResult, error) {
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

//打印凭条收据
//Args: UserA
//      UserB
//      5000
func FuncPrintReceipt(args ...interface{}) (common.OperateResult, error) {
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

//提示交房租日到，请给账户打钱
//Args: UserA
//      UserB
//      5000
func FuncRemindAccount(args ...interface{}) (common.OperateResult, error) {
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

//空动作
func FuncNoAction(args ...interface{}) (common.OperateResult, error) {
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
