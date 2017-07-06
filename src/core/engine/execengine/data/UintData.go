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

//支持int中的各种类型uint8, uint16, uint32, uint64; 不可直接用unit
type UintData struct {
	GeneralData
	ValueUint        uint    `json:"ValueUint"`
	DefaultValueUint uint    `json:"DefaultValueUint"`
	DataRangeUint    [2]uint `json:"DataRangeUint"`
}

const (
	_DataRangeUint    = "_DataRangeUint"
	_ValueUint        = "_ValueUint"
	_DefaultValueUint = "_DefaultValueUint"
)

func NewUintData() *UintData {
	n := &UintData{}
	return n
}

//====================接口方法========================
func (ud UintData) GetName() string {
	return ud.GeneralData.GetName()
}

func (ud UintData) GetValue() interface{} {
	return ud.GetValueUint()
}

func (ud UintData) GetContract() inf.ICognitiveContract {
	return ud.GeneralComponent.GetContract()
}
func (ud UintData) SetContract(p_contract inf.ICognitiveContract) {
	ud.GeneralComponent.SetContract(p_contract)
}
func (ud UintData) GetCtype() string {
	if ud.PropertyTable["_Ctype"] == nil {
		return ""
	}
	ctype_property := ud.PropertyTable["_Ctype"].(property.PropertyT)
	return ctype_property.GetValue().(string)
}
func (ud UintData) SetValue(p_Value interface{}) {
	ud.SetValueUint(p_Value)
}
func (ud UintData) CleanValueInProcess() {
	ud.GeneralData.CleanValueInProcess()
	ud.SetValueUint(0)
	ud.SetDefaultValueUint(0)
}

//====================描述态==========================
//序列化： 需要将运行态结构 序列化到 描述态中
func (ud *UintData) RunningToStatic() {
	ud.GeneralData.RunningToStatic()
	valueUint_property, ok := ud.PropertyTable[_ValueUint].(property.PropertyT)
	if ok {
		ud.ValueUint, _ = valueUint_property.GetValue().(uint)
	}
	defaultValueUint_property, ok := ud.PropertyTable[_DataRangeUint].(property.PropertyT)
	if ok {
		ud.DefaultValueUint, _ = defaultValueUint_property.GetValue().(uint)
	}
	dtaRangeUint_property, ok := ud.PropertyTable[_DataRangeUint].(property.PropertyT)
	if ok {
		ud.DataRangeUint, _ = dtaRangeUint_property.GetValue().([2]uint)
	}
}

func (ud *UintData) Serialize() (string, error) {
	ud.RunningToStatic()
	if s_model, err := json.Marshal(ud); err == nil {
		return string(s_model), err
	} else {
		uniledgerlog.Error("Contract Uint Data fail[" + err.Error() + "]")
		return "", err
	}
}

//====================运行态==========================
func (ud *UintData) InitUintData() error {
	var err error = nil
	err = ud.InitGeneralData()
	if err != nil {
		uniledgerlog.Error("InitUintData fail[" + err.Error() + "]")
		return err
	}
	ud.SetCtype(constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Numeric_Uint])
	fmt.Println(ud.GetCtype())
	var data_range [2]uint = [2]uint{0, 2147483647}
	if ud.DataRangeUint[0] == 0 && ud.DataRangeUint[1] == 0 {
		common.AddProperty(ud, ud.PropertyTable, _DataRangeUint, data_range)
	} else {
		common.AddProperty(ud, ud.PropertyTable, _DataRangeUint, ud.DataRangeUint)
	}
	common.AddProperty(ud, ud.PropertyTable, _ValueUint, ud.ValueUint)
	common.AddProperty(ud, ud.PropertyTable, _DefaultValueUint, ud.DefaultValueUint)
	ud.SetHardConvType("uint")
	return err
}

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func (ud *UintData) GetDataRangeUint() [2]uint {
	datarange_property := ud.PropertyTable[_DataRangeUint].(property.PropertyT)
	return datarange_property.GetValue().([2]uint)
}
func (ud *UintData) GetValueUint() interface{} {
	value_property := ud.PropertyTable[_ValueUint].(property.PropertyT)
	if value_property.GetValue() != nil {
		return value_property.GetValue()
	} else {
		v_default := ud.GetDefaultValueUint()
		return v_default
	}
}
func (ud *UintData) GetDefaultValueUint() interface{} {
	value_property := ud.PropertyTable[_DefaultValueUint].(property.PropertyT)
	if value_property.GetValue() != nil {
		return value_property.GetValue()
	}
	return nil
}

