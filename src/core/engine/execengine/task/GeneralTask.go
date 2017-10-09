package task

//描述态：属性为数组
//运行态：属性为map
//		描述态 =》运行态： 在Init中进行转化
//		运行态 =》描述态： 在反序列化中进行转化
import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	common0 "unicontract/src/common"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/component"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/data"
	"unicontract/src/core/engine/execengine/expression"
	"unicontract/src/core/engine/execengine/function"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
	"unicontract/src/core/engine/gRPCClient"
)

type GeneralTask struct {
	component.GeneralComponent
	TaskId string `json:"TaskId"`
	State  string `json:"State"`
	//type:inf.IExpression
	PreCondition []interface{} `json:"PreCondition"`
	//type:inf.IExpression
	CompleteCondition []interface{} `json:"CompleteCondition"`
	//type:inf.IExpression
	DiscardCondition []interface{} `json:"DiscardCondition"`
	//type:inf.IData
	DataList []interface{} `json:"DataList"`
	//type:inf.IExpression
	DataValueSetterExpressionList []interface{} `json:"DataValueSetterExpressionList"`
	//type:int
	TaskExecuteIdx int      `json:"TaskExecuteIdx"`
	NextTasks      []string `json:"NextTasks"`
	//选择分支条件，查询操作后，根据查询结果进行分支判定，以分支判定的值为最终值，保持多节点一致性
	SelectBranches []common.SelectBranchExpression `json:"SelectBranches"`
}

const (
	_TaskId                        = "_TaskId"
	_State                         = "_State"
	_PreCondition                  = "_PreCondition"
	_CompleteCondition             = "_CompleteCondition"
	_DiscardCondition              = "_DiscardCondition"
	_DataList                      = "_DataList"
	_DataValueSetterExpressionList = "_DataValueSetterExpressionList"
	_TaskExecuteIdx                = "_TaskExecuteIdx"
	_NextTasks                     = "_NextTasks"
	_SelectBranches                = "_SelectBranches"
)

func NewGeneralTask() *GeneralTask {
	v_task := &GeneralTask{}
	return v_task
}

//===============接口实现===================
func (gt GeneralTask) SetContract(p_contract inf.ICognitiveContract) {
	gt.GeneralComponent.SetContract(p_contract)
}

func (gt GeneralTask) GetContract() inf.ICognitiveContract {
	return gt.GeneralComponent.GetContract()
}

func (gt GeneralTask) GetName() string {
	return gt.GeneralComponent.GetCname()
}
func (gt GeneralTask) GetCtype() string {
	return gt.GeneralComponent.GetCtype()
}

func (gt GeneralTask) GetDescription() string {
	return gt.GeneralComponent.GetDescription()
}

func (gt *GeneralTask) GetState() string {
	if gt.PropertyTable[_State] == nil {
		return ""
	}
	state_property, ok := gt.PropertyTable[_State].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	str, ok := state_property.GetValue().(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	return str
}

func (gt *GeneralTask) SetState(p_state string) {
	gt.State = p_state
	state_property, ok := gt.PropertyTable[_State].(property.PropertyT)
	if !ok {
		state_property = *property.NewPropertyT(_State)
	}
	state_property.SetValue(p_state)
	gt.PropertyTable[_State] = state_property
}

func (gt *GeneralTask) GetNextTasks() []string {
	if gt.PropertyTable[_NextTasks] == nil {
		return nil
	}
	nexttask_property, ok := gt.PropertyTable[_NextTasks].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	sl, ok := nexttask_property.GetValue().([]string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	return sl
}

//当前任务生命周期的执行：（根据任务状态选择相应的执行态方法进入）
//入口时机：加载中的任务执行完成Discard,执行下一可执行任务Dormant
//执行过程：PreProcess => Start or Complete or Discard => PostProcess
//执行结果：
//    1. ret = -1：执行失败, 需要回滚
//    2. ret = 0 ：执行条件未达到
//    3. ret = 1 ：执行完成,转入后继任务
func (gt GeneralTask) UpdateState(nBrotherNum int) (int8, error) {
	var r_ret int8 = 0
	var r_err error = nil
	var r_str_error string = ""
	var r_flag int8 = -1
	if &gt == nil {
		r_ret = -1
		r_err = fmt.Errorf("Object pointer is null!")
		return r_ret, r_err
	}

	//预处理
	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), begin to preprocess]",
		uniledgerlog.NO_ERROR, gt.GetContract().GetContractId(), gt.GetName(), gt.GetTaskId()))
	r_err = gt.PreProcess()
	if r_err != nil {
		uniledgerlog.Error("Task[" + gt.GetName() + "] PreProcess fail!")
		return r_ret, r_err
	}

	//处理中
	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), begin to execute]",
		uniledgerlog.NO_ERROR, gt.GetContract().GetContractId(), gt.GetName(), gt.GetTaskId()))
	r_ret, process_err := gt.Start()
	if process_err != nil {
		r_str_error = r_str_error + "[Run_Error]:" + process_err.Error()
	}
	switch r_ret {
	case 1:
		//正常执行，转入下一任务
		r_flag = 1
	case -1:
		//轮询等待后，执行失败，则暂时退出
		r_flag = -1
	case 0:
		//轮询等待后，条件不成立，则暂时退出
		r_flag = 0
	}

	//后处理
	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), begin to postprocess]",
		uniledgerlog.NO_ERROR, gt.GetContract().GetContractId(), gt.GetName(), gt.GetTaskId()))
	postProcess_err := gt.PostProcess(r_flag, nBrotherNum)
	if postProcess_err != nil {
		r_str_error = r_str_error + "[PostProcess_Error]" + postProcess_err.Error()
	}
	if r_str_error != "" {
		r_err = fmt.Errorf(r_str_error)
	}
	return r_ret, r_err
}

func (gt *GeneralTask) GetTaskId() string {
	if gt.PropertyTable[_TaskId] == nil {
		return ""
	}
	taskid_property, ok := gt.PropertyTable[_TaskId].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	str, ok := taskid_property.GetValue().(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	return str
}

func (gt *GeneralTask) GetTaskExecuteIdx() int {
	taskexecuteidx_property, ok := gt.PropertyTable[_TaskExecuteIdx].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return -1
	}
	taskexecuteidx_int, ok := taskexecuteidx_property.GetValue().(int)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return -1
	}
	return taskexecuteidx_int
}

func (gt *GeneralTask) SetTaskId(str_taskId string) {
	//Take case: Setter method need set value for gc.xxxxxx
	gt.TaskId = str_taskId
	taskid_property, ok := gt.PropertyTable[_TaskId].(property.PropertyT)
	if !ok {
		taskid_property = *property.NewPropertyT(_TaskId)
	}
	taskid_property.SetValue(str_taskId)
	//Take case: Setter method need set value for gc.PropertyTable[xxxx]
	gt.PropertyTable[_TaskId] = taskid_property
}

func (gt *GeneralTask) SetTaskExecuteIdx(int_idx int) {
	//Take case: Setter method need set value for gc.xxxxxx
	gt.TaskExecuteIdx = int_idx
	taskexecuteidx_property, ok := gt.PropertyTable[_TaskExecuteIdx].(property.PropertyT)
	if !ok {
		taskexecuteidx_property = *property.NewPropertyT(_TaskExecuteIdx)
	}
	taskexecuteidx_property.SetValue(int_idx)
	//Take case: Setter method need set value for gc.PropertyTable[xxxx]
	gt.PropertyTable[_TaskExecuteIdx] = taskexecuteidx_property
}

//清空任务组件中的中间结果值
//  清空内容：DataList:    DataValueSetterExpressionList:
func (gt *GeneralTask) CleanValueInProcess() {
	if gt.GetDataList() != nil {
		for v_key, v_value := range gt.GetDataList() {
			v_value.CleanValueInProcess()
			gt.GetContract().UpdateComponentRunningState(constdef.ComponentType[constdef.Component_Data], v_key, v_value)
		}
	}
	if gt.GetDataValueSetterExpressionList() != nil {
		for v_key, v_value := range gt.GetDataValueSetterExpressionList() {
			v_value.CleanValueInProcess()
			gt.GetContract().UpdateComponentRunningState(constdef.ComponentType[constdef.Component_Expression], v_key, v_value)
		}
	}
}

