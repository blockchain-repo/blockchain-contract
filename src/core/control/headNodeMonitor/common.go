// common
package headNodeMonitor

import (
	"sync"
)

import (
	"unicontract/src/config"
	"unicontract/src/core/control"
)

//---------------------------------------------------------------------------

var (
	gslPublicKeys       []string
	gnPublicKeysNum     int
	gwgHeadNodeMonitor  sync.WaitGroup
	headNodeMonitorConf map[interface{}]interface{}
)

//---------------------------------------------------------------------------
func Init() {
	headNodeMonitorConf = control.Conf["MonitorHeadNode"].(map[interface{}]interface{})
	gslPublicKeys = config.GetAllPublicKey()
	gnPublicKeysNum = len(gslPublicKeys)
}

//---------------------------------------------------------------------------
