package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
)

func GenDate() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05 PM")
}

func GenTimestamp() string {
	t := time.Now()
	nanos := t.UnixNano()
	millis := nanos / 1000000 //ms len=13
	return strconv.FormatInt(millis, 10)
}

func GenSpecialTimestamp(fullTimeStr string) (string, error) {
	//the_time, err := time.Parse("2006-01-02 15:04:05", "2014-01-08 09:04:41")
	the_time, err := time.Parse("2006-01-02 15:04:05", fullTimeStr)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	unix_time := the_time.UnixNano() / 1000000
	return strconv.FormatInt(unix_time, 10), nil
}

func Serialize(obj interface{}) string {
	str, err := json.Marshal(obj)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return string(str)
}

//only for selfTest, format json output
func SerializePretty(obj interface{}) string {
	input, err := json.Marshal(obj)
	if err != nil {
		log.Fatalf(err.Error())
	}
	var out bytes.Buffer
	err = json.Indent(&out, input, "", "\t")

	if err != nil {
		log.Fatalf(err.Error())
	}
	return string(out.String())
}

func Deserialize(jsonStr string) interface{} {
	var dat interface{}
	err := json.Unmarshal([]byte(jsonStr), &dat)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return dat
}

//三元运算符计算
func TernaryOperator(p_cond bool, p_true_cond interface{}, p_false_cond interface{}) interface{} {
	if p_cond {
		return p_true_cond
	}
	return p_false_cond
}

//反序列化时用到，将table中的默认值设置到对象中
//TODO: 非数组结构的默认值可以实现，类型为数组的再property_table中对应为map类型，不可直接用
//Args: object   => propertyTable的拥有对象
//      str_name => property对应的name
//      value    => property对应的value
func ReflectSetValue(object interface{}, str_name string, value interface{}) {
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

//Args: object          => propertyTable的拥有对象
//      p_propertyTable => propetyTable名称
//      str_name        => property对应的name
//      value           => property对应的value
//NOTE: importance, need support type,one see log "value type not support!!!"
func AddProperty(object interface{}, p_propertyTable map[string]interface{}, str_name string, value interface{}) property.PropertyT {
	var pro_object property.PropertyT
	if p_propertyTable == nil {
		//TODO
		uniledgerlog.Error("param[p_propertyTable] is nil!!!")
		return pro_object
	}
	if value == nil {
		pro_object = *property.NewPropertyT(str_name)
		p_propertyTable[str_name] = pro_object
		return pro_object
	}
	switch value.(type) {
	case string:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(string))
		p_propertyTable[str_name] = pro_object
		ReflectSetValue(object, str_name, value)
	case uint:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(uint))
		p_propertyTable[str_name] = pro_object
		ReflectSetValue(object, str_name, value)
	case int:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(int))
		p_propertyTable[str_name] = pro_object
		ReflectSetValue(object, str_name, value)
	case bool:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(bool))
		p_propertyTable[str_name] = pro_object
		ReflectSetValue(object, str_name, value)
	case [2]int:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.([2]int))
		p_propertyTable[str_name] = pro_object
		ReflectSetValue(object, str_name, value)
	case [2]uint:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.([2]uint))
		p_propertyTable[str_name] = pro_object
		ReflectSetValue(object, str_name, value)
	case [2]float64:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.([2]float64))
		p_propertyTable[str_name] = pro_object
		ReflectSetValue(object, str_name, value)
	case float64:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(float64))
		p_propertyTable[str_name] = pro_object
		ReflectSetValue(object, str_name, value)
	case time.Time:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(time.Time))
		p_propertyTable[str_name] = pro_object
		ReflectSetValue(object, str_name, value)
	case inf.ICognitiveContract:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(inf.ICognitiveContract))
		p_propertyTable[str_name] = pro_object
		ReflectSetValue(object, str_name, value)
	case OperateResult:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(OperateResult))
		p_propertyTable[str_name] = pro_object
		ReflectSetValue(object, str_name, value)
	case []SelectBranchExpression:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.([]SelectBranchExpression))
		p_propertyTable[str_name] = pro_object
		ReflectSetValue(object, str_name, value)
	case []string:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.([]string))
		p_propertyTable[str_name] = pro_object
	case map[string]string:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(map[string]string))
		p_propertyTable[str_name] = pro_object
	case map[string]int:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(map[string]int))
		p_propertyTable[str_name] = pro_object
	case map[string]inf.IExpression:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(map[string]inf.IExpression))
		p_propertyTable[str_name] = pro_object
	case map[string]inf.IData:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(map[string]inf.IData))
		p_propertyTable[str_name] = pro_object
	case map[string]inf.ITask:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(map[string]inf.ITask))
		p_propertyTable[str_name] = pro_object
	case map[string]interface{}: // 只针对决策(Decision)组件
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(map[string]interface{}))
		p_propertyTable[str_name] = pro_object
	case []interface{}:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.([]interface{}))
		p_propertyTable[str_name] = pro_object
		ReflectSetValue(object, str_name, value)
	default:
		//fmt.Println("[", str_name, ":", value, "]value type not support!!!")
		uniledgerlog.Error("[", str_name, ":", value, "]value type not support!!!")
	}
	return pro_object
}
