package execengine

import (
	"encoding/json"
	"errors"
	"fmt"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/contract"
	"unicontract/src/core/engine/execengine/data"
	"unicontract/src/core/engine/execengine/expression"
	"unicontract/src/core/engine/execengine/task"
)

func loadTask(p_contract *contract.CognitiveContract, p_component interface{}) error {
	var err error = nil
	if p_component == nil {
		err = errors.New("Param[component] is null!")
		return err
	}
	map_task := p_component.(map[string]interface{})
	switch map_task["Ctype"] {
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Enquiry]:
		//1.反序列化
		p_task := task.NewEnquiry()
		byte_task, _ := json.Marshal(map_task)
		err = json.Unmarshal(byte_task, &p_task)
		fmt.Println("333333333: ", p_task)
		//2 初始化
		p_task.InitEnquriy()
		fmt.Println("444444444: ", p_task)
		//3 处理数组属性
		if err := loadTaskInnerComponent(p_contract, map_task, p_task); err != nil {
			fmt.Println("err in loadTaskInnerComponent[Enquiry]")
			return err
		}
		//4 添加任务组件到component_table中
		p_contract.AddComponent(p_task)
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Action]:
		p_task := task.NewAction()
		byte_task, _ := json.Marshal(map_task)
		err = json.Unmarshal(byte_task, &p_task)
		p_task.InitAction()
		if err := loadTaskInnerComponent(p_contract, map_task, p_task); err != nil {
			fmt.Println("err in loadTaskInnerComponent[Action]")
			return err
		}
		p_contract.AddComponent(p_task)
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Decision]:
		p_task := task.NewDecision()
		byte_task, _ := json.Marshal(map_task)
		err = json.Unmarshal(byte_task, &p_task)
		p_task.InitDecision()
		if err := loadTaskInnerComponent(p_contract, map_task, p_task); err != nil {
			fmt.Println("err in loadTaskInnerComponent[Decision]")
			return err
		}
		p_contract.AddComponent(p_task)
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Plan]:
		p_task := task.NewPlan()
		byte_task, _ := json.Marshal(map_task)
		err = json.Unmarshal(byte_task, &p_task)
		p_task.InitPlan()
		if err := loadTaskInnerComponent(p_contract, map_task, p_task); err != nil {
			fmt.Println("err in loadTaskInnerComponent[Plan]")
			return err
		}
		p_contract.AddComponent(p_task)
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Unknown]:
		p_task := task.NewGeneralTask()
		byte_task, _ := json.Marshal(map_task)
		err = json.Unmarshal(byte_task, &p_task)
		p_task.InitGeneralTask()
		if err := loadTaskInnerComponent(p_contract, map_task, p_task); err != nil {
			fmt.Println("err in loadTaskInnerComponent[Unknow]")
			return err
		}
		p_contract.AddComponent(p_task)
	}
	return err
}

