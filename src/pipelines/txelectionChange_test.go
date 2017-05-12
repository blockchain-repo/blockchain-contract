package pipelines

import (
	"testing"
	"time"
	"github.com/astaxie/beego/logs"
)

func TestStart(t *testing.T) {
	start()
	time.Sleep(time.Second * 2)
	logs.Info("down")
}
