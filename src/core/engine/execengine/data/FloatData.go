package data

import (
	"unicontract/src/core/engine/execengine/property"
	"errors"
	"math"
	"unicontract/src/core/engine/execengine/inf"
	"strconv"
	"unicontract/src/core/engine/execengine/constdef"
)

//支持int中的各种类型float32, float64
//TODO :先以float64为主
type FloatData struct {
	GeneralData
	DataRange [2]float64  `json:"DataRange"`
}

func NewFloatData()*FloatData{
	n := &FloatData{}
	return n
}
//====================接口方法========================
func (nd FloatData)GetName()string{
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

func (fd FloatData) GetValue() interface{}{
	value_property := fd.PropertyTable[_Value].(property.PropertyT)
	if value_property.GetValue() != nil {
		return value_property.GetValue()
	} else {
		v_contract := fd.GeneralComponent.GetContract()
		v_default := v_contract.ProcessString(fd.GetDefaultValue().(string))
		return v_default
	}
}
func (fd FloatData)SetContract(p_contract inf.ICognitiveContract) {
	fd.GeneralComponent.SetContract(p_contract)
}
func (fd FloatData)GetContract() inf.ICognitiveContract {
	return fd.GeneralComponent.GetContract()
}
func (gc FloatData)GetCtype()string{
	if gc.PropertyTable["_Ctype"] == nil {
		return ""
	}
	ctype_property := gc.PropertyTable["_Ctype"].(property.PropertyT)
	return ctype_property.GetValue().(string)
}
//====================描述态==========================


//====================运行态==========================
func (fd *FloatData) InitFloatData()error{
	var err error = nil
	err = fd.InitGeneralData ()
	if err != nil {
		//TODO log
		return err
	}
	fd.SetCtype(constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Numeric_Float])
	var data_range [2]float64 = [2]float64{-math.MaxFloat64, math.MaxFloat64}
	if fd.DataRange[0] == 0 && fd.DataRange[1] == 0{
		fd.AddProperty(fd, _DataRange, data_range)
	} else {
		fd.AddProperty(fd, _DataRange, fd.DataRange)
	}

	var hard_conv_type string = "float64"
	fd.SetHardConvType(hard_conv_type)
	return err
}
//====属性Get方法
func (fd *FloatData) GetDataRange()[2]float64 {
	datarange_property := fd.PropertyTable[_DataRange].(property.PropertyT)
	return datarange_property.GetValue().([2]float64 )
}
//====属性Set方法
func (fd *FloatData) SetDataRange(data_range [2]float64)error{
	var err error = nil
	if data_range[0] == 0 && data_range[1] == 0 {
		var data_range = [2]float64{-math.MaxFloat64, math.MaxFloat64}
		fd.DataRange = data_range
		datarange_property := fd.PropertyTable[_DataRange].(property.PropertyT)
		datarange_property.SetValue(data_range)
		fd.PropertyTable[_DataRange] = datarange_property
	} else {
		var f_range [2]float64 = data_range
		if f_range[0] <= f_range[1] {
			fd.DataRange = f_range
			datarange_property := fd.PropertyTable[_DataRange].(property.PropertyT)
			datarange_property.SetValue(data_range)
			fd.PropertyTable[_DataRange] = datarange_property
		}else{
			var str_error string = "Data range Error(low:" + strconv.FormatFloat(f_range[0], 'f', -1, 64) +
				", high:" + strconv.FormatFloat(f_range[1], 'f', -1, 64) + ")!"
			err = errors.New(str_error)
		}
	}
	fd.DataRange = data_range
	return err
}
//=====运算
func (fd *FloatData) CheckRange(check_data float64)bool{
	var r_ret = false
	if len(fd.GetDataRange()) == 0 {
		r_ret = true
	} else {
		var f_range = fd.GetDataRange()
		if check_data >= f_range[0] && check_data <= f_range[1] {
			r_ret = true
		} else {
			r_ret = false
		}
	}
	return r_ret
}

func (nd *FloatData) Add(p_data interface{})(float64, error){
	var f_leftdata float64 = nd.GetValue().(float64)
	var f_rightdata float64
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(float64)
	case float64:
		f_rightdata = p_data.(float64)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata + f_rightdata, f_error
}