//===============描述态=====================
//反序列化实现运行态 map结构 到 数组结构的转化
//将任务中的执行态的Data & Expression 属性更新到描述态中
func (gt *GeneralTask) UpdateStaticState() (interface{}, error) {
	var err error = nil
	// State
	gt.State = gt.GetState()
	// TaskExecuteIdx
	gt.TaskExecuteIdx = gt.GetTaskExecuteIdx()
	// Data组件信息 更新到描述态
	var new_data_array []interface{} = make([]interface{}, len(gt.DataList))
	for v_idx := range gt.DataList {
		if gt.DataList[v_idx] == nil {
			err = fmt.Errorf("gt.DataList has nil data!")
			uniledgerlog.Error("UpdateStaticState fail[" + err.Error() + "]")
			return nil, err
		}
		v_data_map := gt.DataList[v_idx].(map[string]interface{})
		new_data := gt.GetContract().GetData(v_data_map["Cname"].(string))

		data_json, _ := new_data.(inf.IData).Serialize()
		new_data_array[v_idx] = common.Deserialize(data_json)
	}
	gt.DataList = new_data_array

	//Expression组件(DataValueSetterExpressionList)信息 更新到描述态
	var new_expression_array []interface{} = make([]interface{}, len(gt.DataValueSetterExpressionList))
	for v_idx := range gt.DataValueSetterExpressionList {
		if gt.DataValueSetterExpressionList[v_idx] == nil {
			err = fmt.Errorf("gt.DataValueSetterExpressionList has nil data!")
			uniledgerlog.Error("UpdateStaticState fail[" + err.Error() + "]")
			return nil, err
		}
		v_expression_map := gt.DataValueSetterExpressionList[v_idx].(map[string]interface{})
		new_expression := gt.GetContract().GetExpression(v_expression_map["Cname"].(string))

		expression_json, _ := new_expression.(inf.IExpression).Serialize()
		new_expression_array[v_idx] = common.Deserialize(expression_json)
	}
	gt.DataValueSetterExpressionList = new_expression_array
	//Expression组件(PreCondition)信息 更新到描述态
	var new_pre_array []interface{} = make([]interface{}, len(gt.PreCondition))
	for v_idx := range gt.PreCondition {
		if gt.PreCondition[v_idx] == nil {
			err = fmt.Errorf("gt.PreCondition has nil data!")
			uniledgerlog.Error("UpdateStaticState fail[" + err.Error() + "]")
			return nil, err
		}
		v_pre_map := gt.PreCondition[v_idx].(map[string]interface{})
		new_pre := gt.GetContract().GetExpression(v_pre_map["Cname"].(string))

		expression_json, _ := new_pre.(inf.IExpression).Serialize()
		new_pre_array[v_idx] = common.Deserialize(expression_json)

	}
	gt.PreCondition = new_pre_array
	//Expression组件(CompleteCondition)信息 更新到描述态
	var new_complete_array []interface{} = make([]interface{}, len(gt.CompleteCondition))
	for v_idx := range gt.CompleteCondition {
		if gt.CompleteCondition[v_idx] == nil {
			err = fmt.Errorf("gt.CompleteCondition has nil data!")
			uniledgerlog.Error("UpdateStaticState fail[" + err.Error() + "]")
			return nil, err
		}
		v_complete_map := gt.CompleteCondition[v_idx].(map[string]interface{})
		new_complete := gt.GetContract().GetExpression(v_complete_map["Cname"].(string))

		expression_json, _ := new_complete.(inf.IExpression).Serialize()
		new_complete_array[v_idx] = common.Deserialize(expression_json)
	}
	gt.CompleteCondition = new_complete_array
	//Expression组件(DiscardCondition)信息 更新到描述态
	var new_discard_array []interface{} = make([]interface{}, len(gt.DiscardCondition))
	for v_idx := range gt.DiscardCondition {
		if gt.DiscardCondition[v_idx] == nil {
			err = fmt.Errorf("gt.DiscardCondition has nil data!")
			uniledgerlog.Error("UpdateStaticState fail[" + err.Error() + "]")
			return nil, err
		}
		v_discard_map := gt.DiscardCondition[v_idx].(map[string]interface{})
		new_discard := gt.GetContract().GetExpression(v_discard_map["Cname"].(string))

		expression_json, _ := new_discard.(inf.IExpression).Serialize()
		new_discard_array[v_idx] = common.Deserialize(expression_json)
	}
	gt.DiscardCondition = new_discard_array
	return gt, err
}

//===============运行态=====================
//Init中实现描述态反序列化后得到的Map[string]map结构到 Map[string]Component对象的转化
func (gt *GeneralTask) InitGeneralTask() error {
	//uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s) %s]",
	//	uniledgerlog.NO_ERROR, gt.GetContract().GetContractId(), "init general task"))
	var err error = nil
	err = gt.InitGeneralComponent()
	if err != nil {
		uniledgerlog.Error("InitGeneralTask fail[" + err.Error() + "]")
		return err
	}
	gt.SetCtype(constdef.ComponentType[constdef.Component_Task])
	common.AddProperty(gt, gt.PropertyTable, _TaskId, gt.TaskId)
	// State default
	gt.State = common.TernaryOperator(gt.State == "", constdef.ComponentType[constdef.TaskState_Dormant], gt.State).(string)
	common.AddProperty(gt, gt.PropertyTable, _State, gt.State)
	common.AddProperty(gt, gt.PropertyTable, _TaskExecuteIdx, gt.TaskExecuteIdx)
	//PreCondition array to map
	if gt.PreCondition == nil {
		gt.PreCondition = make([]interface{}, 0)
	}
	map_precondition := make(map[string]inf.IExpression, 0)
	for _, p_precondition := range gt.PreCondition {
		if p_precondition != nil {
			switch p_precondition.(type) {
			case inf.IExpression:
			case *inf.IExpression:
				tmp_precondition := p_precondition.(inf.IExpression)
				map_precondition[tmp_precondition.GetName()] = tmp_precondition
			case map[string]interface{}:
				tmp_precondition := expression.NewLogicArgument()
				tmp_byte_precondition, _ := json.Marshal(p_precondition)
				err = json.Unmarshal(tmp_byte_precondition, &tmp_precondition)
				if err != nil {
					uniledgerlog.Error("InitGeneralTask(PreCondition) fail[" + err.Error() + "]")
					return err
				}
				tmp_precondition.InitLogicArgument()
				map_precondition[tmp_precondition.GetName()] = tmp_precondition
			}
		}
	}
	common.AddProperty(gt, gt.PropertyTable, _PreCondition, map_precondition)
	//CompleteCondition arrat to map
	if gt.CompleteCondition == nil {
		gt.CompleteCondition = make([]interface{}, 0)
	}
	map_completecondition := make(map[string]inf.IExpression, 0)
	for _, p_completecondition := range gt.CompleteCondition {
		if p_completecondition != nil {
			switch p_completecondition.(type) {
			case inf.IExpression:
			case *inf.IExpression:
				tmp_completecondition := p_completecondition.(inf.IExpression)
				map_completecondition[tmp_completecondition.GetName()] = tmp_completecondition
			case map[string]interface{}:
				tmp_completecondition := expression.NewLogicArgument()
				tmp_byte_completecondition, _ := json.Marshal(p_completecondition)
				err = json.Unmarshal(tmp_byte_completecondition, &tmp_completecondition)
				if err != nil {
					uniledgerlog.Error("InitGeneralTask(CompleteCondition) fail[" + err.Error() + "]")
					return err
				}
				tmp_completecondition.InitLogicArgument()
				map_completecondition[tmp_completecondition.GetName()] = tmp_completecondition
			}
		}
	}
	common.AddProperty(gt, gt.PropertyTable, _CompleteCondition, map_completecondition)
	//DiscardCondition arrat to map
	if gt.DiscardCondition == nil {
		gt.DiscardCondition = make([]interface{}, 0)
	}
	map_discardcondition := make(map[string]inf.IExpression, 0)
	for _, p_discardcondition := range gt.DiscardCondition {
		if p_discardcondition != nil {
			switch p_discardcondition.(type) {
			case inf.IExpression:
			case *inf.IExpression:
				tmp_discardcondition := p_discardcondition.(inf.IExpression)
				map_discardcondition[tmp_discardcondition.GetName()] = tmp_discardcondition
			case map[string]interface{}:
				tmp_discardcondition := expression.NewLogicArgument()
				tmp_byte_discardcondition, _ := json.Marshal(p_discardcondition)
				err = json.Unmarshal(tmp_byte_discardcondition, &tmp_discardcondition)
				if err != nil {
					uniledgerlog.Error("InitGeneralTask(DiscardCondition) fail[" + err.Error() + "]")
					return err
				}
				tmp_discardcondition.InitLogicArgument()
				map_discardcondition[tmp_discardcondition.GetName()] = tmp_discardcondition
			}
		}
	}
	common.AddProperty(gt, gt.PropertyTable, _DiscardCondition, map_discardcondition)
	//DataList arr to map
	if gt.DataList == nil {
		gt.DataList = make([]interface{}, 0)
	}
	map_datalist := make(map[string]inf.IData, 0)
	for _, p_data := range gt.DataList {
		if p_data != nil {
			switch p_data.(type) {
			case inf.IData:
			case *inf.IData:
				tmp_data := p_data.(inf.IData)
				map_datalist[tmp_data.GetName()] = tmp_data
			case map[string]interface{}:
				p_data_map := p_data.(map[string]interface{})
				_, ok := p_data_map["Ctype"].(string)
				if !ok {
					uniledgerlog.Error("assert error!!")
					continue
				}
				switch p_data_map["Ctype"].(string) {
				case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Numeric_Int]:
					tmp_data := data.NewIntData()
					tmp_byte_data, _ := json.Marshal(p_data)
					err = json.Unmarshal(tmp_byte_data, &tmp_data)
					if err != nil {
						uniledgerlog.Error("InitGeneralTask(DataList) fail[" + err.Error() + "]")
						return err
					}
					tmp_data.InitIntData()
					map_datalist[tmp_data.GetName()] = tmp_data
				case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Numeric_Uint]:
					tmp_data := data.NewUintData()
					tmp_byte_data, _ := json.Marshal(p_data)
					err = json.Unmarshal(tmp_byte_data, &tmp_data)
					if err != nil {
						uniledgerlog.Error("InitGeneralTask(DataList) fail[" + err.Error() + "]")
						return err
					}
					tmp_data.InitUintData()
					map_datalist[tmp_data.GetName()] = tmp_data
				case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Numeric_Float]:
					tmp_data := data.NewFloatData()
					tmp_byte_data, _ := json.Marshal(p_data)
					err = json.Unmarshal(tmp_byte_data, &tmp_data)
					if err != nil {
						uniledgerlog.Error("InitGeneralTask(DataList) fail[" + err.Error() + "]")
						return err
					}
					tmp_data.InitFloatData()
					map_datalist[tmp_data.GetName()] = tmp_data
				case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Text]:
					tmp_data := data.NewTextData()
					tmp_byte_data, _ := json.Marshal(p_data)
					err = json.Unmarshal(tmp_byte_data, &tmp_data)
					if err != nil {
						uniledgerlog.Error("InitGeneralTask(DataList) fail[" + err.Error() + "]")
						return err
					}
					tmp_data.InitTextData()
					map_datalist[tmp_data.GetName()] = tmp_data
				case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Date]:
					tmp_data := data.NewDateData()
					tmp_byte_data, _ := json.Marshal(p_data)
					err = json.Unmarshal(tmp_byte_data, &tmp_data)
					if err != nil {
						uniledgerlog.Error("InitGeneralTask(DataList) fail[" + err.Error() + "]")
						return err
					}
					tmp_data.InitDateData()
					map_datalist[tmp_data.GetName()] = tmp_data
				case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Bool]:
					tmp_data := data.NewBoolData()
					tmp_byte_data, _ := json.Marshal(p_data)
					err = json.Unmarshal(tmp_byte_data, &tmp_data)
					if err != nil {
						uniledgerlog.Error("InitGeneralTask(DataList) fail[" + err.Error() + "]")
						return err
					}
					tmp_data.InitBoolData()
					map_datalist[tmp_data.GetName()] = tmp_data
				case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Array]:
					tmp_data := data.NewArrayData()
					tmp_byte_data, _ := json.Marshal(p_data)
					err = json.Unmarshal(tmp_byte_data, &tmp_data)
					if err != nil {
						uniledgerlog.Error("InitGeneralTask(DataList) fail[" + err.Error() + "]")
						return err
					}
					tmp_data.InitArrayData()
					map_datalist[tmp_data.GetName()] = tmp_data
				case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_OperateResult]:
					tmp_data := data.NewOperateResultData()
					tmp_byte_data, _ := json.Marshal(p_data)
					err = json.Unmarshal(tmp_byte_data, &tmp_data)
					if err != nil {
						uniledgerlog.Error("InitGeneralTask(DataList) fail[" + err.Error() + "]")
						return err
					}
					tmp_data.InitOperateResultData()
					map_datalist[tmp_data.GetName()] = tmp_data
				}
			}
		}
	}
	common.AddProperty(gt, gt.PropertyTable, _DataList, map_datalist)
	//DataValueSetterExpressionList arr to map
	if gt.DataValueSetterExpressionList == nil {
		gt.DataValueSetterExpressionList = make([]interface{}, 0)
	}
	map_dataexpressionlist := make(map[string]inf.IExpression, 0)
	for _, p_express := range gt.DataValueSetterExpressionList {
		if p_express != nil {
			switch p_express.(type) {
			case inf.IExpression:
			case *inf.IExpression:
				tmp_express := p_express.(inf.IExpression)
				map_dataexpressionlist[tmp_express.GetName()] = tmp_express
			case map[string]interface{}:
				tmp_express := expression.NewFunction()
				tmp_byte_express, _ := json.Marshal(p_express)
				err = json.Unmarshal(tmp_byte_express, &tmp_express)
				if err != nil {
					uniledgerlog.Error("InitGeneralTask(DataValueSetterExpressionList) fail[" + err.Error() + "]")
					return err
				}
				tmp_express.InitFunction()
				map_dataexpressionlist[tmp_express.GetName()] = tmp_express
			}
		}
	}
	common.AddProperty(gt, gt.PropertyTable, _DataValueSetterExpressionList, map_dataexpressionlist)
	//nextTask array to map
	if gt.NextTasks == nil {
		gt.NextTasks = make([]string, 0)
	}
	common.AddProperty(gt, gt.PropertyTable, _NextTasks, gt.NextTasks)

	if gt.SelectBranches == nil {
		gt.SelectBranches = make([]common.SelectBranchExpression, 0)
	}
	common.AddProperty(gt, gt.PropertyTable, _SelectBranches, gt.SelectBranches)
	return err
}

