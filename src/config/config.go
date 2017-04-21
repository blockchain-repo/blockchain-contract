package config

import (
	"os/user"
	"github.com/astaxie/beego"
	"os"
	"unicontract/src/common"
	"bufio"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

var Config UnicontractConfig
/**
 * function : 智能合约配置
 */
type UnicontractConfig struct {
	Keypair  Keypair
	Keyrings []string
}

/**
 * function : 公私钥
 */
type Keypair struct{
	Public  string
	Private string
}

/**
 * function : 初始化UnicontractConfig struct
 * param   :
 * return :
 */
func init(){
	//获取当前用户目录
	user,err := user.Current()
	if err != nil{
		beego.Error(err.Error())
	}
	fileName := user.HomeDir + "/.unicontract"

	//读取配置文件
	unicontractFile,err := os.Open(fileName)
	defer unicontractFile.Close()
	if err != nil{
		beego.Error(err.Error())
	}
	unicontractStr,err := ioutil.ReadAll(unicontractFile)
	if err != nil{
		beego.Error(err.Error())
	}
	//获取unicontractConfig 结构体
	var unicontractConfig UnicontractConfig
	err = json.Unmarshal(unicontractStr,&unicontractConfig)
	if err != nil{
		beego.Error(err.Error())
	}
	Config = unicontractConfig
}

/**
 * function : 读取UnicontractConfig struct
 * param   :
 * return :
 */
func ReadUnicontractConfig() string{
	//获取当前用户目录
	user,err := user.Current()
	if err != nil{
		beego.Error(err.Error())
	}
	fileName := user.HomeDir + "/.unicontract"

	//读取配置文件
	unicontractFile,err := os.Open(fileName)
	defer unicontractFile.Close()
	if err != nil{
		beego.Error(err.Error())
	}
	unicontractStr,err := ioutil.ReadAll(unicontractFile)
	if err != nil{
		beego.Error(err.Error())
	}
	return string(unicontractStr)
}

/**
 * function : 初始化UnicontractConfig struct
 * param   :
 * return :
 */
func WriteConToFile(){

	//获取当前用户目录
	user,err := user.Current()
	if err != nil{
		beego.Error(err.Error())
	}
	fileName := user.HomeDir + "/.unicontract"

	//判断文件是否存在
	fileInfo,err := os.Stat(fileName)
	if err == nil{	//文件存在
		fmt.Println("unicontractCOnf exist,Do you want to override it?")
		fmt.Println(" please input y(es) or n(o) ")
		inputReader := bufio.NewReader(os.Stdin)
		p := make([]byte,10)
		for i:=1;i<=3;i++{
			fmt.Println(i)
			n,_ := inputReader.Read(p)
			if n == 0 {
				fmt.Println("input error,please input y(es) or n(o) ")
			}else {
				if string(p) == "y"{
					break
				}else if string(p) == "n" {
					beego.Debug("unicontractCOnf exist,fileInfo:",fileInfo)
					return
				}
			}
		}
	}
	//文件不存在则创建unictractConf
	unictractConf,err := os.Create(fileName)
	if err != nil{
		beego.Error(err.Error())
	}
	var unicontractConfig UnicontractConfig


	pub,priv :=common.GenerateKeyPair()
	unicontractConfig.Keypair.Public = pub
	unicontractConfig.Keypair.Private = priv
	unicontractConfig.Keyrings = []string{}

	unictractStr := common.Serialize(unicontractConfig)
	n,err := unictractConf.Write([]byte(unictractStr))
	if err != nil{
		beego.Error(err.Error())
	}else{
		beego.Debug("crate unictractConfong File success",n)
	}
	defer unictractConf.Close()
}








