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

type ArrayData struct {
	GeneralData
	ValueArray []interface{} `json:"-"`
}

const (
	_ValueArray = "ValueArray"
)

func NewArrayData() *ArrayData {
	n := &ArrayData{}
	return n
}

//====================接口方法========================
func (ad ArrayData) GetName() string {
	return ad.GeneralData.GetName()
}

func (ad ArrayData) GetValue() interface{} {
	//uniledgerlog.Debug("++++++++++++++++++++++++++++++++++++++++++++++++++++")
	//uniledgerlog.Debug("我被调用啦～～～～～～～～～～～～～～～～～～～")
	//uniledgerlog.Debug("++++++++++++++++++++++++++++++++++++++++++++++++++++")
	value_property := ad.PropertyTable[_Value].(property.PropertyT)
	return value_property.GetValue()
}

func (ad ArrayData) SetContract(p_contract inf.ICognitiveContract) {
	ad.GeneralComponent.SetContract(p_contract)
}
func (ad ArrayData) GetContract() inf.ICognitiveContract {
	return ad.GeneralComponent.GetContract()
}

func (ad ArrayData) GetCtype() string {
	return ad.GeneralData.GetCtype()
}

func (ad ArrayData) SetValue(p_Value interface{}) {
	//uniledgerlog.Debug("++++++++++++++++++++++++++++++++++++++++++++++++++++")
	//uniledgerlog.Debug("我被调用啦～～～～～～～～～～～～～～～～～～～")
	//uniledgerlog.Debug("++++++++++++++++++++++++++++++++++++++++++++++++++++")
	slData, err := json.Marshal(p_Value)
	if err != nil {
		uniledgerlog.Error("SetValue fail[" + err.Error() + "]")
	}
	ad.GeneralData.SetValue(string(slData))
}
func (ad ArrayData) CleanValueInProcess() {
	ad.GeneralData.CleanValueInProcess()
}

//====================描述态==========================
//序列化： 需要将运行态结构 序列化到 描述态中
func (ad *ArrayData) RunningToStatic() {
	ad.GeneralData.RunningToStatic()
}

func (ad *ArrayData) Serialize() (string, error) {
	ad.RunningToStatic()
	if s_model, err := json.Marshal(ad); err == nil {
		return string(s_model), err
	} else {
		uniledgerlog.Error("Contract Matrix Data fail[" + err.Error() + "]")
		return "", err
	}
}

//====================运行态==========================
func (ad *ArrayData) InitArrayData() error {
	var err error = nil
	err = ad.InitGeneralData()
	if err != nil {
		uniledgerlog.Error("InitArrayData fail[" + err.Error() + "]")
		return err
	}

	v := ad.Value.(string)
	if v != "" {
		//uniledgerlog.Debug("++++++++++++++++++++++++++++++++++++++++++++++++++++")
		//uniledgerlog.Debug("我被调用啦～～～～～～～～～～～～～～～～～～～")
		//uniledgerlog.Debug("++++++++++++++++++++++++++++++++++++++++++++++++++++")
		var interf []interface{}
		err = json.Unmarshal([]byte(v), &interf)
		if err != nil {
			uniledgerlog.Error(err)
			return err
		}
		ad.ValueArray = interf
	}
	common.AddProperty(ad, ad.PropertyTable, _ValueArray, ad.ValueArray)

	ad.SetCtype(constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Array])
	return err
}

func (ad *ArrayData) AppendValue(p_data interface{}) (bool, error) {
	if p_data == nil {
		return false, nil
	}
	value_property, ok := ad.PropertyTable[_Value].(property.PropertyT)
	if !ok {
		value_property = *property.NewPropertyT(_Value)
	}
	var err error = nil
	var a_flag bool = true
	if value_property.GetValue() == nil {
		var v_array []interface{} = make([]interface{}, 0)
		value_property.SetValue(v_array)
	}
	value_property.SetValue(append(value_property.GetValue().([]interface{}), p_data.(interface{})))
	ad.PropertyTable[_Value] = value_property
	ad.Value = value_property
	if ad.Value == nil {
		err = fmt.Errorf("append data Value error!")
		a_flag = false
	}
	return a_flag, err
}

func (ad *ArrayData) RemoveValue(idx int) (bool, error) {
	value_property, ok := ad.PropertyTable[_Value].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return false, nil
	}
	var err error = nil
	var a_flag bool = true
	if value_property.GetValue() == nil {
		err = fmt.Errorf("date Value is nil, remove error!")
		a_flag = false
	} else {
		t_data := make([]interface{}, len(value_property.GetValue().([]interface{})))
		v_data := t_data[0:idx]
		v_data = append(v_data, t_data[idx+1:])
		value_property.SetValue(v_data)
		ad.PropertyTable[_Value] = value_property
		ad.Value = value_property
	}
	return a_flag, err
}

func (ad *ArrayData) GetItem(idx int) (interface{}, error) {
	var err error = nil
	var r_data interface{} = nil
	if ad.GetValue() == nil {
		err = fmt.Errorf("date Value is nil, getitem error!")
	} else {
		if idx >= len(ad.GetValue().([]interface{})) {
			err = fmt.Errorf("index invalid, getitem error!")
		} else {
			r_data = ad.GetValue().([]interface{})[idx]
		}
	}
	return r_data, err
}

func (ad *ArrayData) Len() int {
	var r_len int = 0
	if ad.GetValue() == nil {
		r_len = 0
	} else {
		r_len = len(ad.GetValue().([]interface{}))
	}
	return r_len
}