//====属性Get方法
func (gt *GeneralTask) GetPreCondition() map[string]inf.IExpression {
	if gt.PropertyTable[_PreCondition] == nil {
		return nil
	}
	precondition_property, ok := gt.PropertyTable[_PreCondition].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	precondition_value, ok := precondition_property.GetValue().(map[string]inf.IExpression)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	return precondition_value
}

func (gt *GeneralTask) GetCompleteCondition() map[string]inf.IExpression {
	if gt.PropertyTable[_CompleteCondition] == nil {
		return nil
	}
	completecondition_property, ok := gt.PropertyTable[_CompleteCondition].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	completecondition_value, ok := completecondition_property.GetValue().(map[string]inf.IExpression)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	return completecondition_value
}

func (gt *GeneralTask) GetDiscardCondition() map[string]inf.IExpression {
	if gt.PropertyTable[_DiscardCondition] == nil {
		return nil
	}
	discardcondition_property, ok := gt.PropertyTable[_DiscardCondition].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	discardcondition_value, ok := discardcondition_property.GetValue().(map[string]inf.IExpression)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	return discardcondition_value
}

func (gt *GeneralTask) GetDataValueSetterExpressionList() map[string]inf.IExpression {
	if gt.PropertyTable[_DataValueSetterExpressionList] == nil {
		return nil
	}
	dataexpress_property, ok := gt.PropertyTable[_DataValueSetterExpressionList].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	dataexpress_value, ok := dataexpress_property.GetValue().(map[string]inf.IExpression)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	return dataexpress_value
}

func (gt *GeneralTask) GetDataList() map[string]inf.IData {
	if gt.PropertyTable[_DataList] == nil {
		return nil
	}
	datalist_property, ok := gt.PropertyTable[_DataList].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	datalist_value, ok := datalist_property.GetValue().(map[string]inf.IData)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	return datalist_value
}

func (gt *GeneralTask) GetSelectBranches() []common.SelectBranchExpression {
	if gt.PropertyTable[_SelectBranches] == nil {
		return nil
	}
	selectbranch_property, ok := gt.PropertyTable[_SelectBranches].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	selectbranch_value, ok := selectbranch_property.GetValue().([]common.SelectBranchExpression)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	return selectbranch_value
}

//====属性动态初始化
func (gt *GeneralTask) ReSet() {
	gt.SetState(constdef.TaskState[constdef.TaskState_Dormant])
}

func (gt *GeneralTask) AddNextTasks(task string) {
	nexttask_property, ok := gt.PropertyTable[_NextTasks].(property.PropertyT)
	if !ok {
		nexttask_property = *property.NewPropertyT(_NextTasks)
	}
	if nexttask_property.GetValue() == nil {
		nexttask_property.SetValue(make([]string, 0))
	}
	if task != "" {
		arr_nexttasks := nexttask_property.GetValue().([]string)
		arr_nexttasks = append(arr_nexttasks, task)
		nexttask_property.SetValue(arr_nexttasks)
		gt.PropertyTable[_NextTasks] = nexttask_property
	}
}

func (gt *GeneralTask) AddPreCondition(p_name string, p_condition string) {
	precondition_property, ok := gt.PropertyTable[_PreCondition].(property.PropertyT)
	if !ok {
		precondition_property = *property.NewPropertyT(_PreCondition)
	}
	if precondition_property.GetValue() == nil {
		precondition_property.SetValue(make(map[string]inf.IExpression, 0))
	}
	map_precondition := precondition_property.GetValue().(map[string]inf.IExpression)
	map_precondition[p_name] = expression.NewGeneralExpression(p_name, p_condition)

	precondition_property.SetValue(map_precondition)
	gt.PropertyTable[_PreCondition] = precondition_property
}

func (gt *GeneralTask) AddCompleteCondition(p_name string, p_condition string) {
	completecondition_property, ok := gt.PropertyTable[_CompleteCondition].(property.PropertyT)
	if !ok {
		completecondition_property = *property.NewPropertyT(_CompleteCondition)
	}
	if completecondition_property.GetValue() == nil {
		completecondition_property.SetValue(make(map[string]inf.IExpression, 0))
	}
	map_completecondition := completecondition_property.GetValue().(map[string]inf.IExpression)
	map_completecondition[p_name] = expression.NewGeneralExpression(p_name, p_condition)

	completecondition_property.SetValue(map_completecondition)
	gt.PropertyTable[_CompleteCondition] = completecondition_property
}

func (gt *GeneralTask) AddDiscardCondition(p_name string, p_condition string) {
	Discardcondition_property, ok := gt.PropertyTable[_DiscardCondition].(property.PropertyT)
	if !ok {
		Discardcondition_property = *property.NewPropertyT(_DiscardCondition)
	}
	if Discardcondition_property.GetValue() == nil {
		Discardcondition_property.SetValue(make([]inf.IExpression, 0))
	}
	map_Discardcondition := Discardcondition_property.GetValue().(map[string]inf.IExpression)
	map_Discardcondition[p_name] = expression.NewGeneralExpression(p_name, p_condition)

	Discardcondition_property.SetValue(map_Discardcondition)
	gt.PropertyTable[_DiscardCondition] = Discardcondition_property
}

func (gt *GeneralTask) AddDataSetterExpressionAndData(p_name string, p_dataSetterExpresstionStr string, p_data inf.IData) {
	gt.AddDataSetterExpression(p_name, p_dataSetterExpresstionStr)
	gt.AddData(p_data)
}

