package function

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"strconv"
	"unicontract/src/core/engine/common"
)

var (
	BANK_BALANCE      map[string]float64
	TELEPHONE_BALANCE map[string]float64
	TELEPHONE_CONSUME map[string]float64
)

func init() {
	BANK_BALANCE = make(map[string]float64, 0)
	BANK_BALANCE["AXXXXXXXXXXX"] = 300.0000
	BANK_BALANCE["BXXXXXXXXXXX"] = 100000.0000

	TELEPHONE_BALANCE = make(map[string]float64, 0)
	TELEPHONE_BALANCE["AXXXXXXXXXXX"] = 30.0000
	TELEPHONE_BALANCE["BXXXXXXXXXXX"] = 100000.0000

	TELEPHONE_CONSUME = make(map[string]float64, 0)
	TELEPHONE_CONSUME["AXXXXXXXXXXX"] = 60.0000
	TELEPHONE_CONSUME["BXXXXXXXXXXX"] = 0

}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++
//++++++++++合约机【DEMO】专用扩展方法++++++++++
//++++++++++++++++++++++++++++++++++++++++++++++++++++++
//获取用户余额FuncGetBalance(user_A)
//资金转账FuncTanferMoney(user_A, user_B, amount)
//账户充值FuncDeposit(user_A, amount)
//查询用户月消耗话费FuncQueryMonthConsumption(user_A)
//返还话费FuncBackTelephoneFare(user_B, user_A, amount)

//获取用户余额
func FuncGetBalance(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult = common.OperateResult{}
	var v_err error = nil
	if len(args) != 1 {
		v_result.SetCode(400)
		v_result.SetMessage("Need 1 Param, now is 0!")
		v_err = errors.New("Need 1 Param, now is 0!")
		return v_result, v_err
	}
	arg_0, ok := args[0].(string)
	if !ok {
		logs.Error("assert error")
		return v_result, fmt.Errorf("assert error")
	}
	balance_amount := BANK_BALANCE[arg_0]

	//构建返回值
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(balance_amount)
	return v_result, v_err
}

//资产转移
func FuncTanferMoney(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult = common.OperateResult{}
	var v_err error = nil
	if len(args) != 3 {
		v_result.SetCode(400)
		v_result.SetData(false)
		v_result.SetMessage("Need 1 Param, now is 0!")
		v_err = errors.New("Need 1 Param, now is 0!")
		return v_result, v_err
	}

	arg_A, ok := args[0].(string)
	if !ok {
		logs.Error("assert error")
		return v_result, fmt.Errorf("assert error")
	}
	arg_B, ok := args[1].(string)
	if !ok {
		logs.Error("assert error")
		return v_result, fmt.Errorf("assert error")
	}
	var arg_money float64 = args[2].(float64)

	if BANK_BALANCE[arg_A] < arg_money {
		v_result.SetCode(400)
		v_result.SetData(false)
		v_result.SetMessage("BANK_COUNT[" + arg_A + "] lt TransferMoney(" + strconv.FormatFloat(arg_money, 'f', 6, 64) + ")!")
		v_err = errors.New("Need 1 Param, now is 0!")
		return v_result, v_err
	}
	BANK_BALANCE[arg_A] = BANK_BALANCE[arg_A] - arg_money
	BANK_BALANCE[arg_B] = BANK_BALANCE[arg_B] + arg_money

	v_result.SetCode(200)
	v_result.SetData(true)
	v_result.SetMessage("process success!")
	return v_result, v_err
}

//账户充值
func FuncDeposit(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult = common.OperateResult{}
	var v_err error = nil
	if len(args) != 2 {
		v_result.SetCode(400)
		v_result.SetData(false)
		v_result.SetMessage("Need 1 Param, now is 0!")
		v_err = errors.New("Need 1 Param, now is 0!")
		return v_result, v_err
	}

	var arg_A string = args[0].(string)
	var arg_money float64 = args[1].(float64)

	BANK_BALANCE[arg_A] = BANK_BALANCE[arg_A] + arg_money

	v_result.SetCode(200)
	v_result.SetData(true)
	v_result.SetMessage("process success!")
	return v_result, v_err
}

//查询用户月消耗话费
func FuncQueryMonthConsumption(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult = common.OperateResult{}
	var v_err error = nil
	if len(args) != 1 {
		v_result.SetCode(400)
		v_result.SetMessage("Need 1 Param, now is 0!")
		v_err = errors.New("Need 1 Param, now is 0!")
		return v_result, v_err
	}

	var arg_A string = args[0].(string)

	v_result.SetCode(200)
	v_result.SetData(TELEPHONE_CONSUME[arg_A])
	v_result.SetMessage("process success!")
	return v_result, v_err
}

//返还话费
func FuncBackTelephoneFare(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult = common.OperateResult{}
	var v_err error = nil
	if len(args) != 3 {
		v_result.SetCode(400)
		v_result.SetData(false)
		v_result.SetMessage("Need 1 Param, now is 0!")
		v_err = errors.New("Need 1 Param, now is 0!")
		return v_result, v_err
	}
	var arg_B string = args[0].(string)
	var arg_A string = args[1].(string)
	var arg_money float64 = args[2].(float64)
	TELEPHONE_BALANCE[arg_B] = TELEPHONE_BALANCE[arg_B] - arg_money
	TELEPHONE_BALANCE[arg_A] = TELEPHONE_BALANCE[arg_A] + arg_money

	v_result.SetCode(200)
	v_result.SetData(true)
	v_result.SetMessage("process success!")
	return v_result, v_err
}
