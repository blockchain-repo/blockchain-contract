package data

import (
	"encoding/json"
	"errors"
	"strings"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
)

type TextData struct {
	GeneralData
	ValueString        string `json:"ValueString"`
	DefaultValueString string `json:"DefaultValueString"`
}

const (
	_ValueString        = "_ValueString"
	_DefaultValueString = "_DefaultValueString"
)

func NewTextData() *TextData {
	n := &TextData{}
	return n
}

//====================接口方法========================
func (td TextData) GetName() string {
	return td.GeneralData.GetName()
}

func (td TextData) GetValue() interface{} {
	return td.GetValueString()
}
func (td TextData) SetContract(p_contract inf.ICognitiveContract) {
	td.GeneralComponent.SetContract(p_contract)
}
func (td TextData) GetContract() inf.ICognitiveContract {
	return td.GeneralComponent.GetContract()
}
func (td TextData) GetCtype() string {
	return td.GeneralData.GetCtype()
}
func (td TextData) SetValue(p_Value interface{}) {
	td.SetValueString(p_Value)
}
func (td TextData) CleanValueInProcess() {
	td.GeneralData.CleanValueInProcess()
	td.SetValueString("")
	td.SetDefaultValueString("")
}

//====================描述态==========================
//序列化： 需要将运行态结构 序列化到 描述态中
func (td *TextData) RunningToStatic() {
	td.GeneralData.RunningToStatic()
	valueString_property, ok := td.PropertyTable[_ValueString].(property.PropertyT)
	if ok {
		td.ValueString, _ = valueString_property.GetValue().(string)
	}
	defaultValueString_property, ok := td.PropertyTable[_DefaultValueString].(property.PropertyT)
	if ok {
		td.DefaultValueString, _ = defaultValueString_property.GetValue().(string)
	}
}

func (td *TextData) Serialize() (string, error) {
	td.RunningToStatic()
	if s_model, err := json.Marshal(td); err == nil {
		return string(s_model), err
	} else {
		uniledgerlog.Error("Contract Text Data fail[" + err.Error() + "]")
		return "", err
	}
}

//====================运行态==========================
func (td *TextData) InitTextData() error {
	var err error = nil
	err = td.InitGeneralData()
	if err != nil {
		uniledgerlog.Error("InitTextData fail[" + err.Error() + "]")
		return err
	}
	td.SetCtype(constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Text])
	common.AddProperty(td, td.PropertyTable, _ValueString, td.ValueString)
	common.AddProperty(td, td.PropertyTable, _DefaultValueString, td.DefaultValueString)
	td.SetHardConvType("string")
	return err
}

//=====Getter方法
func (td *TextData) GetValueString() interface{} {
	value_property := td.PropertyTable[_ValueString].(property.PropertyT)
	if value_property.GetValue() != nil {
		return value_property.GetValue()
	} else {
		v_default := td.GetDefaultValueString()
		return v_default
	}
}
func (td *TextData) GetDefaultValueString() interface{} {
	value_property := td.PropertyTable[_DefaultValueString].(property.PropertyT)
	if value_property.GetValue() != nil {
		return value_property.GetValue()
	}
	return nil
}

//=====Setter方法
func (td *TextData) SetValueString(p_ValueString interface{}) {
	if p_ValueString != nil {
		td.ValueString = p_ValueString.(string)
		value_property := td.PropertyTable[_ValueString].(property.PropertyT)
		value_property.SetValue(p_ValueString)
		td.PropertyTable[_ValueString] = value_property
	}
}

func (td *TextData) SetDefaultValueString(p_DefaultValueString interface{}) {
	if p_DefaultValueString != nil {
		td.DefaultValueString = p_DefaultValueString.(string)
		defaultvalue_property := td.PropertyTable[_DefaultValueString].(property.PropertyT)
		defaultvalue_property.SetValue(p_DefaultValueString)
		td.PropertyTable[_DefaultValueString] = defaultvalue_property
	}
}

//=====运算
func (td *TextData) Eq(p_data interface{}) (bool, error) {
	var f_leftdata string = td.GetValue().(string)
	var f_rightdata string
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(string)
	case string:
		f_rightdata = p_data.(string)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return strings.EqualFold(f_leftdata, f_rightdata), f_error
}

func (td *TextData) Add(p_data interface{}) (string, error) {
	var f_leftdata string = td.GetValue().(string)
	var f_rightdata string
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(string)
	case string:
		f_rightdata = p_data.(string)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata + f_rightdata, f_error
}

func (td *TextData) RAdd(p_data interface{}) (string, error) {
	var f_leftdata string = td.GetValue().(string)
	var f_rightdata string
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(string)
	case string:
		f_rightdata = p_data.(string)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_rightdata + f_leftdata, f_error
}

func (td *TextData) Len() int {
	var f_leftdata string = td.GetValue().(string)
	return len(f_leftdata)
}
