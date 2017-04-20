package collectd

import (
	"os"
	"github.com/astaxie/beego"
	"unicontract/src/common/requestHandler"
	"unicontract/src/common"
)

/**
 * function : 
 * param   :
 * return : 
 */
func CreateCollectdConfig(){

	//获取模板配置文件
	configPath := os.Getenv("CONFIGPATH")
	templatePath := configPath + "/collectd.conf.template"
	src,err := os.Open(templatePath)
	if err != nil{
		beego.Error(err.Error())
	}
	defer src.Close()

	//创建配置文件,进行复制
	configPath = configPath + "/collectd.conf"
	collectdConf,err := os.OpenFile(configPath,os.O_WRONLY,0644)
	if err != nil{
		beego.Error(err.Error())
	}
	n,err := collectdConf.Seek(0,os.SEEK_END)
	if err != nil{
		beego.Error(err.Error())
	}

	//获取追加内容
	config := requestHandler.GetYamlConfig("monitorConfig.yaml")
	ip :=common.TypeToString(config["ip"])

	netPlugin :=  `
		<Plugin network>
    		Server "` + ip + `" "25826"
		</Plugin>
	`
	c, err := collectdConf.WriteAt([]byte(netPlugin),n)
	beego.Debug(c)
	if err != nil{
		beego.Error(err.Error())
	}
	defer  collectdConf.Close()
}