package constdef

import (
	"testing"
)

func TestEventPriority(t *testing.T) {
	var ctype0 int = EventPriority_Unknow
	if ctype0 != 0 {
		t.Error("EventPriority_Unknow value Error!")
	}
	var ctype1 int = EventPriority_Immediate
	if ctype1 != 1 {
		t.Error("EventPriority_Immediate value Error!")
	}
	var ctype2 int = EventPriority_AfterEngineCycle
	if ctype2 != 2 {
		t.Error("EventPriority_AfterEngineCycle value Error!")
	}
}
