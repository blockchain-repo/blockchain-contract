package component

import (
	"encoding/json"

	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
)

//描述态只在序列化和反序列化时使用；过程中都是执行态
//描述态：gc.xxxxx 是描述态获取值的方式；Set方法时，需要给gc.xxxx赋值，保证可以正常序列化出来
//运行态：gc.PropertyTable["xxxx"] 是运行态获取值的方式；Get方法返回的是propertyTable中存储的值
type GeneralComponent struct {
	Cname         string            `json:"Cname"`
	Ctype         string            `json:"Ctype"`
	Caption       string            `json:"Caption"`
	Description   string            `json:"Description"`
	MetaAttribute map[string]string `json:"MetaAttribute"`

	//inf.ICognitiveContract
	Contract      inf.ICognitiveContract `json:"-"`
	PropertyTable map[string]interface{} `json:"-"`
}

//TODO: 文件校验时，检查是否缺失，是否对应
const (
	_Cname         = "_Cname"
	_Ctype         = "_Ctype"
	_Caption       = "_Caption"
	_Description   = "_Description"
	_Contract      = "_Contract"
	_MetaAttribute = "_MetaAttribute"
)

//===============接口实现===================
func (gc *GeneralComponent) GetContract() inf.ICognitiveContract {
	var v_contract inf.ICognitiveContract
	if gc.PropertyTable[_Contract] == nil {
		return v_contract
	}
	contract_property := gc.PropertyTable[_Contract].(property.PropertyT)
	if contract_property.GetValue() == nil {
		return v_contract
	}
	return contract_property.GetValue().(inf.ICognitiveContract)
}

func (gc *GeneralComponent) SetContract(contract inf.ICognitiveContract) {
	gc.Contract = contract
	if gc.PropertyTable[_Contract] == nil {
		gc.PropertyTable[_Contract] = property.PropertyT{}
	}
	contract_property := gc.PropertyTable[_Contract].(property.PropertyT)
	contract_property.SetValue(contract)
	gc.PropertyTable[_Contract] = contract_property
}

func (gc *GeneralComponent) GetName() string {
	return gc.GetCname()
}

func (gc *GeneralComponent) GetCtype() string {
	if gc.PropertyTable[_Ctype] == nil {
		return ""
	}
	ctype_property := gc.PropertyTable[_Ctype].(property.PropertyT)
	return ctype_property.GetValue().(string)
}

//===============描述态=====================
//====ToString方法
func (nc *GeneralComponent) ToString() string {
	str_res, err_res := nc.Serialize()
	if err_res != nil {
		uniledgerlog.Error("Component ToString fail[" + err_res.Error() + "]")
		return ""
	}
	return str_res
}

//序列化： 需要将运行态结构 序列化到 描述态中
func (gc *GeneralComponent) RunningToStatic() {
	cname_property, ok := gc.PropertyTable[_Cname].(property.PropertyT)
	if ok {
		gc.Cname, _ = cname_property.GetValue().(string)
	}
	ctype_property, ok := gc.PropertyTable[_Ctype].(property.PropertyT)
	if ok {
		gc.Ctype, _ = ctype_property.GetValue().(string)
	}
	caption_property, ok := gc.PropertyTable[_Caption].(property.PropertyT)
	if ok {
		gc.Caption, _ = caption_property.GetValue().(string)
	}
	description_property, ok := gc.PropertyTable[_Description].(property.PropertyT)
	if ok {
		gc.Description, _ = description_property.GetValue().(string)
	}
	metaAttribute_property, ok := gc.PropertyTable[_MetaAttribute].(property.PropertyT)
	if ok {
		gc.MetaAttribute, _ = metaAttribute_property.GetValue().(map[string]string)
	}
}

func (gc *GeneralComponent) Serialize() (string, error) {
	gc.RunningToStatic()
	if s_model, err := json.Marshal(gc); err == nil {
		return string(s_model), err
	} else {
		uniledgerlog.Error("Component Serialize fail[" + err.Error() + "]")
		return "", err
	}
}

//===============运行态=====================
func (gc *GeneralComponent) InitGeneralComponent() error {
	var err error = nil
	/*
		if gc.Cname == "" {
			uniledgerlog.Warn("GeneralComponent Need Cname!")
			err = fmt.Errorf("GeneralComponent Need Cname!")
			return err

		}
	*/
	if gc.PropertyTable == nil {
		gc.PropertyTable = make(map[string]interface{}, 0)
	}
	if gc.MetaAttribute == nil {
		gc.MetaAttribute = make(map[string]string, 0)
	}
	//将描述态数据加载成运行态，因此value都是gc.xxxx(描述态的)
	common.AddProperty(gc, gc.PropertyTable, _Cname, gc.Cname)
	gc.Ctype = common.TernaryOperator(gc.Ctype == "", constdef.ComponentType[constdef.Component_Unknown], gc.Ctype).(string)
	common.AddProperty(gc, gc.PropertyTable, _Ctype, gc.Ctype)
	common.AddProperty(gc, gc.PropertyTable, _Caption, gc.Caption)
	common.AddProperty(gc, gc.PropertyTable, _Description, gc.Description)
	common.AddProperty(gc, gc.PropertyTable, _Contract, gc.Contract)
	common.AddProperty(gc, gc.PropertyTable, _MetaAttribute, gc.MetaAttribute)
	return err
}

//获取PropertyTable
//return: map[string]property.propertyT
func (gc *GeneralComponent) GetPropertyTable() map[string]interface{} {
	return gc.PropertyTable
}

