package task

import (
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/constdef"
)
//表示场景：1. 执行一个动作 2. 查询一个数据，但不会根据结果分支
type Action struct {
	Enquiry
}

func NewAction()*Action{
	action := &Action{}
	return action
}
//===============接口实现===================
func (a Action)SetContract(p_contract inf.ICognitiveContract){
	a.Enquiry.SetContract(p_contract)
}

func (a Action)GetContract() inf.ICognitiveContract{
	return a.Enquiry.GetContract()
}

func (a Action)GetName()string {
	return a.Enquiry.GetCname()
}
func (a Action)GetCtype()string{
	return a.Enquiry.GetCtype()
}

func (a Action)GetDescription()string {
	return a.Enquiry.GetDescription()
}

func (a Action)GetState()string {
	return a.Enquiry.GetState()
}

func (a Action)SetState(p_state string){
	a.Enquiry.SetState(p_state)
}

func (a Action)GetNextTasks() []string{
	return a.Enquiry.GetNextTasks()
}

func (a Action)UpdateState()(int8, error){
	return a.Enquiry.UpdateState()
}
//===============描述态=====================


//===============运行态=====================
func (a *Action)InitAction() error{
	var err error = nil
	err = a.InitEnquriy()
	if err != nil {
		//TODO log
		return err
	}
	a.SetCtype(constdef.ComponentType[constdef.Component_Task] +  "." + constdef.TaskType[constdef.Task_Action])
	return err
}