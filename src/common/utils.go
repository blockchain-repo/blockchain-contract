package common

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"strconv"
	"time"
	"unicontract/src/common/basic"
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

func StructToMap(obj interface{}) (map[string]interface{}, error) {
	var mapObj map[string]interface{}
	objBytes, err := json.Marshal(obj)
	if err != nil {
		log.Fatalf(err.Error())
		return mapObj, err
	}
	json.Unmarshal(objBytes, &mapObj)
	return mapObj, err
}

func MapToStruct(mapObj map[string]interface{}) (interface{}, error) {
	var obj interface{}
	mapObjBytes, err := json.Marshal(mapObj)
	if err != nil {
		log.Fatalf(err.Error())
		return obj, err
	}
	json.Unmarshal(mapObjBytes, &obj)
	return obj, err
}

/*
The json package always orders keys when marshalling. Specifically:

Maps have their keys sorted lexicographically.
Structs keys are marshalled in the order defined in the struct

*/
/*------------------------------ struct serialize must use this -----------------------------*/
/*------------------------------ Hash and Sign use this -----------------------------*/
func StructSerialize(obj interface{}) string {
	objMap, err := StructToMap(obj)
	if err != nil {
		log.Fatalf(err.Error())
		return ""
	}
	str, err := json.Marshal(objMap)
	if err != nil {
		log.Fatalf(err.Error())
		return ""
	}
	return string(str)
}

//only for selfTest, format json output
func StructSerializePretty(obj interface{}) string {
	objMap, err := StructToMap(obj)
	if err != nil {
		log.Fatalf(err.Error())
		return ""
	}
	input, err := json.Marshal(objMap)
	if err != nil {
		log.Fatalf(err.Error())
		return ""
	}
	var out bytes.Buffer
	err = json.Indent(&out, input, "", "\t")

	if err != nil {
		log.Fatalf(err.Error())
	}
	return string(out.String())
}

/*------------- Structs keys are marshalled in the order defined in the struct ------------------*/
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

func TypeToString(name interface{}) string {

	value, ok := name.(string)
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

func TypeToInt(name interface{}) int {

	value, ok := name.(int)
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

func TypeToFloat32(name interface{}) float32 {

	value, ok := name.(float32)
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

func TypeToFloat64(name interface{}) float64 {

	value, ok := name.(float64)
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

func TypeToMap(name interface{}) map[interface{}]interface{} {

	value, ok := name.(map[interface{}]interface{})
	if !ok {
		log.Fatal("Type conversion error")
	}
	return value
}

// UUID
func GenerateUUID() string {
	return uuid.New().String()
}

// 数组内部具体类型必须为基本类型不可以是，结构体，数组或指针等复杂类型！
// array content must be simple style
func ArrayToHashSet(array []interface{}) *basic.HashSet {
	hashSet := basic.NewHashSet()
	if len(array) == 0 {
		return hashSet
	}

	for _, obj := range array {
		hashSet.Add(obj)
	}
	return hashSet
}

func StrArrayToHashSet(array []string) *basic.HashSet {
	hashSet := basic.NewHashSet()
	if len(array) == 0 {
		return hashSet
	}

	for _, obj := range array {
		hashSet.Add(obj)
	}
	return hashSet
}

/*
 * try...catch 类似函数，起到获取异常作用
 * \param [in]
 * \return
 */
func Try(execFunc func(), afterPanic func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			afterPanic(err)
		}
	}()
	execFunc()
}

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
