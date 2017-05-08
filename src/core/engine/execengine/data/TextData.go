package data

import (
	"errors"
	"strings"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
	"unicontract/src/core/engine/execengine/constdef"
)

type TextData struct {
	GeneralData
}

func NewTextData()*TextData{
	n := &TextData{}
	return n
}
//====================接口方法========================
func (nd TextData)GetName()string{
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

func (td TextData) GetValue() interface{}{
	value_property := td.PropertyTable[_Value].(property.PropertyT)
	if value_property.GetValue() != nil {
		return value_property.GetValue()
	} else {
		v_contract := td.GeneralComponent.GetContract()
		v_default := v_contract.ProcessString(td.GetDefaultValue().(string))
		return v_default
	}
}
func (td TextData)SetContract(p_contract inf.ICognitiveContract) {
	td.GeneralComponent.SetContract(p_contract)
}
func (td TextData)GetContract() inf.ICognitiveContract {
	return td.GeneralComponent.GetContract()
}
func (gc TextData)GetCtype()string{
	if gc.PropertyTable["_Ctype"] == nil {
		return ""
	}
	ctype_property := gc.PropertyTable["_Ctype"].(property.PropertyT)
	return ctype_property.GetValue().(string)
}
//====================描述态==========================


//====================运行态==========================
func (td *TextData) InitTextData()error{
	var err error = nil
	err = td.InitGeneralData()
    if err != nil {
		//TODO log
		return err
	}
	td.SetCtype(constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Text])

	td.SetHardConvType("string")
	return err
}
//=====运算
func (td *TextData) Eq(p_data interface{})(bool, error) {
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

func (td *TextData) Add(p_data interface{})(string, error) {
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

func (td *TextData) RAdd(p_data interface{})(string, error) {
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

func (td *TextData) Len()int{
	var f_leftdata string = td.GetValue().(string)
	return len(f_leftdata)
}