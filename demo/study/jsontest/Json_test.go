package jsontest

import (
	"testing"
	"encoding/json"
	"fmt"
)

func TestJson(t *testing.T ) {
	var str_json string = "{\"a\":{\"name\":\"test name\", \"value\": \"test value\"}, \"description\":\"test description\"}"
	//json 转 struct
	var v_map map[string]interface{}
	if err := json.Unmarshal([]byte(str_json), &v_map); err == nil {
		fmt.Println("=====json to map: =====")
		fmt.Println(v_map)
		fmt.Println(v_map["a"].(map[string]interface{})["name"], v_map["a"].(map[string]interface{})["value"], v_map["description"])
	} else {
		fmt.Println(err)
	}
	//struct 转 json
	var v_b_map B = B{A{"test name", "test value"},"test description"}
	if b_json,err := json.Marshal(v_b_map);err == nil {
		fmt.Println("=====map to json: ======")
		fmt.Println(string(b_json))
	}else {
		fmt.Println(err)
	}
	var v_bb_map BB = BB{AA{"test name", "test value"},"test description"}
	if bb_json,err := json.Marshal(v_bb_map);err == nil {
		fmt.Println("=====map to json: ======")
		fmt.Println(string(bb_json))
	}else {
		fmt.Println(err)
	}
}
