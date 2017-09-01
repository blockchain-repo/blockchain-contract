package execengine

/*
	合约组件加载工具类
	1.Task组件加载:Task_Enquiry, Task_Action, Task_Decision, Task_Plan
	2.Data组件加载:Data_Int, Data_Uint, Data_Float, Data_Date, Data_Text
	               Data_Bool, Data_OperateResult 新增组件
	              (Data_Array, Data_Matrix, Data_Compound) 不常用组件
	3.Expression组件加载：Expression_Condition, Expression_Function
*/
import (
	"bytes"
	"encoding/json"
	"fmt"

	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/contract"
	"unicontract/src/core/engine/execengine/data"
	"unicontract/src/core/engine/execengine/expression"
	"unicontract/src/core/engine/execengine/task"
)

func loadTask(p_contract *contract.CognitiveContract, p_component interface{}) error {
	var err error
	var r_buf bytes.Buffer
	r_buf.WriteString("Contract LoadTask;")
	if p_component == nil {
		r_buf.WriteString("[Result]:LoadTask fail;")
		r_buf.WriteString("[Error]:Param Error!")
		uniledgerlog.Warn(r_buf.String())
		return fmt.Errorf("Param Error!")
	}
	map_task, ok := p_component.(map[string]interface{})
	if !ok {
		r_buf.WriteString("[Result]:LoadTask Fail;")
		r_buf.WriteString("[Error]:Assert Error!")
		uniledgerlog.Error(r_buf.String())
		return fmt.Errorf("Assert Error!")
	}
	switch map_task["Ctype"] {
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Enquiry]:
		//1.反序列化
		p_task := task.NewEnquiry()
		byte_task, _ := json.Marshal(map_task)
		err = json.Unmarshal(byte_task, &p_task)

		//2.初始化
		p_task.InitEnquriy()

		//3.处理数组属性
		if err := loadTaskInnerComponent(p_contract, map_task, p_task); err != nil {
			r_buf.WriteString("[Result]:LoadTask(Enquiry) fail;")
			r_buf.WriteString("[Error]:Error in LoadTask(loadTaskInnerComponent) ," + err.Error())
			uniledgerlog.Warn(r_buf.String())
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
			r_buf.WriteString("[Result]:LoadTask(Action) fail;")
			r_buf.WriteString("[Error]:Error in LoadTask(loadTaskInnerComponent) ," + err.Error())
			uniledgerlog.Warn(r_buf.String())
			return err
		}
		p_contract.AddComponent(p_task)
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Decision]:
		p_task := task.NewDecision()
		byte_task, _ := json.Marshal(map_task)
		err = json.Unmarshal(byte_task, &p_task)
		p_task.InitDecision()
		if err := loadTaskInnerComponent(p_contract, map_task, p_task); err != nil {
			r_buf.WriteString("[Result]:LoadTask(Decision) fail;")
			r_buf.WriteString("[Error]:Error in LoadTask(loadTaskInnerComponent) ," + err.Error())
			uniledgerlog.Warn(r_buf.String())
			return err
		}
		p_contract.AddComponent(p_task)
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Plan]:
		p_task := task.NewPlan()
		byte_task, _ := json.Marshal(map_task)
		err = json.Unmarshal(byte_task, &p_task)
		p_task.InitPlan()
		if err := loadTaskInnerComponent(p_contract, map_task, p_task); err != nil {
			r_buf.WriteString("[Result]:LoadTask(Plan) fail;")
			r_buf.WriteString("[Error]:Error in LoadTask(loadTaskInnerComponent) ," + err.Error())
			uniledgerlog.Warn(r_buf.String())
			return err
		}
		p_contract.AddComponent(p_task)
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Unknown]:
		p_task := task.NewGeneralTask()
		byte_task, _ := json.Marshal(map_task)
		err = json.Unmarshal(byte_task, &p_task)
		p_task.InitGeneralTask()
		if err := loadTaskInnerComponent(p_contract, map_task, p_task); err != nil {
			r_buf.WriteString("[Result]:LoadTask fail;")
			r_buf.WriteString("[Error]:Error in LoadTask(loadTaskInnerComponent) ," + err.Error())
			uniledgerlog.Warn(r_buf.String())
			return err
		}
		p_contract.AddComponent(p_task)
	}
	r_buf.WriteString("[Cname]: " + map_task["Cname"].(string) + "[Ctype]: " + map_task["Ctype"].(string) + "[Result]: LoadTask success;")
	uniledgerlog.Info(r_buf.String())
	return err
}