func loadTaskInnerComponent(p_contract *contract.CognitiveContract, m_task interface{}, p_task interface{}) error {
	var err error = nil
	if m_task == nil || p_task == nil {
		err = errors.New("Param[map_task or object_task] is null!")
		return err
	}
	map_task := m_task.(map[string]interface{})
	switch map_task["Ctype"] {
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Enquiry]:

		pre_conditions := map_task["PreCondition"]
		for _, p_value := range pre_conditions.([]interface{}) {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadExpression[Enquiry.PreCondition]")
				return err
			}
		}
		comp_conditions := map_task["CompleteCondition"]
		for _, p_value := range comp_conditions.([]interface{}) {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadExpression[Enquiry.CompleteCondition]")
				return err
			}
		}
		digard_conditions := map_task["DisgardCondition"]
		for _, p_value := range digard_conditions.([]interface{}) {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadExpression[Enquiry.DisgardCondition]")
				return err
			}
		}
		data_list := map_task["DataList"]
		for _, p_value := range data_list.([]interface{}) {
			if err := loadData(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadData[Enquiry.DataList]")
				return err
			}
		}
		dataexpress_list := map_task["DataValueSetterExpressionList"]
		for _, p_value := range dataexpress_list.([]interface{}) {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadExpression[Enquiry.DataValueSetterExpressionList]")
				return err
			}
		}
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Action]:
		pre_conditions := map_task["PreCondition"]
		for _, p_value := range pre_conditions.([]interface{}) {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadExpression[Action.PreCondition]")
				return err
			}
		}
		comp_conditions := map_task["CompleteCondition"]
		for _, p_value := range comp_conditions.([]interface{}) {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadExpression[Action.CompleteCondition]")
				return err
			}
		}
		digard_conditions := map_task["DisgardCondition"]
		for _, p_value := range digard_conditions.([]interface{}) {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadExpression[Action.DisgardCondition]")
				return err
			}
		}
		data_list := map_task["DataList"]
		for _, p_value := range data_list.([]interface{}) {
			if err := loadData(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadData[Action.DataList]")
				return err
			}
		}
		dataexpress_list := map_task["DataValueSetterExpressionList"]
		for _, p_value := range dataexpress_list.([]interface{}) {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadExpression[Action.DataValueSetterExpressionList]")
				return err
			}
		}
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Decision]:
		pre_conditions := map_task["PreCondition"]
		for _, p_value := range pre_conditions.([]interface{}) {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadExpression[Decision.PreCondition]")
				return err
			}
		}
		comp_conditions := map_task["CompleteCondition"]
		for _, p_value := range comp_conditions.([]interface{}) {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadExpression[Decision.CompleteCondition]")
				return err
			}
		}
		digard_conditions := map_task["DisgardCondition"]
		for _, p_value := range digard_conditions.([]interface{}) {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadExpression[Decision.DisgardCondition]")
				return err
			}
		}
		data_list := map_task["DataList"]
		for _, p_value := range data_list.([]interface{}) {
			if err := loadData(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadData[Decision.DataList]")
				return err
			}
		}
		dataexpress_list := map_task["DataValueSetterExpressionList"]
		for _, p_value := range dataexpress_list.([]interface{}) {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadExpression[Decision.DataValueSetterExpressionList]")
				return err
			}
		}
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Plan]:
		pre_conditions := map_task["PreCondition"]
		for _, p_value := range pre_conditions.([]interface{}) {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadExpression[Plan.PreCondition]")
				return err
			}
		}
		comp_conditions := map_task["CompleteCondition"]
		for _, p_value := range comp_conditions.([]interface{}) {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadExpression[Plan.CompleteCondition]")
				return err
			}
		}
		digard_conditions := map_task["DisgardCondition"]
		for _, p_value := range digard_conditions.([]interface{}) {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadExpression[Plan.DisgardCondition]")
				return err
			}
		}
	default:
		pre_conditions := map_task["PreCondition"]
		for _, p_value := range pre_conditions.([]interface{}) {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadExpression[Unknow.PreCondition]")
				return err
			}
		}
		comp_conditions := map_task["CompleteCondition"]
		for _, p_value := range comp_conditions.([]interface{}) {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadExpression[Unknow.CompleteCondition]")
				return err
			}
		}
		digard_conditions := map_task["DisgardCondition"]
		for _, p_value := range digard_conditions.([]interface{}) {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				fmt.Println("err in loadExpression[Unknow.DisgardCondition]")
				return err
			}
		}
	}
	return err
}

