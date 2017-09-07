package requestHandler

import (
	"os"
	"path/filepath"

	"unicontract/src/common/yaml"
)

var (
	config_file_list = []string{
		"unichainApiConf.yaml",
		"unicontractApiConf.yaml",
	}
	MapConfig map[string]map[interface{}]interface{}
)

func init() {
	MapConfig = make(map[string]map[interface{}]interface{})
	for _, value := range config_file_list {
		MapConfig[value], _ = getYamlConfig(value)
	}
}

func getYamlConfig(yamlName string) (map[interface{}]interface{}, error) {
	//获取环境变量
	requestPath := os.Getenv("CONFIGPATH")
	requestPath = requestPath + string(filepath.Separator) + yamlName
	config := make(map[interface{}]interface{})
	err := yaml.Read(requestPath, config)
	return config, err
}

/**
 * function : function: 获取yamlconfig中内容
 * param   :
 * return : 将数据写到map中并返回
 */
func GetYamlConfig(yamlName string) map[interface{}]interface{} {
	return MapConfig[yamlName]
}
