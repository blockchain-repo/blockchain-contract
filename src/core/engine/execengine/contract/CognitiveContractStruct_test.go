package contract

import (
	"encoding/json"
	"fmt"
	"testing"
)

func PrintContrract(v_contract CognitiveContract) {
	fmt.Println("=======constract init=========")
	v_contract.InitCognitiveContract()
	fmt.Println("=======constract print=========")
	fmt.Println("Cname: ", v_contract.ContractBody.Cname)
	fmt.Println("  PropertyTable[_Cname]: ", v_contract.GetPropertyTable()["_Cname"])
	fmt.Println("Ctype: ", v_contract.ContractBody.Ctype)
	fmt.Println("  PropertyTable[_Ctype]: ", v_contract.GetPropertyTable()["_Ctype"])
	fmt.Println("Caption: ", v_contract.ContractBody.Caption)
	fmt.Println("  PropertyTable[_Caption]: ", v_contract.GetPropertyTable()["_Caption"])
	fmt.Println("Description: ", v_contract.ContractBody.Description)
	fmt.Println("  PropertyTable[_Description]: ", v_contract.GetPropertyTable()["_Description"])
	fmt.Println("Creator: ", v_contract.ContractBody.Creator)
	fmt.Println("  PropertyTable[_Creator]: ", v_contract.GetPropertyTable()["_Creator"])
	fmt.Println("CreateTime: ", v_contract.ContractBody.CreateTime)
	fmt.Println("  PropertyTable[_CreateTime]: ", v_contract.GetPropertyTable()["_CreateTime"])
	fmt.Println("StartTime: ", v_contract.ContractBody.StartTime)
	fmt.Println("  PropertyTable[_StartTime]: ", v_contract.GetPropertyTable()["_StartTime"])
	fmt.Println("EndTime: ", v_contract.ContractBody.EndTime)
	fmt.Println("  PropertyTable[_EndTime]: ", v_contract.GetPropertyTable()["_EndTime"])
	fmt.Println("ContractOwners: ", v_contract.ContractBody.ContractOwners)
	fmt.Println("  PropertyTable[_ContractOwners]: ", v_contract.GetPropertyTable()["_ContractOwners"])
	fmt.Println("  All Owners: ")
	for p_idx, p_owner := range v_contract.ContractBody.ContractOwners {
		fmt.Println("  owner[", p_idx, "]: ", p_owner)
	}
	fmt.Println("")
	fmt.Println("ContractAssets: ", v_contract.ContractBody.ContractAssets)
	fmt.Println("  PropertyTable[_ContractAssets]: ", v_contract.GetPropertyTable()["_ContractAssets"])
	fmt.Println("  All Assets: ")
	for p_idx, p_assert := range v_contract.ContractBody.ContractAssets {
		fmt.Println("  Asset.AssetId[", p_idx, "]: ", p_assert.AssetId)
		fmt.Println("  Asset.Name[", p_idx, "]: ", p_assert.Name)
		fmt.Println("  Asset.Caption[", p_idx, "]: ", p_assert.Caption)
		fmt.Println("  Asset.Description[", p_idx, "]: ", p_assert.Description)
		fmt.Println("  Asset.Unit[", p_idx, "]: ", p_assert.Unit)
		fmt.Println("  Asset.Amount[", p_idx, "]: ", p_assert.Amount)
		fmt.Println("  Asset.MetaData[", p_idx, "]: ", p_assert.MetaData)
	}
	fmt.Println("ContractSignatures: ", v_contract.ContractBody.ContractSignatures)
	fmt.Println("  PropertyTable[_ContractSignatures]: ", v_contract.GetPropertyTable()["_ContractSignatures"])
	fmt.Println("  All Signatures: ")
	for p_idx, p_signature := range v_contract.ContractBody.ContractSignatures {
		fmt.Println("  Signatures.OwnerPubkey[", p_idx, "]: ", p_signature.OwnerPubkey)
		fmt.Println("  Signatures.Signature[", p_idx, "]: ", p_signature.Signature)
		fmt.Println("  Signatures.SignTimestamp[", p_idx, "]: ", p_signature.SignTimestamp)
	}
	fmt.Println("MetaAttribute: ", v_contract.ContractBody.MetaAttribute)
	fmt.Println("  PropertyTable[_MetaAttribute]: ", v_contract.GetPropertyTable()["_MetaAttribute"])
	fmt.Println("  All MetaAttribute: ")
	for p_key, p_value := range v_contract.ContractBody.MetaAttribute {
		fmt.Println("  Attribute[", p_key, "]", p_value)
	}
	fmt.Println("ContractComponent: ", v_contract.ContractBody.ContractComponents)
}

