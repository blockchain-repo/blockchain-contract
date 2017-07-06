package data

import (
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"

	"encoding/json"
	"unicontract/src/common/uniledgerlog"
)

type OperateResultData struct {
	GeneralData
}

func NewOperateResultData() *OperateResultData {
	n := &OperateResultData{}
	return n
}

//====================接口方法========================
func (nd OperateResultData) GetName() string {
	if nd.PropertyTable[_Parent] != nil {
		parent_property, ok := nd.PropertyTable[_Parent].(property.PropertyT)
		if !ok {
			uniledgerlog.Error("assert error")
			return ""
		}
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

func (td OperateResultData) GetValue() interface{} {
	value_property, ok := td.PropertyTable[_Value].(property.PropertyT)
	if !ok {
		uniledgerlog.Error("assert error")
		return nil
	}
	if value_property.GetValue() != nil {
		return value_property.GetValue()
	} else {
		v_contract := td.GeneralComponent.GetContract()
		str, ok := td.GetDefaultValue().(string)
		if !ok {
			uniledgerlog.Error("assert error")
			return ""
		}
		v_default := v_contract.ProcessString(str)

		return v_default
	}
}
func (td OperateResultData) SetContract(p_contract inf.ICognitiveContract) {
	td.GeneralComponent.SetContract(p_contract)
}
func (td OperateResultData) GetContract() inf.ICognitiveContract {
	return td.GeneralComponent.GetContract()
}
func (gc OperateResultData) GetCtype() string {
	if gc.PropertyTable["_Ctype"] == nil {
		return ""
	}
	ctype_property, ok := gc.PropertyTable["_Ctype"].(property.PropertyT)
	if !ok {
		uniledgerlog.Error("assert error")
		return ""
	}
	str, ok := ctype_property.GetValue().(string)
	if !ok {
		uniledgerlog.Error("assert error")
		return ""
	}
	return str
}
func (gc OperateResultData) CleanValueInProcess() {
	gc.GeneralData.CleanValueInProcess()
}

//====================描述态==========================
//序列化： 需要将运行态结构 序列化到 描述态中
func (td *OperateResultData) RunningToStatic() {
	td.GeneralData.RunningToStatic()
}

func (td *OperateResultData) Serialize() (string, error) {
	td.RunningToStatic()
	if s_model, err := json.Marshal(td); err == nil {
		return string(s_model), err
	} else {
		uniledgerlog.Error("Contract OperateResultData Data fail[" + err.Error() + "]")
		return "", err
	}
}

//====================运行态==========================
func (td *OperateResultData) InitOperateResultData() error {
	var err error = nil
	err = td.InitGeneralData()
	if err != nil {
		uniledgerlog.Error("InitOperateResultData fail[" + err.Error() + "]")
		return err
	}
	td.SetCtype(constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_OperateResultData])

	td.SetHardConvType("OperateResultData")
	return err
}
