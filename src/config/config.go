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
	//切片
	Keyrings []string
}

/**
 * function : 公私钥
 */
type Keypair struct{
	PublicKey  string
	PrivateKey string
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

	//判断文件是否存在
	fileInfo,err := os.Stat(fileName)
	if err != nil{	//文件不存在
		//创建unictractConf
		_CreatUnictractConf(fileName)
	}

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
		beego.Error(err.Error(),fileInfo)
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
		inputReader.Read(p)

		if p[0] != []byte("y")[0]{
			beego.Debug("Give Up override unicontract!",fileInfo)
			return
		}
	}
	//创建unictractConf
	_CreatUnictractConf(fileName)
}

/**
 * function : 创建unictractConf
 * param   :
 * return :
 */
func _CreatUnictractConf(fileName string){
	//文件不存在则创建unictractConf
	unictractConf,err := os.Create(fileName)
	defer unictractConf.Close()

	if err != nil{
		beego.Error(err.Error())
	}
	var unicontractConfig UnicontractConfig

	//获取公私钥匙
	pub,priv :=common.GenerateKeyPair()
	unicontractConfig.Keypair.PublicKey = pub
	unicontractConfig.Keypair.PrivateKey = priv
	//添加公钥环 切片
	unicontractConfig.Keyrings = []string{}

	unictractStr := common.Serialize(unicontractConfig)
	n,err := unictractConf.Write([]byte(unictractStr))
	if err != nil{
		beego.Error(err.Error())
	}else{
		beego.Debug("crate unictractConfong File success",n)
	}
}


/**
 * function : 获取所有节点公钥
 * param   :
 * return :
 */
func GetAllPublicKey() []string{

	keyrings := Config.Keyrings
	publicKey := Config.Keypair.PublicKey

	return append(keyrings,publicKey)
}









