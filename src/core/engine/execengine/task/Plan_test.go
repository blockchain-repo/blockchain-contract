package task

import (
	"testing"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/constdef"
)

func TestInitPlan(t *testing.T){
	plan := &Plan{}
	plan.InitPlan()
	plan.SetCname("Plan")
	plan.SetCaption("Plan")
	plan.SetDescription("Plan Test")

	if plan.GetName() != "Plan" {
		t.Error("InitPlan Error,GetName Error!")
	}
	if plan.GetCaption() != "Plan" {
		t.Error("InitPlan Error,GetCaption Error!")
	}
	if plan.GetDescription() != "Plan Test" {
		t.Error("InitPlan Error,GetDescription Error!")
	}
	if plan.GetCtype() != constdef.ComponentType[constdef.Component_Task]+"."+constdef.TaskType[constdef.Task_Plan] {
		t.Error("InitPlan Error,GetCtype Error!")
	}

	var task_1 Enquiry = Enquiry{}
	task_1.InitEnquriy()
	task_1.SetCname("Enquriy")
	task_1.SetCaption("Enquriy")
	task_1.SetDescription("Enquriy Test")
	var task_2 Action = Action{}
	task_2.InitAction()
	task_2.SetCname("Action")
	task_2.SetCaption("Action")
	task_2.SetDescription("Action Test")
	//Test AddTask
	plan.AddTask(task_1)
	plan.AddTask(task_2)
	if len(plan.GetTaskList().(map[string]inf.ITask)) != 2 {
		t.Error("AddTask Error!")
	}
	task_map := plan.GetTaskList().(map[string]inf.ITask)
	if task_map["Action"].GetName() != "Action" {
		t.Error("AddTask Error,check Task Erro!")
	}
	//Test RemoveTask
	plan.RemoveTask("Action")
	if len(plan.GetTaskList().(map[string]inf.ITask)) != 1{
		t.Error("Remove Task Error!")
	}
}
