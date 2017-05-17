package data

import (
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
	"unicontract/src/core/engine/execengine/constdef"
	"github.com/astaxie/beego/logs"
)

//TODO
type CompoundData struct {
	GeneralData
}

func NewCompoundData()*CompoundData{
	cd := &CompoundData{}
	return cd
}
//====================接口方法========================
func (cd CompoundData)GetName()string{
	return cd.GeneralData.GetName()
}

func (cd CompoundData) GetValue() interface{}{
	value_property := cd.PropertyTable[_Value].(property.PropertyT)
	if value_property.GetValue() != nil {
		return value_property.GetValue().(int)
	} else {
		v_contract := cd.GeneralComponent.GetContract()
		v_default := v_contract.ProcessString(cd.GetDefaultValue().(string))
		return v_default
	}
}
func (cd CompoundData)SetContract(p_contract inf.ICognitiveContract) {
	cd.GeneralComponent.SetContract(p_contract)
}
func (cd CompoundData)GetContract() inf.ICognitiveContract {
	return cd.GeneralComponent.GetContract()
}

func (cd CompoundData)GetCtype()string{
	return cd.GeneralData.GetCtype()
}
func (cd CompoundData) SetValue(p_Value interface{}) {
	cd.GeneralData.SetValue(p_Value)
}
//====================描述态==========================


//====================运行态==========================
func (cd *CompoundData) InitCompoundData()error{
	var err error = nil
	err = cd.InitGeneralData ()
	if err == nil {
		logs.Error("InitCompoundData fail["+err.Error()+"]")
		return err
	}
	cd.SetCtype(constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Compound])
	return err
}