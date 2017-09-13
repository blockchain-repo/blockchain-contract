package task

import (
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
)

//表示场景：1. 执行一个动作 2. 查询一个数据，但不会根据结果分支
type Action struct {
	Enquiry
}

func NewAction() *Action {
	action := &Action{}
	return action
}

//===============接口实现===================
func (a Action) SetContract(p_contract inf.ICognitiveContract) {
	a.Enquiry.SetContract(p_contract)
}

func (a Action) GetContract() inf.ICognitiveContract {
	return a.Enquiry.GetContract()
}

func (a Action) GetName() string {
	return a.Enquiry.GetCname()
}
func (a Action) GetCtype() string {
	return a.Enquiry.GetCtype()
}

func (a Action) GetDescription() string {
	return a.Enquiry.GetDescription()
}

func (a Action) GetState() string {
	return a.Enquiry.GetState()
}

func (a Action) SetState(p_state string) {
	a.Enquiry.SetState(p_state)
}

func (a Action) GetNextTasks() []string {
	return a.Enquiry.GetNextTasks()
}

func (a Action) UpdateState(nBrotherNum int) (int8, error) {
	return a.Enquiry.UpdateState(nBrotherNum)
}
func (a Action) GetTaskId() string {
	return a.Enquiry.GetTaskId()
}

func (a Action) GetTaskExecuteIdx() int {
	return a.Enquiry.GetTaskExecuteIdx()
}

func (a Action) SetTaskId(str_taskId string) {
	a.Enquiry.SetTaskId(str_taskId)
}

func (a Action) SetTaskExecuteIdx(int_idx int) {
	a.Enquiry.SetTaskExecuteIdx(int_idx)
}

func (a Action) CleanValueInProcess() {
	a.Enquiry.CleanValueInProcess()
}

func (a Action) UpdateStaticState() (interface{}, error) {
	return a.GeneralTask.UpdateStaticState()
}

//===============描述态=====================

//===============运行态=====================
func (a *Action) InitAction() error {
	var err error = nil
	err = a.InitEnquriy()
	if err != nil {
		uniledgerlog.Error("InitAction fail[" + err.Error() + "]")
		return err
	}
	a.SetCtype(constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Action])
	return err
}
