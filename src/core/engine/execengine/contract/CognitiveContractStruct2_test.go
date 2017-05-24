package contract

import (
	"errors"
	"unicontract/src/core/engine/execengine/property"
	"unicontract/src/core/engine/execengine/table"
	"unicontract/src/core/engine/common"
	"encoding/json"
	"fmt"
	"testing"
	"reflect"
	"strings"
	"unicontract/src/core/engine/execengine/constdef"
)
type GeneralComponentCase struct{
	Cname string  `json:"Cname"`
	Ctype string  `json:"Ctype"`
	Caption string `json:"Caption"`
	Description string `json:"Description"`
	PropertyTable  map[string] interface{}  `json:"-"`
}

func (gc *GeneralComponentCase) InitGeneralComponentCase(p_ctype string)error{
	var err error = nil
	if gc.Cname == "" {
		err = errors.New("GeneralComponentCase Need Cname!")
		return err
	}
	if gc.PropertyTable == nil {
		gc.PropertyTable = make(map[string] interface{}, 0)
	}
	gc.AddProperty(gc,"_Ctype", common.TernaryOperator(gc.Ctype == "", constdef.ComponentType[constdef.Component_Unknown], gc.Ctype))
	gc.AddProperty(gc, "_Cname", gc.Cname)

	gc.AddProperty(gc,"_Caption", gc.Caption)
	gc.AddProperty(gc,"_Description", gc.Description)
	return err
}

func (gc *GeneralComponentCase) AddProperty(v_object interface{}, str_name string, value interface{})property.PropertyT {
	var pro_object property.PropertyT
	if value == nil {
		pro_object = *property.NewPropertyT(str_name)
		gc.PropertyTable[str_name] = pro_object
		return pro_object
	}
	switch value.(type) {
	case string:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(string))
		gc.PropertyTable[str_name] = pro_object
	case uint:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(uint))
		gc.PropertyTable[str_name] = pro_object
	case []string:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.([]string))
		gc.PropertyTable[str_name] = pro_object
	case map[string]ContractAsset:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(map[string]ContractAsset))
		gc.PropertyTable[str_name] = pro_object
	case map[string]ContractSignature:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(map[string]ContractSignature))
		gc.PropertyTable[str_name] = pro_object
	case map[string][]interface{}:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(map[string][]interface{}))
		gc.PropertyTable[str_name] = pro_object
	}
	mutable := reflect.ValueOf(v_object).Elem()
	v_value := reflect.ValueOf(value)
	fmt.Println(mutable.FieldByName(strings.Replace(str_name, "_", "", 1)))
	fmt.Println(strings.Replace(str_name, "_", "", 1), mutable.FieldByName(strings.Replace(str_name, "_", "", 1)).CanSet(), v_value)
	mutable.FieldByName(strings.Replace(str_name, "_", "", 1)).Set(v_value)
	return pro_object
}

//===========================================================================================
type ContractCase struct {
	GeneralComponentCase
	//TODO: need sort struct
	//type: map[string][]property.PropertyT   type:  Unknown, Data, Task, Expression
	ComponentTable table.ComponentTable  `json:"-"`

	//根据实际业务场景增加的属性
	ContractState string `json:"ContractState"`
	Creator string `json:"Creator"`
	CreateTime string `json:"CreateTime"`
	StartTime string `json:"StartTime"`
	EndTime string `json:"EndTime"`
	ContractOwners []string `json:"ContractOwners"`
	ContractAssets []ContractAsset `json:"ContractAssets"`
	ContractSignatures []ContractSignature `json:"ContractSignatures"`
	//TODO: need sort struct
	//type: Unknown, Data, Task, Expression
	ContractComponents map[string][]interface{} `json:"ContractComponents"`
}


//描述态序列化
// 序列化过程中，需要识别所有属性、component或table中的内容，将其序列化到字符串
func (tm *ContractCase)Serialize()string{
	if tm == nil {
		return ""
	}
	var s_model string = ""
	if s_model,err := json.Marshal(tm);err == nil {
		fmt.Println(string(s_model))
	}else {
		fmt.Println(err)
	}
	return s_model
}

