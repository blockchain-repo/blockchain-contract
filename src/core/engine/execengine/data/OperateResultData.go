package data

import (
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
	"unicontract/src/core/engine/execengine/constdef"
	"github.com/astaxie/beego/logs"
)

type OperateResultData struct {
	GeneralData
}

func NewOperateResultData()*OperateResultData{
	n := &OperateResultData{}
	return n
}
//====================接口方法========================
func (nd OperateResultData)GetName()string{
	if nd.PropertyTable[_Parent] != nil  {
		parent_property := nd.PropertyTable[_Parent].(property.PropertyT)
		if parent_property.GetValue() != nil {
			v_general_data := parent_property.GetValue().(inf.IData)
			if v_general_data.GetName() != "" {
				return v_general_data.GetName() + "." + nd.GetCname()
			} else {
				return nd.GetCname()
			}
		}
	}
	return nd.GetCname()
}

func (td OperateResultData) GetValue() interface{}{
	value_property := td.PropertyTable[_Value].(property.PropertyT)
	if value_property.GetValue() != nil {
		return value_property.GetValue()
	} else {
		v_contract := td.GeneralComponent.GetContract()
		v_default := v_contract.ProcessString(td.GetDefaultValue().(string))
		return v_default
	}
}
func (td OperateResultData)SetContract(p_contract inf.ICognitiveContract) {
	td.GeneralComponent.SetContract(p_contract)
}
func (td OperateResultData)GetContract() inf.ICognitiveContract {
	return td.GeneralComponent.GetContract()
}
func (gc OperateResultData)GetCtype()string{
	if gc.PropertyTable["_Ctype"] == nil {
		return ""
	}
	ctype_property := gc.PropertyTable["_Ctype"].(property.PropertyT)
	return ctype_property.GetValue().(string)
}
//====================描述态==========================


//====================运行态==========================
func (td *OperateResultData) InitOperateResultData()error{
	var err error = nil
	err = td.InitGeneralData()
	if err != nil {
		logs.Error("InitOperateResultData fail["+err.Error()+"]")
		return err
	}
	td.SetCtype(constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_OperateResultData])

	td.SetHardConvType("OperateResultData")
	return err
}