func (gt *GeneralTask) AddDataSetterExpression(p_name string, p_dataSetterExpresstionStr string) {
	dataexpressionlist_property, ok := gt.PropertyTable[_DataValueSetterExpressionList].(property.PropertyT)
	if !ok {
		dataexpressionlist_property = *property.NewPropertyT(_DataValueSetterExpressionList)
	}
	if dataexpressionlist_property.GetValue() == nil {
		map_dataexpressionlist := make(map[string]inf.IExpression, 0)
		dataexpressionlist_property.SetValue(map_dataexpressionlist)
	}
	if p_dataSetterExpresstionStr != "" {
		map_dataexpresslist := dataexpressionlist_property.GetValue().(map[string]inf.IExpression)
		map_dataexpresslist[p_name] = expression.NewGeneralExpression(p_name, p_dataSetterExpresstionStr)
		dataexpressionlist_property.SetValue(map_dataexpresslist)
		gt.PropertyTable[_DataValueSetterExpressionList] = dataexpressionlist_property
	}
}

func (gt *GeneralTask) AddData(p_data inf.IData) {
	datalist_property, ok := gt.PropertyTable[_DataList].(property.PropertyT)
	if !ok {
		datalist_property = *property.NewPropertyT(_DataList)
	}
	if datalist_property.GetValue() == nil {
		map_datalist := make(map[string]inf.IData, 0)
		datalist_property.SetValue(map_datalist)
	}
	map_datalist := datalist_property.GetValue().(map[string]inf.IData)
	map_datalist[p_data.GetName()] = p_data
	datalist_property.SetValue(map_datalist)
	gt.PropertyTable[_DataList] = datalist_property
}

func (gt *GeneralTask) RemoveDataSetterExpressionAndData(p_expressionname string, p_dataname string) {
	gt.RemoveDataSetterExpression(p_expressionname)
	gt.RemoveData(p_dataname)
}

func (gt *GeneralTask) RemoveDataSetterExpression(p_expressionname string) {
	if gt.PropertyTable[_DataValueSetterExpressionList] == nil {
		return
	}
	dataExpression_property, ok := gt.PropertyTable[_DataValueSetterExpressionList].(property.PropertyT)
	if !ok {
		dataExpression_property = *property.NewPropertyT(_DataValueSetterExpressionList)
	}
	if dataExpression_property.GetValue() != nil {
		map_dataExpression := dataExpression_property.GetValue().(map[string]inf.IExpression)
		delete(map_dataExpression, p_expressionname)
		dataExpression_property.SetValue(map_dataExpression)
		gt.PropertyTable[_DataValueSetterExpressionList] = dataExpression_property
	}
	return
}

func (gt *GeneralTask) RemoveData(p_name string) {
	if gt.PropertyTable[_DataList] == nil {
		return
	}
	datalist_property, ok := gt.PropertyTable[_DataList].(property.PropertyT)
	if !ok {
		datalist_property = *property.NewPropertyT(_DataList)
	}
	if datalist_property.GetValue() != nil {
		map_datalist := datalist_property.GetValue().(map[string]inf.IData)
		delete(map_datalist, p_name)
		datalist_property.SetValue(map_datalist)
		gt.PropertyTable[_DataList] = datalist_property
	}
	return
}

//====属性Set方法
func (gt *GeneralTask) SetSelectBranches(p_selectbranchs []common.SelectBranchExpression) {
	gt.SelectBranches = p_selectbranchs
	selectbranch_property, ok := gt.PropertyTable[_SelectBranches].(property.PropertyT)
	if !ok {
		selectbranch_property = *property.NewPropertyT(_SelectBranches)
	}
	selectbranch_property.SetValue(p_selectbranchs)
	gt.PropertyTable[_SelectBranches] = selectbranch_property
}

func (gt *GeneralTask) GetData(p_name string) (interface{}, error) {
	var err error = nil
	datalist_property, ok := gt.PropertyTable[_DataList].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil, err
	}
	if datalist_property.GetValue() != nil {
		data_map, ok := datalist_property.GetValue().(map[string]inf.IData)
		if !ok {
			err = fmt.Errorf("Find data[" + p_name + "] Error!")
			return nil, err
		}
		r_data, ok := data_map[p_name]
		if !ok {
			err = fmt.Errorf("Find data[" + p_name + "] Error!")
		}
		return r_data, err
	} else {
		err = fmt.Errorf("DataList is nil,find data[" + p_name + "] Error!")
		return nil, err
	}
}

func (gt *GeneralTask) GetDataExpression(p_name string) (interface{}, error) {
	var err error = nil
	dataexpressionlist_property, ok := gt.PropertyTable[_DataValueSetterExpressionList].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil, err
	}
	if dataexpressionlist_property.GetValue() != nil {
		dataexpression_map, ok := dataexpressionlist_property.GetValue().(map[string]inf.IExpression)
		if !ok {
			err = fmt.Errorf("Find dataExpression[" + p_name + "] Error!")
			return nil, err
		}
		r_data, ok := dataexpression_map[p_name]
		if !ok {
			err = fmt.Errorf("Find dataExpression[" + p_name + "] Error!")
		}
		return r_data, err
	} else {
		err = fmt.Errorf("DataValueSetterExpressionList is nil,find dataExpression[" + p_name + "] Error!")
		return nil, err
	}
}

//====运行条件判断
func (gt *GeneralTask) testCompleteCondition() bool {
	var r_flag bool = true

	if len(gt.GetCompleteCondition()) == 0 {
		r_flag = true
	}

	for _, value := range gt.GetCompleteCondition() {
		v_contract := gt.GetContract()
		v_bool, v_err := v_contract.EvaluateExpression(constdef.ExpressionType[constdef.Expression_Condition], value.GetExpressionStr())
		if v_err != nil {
			uniledgerlog.Error("CompleteCondition EvaluateExpression[" + value.GetExpressionStr() + " fail, [" + v_err.Error() + "]")
			r_flag = false
			break
		}
		if !v_bool.(bool) {
			r_flag = false
			break
		}
	}

	return r_flag
}

func (gt *GeneralTask) testDiscardCondition() bool {
	var r_flag bool = true
	if len(gt.GetDiscardCondition()) == 0 {
		r_flag = true
	}

	//合约录入的终止条件
	for _, value := range gt.GetDiscardCondition() {
		v_contract := gt.GetContract()
		v_bool, v_err := v_contract.EvaluateExpression(constdef.ExpressionType[constdef.Expression_Condition], value.GetExpressionStr())
		if v_err != nil {
			uniledgerlog.Error("DiscardCondition EvaluateExpression[" + value.GetExpressionStr() + "] fail, [" + v_err.Error() + "]")
			r_flag = false
			break
		}
		if !v_bool.(bool) {
			r_flag = false
			break
		}
	}
	if r_flag == false {
		return r_flag
	}

	//默认的合约终止条件（当前合约步骤入链查询判定）
	v_contract := gt.GetContract()
	v_default_function := "FuncIsConPutInUnichian(\"" + v_contract.GetOutputId() + "\")"
	v_result, v_err := v_contract.EvaluateExpression(constdef.ExpressionType[constdef.Expression_Condition], v_default_function)
	if v_err != nil {
		uniledgerlog.Error("DiscardCondition EvaluateExpression[" + v_default_function + "] fail, [" + v_err.Error() + "]")
		r_flag = false
	}
	if !v_result.(bool) {
		uniledgerlog.Error("DiscardCondition EvaluateExpression[" + v_default_function + "] result is false!")
		r_flag = false
	}
	return r_flag
}

func (gt *GeneralTask) testPreCondition() bool {
	var r_flag bool = true
	if len(gt.GetPreCondition()) == 0 {
		r_flag = true
	}
	for _, value := range gt.GetPreCondition() {
		v_contract := gt.GetContract()
		uniledgerlog.Debug("PreCondition is [%s]", value.GetExpressionStr())
		v_bool, v_err := v_contract.EvaluateExpression(constdef.ExpressionType[constdef.Expression_Condition], value.GetExpressionStr())
		if v_err != nil {
			uniledgerlog.Error("PreCondition EvaluateExpression[" + value.GetExpressionStr() + " fail, [" + v_err.Error() + "]")
			r_flag = false
			break
		}
		if !v_bool.(bool) {
			r_flag = false
			break
		}
	}
	return r_flag
}

//====运行状态控制
func (gt *GeneralTask) IsDormant() bool {
	return gt.GetState() == constdef.TaskState[constdef.TaskState_Dormant]
}

func (gt *GeneralTask) IsInProgress() bool {
	return gt.GetState() == constdef.TaskState[constdef.TaskState_In_Progress]
}

func (gt *GeneralTask) IsCompleted() bool {
	return gt.GetState() == constdef.TaskState[constdef.TaskState_Completed]
}

func (gt *GeneralTask) IsDiscarded() bool {
	return gt.GetState() == constdef.TaskState[constdef.TaskState_Discard]
}

//任务运行前进行的预处理
func (gt *GeneralTask) PreProcess() error {
	var r_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Task PreProcess Runing:")
	//当前合约执行任务为新任务,即OutputStruct
	if gt.GetContract() == nil {
		r_err = fmt.Errorf("PreProcess fail, GetContract is nil!")
		r_buf.WriteString("PreProcess fail, GetContract is nil!")
		uniledgerlog.Error(r_buf.String())
		return r_err
	}
	//outputTaskId, outputTaskExecuteIdx赋值
	gt.GetContract().SetOutputTaskId(gt.GetTaskId())
	gt.GetContract().SetOutputTaskExecuteIdx(gt.GetTaskExecuteIdx())
	return r_err
}

//用于执行回滚操作，回滚后将任务状态改为dormant
func (gt *GeneralTask) Dormant() (int8, error) {
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Task Process Runing:InProcess or Complete State.")
	r_buf.WriteString("[ContractID]: " + gt.GetContract().GetContractId() + ";")
	r_buf.WriteString("[TaskName]: " + gt.GetName() + ";")
	uniledgerlog.Info(r_buf.String(), " begin....")
	var r_ret int8 = 0
	var r_err error = nil
	if gt.IsInProgress() || gt.IsCompleted() {
		uniledgerlog.Info(r_buf.String(), " InProcess|Completed to Dormant....")
		gt.SetState(constdef.TaskState[constdef.TaskState_Dormant])
		gt.CleanValueInProcess()
	}
	return r_ret, r_err
}

