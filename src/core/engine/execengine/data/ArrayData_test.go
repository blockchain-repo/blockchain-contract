package data

import (
	"fmt"
	"testing"
)

func CreateArrayDataObject() *ArrayData {
	t_array := new(ArrayData)
	t_array.InitArrayData()
	t_array.SetCname("TestArray")
	t_array.SetUnit("")
	t_array.SetCaption("array")
	t_array.SetDescription("test array")

	return t_array
}

func TestArrayInit(t *testing.T) {
	t_array := new(ArrayData)
	t_array.InitArrayData()
	t_array.SetCname("TestArray")
	t_array.SetUnit("")
	t_array.SetCaption("array")
	t_array.SetDescription("test array")

	if t_array == nil {
		t.Error("ArrayData init Error!")
	}
	if t_array.GetCname() != "TestArray" {
		t.Error("t_name value Error!")
	}
	if t_array.GetDefaultValue() != nil {
		t.Error("t_default value Error!")
	}
	if t_array.GetUnit() != "" {
		t.Error("t_unit value Error!")
	}
	if t_array.GetCaption() != "array" {
		t.Error("t_caption value Error!")
	}
	if t_array.GetDescription() != "test array" {
		t.Error("t_description value Error!")
	}
	fmt.Println(t_array.GetCname(), " ", t_array.GetDefaultValue(), " ", t_array.GetUnit(), " ", t_array.GetCaption(), " ", t_array.GetDescription())
}

func TestAppendValue(t *testing.T) {
	t_array := CreateArrayDataObject()
	if v_bool, err := t_array.AppendValue(100); err == nil {
		if v_bool {
			if t_array.Len() != 1 {
				t.Error("Array appendValue Error, get len Error")
			}
		} else {
			t.Error("Array appendValue Error!")
		}
	} else {
		t.Error("Array appendValue Error!")
	}
	if v_bool, err := t_array.AppendValue("test array append"); err == nil {
		if v_bool {
			if t_array.Len() != 2 {
				t.Error("Array appendValue Error, get len Error")
			}
		} else {
			t.Error("Array appendValue Error!")
		}
	} else {
		t.Error("Array appendValue Error!")
	}
	if t_array.GetValue().([]interface{})[0].(int) != 100 {
		t.Error("Array append[0] error!")
	}
	if t_array.GetValue().([]interface{})[1].(string) != "test array append" {
		t.Error("Array append[1] error!")
	}
}

func TestRemoveValue(t *testing.T) {
	t_array := CreateArrayDataObject()
	t_array.AppendValue(100)
	t_array.AppendValue(200)
	t_array.AppendValue(300)
	if v_bool, err := t_array.RemoveValue(1); err == nil {
		if v_bool {
			if t_array.Len() != 2 {
				t.Error("Array RemoveValue Error, get len Error")
			}
		} else {
			t.Error("Array RemoveValue Error!")
		}
	} else {
		t.Error("Array RemoveValue Error!")
	}
}

func TestRemoveValueNull(t *testing.T) {
	t_array := CreateArrayDataObject()
	if _, err := t_array.RemoveValue(1); err == nil {
		t.Error("Array RemoveValue Error!")
	}
	if v_bool, _ := t_array.RemoveValue(2); v_bool {
		t.Error("Array RemoveValue Error!")
	}
}

/*
func TestArrayGetItem(t *testing.T) {
	t_array := CreateArrayDataObject()
	_,err := t_array.GetItem(1)
	if err == nil {
		t.Error("Array is Null, GetItem Error!")
	}
	t_array.AppendValue(100)
	t_array.AppendValue(200)
	t_array.AppendValue(300)
	v_len := t_array.Len()
	v_res,err := t_array.GetItem(1)
	if err != nil {
		t.Error("Array len:", v_len, ",Get Item[1] Error!")
	}
	if v_res != 200 {
		t.Error("Array len:", v_len, ",Get Item[1] Error!")
	}
	_, err1 := t_array.GetItem(4)
	if err1 == nil {
		t.Error("Array len:", v_len, "Get Item[4] out of range!")
	}
}
*/
func TestArrayLen(t *testing.T) {
	t_array := CreateArrayDataObject()
	if t_array.Len() != 0 {
		t.Error("Array is nil, Len[0] Error!")
	}
	t_array.AppendValue(100)
	if t_array.Len() != 1 {
		t.Error("Array len[1] Error!")
	}

}
