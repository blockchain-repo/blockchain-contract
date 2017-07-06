package data

import (
	"encoding/json"
	"errors"
	"math"
	"strconv"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
)

//支持int中的各种类型float32, float64
//TODO :先以float64为主
type FloatData struct {
	GeneralData
	ValueFloat        float64    `json:"ValueFloat"`
	DefaultValueFloat float64    `json:"DefaultValueFloat"`
	DataRangeFloat    [2]float64 `json:"DataRangeFloat"`
}

const (
	_DataRangeFloat    = "_DataRangeFloat"
	_ValueFloat        = "_ValueFloat"
	_DefaultValueFloat = "_DefaultValueFloat"
)

func NewFloatData() *FloatData {
	n := &FloatData{}
	return n
}

//====================接口方法========================
func (fd FloatData) GetName() string {
	return fd.GeneralData.GetName()
}

func (fd FloatData) GetValue() interface{} {
	return fd.GetValueFloat()
}
func (fd FloatData) SetContract(p_contract inf.ICognitiveContract) {
	fd.GeneralComponent.SetContract(p_contract)
}
func (fd FloatData) GetContract() inf.ICognitiveContract {
	return fd.GeneralComponent.GetContract()
}
func (fd FloatData) GetCtype() string {
	return fd.GeneralData.GetCtype()
}
func (fd FloatData) SetValue(p_Value interface{}) {
	fd.SetValueFloat(p_Value)
}
func (fd FloatData) CleanValueInProcess() {
	fd.GeneralData.CleanValueInProcess()
	fd.SetValueFloat(0.0)
	fd.SetDefaultValueFloat(0.0)
}

//====================描述态==========================
//序列化： 需要将运行态结构 序列化到 描述态中
func (fd *FloatData) RunningToStatic() {
	fd.GeneralData.RunningToStatic()
	valueFloat_property, ok := fd.PropertyTable[_ValueFloat].(property.PropertyT)
	if ok {
		fd.ValueFloat, _ = valueFloat_property.GetValue().(float64)
	}
	defaultValueFloat_property, ok := fd.PropertyTable[_DefaultValueFloat].(property.PropertyT)
	if ok {
		fd.DefaultValueFloat, _ = defaultValueFloat_property.GetValue().(float64)
	}
	dataRangeFloat_property, ok := fd.PropertyTable[_DataRangeFloat].(property.PropertyT)
	if ok {
		fd.DataRangeFloat, _ = dataRangeFloat_property.GetValue().([2]float64)
	}
}

func (fd *FloatData) Serialize() (string, error) {
	fd.RunningToStatic()
	if s_model, err := json.Marshal(fd); err == nil {
		return string(s_model), err
	} else {
		uniledgerlog.Error("Contract Float Data fail[" + err.Error() + "]")
		return "", err
	}
}

//====================运行态==========================
func (fd *FloatData) InitFloatData() error {
	var err error = nil
	err = fd.InitGeneralData()
	if err != nil {
		uniledgerlog.Error("InitFloatData fail[" + err.Error() + "]")
		return err
	}
	fd.SetCtype(constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Numeric_Float])
	var data_range [2]float64 = [2]float64{-math.MaxFloat64, math.MaxFloat64}
	if fd.DataRangeFloat[0] == 0 && fd.DataRangeFloat[1] == 0 {
		common.AddProperty(fd, fd.PropertyTable, _DataRangeFloat, data_range)
	} else {
		common.AddProperty(fd, fd.PropertyTable, _DataRangeFloat, fd.DataRangeFloat)
	}
	common.AddProperty(fd, fd.PropertyTable, _ValueFloat, fd.ValueFloat)
	common.AddProperty(fd, fd.PropertyTable, _DefaultValueFloat, fd.DefaultValueFloat)
	var hard_conv_type string = "float64"
	fd.SetHardConvType(hard_conv_type)
	return err
}

//====属性Get方法
func (fd *FloatData) GetDataRangeFloat() [2]float64 {
	datarange_property := fd.PropertyTable[_DataRangeFloat].(property.PropertyT)
	return datarange_property.GetValue().([2]float64)
}
func (fd *FloatData) GetValueFloat() interface{} {
	value_property := fd.PropertyTable[_ValueFloat].(property.PropertyT)
	if value_property.GetValue() != nil {
		return value_property.GetValue()
	} else {
		v_default := fd.GetDefaultValueFloat()
		return v_default
	}
}
func (fd *FloatData) GetDefaultValueFloat() interface{} {
	value_property := fd.PropertyTable[_DefaultValueFloat].(property.PropertyT)
	if value_property.GetValue() != nil {
		return value_property.GetValue()
	}
	return nil
}

