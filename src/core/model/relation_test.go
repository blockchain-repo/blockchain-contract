package model

import (
	"github.com/astaxie/beego/logs"
	"testing"
	"unicontract/src/common"
)

func TestGenerateRelation(t *testing.T) {
	r := Relation{}
	contracHashId := "1"
	contractid := "2"
	taskid := "3"
	taskExecuteIdx := 4

	r.GenerateRelation(contracHashId, contractid, taskid, taskExecuteIdx)
	logs.Info(common.StructSerialize(r))
}
