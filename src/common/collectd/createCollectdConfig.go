package collectd

import (
	"os"
	"github.com/astaxie/beego"
	"unicontract/src/common/requestHandler"
	"unicontract/src/common"
	"io"
)

/**
 * function : 
 * param   :
 * return : 
 */
func CreateCollectdConfig(){
	//获取模板配置文件
	configPath := os.Getenv("CONFIGPATH")
	templateConfig :=_GetTemplateConfig(configPath)
	//删除旧的配置文件
	_DeletOldConfigFile(configPath)
	//创建新的配置文件
	collectdConf := _CreatNeWCollectdConf(configPath,templateConfig)
	defer templateConfig.Close()
	//获取追加内容
	netPlugin := _GetContent()
	//文件内容追加
	_AddContentToConf(collectdConf,netPlugin)
}

/**
 * function : 获取模板配置文件
 * param   :
 * return : *os.File
 */
func _GetTemplateConfig(configPath string) *os.File{

	templatePath := configPath + "/collectd.conf.template"
	templateConfig,err := os.Open(templatePath)
	if err != nil{
		beego.Error(err.Error())
	}
	return templateConfig
}

/**
 * function : 删除旧的配置文件
 * param   :
 * return : *os.File
 */
func _DeletOldConfigFile(configPath string) {

	configFile := configPath + "/collectd.conf"
	_,err := os.Stat(configFile)
	if err != nil{
		beego.Error("create new collectd.conf")
	}
	//文件存在则删除
	if !os.IsNotExist(err){
		err := os.Remove(configFile)
		beego.Debug("delete old collectd.conf")
		if err != nil{
			beego.Error(err.Error())
		}
	}
}

/**
 * function : 创建新的配置文件
 * param   :
 * return : *os.File
 */
func _CreatNeWCollectdConf(configPath string,templateConfig *os.File) *os.File{
	configFile := configPath + "/collectd.conf"
	collectdConf,err := os.Create(configFile)
	if err != nil{
		beego.Error(err.Error())
	}
	io.Copy(collectdConf,templateConfig)
	return collectdConf
}

/**
 * function : 获取追加内容
 * param   :
 * return : string
 */
func _GetContent() string{
	config := requestHandler.GetYamlConfig("monitorConfig.yaml")
	ip :=common.TypeToString(config["ip"])

	netPlugin :=  `
<Plugin network>
    Server "` + ip + `" "25826"
</Plugin>
`
	return netPlugin
}

/**
 * function : 文件内容追加
 * param   :
 * return : string
 */
func _AddContentToConf(collectdConf *os.File,netPlugin string){
	n,err := collectdConf.Seek(0,os.SEEK_END)
	if err != nil{
		beego.Error(err.Error())
	}
	c, err := collectdConf.WriteAt([]byte(netPlugin),n)
	if err != nil{
		beego.Error(err.Error())
	}
	beego.Debug("create collectd.conf success...",c)
}

