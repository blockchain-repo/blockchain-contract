package data

import (
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
	"unicontract/src/core/engine/execengine/constdef"
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
	parent_property := cd.PropertyTable[_Parent].(property.PropertyT)
	if parent_property.GetValue() != nil {
		var v_general_data GeneralData = parent_property.GetValue().(GeneralData)
		if v_general_data.GetCname() != "" {
			return v_general_data.GetCname() + "." + cd.GetCname()
		} else {
			return cd.GetCname()
		}
	} else {
		return cd.GetCname()
	}
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

func (gc CompoundData)GetCtype()string{
	if gc.PropertyTable["_Ctype"] == nil {
		return ""
	}
	ctype_property := gc.PropertyTable["_Ctype"].(property.PropertyT)
	return ctype_property.GetValue().(string)
}
//====================描述态==========================


//====================运行态==========================
func (cd *CompoundData) InitCompoundData()error{
	var err error = nil
	err = cd.InitGeneralData ()
	if err == nil {
		//TODO nil
		return err
	}
	cd.SetCtype(constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Compound])
	return err
}