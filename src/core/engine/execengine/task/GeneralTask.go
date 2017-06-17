package task

//描述态：属性为数组
//运行态：属性为map
//		描述态 =》运行态： 在Init中进行转化
//		运行态 =》描述态： 在反序列化中进行转化
import (
	"bytes"
	"errors"
	"time"
	"unicontract/src/core/engine"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/component"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/expression"
	"unicontract/src/core/engine/execengine/function"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"

	"fmt"
	"github.com/astaxie/beego/logs"
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
	state_property := gt.PropertyTable[_State].(property.PropertyT)
	return state_property.GetValue().(string)
}

func (gt *GeneralTask) SetState(p_state string) {
	gt.State = p_state
	state_property := gt.PropertyTable[_State].(property.PropertyT)
	state_property.SetValue(p_state)
	gt.PropertyTable[_State] = state_property
}

func (gt *GeneralTask) GetNextTasks() []string {
	if gt.PropertyTable[_NextTasks] == nil {
		return nil
	}
	nexttask_property := gt.PropertyTable[_NextTasks].(property.PropertyT)
	return nexttask_property.GetValue().([]string)
}

//当前任务生命周期的执行：（根据任务状态选择相应的执行态方法进入）
//入口时机：加载中的任务执行完成Discard,执行下一可执行任务Dormant
//执行过程：PreProcess => Start or Complete or Discard => PostProcess
//执行结果：
//    1. ret = -1：执行失败, 需要回滚
//    2. ret = 0 ：执行条件未达到
//    3. ret = 1 ：执行完成,转入后继任务
func (gt GeneralTask) UpdateState() (int8, error) {
	var r_ret int8 = 0
	var r_err error = nil
	var r_str_error string = ""
	var r_flag int8 = -1
	if &gt == nil {
		r_ret = -1
		r_err = errors.New("Object pointer is null!")
		return r_ret, r_err
	}
	//预处理
	r_err = gt.PreProcess()
	if r_err != nil {
		logs.Error("Task[" + gt.GetName() + "] PreProcess fail!")
		return r_ret, r_err
	}
	//处理中
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
	postProcess_err := gt.PostProcess(r_flag)
	if postProcess_err != nil {
		r_str_error = r_str_error + "[PostProcess_Error]" + postProcess_err.Error()
	}
	if r_str_error != "" {
		r_err = errors.New(r_str_error)
	}
	return r_ret, r_err
}
func (gt *GeneralTask) GetTaskId() string {
	if gt.PropertyTable[_TaskId] == nil {
		return ""
	}
	taskid_property := gt.PropertyTable[_TaskId].(property.PropertyT)
	return taskid_property.GetValue().(string)
}

func (gt *GeneralTask) GetTaskExecuteIdx() int {
	taskexecuteidx_property := gt.PropertyTable[_TaskExecuteIdx].(property.PropertyT)
	return taskexecuteidx_property.GetValue().(int)
}

func (gt *GeneralTask) SetTaskId(str_taskId string) {
	//Take case: Setter method need set value for gc.xxxxxx
	gt.TaskId = str_taskId
	taskid_property := gt.PropertyTable[_TaskId].(property.PropertyT)
	taskid_property.SetValue(str_taskId)
	//Take case: Setter method need set value for gc.PropertyTable[xxxx]
	gt.PropertyTable[_TaskId] = taskid_property
}

