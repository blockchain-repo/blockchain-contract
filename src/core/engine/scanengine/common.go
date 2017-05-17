// common
package scanengine

import (
	"sync"
)

import (
	"unicontract/src/config"
	"unicontract/src/core/engine"
	"unicontract/src/core/model"
)

//---------------------------------------------------------------------------
const (
	_CONFIG_FILE_ENV = "CONFIGPATH"
	_HTTP_OK         = 200
)

var (
	gchTaskQueue    chan model.TaskSchedule
	gwgTaskExe      sync.WaitGroup
	gnPublicKeysNum int
	gslPublicKeys   []string
	scanEngineConf  map[interface{}]interface{}
)

//---------------------------------------------------------------------------
func init() {
	scanEngineConf = engine.UCVMConf["ScanEngine"].(map[interface{}]interface{})
	gchTaskQueue = make(chan model.TaskSchedule, scanEngineConf["task_queue_len"].(int))
	gslPublicKeys = config.GetAllPublicKey()
	gnPublicKeysNum = len(gslPublicKeys)
}

//---------------------------------------------------------------------------
