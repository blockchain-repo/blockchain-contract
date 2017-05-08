package property

import (
	"testing"
	"fmt"
	"strconv"
)

func TestInitName(t *testing.T){
	var x,y string = "testing_string", "testing_value"
	p := &PropertyT{Name:x}
	if p == nil {
		t.Error("Init String Property Error!")
	} else {
		fmt.Println("InitAndTostring: " + p.ToString())
	}
	if p.GetName() != x {
		t.Error("Name value error!")
	}
	p.SetValue(y)
	p.SetNewValue(y)
	p.SetOldValue(y)
	if p.GetValue() != y {
		t.Error("Value error!")
	}
	if p.GetValue() != y {
		t.Error("NewValue error!")
	}
	if p.GetValue() != y {
		t.Error("OldValue error!")
	}
}

func TestInitTString(t *testing.T){
	var x,y string = "testing_string", "testing_value"
	p := &PropertyT{x, y,y,y}
	if p == nil {
		t.Error("Init String Property Error!")
	} else {
		fmt.Println("InitAndTostring: " + p.ToString())
	}
}

func TestInitTInt(t *testing.T){
	var x,y int = 100, 200
	p := &PropertyT{strconv.Itoa(x), y,y,y}
	if p == nil {
		t.Error("Init Int Property Error!")
	} else {
		fmt.Println("InitAndTostring: " + p.ToString())
	}
}

func TestInitTBool(t *testing.T){
	var x,y bool = true, true
	p := &PropertyT{strconv.FormatBool(x), y,y,y}
	if p == nil {
		t.Error("Init Bool Property Error!")
	} else {
		fmt.Println("InitAndTostring: " + p.ToString())
	}
}

func TestOpOldvalueTString(t *testing.T){
	var x string = "testing_old_value"
	p := &PropertyT{x, x,x,x}
	p.SetOldValue(x)
	if p.GetOldValue() == "" {
		t.Error("set old_value")
	} else {
		fmt.Println("OpOldValue: " + p.GetOldValue().(string))
	}
}

func TestOpOldvalueTInt(t *testing.T){
	var x int = 100
	p := &PropertyT{strconv.Itoa(x), x,x,x}
	p.SetOldValue(x)
	if p.GetOldValue() == "" {
		t.Error("set old_value")
	} else {
		fmt.Println("OpOldValue: ", p.GetOldValue().(int))
	}
}
func TestOpOldvalueTBool(t *testing.T){
	var x bool = true
	p := &PropertyT{strconv.FormatBool(x), x,x,x}
	p.SetOldValue(x)
	if p.GetOldValue() == "" {
		t.Error("set old_value")
	} else {
		fmt.Println("OpOldValue: ", p.GetOldValue().(bool))
	}
}

func TestOpNewvalueTString(t *testing.T){
	var x string = "testing_new_value"
	p := &PropertyT{x, x,x,x}
	p.SetNewValue(x)
	if p.GetNewValue() == "" {
		t.Error("set new_value")
	} else {
		fmt.Println("OpNewValue: " + p.GetNewValue().(string))
	}
}

func TestOpNewvalueTInt(t *testing.T){
	var x int = 300
	p := &PropertyT{strconv.Itoa(x), x,x,x}
	p.SetNewValue(x)
	if p.GetNewValue() == "" {
		t.Error("set new_value")
	} else {
		fmt.Println("OpNewValue: ", p.GetNewValue().(int))
	}
}

func TestOpNewvalueTBool(t *testing.T){
	var x bool = true
	p := &PropertyT{strconv.FormatBool(x), x,x,x}
	p.SetNewValue(x)
	if p.GetNewValue() == "" {
		t.Error("set new_value")
	} else {
		fmt.Println("OpNewValue: ", p.GetNewValue().(bool))
	}
}