func (gt *GeneralTask) SetTaskExecuteIdx(int_idx int) {
	//Take case: Setter method need set value for gc.xxxxxx
	gt.TaskExecuteIdx = int_idx
	taskexecuteidx_property := gt.PropertyTable[_TaskExecuteIdx].(property.PropertyT)
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
	for v_idx, v_data := range gt.DataList {
		if v_data == nil {
			err = fmt.Errorf("gt.DataList has nil data!")
			logs.Error("UpdateStaticState fail[" + err.Error() + "]")
			return nil, err
		}
		v_data_map := v_data.(map[string]interface{})
		new_data := gt.GetContract().GetData(v_data_map["Cname"].(string))
		new_data_array[v_idx] = new_data
	}
	gt.DataList = new_data_array

	//Expression组件(DataValueSetterExpressionList)信息 更新到描述态
	var new_expression_array []interface{} = make([]interface{}, len(gt.DataValueSetterExpressionList))
	for v_idx, v_expression := range gt.DataValueSetterExpressionList {
		if v_expression == nil {
			err = fmt.Errorf("gt.DataValueSetterExpressionList has nil data!")
			logs.Error("UpdateStaticState fail[" + err.Error() + "]")
			return nil, err
		}
		v_expression_map := v_expression.(map[string]interface{})
		new_expression := gt.GetContract().GetExpression(v_expression_map["Cname"].(string))
		new_expression_array[v_idx] = new_expression
	}
	gt.DataValueSetterExpressionList = new_expression_array
	//Expression组件(PreCondition)信息 更新到描述态
	var new_pre_array []interface{} = make([]interface{}, len(gt.PreCondition))
	for v_idx, v_pre := range gt.PreCondition {
		if v_pre == nil {
			err = fmt.Errorf("gt.PreCondition has nil data!")
			logs.Error("UpdateStaticState fail[" + err.Error() + "]")
			return nil, err
		}
		v_pre_map := v_pre.(map[string]interface{})
		new_pre := gt.GetContract().GetExpression(v_pre_map["Cname"].(string))
		new_pre_array[v_idx] = new_pre
	}
	gt.PreCondition = new_pre_array
	//Expression组件(CompleteCondition)信息 更新到描述态
	var new_complete_array []interface{} = make([]interface{}, len(gt.CompleteCondition))
	for v_idx, v_complete := range gt.CompleteCondition {
		if v_complete == nil {
			err = fmt.Errorf("gt.CompleteCondition has nil data!")
			logs.Error("UpdateStaticState fail[" + err.Error() + "]")
			return nil, err
		}
		v_complete_map := v_complete.(map[string]interface{})
		new_complete := gt.GetContract().GetExpression(v_complete_map["Cname"].(string))
		new_complete_array[v_idx] = new_complete
	}
	gt.CompleteCondition = new_complete_array
	//Expression组件(DiscardCondition)信息 更新到描述态
	var new_discard_array []interface{} = make([]interface{}, len(gt.DiscardCondition))
	for v_idx, v_discard := range gt.DiscardCondition {
		if v_discard == nil {
			err = fmt.Errorf("gt.DiscardCondition has nil data!")
			logs.Error("UpdateStaticState fail[" + err.Error() + "]")
			return nil, err
		}
		v_discard_map := v_discard.(map[string]interface{})
		new_discard := gt.GetContract().GetExpression(v_discard_map["Cname"].(string))
		new_discard_array[v_idx] = new_discard
	}
	gt.DiscardCondition = new_discard_array
	return gt, err
}

