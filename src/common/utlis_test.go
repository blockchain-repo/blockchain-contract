package common

import (
	"errors"
	"fmt"
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
	jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`
	data := Deserialize(jsonStr)
	data2 := Serialize(data)
	fmt.Println("----------------------ser:", data2)

	mapStr := map[string]interface{}{"host": "http://localhost:9090", "name": "test", "age": 22}
	data_Map := Serialize(mapStr)
	fmt.Println("----------------------ser:", data_Map)
}

//only for selfTest
func Test_SerializePretty(t *testing.T) {
	jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`
	data := Deserialize(jsonStr)
	data3 := SerializePretty(data)
	fmt.Println("----------------------ser:\n", data3)

	mapStr := map[string]interface{}{"host": "http://localhost:9090", "name": "test", "age": 22}
	data_Map := SerializePretty(mapStr)
	fmt.Println("----------------------ser:", data_Map)
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