func loadTaskInnerComponent(p_contract *contract.CognitiveContract, m_task interface{}, p_task interface{}) error {
	var err error
	var r_buf bytes.Buffer
	r_buf.WriteString("loadTaskInnerComponent;")
	if m_task == nil || p_task == nil {
		r_buf.WriteString("[Result]:loadTaskInnerComponent fail;")
		r_buf.WriteString("[Error]:Param Error!")
		uniledgerlog.Warn(r_buf.String())
		return fmt.Errorf("Param Error!")
	}
	map_task, ok := m_task.(map[string]interface{})
	if !ok {
		r_buf.WriteString("[Result]:loadTaskInnerComponent Fail;")
		r_buf.WriteString("[Error]:Assert Error!")
		uniledgerlog.Error(r_buf.String())
		return fmt.Errorf("Assert Error!")
	}
	switch map_task["Ctype"] {
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Enquiry]:
		pre_conditions := map_task["PreCondition"]
		if pre_conditions != nil {
			sl1, ok := pre_conditions.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Enquiry) Fail;")
				r_buf.WriteString("[Error]:Enquiry.PreCondition Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Enquiry.PreCondition Assert Error!")
			}
			for _, p_value := range sl1 {
				if err := loadExpression(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Enquiry) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadExpression[Enquiry.PreCondition]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
		comp_conditions := map_task["CompleteCondition"]
		if comp_conditions != nil {
			sl2, ok := comp_conditions.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Enquiry) Fail;")
				r_buf.WriteString("[Error]:Enquiry.CompleteCondition Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Enquiry.CompleteCondition Assert Error!")
			}
			for _, p_value := range sl2 {
				if err := loadExpression(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Enquiry) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadExpression[Enquiry.CompleteCondition]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
		digard_conditions := map_task["DiscardCondition"]
		if digard_conditions != nil {
			sl3, ok := digard_conditions.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Enquiry) Fail;")
				r_buf.WriteString("[Error]:Enquiry.DiscardCondition Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Enquiry.DiscardCondition Assert Error!")
			}
			for _, p_value := range sl3 {
				if err := loadExpression(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Enquiry) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadExpression[Enquiry.DiscardCondition]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
		data_list := map_task["DataList"]
		if data_list != nil {
			sl4, ok := data_list.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Enquiry) Fail;")
				r_buf.WriteString("[Error]:Enquiry.DataList Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Enquiry.DataList Assert Error!")
			}
			for _, p_value := range sl4 {
				if err := loadData(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Enquiry) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadData[Enquiry.DataList]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
		dataexpress_list := map_task["DataValueSetterExpressionList"]
		if dataexpress_list != nil {
			sl5, ok := dataexpress_list.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Enquiry) Fail;")
				r_buf.WriteString("[Error]:Enquiry.DataValueSetterExpressionList Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Enquiry.DataValueSetterExpressionList Assert Error!")
			}
			for _, p_value := range sl5 {
				if err := loadExpression(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Enquiry) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadExpression[Enquiry.DataValueSetterExpressionList]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Action]:
		pre_conditions := map_task["PreCondition"]
		if pre_conditions != nil {
			sl6, ok := pre_conditions.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Action) Fail;")
				r_buf.WriteString("[Error]:Action.PreCondition Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Action.PreCondition Assert Error!")
			}
			for _, p_value := range sl6 {
				if err := loadExpression(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Action) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadExpression[Action.PreCondition]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
		comp_conditions := map_task["CompleteCondition"]
		if comp_conditions != nil {
			sl7, ok := comp_conditions.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Action) Fail;")
				r_buf.WriteString("[Error]:Action.CompleteCondition Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Action.CompleteCondition Assert Error!")
			}
			for _, p_value := range sl7 {
				if err := loadExpression(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Action) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadExpression[Action.CompleteCondition]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
		digard_conditions := map_task["DiscardCondition"]
		if digard_conditions != nil {
			sl8, ok := digard_conditions.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Action) Fail;")
				r_buf.WriteString("[Error]:Action.DiscardCondition Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Action.DiscardCondition Assert Error!")
			}
			for _, p_value := range sl8 {
				if err := loadExpression(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Action) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadExpression[Action.DiscardCondition]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
		data_list := map_task["DataList"]
		if data_list != nil {
			sl9, ok := data_list.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Action) Fail;")
				r_buf.WriteString("[Error]:Action.DataList Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Action.DataList Assert Error!")
			}
			for _, p_value := range sl9 {
				if err := loadData(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Action) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadData[Action.DataList]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
		dataexpress_list := map_task["DataValueSetterExpressionList"]
		if dataexpress_list != nil {
			sl10, ok := dataexpress_list.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Action) Fail;")
				r_buf.WriteString("[Error]:Action.DataValueSetterExpressionList Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Action.DataValueSetterExpressionList Assert Error!")
			}
			for _, p_value := range sl10 {
				if err := loadExpression(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Action) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadExpression[Action.DataValueSetterExpressionList]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Decision]:
		pre_conditions := map_task["PreCondition"]
		if pre_conditions != nil {
			sl11, ok := pre_conditions.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Decision) Fail;")
				r_buf.WriteString("[Error]:Decision.PreCondition Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Decision.PreCondition Assert Error!")
			}
			for _, p_value := range sl11 {
				if err := loadExpression(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Decision) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadExpression[Decision.PreCondition]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
		comp_conditions := map_task["CompleteCondition"]
		if comp_conditions != nil {
			sl12, ok := comp_conditions.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Decision) Fail;")
				r_buf.WriteString("[Error]:Decision.CompleteCondition Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Decision.CompleteCondition Assert Error!")
			}
			for _, p_value := range sl12 {
				if err := loadExpression(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Decision) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadExpression[Decision.CompleteCondition]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
		digard_conditions := map_task["DiscardCondition"]
		if digard_conditions != nil {
			sl13, ok := digard_conditions.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Decision) Fail;")
				r_buf.WriteString("[Error]:Decision.DiscardCondition Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Decision.DiscardCondition Assert Error!")
			}
			for _, p_value := range sl13 {
				if err := loadExpression(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Decision) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadExpression[Decision.DiscardCondition]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
		candidate_list := map_task["CandidateList"]
		if candidate_list != nil {
			sl14, ok := candidate_list.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Decision) Fail;")
				r_buf.WriteString("[Error]:Decision.CandidateList Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Decision.CandidateList Assert Error!")
			}
			for _, p_value := range sl14 {
				if err := loadCandidate(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Decision) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadCandidate[Decision.CandidateList]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Plan]:
		pre_conditions := map_task["PreCondition"]
		if pre_conditions != nil {
			sl16, ok := pre_conditions.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Plan) Fail;")
				r_buf.WriteString("[Error]:Plan.PreCondition Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Plan.PreCondition Assert Error!")
			}
			for _, p_value := range sl16 {
				if err := loadExpression(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Plan) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadExpression[Plan.PreCondition]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
		comp_conditions := map_task["CompleteCondition"]
		if comp_conditions != nil {
			sl100, ok := comp_conditions.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Plan) Fail;")
				r_buf.WriteString("[Error]:Plan.CompleteCondition Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Plan.CompleteCondition Assert Error!")
			}
			for _, p_value := range sl100 {
				if err := loadExpression(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Plan) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadExpression[Plan.CompleteCondition]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
		digard_conditions := map_task["DiscardCondition"]
		if digard_conditions != nil {
			sl17, ok := digard_conditions.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Plan) Fail;")
				r_buf.WriteString("[Error]:Plan.DiscardCondition Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Plan.DiscardCondition Assert Error!")
			}
			for _, p_value := range sl17 {
				if err := loadExpression(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Plan) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadExpression[Plan.DiscardCondition]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
	default:
		pre_conditions := map_task["PreCondition"]
		if pre_conditions != nil {
			sl18, ok := pre_conditions.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Unknow) Fail;")
				r_buf.WriteString("[Error]:Task.PreCondition Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Task.PreCondition Assert Error!")
			}
			for _, p_value := range sl18 {
				if err := loadExpression(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Unknow) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadExpression[Unknow.PreCondition]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
		comp_conditions := map_task["CompleteCondition"]
		if comp_conditions != nil {
			sl19, ok := comp_conditions.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Unknow) Fail;")
				r_buf.WriteString("[Error]:Task.CompleteCondition Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Task.CompleteCondition Assert Error!")
			}
			for _, p_value := range sl19 {
				if err := loadExpression(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Unknow) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadExpression[Unknow.CompleteCondition]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
		digard_conditions := map_task["DiscardCondition"]
		if digard_conditions != nil {
			sl20, ok := digard_conditions.([]interface{})
			if !ok {
				r_buf.WriteString("[Result]:loadTaskInnerComponent(Unknow) Fail;")
				r_buf.WriteString("[Error]:Task.DiscardCondition Assert Error!")
				uniledgerlog.Error(r_buf.String())
				return fmt.Errorf("Task.DiscardCondition Assert Error!")
			}
			for _, p_value := range sl20 {
				if err := loadExpression(p_contract, p_value, p_task); err != nil {
					r_buf.WriteString("[Result]:loadTaskInnerComponent(Unknow) Fail;")
					r_buf.WriteString("[Error]:Error in loadTaskInnerComponent(loadExpression[Unknow.DiscardCondition]) ," + err.Error())
					uniledgerlog.Warn(r_buf.String())
					return err
				}
			}
		}
	}
	r_buf.WriteString("[Cname]: " + map_task["Cname"].(string) + "[Ctype]: " + map_task["Ctype"].(string) + "[Result]: loadTaskInnerComponent success;")
	uniledgerlog.Info(r_buf.String())
	return err
}