func (gt *GeneralTask) Start() (int8, error) {
	var r_ret int8
	var r_err error
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Task Process Runing:Dormant State.")
	r_buf.WriteString("[ContractID]: " + gt.GetContract().GetContractId() + ";")
	r_buf.WriteString("[ContractHashID]: " + gt.GetContract().GetId() + ";")
	r_buf.WriteString("[TaskName]: " + gt.GetName() + ";")
	uniledgerlog.Info(r_buf.String(), "Start begin....")

	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), check start precondition]",
		uniledgerlog.NO_ERROR, gt.GetContract().GetContractId(), gt.GetName(), gt.GetTaskId()))
	if gt.IsDormant() && gt.testPreCondition() {
		var exec_flag bool = true
		//如果没有后继任务，则最后一个任务执行结束后，合约完成
		//在此处更新合约状态为【完成】,随着任务执行的产出，完成入链
		if len(gt.GetNextTasks()) == 0 {
			gt.GetContract().UpdateContractState(constdef.ContractState[constdef.Contract_Completed])
		}

		//var data_array []interface{} = gt.DataList
		//循环遍历函数表达式列表，执行函数
		//注意：限制只可有一个Output交易产出
		//TODO 待处理，避免一般操作任务，重复执行
		//TODO DataValueSetterExpressionList 和 Data的对应（通过 Cname进行对应， expression_function_A\data_int_expression_function_A）
		uniledgerlog.Debug("Task %s DataSetterExpressionList() size is %d", gt.GetName(), len(gt.GetDataValueSetterExpressionList()))
		for v_key := range gt.GetDataValueSetterExpressionList() {
			v_expr_object := gt.GetDataValueSetterExpressionList()[v_key]

			//1 执行
			// 1.1 函数识别 & 参数补齐
			str_name := v_expr_object.GetName()
			str_function := v_expr_object.GetExpressionStr()
			uniledgerlog.Info("%s execute Function name is %s", str_name, str_function)
			str_function = strings.TrimSpace(str_function)
			var v_beginwith_flag bool
			if !gRPCClient.On {
				reg := regexp.MustCompile("FuncTransferAsset\\(")
				v_str := reg.FindString(str_function)
				if "" != v_str {
					v_beginwith_flag = true
				} else {
					v_beginwith_flag = false
				}
			} else {
				slString := strings.Split(str_function, `(`)
				funcType, r_err := gRPCClient.QueryFuncType(slString[0])
				if r_err != nil {
					r_ret = -1
					r_buf.WriteString("[Result]: gRPC QueryFuncType failed;")
					r_buf.WriteString("[Error]: " + r_err.Error() + ";")
					r_buf.WriteString("Start fail....")
					uniledgerlog.Error(r_buf.String())
					return r_ret, r_err
				}
				uniledgerlog.Info("%s function %s type is %d", str_name, slString[0], funcType)

				if funcType == 2 { // 需要补齐参数
					v_beginwith_flag = true
				} else { // 不需要补齐参数
					v_beginwith_flag = false
				}
			}

			if v_beginwith_flag {
				str_json_contract, r_err := gt.GetContract().Serialize()
				if r_err != nil || len(str_json_contract) == 0 {
					r_ret = -1
					r_buf.WriteString("[Result]: Generate OutputStruct fail, str_json_contract Serialize fail;")
					if r_err != nil {
						r_buf.WriteString("[Error]: " + r_err.Error() + ";")
					}
					if len(str_json_contract) == 0 {
						r_buf.WriteString("[Error]: constract json string is null;")
					}
					r_buf.WriteString("Start fail...")
					uniledgerlog.Error(r_buf.String())
					return r_ret, r_err
				}
				uniledgerlog.Debug("%s before transfer asset contract is %s", str_name, str_json_contract)

				var func_buf bytes.Buffer = bytes.Buffer{}
				str_json_contract = strings.Replace(str_json_contract, "\\", "\\\\", -1)
				str_json_contract = strings.Replace(str_json_contract, "\"", "\\\"", -1)
				func_buf.WriteString(strings.Trim(str_function, ")"))
				func_buf.WriteString("@\"")
				func_buf.WriteString(str_json_contract)
				func_buf.WriteString("\"@\"")
				if gRPCClient.On {
					func_buf.WriteString("contractHashId")
					func_buf.WriteString("\",\"")
				}
				func_buf.WriteString(gt.GetContract().GetContractId())
				func_buf.WriteString("\",\"")
				func_buf.WriteString(gt.GetTaskId())
				func_buf.WriteString("\",")
				func_buf.WriteString(strconv.FormatInt(int64(gt.GetTaskExecuteIdx()), 10))
				func_buf.WriteString(",\"")
				func_buf.WriteString(gt.GetContract().GetMainPubkey())
				func_buf.WriteString("\")")
				str_function = func_buf.String()

				uniledgerlog.Debug("%s after add params function is %s", str_name, str_function)
			}

			// 1.2 函数执行
			uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), evaluate expression]",
				uniledgerlog.NO_ERROR, gt.GetContract().GetContractId(), gt.GetName(), gt.GetTaskId()))
			v_result, r_err := gt.GetContract().EvaluateExpression(constdef.ExpressionType[constdef.Expression_Function], str_function)
			if r_err != nil {
				r_ret = -1
				r_buf.WriteString("[Result]: EvaluateExpression run error;")
				r_buf.WriteString("[Error]: " + r_err.Error() + ";")
				r_buf.WriteString("Start fail...")
				uniledgerlog.Error(r_buf.String())
				return r_ret, r_err
			}
			v_result_object, ok := v_result.(common.OperateResult)
			if !ok {
				r_ret = -1
				r_buf.WriteString("[Result]: EvaluateExpression return value type is error;")
				r_buf.WriteString("[Error]: v_result.(common.OperateResult) assert error;")
				r_buf.WriteString("Start fail...")
				uniledgerlog.Error(r_buf.String())
				return r_ret, r_err
			}

			//2 执行结果赋值
			//  2.1 结果赋值到 data中,针对Enquiry Task，需要根据分支条件一致性化查询结果值
			uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), update consistent value]",
				uniledgerlog.NO_ERROR, gt.GetContract().GetContractId(), gt.GetName(), gt.GetTaskId()))
			gt.ConsistentValue(gt.GetDataList(), str_name, v_result_object)

			//  2.2 结果赋值到 dataSetterValue函数结果
			uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), update expression result]",
				uniledgerlog.NO_ERROR, gt.GetContract().GetContractId(), gt.GetName(), gt.GetTaskId()))
			v_expr_object.SetExpressionResult(v_result_object)
			gt.GetContract().UpdateComponentRunningState(constdef.ComponentType[constdef.Component_Expression],
				v_expr_object.GetName(), v_expr_object)

			now_json, _ := gt.GetContract().Serialize()
			uniledgerlog.Debug("%s after update component contract is %s", str_name, now_json)

			//  2.3 Output交易产出结构体赋值
			uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), update output]",
				uniledgerlog.NO_ERROR, gt.GetContract().GetContractId(), gt.GetName(), gt.GetTaskId()))
			if v_result_object.GetOutput() != nil {
				if !gRPCClient.On {
					strOutput, ok := v_result_object.GetOutput().(string)
					if ok {
						gt.GetContract().SetOutputStruct(strOutput)
						uniledgerlog.Debug("%s after transfer asset output is %s", str_name, strOutput)
					} else {
						r_ret = -1
						r_err = fmt.Errorf("v_result_object.GetOutput().(string) assert error")
						uniledgerlog.Error(r_err.Error())
						return r_ret, r_err
					}
				} else {
					slOutput, ok := v_result_object.GetOutput().([]interface{})
					if ok {
						slData, r_err := json.Marshal(slOutput)
						if r_err != nil {
							r_ret = -1
							uniledgerlog.Error(r_err)
							return r_ret, r_err
						}
						gt.GetContract().SetOutputStruct(string(slData))
						uniledgerlog.Debug("%s after transfer asset output is %s", str_name, string(slData))
					} else {
						r_ret = -1
						r_err = fmt.Errorf("v_result_object.GetOutput().([]interface{}) assert error")
						uniledgerlog.Error(r_err.Error())
						return r_ret, r_err
					}
				}
			}

			//3 执行结果判断
			if v_result_object.GetCode() != 200 {
				r_err = fmt.Errorf("%s execute failed, return code is (%d), return message is (%s)",
					str_name, v_result_object.GetCode(), v_result_object.GetMessage())
				uniledgerlog.Error(r_err.Error())
				exec_flag = false
				break
			}
		}

		//执行失败，返回 -1
		if !exec_flag {
			r_ret = -1
			r_buf.WriteString("[Result]: Task execute fail;")
			r_buf.WriteString("start fail....")
			uniledgerlog.Error(r_buf.String())
			return r_ret, r_err
		}
		r_buf.WriteString("[Result]: Task execute success;")
		uniledgerlog.Info(r_buf.String(), " Dormant to Inprocess....")
		gt.SetState(constdef.TaskState[constdef.TaskState_In_Progress])
	} else if gt.IsDormant() && !gt.testPreCondition() { //未达到执行条件，返回 0
		r_ret = 0
		r_buf.WriteString("[Result]: preCondition not true;")
		uniledgerlog.Warn(r_buf.String(), "start exit....")
		return r_ret, r_err
	}

	//执行完动作后需要等待执行完成
	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), begin to complete]",
		uniledgerlog.NO_ERROR, gt.GetContract().GetContractId(), gt.GetName(), gt.GetTaskId()))
	r_ret, r_err = gt.Complete()
	return r_ret, r_err
}

