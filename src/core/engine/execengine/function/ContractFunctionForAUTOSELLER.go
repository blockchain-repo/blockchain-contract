package function

import "time"
import "unicontract/src/core/engine/common"

//++++++++++++++++++++++++++++++++++++++++++++++++++++++
//++++++++++合约机【自动售卖机】专用扩展方法++++++++++
//++++++++++++++++++++++++++++++++++++++++++++++++++++++
//获取用户输入的饮料种类FuncGetUserSelectedStyle()
//获取用户输入的购买数量FuncGetUserSelectedCount()
//查询购买饮料品种余量FuncQueryRemainingCount(style_A, count_A)
//退出操作，提供机器饮料不足FuncExitForNoRemaining(style_A, count_A)
//计算消耗总金额FuncCalculatedCost(style_A, count_A)
//休眠等待用户支付FuncWaitPayMoney(sleeptime)
//查询用户支付额度FuncQueryUserPayCount(user_A)
//售卖机出饮料FuncSupplyGoods(style_A, count_A)
//查询消耗余额FuncQueryRemainingMoney(user_A)
//售卖机找零FuncOddChange(user_A)
//退出操作，打印欢迎使用售卖机FuncExitForSuccess(user_A)
//购买流程结束FuncExitForTerminal(user_A)

//获取用户输入的饮料种类
func FuncGetUserSelectedStyle(args ...interface{}) (common.OperateResult, error) {
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

//获取用户输入的购买数量
func FuncGetUserSelectedCount(args ...interface{}) (common.OperateResult, error) {
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

//查询购买饮料品种余量
//Args: data_selected_style_A.Value
//      data_selected_count_B.Value
func FuncQueryRemainingCount(args ...interface{}) (common.OperateResult, error) {
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

//退出操作，提供机器饮料不足
//Args: data_selected_style_A.Value
//      data_selected_count_B.Value
func FuncExitForNoRemaining(args ...interface{}) (common.OperateResult, error) {
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

//计算消耗总金额
//Args: data_selected_style_A.Value
//      data_selected_count_B.Value)(
func FuncCalculatedCost(args ...interface{}) (common.OperateResult, error) {
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

//休眠等待用户支付
//Args: sleeptime
func FuncWaitPayMoney(args ...interface{}) (common.OperateResult, error) {
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

//查询用户支付额度
//Args: User_A
func FuncQueryUserPayCount(args ...interface{}) (common.OperateResult, error) {
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

//售卖机出饮料
//Args: data_selected_style_A.Value
//      data_selected_count_B.Value
func FuncSupplyGoods(args ...interface{}) (common.OperateResult, error) {
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

//查询消耗余额
//Args： User_A
func FuncQueryRemainingMoney(args ...interface{}) (common.OperateResult, error) {
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

//售卖机找零
//Args: User_A
func FuncOddChange(args ...interface{}) (common.OperateResult, error) {
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

//退出操作，打印欢迎使用售卖机
//用户在超时范围内没有支付，购买流程结束
//Args: User_A
func FuncExitForSuccess(args ...interface{}) (common.OperateResult, error) {
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

//用户在超时范围内没有支付，购买流程结束
//Args: User_A
func FuncExitForTerminal(args ...interface{}) (common.OperateResult, error) {
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
