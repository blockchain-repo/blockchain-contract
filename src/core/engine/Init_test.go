package engine

import (
	"fmt"
	"os"
	"testing"
	"unicontract/src/common/yaml"
)

func TestUCVMconfigure(t *testing.T) {
	strConfigPath := "./" + _CONFIG_FILE_NAME
	if err := yaml.Read(strConfigPath, &UCVMConf); err != nil {
		t.Error(err)
		os.Exit(-1)
	}
	fmt.Println(UCVMConf)
	var ExecuteEngineConf map[interface{}]interface{} = UCVMConf["ExecuteEngine"].(map[interface{}]interface{})
	fmt.Println(ExecuteEngineConf["function_source"])
	fmt.Println(ExecuteEngineConf)

	var ScanEngineConf map[interface{}]interface{} = UCVMConf["ScanEngine"].(map[interface{}]interface{})
	fmt.Println(ScanEngineConf["clean_data_on"])
	fmt.Println(ScanEngineConf["clean_data_time"])
	fmt.Println(ScanEngineConf["clean_time"])
	fmt.Println(ScanEngineConf["sleep_time"])
	fmt.Println(ScanEngineConf["task_queue_len"])
	fmt.Println(ScanEngineConf)
}
