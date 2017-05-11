// common
package scanengine

import (
	"os"
	"sync"
)

import (
	beegoLog "github.com/astaxie/beego/logs"
)

import (
	"unicontract/src/common/yaml"
	"unicontract/src/config"
	"unicontract/src/core/model"
)

//---------------------------------------------------------------------------
type scanEngineParam struct {
	SleepTime     int  `yaml:"sleep_time"`      // 数据表扫描间隔时间，单位是秒
	TaskQueueLen  int  `yaml:"task_queue_len"`  // 待执行队列最大长度
	CleanData     bool `yaml:"clean_data"`      // 是否进行数据清理
	CleanTime     int  `yaml:"clean_time"`      // 数据表清理扫描间隔时间，单位是分钟
	CleanDataTime int  `yaml:"clean_data_time"` // 数据表清理间隔时间，单位是天
}

const (
	_CONFIG_FILE_NAME = "scanEngineConfig.yaml"
	_CONFIG_FILE_ENV  = "CONFIGPATH"
	_HTTP_OK          = 200
)

var (
	gchTaskQueue    chan model.TaskSchedule
	gwgTaskExe      sync.WaitGroup
	gnPublicKeysNum int
	gslPublicKeys   []string
	gParam          scanEngineParam
)

//---------------------------------------------------------------------------
func init() {
	strConfigOSPath := os.Getenv(_CONFIG_FILE_ENV)
	strConfigPath := strConfigOSPath + string(os.PathSeparator) + _CONFIG_FILE_NAME
	if err := yaml.Read(strConfigPath, &gParam); err != nil {
		beegoLog.Error(err)
		os.Exit(-1)
	}

	gchTaskQueue = make(chan model.TaskSchedule, gParam.TaskQueueLen)
	gslPublicKeys = config.GetAllPublicKey()
	gnPublicKeysNum = len(gslPublicKeys)
}

//---------------------------------------------------------------------------
