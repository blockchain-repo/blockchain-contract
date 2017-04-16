package requestHandler

import (
	"testing"
	"fmt"
)

/**
 * function : 获取yamlconfig对象
 * param   :
 * return : 
 */
func TestGetYamlConfig(t *testing.T) {
	config := GetYamlConfig("unichainApiConf.yaml")
	fmt.Println(config)
}

/**
 * function : 获取对应api的参数
 * param   :
 * return :
 */
func TestGetParam(t *testing.T) {

	config := GetYamlConfig("unichainApiConf.yaml")
	method,url,head,body :=GetParam(config,"ContractAssetFreeze")

	fmt.Println(method)
	fmt.Println(url)
	fmt.Println(head)
	fmt.Println(body)
}

/**
 * function : 获取对应api的参数
 * param   :
 * return :
 */
func TestGetParam1(t *testing.T) {

	config := GetYamlConfig("unicontractApiConf.yaml")
	method,url,head,body :=GetParam(config,"ContractAssetFreeze")

	fmt.Println(method)
	fmt.Println(url)
	fmt.Println(head)
	fmt.Println(body)
}

/**
 * function : 获取对应api的参数
 * param   :
 * return :
 */
func TestGetParam2(t *testing.T) {

	config := GetYamlConfig("unichainApiConf.yaml")
	method,url,head,body :=GetParam(config,"ContractAssetFreeze")

	fmt.Println(method)
	fmt.Println(url)
	fmt.Println(head)
	fmt.Println(body)

	fmt.Println("===================================")
	config1 := GetYamlConfig("unicontractApiConf.yaml")
	method1,url1,head1,body1 :=GetParam(config1,"ContractAssetFreeze")

	fmt.Println(method1)
	fmt.Println(url1)
	fmt.Println(head1)
	fmt.Println(body1)
}
