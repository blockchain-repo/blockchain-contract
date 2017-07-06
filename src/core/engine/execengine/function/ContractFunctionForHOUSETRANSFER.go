package function

import (
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/common"
)

//++++++++++++++++++++++++++++++++++++++++++++++++++++++
//++++++++++合约机【房屋跨链交易】专用扩展方法++++++++++
//++++++++++++++++++++++++++++++++++++++++++++++++++++++
//查询要购买的房产 FuncQueryHouse()
//无合适房产，退出 FuncExitForNoHouse()
//查询账户余额 FuncUserBalance(user_A)
//购房款转账 FuncTransferHouseFees(user_A, user_B, amount)
//提示钱不足，需要充钱 FuncNoticeRecharge()
//查看购房款转账结果 FuncQueryHouseFeesResult()
//购房款转账失败，退出 FuncExitForTransferFail()
//等待用户充钱 FuncSleepTime()
//房产转移 FuncTransferHouse(user_B, user_A, amount)
//查询房产交易结果 FuncQueryHouseResult()
//购房交易成功 FuncExitForSuccess()
//交易失败购房款回退 FuncExitForHouseTransferFail(fees_transfer_id)

//查询要购买的房产
//Return: result int 0,无 1,有
func FuncQueryHouse(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	uniledgerlog.Warn("Business Operate.[Result] 查询到房产。")
	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(1)
	return v_result, v_err
}

//无合适房产，退出
func FuncExitForNoHouse(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	uniledgerlog.Warn("Business Operate.[Result] 未查询到合适房产，合约退出。")
	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(1)
	return v_result, v_err
}

//查询账户余额
//Args:   user_A string
//Return: balance float
func FuncUserBalance(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	var user_balance float32 = 100000000.0
	uniledgerlog.Warn("Business Operate.[Result]查询用户账户余额为", user_balance)
	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(user_balance)
	return v_result, v_err
}

//购房款转账
//Args: user_A string
//      user_B string
//      amount float32
func FuncTransferHouseFees(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	uniledgerlog.Warn("Business Operate.[Result]购房款转账成功", args[0], "转账给", args[1], "人民币", args[2], "元")
	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(1)
	return v_result, v_err
}

//提示钱不足，需要充钱
func FuncNoticeRecharge(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	uniledgerlog.Warn("Business Operate.[Result]账户余额不足，请充钱.")

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(1)
	return v_result, v_err
}

//查看购房款转账结果
func FuncQueryHouseFeesResult(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	uniledgerlog.Warn("Business Operate.[Result]查询购房款转账结果.")

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(1)
	return v_result, v_err
}

//购房款转账失败，退出
func FuncExitForTransferFail(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	uniledgerlog.Warn("Business Operate.[Result]购房款多次转账失败, 合约退出.")
	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(1)
	return v_result, v_err
}

//等待用户充钱 FuncSleepTime

//房产转移
//Args: user_B string
//      user_A string
//      amount float
func FuncTransferHouse(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	uniledgerlog.Warn("Business Operate.[Result]房产转移成功", args[0], "转移资产给", args[1], "房屋", args[2], "套.")

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(1)
	return v_result, v_err
}

//查询房产交易结果
func FuncQueryHouseResult(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	uniledgerlog.Warn("Business Operate.[Result]查询房产转移结果.")

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(1)
	return v_result, v_err
}

//购房交易成功
func FuncExitForHouseTransferSuccess(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	uniledgerlog.Warn("Business Operate.[Result]房产转移成功, 购房款支付成功, 合约执行完成.")

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(1)
	return v_result, v_err
}

//交易失败购房款回退
//Args: fees_transfer_id string
func FuncExitForHouseTransferFail(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	uniledgerlog.Warn("Business Operate.[Result]房产转移失败, 交付房款退回.")

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(1)
	return v_result, v_err
}
