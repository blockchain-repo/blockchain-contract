package constdef

import (
	"testing"
)

func TestComponentType(t *testing.T) {
	var ctype0 int = Component_Unknown
	if ctype0 != 0 {
		t.Error("Component_Unknown value Error!")
	}
	var ctype1 int = Component_Contract
	if ctype1 != 1 {
		t.Error("Component_Contract value Error!")
	}
	var ctype2 int = Component_Task
	if ctype2 != 2 {
		t.Error("Component_Task value Error!")
	}
	var ctype3 int = Component_Data
	if ctype3 != 3 {
		t.Error("Component_Data value Error!")
	}
	var ctype4 int = Component_Expression
	if ctype4 != 4 {
		t.Error("Component_Expression value Error!")
	}
}
