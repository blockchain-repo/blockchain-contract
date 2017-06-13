package function

import "time"
import "unicontract/src/core/engine/common"

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//++++++++++合约机【TRANSFER】自动转账场景专用扩展方法++++++++++

//++++++++++++++++自动转账合约+++++++++++++++++++++++++++++++++++
//在指定时间，交易账户A 给交易账户B 转账500元
//Args: user_A  交易用户From
//      user_B  交易用户To
//      transfer_time 转账时间
//      amount  转账金额
func FuncAutoTransferAssetAtTime(args ...interface{}) (common.OperateResult, error) {
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
