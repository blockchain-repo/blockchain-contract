package data

import (
	"unicontract/src/core/engine/execengine/property"
	"strconv"
	"errors"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/constdef"
	"fmt"
)

//支持int中的各种类型uint8, uint16, uint32, uint64; 不可直接用unit
type UintData struct {
	GeneralData
	DataRange [2]uint `json:"DataRange"`
}

func NewUintData()*UintData{
	n := &UintData{}
	return n
}
//====================接口方法========================
func (nd UintData)GetName()string{
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

func (ud UintData) GetValue() interface{}{
	value_property := ud.PropertyTable[_Value].(property.PropertyT)
	if value_property.GetValue() != nil {
		return value_property.GetValue()
	} else {
		v_contract := ud.GeneralComponent.GetContract()
		v_default := v_contract.ProcessString(ud.GetDefaultValue().(string))
		return v_default
	}
}

func (ud UintData)GetContract() inf.ICognitiveContract {
	return ud.GeneralComponent.GetContract()
}
func (ud UintData)SetContract(p_contract inf.ICognitiveContract) {
	ud.GeneralComponent.SetContract(p_contract)
}
func (gc UintData)GetCtype()string{
	if gc.PropertyTable["_Ctype"] == nil {
		return ""
	}
	ctype_property := gc.PropertyTable["_Ctype"].(property.PropertyT)
	return ctype_property.GetValue().(string)
}
//====================描述态==========================

//====================运行态==========================
func (ud *UintData) InitUintData()error{
	var err error = nil
	err = ud.InitGeneralData()
	if err != nil {
        //TODO log
		return err
	}
	ud.SetCtype(constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Numeric_Uint])
	fmt.Println(ud.GetCtype())
	var data_range [2]uint = [2]uint{0, 2147483647}
	if ud.DataRange[0] == 0 && ud.DataRange[1] == 0  {
		ud.AddProperty(ud, _DataRange, data_range)
	} else {
		ud.AddProperty(ud, _DataRange, ud.DataRange)
	}
	ud.SetHardConvType("uint")
	return err
}
//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func (ud *UintData) GetDataRange()[2]uint{
	datarange_property := ud.PropertyTable[_DataRange].(property.PropertyT)
	return datarange_property.GetValue().([2]uint)
}

func (ud *UintData) SetDataRange(data_range [2]uint)error{
	var err error = nil
	if data_range[0] == 0 && data_range[1] == 0 {
		var data_range = [2]uint{0, 2147483647}
		ud.DataRange = data_range
		datarange_property := ud.PropertyTable[_DataRange].(property.PropertyT)
		datarange_property.SetValue(data_range)
		ud.PropertyTable[_DataRange] = datarange_property
	} else {
		var f_range [2]uint = data_range
		if f_range[0] < 0 || f_range[1] < 0 {
			err = errors.New("range must > 0")
		} else if f_range[0] <= f_range[1] {
			ud.DataRange = f_range
			datarange_property := ud.PropertyTable[_DataRange].(property.PropertyT)
			datarange_property.SetValue(data_range)
			ud.PropertyTable[_DataRange] = datarange_property
		}else{
			var str_error string = "Data range Error(low:" + strconv.FormatUint(uint64(f_range[0]), 10) +
				", high:" + strconv.FormatUint(uint64(f_range[1]), 10) + ")!"
			err = errors.New(str_error)
		}
	}
	ud.DataRange = data_range
	return err
}

func (ud *UintData) CheckRange(check_data uint)bool{
	var r_ret = false
	if len(ud.GetDataRange()) == 0 {
		r_ret = true
	} else {
		var f_range = ud.GetDataRange()
		if check_data >= f_range[0] && check_data <= f_range[1] {
			r_ret = true
		} else {
			r_ret = false
		}
	}
	return r_ret
}

func (ud *UintData) Add(p_data interface{})(uint, error){
	var f_leftdata uint = ud.GetValue().(uint)
	var f_rightdata uint
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(uint)
	case uint:
		f_rightdata = p_data.(uint)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata + f_rightdata, f_error
}

func (ud *UintData) RAdd(p_data interface{})(uint, error){
	var f_leftdata uint = ud.GetValue().(uint)
	var f_rightdata uint
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(uint)
	case uint:
		f_rightdata = p_data.(uint)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata + f_rightdata, f_error
}

