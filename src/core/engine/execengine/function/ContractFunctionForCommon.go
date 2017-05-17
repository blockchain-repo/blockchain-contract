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

//查询冻结资产方法
func GetFreezeAsset() (common.OperateResult, error) {
	//userPubKey string, contractId string, taskId string, taskNum int
	var v_err error = nil

	v_result := common.OperateResult{}
	return v_result, v_err
}

//冻结资产方法
func FreezeAsset(userPubKey string, amount int, contractId string, taskId string, taskNum int) (common.OperateResult, error) {
	var v_err error = nil

	v_result := common.OperateResult{}
	return v_result, v_err
}

//资产转移方法
func TransferAsset(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult = common.OperateResult{}
	var v_err error = nil

	//var v_map_args map[string]interface{} = nil
	if len(args) != 4 {
		v_err = errors.New("param num error")
		return v_result, v_err
	}

	//user provide
	var ownerBefore string = args[0].(string)
	var recipients [][2]interface{} = [][2]interface{}{
		[2]interface{}{"5XAJvuRGb8B3hUesjREL7zdZ82ahZqHuBV6ttf3UEhyL", 100},
	}
	//executer provide
	var contractStr string = args[2].(string)
	var contractHashId string = args[3].(string)
	var contractId string = args[4].(string)
	var taskId string = args[5].(string)
	var taskIndex int = args[6].(int)
	var mainPubkey string = args[7].(string)
	var metadataStr string = ""
	//TODO generate by contractHashId contractId taskId taskIndex
	var relationStr string = GenerateRelation(contractHashId, contractId, taskId, taskIndex)

	//select freeze asset
	input, bal, spentFlag := transaction.GetFrozenUnspent(ownerBefore, contractId, taskId, taskIndex)
	inputs := common.Serialize(input)
	logs.Info(inputs)
	logs.Info(bal)

	if spentFlag == 0 || spentFlag == 2 {
		// no freeze asset or the freeze asset had be unfreezed
		mykey := config.Config.Keypair.PublicKey
		//var err error = nil
		//check main pubkey
		if mainPubkey == mykey {
			//if mainNode, do freeze;
			transaction.ExecuteTransfer("FREEZE", ownerBefore, recipients, metadataStr, relationStr, contractStr)
			//TODO if error do something
			//wait for the freeze asset write into the unichain
			time.Sleep(time.Second * 2)
		} else {
			// not mainNode, wait for the main node write the freeze-asset into the unchain
			time.Sleep(time.Second * 3)
		}

	} else if spentFlag == 3 {
		// The frozen asset had be transfered
		err := errors.New("The frozen asset had be transfered !")
		//TODO the transfer had write into the unchain, do nothing and return
		return v_result, err
	} else if spentFlag == 4 {
		//muti-frozen assets are found in unichian.
		err := errors.New("There are muti-frozen assets ,please check on !")
		//TODO should unfreeze all assets and then freeze only one asset
		return v_result, err
	}

	//todo if the freeze is not writen into the unchain, get the return and wait for another 3 seconds
	for i := 0; i <= 1; i++ {
		//transfer asset
		transaction.ExecuteTransfer("TRANSFER", ownerBefore, recipients, metadataStr, relationStr, contractStr)
		time.Sleep(time.Second * 3)
	}
	//todo make the result to return
	return v_result, v_err
}

//解冻资产方法
func UnfreezeAsset(args ...interface{}) (common.OperateResult, error) {
	//userPubKey string, contractId string, taskId string, taskNum int
	var v_err error = nil

	v_result := common.OperateResult{}
	return v_result, v_err
}

//根据合约ContractID查找合约
func GetContractById() (common.OperateResult, error) {
	//contractId string
	var v_err error = nil

	v_result := common.OperateResult{}
	return v_result, v_err
}

func GenerateRelation(contractHashId string, contractId string, taskId string, taskIndex int) string {
	//TODO generate relation with the execute strut

	return ""
}
