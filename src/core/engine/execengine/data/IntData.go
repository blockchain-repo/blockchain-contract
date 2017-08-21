package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
)

//支持int中的各种类型int8, int16, int32, int64; 不可直接用int
//TODO :先以int为主
type IntData struct {
	GeneralData
	ValueInt        int    `json:"ValueInt"`
	DefaultValueInt int    `json:"DefaultValueInt"`
	DataRangeInt    [2]int `json:"DataRangeInt"`
}

const (
	_DataRangeInt    = "_DataRangeInt"
	_ValueInt        = "_ValueInt"
	_DefaultValueInt = "_DefaultValueInt"
)

func NewIntData() *IntData {
	n := &IntData{}
	return n
}

//====================接口方法========================
func (nd IntData) GetName() string {
	return nd.GeneralData.GetName()
}

func (nd IntData) GetValue() interface{} {
	return nd.GetValueInt()
}

func (nd IntData) GetContract() inf.ICognitiveContract {
	return nd.GeneralData.GetContract()
}
func (nd IntData) SetContract(p_contract inf.ICognitiveContract) {
	nd.GeneralData.SetContract(p_contract)
}
func (nd IntData) GetCtype() string {
	return nd.GeneralData.GetCtype()
}
func (nd IntData) SetValue(p_Value interface{}) {
	nd.SetValueInt(p_Value)
}

func (nd IntData) CleanValueInProcess() {
	nd.GeneralData.CleanValueInProcess()
	nd.SetDefaultValueInt(0)
	nd.SetValueInt(0)
}

//====================描述态==========================
//序列化： 需要将运行态结构 序列化到 描述态中
func (nd *IntData) RunningToStatic() {
	nd.GeneralData.RunningToStatic()
	valueInt_property, ok := nd.PropertyTable[_ValueInt].(property.PropertyT)
	if ok {
		nd.ValueInt, _ = valueInt_property.GetValue().(int)
	}
	defaultValueInt_property, ok := nd.PropertyTable[_DefaultValueInt].(property.PropertyT)
	if ok {
		nd.DefaultValueInt, _ = defaultValueInt_property.GetValue().(int)
	}
	dtaRangeInt_property, ok := nd.PropertyTable[_DataRangeInt].(property.PropertyT)
	if ok {
		nd.DataRangeInt, _ = dtaRangeInt_property.GetValue().([2]int)
	}
}

func (nd *IntData) Serialize() (string, error) {
	nd.RunningToStatic()
	if s_model, err := json.Marshal(nd); err == nil {
		return string(s_model), err
	} else {
		uniledgerlog.Error("Contract Int Data fail[" + err.Error() + "]")
		return "", err
	}
}

//====================运行态==========================
func (nd *IntData) InitIntData() error {
	var err error = nil
	err = nd.InitGeneralData()
	if err != nil {
		uniledgerlog.Error("InitIntData fail[" + err.Error() + "]")
		return err
	}
	nd.SetCtype(constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Numeric_Int])
	var data_range [2]int = [2]int{-2147483647, 2147483647}
	if nd.DataRangeInt[0] == 0 && nd.DataRangeInt[1] == 0 {
		common.AddProperty(nd, nd.PropertyTable, _DataRangeInt, data_range)
	} else {
		common.AddProperty(nd, nd.PropertyTable, _DataRangeInt, nd.DataRangeInt)
	}
	common.AddProperty(nd, nd.PropertyTable, _ValueInt, nd.ValueInt)
	common.AddProperty(nd, nd.PropertyTable, _DefaultValueInt, nd.DefaultValueInt)
	nd.SetHardConvType("int")
	return err
}

//====属性Get方法
func (nd *IntData) GetDataRangeInt() [2]int {
	dataRangeInt_property, ok := nd.PropertyTable[_DataRangeInt].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return [2]int{0, 0}
	}
	dataRangeInt_value, ok := dataRangeInt_property.GetValue().([2]int)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return [2]int{0, 0}
	}
	return dataRangeInt_value
}
func (nd *IntData) GetValueInt() interface{} {
	value_property, ok := nd.PropertyTable[_ValueInt].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	if value_property.GetValue() != nil {
		return value_property.GetValue()
	} else {
		v_default := nd.GetDefaultValueInt()
		return v_default
	}
}
func (nd *IntData) GetDefaultValueInt() interface{} {
	value_property, ok := nd.PropertyTable[_DefaultValueInt].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	if value_property.GetValue() != nil {
		return value_property.GetValue()
	}
	return nil
}

//====属性Set方法
func (nd *IntData) SetValueInt(p_ValueInt interface{}) {
	if p_ValueInt != nil {
		ok := false
		nd.ValueInt, ok = p_ValueInt.(int)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return
		}
		value_property, ok := nd.PropertyTable[_ValueInt].(property.PropertyT)
		if !ok {
			value_property = *property.NewPropertyT(_ValueInt)
		}
		value_property.SetValue(p_ValueInt)
		nd.PropertyTable[_ValueInt] = value_property
	}
}

