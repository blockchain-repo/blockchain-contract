package constdef

import (
	"testing"
)

func TestTaskStateConst(t *testing.T) {
	var ctype0 int = TaskState_Dormant
	if ctype0 != 0 {
		t.Error("TaskState_Dormant value Error!")
	}
	var ctype1 int = TaskState_In_Progress
	if ctype1 != 1 {
		t.Error("TaskState_In_Progress value Error!")
	}
	var ctype2 int = TaskState_Completed
	if ctype2 != 2 {
		t.Error("TaskState_Completed value Error!")
	}
	var ctype3 int = TaskState_Discard
	if ctype3 != 3 {
		t.Error("TaskState_Discard value Error!")
	}
}
