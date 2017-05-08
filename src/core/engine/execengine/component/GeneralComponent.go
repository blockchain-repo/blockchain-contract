package component

import (
	"fmt"
	"time"

	"unicontract/src/core/engine/execengine/property"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/common"
	"strings"
	"unicontract/src/core/engine/execengine/constdef"
	"reflect"
)

type GeneralComponent struct{
	Cname string  `json:"Cname"`
	Ctype string  `json:"Ctype"`
	Caption string `json:"Caption"`
    Description string `json:"Description"`
	MetaAttribute map[string]string `json:"MetaAttribute"`
    //inf.ICognitiveContract
	Contract inf.ICognitiveContract  `json:"-"`
	//TODO
	//EnactmentComponent property.PropertyT `json:"EnactmentComponent"`
	//TODO
	//Event_handler property.PropertyAttributeEvent `json:"Event_handler"`

	PropertyTable map[string] interface{}  `json:"-"`
}

//TODO: 文件校验时，检查是否缺失，是否对应
const (
	_Cname = "_Cname"
	_Ctype = "_Ctype"
	_Caption =  "_Caption"
	_Description = "_Description"
    _Contract = "_Contract"
    _MetaAttribute = "_MetaAttribute"
)
//===============接口实现===================
func (gc *GeneralComponent) GetContract() inf.ICognitiveContract{
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

func (gc *GeneralComponent) SetContract(contract  inf.ICognitiveContract){
	gc.Contract = contract
	if gc.PropertyTable[_Contract] == nil {
		//TODO: need
	}
	contract_property := gc.PropertyTable[_Contract].(property.PropertyT)
	contract_property.SetValue(contract)
	gc.PropertyTable[_Contract] = contract_property
}

func (gc *GeneralComponent)GetName()string  {
	return gc.GetCname()
}

func (gc *GeneralComponent)GetCtype()string{
	if gc.PropertyTable[_Ctype] == nil {
		return ""
	}
	ctype_property := gc.PropertyTable[_Ctype].(property.PropertyT)
	return ctype_property.GetValue().(string)
}
//===============描述态=====================
//====ToString方法
func (nc *GeneralComponent)ToString() string{
	var str_res string
	str_res = strings.Join([]string{"Cname:", nc.Cname,
									", Ctype:", nc.Ctype,
									", Caption:", nc.Caption,
									", Description:", nc.Description},  "")
	return str_res
}
//===============运行态=====================
func (gc *GeneralComponent) InitGeneralComponent()error{
	var err error = nil
	/*
	if gc.Cname == "" {
		//TODO log
		err = errors.New("GeneralComponent Need Cname!")
		return err
	}*/
	if gc.PropertyTable == nil {
		gc.PropertyTable = make(map[string]interface{}, 0)
	}
	//Take Care: map or [] need init
	if gc.MetaAttribute == nil || len(gc.MetaAttribute) == 0 {
		gc.MetaAttribute = make(map[string]string, 0)
	}
	//Take Care: Value must be gc.xxxxxx, because method init is State[runnable], need use value of State[description]
	gc.AddProperty(gc, _Cname, gc.Cname)
	gc.Ctype = common.TernaryOperator(gc.Ctype == "", constdef.ComponentType[constdef.Component_Unknown], gc.Ctype).(string)
	gc.AddProperty(gc, _Ctype, gc.Ctype)
	gc.AddProperty(gc, _Caption, gc.Caption)
	gc.AddProperty(gc, _Description, gc.Description)
	gc.AddProperty(gc, _Contract, gc.Contract)
	gc.AddProperty(gc, _MetaAttribute, gc.MetaAttribute)
	return err
}
//反序列化时用到，将table中的默认值设置到对象中
//TODO: 非数组结构的默认值可以实现，类型为数组的再property_table中对应为map类型，不可直接用
func (gc *GeneralComponent)ReflectSetValue(object interface{}, str_name string, value interface{}) {
	v_value := reflect.ValueOf(value)
	if reflect.ValueOf(object).Elem().CanSet() {
		mutable := reflect.ValueOf(object).Elem()
		if mutable.FieldByName(strings.Replace(str_name, "_", "", 1)).IsValid() {
			mutable.FieldByName(strings.Replace(str_name, "_", "", 1)).Set(v_value)
		}
	} else {
		mutable := reflect.ValueOf(&object).Elem()
		if mutable.FieldByName(strings.Replace(str_name, "_", "", 1)).IsValid() {
			mutable.FieldByName(strings.Replace(str_name, "_", "", 1)).Set(v_value)
		}
	}
}

//====属性动态初始化
//NOTE: importance, need support type,one see log "value type not support!!!"
func (gc *GeneralComponent) AddProperty(object interface{}, str_name string, value interface{})property.PropertyT {
	var pro_object property.PropertyT
	if value == nil {
		pro_object = *property.NewPropertyT(str_name)
		gc.PropertyTable[str_name] = pro_object
		return pro_object
	}
	switch value.(type) {
	case string:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(string))
		gc.PropertyTable[str_name] = pro_object
		gc.ReflectSetValue(object, str_name, value)
	case uint:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(uint))
		gc.PropertyTable[str_name] = pro_object
		gc.ReflectSetValue(object, str_name, value)
	case int:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(int))
		gc.PropertyTable[str_name] = pro_object
		gc.ReflectSetValue(object, str_name, value)
	case bool:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(bool))
		gc.PropertyTable[str_name] = pro_object
		gc.ReflectSetValue(object, str_name, value)
	case [2]int:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.([2]int))
		gc.PropertyTable[str_name] = pro_object
		gc.ReflectSetValue(object, str_name, value)
	case [2]uint:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.([2]uint))
		gc.PropertyTable[str_name] = pro_object
		gc.ReflectSetValue(object, str_name, value)
	case [2]float64:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.([2]float64))
		gc.PropertyTable[str_name] = pro_object
		gc.ReflectSetValue(object, str_name, value)
	case float64:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(float64))
		gc.PropertyTable[str_name] = pro_object
		gc.ReflectSetValue(object, str_name, value)
	case time.Time:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(time.Time))
		gc.PropertyTable[str_name] = pro_object
		gc.ReflectSetValue(object, str_name, value)
	case inf.ICognitiveContract:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(inf.ICognitiveContract))
		gc.PropertyTable[str_name] = pro_object
		gc.ReflectSetValue(object, str_name, value)
	case common.OperateResult:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(common.OperateResult))
		gc.PropertyTable[str_name] = pro_object
		gc.ReflectSetValue(object, str_name, value)
	case []string:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.([]string))
		gc.PropertyTable[str_name] = pro_object
	case map[string]string:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(map[string]string))
		gc.PropertyTable[str_name] = pro_object
	case map[string]int:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(map[string]int))
		gc.PropertyTable[str_name] = pro_object
	case map[string]inf.IExpression:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(map[string]inf.IExpression))
		gc.PropertyTable[str_name] = pro_object
	case map[string]inf.IData:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(map[string]inf.IData))
		gc.PropertyTable[str_name] = pro_object
	case map[string]inf.ITask:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(map[string]inf.ITask))
		gc.PropertyTable[str_name] = pro_object
	default:
		fmt.Println("[", str_name, ":", value, "]value type not support!!!")
	}
	return pro_object
}
//Note: 获取PropertyTable
//Return: map[string]property.PropertyT
func (gc *GeneralComponent)GetPropertyTable()map[string] interface{}{
	return gc.PropertyTable
}
//Note: PropertyTable的key为属性变量名大写加_前缀，如：_NAME
//Return: property.PropertyT
func (gc *GeneralComponent) GetProperty(p_name string)interface{}{
	if p_name != "" && gc.PropertyTable != nil {
		return gc.PropertyTable[p_name].(property.PropertyT)
	}
	return nil
}
//Note:获取PropertyTable中的属性
//Return: property.PropertyT
func (gc *GeneralComponent)GetItem(p_name string)interface{}{
	if p_name != "" && gc.PropertyTable != nil {
		return gc.PropertyTable[p_name].(property.PropertyT)
	}
	return nil
}

