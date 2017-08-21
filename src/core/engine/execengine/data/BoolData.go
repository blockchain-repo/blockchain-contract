package data

import (
	"encoding/json"
	"fmt"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
)

type BoolData struct {
	GeneralData
	ValueBool        bool `json:"ValueBool"`
	DefaultValueBool bool `json:"DefaultValueBool"`
}

const (
	_ValueBool        = "_ValueBool"
	_DefaultValueBool = "_DefaultValueBool"
)

func NewBoolData() *BoolData {
	n := &BoolData{}
	return n
}

//====================接口方法========================
func (td BoolData) GetName() string {
	return td.GeneralData.GetName()
}

func (td BoolData) GetValue() interface{} {
	return td.GetValueBool()
}
func (td BoolData) SetContract(p_contract inf.ICognitiveContract) {
	td.GeneralComponent.SetContract(p_contract)
}
func (td BoolData) GetContract() inf.ICognitiveContract {
	return td.GeneralComponent.GetContract()
}
func (td BoolData) GetCtype() string {
	return td.GeneralData.GetCtype()
}
func (td BoolData) SetValue(p_Value interface{}) {
	td.SetValueBool(p_Value)
}
func (td BoolData) CleanValueInProcess() {
	td.GeneralData.CleanValueInProcess()
	td.SetValueBool("")
	td.SetDefaultValueBool("")
}

//====================描述态==========================
//序列化： 需要将运行态结构 序列化到 描述态中
func (td *BoolData) RunningToStatic() {
	td.GeneralData.RunningToStatic()
	valueBool_property, ok := td.PropertyTable[_ValueBool].(property.PropertyT)
	if ok {
		td.ValueBool, _ = valueBool_property.GetValue().(bool)
	}
	defaultValueBool_property, ok := td.PropertyTable[_DefaultValueBool].(property.PropertyT)
	if ok {
		td.DefaultValueBool, _ = defaultValueBool_property.GetValue().(bool)
	}
}

func (td *BoolData) Serialize() (string, error) {
	td.RunningToStatic()
	if s_model, err := json.Marshal(td); err == nil {
		return string(s_model), err
	} else {
		uniledgerlog.Error("Bool Data fail[" + err.Error() + "]")
		return "", err
	}
}

//====================运行态==========================
func (td *BoolData) InitBoolData() error {
	var err error = nil
	err = td.InitGeneralData()
	if err != nil {
		uniledgerlog.Error("InitBoolData fail[" + err.Error() + "]")
		return err
	}
	td.SetCtype(constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Bool])
	common.AddProperty(td, td.PropertyTable, _ValueBool, td.ValueBool)
	common.AddProperty(td, td.PropertyTable, _DefaultValueBool, td.DefaultValueBool)
	td.SetHardConvType("bool")
	return err
}

//=====Getter方法
func (td *BoolData) GetValueBool() interface{} {
	value_property, ok := td.PropertyTable[_ValueBool].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	if value_property.GetValue() != nil {
		return value_property.GetValue()
	} else {
		v_default := td.GetDefaultValueBool()
		return v_default
	}
}
func (td *BoolData) GetDefaultValueBool() interface{} {
	value_property, ok := td.PropertyTable[_DefaultValueBool].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	if value_property.GetValue() != nil {
		return value_property.GetValue()
	}
	return nil
}

//=====Setter方法
func (td *BoolData) SetValueBool(p_ValueBool interface{}) {
	if p_ValueBool != nil {
		p_value, ok := p_ValueBool.(bool)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return
		}
		td.ValueBool = p_value
		value_property, ok := td.PropertyTable[_ValueBool].(property.PropertyT)
		if !ok {
			value_property = *property.NewPropertyT(_ValueBool)
		}
		value_property.SetValue(p_ValueBool)
		td.PropertyTable[_ValueBool] = value_property
	}
}

func (td *BoolData) SetDefaultValueBool(p_DefaultValueBool interface{}) {
	if p_DefaultValueBool != nil {
		p_defaultvalue, ok := p_DefaultValueBool.(bool)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return
		}
		td.DefaultValueBool = p_defaultvalue
		defaultvalue_property, ok := td.PropertyTable[_DefaultValueBool].(property.PropertyT)
		if !ok {
			defaultvalue_property = *property.NewPropertyT(_DefaultValueBool)
		}
		defaultvalue_property.SetValue(p_DefaultValueBool)
		td.PropertyTable[_DefaultValueBool] = defaultvalue_property
	}
}
