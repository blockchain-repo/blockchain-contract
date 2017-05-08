package task

import (
	"unicontract/src/core/engine/execengine/property"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/constdef"
	"fmt"
)

type Decision struct {
	Enquiry
	CandidateList []DecisionCandidate `json:"CandidateList"`
	DecisionResult []DecisionCandidate `json:"DecisionResult"`
}

const (
	_CandidateList = "_CandidateList"
	_DecisionResult = "_DecisionResult"
)

func NewDecision()*Decision{
	decision := &Decision{}
	return decision
}
//====================接口方法========================
func (d Decision)SetContract(p_contract inf.ICognitiveContract){
	d.Enquiry.SetContract(p_contract)
}

func (d Decision)GetContract() inf.ICognitiveContract{
	return d.Enquiry.GetContract()
}

func (d Decision)GetName()string {
	return d.Enquiry.GetName()
}

func (d Decision)GetCtype()string{
	return d.Enquiry.GetCtype()
}

func (d Decision)GetDescription()string {
	return d.Enquiry.GetDescription()
}

func (d Decision)GetState()string {
	return d.Enquiry.GetState()
}

func (d Decision)SetState(p_state string){
	d.Enquiry.SetState(p_state)
}

func (d Decision)GetNextTasks() []string{
	return d.Enquiry.GetNextTasks()
}

func (d Decision)UpdateState()(int8, error){
	var r_ret int8 = 0
	var r_err error = nil
	switch d.GetState(){
	case constdef.TaskState[constdef.TaskState_Dormant]:
		r_ret,r_err = d.Start()
		//正常执行，转入下一任务
		if r_ret == 1 {
			// TODO 调用公用方法，执行后继任务
		}else if r_ret == -1{ //轮询等待后，执行失败，则进行回滚
			//TODO 回滚
		}else if r_ret == 0 { //轮询等待后，条件不成立，则暂时退出
			//TODO log打印
		}
	case constdef.TaskState[constdef.TaskState_In_Progress]:
		r_ret,r_err = d.Complete()
		//正常执行，转入下一任务
		if r_ret == 1 {
			// TODO 调用公用方法，执行后继任务
		}else if r_ret == -1{ //轮询等待后，执行失败，则进行回滚
			//TODO 回滚
		}else if r_ret == 0 {//轮询等待后，条件不成立，则暂时退出
			//TODO log打印
		}
	case constdef.TaskState[constdef.TaskState_Completed]:
		r_ret,r_err = d.Disgard()
		//正常执行，转入下一任务
		if r_ret == 1 {
			// TODO 调用公用方法，执行后继任务
		}else { //轮询等待后，条件不成立或执行失败，则暂时退出
			//TODO log打印
		}
	}
	return r_ret,r_err
}

//====================描述态==========================


//====================运行态==========================
func (d *Decision) AddProperty(object interface{}, str_name string, value interface{})property.PropertyT {
	var pro_object property.PropertyT
	if value == nil {
		pro_object = *property.NewPropertyT(str_name)
		d.PropertyTable[str_name] = pro_object
		return pro_object
	}
	switch value.(type) {
	case map[string]DecisionCandidate:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(map[string]DecisionCandidate))
		d.PropertyTable[str_name] = pro_object
	}
	return pro_object
}

func (d *Decision)InitDecision() error{
	var err error = nil
	err = d.InitEnquriy()
	if err != nil {
		//TODO log
		return err
	}
	d.SetCtype(constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Decision])
	//condidatelist arrar to map
	if d.CandidateList == nil {
		d.CandidateList = make([]DecisionCandidate, 0)
	}
	map_candidatelist := make(map[string]DecisionCandidate, 0)
	for _,p_cand := range d.CandidateList {
		map_candidatelist[p_cand.GetName()] = p_cand
	}
	d.AddProperty(d, _CandidateList, map_candidatelist)
	if map_candidatelist == nil {
		fmt.Println("1111111111111111")
	}
	if d.PropertyTable[_CandidateList] == nil {
		fmt.Println("2222222222222")
	}
    fmt.Println(d.PropertyTable[_CandidateList])
	//decisionresult arrar to map
	if d.DecisionResult == nil {
		d.DecisionResult = make([]DecisionCandidate, 0)
	}
	map_decisionResult := make(map[string]DecisionCandidate, 0)
	for _,p_result := range d.DecisionResult{
		map_decisionResult[p_result.GetName()] = p_result
	}
	d.AddProperty(d,  _DecisionResult, map_decisionResult)
	return err
}