func (gt *GeneralTask) Complete() (int8, error) {
	var r_ret int8
	var r_err error
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Task Process Runing:Inprogress State.")
	r_buf.WriteString("[ContractID]: " + gt.GetContract().GetContractId() + ";")
	r_buf.WriteString("[ContractHashID]: " + gt.GetContract().GetId() + ";")
	r_buf.WriteString("[TaskName]: " + gt.GetName() + ";")
	uniledgerlog.Info(r_buf.String(), "Complete begin....")

	//   任务执行成功，继续往下执行
	//   任务执行失败，该任务需要重新执行
	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), check complete precondition]",
		uniledgerlog.NO_ERROR, gt.GetContract().GetContractId(), gt.GetName(), gt.GetTaskId()))
	if gt.IsInProgress() && gt.testCompleteCondition() {
		//Dormant方法中生成交易产出Output（针对资产方法，合约执行状态+交易产出）；如果没有交易产出Output，则在Complete中生成Output（纯合约执行状态）
		//1 判断OutputStruct 是否为空，为空则需要在此构造产出结构体
		var output_null_flag bool = false
		if len(gt.GetContract().GetOutputStruct()) == 0 {
			output_null_flag = true
			var tmp_output common.OperateResult = common.OperateResult{}
			str_json_contract, r_err := gt.GetContract().Serialize()
			if r_err != nil || len(str_json_contract) == 0 {
				r_ret = -1
				r_buf.WriteString("[Result]: Generate OutputStruct fail, str_json_contract Serialize fail;")
				if r_err != nil {
					r_buf.WriteString("[Error]: " + r_err.Error() + ";")
				}
				if len(str_json_contract) == 0 {
					r_buf.WriteString("[Error]: constract json string is null;")
				}
				r_buf.WriteString("Complete fail...")
				uniledgerlog.Error(r_buf.String())
				return r_ret, r_err
			}

			if !gRPCClient.On {
				tmp_output, r_err = function.FuncInterim(str_json_contract,
					gt.GetContract().GetContractId(),
					gt.GetTaskId(),
					gt.GetTaskExecuteIdx())
			} else {
				var func_params map[string]interface{}
				func_params = make(map[string]interface{})
				func_params["Param01"] = str_json_contract
				func_params["Param02"] = ""
				func_params["Param03"] = gt.GetContract().GetContractId()
				func_params["Param04"] = gt.GetTaskId()
				func_params["Param05"] = gt.GetTaskExecuteIdx()
				func_params["Param06"] = ""

				slData, _ := json.Marshal(func_params)
				hostname, _ := os.Hostname()
				tmp_output, r_err = gRPCClient.FunctionRun(hostname+"|"+common0.GenTimestamp(),
					"FuncInterim", string(slData))
			}
			if r_err != nil {
				r_ret = -1
				r_buf.WriteString("[Result]: Generate OutputStruct fail, FuncInterim generage output error;")
				r_buf.WriteString("[Error]: " + r_err.Error() + ";")
				r_buf.WriteString("Complete fail....")
				uniledgerlog.Error(r_buf.String())
				return r_ret, r_err
			}

			if tmp_output.GetOutput() == nil {
				r_ret = -1
				r_buf.WriteString("[Result]: Generate OutputStruct fail,FuncInterim generage output is nil;")
				r_buf.WriteString("[Error]: outputStruct is nil;")
				r_buf.WriteString("Complete fail....")
				uniledgerlog.Error(r_buf.String())
				return r_ret, r_err
			}

			if !gRPCClient.On {
				strOutput, ok := tmp_output.GetOutput().(string)
				if ok {
					gt.GetContract().SetOutputStruct(strOutput)
					uniledgerlog.Debug("Task %s after transfer asset output is %s", gt.GetName(), strOutput)
				} else {
					r_ret = -1
					r_err = fmt.Errorf("tmp_output.GetOutput().(string) assert error")
					uniledgerlog.Error(r_err.Error())
					return r_ret, r_err
				}

			} else {
				slOutput, ok := tmp_output.GetOutput().([]interface{})
				if ok {
					slData, r_err := json.Marshal(slOutput)
					if r_err != nil {
						r_ret = -1
						uniledgerlog.Error(r_err)
						return r_ret, r_err
					}
					gt.GetContract().SetOutputStruct(string(slData))
					uniledgerlog.Debug("Task %s after transfer asset output is %s", gt.GetName(), string(slData))
				} else {
					r_ret = -1
					r_err = fmt.Errorf("tmp_output.GetOutput().([]interface{}) assert error")
					uniledgerlog.Error(r_err.Error())
					return r_ret, r_err
				}
			}
		}

		//4 OutputStruct插入到Output表中
		uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), complete execute]",
			uniledgerlog.NO_ERROR, gt.GetContract().GetContractId(), gt.GetName(), gt.GetTaskId()))
		var v_result common.OperateResult = common.OperateResult{}
		if output_null_flag {
			if !gRPCClient.On {
				v_result, r_err = function.FuncInterimComplete(gt.GetContract().GetOutputStruct(),
					constdef.TaskState[constdef.TaskState_Completed],
					gt.GetContract().GetContractState())
			} else {
				var func_params map[string]interface{}
				func_params = make(map[string]interface{})
				var interf []interface{}
				r_err = json.Unmarshal([]byte(gt.GetContract().GetOutputStruct()), &interf)
				if r_err != nil {
					r_ret = -1
					uniledgerlog.Error(r_err)
					return r_ret, r_err
				}
				func_params["Param01"] = interf
				func_params["Param02"] = constdef.TaskState[constdef.TaskState_Completed]
				func_params["Param03"] = gt.GetContract().GetContractState()

				slData, _ := json.Marshal(func_params)
				hostname, _ := os.Hostname()
				v_result, r_err = gRPCClient.FunctionRun(hostname+"|"+common0.GenTimestamp(),
					"FuncInterimComplete", string(slData))
			}
		} else {
			if !gRPCClient.On {
				// TODO : 目前仅支持一个datalist的情况
				var output_string string
				for v_key := range gt.GetDataList() {
					slData, r_err := json.Marshal(gt.GetContract().GetComponentTtem(v_key))
					if r_err != nil {
						r_ret = -1
						uniledgerlog.Error(r_err)
						return r_ret, r_err
					}
					var m_datalist map[string]interface{}
					m_datalist = make(map[string]interface{})
					if r_err = json.Unmarshal([]byte(slData), &m_datalist); r_err != nil {
						r_ret = -1
						uniledgerlog.Error(r_err)
						return r_ret, r_err
					}
					var m_destination map[string]interface{}
					m_destination = make(map[string]interface{})
					if err := json.Unmarshal([]byte(gt.GetContract().GetOutputStruct()), &m_destination); err != nil {
						uniledgerlog.Error(err)
					}
					//uniledgerlog.Debug("=== m_destination ===", m_destination)
					//uniledgerlog.Debug("=== m_datalist ===", m_datalist)
					//uniledgerlog.Debug("=== v_key ===", v_key)
					_InsertDataListToOutput(m_destination, m_datalist, v_key)
					slData, r_err = json.Marshal(m_destination)
					if r_err != nil {
						r_ret = -1
						uniledgerlog.Error(r_err)
						return r_ret, r_err
					}
					output_string = string(slData)
				}
				gt.GetContract().SetOutputStruct(output_string)
				v_result, r_err = function.FuncTransferAssetComplete(output_string, constdef.TaskState[constdef.TaskState_Completed])
			} else {
				var func_params map[string]interface{}
				func_params = make(map[string]interface{})
				var interf []interface{}
				if r_err = json.Unmarshal([]byte(gt.GetContract().GetOutputStruct()), &interf); r_err != nil {
					r_ret = -1
					uniledgerlog.Error(r_err)
					return r_ret, r_err
				}

				var datalistCname string
				for i := 0; i < len(interf); i++ {
					if interf[i].([]interface{})[1] == "ContractExecute" || len(interf) == 1 {
						// TODO : 目前仅支持一个datalist的情况
						for v_key := range gt.GetDataList() {
							slData, r_err := json.Marshal(gt.GetContract().GetComponentTtem(v_key))
							if r_err != nil {
								r_ret = -1
								uniledgerlog.Error(r_err)
								return r_ret, r_err
							}
							var m_datalist map[string]interface{}
							m_datalist = make(map[string]interface{})
							if r_err = json.Unmarshal([]byte(slData), &m_datalist); r_err != nil {
								r_ret = -1
								uniledgerlog.Error(r_err)
								return r_ret, r_err
							}
							var m_destination map[string]interface{}
							m_destination = make(map[string]interface{})
							if r_err = json.Unmarshal([]byte(interf[i].([]interface{})[2].(string)), &m_destination); r_err != nil {
								r_ret = -1
								uniledgerlog.Error(r_err)
								return r_ret, r_err
							}
							uniledgerlog.Debug("=== m_destination ===", m_destination)
							uniledgerlog.Debug("=== m_datalist ===", m_datalist)
							uniledgerlog.Debug("=== v_key ===", v_key)
							datalistCname = v_key
							_InsertDataListToOutput(m_destination, m_datalist, v_key)
							slData, r_err = json.Marshal(m_destination)
							if r_err != nil {
								r_ret = -1
								uniledgerlog.Error(r_err)
								return r_ret, r_err
							}
							//uniledgerlog.Debug("-------------------------------------")
							//uniledgerlog.Debug(interf[i].([]interface{})[2])
							interf[i].([]interface{})[2] = string(slData)
							//uniledgerlog.Debug("-------------------------------------")
							//uniledgerlog.Debug(interf[i].([]interface{})[2])
							//uniledgerlog.Debug("-------------------------------------")
						}
					}
				}

				//uniledgerlog.Debug("------------------------------------------------------------")
				//uniledgerlog.Debug("interf : ", interf)
				//uniledgerlog.Debug("------------------------------------------------------------")

				func_params["Param01"] = interf
				func_params["Param02"] = constdef.TaskState[constdef.TaskState_Completed]
				func_params["Param03"] = datalistCname

				slData, _ := json.Marshal(func_params)
				hostname, _ := os.Hostname()
				v_result, r_err = gRPCClient.FunctionRun(hostname+"|"+common0.GenTimestamp(),
					"FuncAllTransferComplete", string(slData))
			}
		}
		//执行结果判断
		if r_err != nil || v_result.GetCode() != 200 {
			r_ret = -1
			r_buf.WriteString("[Result]: Task execute fail;")
			if r_err != nil {
				r_buf.WriteString("[Error]: " + r_err.Error() + ";")
			}
			if v_result.GetCode() != 200 {
				r_buf.WriteString(fmt.Sprintf("%s execute failed, return code is (%d), return message is (%s) ;",
					gt.GetName(), v_result.GetCode(), v_result.GetMessage()))
			}
			r_buf.WriteString("Complete fail....")
			uniledgerlog.Error(r_buf.String())
			return r_ret, r_err
		}

		//5 设置OutputStruct的部分字段更新: OutputId  OutputTaskId, OutputTaskExecuteIdx, OutputStruct
		uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), update complete output]",
			uniledgerlog.NO_ERROR, gt.GetContract().GetContractId(), gt.GetName(), gt.GetTaskId()))
		if !gRPCClient.On {
			gt.GetContract().SetOutputStruct(v_result.GetOutput().(string))
			uniledgerlog.Debug("Task %s after complete operate output is %s", gt.GetName(), v_result.GetOutput().(string))
		} else {
			gt.GetContract().SetOutputStruct(v_result.GetOutput().(string))
			uniledgerlog.Debug("Task %s after complete operate output is %s", gt.GetName(), v_result.GetOutput().(string))

			//output, ok := v_result.GetOutput().([]interface{})
			//if ok {
			//	str_output, ok := output[2].(string)
			//	if !ok {
			//		r_ret = -1
			//		uniledgerlog.Error("utput[2].(string) assert error")
			//		return r_ret, r_err
			//
			//	}
			//	gt.GetContract().SetOutputStruct(str_output)
			//	uniledgerlog.Info("====after transfer operate==" + str_output)
			//} else {
			//	r_ret = -1
			//	uniledgerlog.Error("v_result.GetOutput().([]interface{}) assert error")
			//	return r_ret, r_err
			//}
		}

		var map_output_first interface{} = common.Deserialize(gt.GetContract().GetOutputStruct())
		if map_output_first == nil {
			r_ret = -1
			r_err = fmt.Errorf("Contract Output Deserialize fail!")
			r_buf.WriteString("[Result]: CompleteCondition not true;")
			r_buf.WriteString("[Error]: " + r_err.Error() + ";")
			uniledgerlog.Error(r_buf.String(), "Complete exit....")
			return r_ret, r_err
		}

		var map_output_second map[string]interface{} = map_output_first.(map[string]interface{})
		if map_output_second == nil || len(map_output_second) == 0 || map_output_second["transaction"] == nil {
			r_ret = -1
			r_err = fmt.Errorf("Contract Output Struct Get fail!")
			r_buf.WriteString("[Result]: CompleteCondition not true;")
			r_buf.WriteString("[Error]: " + r_err.Error() + ";")
			uniledgerlog.Error(r_buf.String(), "Complete exit....")
			return r_ret, r_err
		}

		var map_transaction map[string]interface{} = map_output_second["transaction"].(map[string]interface{})
		if map_transaction["Contract"] == nil {
			r_ret = -1
			r_err = fmt.Errorf("Contract HashId Get fail!")
			r_buf.WriteString("[Result]: CompleteCondition not true;")
			r_buf.WriteString("[Error]: " + r_err.Error() + ";")
			uniledgerlog.Error(r_buf.String(), "Complete exit....")
			return r_ret, r_err
		}

		var map_contract map[string]interface{} = map_transaction["Contract"].(map[string]interface{})
		gt.GetContract().SetOutputId(map_contract["id"].(string))
		gt.GetContract().SetOutputTaskId(gt.GetTaskId())
		gt.GetContract().SetOutputTaskExecuteIdx(gt.GetTaskExecuteIdx())
		uniledgerlog.Debug("gt.GetContract().GetOutputId():", gt.GetContract().GetOutputId())

		gt.SetState(constdef.TaskState[constdef.TaskState_Completed])
		uniledgerlog.Info(r_buf.String(), " Inprocess to Complete....")
	} else if gt.IsInProgress() && !gt.testCompleteCondition() {
		r_ret = 0
		r_buf.WriteString("[Result]: CompleteCondition not true;")
		uniledgerlog.Warn(r_buf.String(), "Complete exit....")
		return r_ret, r_err
	}

	//保证顺利执行，给执行方法留下执行时间，需要多次sleep等待执行
	var executeEngineConf map[interface{}]interface{}
	if engine.UCVMConf == nil {
		executeEngineConf = make(map[interface{}]interface{}, 0)
		executeEngineConf["task_complete_sleep_count"] = 3
		executeEngineConf["task_complete_sleep_time"] = 5
	} else {
		executeEngineConf = engine.UCVMConf["ExecuteEngine"].(map[interface{}]interface{})
	}

	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), begin to discard]",
		uniledgerlog.NO_ERROR, gt.GetContract().GetContractId(), gt.GetName(), gt.GetTaskId()))
	var v_sleep_num, v_sleep_time int
	var ok bool
	if v_sleep_num, ok = executeEngineConf["task_complete_sleep_count"].(int); !ok {
		v_sleep_num = 3
	}
	if v_sleep_time, ok = executeEngineConf["task_complete_sleep_time"].(int); !ok {
		v_sleep_time = 5
	}
	time.Sleep(time.Second * time.Duration(v_sleep_time))
	for v_sleep_num > 0 {
		v_sleep_num--
		r_ret, r_err = gt.Discard()
		if r_ret == 0 {
			time.Sleep(time.Second * time.Duration(v_sleep_time))
		} else {
			break
		}
	}

	return r_ret, r_err
}

