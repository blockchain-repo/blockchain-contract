package constdef

import "testing"

func TestEventTriggerType(t *testing.T){
	var ctype0 int = Engine_Unknown
	if ctype0 != 0{
		t.Error("Engine_Unknown value Error!")
	}
	var ctype1 int = Engine_Initialised
	if ctype1 != 1{
		t.Error("Engine_Initialised value Error!")
	}
	var ctype2 int = Engine_BeforeStarted
	if ctype2 != 2{
		t.Error("Engine_BeforeStarted value Error!")
	}
	var ctype3 int = Engine_AfterStarted
	if ctype3 != 3{
		t.Error("Engine_AfterStarted value Error!")
	}
	var ctype4 int = Engine_BeforeRun
	if ctype4 != 4{
		t.Error("Engine_BeforeRun value Error!")
	}
	var ctype5 int = Engine_AfterRun
	if ctype5 != 5{
		t.Error("Engine_AfterRun value Error!")
	}
	var ctype6 int = Engine_Changed
	if ctype6 != 6{
		t.Error("Engine_Changed value Error!")
	}
}
