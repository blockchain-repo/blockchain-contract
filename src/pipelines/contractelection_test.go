// contractelection_test
package pipelines

import (
	"testing"
	"unicontract/src/config"
)

func Test_startContractElection(t *testing.T) {
	config.Init()
	startContractElection()
}