//====属性Get方法
//TODO： map本身是无序的，不需排序
func (d *Decision) GetCandidateList() map[string]DecisionCandidate{
	candlist_property := d.PropertyTable[_CandidateList].(property.PropertyT)
	return candlist_property.GetValue().(map[string]DecisionCandidate)
}

func (d *Decision) GetDecisionResult()map[string]DecisionCandidate{
	resultlist_property := d.PropertyTable[_DecisionResult].(property.PropertyT)
	return resultlist_property.GetValue().(map[string]DecisionCandidate)
}

func (d *Decision)GetCandidate(p_name string) DecisionCandidate{
	candlist_property := d.PropertyTable[_CandidateList].(property.PropertyT)
	if candlist_property.GetValue() != nil {
		map_candlist := candlist_property.GetValue().(map[string]DecisionCandidate)
		return map_candlist[p_name]
	}
	return DecisionCandidate{}
}
//====动态添加方法
func (d *Decision) AddCandidate(p_candidate interface{}){
	if p_candidate != nil {
		candlist_property := d.PropertyTable[_CandidateList].(property.PropertyT)
		if candlist_property.GetValue() == nil {
			candlist_property.SetValue(make(map[string]DecisionCandidate, 0))
		}
		v_candidate := p_candidate.(DecisionCandidate)
        if d.GetContract() != nil {
			v_candidate.SetContract(d.GetContract())
		}
		map_candlist := candlist_property.GetValue().(map[string]DecisionCandidate)
		map_candlist[v_candidate.GetCname()] = v_candidate
		candlist_property.SetValue(map_candlist)
		d.PropertyTable[_CandidateList] = candlist_property
	}
}

func (d *Decision) RemoveCandidate(p_candidate interface{}){
	if p_candidate != nil {
		candlist_property := d.PropertyTable[_CandidateList].(property.PropertyT)
		if candlist_property.GetValue() != nil {
			v_candidate := p_candidate.(DecisionCandidate)
			map_candlist := candlist_property.GetValue().(map[string]DecisionCandidate)
			delete(map_candlist, v_candidate.GetCname())
			candlist_property.SetValue(map_candlist)
			d.PropertyTable[_CandidateList] = candlist_property
		}
	}
}

func (d *Decision) evaluateCandidate(){
	candlist_property := d.PropertyTable[_CandidateList].(property.PropertyT)
	if candlist_property.GetValue() != nil {
		for _,v_value := range candlist_property.GetValue().(map[string]DecisionCandidate) {
			 v_value.SetContract(d.GetContract())
			if v_value.Eval() > 0 {
				d.AddDecisionResult(v_value)
			}
		}
	}
}

func (d *Decision)ResetDecisionResult(){
	resultlist_property := d.PropertyTable[_DecisionResult].(property.PropertyT)
	if resultlist_property.GetValue() != nil {
		resultlist_property.SetValue(make(map[string]DecisionCandidate, 0))
	}
	d.PropertyTable[_DecisionResult] = resultlist_property
}

func (d *Decision) AddDecisionResult(p_cand DecisionCandidate){
	resultlist_property := d.PropertyTable[_DecisionResult].(property.PropertyT)
	if resultlist_property.GetValue() == nil {
		resultlist_property.SetValue(make(map[string]DecisionCandidate, 0))
	}
	map_resultlist := resultlist_property.GetValue().(map[string]DecisionCandidate)
	map_resultlist[p_cand.GetCname()] = p_cand
	resultlist_property.SetValue(map_resultlist)
	d.PropertyTable[_DecisionResult] = resultlist_property
}

func (d *Decision)RemoveDecisionResult(p_cand []DecisionCandidate){
	resultlist_property := d.PropertyTable[_DecisionResult].(property.PropertyT)
	if p_cand != nil {
		map_resultlist := resultlist_property.GetValue().(map[string]DecisionCandidate)
		for _,v_cand := range p_cand {
			delete(map_resultlist, v_cand.GetCname())
		}
		resultlist_property.SetValue(map_resultlist)
		d.PropertyTable[_DecisionResult] = resultlist_property
	}
}