//====属性Set方法
func (fd *FloatData) SetDataRangeFloat(data_range [2]float64) error {
	var err error = nil
	if data_range[0] == 0 && data_range[1] == 0 {
		var data_range = [2]float64{-math.MaxFloat64, math.MaxFloat64}
		fd.DataRangeFloat = data_range
		datarange_property := fd.PropertyTable[_DataRangeFloat].(property.PropertyT)
		datarange_property.SetValue(data_range)
		fd.PropertyTable[_DataRangeFloat] = datarange_property
	} else {
		var f_range [2]float64 = data_range
		if f_range[0] <= f_range[1] {
			fd.DataRangeFloat = f_range
			datarange_property := fd.PropertyTable[_DataRangeFloat].(property.PropertyT)
			datarange_property.SetValue(data_range)
			fd.PropertyTable[_DataRangeFloat] = datarange_property
		} else {
			var str_error string = "Data range Error(low:" + strconv.FormatFloat(f_range[0], 'f', -1, 64) +
				", high:" + strconv.FormatFloat(f_range[1], 'f', -1, 64) + ")!"
			err = errors.New(str_error)
		}
	}
	fd.DataRangeFloat = data_range
	return err
}
func (fd *FloatData) SetValueFloat(p_ValueFloat interface{}) {
	if p_ValueFloat != nil {
		fd.ValueFloat = p_ValueFloat.(float64)
		value_property := fd.PropertyTable[_ValueFloat].(property.PropertyT)
		value_property.SetValue(p_ValueFloat)
		fd.PropertyTable[_ValueFloat] = value_property
	}
}

func (fd *FloatData) SetDefaultValueFloat(p_DefaultValueFloat interface{}) {
	if p_DefaultValueFloat != nil {
		fd.DefaultValueFloat = p_DefaultValueFloat.(float64)
		defaultvalue_property := fd.PropertyTable[_DefaultValueFloat].(property.PropertyT)
		defaultvalue_property.SetValue(p_DefaultValueFloat)
		fd.PropertyTable[_DefaultValueFloat] = defaultvalue_property
	}
}

//=====运算
func (fd *FloatData) CheckRange(check_data float64) bool {
	var r_ret = false
	if len(fd.GetDataRangeFloat()) == 0 {
		r_ret = true
	} else {
		var f_range = fd.GetDataRangeFloat()
		if check_data >= f_range[0] && check_data <= f_range[1] {
			r_ret = true
		} else {
			r_ret = false
		}
	}
	return r_ret
}

func (fd *FloatData) Add(p_data interface{}) (float64, error) {
	var f_leftdata float64 = fd.GetValue().(float64)
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

func (fd *FloatData) RAdd(p_data interface{}) (float64, error) {
	var f_leftdata float64 = fd.GetValue().(float64)
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

func (fd *FloatData) Sub(p_data interface{}) (float64, error) {
	var f_leftdata float64 = fd.GetValue().(float64)
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

func (fd *FloatData) RSub(p_data interface{}) (float64, error) {
	var f_leftdata float64 = fd.GetValue().(float64)
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

func (fd *FloatData) Mul(p_data interface{}) (float64, error) {
	var f_leftdata float64 = fd.GetValue().(float64)
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

func (fd *FloatData) RMul(p_data interface{}) (float64, error) {
	var f_leftdata float64 = fd.GetValue().(float64)
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

func (fd *FloatData) Div(p_data interface{}) (float64, error) {
	var f_leftdata float64 = fd.GetValue().(float64)
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

func (fd *FloatData) RDiv(p_data interface{}) (float64, error) {
	var f_leftdata float64 = fd.GetValue().(float64)
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

func (fd *FloatData) Neg() float64 {
	var f_leftdata float64 = fd.GetValue().(float64)
	return -f_leftdata
}

func (fd *FloatData) Lt(p_data interface{}) (bool, error) {
	var f_leftdata float64 = fd.GetValue().(float64)
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

func (fd *FloatData) Le(p_data interface{}) (bool, error) {
	var f_leftdata float64 = fd.GetValue().(float64)
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

func (fd *FloatData) Eq(p_data interface{}) (bool, error) {
	var f_leftdata float64 = fd.GetValue().(float64)
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

func (fd *FloatData) Ne(p_data interface{}) (bool, error) {
	var f_leftdata float64 = fd.GetValue().(float64)
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

func (fd *FloatData) Ge(p_data interface{}) (bool, error) {
	var f_leftdata float64 = fd.GetValue().(float64)
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

func (fd *FloatData) Gt(p_data interface{}) (bool, error) {
	var f_leftdata float64 = fd.GetValue().(float64)
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