//Note: PropertyTable的key为属性变量名大写加_前缀，如：_NAME
//return: property.propertyT
func (gc *GeneralComponent) GetProperty(p_name string) interface{} {
	if p_name != "" && gc.PropertyTable != nil {
		v_property, ok := gc.PropertyTable[p_name].(property.PropertyT)
		if ok {
			return v_property.GetValue()
		}
	}
	return nil
}

//Note:获取PropertyTable中的属性的值，为了保持统一的获取对象元素的方法
//Return: interface{}
func (gc *GeneralComponent) GetItem(p_name string) interface{} {
	if p_name != "" && gc.PropertyTable != nil {
		v_property, ok := gc.PropertyTable[p_name].(property.PropertyT)
		if ok {
			return v_property.GetValue()
		}
	}
	return nil
}

func (gc *GeneralComponent) AddMetaAttribute(metaProperty interface{}) {
	map_metaattribute, ok := metaProperty.(map[string]string)
	if !ok {
		uniledgerlog.Error("AddMetaAttribute fail")
		return
	}
	if metaProperty != nil && len(map_metaattribute) != 0 {
		metaAttribute_property, ok := gc.PropertyTable[_MetaAttribute].(property.PropertyT)
		if ok {
			map_property_metaattribute, m_ok := metaAttribute_property.GetValue().(map[string]string)
			if !m_ok {
				uniledgerlog.Error("AddMetaAttribute fail")
				return
			}
			if metaAttribute_property.GetValue() == nil || len(map_property_metaattribute) == 0 {
				metaAttribute_property.SetValue(make(map[string]string, 0))
			}
			for key, value := range map_metaattribute {
				metaAttribute_property.GetValue().(map[string]string)[key] = value
			}
			gc.PropertyTable[_MetaAttribute] = metaAttribute_property
			gc.MetaAttribute = metaAttribute_property.GetValue().(map[string]string)
		}
	}
}

//====属性Get方法
func (gc *GeneralComponent) GetCname() string {
	var r_res string = ""
	if gc.PropertyTable[_Cname] != nil {
		cname_property := gc.PropertyTable[_Cname].(property.PropertyT)
		r_res, _ := cname_property.GetValue().(string)
		return r_res
	}
	return r_res
}

func (gc *GeneralComponent) GetCaption() string {
	var r_res string = ""
	if gc.PropertyTable[_Caption] != nil {
		caption_property := gc.PropertyTable[_Caption].(property.PropertyT)
		r_res, _ = caption_property.GetValue().(string)
		if gc.Contract != nil {
			r_res = gc.Contract.ProcessString(gc.Caption)
		}
	}
	return r_res
}

func (gc *GeneralComponent) GetDescription() string {
	var r_res string = ""
	if gc.PropertyTable[_Description] != nil {
		description_property := gc.PropertyTable[_Description].(property.PropertyT)
		r_res, _ = description_property.GetValue().(string)
		if gc.Contract != nil {
			r_res = gc.Contract.ProcessString(gc.Description)
		}
	}
	return r_res
}

func (gc *GeneralComponent) GetMetaAttribute() map[string]string {
	if gc.PropertyTable[_MetaAttribute] == nil {
		return nil
	}
	metaattribute_property := gc.PropertyTable[_MetaAttribute].(property.PropertyT)
	r_metaattribute, _ := metaattribute_property.GetValue().(map[string]string)
	return r_metaattribute
}

//属性Set方法
func (gc *GeneralComponent) SetCname(str_name string) {
	gc.Cname = str_name
	cname_property, ok := gc.PropertyTable[_Cname].(property.PropertyT)
	if !ok {
		cname_property = *property.NewPropertyT(_Cname)
	}
	cname_property.SetValue(str_name)
	gc.PropertyTable[_Cname] = cname_property
}

func (gc *GeneralComponent) SetCtype(str_type string) {
	gc.Ctype = str_type
	ctype_property, ok := gc.PropertyTable[_Ctype].(property.PropertyT)
	if !ok {
		ctype_property = *property.NewPropertyT(_Ctype)
	}
	ctype_property.SetValue(str_type)
	gc.PropertyTable[_Ctype] = ctype_property
}

func (gc *GeneralComponent) SetCaption(str_Caption string) {
	gc.Caption = str_Caption
	caption_property, ok := gc.PropertyTable[_Caption].(property.PropertyT)
	if !ok {
		caption_property = *property.NewPropertyT(_Caption)
	}
	caption_property.SetValue(str_Caption)
	gc.PropertyTable[_Caption] = caption_property
}

func (gc *GeneralComponent) SetDescription(str_Description string) {
	gc.Description = str_Description
	description_property, ok := gc.PropertyTable[_Description].(property.PropertyT)
	if !ok {
		description_property = *property.NewPropertyT(_Description)
	}
	description_property.SetValue(str_Description)
	gc.PropertyTable[_Description] = description_property
}

func (gc *GeneralComponent) SetMetaAttribute(p_metaAttribute map[string]string) {
	gc.MetaAttribute = p_metaAttribute
	metaAttribute_property, ok := gc.PropertyTable[_MetaAttribute].(property.PropertyT)
	if !ok {
		metaAttribute_property = *property.NewPropertyT(_MetaAttribute)
	}
	metaAttribute_property.SetValue(p_metaAttribute)
	gc.PropertyTable[_MetaAttribute] = metaAttribute_property
}