//输入串后正常反序列化
func (tm *ContractCase)Deserialize(p_str string) *ContractCase{
	if p_str == "" || tm == nil{
		return nil
	}

	if err := json.Unmarshal([]byte(p_str), &tm); err != nil {
		fmt.Println(err)
	}
	return tm
}
//运行态
func (cc *ContractCase)InitContractCase()error{
	var err error = nil
	if cc.Cname == "" {
		//TODO log
		err = errors.New("Contract Need Cname!")
		return err
	}
	if cc.Caption == "" {
		//TODO log
		err = errors.New("Contract Need Caption!")
		return err
	}
	if cc.Description == "" {
		//TODO log
		err = errors.New("Contract Need Description!")
		return err
	}
	if cc.Creator == "" {
		//TODO log
		err = errors.New("Contract Need Creator!")
		return err
	}
	if cc.CreateTime == "" {
		//TODO log
		err = errors.New("Contract Need CreateTime!")
		return err
	}
	if cc.StartTime == "" {
		//TODO log
		err = errors.New("Contract Need StartTime!")
		return err
	}
	if cc.EndTime == "" {
		//TODO log
		err = errors.New("Contract Need EndTime!")
		return err
	}
	if cc.ContractOwners == nil || len(cc.ContractOwners) == 0 {
		//TODO log
		err = errors.New("Contract Need ContractOwners!")
		return err
	}
	if cc.ContractAssets == nil || len(cc.ContractAssets) == 0  {
		//TODO log
		err = errors.New("Contract Need ContractAssets!")
		return err
	}
	if cc.ContractSignatures == nil || len(cc.ContractSignatures) == 0  {
		//TODO log
		err = errors.New("Contract Need ContractOwners!")
		return err
	}
	cc.InitGeneralComponentCase(constdef.ComponentType[constdef.Component_Contract])
	cc.AddProperty(cc,"_ContractState", common.TernaryOperator(cc.ContractState == "", constdef.ContractState[constdef.Contract_Create], cc.ContractState))
	cc.AddProperty(cc,"_Creator", cc.Creator)
	cc.AddProperty(cc,"_CreateTime", cc.CreateTime)
	cc.AddProperty(cc,"_StartTime", cc.StartTime)
	cc.AddProperty(cc,"_EndTime", cc.EndTime)
	cc.AddProperty(cc,"_ContractOwners", cc.ContractOwners)
	cc.AddProperty(cc,"_ContractAssets", cc.ContractAssets)
	cc.AddProperty(cc,"_ContractSignatures", cc.ContractSignatures)

	var meta_map map[string]string = make(map[string]string, 0)
	meta_map["_UCVM_Version"] = constdef.UCVM_Version
	meta_map["_UCVM_CopyRight"] = constdef.UCVM_CopyRight
	meta_map["_UCVM_Date"] = constdef.UCVM_Date
	//cc.GeneralComponent.AddMetaAttribute(meta_map)
    //component table初始化 TODO
	cc.ComponentTable = *new(table.ComponentTable)
	cc.InitComponentTable()
	//cc.loadBuildInFunctions()
	//cc.loadExpressionParser()
	return err
}

func (cc *ContractCase)InitComponentTable(){
	if cc.ContractComponents == nil || len(cc.ContractComponents) == 0 {
		return
	}
	for v_key,v_value := range cc.ContractComponents {
		fmt.Println(v_key)
		fmt.Println(v_value)
	}
}


func TestM(t *testing.T)  {
	var p_str string = `{
"ContractId":"xxxxxxxxxxxxxxxxxxxxx",
"Cname":"contract_mobilecallback",
"Caption":"购智能手机返话费合约产品协议",
"Description":"移动用户A花费500元购买移动运营商B的提供的合约智能手机C后，要求用户每月消费58元以上通信费，移动运营商B便可按月返还话费（即：每月1号返还移动用户A20元话费），连续返还12个月",
"Creator":"ABCDEFGHIJKLMNIPQRSTUVWXYZ",
"CreateTime":"2016-12-2012:00:00",
"StartTime":"2017-01-0112:00:00",
"EndTime":"2017-01-0112:00:00",
"TestAA":"",
"ContractOwners":[
"AXXXXXXXXXXX",
"BXXXXXXXXXXX"
],
"ContractAssets":[
{
"Name":"asset_money",
"Caption":"理财产品",
"Description":"理财资产",
"Unit":"份",
"Amount":1000
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
"ContractComponents":{
"TaskType_Enquiry": [],
"TaskType_Action": [],
"TaskType_Decision":[]
}}`
	fmt.Println("描述态 =》 运行态")
	fmt.Println("反序列化：")
	var v_model ContractCase = *new(ContractCase)
	v_model.Deserialize(p_str)
	fmt.Println(v_model)
	fmt.Println("Cname: ", v_model.Cname)
	fmt.Println("Ctype: ", v_model.Ctype)
	fmt.Println("Caption: ", v_model.Caption)
	fmt.Println("Description: ", v_model.Description)
	fmt.Println("ContractState: ", v_model.ContractState)
	fmt.Println("Creator: ", v_model.Creator)
	fmt.Println("CreateTime: ", v_model.CreateTime)
	fmt.Println("StartTime: ", v_model.StartTime)
	fmt.Println("EndTime: ", v_model.EndTime)
	fmt.Println("ContractOwners: ", v_model.ContractOwners)
	fmt.Println("ContractAssets: ", v_model.ContractAssets)
	fmt.Println("ContractSignatures", v_model.ContractSignatures)
    fmt.Println("运行态初始化：")
	v_model.InitContractCase()
	v_model.InitComponentTable()
	fmt.Println("    property_table: ")
	fmt.Println(v_model.PropertyTable)
	fmt.Println("Cname: ", v_model.Cname)
	fmt.Println("Ctype: ", v_model.Ctype)
	fmt.Println("Caption: ", v_model.Caption)
	fmt.Println("Description: ", v_model.Description)
	fmt.Println("ContractState: ", v_model.ContractState)
	fmt.Println("Creator: ", v_model.Creator)
	fmt.Println("CreateTime: ", v_model.CreateTime)
	fmt.Println("StartTime: ", v_model.StartTime)
	fmt.Println("EndTime: ", v_model.EndTime)
	fmt.Println("ContractOwners: ", v_model.ContractOwners)
	fmt.Println("ContractAssets: ", v_model.ContractAssets)
	fmt.Println("ContractSignatures", v_model.ContractSignatures)
	fmt.Println("序列化：")
	fmt.Println(v_model.Serialize())
}