func (ud *UintData) Sub(p_data interface{})(uint, error) {
	var f_leftdata uint = ud.GetValue().(uint)
	var f_rightdata uint
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(uint)
	case uint:
		f_rightdata = p_data.(uint)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata - f_rightdata, f_error
}

func (ud *UintData) RSub(p_data interface{})(uint, error) {
	var f_leftdata uint = ud.GetValue().(uint)
	var f_rightdata uint
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(uint)
	case uint:
		f_rightdata = p_data.(uint)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_rightdata - f_leftdata, f_error
}

func (ud *UintData) Mul(p_data interface{})(uint, error) {
	var f_leftdata uint = ud.GetValue().(uint)
	var f_rightdata uint
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(uint)
	case uint:
		f_rightdata = p_data.(uint)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata * f_rightdata, f_error
}

func (ud *UintData) RMul(p_data interface{})(uint, error) {
	var f_leftdata uint = ud.GetValue().(uint)
	var f_rightdata uint
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(uint)
	case uint:
		f_rightdata = p_data.(uint)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata * f_rightdata, f_error
}

func (ud *UintData) Div(p_data interface{})(uint, error) {
	var f_leftdata uint = ud.GetValue().(uint)
	var f_rightdata uint
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(uint)
	case uint:
		f_rightdata = p_data.(uint)
	default:
		f_error = errors.New("Param Type Error!")
	}
	if f_rightdata == 0 {
		return 0, errors.New("Div right num is zero!")
	}
	return f_leftdata / f_rightdata, f_error
}

func (ud *UintData) RDiv(p_data interface{})(uint, error) {
	var f_leftdata uint = ud.GetValue().(uint)
	var f_rightdata uint
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(uint)
	case uint:
		f_rightdata = p_data.(uint)
	default:
		f_error = errors.New("Param Type Error!")
	}
	if f_leftdata == 0 {
		return 0, errors.New("Div right num is zero!")
	}
	return f_rightdata / f_leftdata, f_error
}

func (ud *UintData) Mod(p_data interface{})(uint, error) {
	var f_leftdata uint = ud.GetValue().(uint)
	var f_rightdata uint
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(uint)
	case uint:
		f_rightdata = p_data.(uint)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata % f_rightdata, f_error
}

func (ud *UintData) RMod(p_data interface{})(uint, error) {
	var f_leftdata uint = ud.GetValue().(uint)
	var f_rightdata uint
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(uint)
	case uint:
		f_rightdata = p_data.(uint)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_rightdata % f_leftdata, f_error
}

func (ud *UintData) Lt(p_data interface{})(bool, error) {
	var f_leftdata uint = ud.GetValue().(uint)
	var f_rightdata uint
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(uint)
	case uint:
		f_rightdata = p_data.(uint)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata < f_rightdata, f_error
}

func (ud *UintData) Le(p_data interface{})(bool, error) {
	var f_leftdata uint = ud.GetValue().(uint)
	var f_rightdata uint
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(uint)
	case uint:
		f_rightdata = p_data.(uint)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata <= f_rightdata, f_error
}

func (ud *UintData) Eq(p_data interface{})(bool, error) {
	var f_leftdata uint = ud.GetValue().(uint)
	var f_rightdata uint
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(uint)
	case uint:
		f_rightdata = p_data.(uint)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata == f_rightdata, f_error
}

func (ud *UintData) Ne(p_data interface{})(bool, error) {
	var f_leftdata uint = ud.GetValue().(uint)
	var f_rightdata uint
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(uint)
	case uint:
		f_rightdata = p_data.(uint)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata != f_rightdata, f_error
}

func (ud *UintData) Ge(p_data interface{})(bool, error) {
	var f_leftdata uint = ud.GetValue().(uint)
	var f_rightdata uint
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(uint)
	case uint:
		f_rightdata = p_data.(uint)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata >= f_rightdata, f_error
}

func (ud *UintData) Gt(p_data interface{})(bool, error) {
	var f_leftdata uint = ud.GetValue().(uint)
	var f_rightdata uint
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(uint)
	case uint:
		f_rightdata = p_data.(uint)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata > f_rightdata, f_error
}