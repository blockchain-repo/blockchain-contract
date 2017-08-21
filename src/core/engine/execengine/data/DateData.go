package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
)

type DateData struct {
	GeneralData
	Format string `json:"Format"`
	//Value type:string(format)
}

const (
	_Format = "_Format"
)

func NewDateData() *DateData {
	n := &DateData{}
	return n
}

//====================接口方法========================
func (dd DateData) GetName() string {
	return dd.GeneralData.GetName()
}

func (dd DateData) GetValue() interface{} {
	value_property := dd.PropertyTable[_Value].(property.PropertyT)
	return value_property.GetValue()
}
func (dd DateData) SetContract(p_contract inf.ICognitiveContract) {
	dd.GeneralComponent.SetContract(p_contract)
}
func (dd DateData) GetContract() inf.ICognitiveContract {
	return dd.GeneralComponent.GetContract()
}
func (dd DateData) GetCtype() string {
	return dd.GeneralData.GetCtype()
}
func (dd DateData) SetValue(p_Value interface{}) {
	dd.GeneralData.SetValue(p_Value)
}
func (dd DateData) CleanValueInProcess() {
	dd.GeneralData.CleanValueInProcess()
}

//====================描述态==========================
//序列化： 需要将运行态结构 序列化到 描述态中
func (dd *DateData) RunningToStatic() {
	dd.GeneralData.RunningToStatic()
	format_property, ok := dd.PropertyTable[_Format].(property.PropertyT)
	if ok {
		dd.Format, _ = format_property.GetValue().(string)
	}
}

func (dd *DateData) Serialize() (string, error) {
	dd.RunningToStatic()
	if s_model, err := json.Marshal(dd); err == nil {
		return string(s_model), err
	} else {
		uniledgerlog.Error("Contract Date Data fail[" + err.Error() + "]")
		return "", err
	}
}

//====================运行态==========================
func (dd *DateData) InitDateData() error {
	var err error = nil
	err = dd.InitGeneralData()
	if err != nil {
		uniledgerlog.Error("InitGeneralData fail[" + err.Error() + "]")
		return err
	}
	dd.SetCtype(constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Date])
	if dd.Format == "" {
		dd.Format = "2006-01-02 15:04:05"
	}
	common.AddProperty(dd, dd.PropertyTable, _Format, dd.Format)

	//default : now date format
	if dd.GetDefaultValue() == nil {
		dd.SetDefaultValue(time.Unix(time.Now().Unix(), 0).Format(dd.Format))
	}

	var hard_conv_type string = "strToDate"
	dd.SetHardConvType(hard_conv_type)
	return err
}

//====属性Get方法
func (dd *DateData) GetFormat() string {
	format_property, ok := dd.PropertyTable[_Format].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	format_value, ok := format_property.GetValue().(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	return format_value
}

//====属性Set方法
func (dd *DateData) SetFormat(p_formate string) {
	dd.Format = p_formate
	format_property, ok := dd.PropertyTable[_Format].(property.PropertyT)
	if !ok {
		format_property = *property.NewPropertyT(_Format)
	}
	format_property.SetValue(p_formate)
	dd.PropertyTable[_Format] = format_property
}

//=====运算
// param: p_date 对应为format格式的字符串，ex 2017-04-14 16:30:30 400
func (dd *DateData) strToDate(p_date string) (time.Time, error) {
	var v_time time.Time
	var err error = nil
	if p_date == "" {
		err = errors.New("Param is null!")
		return time.Time{}, err
	}
	format_property, ok := dd.PropertyTable[_Format].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return v_time, err
	}
	format_value, ok := format_property.GetValue().(string)
	v_time, err = time.Parse(format_value, p_date)
	return v_time, err
}

//param Value: format string
func (dd *DateData) GetValueInt() (int64, error) {
	var err error = nil
	if dd.GetValue() != nil {
		var v_time time.Time
		v_time, err = time.Parse(dd.GetFormat(), dd.GetValue().(string))
		return v_time.Unix(), err
	} else {
		return 0, errors.New("Value is nil!")
	}
}

//param Value: format string
func (dd *DateData) GetValueFormat() string {
	v_value, ok := dd.GetValue().(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	return v_value
}

func (dd *DateData) Add(p_day int) (time.Time, error) {
	var err error
	if dd.GetValue() != nil {
		day, _ := time.ParseDuration("24h")
		var date_Value_str string = dd.GetValue().(string)
		date_Value_time, err := time.Parse(dd.GetFormat(), date_Value_str)
		if err != nil {
			return time.Time{}, err
		}
		return date_Value_time.Add(day * time.Duration(p_day)), err
	} else {
		err = errors.New("Date Value is nil!")
		return time.Time{}, err
	}
}

func (dd *DateData) Lt(p_date interface{}) (bool, error) {
	var f_leftdata_str string = dd.GetValue().(string)
	f_leftdata, err := time.Parse(dd.GetFormat(), f_leftdata_str)
	if err != nil {
		return false, err
	}
	var f_rightdata time.Time
	var f_error error
	switch p_date.(type) {
	case GeneralData:
		v_data := p_date.(GeneralData)
		f_rightdata, err = time.Parse(dd.GetFormat(), v_data.GetValue().(string))
		if err != nil {
			return false, err
		}
	case time.Time:
		f_rightdata = p_date.(time.Time)
	case string:
		f_rightdata, err = time.Parse(dd.GetFormat(), p_date.(string))
		if err != nil {
			return false, err
		}
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata.Before(f_rightdata), f_error
}

func (dd *DateData) Gt(p_date interface{}) (bool, error) {
	var f_leftdata_str string = dd.GetValue().(string)
	f_leftdata, err := time.Parse(dd.GetFormat(), f_leftdata_str)
	if err != nil {
		return false, err
	}
	var f_rightdata time.Time
	var f_error error
	switch p_date.(type) {
	case GeneralData:
		v_data := p_date.(GeneralData)
		f_rightdata, err = time.Parse(dd.GetFormat(), v_data.GetValue().(string))
		if err != nil {
			return false, err
		}
	case time.Time:
		f_rightdata = p_date.(time.Time)
	case string:
		f_rightdata, err = time.Parse(dd.GetFormat(), p_date.(string))
		if err != nil {
			return false, err
		}
	default:
		f_error = errors.New("Param Type Error!")
	}
	return f_leftdata.After(f_rightdata), f_error
}