func (nd *IntData) SetDefaultValueInt(p_DefaultValueInt interface{}) {
	if p_DefaultValueInt != nil {
		ok := false
		nd.DefaultValueInt, ok = p_DefaultValueInt.(int)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return
		}
		defaultvalue_property, ok := nd.PropertyTable[_DefaultValueInt].(property.PropertyT)
		if !ok {
			defaultvalue_property = *property.NewPropertyT(_DefaultValueInt)
		}
		defaultvalue_property.SetValue(p_DefaultValueInt)
		nd.PropertyTable[_DefaultValueInt] = defaultvalue_property
	}
}
func (nd *IntData) SetDataRangeInt(data_range [2]int) error {
	var err error = nil
	if data_range[0] == 0 && data_range[1] == 0 {
		var data_range [2]int = [2]int{-2147483647, 2147483647}
		nd.DataRangeInt = data_range
		datarangeint_property, ok := nd.PropertyTable[_DataRangeInt].(property.PropertyT)
		if !ok {
			datarangeint_property = *property.NewPropertyT(_DataRangeInt)
		}
		datarangeint_property.SetValue(data_range)
		nd.PropertyTable[_DataRangeInt] = datarangeint_property
	} else {
		var f_range [2]int = data_range
		if f_range[0] <= f_range[1] {
			nd.DataRangeInt = f_range
			datarangeint_property, ok := nd.PropertyTable[_DataRangeInt].(property.PropertyT)
			if !ok {
				datarangeint_property = *property.NewPropertyT(_DataRangeInt)
			}
			datarangeint_property.SetValue(data_range)
			nd.PropertyTable[_DataRangeInt] = datarangeint_property
		} else {
			var str_error string = "Data range Error(low:" + strconv.Itoa(f_range[0]) +
				", high:" + strconv.Itoa(f_range[1]) + ")!"
			err = errors.New(str_error)
		}
	}
	nd.DataRangeInt = data_range
	return err
}

//=====运算
func (nd *IntData) CheckRange(check_data int) bool {
	var r_ret = false
	if len(nd.GetDataRangeInt()) == 0 {
		r_ret = true
	} else {
		var f_range = nd.GetDataRangeInt()
		if check_data >= f_range[0] && check_data <= f_range[1] {
			r_ret = true
		} else {
			r_ret = false
		}
	}
	return r_ret
}

func (nd *IntData) Add(p_data interface{}) (int, error) {
	f_leftdata, ok := nd.GetValue().(int)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0, fmt.Errorf("assert error")
	}
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		ok := false
		f_rightdata, ok = v_data.GetValue().(int)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return 0, fmt.Errorf("assert error")
		}
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata + f_rightdata, f_error
}

func (nd *IntData) RAdd(p_data interface{}) (int, error) {
	f_leftdata, ok := nd.GetValue().(int)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0, fmt.Errorf("assert error")
	}
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		ok := false
		f_rightdata, ok = v_data.GetValue().(int)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return 0, fmt.Errorf("assert error")
		}
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata + f_rightdata, f_error
}

func (nd *IntData) Sub(p_data interface{}) (int, error) {
	f_leftdata, ok := nd.GetValue().(int)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0, fmt.Errorf("assert error")
	}
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		ok := false
		f_rightdata, ok = v_data.GetValue().(int)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return 0, fmt.Errorf("assert error")
		}
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata - f_rightdata, f_error
}

func (nd *IntData) RSub(p_data interface{}) (int, error) {
	f_leftdata, ok := nd.GetValue().(int)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0, fmt.Errorf("assert error")
	}
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		ok := false
		f_rightdata, ok = v_data.GetValue().(int)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return 0, fmt.Errorf("assert error")
		}
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_rightdata - f_leftdata, f_error
}

func (nd *IntData) Mul(p_data interface{}) (int, error) {
	f_leftdata, ok := nd.GetValue().(int)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0, fmt.Errorf("assert error")
	}
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		ok := false
		f_rightdata, ok = v_data.GetValue().(int)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return 0, fmt.Errorf("assert error")
		}
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata * f_rightdata, f_error
}

func (nd *IntData) RMul(p_data interface{}) (int, error) {
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

func (nd *IntData) Div(p_data interface{}) (int, error) {
	f_leftdata, ok := nd.GetValue().(int)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0, fmt.Errorf("assert error")
	}
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

func (nd *IntData) RDiv(p_data interface{}) (int, error) {
	f_leftdata, ok := nd.GetValue().(int)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0, fmt.Errorf("assert error")
	}
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

func (nd *IntData) Mod(p_data interface{}) (int, error) {
	f_leftdata, ok := nd.GetValue().(int)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0, fmt.Errorf("assert error")
	}
	var f_rightdata int
	var f_error error
	switch p_data.(type) {
	case GeneralData:
		v_data := p_data.(GeneralData)
		ok := false
		f_rightdata, ok = v_data.GetValue().(int)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return 0, fmt.Errorf("assert error")
		}
	case int:
		f_rightdata = p_data.(int)
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata % f_rightdata, f_error
}

func (nd *IntData) RMod(p_data interface{}) (int, error) {
	f_leftdata, ok := nd.GetValue().(int)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0, fmt.Errorf("assert error")
	}
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

func (nd *IntData) Neg() int {
	var f_leftdata int = nd.GetValue().(int)
	return -f_leftdata
}

func (nd *IntData) Lt(p_data interface{}) (bool, error) {
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

func (nd *IntData) Le(p_data interface{}) (bool, error) {
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

func (nd *IntData) Eq(p_data interface{}) (bool, error) {
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

func (nd *IntData) Ne(p_data interface{}) (bool, error) {
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

func (nd *IntData) Ge(p_data interface{}) (bool, error) {
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

func (nd *IntData) Gt(p_data interface{}) (bool, error) {
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
