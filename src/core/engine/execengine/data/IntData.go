package data

import (
	"unicontract/src/core/engine/execengine/property"
	"strconv"
	"errors"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/constdef"
	"github.com/astaxie/beego/logs"
	"unicontract/src/core/engine/common"
)

//支持int中的各种类型int8, int16, int32, int64; 不可直接用int
//TODO :先以int为主
type IntData struct {
	GeneralData
	DataRange [2]int `json:"DataRange"`
}

const (
	_Contract = "_Contract"
)

func NewIntData()*IntData{
	n := &IntData{}
	return n
}
//====================接口方法========================
func (nd IntData)GetName()string{
	return nd.GeneralData.GetName()
}

func (nd IntData) GetValue() interface{}{
	return nd.GeneralData.GetValue()
}

func (nd IntData)GetContract() inf.ICognitiveContract {
	return nd.GeneralData.GetContract()
}
func (nd IntData)SetContract(p_contract inf.ICognitiveContract) {
	nd.GeneralData.SetContract(p_contract)
}
func (nd IntData)GetCtype()string{
	return nd.GeneralData.GetCtype()
}
func (nd IntData) SetValue(p_Value interface{}){
	nd.GeneralData.SetValue(p_Value)
}
//====================描述态==========================


//====================运行态==========================
func (nd *IntData) InitIntData()error{
	var err error = nil
	err = nd.InitGeneralData()
    if err != nil {
		logs.Error("InitIntData fail["+err.Error()+"]")
		return err
	}
	nd.SetCtype(constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Numeric_Int])
	var data_range [2]int = [2]int{-2147483647, 2147483647}
	if nd.DataRange[0] ==  0 && nd.DataRange[1] == 0 {
		common.AddProperty(nd, nd.PropertyTable, _DataRange, data_range)
	} else {
		common.AddProperty(nd, nd.PropertyTable, _DataRange, nd.DataRange)
	}
	nd.SetHardConvType("int")
	return err
}

//====属性Get方法
func (nd *IntData) GetDataRange()[2]int{
	datarange_property := nd.PropertyTable[_DataRange].(property.PropertyT)
	return datarange_property.GetValue().([2]int)
}

//====属性Set方法
func (nd *IntData) SetDataRange(data_range [2]int)error{
	var err error = nil
	if data_range[0] == 0 && data_range[1] == 0 {
		var data_range [2]int = [2]int{-2147483647, 2147483647}
		nd.DataRange = data_range
		datarange_property := nd.PropertyTable[_DataRange].(property.PropertyT)
		datarange_property.SetValue(data_range)
		nd.PropertyTable[_DataRange] = datarange_property
	} else {
		var f_range [2]int = data_range
		if f_range[0] <= f_range[1] {
			nd.DataRange = f_range
			datarange_property := nd.PropertyTable[_DataRange].(property.PropertyT)
			datarange_property.SetValue(data_range)
			nd.PropertyTable[_DataRange] = datarange_property
		}else{
			var str_error string = "Data range Error(low:" + strconv.Itoa(f_range[0]) +
				", high:" + strconv.Itoa(f_range[1]) + ")!"
			err = errors.New(str_error)
		}
	}
	nd.DataRange = data_range
	return err
}
//=====运算
func (nd *IntData) CheckRange(check_data int)bool{
	var r_ret = false
	if len(nd.GetDataRange()) == 0 {
		r_ret = true
	} else {
		var f_range = nd.GetDataRange()
		if check_data >= f_range[0] && check_data <= f_range[1] {
			r_ret = true
		} else {
			r_ret = false
		}
	}
	return r_ret
}

func (nd *IntData) Add(p_data interface{})(int, error){
	var f_leftdata int = nd.GetValue().(int)
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(int)
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata + f_rightdata, f_error
}

func (nd *IntData) RAdd(p_data interface{})(int, error){
	var f_leftdata int = nd.GetValue().(int)
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(int)
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata + f_rightdata, f_error
}

func (nd *IntData) Sub(p_data interface{})(int, error) {
	var f_leftdata int = nd.GetValue().(int)
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(int)
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata - f_rightdata, f_error
}

func (nd *IntData) RSub(p_data interface{})(int, error) {
	var f_leftdata int = nd.GetValue().(int)
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(int)
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_rightdata - f_leftdata, f_error
}

func (nd *IntData) Mul(p_data interface{})(int, error) {
	var f_leftdata int = nd.GetValue().(int)
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(int)
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata * f_rightdata, f_error
}

func (nd *IntData) RMul(p_data interface{})(int, error) {
	var f_leftdata int = nd.GetValue().(int)
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(int)
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata * f_rightdata, f_error
}

func (nd *IntData) Div(p_data interface{})(int, error) {
	var f_leftdata int = nd.GetValue().(int)
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(int)
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	if f_rightdata == 0 {
		return -1, errors.New("Div right num is zero!")
	}
	return f_leftdata / f_rightdata, f_error
}

func (nd *IntData) RDiv(p_data interface{})(int, error) {
	var f_leftdata int = nd.GetValue().(int)
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(int)
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	if f_leftdata == 0 {
		return -1, errors.New("Div right num is zero!")
	}
	return f_rightdata / f_leftdata, f_error
}

func (nd *IntData) Mod(p_data interface{})(int, error) {
	var f_leftdata int = nd.GetValue().(int)
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(int)
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata % f_rightdata, f_error
}

func (nd *IntData) RMod(p_data interface{})(int, error) {
	var f_leftdata int = nd.GetValue().(int)
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(int)
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_rightdata % f_leftdata, f_error
}

func (nd *IntData) Neg()(int) {
	var f_leftdata int = nd.GetValue().(int)
	return -f_leftdata
}

func (nd *IntData) Lt(p_data interface{})(bool, error) {
	var f_leftdata int = nd.GetValue().(int)
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(int)
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata < f_rightdata, f_error
}

func (nd *IntData) Le(p_data interface{})(bool, error) {
	var f_leftdata int = nd.GetValue().(int)
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(int)
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata <= f_rightdata, f_error
}

func (nd *IntData) Eq(p_data interface{})(bool, error) {
	var f_leftdata int = nd.GetValue().(int)
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(int)
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata == f_rightdata, f_error
}

func (nd *IntData) Ne(p_data interface{})(bool, error) {
	var f_leftdata int = nd.GetValue().(int)
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(int)
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata != f_rightdata, f_error
}

func (nd *IntData) Ge(p_data interface{})(bool, error) {
	var f_leftdata int = nd.GetValue().(int)
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(int)
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata >= f_rightdata, f_error
}

func (nd *IntData) Gt(p_data interface{})(bool, error) {
	var f_leftdata int = nd.GetValue().(int)
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(int)
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata > f_rightdata, f_error
}