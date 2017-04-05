package common

import (
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

	map_obj_input := map[string]interface{}{"host":"http://localhost:9090","name":"test", "age":22}
	map_obj_json := Serialize(map_obj_input)
	map_obj := Deserialize(map_obj_json)
	fmt.Println("----------------------des:", map_obj)
}

func Test_Serialize(t *testing.T) {
	jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`
	data := Deserialize(jsonStr)
	data2 := Serialize(data)
	fmt.Println("----------------------ser:", data2)

	mapStr := map[string]interface{}{"host":"http://localhost:9090","name":"test", "age":22}
	data_Map := Serialize(mapStr)
	fmt.Println("----------------------ser:", data_Map)
}

//only for selfTest
func Test_SerializePretty(t *testing.T) {
	jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`
	data := Deserialize(jsonStr)
	data3 := _SerializePretty(data)
	fmt.Println("----------------------ser:\n", data3)

	mapStr := map[string]interface{}{"host":"http://localhost:9090","name":"test", "age":22}
	data_Map := _SerializePretty(mapStr)
	fmt.Println("----------------------ser:", data_Map)
}
