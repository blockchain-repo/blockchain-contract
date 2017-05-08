package constdef

import (
	"testing"
)

func TestEventType(t *testing.T){
	var ctype0 int = EventType_Unknow
	if ctype0 != 0{
		t.Error("EventType_Unknow value Error!")
	}
	var ctype1 int = EventType_Engine
	if ctype1 != 1{
		t.Error("EventType_Engine value Error!")
	}
	var ctype2 int = EventType_Attribute
	if ctype2 != 2{
		t.Error("EventType_Attribute value Error!")
	}
	var ctype3 int = EventType_Component
	if ctype3 != 3{
		t.Error("EventType_Component value Error!")
	}
}