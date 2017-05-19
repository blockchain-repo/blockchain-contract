package constdef

import (
	"testing"
)

func TestTaskType(t *testing.T){
	var ctype0 int = Task_Unknown
	if ctype0 != 0{
		t.Error("Task_Unknown value Error!")
	}
	var ctype1 int = Task_Enquiry
	if ctype1 != 1{
		t.Error("Component_Data value Error!")
	}
	var ctype2 int = Task_Action
	if ctype2 !=2{
		t.Error("Task_Action value Error!")
	}
	var ctype3 int = Task_Decision
	if ctype3 != 3{
		t.Error("Task_Decision value Error!")
	}
	var ctype5 int = Task_Plan
	if ctype5 != 5{
		t.Error("Task_Plan value Error!")
	}

	var ctype6 int = Task_DelegateTaskGroup
	if ctype6 != 6{
		t.Error("Task_DelegateTaskGroup value Error!")
	}
}