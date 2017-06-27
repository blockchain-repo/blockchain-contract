package transaction

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	//data := strings
	//a := StringSlice(data[0:])
	//Sort(a)
	//var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}

	var silce = []string{"a", "c", "b", "A", "C", "B", "1", "3", "2"}
	sort.Strings(silce)
	//sliceSort := sort.StringSlice(silce)
	logs.Info(silce)
	//logs.Info(sliceSort)
}

func TestTransferAssetComplete(t *testing.T) {
	//TransferAssetComplete("", ",")
}

func Test_GetInterestCount(t *testing.T) {
	GetInterestCount("key...")
}

func Test_GetPurchaseAmount(t *testing.T) {
	logs.Info(GetPurchaseAmount("key..."))
}

func Test_SaveEarnings(t *testing.T) {
	//SaveEarnings("key..." , 22 , 0.03 , 3 , false)
}

func TestCon(t *testing.T) {
	contract := "{\"id\":\"f47d70d931c1001e3e105bb30ae3c6e8da17a0653f7a29516b5921b3a689f5df\",\"ContractHead\":{\"AssignTime\":\"1498586538948\",\"MainPubkey\":\"F7wLyzThjzRH3rruP72iV9mikjWCrznAvC7pWW4r51rz\",\"OperateTime\":\"\",\"Version\":1,\"ConsensusResult\":1},\"ContractBody\":{\"ContractId\":\"170628020217127631\",\"Cname\":\"contract_transfer\",\"Ctype\":\"Component_Contract\",\"Caption\":\"单次转账合约\",\"Description\":\"在指定时间内自动完成转账给某人\",\"ContractState\":\"Contract_Completed\",\"Creator\":\"6p7waxWGKDYKDJPve4v5oQyFV9Sj2a8Zrw6EHVEHZhGu\",\"CreateTime\":\"1498579155000\",\"StartTime\":\"1497937998000\",\"EndTime\":\"1498802001000\",\"ContractOwners\":[\"6p7waxWGKDYKDJPve4v5oQyFV9Sj2a8Zrw6EHVEHZhGu\"],\"ContractAssets\":[{\"AssetId\":\"91327563-2378-4b4e-865d-71f5ee3d8921\",\"Name\":\"asset_transfer\",\"Caption\":\"\",\"Description\":\"自动转账的资金\",\"Unit\":\"元\",\"Amount\":0,\"MetaData\":{\"\":\"\"}}],\"ContractSignatures\":[{\"OwnerPubkey\":\"6p7waxWGKDYKDJPve4v5oQyFV9Sj2a8Zrw6EHVEHZhGu\",\"Signature\":\"3VAt1Sc8k2fpH4C7efodgwR7u8cqQbQxNMRRWjm1aEMR918PwPRHjaJJTWxVAJ2vp1YA6NCyZHXxoh7zbdWE44F6\",\"SignTimestamp\":\"1498586537912\"}],\"ContractComponents\":[{\"Caption\":\"自动转账\",\"Cname\":\"action_transfer\",\"CompleteCondition\":[{\"Caption\":\"\",\"Cname\":\"expression_complete\",\"Ctype\":\"Component_Expression.Expression_Condition\",\"Description\":\"\",\"ExpressionResult\":{\"Code\":0,\"Data\":\"\",\"Message\":\"\",\"Output\":\"\"},\"ExpressionStr\":\"true\",\"LogicValue\":0,\"MetaAttribute\":{}}],\"Ctype\":\"Component_Task.Task_Action\",\"DataList\":[{\"Caption\":\"\",\"Category\":[],\"Cname\":\"data_action\",\"Ctype\":\"Component_Data\",\"DefaultValue\":\"\",\"Description\":\"\",\"HardConvType\":\"Data_OperateResultData\",\"Mandatory\":false,\"MetaAttribute\":{},\"ModifyDate\":\"\",\"Options\":{},\"Parent\":{\"Caption\":\"\",\"Category\":null,\"Cname\":\"\",\"Ctype\":\"\",\"DataRangeFloat\":null,\"DataRangeInt\":null,\"DataRangeUint\":null,\"DefaultValueFloat\":0,\"DefaultValueInt\":0,\"DefaultValueString\":\"\",\"DefaultValueUint\":0,\"Description\":\"\",\"Format\":\"\",\"HardConvType\":\"\",\"Mandatory\":false,\"ModifyDate\":\"\",\"Options\":null,\"Unit\":\"\",\"ValueFloat\":0,\"ValueInt\":0,\"ValueString\":\"\",\"ValueUint\":0},\"Unit\":\"\",\"Value\":\"\"}],\"DataValueSetterExpressionList\":[{\"Caption\":\"\",\"Cname\":\"expression_function_action\",\"Ctype\":\"Component_Expression.Expression_Function\",\"Description\":\"转账\",\"ExpressionResult\":{\"Code\":0,\"Data\":\"\",\"Message\":\"\",\"Output\":\"\"},\"ExpressionStr\":\"FuncTransferAsset(contract_transfer.ContractBody.ContractOwners.0,contract_transfer.ContractBody.MetaAttribute.TransferTo,contract_transfer.ContractBody.MetaAttribute.TransferAmount)\",\"MetaAttribute\":{}}],\"Description\":\"达到指定转账时间，自动转账给某人\",\"DiscardCondition\":[{\"Caption\":\"\",\"Cname\":\"expression_discard\",\"Ctype\":\"Component_Expression.Expression_Condition\",\"Description\":\"\",\"ExpressionResult\":{\"Code\":0,\"Data\":\"\",\"Message\":\"\",\"Output\":\"\"},\"ExpressionStr\":\"true\",\"LogicValue\":0,\"MetaAttribute\":{}}],\"MetaAttribute\":{},\"NextTasks\":[],\"PreCondition\":[{\"Caption\":\"\",\"Cname\":\"expression_pre_transfer\",\"Ctype\":\"Component_Expression.Expression_Condition\",\"Description\":\"在指定时间转账\",\"ExpressionResult\":{\"Code\":0,\"Data\":\"\",\"Message\":\"\",\"Output\":\"\"},\"ExpressionStr\":\"data_date_expression_function_nowdate.Value \u003e=contract_transfer.ContractBody.MetaAttribute.TransferDate\",\"LogicValue\":0,\"MetaAttribute\":{}}],\"SelectBranches\":[],\"State\":\"TaskState_Dormant\",\"TaskExecuteIdx\":0,\"TaskId\":\"b8ef012f-031b-43ce-a0b4-9f9989b082a3\"},{\"Caption\":\"不达到转账日期，休眠5s\",\"Cname\":\"task_action_sleep\",\"CompleteCondition\":[],\"Ctype\":\"Component_Task.Task_Action\",\"DataList\":[],\"DataValueSetterExpressionList\":[{\"Caption\":\"\",\"Cname\":\"expression_function_sleep\",\"Ctype\":\"Component_Expression.Expression_Function\",\"Description\":\"\",\"ExpressionResult\":{\"Code\":0,\"Data\":\"\",\"Message\":\"\",\"Output\":\"\"},\"ExpressionStr\":\"FuncSleepTime(5)\",\"MetaAttribute\":{}}],\"Description\":\"判定当前时间是否达到转账日期，没有达到，需要休眠5是等待\",\"DiscardCondition\":[],\"MetaAttribute\":{},\"NextTasks\":[\"task_enquiry_nowdate\"],\"PreCondition\":[{\"Caption\":\"\",\"Cname\":\"expression_condition_pre_sleep\",\"Ctype\":\"Component_Expression.Expression_Condition\",\"Description\":\"\",\"ExpressionResult\":{\"Code\":0,\"Data\":\"\",\"Message\":\"\",\"Output\":\"\"},\"ExpressionStr\":\"data_date_expression_function_nowdate.Value \u003c contract_transfer.ContractBody.MetaAttribute.TransferDate \",\"LogicValue\":0,\"MetaAttribute\":{}}],\"SelectBranches\":[],\"State\":\"TaskState_Dormant\",\"TaskExecuteIdx\":0,\"TaskId\":\"c18ad533-fcc6-4e5a-84b8-9a86a97a7b0c\"},{\"Caption\":\"查询当前日期\",\"Cname\":\"task_enquiry_nowdate\",\"CompleteCondition\":[],\"Ctype\":\"Component_Task.Task_Enquiry\",\"DataList\":[{\"Caption\":\"\",\"Category\":[],\"Cname\":\"data_date_expression_function_nowdate\",\"Ctype\":\"Component_Data.Data_Date\",\"DefaultValue\":\"\",\"Description\":\"\",\"Format\":\"2006-01-02 15:04:05\",\"HardConvType\":\"strToDate\",\"Mandatory\":false,\"MetaAttribute\":{},\"ModifyDate\":\"\",\"Options\":{},\"Parent\":{\"Caption\":\"\",\"Category\":null,\"Cname\":\"\",\"Ctype\":\"\",\"DataRangeFloat\":null,\"DataRangeInt\":null,\"DataRangeUint\":null,\"DefaultValueFloat\":0,\"DefaultValueInt\":0,\"DefaultValueString\":\"\",\"DefaultValueUint\":0,\"Description\":\"\",\"Format\":\"\",\"HardConvType\":\"\",\"Mandatory\":false,\"ModifyDate\":\"\",\"Options\":null,\"Unit\":\"\",\"ValueFloat\":0,\"ValueInt\":0,\"ValueString\":\"\",\"ValueUint\":0},\"Unit\":\"\",\"Value\":\"1498586542764\"}],\"DataValueSetterExpressionList\":[{\"Caption\":\"\",\"Cname\":\"expression_function_nowdate\",\"Ctype\":\"Component_Expression.Expression_Function\",\"Description\":\"\",\"ExpressionResult\":{\"Code\":200,\"Data\":\"1498586542764\",\"Message\":\"process success!\",\"Output\":\"\"},\"ExpressionStr\":\"FuncGetNowDateTimestamp()\",\"MetaAttribute\":{}}],\"Description\":\"获取当期日期，用于判定是否达到转账日期\",\"DiscardCondition\":[],\"MetaAttribute\":{},\"NextTasks\":[\"action_transfer\",\"task_action_sleep\"],\"PreCondition\":[{\"Caption\":\"\",\"Cname\":\"expression_condition_pre_nowdate\",\"Ctype\":\"Component_Expression.Expression_Condition\",\"Description\":\"\",\"ExpressionResult\":{\"Code\":0,\"Data\":\"\",\"Message\":\"\",\"Output\":\"\"},\"ExpressionStr\":\"true\",\"LogicValue\":0,\"MetaAttribute\":{}}],\"SelectBranches\":[],\"State\":\"TaskState_Completed\",\"TaskExecuteIdx\":0,\"TaskId\":\"ada66d9d-a21c-4fea-8b4d-8f44c9c368ac\"}],\"NextTasks\":[\"task_enquiry_nowdate\"],\"MetaAttribute\":{\"Copyright\":\"uni-ledger\",\"TransferAmount\":\"100\",\"TransferDate\":\"1498586531000\",\"TransferTo\":\"F2P8cmiNbzr79QserzAh2LktZLdR6AgnNRfjQd6eMbB9\",\"Version\":\"v1.0\",\"_UCVM_CopyRight\":\"uni-ledger\",\"_UCVM_Date\":\"2017-06-01 12:00:00\",\"_UCVM_Version\":\"v1.0\"}}}"
	v_model, v_err := GenContractByExecStr(contract)
	fmt.Println("1111:", v_model, v_err)

	//contract = strings.TrimLeft(contract, "\"")
	//contract = strings.TrimRight(contract, "\"")
	//fmt.Println(contract)

	v_model1, v_err1 := GenContractByExecStr(contract)
	fmt.Println("2222:", v_model1, v_err1)
}
