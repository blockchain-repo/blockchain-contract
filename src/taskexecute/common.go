// common
package taskexecute

import (
	"sync"
)

import (
	"unicontract/src/config"
	"unicontract/src/core/model"
)

//---------------------------------------------------------------------------
const (
	_TASKQUEUELEN  = 20
	_THRESHOLD     = 50
	_SLEEPTIME     = 30 // 数据表扫描间隔时间，单位是秒
	_CLEANTIME     = 30 // 数据表清理扫描间隔时间，单位是分钟
	_CLEANDATATIME = 30 // 数据表清理间隔时间，单位是天
)

var (
	gchTaskQueue    chan model.TaskSchedule
	gwgTaskExe      sync.WaitGroup
	gnPublicKeysNum int
	gslPublicKeys   []string
)

//---------------------------------------------------------------------------
func init() {
	gchTaskQueue = make(chan model.TaskSchedule, _TASKQUEUELEN)
	gslPublicKeys = config.GetAllPublicKey()
	gnPublicKeysNum = len(gslPublicKeys)
}

//---------------------------------------------------------------------------
