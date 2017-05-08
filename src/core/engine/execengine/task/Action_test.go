package task

import (
	"testing"
	"unicontract/src/core/engine/execengine/constdef"
)

func TestAction(t *testing.T){
	action := &Action{}
	action.InitAction()
	action.SetCname("Action")
	action.SetCaption("Action")
	action.SetDescription("Action Test")

	if action.GetName() != "Action" {
		t.Error("InitAction getName Erro!")
	}
	if action.GetCtype() != constdef.ComponentType[constdef.Component_Task] +"."+constdef.TaskType[constdef.Task_Action]{
		t.Error("InitAction GetCtype Erro!")
	}
	if action.GetCaption() != "Action" {
		t.Error("InitAction GetCaption Erro!")
	}
	if action.GetDescription() != "Action Test" {
		t.Error("InitAction GetDescription Erro!")
	}
}
