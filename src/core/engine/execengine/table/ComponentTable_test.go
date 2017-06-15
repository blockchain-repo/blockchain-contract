package table

import (
	"fmt"
	"testing"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/data"
	"unicontract/src/core/engine/execengine/expression"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/task"
)

func CreateGeneralDataObject() data.GeneralData {
	t_data := new(data.GeneralData)
	t_data.InitGeneralData()
	t_data.SetCname("TestGeneralData")
	t_data.SetCaption("data")
	t_data.SetDescription("test general data")
	t_data.SetUnit("")

	return *t_data
}

func CreateGeneralTaskObject() task.GeneralTask {
	t_task := new(task.GeneralTask)
	t_task.InitGeneralTask()
	t_task.SetTaskId("Task-UUID-0001")
	t_task.SetCname("TestGeneralTask")
	t_task.SetCaption("task")
	t_task.SetDescription("test general task")

	return *t_task
}

func CreateGeneralExpressionObject() expression.GeneralExpression {
	t_expression := new(expression.GeneralExpression)
	t_expression.InitExpression()
	t_expression.SetCname("TestGeneralExpression")
	t_expression.SetCaption("expression")
	t_expression.SetDescription("test general expression")

	return *t_expression
}

func TestGetComponentType(t *testing.T) {
	v_comp_table := new(ComponentTable)
	//v_type, _ := v_comp_table.getComponentType(v_str)
	//if v_type != constdef.ComponentType[constdef.Component_Unknown] {
	//	t.Error("Type is not Unknow!")
	//}
	var inf_data inf.IData = data.NewGeneralData()
	v_type, _ := v_comp_table.getComponentType(inf_data)
	if v_type != constdef.ComponentType[constdef.Component_Data] {
		t.Error("Type is not Data!")
	}
	var inf_task inf.ITask = task.NewGeneralTask()
	v_type, _ = v_comp_table.getComponentType(inf_task)
	if v_type != constdef.ComponentType[constdef.Component_Task] {
		t.Error("Type is not Task!")
	}
	var inf_expression inf.IExpression = expression.NewGeneralExpression("")
	v_type, _ = v_comp_table.getComponentType(inf_expression)
	if v_type != constdef.ComponentType[constdef.Component_Expression] {
		t.Error("Type is not Expression!")
	}
}

func TestComponentTableAll(t *testing.T) {
	//compTable map[string][]map[string]component.GeneralComponent
	t_comp_table := &ComponentTable{}
	v_data := CreateGeneralDataObject()
	v_task := CreateGeneralTaskObject()
	v_expression := CreateGeneralExpressionObject()
	//test AddComponent
	t_comp_table.AddComponent(v_data)
	fmt.Println(t_comp_table)
	if v_value, ok := t_comp_table.CompTable["1"]; !ok {
		t.Error("component_table add Data Error!")
	} else if len(v_value) != 1 {
		t.Error("component_table add Data, element length Error")
	} else {
		for t_key, t_value := range v_value[0] {
			tt_value := t_value.(data.GeneralData)
			if t_key != "TestGeneralData" || tt_value.GetCname() != "TestGeneralData" {
				t.Error("component_table add Data, data info Error!")
			}
		}
	}
	t_comp_table.AddComponent(v_task)
	fmt.Println(t_comp_table)
	if v_value, ok := t_comp_table.CompTable["2"]; !ok {
		t.Error("component_table add Task Error!")
	} else if len(v_value) != 1 {
		t.Error("component_table add Task, element length Error")
	} else {
		for t_key, t_value := range v_value[0] {
			tt_value := t_value.(task.GeneralTask)
			if t_key != "TestGeneralTask" || tt_value.GetCname() != "TestGeneralTask" {
				t.Error("component_table add Task, task info Error!")
			}
		}
	}
	t_comp_table.AddComponent(v_expression)
	fmt.Println(t_comp_table)
	if v_value, ok := t_comp_table.CompTable["3"]; !ok {
		t.Error("component_table add Expression Error!")
	} else if len(v_value) != 1 {
		t.Error("component_table add Expression, element length Error")
	} else {
		for t_key, t_value := range v_value[0] {
			tt_value := t_value.(expression.GeneralExpression)
			if t_key != "TestGeneralExpression" || tt_value.GetCname() != "TestGeneralExpression" {
				t.Error("component_table add Expression, expression info Error!")
			}
		}
	}
	//test GetComponent
	q_component_1 := t_comp_table.GetComponent("TestGeneralData", "1")
	if q_component_1 == nil {
		t.Error("GetComponent Data Error!")
	} else {
		v_comp_1 := q_component_1.(data.GeneralData)
		if v_comp_1.GetCname() != "TestGeneralData" {
			t.Error("GetComponent Data, check name Error!")
		}
	}
	q_component_2 := t_comp_table.GetComponent("TestGeneralTask", "2")
	if q_component_2 == nil {
		t.Error("GetComponent Task Error!")
	} else {
		v_comp_2 := q_component_2.(task.GeneralTask)
		if v_comp_2.GetCname() != "TestGeneralTask" {
			t.Error("GetComponent Task, check name Error!")
		}
	}
	q_component_3 := t_comp_table.GetComponent("TestGeneralExpression", "3")
	if q_component_3 == nil {
		t.Error("GetComponent Expression Error!")
	} else {
		v_comp_3 := q_component_3.(expression.GeneralExpression)
		if v_comp_3.GetCname() != "TestGeneralExpression" {
			t.Error("GetComponent Expression, check name Error!")
		}
	}
	//test GetComponentByType
	ct_component_1 := t_comp_table.GetComponentByType("1")
	if len(ct_component_1) != 1 {
		t.Error("GetComponentByType Data Error,length error!")
	} else {
		if _, ok := ct_component_1[0]["TestGeneralData"]; !ok {
			t.Error("GetComponentByType Data Error, element not exist!")
		}
	}
	ct_component_2 := t_comp_table.GetComponentByType("2")
	if len(ct_component_2) != 1 {
		t.Error("GetComponentByType Task Error,length error!")
	} else {
		if _, ok := ct_component_2[0]["TestGeneralTask"]; !ok {
			t.Error("GetComponentByType Task Error, element not exist!")
		}
	}
	ct_component_3 := t_comp_table.GetComponentByType("3")
	if len(ct_component_3) != 1 {
		t.Error("GetComponentByType Expression Error,length error!")
	} else {
		if _, ok := ct_component_3[0]["TestGeneralExpression"]; !ok {
			t.Error("GetComponentByType Expression Error, element not exist!")
		}
	}

	//test GetTaskByID
	id_component_1 := t_comp_table.GetTaskByID("Task-UUID-0001", "2")
	if id_component_1 == nil {
		t.Error("GetTaskByID Error!")
	} else {
		id_comp_1 := id_component_1.(task.GeneralTask)
		if id_comp_1.GetCname() != "TestGeneralTask" {
			t.Error("GetTaskByID, check name Error!")
		}
		t.Error("GetTaskByID, check name Error!")
	}
}
