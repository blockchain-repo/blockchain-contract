package yaml

import (
	"testing"
)

type T struct {
	A string `yaml:"a"`
	B struct {
		RenamedC bool    `yaml:"c"`
		D        []int   `yaml:",flow"`
		E        float32 `yaml:"e"`
	}
	F int `yaml:"f"`
}

var testConfigName string = "testConfig.yaml"
var testStruct = T{}
var testMap = make(map[interface{}]interface{})

func Test_Unmarshal_to_struct(t *testing.T) {
	if err := Unmarshal(testConfigName, &testStruct); err != nil {
		t.Errorf("Test_Unmarshal_to_struct is failed, err is %v\n", err)
	} else {
		t.Log("Test_Unmarshal_to_struct is pass.")
		t.Logf("Test data is %+v\n", testStruct)
	}
}

func Test_Unmarshal_to_map(t *testing.T) {
	if err := Unmarshal(testConfigName, &testMap); err != nil {
		t.Errorf("Test_Unmarshal_to_map is failed, err is %v\n", err)
	} else {
		t.Log("Test_Unmarshal_to_map is pass.")
		t.Logf("Test data is %+v\n", testMap)
	}
}

func Test_Marshal_from_struct(t *testing.T) {
	if err := Marshal(testConfigName, &testStruct); err != nil {
		t.Errorf("Test_Marshal_from_struct is failed, err is %v\n", err)
	} else {
		t.Log("Test_Marshal_from_struct is pass.")
	}
}

func Test_Marshal_from_map(t *testing.T) {
	if err := Marshal(testConfigName, &testMap); err != nil {
		t.Errorf("Test_Marshal_from_map is failed, err is %v\n", err)
	} else {
		t.Log("Test_Marshal_from_map is pass.")
	}
}