func (gc *GeneralComponent) AddMetaAttribute(metaProperty interface{}){
	if metaProperty != nil && len(metaProperty.(map[string]string)) != 0{
		metaAttribute_property := gc.PropertyTable[_MetaAttribute].(property.PropertyT)
		if metaAttribute_property.GetValue() == nil || len(metaAttribute_property.GetValue().(map[string]string)) == 0{
			metaAttribute_property.SetValue(make(map[string]string, 0))
		}
		v_metaProperty := metaProperty.(map[string]string)
		for key,value := range v_metaProperty {
			metaAttribute_property.GetValue().(map[string]string)[key] = value
		}
		gc.PropertyTable[_MetaAttribute] = metaAttribute_property
		gc.MetaAttribute = metaAttribute_property.GetValue().(map[string]string)
	}
}
//====属性Get方法
func (gc *GeneralComponent)GetCname()string {
	if gc.PropertyTable[_Cname] == nil {
		return ""
	}
	cname_property := gc.PropertyTable[_Cname].(property.PropertyT)
	return cname_property.GetValue().(string)
}

func (gc *GeneralComponent) GetCaption()string {
	var r_res string = ""
	if gc.PropertyTable[_Caption] == nil {
		r_res = ""
	} else {
		caption_property := gc.PropertyTable[_Caption].(property.PropertyT)
		r_res = caption_property.GetValue().(string)
		/*
		if gc.Contract.GetValue() != nil {
			v_Contract := gc.Contract.GetValue().(inf.ICognitiveContract)
			r_res = v_Contract.ProcessString(gc.Caption.GetValue().(string))
		}*/
	}
	return r_res
}

