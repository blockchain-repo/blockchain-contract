package data

import (
	"testing"
	"fmt"
	"unicontract/src/core/engine/execengine/constdef"
)

func CreateTextDataObject() *TextData{
	t_text := new(TextData)
	t_text.InitTextData()
    t_text.SetCname("TestText")
	t_text.SetCaption("text")
	t_text.SetDescription("text description")
	t_text.SetDefaultValue("")
	t_text.SetUnit("")
	return t_text
}

func TestTextInit(t *testing.T){
	t_text := &TextData{}
	t_text.InitTextData()
	t_text.SetCname("TestText")
	t_text.SetCaption("text")
	t_text.SetDescription("text description")
	t_text.SetDefaultValue("")
	t_text.SetUnit("")
	if t_text == nil {
		t.Error("TextData init Error!")
	}
	fmt.Println(t_text.GetCtype())
	fmt.Println(t_text.GetHardConvType())
	if t_text.GetCtype() != constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Text] {
		t.Error("t_ctype value Error!")
	}
	if t_text.GetCname() != "TestText" {
		t.Error("t_name value Error!")
	}
	if t_text.GetCaption() != "text" {
		t.Error("t_caption value Error!")
	}
	if t_text.GetDescription() != "text description" {
		t.Error("t_description value Error!")
	}
	if t_text.GetHardConvType() != "string" {
		t.Error("t_hardconvtype value Error")
	}
}

func TestTextEq(t *testing.T) {
	t_text := CreateTextDataObject()
	t_text.SetValue("test_eq")
	var v_text_eq string = "test_eq"
	if v_bool,err := t_text.Eq(v_text_eq) ; err == nil {
		if !v_bool {
			t.Error("TextData Eq error,result not equal true!")
		} else {
			fmt.Println(t_text.GetValue().(string), " equals ", v_text_eq, ": ", v_bool)
		}
	} else {
		t.Error("TextData Eq error!")
	}
}

func TestTextAdd(t *testing.T) {
	t_text := CreateTextDataObject()
	t_text.SetValue("test_add")
	var v_text_add string = " and test_check"
	if v_res,err := t_text.Add(v_text_add) ; err == nil {
		if v_res != "test_add and test_check" {
			t.Error("TextData Add error!")
		} else {
			fmt.Println(t_text.GetValue(), " add ", v_text_add, ": ", v_res)
		}
	} else {
		t.Error("TextData Add error!")
	}
}

func TestTextRAdd(t *testing.T) {
	t_text := CreateTextDataObject()
	t_text.SetValue("test_add")
	var v_text_add string = " and test_check "
	if v_res,err := t_text.RAdd(v_text_add) ; err == nil {
		if v_res != " and test_check test_add" {
			t.Error("TextData Add error!")
		} else {
			fmt.Println(v_text_add, " add ", t_text.GetValue(), ": ", v_res)
		}
	} else {
		t.Error("TextData Add error!")
	}
}

func TestTextLen(t *testing.T) {
	t_text := CreateTextDataObject()
	t_text.SetValue("test_len")
	if t_text.Len() != len(t_text.GetValue().(string)) {
		t.Error("IntData Le error!")
	}
}