func (gt *GeneralTask) Discard() (int8, error) {
	var r_ret int8
	var r_err error
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Task Process Runing:Complete State.")
	r_buf.WriteString("[ContractID]: " + gt.GetContract().GetContractId() + ";")
	r_buf.WriteString("[ContractHashID]: " + gt.GetContract().GetId() + ";")
	r_buf.WriteString("[TaskName]: " + gt.GetName() + ";")
	uniledgerlog.Info(r_buf.String(), "Discard begin....")

	// DiscardCondition 需要包含多节点共识结果标识 (默认条件)
	//   任务执行结果共识通过后，继续往下执行；
	//   任务执行结果共识不通过，该任务需要重新执行；
	if gt.IsCompleted() && gt.testDiscardCondition() {
		//DiscardCondition中需要默认添加任务执行结果入链判断条件
		uniledgerlog.Info(r_buf.String(), " Complete to Discard....")
		gt.SetState(constdef.TaskState[constdef.TaskState_Discard])
		r_ret = 1
	}

	return r_ret, r_err
}

//任务运行后进行的后处理
func (gt *GeneralTask) PostProcess(p_flag int8, nBrotherNum int) error {
	var r_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}

	//获取当前合约HashID(contract.Id），新建合约HashID(contract.outputId)
	v_contract := gt.GetContract()
	r_buf.WriteString("Contract Runing:PostProcess.")
	r_buf.WriteString("[ContractID]: " + v_contract.GetContractId() + ";")
	r_buf.WriteString("[ContractHashID]: " + v_contract.GetId() + ";")
	r_buf.WriteString("[TaskName]: " + gt.GetName() + ";")

	uniledgerlog.Debug("-----------------------------------------------")
	uniledgerlog.Debug("ContractId:" + v_contract.GetContractId())
	uniledgerlog.Debug("Id:" + v_contract.GetId())
	uniledgerlog.Debug("TaskId:" + gt.GetTaskId())
	uniledgerlog.Debug("State:" + gt.GetState())
	uniledgerlog.Debug("TaskExecuteIdx:%d", gt.GetTaskExecuteIdx())
	uniledgerlog.Debug("-----------------------------------------------")

	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), update task schedule state]",
		uniledgerlog.NO_ERROR, gt.GetContract().GetContractId(), gt.GetName(), gt.GetTaskId()))
	switch p_flag {
	case -1:
		//执行失败：1.更新contractID1 的flag=0, failNum+1, timestamp
		//    调用扫描引擎接口： UpdateMonitorFail(contractID_old)
		var failStruct common.UpdateMonitorFailStruct
		failStruct.FstrContractID = v_contract.GetContractId()
		failStruct.FstrContractHashID = v_contract.GetId()
		failStruct.FstrTaskId = gt.GetTaskId()
		failStruct.FstrTaskState = gt.GetState()
		failStruct.FnTaskExecuteIndex = gt.GetTaskExecuteIdx()
		slFailData, _ := json.Marshal(failStruct)
		r_err = common.UpdateMonitorFail(string(slFailData))
		if r_err != nil {
			r_buf.WriteString("[Result]: PostProcess[UpdateMonitorFail] Fail;")
			r_buf.WriteString("[Error]: " + r_err.Error() + ";")
			uniledgerlog.Error(r_buf.String())
		} else {
			r_buf.WriteString("[Result]: PostProcess[UpdateMonitorFail] Succ;")
			uniledgerlog.Info(r_buf.String())
		}
	case 0:
		//执行条件不满足：
		//    case1: State=Dormant or Inprocess .更新contractID1 的flag=0，waitNum+1, timestamp
		//    case2: State=Complete 更新 contractID1 的flag=1,successNum+1, timestamp; 添加 contractID2 的记录 flag=0
		//    调用扫描引擎接口： UpdateMonitorWait(contractID_old)
		if nBrotherNum == 0 { // 只有当是同级task中最后一个时，才会更新状态；否则等待其他同级task的执行。
			if gt.GetState() == constdef.TaskState[constdef.TaskState_Dormant] ||
				gt.GetState() == constdef.TaskState[constdef.TaskState_In_Progress] {
				var waitStruct common.UpdateMonitorWaitStruct
				waitStruct.WstrContractID = v_contract.GetContractId()
				waitStruct.WstrContractHashID = v_contract.GetId()
				waitStruct.WstrTaskId = gt.GetTaskId()
				waitStruct.WstrTaskState = gt.GetState()
				waitStruct.WnTaskExecuteIndex = gt.GetTaskExecuteIdx()
				slWaitData, _ := json.Marshal(waitStruct)
				r_err = common.UpdateMonitorWait(string(slWaitData))
				if r_err != nil {
					r_buf.WriteString("[Result]: PostProcess[UpdateMonitorWait] Fail;")
					r_buf.WriteString("[Error]: " + r_err.Error() + ";")
					uniledgerlog.Error(r_buf.String())
				} else {
					r_buf.WriteString("[Result]: PostProcess[UpdateMonitorWait] Succ;")
					uniledgerlog.Info(r_buf.String())
				}
			} else if gt.GetState() == constdef.TaskState[constdef.TaskState_Completed] {
				r_buf.WriteString("[ContractHashID_new]: " + v_contract.GetOutputId() + ";")
				var succStruct common.UpdateMonitorSuccStruct
				succStruct.SstrContractID = v_contract.GetContractId()
				succStruct.SstrContractHashIdOld = v_contract.GetId()
				succStruct.SstrTaskStateOld = gt.GetState()
				succStruct.SstrTaskIdOld = v_contract.GetOrgTaskId()
				succStruct.SnTaskExecuteIndexOld = v_contract.GetOrgTaskExecuteIdx()
				succStruct.SstrContractHashIDNew = v_contract.GetOutputId()
				succStruct.SstrTaskIdNew = v_contract.GetOutputTaskId()
				succStruct.SstrTaskStateNew = gt.GetState()
				succStruct.SnTaskExecuteIndexNew = v_contract.GetOutputTaskExecuteIdx()
				succStruct.SnFlag = 0
				slSuccData, _ := json.Marshal(succStruct)
				r_err = common.UpdateMonitorSucc(string(slSuccData))
				if r_err != nil {
					r_buf.WriteString("[Result]: PostProcess[0][UpdateMonitorSucc] Fail;")
					r_buf.WriteString("[Error]: " + r_err.Error() + ";")
					uniledgerlog.Error(r_buf.String())
				} else {
					r_buf.WriteString("[Result]: PostProcess[0][UpdateMonitorSucc] Succ;")
					uniledgerlog.Info(r_buf.String())
				}
			}
		}
	case 1:
		//执行成功：1 更新contractID1 的flag=1, succNum+1, timestamp, 2.将contractID2插入到扫描监控表中 flag=1
		//    调用扫描引擎接口： UpdateMonitorSucc(contractID_old, contractID_new)
		r_buf.WriteString("[ContractHashID_new]: " + v_contract.GetOutputId() + ";")
		var succStruct common.UpdateMonitorSuccStruct
		succStruct.SstrContractID = v_contract.GetContractId()
		succStruct.SstrContractHashIdOld = v_contract.GetId()
		succStruct.SstrTaskStateOld = gt.GetState()
		succStruct.SstrTaskIdOld = v_contract.GetOrgTaskId()
		succStruct.SnTaskExecuteIndexOld = v_contract.GetOrgTaskExecuteIdx()
		succStruct.SstrContractHashIDNew = v_contract.GetOutputId()
		succStruct.SstrTaskIdNew = v_contract.GetOutputTaskId()
		succStruct.SstrTaskStateNew = gt.GetState()
		succStruct.SnTaskExecuteIndexNew = v_contract.GetOutputTaskExecuteIdx()
		succStruct.SnFlag = 0
		slSuccData, _ := json.Marshal(succStruct)
		r_err = common.UpdateMonitorSucc(string(slSuccData))
		if r_err != nil {
			r_buf.WriteString("[Result]: PostProcess[1][UpdateMonitorSucc] Fail;")
			r_buf.WriteString("[Error]: " + r_err.Error() + ";")
			uniledgerlog.Error(r_buf.String())
		} else {
			r_buf.WriteString("[Result]: PostProcess[1][UpdateMonitorSucc] Succ;")
			uniledgerlog.Info(r_buf.String())
		}
	}
	return r_err
}