func (ud *UintData) SetDataRangeUint(data_range [2]uint) error {
	var err error = nil
	if data_range[0] == 0 && data_range[1] == 0 {
		var data_range = [2]uint{0, 2147483647}
		ud.DataRangeUint = data_range
		datarange_property := ud.PropertyTable[_DataRangeUint].(property.PropertyT)
		datarange_property.SetValue(data_range)
		ud.PropertyTable[_DataRangeUint] = datarange_property
	} else {
		var f_range [2]uint = data_range
		if f_range[0] < 0 || f_range[1] < 0 {
			err = errors.New("range must > 0")
		} else if f_range[0] <= f_range[1] {
			ud.DataRangeUint = f_range
			datarange_property := ud.PropertyTable[_DataRangeUint].(property.PropertyT)
			datarange_property.SetValue(data_range)
			ud.PropertyTable[_DataRangeUint] = datarange_property
		} else {
			var str_error string = "Data range Error(low:" + strconv.FormatUint(uint64(f_range[0]), 10) +
				", high:" + strconv.FormatUint(uint64(f_range[1]), 10) + ")!"
			err = errors.New(str_error)
		}
	}
	ud.DataRangeUint = data_range
	return err
}
func (ud *UintData) SetValueUint(p_ValueUint interface{}) {
	if p_ValueUint != nil {
		ud.ValueUint = p_ValueUint.(uint)
		value_property := ud.PropertyTable[_ValueUint].(property.PropertyT)
		value_property.SetValue(p_ValueUint)
		ud.PropertyTable[_ValueUint] = value_property
	}
}

func (ud *UintData) SetDefaultValueUint(p_DefaultValueUint interface{}) {
	if p_DefaultValueUint != nil {
		ud.DefaultValueUint = p_DefaultValueUint.(uint)
		defaultvalue_property := ud.PropertyTable[_DefaultValueUint].(property.PropertyT)
		defaultvalue_property.SetValue(p_DefaultValueUint)
		ud.PropertyTable[_DefaultValueUint] = defaultvalue_property
	}
}
func (ud *UintData) CheckRange(check_data uint) bool {
	var r_ret = false
	if len(ud.GetDataRangeUint()) == 0 {
		r_ret = true
	} else {
		var f_range = ud.GetDataRangeUint()
		if check_data >= f_range[0] && check_data <= f_range[1] {
			r_ret = true
		} else {
			r_ret = false
		}
	}
	return r_ret
}

func (ud *UintData) Add(p_data interface{}) (uint, error) {
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

func (ud *UintData) RAdd(p_data interface{}) (uint, error) {
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

func (ud *UintData) Sub(p_data interface{}) (uint, error) {
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

func (ud *UintData) RSub(p_data interface{}) (uint, error) {
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

func (ud *UintData) Mul(p_data interface{}) (uint, error) {
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

func (ud *UintData) RMul(p_data interface{}) (uint, error) {
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

func (ud *UintData) Div(p_data interface{}) (uint, error) {
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

func (ud *UintData) RDiv(p_data interface{}) (uint, error) {
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

func (ud *UintData) Mod(p_data interface{}) (uint, error) {
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

func (ud *UintData) RMod(p_data interface{}) (uint, error) {
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

func (ud *UintData) Lt(p_data interface{}) (bool, error) {
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

func (ud *UintData) Le(p_data interface{}) (bool, error) {
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

func (ud *UintData) Eq(p_data interface{}) (bool, error) {
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

func (ud *UintData) Ne(p_data interface{}) (bool, error) {
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

func (ud *UintData) Ge(p_data interface{}) (bool, error) {
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

func (ud *UintData) Gt(p_data interface{}) (bool, error) {
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