func loadData(p_contract *contract.CognitiveContract, m_data interface{}, p_task interface{}) error {
	var err error = nil
	if p_contract == nil || m_data == nil || p_task == nil {
		err = errors.New("Param[p_contract || m_data || p_task] is null!")
		return err
	}
	map_data := m_data.(map[string]interface{})
	switch map_data["Ctype"] {
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Numeric_Int]:
		//1.反序列化
		p_data := data.NewIntData()
		byte_data, _ := json.Marshal(map_data)
		err = json.Unmarshal(byte_data, &p_data)
		fmt.Println("44444444: ", p_data)
		//2 初始化
		p_data.InitIntData()
		fmt.Println("55555555: ", p_data)
		//3 添加data组件到component_table中
		p_contract.AddComponent(p_data)
		//4 添加data组件到task中
		//v_task := p_task.(inf.ITask)
		//v_task.AddData()
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Numeric_Uint]:
		//1.反序列化
		p_data := data.NewUintData()
		byte_data, _ := json.Marshal(map_data)
		err = json.Unmarshal(byte_data, &p_data)
		fmt.Println("44444444: ", p_data)
		//2 初始化
		p_data.InitUintData()
		fmt.Println("55555555: ", p_data)
		//3 添加data组件到component_table中
		p_contract.AddComponent(p_data)
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Numeric_Float]:
		//1.反序列化
		p_data := data.NewFloatData()
		byte_data, _ := json.Marshal(map_data)
		err = json.Unmarshal(byte_data, &p_data)
		fmt.Println("44444444: ", p_data)
		//2 初始化
		p_data.InitFloatData()
		fmt.Println("55555555: ", p_data)
		//3 添加data组件到component_table中
		p_contract.AddComponent(p_data)
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Text]:
		//1.反序列化
		p_data := data.NewTextData()
		byte_data, _ := json.Marshal(map_data)
		err = json.Unmarshal(byte_data, &p_data)
		fmt.Println("44444444: ", p_data)
		//2 初始化
		p_data.InitTextData()
		fmt.Println("55555555: ", p_data)
		//3 添加data组件到component_table中
		p_contract.AddComponent(p_data)
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Date]:
		//1.反序列化
		p_data := data.NewDateData()
		byte_data, _ := json.Marshal(map_data)
		err = json.Unmarshal(byte_data, &p_data)
		fmt.Println("44444444: ", p_data)
		//2 初始化
		p_data.InitDateData()
		fmt.Println("55555555: ", p_data)
		//3 添加data组件到component_table中
		p_contract.AddComponent(p_data)
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Array]:
		//1.反序列化
		p_data := data.NewArrayData()
		byte_data, _ := json.Marshal(map_data)
		err = json.Unmarshal(byte_data, &p_data)
		fmt.Println("44444444: ", p_data)
		//2 初始化
		p_data.InitArrayData()
		fmt.Println("55555555: ", p_data)
		//3 添加data组件到component_table中
		p_contract.AddComponent(p_data)
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Matrix]:
		//1.反序列化
		p_data := data.NewMatrixData()
		byte_data, _ := json.Marshal(map_data)
		err = json.Unmarshal(byte_data, &p_data)
		fmt.Println("44444444: ", p_data)
		//2 初始化
		p_data.InitMatrixData()
		fmt.Println("55555555: ", p_data)
		//3 添加data组件到component_table中
		p_contract.AddComponent(p_data)
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Compound]:
		//1.反序列化
		p_data := data.NewCompoundData()
		byte_data, _ := json.Marshal(map_data)
		err = json.Unmarshal(byte_data, &p_data)
		fmt.Println("44444444: ", p_data)
		//2 初始化
		p_data.InitCompoundData()
		fmt.Println("55555555: ", p_data)
		//3 添加data组件到component_table中
		p_contract.AddComponent(p_data)
	default:
		//1.反序列化
		p_data := data.NewGeneralData()
		byte_data, _ := json.Marshal(map_data)
		err = json.Unmarshal(byte_data, &p_data)
		fmt.Println("44444444: ", p_data)
		//2 初始化
		p_data.InitGeneralData()
		fmt.Println("55555555: ", p_data)
		//3 添加data组件到component_table中
		p_contract.AddComponent(p_data)
	}
	return err
}

func loadExpression(p_contract *contract.CognitiveContract, p_expression interface{}, p_task interface{}) error {
	var err error = nil
	if p_contract == nil || p_task == nil || p_expression == nil {
		err = errors.New("Param[p_contract or p_task or expression] is null!")
		return err
	}
	map_expression := p_expression.(map[string]interface{})
	switch map_expression["Ctype"] {
	case constdef.ComponentType[constdef.Component_Expression] + "." + constdef.ExpressionType[constdef.Expression_Condition]:
		object_expression := &expression.LogicArgument{}

		byte_task, _ := json.Marshal(map_expression)
		err = json.Unmarshal(byte_task, &object_expression)
		fmt.Println("777777777: ", object_expression)
		object_expression.InitLogicArgument()
		p_contract.AddComponent(object_expression)
	case constdef.ComponentType[constdef.Component_Expression] + "." + constdef.ExpressionType[constdef.Expression_Function]:
		object_expression := &expression.Function{}
		byte_task, _ := json.Marshal(map_expression)
		err = json.Unmarshal(byte_task, &object_expression)
		fmt.Println("777777777: ", object_expression)
		object_expression.InitFunction()
		p_contract.AddComponent(object_expression)
	default:
		object_expression := &expression.GeneralExpression{}
		byte_task, _ := json.Marshal(map_expression)
		err = json.Unmarshal(byte_task, &object_expression)
		fmt.Println("777777777: ", object_expression)
		object_expression.InitExpression()
		p_contract.AddComponent(object_expression)
	}
	return err
}