func loadData(p_contract *contract.CognitiveContract, m_data interface{}, p_task interface{}) error {
	var err error
	var r_buf bytes.Buffer
	r_buf.WriteString("loadData;")
	if p_contract == nil || m_data == nil || p_task == nil {
		r_buf.WriteString("[Result]:loadData fail;")
		r_buf.WriteString("[Error]:Param Error!")
		uniledgerlog.Warn(r_buf.String())
		return fmt.Errorf("Param Error!")
	}
	map_data, ok := m_data.(map[string]interface{})
	if !ok {
		r_buf.WriteString("[Result]:loadData Fail;")
		r_buf.WriteString("[Error]:Assert Error!")
		uniledgerlog.Error(r_buf.String())
		return fmt.Errorf("Assert Error!")
	}
	switch map_data["Ctype"] {
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Numeric_Int]:
		//1.反序列化
		p_data := data.NewIntData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Int) Fail;")
			r_buf.WriteString("[Error]:Data.Int Serialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Int) Fail;")
			r_buf.WriteString("[Error]:Data.Int Deserialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//2.初始化
		err = p_data.InitIntData()
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Int) Fail;")
			r_buf.WriteString("[Error]:Error in loadData(InitIntData) ," + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//3.添加data组件到component_table中
		p_contract.AddComponent(p_data)

		//4.添加data组件到task中
		//v_task := p_task.(inf.ITask)
		//v_task.AddData()
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Numeric_Uint]:
		//1.反序列化
		p_data := data.NewUintData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Uint) Fail;")
			r_buf.WriteString("[Error]:Data.Uint Serialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Uint) Fail;")
			r_buf.WriteString("[Error]:Data.Uint Deserialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//2.初始化
		err = p_data.InitUintData()
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Uint) Fail;")
			r_buf.WriteString("[Error]:Error in loadData(InitUintData) ," + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//3.添加data组件到component_table中
		p_contract.AddComponent(p_data)
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Numeric_Float]:
		//1.反序列化
		p_data := data.NewFloatData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Float) Fail;")
			r_buf.WriteString("[Error]:Data.Float Serialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Float) Fail;")
			r_buf.WriteString("[Error]:Data.Float Deserialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//2.初始化
		err = p_data.InitFloatData()
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Float) Fail;")
			r_buf.WriteString("[Error]:Error in loadData(InitFloatData) ," + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//3.添加data组件到component_table中
		p_contract.AddComponent(p_data)
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Text]:
		//1.反序列化
		p_data := data.NewTextData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Text) Fail;")
			r_buf.WriteString("[Error]:Data.Text Serialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Text) Fail;")
			r_buf.WriteString("[Error]:Data.Text Deserialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//2.初始化
		err = p_data.InitTextData()
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Text) Fail;")
			r_buf.WriteString("[Error]:Error in loadData(InitTextData) ," + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//3.添加data组件到component_table中
		p_contract.AddComponent(p_data)
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Date]:
		//1.反序列化
		p_data := data.NewDateData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Date) Fail;")
			r_buf.WriteString("[Error]:Data.Date Serialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Date) Fail;")
			r_buf.WriteString("[Error]:Data.Date Deserialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//2.初始化
		err = p_data.InitDateData()
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Date) Fail;")
			r_buf.WriteString("[Error]:Error in loadData(InitDateData) ," + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//3.添加data组件到component_table中
		p_contract.AddComponent(p_data)
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Bool]:
		//1.反序列化
		p_data := data.NewBoolData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Bool) Fail;")
			r_buf.WriteString("[Error]:Data.Bool Serialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Bool) Fail;")
			r_buf.WriteString("[Error]:Data.Bool Deserialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//2.初始化
		err = p_data.InitBoolData()
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Bool) Fail;")
			r_buf.WriteString("[Error]:Error in loadData(InitBoolData) ," + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//3.添加data组件到component_table中
		p_contract.AddComponent(p_data)

		//4.添加data组件到task中
		//v_task := p_task.(inf.ITask)
		//v_task.AddData()
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_OperateResult]:
		//1.反序列化
		p_data := data.NewOperateResultData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(OperateResult) Fail;")
			r_buf.WriteString("[Error]:Data.OperateResult Serialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(OperateResult) Fail;")
			r_buf.WriteString("[Error]:Data.OperateResult Deserialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//2.初始化
		err = p_data.InitOperateResultData()
		if err != nil {
			r_buf.WriteString("[Result]:loadData(OperateResult) Fail;")
			r_buf.WriteString("[Error]:Error in loadData(InitOperateResultData) ," + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//3.添加data组件到component_table中
		p_contract.AddComponent(p_data)

		//4.添加data组件到task中
		//v_task := p_task.(inf.ITask)
		//v_task.AddData()
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Array]:
		//1.反序列化
		p_data := data.NewArrayData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Array) Fail;")
			r_buf.WriteString("[Error]:Data.Array Serialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Array) Fail;")
			r_buf.WriteString("[Error]:Data.Array Deserialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//2.初始化
		err = p_data.InitArrayData()
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Array) Fail;")
			r_buf.WriteString("[Error]:Error in loadData(InitArrayData) ," + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//3.添加data组件到component_table中
		p_contract.AddComponent(p_data)
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Matrix]:
		//1.反序列化
		p_data := data.NewMatrixData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Matrix) Fail;")
			r_buf.WriteString("[Error]:Data.Matrix Serialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Matrix) Fail;")
			r_buf.WriteString("[Error]:Data.Matrix Deserialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//2.初始化
		err = p_data.InitMatrixData()
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Matrix) Fail;")
			r_buf.WriteString("[Error]:Error in loadData(InitMatrixData) ," + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//3.添加data组件到component_table中
		p_contract.AddComponent(p_data)
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Compound]:
		//1.反序列化
		p_data := data.NewCompoundData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Compound) Fail;")
			r_buf.WriteString("[Error]:Data.Compound Serialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Compound) Fail;")
			r_buf.WriteString("[Error]:Data.Compound Deserialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//2.初始化
		err = p_data.InitCompoundData()
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Compound) Fail;")
			r_buf.WriteString("[Error]:Error in loadData(InitCompoundData) ," + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//3.添加data组件到component_table中
		p_contract.AddComponent(p_data)
	default:
		//1.反序列化
		p_data := data.NewGeneralData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Unknow) Fail;")
			r_buf.WriteString("[Error]:Data.Unknow Serialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Unknow) Fail;")
			r_buf.WriteString("[Error]:Data.Unknow Deserialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//2.初始化
		err = p_data.InitGeneralData()
		if err != nil {
			r_buf.WriteString("[Result]:loadData(Unknow) Fail;")
			r_buf.WriteString("[Error]:Error in loadData(InitGeneralData) ," + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		//3.添加data组件到component_table中
		p_contract.AddComponent(p_data)
	}
	r_buf.WriteString("[Cname]: " + map_data["Cname"].(string) + "[Ctype]: " + map_data["Ctype"].(string) + "[Result]: loadData success;")
	//uniledgerlog.Info(r_buf.String())
	return err
}

func loadExpression(p_contract *contract.CognitiveContract, p_expression interface{}, p_task interface{}) error {
	var err error
	var r_buf bytes.Buffer
	r_buf.WriteString("loadExpression;")
	if p_contract == nil || p_task == nil || p_expression == nil {
		r_buf.WriteString("[Result]:loadExpression fail;")
		r_buf.WriteString("[Error]:Param Error!")
		uniledgerlog.Warn(r_buf.String())
		return fmt.Errorf("Param Error!")
	}
	map_expression, ok := p_expression.(map[string]interface{})
	if !ok {
		r_buf.WriteString("[Result]:loadExpression Fail;")
		r_buf.WriteString("[Error]:Assert Error!")
		uniledgerlog.Error(r_buf.String())
		return fmt.Errorf("Assert Error!")
	}

	switch map_expression["Ctype"] {
	case constdef.ComponentType[constdef.Component_Expression] + "." + constdef.ExpressionType[constdef.Expression_Condition]:
		object_expression := expression.NewLogicArgument()
		byte_task, err := json.Marshal(map_expression)
		if err != nil {
			r_buf.WriteString("[Result]:loadExpression(Condition) Fail;")
			r_buf.WriteString("[Error]:Component_Expression(Expression_Condition) Serialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_task, &object_expression)
		if err != nil {
			r_buf.WriteString("[Result]:loadExpression(Condition) Fail;")
			r_buf.WriteString("[Error]:Component_Expression(Expression_Condition) Deserialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		err = object_expression.InitLogicArgument()
		if err != nil {
			r_buf.WriteString("[Result]:loadExpression(Condition) Fail;")
			r_buf.WriteString("[Error]:Error in loadExpression(Component_Expression(Expression_Condition)InitLogicArgument) ," + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		p_contract.AddComponent(object_expression)
	case constdef.ComponentType[constdef.Component_Expression] + "." + constdef.ExpressionType[constdef.Expression_Function]:
		object_expression := expression.NewFunction()
		byte_task, err := json.Marshal(map_expression)
		if err != nil {
			r_buf.WriteString("[Result]:loadExpression(Function) Fail;")
			r_buf.WriteString("[Error]:Component_Expression(Expression_Function) Serialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_task, &object_expression)
		if err != nil {
			r_buf.WriteString("[Result]:loadExpression(Function) Fail;")
			r_buf.WriteString("[Error]:Component_Expression(Expression_Function) Deserialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		err = object_expression.InitFunction()
		if err != nil {
			r_buf.WriteString("[Result]:loadExpression(Function) Fail;")
			r_buf.WriteString("[Error]:Error in loadExpression(Component_Expression(Expression_Function)InitFunction) ," + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		p_contract.AddComponent(object_expression)
	default:
		object_expression := &expression.GeneralExpression{}
		byte_task, err := json.Marshal(map_expression)
		if err != nil {
			r_buf.WriteString("[Result]:loadExpression(Unknow) Fail;")
			r_buf.WriteString("[Error]:Component_Expression(Expression_Unknow) Serialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_task, &object_expression)
		if err != nil {
			r_buf.WriteString("[Result]:loadExpression(Unknow) Fail;")
			r_buf.WriteString("[Error]:Component_Expression(Expression_Unknow) Deserialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}

		err = object_expression.InitExpression()
		if err != nil {
			r_buf.WriteString("[Result]:loadExpression(Unknow) Fail;")
			r_buf.WriteString("[Error]:Error in loadExpression(Component_Expression(Expression_Unknow)InitFunction) ," + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		p_contract.AddComponent(object_expression)
	}
	r_buf.WriteString("[Cname]: " + map_expression["Cname"].(string) + "[Ctype]: " + map_expression["Ctype"].(string) + "[Result]: loadExpression success;")
	//uniledgerlog.Info(r_buf.String())
	return err
}

func loadCandidate(p_contract *contract.CognitiveContract, p_candidate interface{}, p_task interface{}) error {
	var err error
	var r_buf bytes.Buffer
	r_buf.WriteString("loadCandidate...;")
	if p_contract == nil || p_task == nil || p_candidate == nil {
		r_buf.WriteString("[Result]:loadCandidate fail;")
		r_buf.WriteString("[Error]:Param Error!")
		uniledgerlog.Warn(r_buf.String())
		return fmt.Errorf("Param Error!")
	}
	map_candidate := p_candidate.(map[string]interface{})
	switch map_candidate["Ctype"] {
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.ExpressionType[constdef.Task_DecisionCandidate]:
		object_candidate := task.NewDecisionCandidate()
		byte_task, err := json.Marshal(map_candidate)
		if err != nil {
			r_buf.WriteString("[Result]:loadCandidate Fail;")
			r_buf.WriteString("[Error]:Component_Task(Task_DecisionCandidate) Serialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_task, &object_candidate)
		if err != nil {
			r_buf.WriteString("[Result]:loadCandidate Fail;")
			r_buf.WriteString("[Error]:Component_Task(Task_DecisionCandidate) Deserialize Error! " + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		err = object_candidate.InitDecisionCandidate()
		if err != nil {
			r_buf.WriteString("[Result]:loadCandidate Fail;")
			r_buf.WriteString("[Error]:Error in loadExpression(Component_Task(Task_DecisionCandidate)InitDecisionCandidate) ," + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
		p_contract.AddComponent(object_candidate)
	}
	r_buf.WriteString("[Cname]: " + map_candidate["Cname"].(string) + "[Ctype]: " + map_candidate["Ctype"].(string) + "[Result]: loadCandidate success;")
	//uniledgerlog.Info(r_buf.String())
	return err
}
