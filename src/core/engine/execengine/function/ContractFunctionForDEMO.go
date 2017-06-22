package function

import (
	"errors"
	"strconv"
	"time"
	"unicontract/src/core/engine/common"
	"github.com/astaxie/beego/logs"
	"fmt"
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

//获取当前日期的Day int
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

func FuncGetBalance(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult = common.OperateResult{}
	var v_err error = nil
	if len(args) != 1 {
		v_result.SetCode(400)
		v_result.SetMessage("Need 1 Param, now is 0!")
		v_err = errors.New("Need 1 Param, now is 0!")
		return v_result, v_err
	}
	arg_0,ok := args[0].(string)
	if !ok{
		logs.Error("assert error")
		return v_result,fmt.Errorf("assert error")
	}
	balance_amount := BANK_BALANCE[arg_0]

	//构建返回值
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(balance_amount)
	return v_result, v_err
}

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

	arg_A,ok := args[0].(string)
	if !ok{
		logs.Error("assert error")
		return v_result,fmt.Errorf("assert error")
	}
	arg_B,ok := args[1].(string)
	if !ok{
		logs.Error("assert error")
		return v_result,fmt.Errorf("assert error")
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
