// scantaskschedule_test
package scanengine

import (
	"testing"
	"unicontract/src/config"
)

func Test_Start(t *testing.T) {
	config.Init()

	Start()
}
