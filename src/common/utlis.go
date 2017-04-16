package common

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"
	"time"
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


/**
 * function : 断言-类型转换string
 * param   :
 * return : 返回string
 */

func TypeToString(name interface{}) string{

	value,ok := name.(string)
	if !ok {
		log.Fatal("Type conversion error")
	}
	return value
}

/**
 * function : 断言-类型转换int
 * param   :
 * return : 返回int
 */

func TypeToInt(name interface{}) int{

	value,ok := name.(int)
	if !ok {
		log.Fatal("Type conversion error")
	}
	return value
}

/**
 * function : 断言-类型转换float32
 * param   :
 * return : 返回int
 */

func TypeToFloat32(name interface{}) float32{

	value,ok := name.(float32)
	if !ok {
		log.Fatal("Type conversion error")
	}
	return value
}

/**
 * function : 断言-类型转换float64
 * param   :
 * return : 返回int
 */

func TypeToFloat64(name interface{}) float64{

	value,ok := name.(float64)
	if !ok {
		log.Fatal("Type conversion error")
	}
	return value
}

/**
 * function : 断言-类型转换map[interface{}]interface{}
 * param   :
 * return : 返回int
 */

func TypeToMap(name interface{}) map[interface{}]interface{}{

	value,ok := name.(map[interface{}]interface{})
	if !ok {
		log.Fatal("Type conversion error")
	}
	return value
}