func (gc *GeneralComponent) GetDescription()string{
	var r_res string = ""
	if gc.PropertyTable[_Description] != nil {
		description_property := gc.PropertyTable[_Description].(property.PropertyT)
		r_res = description_property.GetValue().(string)
		/*
		if gc.Contract.GetValue() != nil {
			v_Contract := gc.Contract.GetValue().(inf.ICognitiveContract)
			r_res = v_Contract.ProcessString(gc.Description.GetValue().(string))
		}*/
	}
	return r_res
}

func (gc *GeneralComponent) GetMetaAttribute()map[string]string {
	if gc.PropertyTable[_MetaAttribute] == nil {
		return nil
	}
	metaattribute_property := gc.PropertyTable[_MetaAttribute].(property.PropertyT)
	return metaattribute_property.GetValue().(map[string]string)
}
//属性Set方法
func (gc *GeneralComponent)SetCname(str_name string){
	//Take case: Setter method need set value for gc.xxxxxx
	gc.Cname = str_name
	cname_property := gc.PropertyTable[_Cname].(property.PropertyT)
	cname_property.SetValue(str_name)
	//Take case: Setter method need set value for gc.PropertyTable[xxxx]
	gc.PropertyTable[_Cname] = cname_property
}

func (gc *GeneralComponent)SetCtype(str_type string){
	gc.Ctype = str_type
	ctype_property := gc.PropertyTable[_Ctype].(property.PropertyT)
	ctype_property.SetValue(str_type)
	gc.PropertyTable[_Ctype] = ctype_property
}

func (gc *GeneralComponent) SetCaption(str_Caption string) {
	gc.Caption = str_Caption
	caption_property := gc.PropertyTable[_Caption].(property.PropertyT)
	caption_property.SetValue(str_Caption)
	gc.PropertyTable[_Caption] = caption_property
}

func (gc *GeneralComponent) SetDescription(str_Description string){
	gc.Description = str_Description
	description_property := gc.PropertyTable[_Description].(property.PropertyT)
	description_property.SetValue(str_Description)
	gc.PropertyTable[_Description] = description_property
}

func (gc *GeneralComponent) SetMetaAttribute(p_metaAttribute map[string]string) {
	gc.MetaAttribute = p_metaAttribute
	metaAttribute_property := gc.PropertyTable[_MetaAttribute].(property.PropertyT)
	metaAttribute_property.SetValue(p_metaAttribute)
	gc.PropertyTable[_MetaAttribute] = metaAttribute_property
}