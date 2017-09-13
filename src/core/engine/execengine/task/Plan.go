package task

import (
	"fmt"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
)

type Plan struct {
	GeneralTask
	TaskList []string `json:"TaskList"` //TODO process map[string]Task
}

const (
	_TaskList = "_TaskList"
)

func NewPlan() *Plan {
	plan := &Plan{}
	return plan
}

//====================接口方法========================
func (p Plan) SetContract(p_contract inf.ICognitiveContract) {
	p.GeneralTask.SetContract(p_contract)
}

func (p Plan) GetContract() inf.ICognitiveContract {
	return p.GeneralTask.GetContract()
}

func (p Plan) GetName() string {
	return p.GeneralTask.GetCname()
}

func (p Plan) GetCtype() string {
	return p.GeneralTask.GetCtype()
}

func (p Plan) GetDescription() string {
	return p.GeneralTask.GetDescription()
}

func (p Plan) GetState() string {
	return p.GeneralTask.GetState()
}

func (p Plan) SetState(p_state string) {
	p.GeneralTask.SetState(p_state)
}

func (p Plan) GetNextTasks() []string {
	return p.GeneralTask.GetNextTasks()
}

func (p Plan) UpdateState(nBrotherNum int) (int8, error) {
	return p.GeneralTask.UpdateState(nBrotherNum)
}
func (p Plan) GetTaskId() string {
	return p.GeneralTask.GetTaskId()
}

func (p Plan) GetTaskExecuteIdx() int {
	return p.GeneralTask.GetTaskExecuteIdx()
}

func (p Plan) SetTaskId(str_taskId string) {
	p.GeneralTask.SetTaskId(str_taskId)
}

func (p Plan) SetTaskExecuteIdx(int_idx int) {
	p.GeneralTask.SetTaskExecuteIdx(int_idx)
}

func (p Plan) CleanValueInProcess() {
	p.GeneralTask.CleanValueInProcess()
}

func (p Plan) UpdateStaticState() (interface{}, error) {
	return p.GeneralTask.UpdateStaticState()
}

//====================描述态==========================
//TODO 反序列化任务数组

//====================运行态==========================
func (p *Plan) InitPlan() error {
	var err error = nil
	err = p.InitGeneralTask()
	if err != nil {
		uniledgerlog.Error("InitPlan fail[" + err.Error() + "]")
		return err
	}
	p.SetCtype(constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Plan])
	//TaskList array to map
	map_tasklist := make(map[string]inf.ITask, 0)
	if p.TaskList != nil && len(p.TaskList) > 0 {
		for _, p_task := range p.TaskList {
			//TODO read from array
			//map_tasklist[p_task.GetName()] = p_task
			fmt.Println(p_task)
		}
	}
	common.AddProperty(p, p.PropertyTable, _TaskList, map_tasklist)
	return err
}

func (p *Plan) GetTaskList() interface{} {
	tasklist_property := p.PropertyTable[_TaskList].(property.PropertyT)
	return tasklist_property.GetValue().(map[string]inf.ITask)
}

//property_table update
func (p *Plan) AddTask(p_task inf.ITask) {
	tasklist_property := p.PropertyTable[_TaskList].(property.PropertyT)
	if tasklist_property.GetValue() == nil {
		tasklist_property.SetValue(make(map[string]inf.ITask, 0))
	}
	tasklist_map := tasklist_property.GetValue().(map[string]inf.ITask)
	tasklist_map[p_task.GetName()] = p_task
	tasklist_property.SetValue(tasklist_map)
	p.PropertyTable[_TaskList] = tasklist_property
	//TODO contrat add component
	//p.GetContract().AddComponent(p_task.GetName(), p_task)
}

func (p *Plan) RemoveTask(p_taskname string) {
	tasklist_property := p.PropertyTable[_TaskList].(property.PropertyT)
	if tasklist_property.GetValue() != nil {
		tasklist_map := tasklist_property.GetValue().(map[string]inf.ITask)
		delete(tasklist_map, p_taskname)
		tasklist_property.SetValue(tasklist_map)
		p.PropertyTable[_TaskList] = tasklist_property
	}
}
