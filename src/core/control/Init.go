// Init
package control

import (
	"os"
)

import (
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/common/yaml"
)

//---------------------------------------------------------------------------
const (
	_CONFIG_FILE_NAME = "faultToleranceConfig.yaml"
	_CONFIG_FILE_ENV  = "CONFIGPATH"
)

var (
	Conf map[interface{}]interface{}
)

//---------------------------------------------------------------------------
func Init() {
	strConfigOSPath := os.Getenv(_CONFIG_FILE_ENV)
	strConfigPath := strConfigOSPath + string(os.PathSeparator) + _CONFIG_FILE_NAME
	if err := yaml.Read(strConfigPath, &Conf); err != nil {
		uniledgerlog.Error(err)
		os.Exit(-1)
	}
}

//---------------------------------------------------------------------------