/*
func TestContractStruct_simple(t *testing.T)  {
	str_contract :=  `{
"ContractId":"xxxxxxxxxxxxxxxxxx",
"Cname":"contract_mobilecallback",
"Caption":"购智能手机返话费合约产品协议",
"Description":"移动用户A花费500元购买移动运营商B的提供的合约智能手机C后，要求用户每月消费58元以上通信费，移动运营商B便可按月返还话费（即：每月1号返还移动用户A20元话费），连续返还12个月",
"Creator":"ABCDEFGHIJKLMNIPQRSTUVWXYZ",
"CreateTime":"2016-12-20 12:00:00",
"StartTime":"2017-01-01 12:00:00",
"EndTime":"2017-01-01 12:00:00",
"ContractOwners":[
"AXXXXXXXXXXX",
"BXXXXXXXXXXX"
],
"ContractAssets":[
{
"AssetId":"xxxxxxxxxxx",
"Name":"asset_money",
"Caption":"理财产品",
"Description":"理财资产",
"Unit":"份",
"Amount":1000,
"MetaData":{
"TestAsset1":"1111111111",
"TestAsset2":"2222222222"}
}
],
"ContractSignatures":[
{
"OwnerPubkey":"AXXXXXXXXXXX",
"Signature":"Axxxxxxxxxxxxxxxxxxxxxx",
"SignTimestamp":"1492619683"
},
{
"OwnerPubkey":"BXXXXXXXXXXX",
"Signature":"Bxxxxxxxxxxxxxxxxxxxxxx",
"SignTimestamp":"1492619983"
}
],
"MetaAttribute": {
  "Test1":"aaaaaa",
  "Test2":"bbbbbb"
}}`
	fmt.Println("====================================")
	var v_contract.ContractBody CognitiveContract = *new(CognitiveContract)
	if err := json.Unmarshal([]byte(str_contract), &v_contract.ContractBody); err == nil {
		fmt.Println("=======constract object=========")
		fmt.Println(v_contract.ContractBody)
	}else {
		fmt.Println(err)
		return
	}
	fmt.Println("=======constract init=========")
	v_contract.ContractBody.InitCognitiveContract()
	fmt.Println("=======constract print=========")
	fmt.Println("Cname: ", v_contract.ContractBody.Cname)
	fmt.Println("  PropertyTable[_Cname]: ", v_contract.ContractBody.GetPropertyTable()["_Cname"])
	fmt.Println("Ctype: ", v_contract.ContractBody.Ctype)
	fmt.Println("  PropertyTable[_Ctype]: ", v_contract.ContractBody.GetPropertyTable()["_Ctype"])
	fmt.Println("Caption: ", v_contract.ContractBody.Caption)
	fmt.Println("  PropertyTable[_Caption]: ", v_contract.ContractBody.GetPropertyTable()["_Caption"])
	fmt.Println("Description: ", v_contract.ContractBody.Description)
	fmt.Println("  PropertyTable[_Description]: ", v_contract.ContractBody.GetPropertyTable()["_Description"])
	fmt.Println("Creator: ", v_contract.ContractBody.Creator)
	fmt.Println("  PropertyTable[_Creator]: ", v_contract.ContractBody.GetPropertyTable()["_Creator"])
	fmt.Println("CreateTime: ", v_contract.ContractBody.CreateTime)
	fmt.Println("  PropertyTable[_CreateTime]: ", v_contract.ContractBody.GetPropertyTable()["_CreateTime"])
	fmt.Println("StartTime: ", v_contract.ContractBody.StartTime)
	fmt.Println("  PropertyTable[_StartTime]: ", v_contract.ContractBody.GetPropertyTable()["_StartTime"])
	fmt.Println("EndTime: ", v_contract.ContractBody.EndTime)
	fmt.Println("  PropertyTable[_EndTime]: ", v_contract.ContractBody.GetPropertyTable()["_EndTime"])
	fmt.Println("ContractOwners: ", v_contract.ContractBody.ContractOwners)
	fmt.Println("  PropertyTable[_ContractOwners]: ", v_contract.ContractBody.GetPropertyTable()["_ContractOwners"])
	fmt.Println("  All Owners: ")
	for p_idx,p_owner := range v_contract.ContractBody.ContractOwners {
		fmt.Println("  owner[",p_idx, "]: ", p_owner)
	}
	fmt.Println("")
	fmt.Println("ContractAssets: ", v_contract.ContractBody.ContractAssets)
	fmt.Println("  PropertyTable[_ContractAssets]: ", v_contract.ContractBody.GetPropertyTable()["_ContractAssets"])
	fmt.Println("  All Assets: ")
	for p_idx,p_assert := range v_contract.ContractBody.ContractAssets {
		fmt.Println("  Asset.AssetId[",p_idx, "]: ", p_assert.AssetId)
		fmt.Println("  Asset.Name[",p_idx, "]: ", p_assert.Name)
		fmt.Println("  Asset.Caption[",p_idx, "]: ", p_assert.Caption)
		fmt.Println("  Asset.Description[",p_idx, "]: ", p_assert.Description)
		fmt.Println("  Asset.Unit[",p_idx, "]: ", p_assert.Unit)
		fmt.Println("  Asset.Amount[",p_idx, "]: ", p_assert.Amount)
		fmt.Println("  Asset.MetaData[",p_idx, "]: ", p_assert.MetaData)
	}
	fmt.Println("ContractSignatures: ", v_contract.ContractBody.ContractSignatures)
	fmt.Println("  PropertyTable[_ContractSignatures]: ", v_contract.ContractBody.GetPropertyTable()["_ContractSignatures"])
	fmt.Println("  All Signatures: ")
	for p_idx,p_signature := range v_contract.ContractBody.ContractSignatures {
		fmt.Println("  Signatures.OwnerPubkey[",p_idx, "]: ", p_signature.OwnerPubkey)
		fmt.Println("  Signatures.Signature[",p_idx, "]: ", p_signature.Signature)
		fmt.Println("  Signatures.SignTimestamp[",p_idx, "]: ", p_signature.SignTimestamp)
	}
	fmt.Println("MetaAttribute: ", v_contract.ContractBody.MetaAttribute)
	fmt.Println("  PropertyTable[_MetaAttribute]: ", v_contract.ContractBody.GetPropertyTable()["_MetaAttribute"])
	fmt.Println("  All MetaAttribute: ")
	for p_key,p_value := range v_contract.ContractBody.MetaAttribute {
		fmt.Println("  Attribute[", p_key, "]", p_value)
	}
}*/

