package task

import (
	"testing"
	"unicontract/src/core/engine/execengine/constdef"
)

func TestGeneralTask(t *testing.T){
	g_task := &GeneralTask{}
	g_task.InitGeneralTask()
	g_task.SetCname("GeneralTask")
	g_task.SetCaption("GeneralTask")
	g_task.SetDescription("GeneralTask Test")
	g_task.SetState(constdef.TaskState[constdef.TaskState_Dormant])

	if g_task.GetName() != "GeneralTask" || g_task.Cname != "GeneralTask"{
		t.Error("InitGeneralTask Error,GetName Error!")
	}
	if g_task.GetCtype() != constdef.ComponentType[constdef.Component_Task] || g_task.GeneralComponent.Ctype != constdef.ComponentType[constdef.Component_Task]{
		t.Error("InitGeneralTask Error,GetCtype Error!")
	}
	if g_task.GetCaption() != "GeneralTask" || g_task.Caption != "GeneralTask"{
		t.Error("InitGeneralTask Error,GetCaption Error!")
	}
	if g_task.GetDescription() != "GeneralTask Test" || g_task.GeneralComponent.Description != "GeneralTask Test"{
		t.Error("InitGeneralTask Error,GetDescription Error!")
	}
	if g_task.GetState() != constdef.TaskState[constdef.TaskState_Dormant] || g_task.State != constdef.TaskState[constdef.TaskState_Dormant]{
		t.Error("InitGeneralTask Error,GetState Error!")
	}
	if len(g_task.GetNextTasks()) != 0 {
		t.Error("InitGeneralTask Error,GetNextTasks Erro!")
	}
	//Test AddTask
	g_task.AddNextTasks("task_1")
	g_task.AddNextTasks("task_2")
	if len(g_task.GetNextTasks()) != 2 {
		t.Error("AddTask Error!")
	}
}
