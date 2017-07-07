package model

import (
	"unicontract/src/common/uniledgerlog"
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
	uniledgerlog.Info(common.StructSerialize(r))
}
