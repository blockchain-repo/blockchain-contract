package pipelines

import (
	"testing"
	"unicontract/src/config"
)

func Test_pipStart(t *testing.T) {
	config.Init()
	starttxElection()
}
