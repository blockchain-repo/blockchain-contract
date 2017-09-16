package function

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

	/*
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
	//FuncTransferAsset(args ...interface{}) (common.OperateResult, error) {

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	return v_result, v_err
}