func (nd *FloatData) RAdd(p_data interface{})(float64, error){
	var f_leftdata float64 = nd.GetValue().(float64)
	var f_rightdata float64
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(float64)
	case float64:
		f_rightdata = p_data.(float64)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata + f_rightdata, f_error
}

func (nd *FloatData) Sub(p_data interface{})(float64, error) {
	var f_leftdata float64 = nd.GetValue().(float64)
	var f_rightdata float64
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(float64)
	case float64:
		f_rightdata = p_data.(float64)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata - f_rightdata, f_error
}

func (nd *FloatData) RSub(p_data interface{})(float64, error) {
	var f_leftdata float64 = nd.GetValue().(float64)
	var f_rightdata float64
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(float64)
	case float64:
		f_rightdata = p_data.(float64)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_rightdata - f_leftdata, f_error
}

func (nd *FloatData) Mul(p_data interface{})(float64, error) {
	var f_leftdata float64 = nd.GetValue().(float64)
	var f_rightdata float64
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(float64)
	case float64:
		f_rightdata = p_data.(float64)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata * f_rightdata, f_error
}

func (nd *FloatData) RMul(p_data interface{})(float64, error) {
	var f_leftdata float64 = nd.GetValue().(float64)
	var f_rightdata float64
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(float64)
	case float64:
		f_rightdata = p_data.(float64)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata * f_rightdata, f_error
}

func (nd *FloatData) Div(p_data interface{})(float64, error) {
	var f_leftdata float64 = nd.GetValue().(float64)
	var f_rightdata float64
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(float64)
	case float64:
		f_rightdata = p_data.(float64)
	default:
		f_error = errors.New("Param Type Error!")
	}
	if f_rightdata == 0 {
		return -1, errors.New("Div right num is zero!")
	}
	return f_leftdata / f_rightdata, f_error
}

func (nd *FloatData) RDiv(p_data interface{})(float64, error) {
	var f_leftdata float64 = nd.GetValue().(float64)
	var f_rightdata float64
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(float64)
	case float64:
		f_rightdata = p_data.(float64)
	default:
		f_error = errors.New("Param Type Error!")
	}
	if f_leftdata == 0 {
		return -1, errors.New("Div right num is zero!")
	}
	return f_rightdata / f_leftdata, f_error
}

func (nd *FloatData) Neg()(float64) {
	var f_leftdata float64 = nd.GetValue().(float64)
	return -f_leftdata
}

func (nd *FloatData) Lt(p_data interface{})(bool, error) {
	var f_leftdata float64 = nd.GetValue().(float64)
	var f_rightdata float64
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(float64)
	case float64:
		f_rightdata = p_data.(float64)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata < f_rightdata, f_error
}

func (nd *FloatData) Le(p_data interface{})(bool, error) {
	var f_leftdata float64 = nd.GetValue().(float64)
	var f_rightdata float64
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(float64)
	case float64:
		f_rightdata = p_data.(float64)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata <= f_rightdata, f_error
}

func (nd *FloatData) Eq(p_data interface{})(bool, error) {
	var f_leftdata float64 = nd.GetValue().(float64)
	var f_rightdata float64
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(float64)
	case float64:
		f_rightdata = p_data.(float64)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata == f_rightdata, f_error
}

func (nd *FloatData) Ne(p_data interface{})(bool, error) {
	var f_leftdata float64 = nd.GetValue().(float64)
	var f_rightdata float64
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(float64)
	case float64:
		f_rightdata = p_data.(float64)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata != f_rightdata, f_error
}

func (nd *FloatData) Ge(p_data interface{})(bool, error) {
	var f_leftdata float64 = nd.GetValue().(float64)
	var f_rightdata float64
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(float64)
	case float64:
		f_rightdata = p_data.(float64)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata >= f_rightdata, f_error
}

func (nd *FloatData) Gt(p_data interface{})(bool, error) {
	var f_leftdata float64 = nd.GetValue().(float64)
	var f_rightdata float64
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		f_rightdata = v_data.GetValue().(float64)
	case float64:
		f_rightdata = p_data.(float64)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata > f_rightdata, f_error
}