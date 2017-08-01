// common
package scanengine

import (
	"sync"
)

import (
	"unicontract/src/config"
	"unicontract/src/core/engine"
	"unicontract/src/core/engine/common/db"
)

//---------------------------------------------------------------------------
const (
	_CONFIG_FILE_ENV = "CONFIGPATH"
	_HTTP_OK         = 200
)

var (
	gchTaskQueue      chan db.TaskSchedule
	gchExecParamQueue chan executeParam
	gwgTaskExe        sync.WaitGroup
	gnPublicKeysNum   int
	gslPublicKeys     []string
	scanEngineConf    map[interface{}]interface{}
)

//---------------------------------------------------------------------------
func Init() {
	scanEngineConf = engine.UCVMConf["ScanEngine"].(map[interface{}]interface{})
	gchTaskQueue = make(chan db.TaskSchedule, scanEngineConf["task_queue_len"].(int))
	gchExecParamQueue = make(chan executeParam, scanEngineConf["task_queue_len"].(int))
	gslPublicKeys = config.GetAllPublicKey()
	gnPublicKeysNum = len(gslPublicKeys)
}

//---------------------------------------------------------------------------
