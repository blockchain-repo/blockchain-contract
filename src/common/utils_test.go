package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"testing"
)

func Test_GenDate(t *testing.T) {
	dat := GenDate()
	fmt.Println("----------------------dat:", dat)
}

func Test_GenTimestamp(t *testing.T) {
	tim := GenTimestamp()
	fmt.Println("----------------------tim:", tim)
}

func Test_GenSpecialTimestamp(t *testing.T) {
	//the_time, err := time.Parse("2006-01-02 15:04:05", "2014-01-08 09:04:41")
	fullTimeStr := "2017-06-03 08:12:00"
	str, err := GenSpecialTimestamp(fullTimeStr)
	if err != nil {
		fmt.Println("Test_GenSpecialTimestamp error")
		return
	}
	fmt.Println(fullTimeStr, "对应的时间戳为", str)
}

func Test_Deserialize(t *testing.T) {
	jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`
	data := Deserialize(jsonStr)
	fmt.Println("----------------------des:", data)

	jsonStr_Map := `{"host":"http://localhost:9090","name":"test", "age":22}`
	data_Map := Deserialize(jsonStr_Map)
	fmt.Println("----------------------des:", data_Map)

	map_obj_input := map[string]interface{}{"host": "http://localhost:9090", "name": "test", "age": 22}
	map_obj_json := Serialize(map_obj_input)
	map_obj := Deserialize(map_obj_json)
	fmt.Println("----------------------des:", map_obj)
}

func Test_Serialize(t *testing.T) {
	jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "1>=0 && 3+2 <=5 || 4&2 || 2^1","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`
	data := Deserialize(jsonStr)
	//data2 := Serialize(data, false)
	data2 := Serialize(data)
	fmt.Println("----------------------ser:", data2)

	mapStr := map[string]interface{}{"host": "http://localhost:9090", "name": "test", "age": 22}
	data_Map := Serialize(mapStr)
	fmt.Println("----------------------ser:", data_Map)
}

//only for selfTest
func Test_SerializePretty(t *testing.T) {
	jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "1>=0 && 3+2 <=5 || 4&2 || 2^1","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`
	data := Deserialize(jsonStr)
	data3 := SerializePretty(data)
	fmt.Println("----------------------ser:\n", data3)

	mapStr := map[string]interface{}{"host": "http://localhost:9090", "name": "test", "age": 22}
	data_Map := Serialize(mapStr)
	//data_Map := SerializePretty(mapStr)
	fmt.Println("----------------------ser:", data_Map)

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.Encode(jsonStr)
	str := buf.String()
	str = strings.Replace(str, "\\", "", -1)
	t.Log(strings.TrimRight(str, "\""))
	t.Logf("%s", str[1:len(str)-2])
	str2 := string(str)
	t.Log(strings.Trim(str2, "\""))

	aa := "{\"host\": \"http://localhost:9090\",\"port\": 9090,\"analytics_file\": \"1>=2问问的<233 ==23 &23 \",\"static_file_version\": 1,\"static_dir\": \"E:/Project/goTest/src/\",\"templates_dir\": \"E:/Project/goTest/src/templates/\",\"serTcpSocketHost\": \":12340\",\"serTcpSocketPort\": 12340,\"fruits\":[\"apple\", \"peach\"]}"
	t.Log(aa)
	t.Log(strings.Trim(aa, "\""))

}

type C int

func (C) MarshalJSON() ([]byte, error) {
	return []byte(`"<&>"`), nil
}

// CText implements Marshaler and returns unescaped text.
type CText int

func (CText) MarshalText() ([]byte, error) {
	return []byte(`"<&>"`), nil
}

func Test_ArrayToHashSet(t *testing.T) {
	//maps := map[string]interface{}{"name":"wang","age":123123}
	//mapstrs := []string{"name","wang","age","123"}
	//maps["name"] = "wang"
	//maps["age"] = 11111
	//array := []interface{}{"22",242,"4", }

	pub_keys := []interface{}{
		"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
		"22",
		42342,
		nil,
		nil,
	}
	resultHashSet := ArrayToHashSet(pub_keys)
	fmt.Println(resultHashSet)
	fmt.Println(resultHashSet.Contains("qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3"))
	//fmt.Println(common.Serialize(resultHashSet))
	fmt.Println(resultHashSet.Len())
}

func Test_StrArrayToHashSet(t *testing.T) {
	pub_keys := []string{
		"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
	}
	_ = pub_keys
	resultHashSet := StrArrayToHashSet([]string{})
	fmt.Println(resultHashSet)
	fmt.Println(resultHashSet.Contains("qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3"))
	// not exist builtin struct hashSet, so serialize output nil or {}
	//fmt.Println(common.Serialize(resultHashSet))

	pub_keys2 := []string{
		//"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		//"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
		"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet2",
	}
	resultHashSet2 := StrArrayToHashSet(pub_keys2)
	fmt.Println(resultHashSet.Intersect(resultHashSet2))
	fmt.Println(resultHashSet2.Len())
}

func Test_Try(t *testing.T) {
	var err error
	Try(func() {
		panic("我就是为了测试......")
	}, func(e interface{}) {
		fmt.Printf("%+v\n", e)
		err = errors.New(e.(string))
	})
	fmt.Printf("%+v\n", err)
}

func Test_JsonStrTrim(t *testing.T) {
	jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "1>=0 && 3+2 <=5 || 4&2 || 2^1","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	// disabled the HTMLEscape for &, <, and > to \u0026, \u003c, and \u003e in json string
	enc.SetEscapeHTML(false)
	/*--------------- 1. trim the result maybe have the extra double quotation in last--------------------*/
	enc.Encode(jsonStr)

	/*---------------- 2. it`s ok change to interface and encode--------------------------*/
	//var jsonStrData interface{}
	//json.Unmarshal([]byte(jsonStr), &jsonStrData)
	//enc.Encode(jsonStrData)

	strJson := buf.String()
	strJson = strings.TrimSpace(strJson)
	//t.Log("disableEscapeHTML and output len is", len(strJson), ",content is:\n", strJson)
	strJson = strings.Replace(strJson, "\\", "", -1)
	//t.Log("remove the backslash and output len is", len(strJson), ",content is:\n", strJson)
	strJsonWithOutBacklashLen := len(strJson)
	strJson = strings.Trim(strJson, "\"")
	t.Log("after strings trim, the len is", len(strJson), ",expect is", strJsonWithOutBacklashLen-2, ",content is:\n", strJson)
	// the last double quotation except removed
}