//===============运行态=====================
//Init中实现描述态 数组格式 到 map结构的转化
func (gt *GeneralTask) InitGeneralTask() error {
	var err error = nil
	err = gt.InitGeneralComponent()
	if err != nil {
		logs.Error("InitGeneralTask fail[" + err.Error() + "]")
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
		//TODO 转化
		if p_precondition != nil {
			switch p_precondition.(type) {
			case inf.IExpression:
			case *inf.IExpression:
				tmp_precondition := p_precondition.(inf.IExpression)
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
	precondition_property := gt.PropertyTable[_PreCondition].(property.PropertyT)
	return precondition_property.GetValue().(map[string]inf.IExpression)
}

func (gt *GeneralTask) GetCompleteCondition() map[string]inf.IExpression {
	if gt.PropertyTable[_CompleteCondition] == nil {
		return nil
	}
	completecondition_property := gt.PropertyTable[_CompleteCondition].(property.PropertyT)
	return completecondition_property.GetValue().(map[string]inf.IExpression)
}

func (gt *GeneralTask) GetDiscardCondition() map[string]inf.IExpression {
	if gt.PropertyTable[_DiscardCondition] == nil {
		return nil
	}
	Discardcondition_property := gt.PropertyTable[_DiscardCondition].(property.PropertyT)
	return Discardcondition_property.GetValue().(map[string]inf.IExpression)
}

func (gt *GeneralTask) GetDataList() map[string]inf.IData {
	datalist_property := gt.PropertyTable[_DataList].(property.PropertyT)
	return datalist_property.GetValue().(map[string]inf.IData)
}

func (gt *GeneralTask) GetDataValueSetterExpressionList() map[string]inf.IExpression {
	dataexpress_property := gt.PropertyTable[_DataValueSetterExpressionList].(property.PropertyT)
	return dataexpress_property.GetValue().(map[string]inf.IExpression)
}

func (gt *GeneralTask) GetSelectBranches() []common.SelectBranchExpression {
	selectbranch_property := gt.PropertyTable[_SelectBranches].(property.PropertyT)
	return selectbranch_property.GetValue().([]common.SelectBranchExpression)
}

//====属性动态初始化
func (gt *GeneralTask) ReSet() {
	gt.SetState(constdef.TaskState[constdef.TaskState_Dormant])
}

func (gt *GeneralTask) AddNextTasks(task string) {
	nexttask_property := gt.PropertyTable[_NextTasks].(property.PropertyT)
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
	precondition_property := gt.PropertyTable[_PreCondition].(property.PropertyT)
	if precondition_property.GetValue() == nil {
		precondition_property.SetValue(make(map[string]inf.IExpression, 0))
	}
	map_precondition := precondition_property.GetValue().(map[string]inf.IExpression)
	map_precondition[p_name] = expression.NewGeneralExpression(p_condition)

	precondition_property.SetValue(map_precondition)
	gt.PropertyTable[_PreCondition] = precondition_property
}

func (gt *GeneralTask) AddCompleteCondition(p_name string, p_condition string) {
	completecondition_property := gt.PropertyTable[_CompleteCondition].(property.PropertyT)
	if completecondition_property.GetValue() == nil {
		completecondition_property.SetValue(make(map[string]inf.IExpression, 0))
	}
	map_completecondition := completecondition_property.GetValue().(map[string]inf.IExpression)
	map_completecondition[p_name] = expression.NewGeneralExpression(p_condition)

	completecondition_property.SetValue(map_completecondition)
	gt.PropertyTable[_CompleteCondition] = completecondition_property
}

func (gt *GeneralTask) AddDiscardCondition(p_name string, p_condition string) {
	Discardcondition_property := gt.PropertyTable[_DiscardCondition].(property.PropertyT)
	if Discardcondition_property.GetValue() == nil {
		Discardcondition_property.SetValue(make([]inf.IExpression, 0))
	}
	map_Discardcondition := Discardcondition_property.GetValue().(map[string]inf.IExpression)
	map_Discardcondition[p_name] = expression.NewGeneralExpression(p_condition)

	Discardcondition_property.SetValue(map_Discardcondition)
	gt.PropertyTable[_DiscardCondition] = Discardcondition_property
}

func (gt *GeneralTask) AddDataSetterExpressionAndData(p_name string, p_dataSetterExpresstionStr string, p_data inf.IData) {
	gt.AddDataSetterExpression(p_name, p_dataSetterExpresstionStr)
	gt.AddData(p_data)
}

func (gt *GeneralTask) AddDataSetterExpression(p_name string, p_dataSetterExpresstionStr string) {
	if gt.PropertyTable[_DataValueSetterExpressionList] == nil {
		return
	}
	dataexpressionlist_property := gt.PropertyTable[_DataValueSetterExpressionList].(property.PropertyT)
	if dataexpressionlist_property.GetValue() == nil {
		map_dataexpressionlist := make(map[string]inf.IExpression, 0)
		dataexpressionlist_property.SetValue(map_dataexpressionlist)
	}
	if p_dataSetterExpresstionStr != "" {
		map_dataexpresslist := dataexpressionlist_property.GetValue().(map[string]inf.IExpression)
		map_dataexpresslist[p_name] = expression.NewGeneralExpression(p_dataSetterExpresstionStr)
		dataexpressionlist_property.SetValue(map_dataexpresslist)
		gt.PropertyTable[_DataValueSetterExpressionList] = dataexpressionlist_property
	}
}

func (gt *GeneralTask) AddData(p_data inf.IData) {
	if gt.PropertyTable[_DataList] == nil {
		return
	}
	datalist_property := gt.PropertyTable[_DataList].(property.PropertyT)
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
	dataExpression_property := gt.PropertyTable[_DataValueSetterExpressionList].(property.PropertyT)
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
	datalist_property := gt.PropertyTable[_DataList].(property.PropertyT)
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
	selectbranch_property := gt.PropertyTable[_SelectBranches].(property.PropertyT)
	selectbranch_property.SetValue(p_selectbranchs)
	gt.PropertyTable[_SelectBranches] = selectbranch_property
}

//TODO: 缺少Compounddata考虑
func (gt *GeneralTask) GetData(p_name string) (interface{}, error) {
	var err error = nil
	datalist_property := gt.PropertyTable[_DataList].(property.PropertyT)
	if datalist_property.GetValue() != nil {
		var data_map map[string]inf.IData = datalist_property.GetValue().(map[string]inf.IData)
		r_data, ok := data_map[p_name]
		if !ok {
			err = errors.New("Find data[" + p_name + "] Error!")
		}
		return r_data, err
	} else {
		err = errors.New("DataList is nil,find data[" + p_name + "] Error!")
		return nil, err
	}
}

func (gt *GeneralTask) GetDataExpression(p_name string) (interface{}, error) {
	var err error = nil
	dataexpressionlist_property := gt.PropertyTable[_DataValueSetterExpressionList].(property.PropertyT)
	if dataexpressionlist_property.GetValue() != nil {
		var dataexpression_map map[string]inf.IExpression = dataexpressionlist_property.GetValue().(map[string]inf.IExpression)
		r_data, ok := dataexpression_map[p_name]
		if !ok {
			err = errors.New("Find dataExpression[" + p_name + "] Error!")
		}
		return r_data, err
	} else {
		err = errors.New("DataValueSetterExpressionList is nil,find dataExpression[" + p_name + "] Error!")
		return nil, err
	}
}

//====运行条件判断
func (gt *GeneralTask) testCompleteCondition() bool {
	var r_flag bool = false
	if len(gt.GetPreCondition()) == 0 {
		r_flag = true
	}
	for _, value := range gt.GetPreCondition() {
		v_contract := gt.GetContract()
		v_bool, v_err := v_contract.EvaluateExpression(constdef.ExpressionType[constdef.Expression_Condition], value.GetExpressionStr())
		if v_err != nil {
			logs.Warning("CompleteCondition EvaluateExpression[" + value.GetExpressionStr() + " fail, [" + v_err.Error() + "]")
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
			logs.Warning("DiscardCondition EvaluateExpression[" + value.GetExpressionStr() + "] fail, [" + v_err.Error() + "]")
			r_flag = false
			break
		}
		if !v_bool.(bool) {
			r_flag = false
			break
		}
		r_flag = v_bool.(bool)
	}
	//默认的合约终止条件（当前合约步骤入链查询判定）
	v_contract := gt.GetContract()
	v_default_function := "FuncIsConPutInUnichian(\"" + v_contract.GetOutputId() + "\")"
	v_result, v_err := v_contract.EvaluateExpression(constdef.ExpressionType[constdef.Expression_Condition], v_default_function)
	if v_err != nil {
		logs.Warning("DiscardCondition EvaluateExpression[" + v_default_function + "] fail, [" + v_err.Error() + "]")
		r_flag = false
	}
	if !v_result.(bool) {
		logs.Warning("DiscardCondition EvaluateExpression[" + v_default_function + "] result is false!")
		r_flag = false
	}
	return r_flag
}

func (gt *GeneralTask) testPreCondition() bool {
	var r_flag bool = false
	if len(gt.GetPreCondition()) == 0 {
		r_flag = true
	}
	for _, value := range gt.GetPreCondition() {
		v_contract := gt.GetContract()
		v_bool, v_err := v_contract.EvaluateExpression(constdef.ExpressionType[constdef.Expression_Condition], value.GetExpressionStr())
		if v_err != nil {
			logs.Warning("PreCondition EvaluateExpression[" + value.GetExpressionStr() + " fail, [" + v_err.Error() + "]")
			r_flag = false
			break
		}
		if !v_bool.(bool) {
			r_flag = false
			break
		}
		r_flag = v_bool.(bool)
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
		r_err = errors.New("PreProcess fail, GetContract is nil!")
		r_buf.WriteString("PreProcess fail, GetContract is nil!")
		logs.Warning(r_buf.String())
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
	r_buf.WriteString("Contract Runing:Dormant State.")
	r_buf.WriteString("[ContractID]: " + gt.GetContract().GetContractId() + ";")
	r_buf.WriteString("[TaskName]: " + gt.GetName() + ";")
	logs.Info(r_buf.String(), " begin....")
	var r_ret int8 = 0
	var r_err error = nil
	if gt.IsInProgress() || gt.IsCompleted() {
		logs.Info("Task[", gt.GetName(), "] State[Start to Dormant]....")
		logs.Info(r_buf.String(), " InProgress|Completed to Dormant....")
		gt.SetState(constdef.TaskState[constdef.TaskState_Dormant])
		//TODO 回滚需求清空中间变量的值
	}
	return r_ret, r_err
}

func (gt *GeneralTask) Start() (int8, error) {
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Contract Runing:Dormant State.")
	r_buf.WriteString("[ContractID]: " + gt.GetContract().GetContractId() + ";")
	r_buf.WriteString("[TaskName]: " + gt.GetName() + ";")
	logs.Info(r_buf.String(), " begin....")
	var r_ret int8 = 0
	var r_err error = nil
	if gt.IsDormant() && gt.testPreCondition() {
		var exec_flag bool = true
		//var data_array []interface{} = gt.DataList
		//循环遍历函数表达式列表，执行函数
		//注意：限制只可有一个Output交易产出
		// TODO 待处理，避免一般操作任务，重复执行
		var v_idx int8 = 0
		for _, v_dataValueSetterExpression := range gt.GetDataValueSetterExpressionList() {
			v_expr_object := v_dataValueSetterExpression.(inf.IExpression)
			//1 函数识别 & 执行
			v_result, r_err := gt.GetContract().EvaluateExpression(constdef.ExpressionType[constdef.Expression_Function], v_expr_object.GetExpressionStr())
			v_result_object := v_result.(common.OperateResult)
			//2 执行结果赋值
			//  2.1 结果赋值到 data中,针对Enquiry Task，需要根据分支条件一致性化查询结果值
			gt.ConsistentValue(gt.DataList, v_idx, v_result_object)
			//  2.2 结果赋值到 dataSetterValue函数结果
			v_expr_object.SetExpressionResult(v_result_object)
			gt.GetContract().UpdateComponentRunningState(constdef.ComponentType[constdef.Component_Expression], v_expr_object.GetName(), v_result_object)
			//  2.3 Output交易产出结构体赋值
			if v_result_object.GetOutput() != "" {
				gt.GetContract().SetOutputStruct(v_result_object.GetOutput().(string))
			}
			//3 执行结果判断
			if r_err != nil || v_result_object.GetCode() != 200 {
				exec_flag = false
				break
			}
			v_idx = v_idx + 1
		}
		//执行失败，返回 -1
		if !exec_flag {
			r_ret = -1
			r_buf.WriteString("[Result]: Task execute fail;")
			r_buf.WriteString("[Error]: " + r_err.Error() + ";")
			r_buf.WriteString("fail....")
			logs.Error(r_buf.String())
			return r_ret, r_err
		}
		r_buf.WriteString("[Result]: Task execute success;")
		logs.Info(r_buf.String(), " Dormant to Inprocess....")
		gt.SetState(constdef.TaskState[constdef.TaskState_In_Progress])
	} else if gt.IsDormant() && !gt.testPreCondition() { //未达到执行条件，返回 0
		r_ret = 0
		r_buf.WriteString("[Result]: preCondition not true;")
		logs.Warning(r_buf.String(), " exit....")
		return r_ret, r_err
	}
	//执行完动作后需要等待执行完成
	var v_exit_flag int8 = 0
	for v_exit_flag == 0 {
		r_ret, r_err = gt.Complete()
		if r_ret == 0 {
			continue
		} else {
			break
		}
	}
	return r_ret, r_err
}

func (gt *GeneralTask) Complete() (int8, error) {
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Contract Runing:Inprogress State.")
	r_buf.WriteString("[ContractID]: " + gt.GetContract().GetContractId() + ";")
	r_buf.WriteString("[ContractHashID]: " + gt.GetContract().GetId() + ";")
	r_buf.WriteString("[TaskName]: " + gt.GetName() + ";")
	logs.Info(r_buf.String(), "Complete begin....")
	var r_ret int8 = 0
	var r_err error = nil
	// CompleteCondition 需要包含单节点任务执行结果条件
	//   任务执行成功，继续往下执行
	//   任务执行失败，该任务需要重新执行
	if gt.IsInProgress() && gt.testCompleteCondition() {
		//Dormant方法中生成交易产出Output（针对资产方法，合约执行状态+交易产出）；如果没有交易产出Output，则在Complete中生成Output（纯合约执行状态）
		//1 判断OutputStruct 是否为空，为空则需要在此构造产出结构体
		var output_null_flag bool = false
		if gt.GetContract().GetOutputStruct() == "" {
			output_null_flag = true
			var tmp_output common.OperateResult = common.OperateResult{}
			str_json_contract, r_err := gt.GetContract().Serialize()
			if r_err != nil || str_json_contract == "" {
				r_ret = -1
				r_buf.WriteString("[Result]: Generate OutputStruct fail, str_json_contract Serialize fail;")
				r_buf.WriteString("[Error]: " + r_err.Error() + ";")
				r_buf.WriteString("Complete fail....")
				logs.Error(r_buf.String())
				return r_ret, r_err
			}
			tmp_output, r_err = function.FuncInterim(str_json_contract, gt.GetContract().GetContractId(), gt.GetTaskId(), gt.GetTaskExecuteIdx())
			if r_err != nil {
				r_ret = -1
				r_buf.WriteString("[Result]: Generate OutputStruct fail, FuncInterim generage output error;")
				r_buf.WriteString("[Error]: " + r_err.Error() + ";")
				r_buf.WriteString("Complete fail....")
				logs.Error(r_buf.String())
				return r_ret, r_err
			}

			if tmp_output.GetOutput() == nil || tmp_output.GetOutput().(string) == "" {
				r_ret = -1
				r_buf.WriteString("[Result]: Generate OutputStruct fail,FuncInterim generage output is nil;")
				r_buf.WriteString("[Error]: outputStruct is nil;")
				r_buf.WriteString("Complete fail....")
				logs.Error(r_buf.String())
				return r_ret, r_err
			}
			gt.GetContract().SetOutputStruct(tmp_output.GetOutput().(string))
		}
		//4 OutputStruct插入到Output表中
		var v_result common.OperateResult = common.OperateResult{}
		if output_null_flag {
			v_result, r_err = function.FuncInterimComplete(gt.GetContract().GetOutputStruct(), constdef.TaskState[constdef.TaskState_Completed])
		} else {
			v_result, r_err = function.FuncTransferAssetComplete(gt.GetContract().GetOutputStruct(), constdef.TaskState[constdef.TaskState_Completed])
		}
		//执行结果判断
		if r_err != nil || v_result.GetCode() != 200 {
			r_ret = -1
			r_buf.WriteString("[Result]: Task execute fail;")
			r_buf.WriteString("[Error]: " + r_err.Error() + ";")
			r_buf.WriteString("Complete fail....")
			logs.Error(r_buf.String())
			return r_ret, r_err
		}
		//5 设置OutputStruct的部分字段更新: OutputId  OutputTaskId, OutputTaskExecuteIdx, OutputStruct
		gt.GetContract().SetOutputStruct(v_result.GetData().(string))
		var map_output_first interface{} = common.Deserialize(gt.GetContract().GetOutputStruct())
		if map_output_first == nil {
			r_ret = -1
			r_err = errors.New("Contract Output Deserialize fail!")
			r_buf.WriteString("[Result]: CompleteCondition not true;")
			r_buf.WriteString("[Error]: " + r_err.Error() + ";")
			logs.Warning(r_buf.String(), "Complete exit....")
			return r_ret, r_err
		}
		var map_output_second map[string]interface{} = map_output_first.(map[string]interface{})
		if map_output_second == nil || len(map_output_second) == 0 || map_output_second["transaction"] == nil {
			r_ret = -1
			r_err = errors.New("Contract Output Struct Get fail!")
			r_buf.WriteString("[Result]: CompleteCondition not true;")
			r_buf.WriteString("[Error]: " + r_err.Error() + ";")
			logs.Warning(r_buf.String(), "Complete exit....")
			return r_ret, r_err
		}
		var map_transaction map[string]interface{} = map_output_second["transaction"].(map[string]interface{})
		if map_transaction["Contract"] == nil {
			r_ret = -1
			r_err = errors.New("Contract HashId Get fail!")
			r_buf.WriteString("[Result]: CompleteCondition not true;")
			r_buf.WriteString("[Error]: " + r_err.Error() + ";")
			logs.Warning(r_buf.String(), "Complete exit....")
			return r_ret, r_err
		}
		var map_contract map[string]interface{} = map_transaction["Contract"].(map[string]interface{})
		gt.GetContract().SetOutputId(map_contract["id"].(string))
		gt.GetContract().SetOutputTaskId(gt.GetTaskId())
		gt.GetContract().SetOutputTaskExecuteIdx(gt.GetTaskExecuteIdx())
		//logs.Error("gt.GetContract():", gt.GetContract())
		logs.Error("gt.GetContract().GetOutputId():", gt.GetContract().GetOutputId())
		logs.Info(r_buf.String(), " Inprocess to Complete....")
		gt.SetState(constdef.TaskState[constdef.TaskState_Completed])
	} else if gt.IsInProgress() && !gt.testCompleteCondition() {
		r_ret = 0
		r_buf.WriteString("[Result]: CompleteCondition not true;")
		logs.Warning(r_buf.String(), "Complete exit....")
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

	var v_sleep_num int = executeEngineConf["task_complete_sleep_count"].(int)
	time.Sleep(time.Second * time.Duration(executeEngineConf["task_complete_sleep_time"].(int)))
	for v_sleep_num > 0 {
		v_sleep_num = v_sleep_num - 1
		r_ret, r_err = gt.Discard()
		if r_ret == 0 {
			time.Sleep(time.Second * time.Duration(executeEngineConf["task_complete_sleep_time"].(int)))
		} else {
			break
		}
	}
	return r_ret, r_err
}

func (gt *GeneralTask) Discard() (int8, error) {
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Contract Runing:Complete State.")
	r_buf.WriteString("[ContractID]: " + gt.GetContract().GetContractId() + ";")
	r_buf.WriteString("[ContractHashID]: " + gt.GetContract().GetId() + ";")
	r_buf.WriteString("[TaskName]: " + gt.GetName() + ";")
	logs.Info(r_buf.String(), " begin....")
	var r_ret int8 = 0
	var r_err error = nil
	// DiscardCondition 需要包含多节点共识结果标识 (默认条件)
	//   任务执行结果共识通过后，继续往下执行；
	//   任务执行结果共识不通过，该任务需要重新执行；
	if gt.IsCompleted() && gt.testDiscardCondition() {
		//DiscardCondition中需要默认添加任务执行结果入链判断条件
		logs.Info(r_buf.String(), " Complete to Discard....")
		gt.SetState(constdef.TaskState[constdef.TaskState_Discard])
		r_ret = 1
	}
	return r_ret, r_err
}

//任务运行后进行的后处理
func (gt *GeneralTask) PostProcess(p_flag int8) error {
	var r_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	//获取当前合约HashID(contract.Id），新建合约HashID(contract.outputId)
	v_contract := gt.GetContract()
	r_buf.WriteString("Contract Runing:PostProcess.")
	r_buf.WriteString("[ContractID]: " + v_contract.GetContractId() + ";")
	r_buf.WriteString("[ContractHashID]: " + v_contract.GetId() + ";")
	r_buf.WriteString("[TaskName]: " + gt.GetName() + ";")
	switch p_flag {
	case -1:
		//执行失败：1.更新contractID1 的flag=0, failNum+1, timestamp
		//    调用扫描引擎接口： UpdateMonitorFail(contractID_old)
		r_err = common.UpdateMonitorFail(v_contract.GetContractId(), v_contract.GetId(), gt.GetTaskId(), gt.GetState(), gt.GetTaskExecuteIdx())
		logs.Error("-----------------------------------------------")
		logs.Error("ContractId:" + v_contract.GetContractId())
		logs.Error("Id:" + v_contract.GetId())
		logs.Error("TaskId:" + gt.GetTaskId())
		logs.Error("State:" + gt.GetState())
		logs.Error("TaskExecuteIdx:%d", gt.GetTaskExecuteIdx())
		logs.Error("-----------------------------------------------")
		if r_err != nil {
			r_buf.WriteString("[Result]: PostProcess[UpdateMonitorFail] Fail;")
			r_buf.WriteString("[Error]: " + r_err.Error() + ";")
			logs.Warning(r_buf.String())
		} else {
			r_buf.WriteString("[Result]: PostProcess[UpdateMonitorFail] Succ;")
			logs.Info(r_buf.String())
		}
	case 0:
		//执行条件不满足：
		//    case1: State=Dormant or Inprocess .更新contractID1 的flag=0，waitNum+1, timestamp
		//    case2: State=Complete 更新 contractID1 的flag=1,successNum+1, timestamp; 添加 contractID2 的记录 flag=0
		//    调用扫描引擎接口： UpdateMonitorWait(contractID_old)
		if gt.GetState() == constdef.TaskState[constdef.TaskState_Dormant] || gt.GetState() == constdef.TaskState[constdef.TaskState_In_Progress] {
			r_err = common.UpdateMonitorWait(v_contract.GetContractId(), v_contract.GetId(), gt.GetTaskId(), gt.GetState(), gt.GetTaskExecuteIdx())
			if r_err != nil {
				r_buf.WriteString("[Result]: PostProcess[UpdateMonitorWait] Fail;")
				r_buf.WriteString("[Error]: " + r_err.Error() + ";")
				logs.Warning(r_buf.String())
			} else {
				r_buf.WriteString("[Result]: PostProcess[UpdateMonitorWait] Succ;")
				logs.Info(r_buf.String())
			}
		} else if gt.GetState() == constdef.TaskState[constdef.TaskState_Completed] {
			r_buf.WriteString("[ContractHashID_new]: " + v_contract.GetOutputId() + ";")
			r_err = common.UpdateMonitorSucc(
				v_contract.GetContractId(),
				v_contract.GetId(),
				gt.GetState(),
				v_contract.GetOrgTaskId(),
				v_contract.GetOrgTaskExecuteIdx(),
				v_contract.GetOutputId(),
				v_contract.GetOutputTaskId(),
				gt.GetState(),
				v_contract.GetOutputTaskExecuteIdx(),
				0,
			)
			if r_err != nil {
				r_buf.WriteString("[Result]: PostProcess[0][UpdateMonitorSucc] Fail;")
				r_buf.WriteString("[Error]: " + r_err.Error() + ";")
				logs.Warning(r_buf.String())
			} else {
				r_buf.WriteString("[Result]: PostProcess[0][UpdateMonitorSucc] Succ;")
				logs.Info(r_buf.String())
			}
		}
	case 1:
		//执行成功：1 更新contractID1 的flag=1, succNum+1, timestamp, 2.将contractID2插入到扫描监控表中 flag=1
		//    调用扫描引擎接口： UpdateMonitorSucc(contractID_old, contractID_new)
		r_buf.WriteString("[ContractHashID_new]: " + v_contract.GetOutputId() + ";")
		r_err = common.UpdateMonitorSucc(
			v_contract.GetContractId(),
			v_contract.GetId(),
			gt.GetState(),
			v_contract.GetOrgTaskId(),
			v_contract.GetOrgTaskExecuteIdx(),
			v_contract.GetOutputId(),
			v_contract.GetOutputTaskId(),
			gt.GetState(),
			v_contract.GetOutputTaskExecuteIdx(),
			0,
		)
		if r_err != nil {
			r_buf.WriteString("[Result]: PostProcess[1][UpdateMonitorSucc] Fail;")
			r_buf.WriteString("[Error]: " + r_err.Error() + ";")
			logs.Warning(r_buf.String())
		} else {
			r_buf.WriteString("[Result]: PostProcess[1][UpdateMonitorSucc] Succ;")
			logs.Info(r_buf.String())
		}
	}
	return r_err
}

//由于查询分支结果的不确定性，使用分支条件赋予预估值，使得多节点 不同时运行结果一致性
func (gt *GeneralTask) ConsistentValue(p_dataList []interface{}, p_idx int8, p_result common.OperateResult) {
	var r_buf bytes.Buffer = bytes.Buffer{}
	var v_data inf.IData
	switch gt.GetCtype() {
	case constdef.TaskType[constdef.Task_Enquiry]:
		// 根据函数执行结果和分支情况决定最终的结果值
		v_data := p_dataList[p_idx].(inf.IData)

		select_branchs := gt.GetSelectBranches()
		if len(select_branchs) != 0 {
			for _, select_expression := range select_branchs {
				select_object := select_expression
				select_value, select_err := gt.GetContract().EvaluateExpression(constdef.ExpressionType[constdef.Expression_Condition], select_object.GetBranchExpressionStr())
				if select_err != nil {
					r_buf.WriteString("[Result]: ConsistentValue fail;")
					r_buf.WriteString("[ContractId]: " + gt.GetContract().GetContractId() + ";")
					r_buf.WriteString("[ConstractHashId]: " + gt.GetContract().GetOutputId() + ";")
					r_buf.WriteString("[Error]: " + select_err.Error() + ";")
					logs.Error(r_buf.String())
					break
				}
				if select_value.(bool) {
					v_data.SetValue(select_object.GetBranchExpressionValue())
					break
				}
			}
		} else {
			v_data.SetValue(p_result.GetData())
		}
	case constdef.TaskType[constdef.Task_Action]:
		v_data := p_dataList[p_idx].(inf.IData)
		v_data.SetValue(p_result.GetData())
	case constdef.TaskType[constdef.Task_Decision]:
		v_data := p_dataList[p_idx].(inf.IData)
		v_data.SetValue(p_result.GetData())
	case constdef.TaskType[constdef.Task_Plan]:
		v_data := p_dataList[p_idx].(inf.IData)
		v_data.SetValue(p_result.GetData())
	default:
		v_data := p_dataList[p_idx].(inf.IData)
		v_data.SetValue(p_result.GetData())
	}
	gt.GetContract().UpdateComponentRunningState(constdef.ComponentType[constdef.Component_Data], v_data.GetName(), v_data)
}