func TestContractStruct_all(t *testing.T) {

	str_contract := `{
"ContractId":"xxxxxxxxxxxxxxxxxxxxx",
"Cname":"contract_mobilecallback",
"Caption":"购智能手机返话费合约产品协议",
"Description":"移动用户A花费500元购买移动运营商B的提供的合约智能手机C后，要求用户每月消费58元以上通信费，移动运营商B便可按月返还话费（即：每月1号返还移动用户A20元话费），连续返还12个月",
"MetaAttribute":{
"Version":"v1.0",
"Copyright":"uni-ledger"
},
"Creator":"ABCDEFGHIJKLMNIPQRSTUVWXYZ",
"CreateTime":"2016-12-20 12:00:00",
"StartTime":"2017-01-01 12:00:00",
"EndTime":"2017-01-01 12:00:00",
"ContractOwners":[
"AXXXXXXXXXXX",
"BXXXXXXXXXXX"
],
"ContractAssets":[
{
"AssetId":"xxxxxxxxxxx",
"Name":"asset_money",
"Caption":"理财产品",
"Description":"理财资产",
"Unit":"份",
"Amount":1000,
"MetaData":{
"TestAsset1":"1111111111",
"TestAsset2":"2222222222"}
}
],
"ContractSignatures":[
{
"OwnerPubkey":"AXXXXXXXXXXX",
"Signature":"Axxxxxxxxxxxxxxxxxxxxxx",
"SignTimestamp":"1492619683"
},
{
"OwnerPubkey":"BXXXXXXXXXXX",
"Signature":"Bxxxxxxxxxxxxxxxxxxxxxx",
"SignTimestamp":"1492619983"
}
],
"MetaAttribute": {
  "Test1":"aaaaaa",
  "Test2":"bbbbbb"
},
"ContractComponents":[
{
"Cname":"enquiry_A",
"Caption":"查询用户账户",
"Description":"查询移动用户A账户是否有500元",
"Ctype":"TASK_ENQUIRY",
"State":"Dromant",
"PreCondition":[{
"Cname":"expression_precond_A",
"Caption":"查询用户账户前提条件",
"Description":"当前日期大于等于合约生效起始日期",
"Ctype":"Expression_LogicArgument",
"ExpressionStr":"contract_mobilecallback.StartTime()>=\"2017-01-0112:00:00\"&&contract_mobilecallback.EndTime()<=\"2017-12-3123:59:59\"",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}],
"DisgardCondition":[{
"Cname":"expression_discond_A",
"Caption":"",
"Description":"",
"Ctype":"Expression_LogicArgument",
"ExpressionStr":"true",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}],
"CompleteCondition":[{
"Cname":"expression_comcond_A",
"Caption":"",
"Description":"",
"Ctype":"Expression_LogicArgument",
"ExpressionStr":"expression_data_A.ExpressionResult['code']==200",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}],
"NextTasks":[
"action_B",
"action_C"
],
"DataValueSetterExpressionList":[
{
"Cname":"expression_data_A",
"Caption":"",
"Description":"",
"Ctype":"Expression_Function",
"ExpressionStr":"getBalance(AXXXXXXXXXX)",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}
],
"DataList":[
{
"Cname":"data_A",
"Caption":"",
"Description":"",
"Ctype":"Data_Float",
"Unit":"元",
"Value":"600",
"DefaultValue":"0",
"HardConvType":"float64"
}
]
},
{
"Cname":"action_B",
"Caption":"A购置手机",
"Description":"移动用户A账户存在500元：用户A将500元转账给移动运营商B，运营商B将手机快递给用户A",
"Ctype":"TASK_ACTION",
"State":"Dromant",
"PreCondition":[{
"Cname":"expression_precond_B",
"Caption":"",
"Description":"",
"Ctype":"Expression_LogicArgument",
"ExpressionStr":"contract_mobilecallback.StartTime()>=\"2017-01-0112:00:00\"&&contract_mobilecallback.EndTime()<=\"2017-12-3123:59:59\"&&enquiry_A.data_A>500",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}],
"DisgardCondition":[{
"Cname":"expression_discond_B",
"Caption":"",
"Description":"",
"Ctype":"Expression_LogicArgument",
"ExpressionStr":"true",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}],
"CompleteCondition":[{
"Cname":"expression_comcond_B",
"Caption":"",
"Description":"",
"Ctype":"Expression_LogicArgument",
"ExpressionStr":"expression_data_B.ExpressionResult['code']==200",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}],
"NextTasks":[
"enquiry_D"
],
"DataValueSetterExpressionList":[
{
"Cname":"expression_data_B",
"Caption":"A转账500给B",
"Description":"用户A转账500元给移动运营商B(运营商B将手机快递给用户A,不在线上确认)",
"Ctype":"Expression_Function",
"ExpressionStr":"tranferMoney(AXXXXXXXXXX,BXXXXXXXXXX,500)",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}
],
"DataList":[]
},
{
"Cname":"action_C",
"Caption":"A账户存款500元",
"Description":"移动用户A账户不存在500元：用户Ａ往账户存入500元，然后将500元转账给移动运营商B",
"Ctype":"TASK_ACTION",
"State":"Dromant",
"PreCondition":[{
"Cname":"expression_precond_C",
"Caption":"",
"Description":"",
"Ctype":"Exporess_LogicArgument",
"ExpressionStr":"contract_mobilecallback.StartTime()>=\"2017-01-0112:00:00\"&&contract_mobilecallback.EndTime()<=\"2017-12-3123:59:59\"&&enquiry_A.data_A<500",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}],
"DisgardCondition":[{
"Cname":"expression_discond_C",
"Caption":"",
"Description":"",
"Ctype":"Exporess_LogicArgument",
"ExpressionStr":"true",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}],
"CompleteCondition":[{
"Cname":"expression_comcond_C",
"Caption":"",
"Description":"",
"Ctype":"Exporess_LogicArgument",
"ExpressionStr":"expression_data_C.ExpressionResult['code']==200",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}],
"NextTasks":[
"action_B"
],
"DataValueSetterExpressionList":[
{
"Cname":"expression_data_C",
"Caption":"",
"Description":"",
"Ctype":"Exporess_Function",
"ExpressionStr":"deposit(AXXXXXXXXXX)",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}
],
"DataList":[]
},
{
"Cname":"enquiry_D",
"Caption":"查询用户月消费额",
"Description":"查询用户A当月消费额度，是否大于58元",
"Ctype":"TASK_ENQUIRY",
"State":"Dromant",
"PreCondition":[{
"Cname":"expression_precond_D",
"Caption":"",
"Description":"",
"Ctype":"Expression_LogicArgument",
"ExpressionStr":"contract_mobilecallback.StartTime()>=\"2017-01-0112:00:00\"&&contract_mobilecallback.EndTime()<=\"2017-12-3123:59:59\"",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}],
"DisgardCondition":[{
"Cname":"expression_discond_D",
"Caption":"",
"Description":"",
"Ctype":"Expression_LogicArgument",
"ExpressionStr":"true",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}],
"CompleteCondition":[{
"Cname":"expression_comcond_D",
"Caption":"",
"Description":"",
"Ctype":"Expression_LogicArgument",
"ExpressionStr":"expression_data_D.ExpressionResult['code']==200",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}],
"NextTasks":[
"action_E",
"actionF"
],
"DataValueSetterExpressionList":[
{
"Cname":"expression_data_D",
"Caption":"",
"Description":"",
"Ctype":"Expression_Function",
"ExpressionStr":"queryMonthConsumption(AXXXXXXXXXX)",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}
],
"DataList":[
{
"Cname":"data_D",
"Caption":"用户A月消费额",
"Description":"",
"Ctype":"Data_Float",
"Unit":"元",
"Value":"80",
"DefaultValue":"0",
"HardConvType":"float64"
}
]
},
{
"Cname":"action_E",
"Caption":"移动运营商下月1号返移动用户A20元",
"Description":"移动用户A当月消费58元以上：移动运营商B下月1号返还话费20元；连续返还12个月",
"Ctype":"TASK_ACTION",
"State":"Dromant",
"PreCondition":[{
"Cname":"expression_precond_E",
"Caption":"",
"Description":"",
"Ctype":"Expression_LogicArgument",
"ExpressionStr":"contract_mobilecallback.StartTime()>=\"2017-01-0112:00:00\"&&contract_mobilecallback.EndTime()<=\"2017-12-3123:59:59\"&&getNowDay()==1&&getNowDate()!=action_E.data_E&&enquiry_D.data_D>=58",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}],
"DisgardCondition":[{
"Cname":"expression_discond_E",
"Caption":"",
"Description":"",
"Ctype":"Expression_LogicArgument",
"ExpressionStr":"true",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}],
"CompleteCondition":[{
"Cname":"expression_comcond_E",
"Caption":"",
"Description":"",
"Ctype":"Expression_LogicArgument",
"ExpressionStr":"expression_data_E.ExpressionResult['code']==200",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}],
"NextTasks":[
"enquiry_D"
],
"DataValueSetterExpressionList":[
{
"Cname":"expression_data_E",
"Caption":"B返话费给A20元",
"Description":"",
"Ctype":"Expression_Function",
"ExpressionStr":"backTelephoneFare(BXXXXXXXXXX,AXXXXXXXXXX,20)",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}
],
"DataList":[
{
"Cname":"data_E",
"Caption":"B返话费给A日期",
"Description":"",
"Ctype":"Data_Date",
"Unit":"",
"Value":"2017-02-0112:00:00",
"DefaultValue":"",
"HardConvType":"string"
}
]
},
{
"Cname":"action_F",
"Caption":"移动运行商不返话费",
"Description":"移动用户A当月消费不足58元：移动运营商B下月1号不返还话费",
"Ctype":"TASK_ACTION",
"State":"Dromant",
"PreCondition":[{
"Cname":"expression_precond_F",
"Caption":"",
"Description":"",
"Ctype":"Expression_LogicArgument",
"ExpressionStr":"contract_mobilecallback.StartTime()>=\"2017-01-0112:00:00\"&&contract_mobilecallback.EndTime()<=\"2017-12-3123:59:59\"&&getNowDay()==1&&getNowDate()!=action_E.data_E&&enquiry_D.data_D<58",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}],
"DisgardCondition":[{
"Cname":"expression_discond_F",
"Caption":"",
"Description":"",
"Ctype":"Expression_LogicArgument",
"ExpressionStr":"true",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}],
"CompleteCondition":[{
"Cname":"expression_discond_F",
"Caption":"",
"Description":"",
"Ctype":"Expression_LogicArgument",
"ExpressionStr":"expression_data_F.ExpressionResult['code']==200",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}],
"NextTasks":[
"enquiry_D"
],
"DataValueSetterExpressionList":[
{
"Cname":"expression_data_F",
"Caption":"不执行动作",
"Description":"消费不足58元，不执行动作",
"Ctype":"Expression_Function",
"ExpressionStr":"true",
"ExpressionResult":{
"Message":"Operatesuccess.",
"Code":200
},
"LogicValue":"1"
}
],
"DataList":[
{
"Cname":"data_F",
"Caption":"B返话费给A执行日期",
"Description":"移动运营商B返话费给用户A的操作日期",
"Ctype":"Data_Date",
"Unit":"",
"Value":"2017-02-0112:00:00",
"DefaultValue":"",
"HardConvType":"string"
}
]
}
]
}
`
	fmt.Println("+++++++++++manual Unmarshal++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	var v_contract CognitiveContract = *new(CognitiveContract)
	if err := json.Unmarshal([]byte(str_contract), &v_contract.ContractBody); err == nil {
		fmt.Println("=======constract object=========")
		fmt.Println(v_contract.ContractBody)
	} else {
		fmt.Println(err)
		return
	}
	PrintContrract(v_contract)

	fmt.Println("+++++++++++Deserialize method+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	v2_contract := new(CognitiveContract)
	tmp_contract, err := v2_contract.Deserialize(str_contract)
	v2_contract = tmp_contract.(*CognitiveContract)
	if err != nil {
		t.Error("Deserialize Error!")
	}
	PrintContrract(*v2_contract)
	fmt.Println("+++++++++++Serialize method++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	v_str, err := v2_contract.Serialize()
	if err != nil {
		t.Error("Serialize Error!")
	}
	fmt.Println(v_str)
}
