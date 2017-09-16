package data

import (
	"testing"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
)

//func (gd *GeneralData) InitGeneralData(str_cname string, default_value  interface{}, str_unit string,
//str_ctype string, str_caption string, str_description string){

func CreateGeneralDataObject() *GeneralData {
	t_data := new(GeneralData)
	t_data.InitGeneralData()

	t_data.SetCname("TestGeneralData")
	t_data.SetCaption("general")
	t_data.SetDescription("test general")
	t_data.SetUnit("")

	return t_data
}

func TestGeneralDataInit(t *testing.T) {
	t_data := new(GeneralData)
	t_data.InitGeneralData()

	t_data.SetCname("TestGeneralData")
	t_data.SetCaption("general")
	t_data.SetDescription("test general")
	t_data.SetUnit("")

	if t_data == nil {
		t.Error("GeneralData init Error!")
	}
	if t_data.GetName() != "TestGeneralData" {
		t.Error("GeneralData getName Error!")
	}
	if t_data.GetCaption() != "general" {
		t.Error("GeneralData getCaption Error!")
	}
	if t_data.GetDescription() != "test general" {
		t.Error("GeneralData getDescription Error!")
	}
	if t_data.GetCtype() != constdef.ComponentType[constdef.Component_Data] {
		t.Error("GeneralData getCtype Error!")
	}
}

func TestGetName(t *testing.T) {
	t_data := CreateGeneralDataObject()
	t_data.SetCname("TestGeneralData")
	t_data.SetCaption("generalData")
	t_data.SetDescription("test general data")
	t_data.SetUnit("")

	i_parent := &GeneralData{}
	err := i_parent.InitGeneralData()
	i_parent.SetCname("TestParentData")
	if err != nil {
		t.Error("InitGeneralData Error!")
	}
	t_data.SetParent(i_parent)
	if t_data.GetName() != "TestParentData.TestGeneralData" {
		t.Error("GeneralData getName Error!")
	}
	var t_parent_2 inf.IData = new(GeneralData)
	t_data.SetParent(t_parent_2)
	if t_data.GetName() != "TestGeneralData" {
		t.Error("GeneralData getName nil Error")
	}
}