//由于查询分支结果的不确定性，使用分支条件赋予预估值，使得多节点 不同时运行结果一致性
//   通过 Cname进行对应function和data， expression_function_A \ data_int_expression_function_A
func (gt *GeneralTask) ConsistentValue(p_dataList map[string]inf.IData, p_name string, p_result common.OperateResult) {
	var r_buf bytes.Buffer = bytes.Buffer{}
	var v_data inf.IData
	//TODO :临时处理,把DataValueSetter 和 Data的名称保持一样的规则
	if len(p_dataList) == 0 {
		return
	}
	for v_key := range p_dataList {
		if strings.Contains(v_key, p_name) {
			v_data = p_dataList[v_key]
			v_data.SetValue(p_result.GetData())
			gt.GetContract().UpdateComponentRunningState(constdef.ComponentType[constdef.Component_Data], v_data.GetName(), v_data)
		}
	}

	if v_data == nil {
		uniledgerlog.Error("function and datalist Cname are different !!")
		return
	}

	switch gt.GetCtype() {
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Enquiry]:
		// 根据函数执行结果和分支情况决定最终的结果值
		select_branchs := gt.GetSelectBranches()
		if len(select_branchs) != 0 {
			for v_idx := range select_branchs {
				select_object := select_branchs[v_idx]
				select_value, select_err := gt.GetContract().EvaluateExpression(constdef.ExpressionType[constdef.Expression_Condition], select_object.GetBranchExpressionStr())
				if select_err != nil {
					r_buf.WriteString("[Result]: ConsistentValue fail;")
					r_buf.WriteString("[ContractId]: " + gt.GetContract().GetContractId() + ";")
					r_buf.WriteString("[ConstractHashId]: " + gt.GetContract().GetOutputId() + ";")
					r_buf.WriteString("[Error]: " + select_err.Error() + ";")
					uniledgerlog.Error(r_buf.String())
					break
				}
				if select_value.(bool) {
					select_final_value, select_final_err := gt.GetContract().EvaluateExpression("", select_object.GetBranchExpressionValue().(string))
					if select_final_err != nil {
						r_buf.WriteString("[Result]: ConsistentValue fail;")
						r_buf.WriteString("[ContractId]: " + gt.GetContract().GetContractId() + ";")
						r_buf.WriteString("[ConstractHashId]: " + gt.GetContract().GetOutputId() + ";")
						r_buf.WriteString("[Error]: " + select_final_err.Error() + ";")
						uniledgerlog.Error(r_buf.String())
						break
					}
					v_data.SetValue(select_final_value)
					break
				}
			}
		} else {
			v_data.SetValue(p_result.GetData())

		}
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Action]:
		v_data.SetValue(p_result.GetData())
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Decision]:
		v_data.SetValue(p_result.GetData())
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Plan]:
		v_data.SetValue(p_result.GetData())
	default:
		v_data.SetValue(p_result.GetData())
	}
	gt.GetContract().UpdateComponentRunningState(constdef.ComponentType[constdef.Component_Data], v_data.GetName(), v_data)
}

func _InsertDataListToOutput(map_destination, map_datalist map[string]interface{}, datalist_cname string) {
	length := len(map_destination["transaction"].(map[string]interface{})["Contract"].(map[string]interface{})["ContractBody"].(map[string]interface{})["ContractComponents"].([]interface{}))
	for i := 0; i < length; i++ {
		if map_destination["transaction"].(map[string]interface{})["Contract"].(map[string]interface{})["ContractBody"].(map[string]interface{})["ContractComponents"].([]interface{})[i].(map[string]interface{})["DataList"] != nil {
			datalist_len := len(map_destination["transaction"].(map[string]interface{})["Contract"].(map[string]interface{})["ContractBody"].(map[string]interface{})["ContractComponents"].([]interface{})[i].(map[string]interface{})["DataList"].([]interface{}))
			if datalist_len > 0 {
				// TODO : 目前仅支持一个datalist的情况
				cname := map_destination["transaction"].(map[string]interface{})["Contract"].(map[string]interface{})["ContractBody"].(map[string]interface{})["ContractComponents"].([]interface{})[i].(map[string]interface{})["DataList"].([]interface{})[0].(map[string]interface{})["Cname"].(string)
				if cname == datalist_cname {
					map_destination["transaction"].(map[string]interface{})["Contract"].(map[string]interface{})["ContractBody"].(map[string]interface{})["ContractComponents"].([]interface{})[i].(map[string]interface{})["DataList"].([]interface{})[0] = map_datalist
				}
			}
		}
	}
}
