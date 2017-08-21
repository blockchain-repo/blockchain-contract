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

func CreateGeneralDataObject() *data.GeneralData {
	t_data := new(data.GeneralData)
	t_data.InitGeneralData()
	t_data.SetCname("TestGeneralData")
	t_data.SetCaption("data")
	t_data.SetDescription("test general data")
	t_data.SetUnit("")

	return t_data
}

func CreateGeneralTaskObject() *task.GeneralTask {
	t_task := new(task.GeneralTask)
	t_task.InitGeneralTask()
	t_task.SetTaskId("Task-UUID-0001")
	t_task.SetCname("TestGeneralTask")
	t_task.SetCaption("task")
	t_task.SetDescription("test general task")

	return t_task
}

func CreateGeneralExpressionObject() *expression.GeneralExpression {
	t_expression := expression.NewGeneralExpression("TestGeneralExpression", "Test Geneeral expression")
	t_expression.InitExpression()
	t_expression.SetCaption("expression")
	t_expression.SetDescription("test general expression")

	return t_expression
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
	var inf_expression inf.IExpression = expression.NewGeneralExpression("", "")
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
	var inf_data inf.IData = inf.IData(v_data)
	t_comp_table.AddComponent(inf_data)
	fmt.Println(t_comp_table)
	if v_value, ok := t_comp_table.CompTable[constdef.ComponentType[constdef.Component_Data]]; !ok {
		t.Error("component_table add Data Error!")
	} else if len(v_value) != 1 {
		t.Error("component_table add Data, element length Error")
	} else {
		for t_key, t_value := range v_value[0] {
			tt_value := t_value.(inf.IData)
			if t_key != "TestGeneralData" || tt_value.GetName() != "TestGeneralData" {
				t.Error("component_table add Data, data info Error!")
			}
		}
	}
	var inf_task inf.ITask = inf.ITask(v_task)
	t_comp_table.AddComponent(inf_task)
	fmt.Println(t_comp_table)
	if v_value, ok := t_comp_table.CompTable[constdef.ComponentType[constdef.Component_Task]]; !ok {
		t.Error("component_table add Task Error!")
	} else if len(v_value) != 1 {
		t.Error("component_table add Task, element length Error")
	} else {
		for t_key, t_value := range v_value[0] {
			tt_value := t_value.(inf.ITask)
			if t_key != "TestGeneralTask" || tt_value.GetName() != "TestGeneralTask" {
				t.Error("component_table add Task, task info Error!")
			}
		}
	}
	var inf_expression inf.IExpression = inf.IExpression(v_expression)
	t_comp_table.AddComponent(inf_expression)
	fmt.Println(t_comp_table)
	if v_value, ok := t_comp_table.CompTable[constdef.ComponentType[constdef.Component_Expression]]; !ok {
		t.Error("component_table add Expression Error!")
	} else if len(v_value) != 1 {
		t.Error("component_table add Expression, element length Error")
	} else {
		for t_key, t_value := range v_value[0] {
			tt_value := t_value.(inf.IExpression)
			if t_key != "TestGeneralExpression" || tt_value.GetName() != "TestGeneralExpression" {
				t.Error("component_table add Expression, expression info Error!")
			}
		}
	}
	//test GetComponent
	q_component_1 := t_comp_table.GetComponent("TestGeneralData", constdef.ComponentType[constdef.Component_Data])
	if q_component_1 == nil {
		t.Error("GetComponent Data Error!")
	} else {
		v_comp_1 := q_component_1.(inf.IData)
		if v_comp_1.GetName() != "TestGeneralData" {
			t.Error("GetComponent Data, check name Error!")
		}
	}
	q_component_2 := t_comp_table.GetComponent("TestGeneralTask", constdef.ComponentType[constdef.Component_Task])
	if q_component_2 == nil {
		t.Error("GetComponent Task Error!")
	} else {
		v_comp_2 := q_component_2.(inf.ITask)
		if v_comp_2.GetName() != "TestGeneralTask" {
			t.Error("GetComponent Task, check name Error!")
		}
	}
	q_component_3 := t_comp_table.GetComponent("TestGeneralExpression", constdef.ComponentType[constdef.Component_Expression])
	if q_component_3 == nil {
		t.Error("GetComponent Expression Error!")
	} else {
		v_comp_3 := q_component_3.(inf.IExpression)
		if v_comp_3.GetName() != "TestGeneralExpression" {
			t.Error("GetComponent Expression, check name Error!")
		}
	}
	//test GetComponentByType
	ct_component_1 := t_comp_table.GetComponentByType(constdef.ComponentType[constdef.Component_Data])
	if len(ct_component_1) != 1 {
		t.Error("GetComponentByType Data Error,length error!")
	} else {
		if _, ok := ct_component_1[0]["TestGeneralData"]; !ok {
			t.Error("GetComponentByType Data Error, element not exist!")
		}
	}
	ct_component_2 := t_comp_table.GetComponentByType(constdef.ComponentType[constdef.Component_Task])
	if len(ct_component_2) != 1 {
		t.Error("GetComponentByType Task Error,length error!")
	} else {
		if _, ok := ct_component_2[0]["TestGeneralTask"]; !ok {
			t.Error("GetComponentByType Task Error, element not exist!")
		}
	}
	ct_component_3 := t_comp_table.GetComponentByType(constdef.ComponentType[constdef.Component_Expression])
	if len(ct_component_3) != 1 {
		t.Error("GetComponentByType Expression Error,length error!")
	} else {
		if _, ok := ct_component_3[0]["TestGeneralExpression"]; !ok {
			t.Error("GetComponentByType Expression Error, element not exist!")
		}
	}

	//test GetTaskByID
	id_component_1 := t_comp_table.GetTaskByID("Task-UUID-0001", constdef.ComponentType[constdef.Component_Task])
	if id_component_1 == nil {
		t.Error("GetTaskByID Error!")
	} else {
		id_comp_1 := id_component_1.(inf.ITask)
		if id_comp_1.GetName() != "TestGeneralTask" {
			t.Error("GetTaskByID, check name Error!")
		}
	}
	//test UpdateComponent
	id_component_2 := t_comp_table.GetTaskByID("Task-UUID-0001", constdef.ComponentType[constdef.Component_Task])
	if id_component_2 == nil {
		t.Error("GetTaskByID Error!")
	} else {
		id_component_2.(inf.ITask).SetTaskId("Task-UUID-0002")
		t_comp_table.UpdateComponent(constdef.ComponentType[constdef.Component_Task], id_component_2.(inf.ITask).GetName(), id_component_2)
		id_component_3 := t_comp_table.GetTaskByID("Task-UUID-0002", constdef.ComponentType[constdef.Component_Task])
		id_comp_3 := id_component_3.(inf.ITask)
		if id_comp_3.GetName() != "TestGeneralTask" {
			t.Error("UpdateComponent Error!")
		}
	}
}
