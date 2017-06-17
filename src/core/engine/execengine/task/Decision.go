package task

import (
	"github.com/astaxie/beego/logs"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
)

type Decision struct {
	Enquiry
	CandidateList  []DecisionCandidate `json:"CandidateList"`
	DecisionResult []DecisionCandidate `json:"DecisionResult"`
}

const (
	_CandidateList  = "_CandidateList"
	_DecisionResult = "_DecisionResult"
)

func NewDecision() *Decision {
	decision := &Decision{}
	return decision
}

//====================接口方法========================
func (d Decision) SetContract(p_contract inf.ICognitiveContract) {
	d.Enquiry.SetContract(p_contract)
}

func (d Decision) GetContract() inf.ICognitiveContract {
	return d.Enquiry.GetContract()
}

func (d Decision) GetName() string {
	return d.Enquiry.GetName()
}

func (d Decision) GetCtype() string {
	return d.Enquiry.GetCtype()
}

func (d Decision) GetDescription() string {
	return d.Enquiry.GetDescription()
}

func (d Decision) GetState() string {
	return d.Enquiry.GetState()
}

func (d Decision) SetState(p_state string) {
	d.Enquiry.SetState(p_state)
}

func (d Decision) GetNextTasks() []string {
	return d.Enquiry.GetNextTasks()
}

func (d Decision) UpdateState() (int8, error) {
	return d.Enquiry.UpdateState()
}
func (d Decision) GetTaskId() string {
	return d.Enquiry.GetTaskId()
}

func (d Decision) GetTaskExecuteIdx() int {
	return d.Enquiry.GetTaskExecuteIdx()
}

func (d Decision) SetTaskId(str_taskId string) {
	d.Enquiry.SetTaskId(str_taskId)
}

func (d Decision) SetTaskExecuteIdx(int_idx int) {
	d.Enquiry.SetTaskExecuteIdx(int_idx)
}

func (d Decision) CleanValueInProcess() {
	d.Enquiry.CleanValueInProcess()
	d.ResetDecisionResult()
}

func (d Decision) UpdateStaticState() (interface{}, error) {
	return d.Enquiry.UpdateStaticState()
}

//====================描述态==========================

//====================运行态==========================
func (d *Decision) AddProperty(object interface{}, str_name string, value interface{}) property.PropertyT {
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

func (d *Decision) InitDecision() error {
	var err error = nil
	err = d.InitEnquriy()
	if err != nil {
		logs.Error("InitDecision fail[" + err.Error() + "]")
		return err
	}
	d.SetCtype(constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Decision])
	//condidatelist arrar to map
	if d.CandidateList == nil {
		d.CandidateList = make([]DecisionCandidate, 0)
	}
	map_candidatelist := make(map[string]DecisionCandidate, 0)
	for _, p_cand := range d.CandidateList {
		map_candidatelist[p_cand.GetName()] = p_cand
	}
	common.AddProperty(d, d.PropertyTable, _CandidateList, map_candidatelist)
	//decisionresult arrar to map
	if d.DecisionResult == nil {
		d.DecisionResult = make([]DecisionCandidate, 0)
	}
	map_decisionResult := make(map[string]DecisionCandidate, 0)
	for _, p_result := range d.DecisionResult {
		map_decisionResult[p_result.GetName()] = p_result
	}
	common.AddProperty(d, d.PropertyTable, _DecisionResult, map_decisionResult)
	return err
}

//====属性Get方法
//TODO： map本身是无序的，不需排序
func (d *Decision) GetCandidateList() map[string]DecisionCandidate {
	candlist_property := d.PropertyTable[_CandidateList].(property.PropertyT)
	return candlist_property.GetValue().(map[string]DecisionCandidate)
}

func (d *Decision) GetDecisionResult() map[string]DecisionCandidate {
	resultlist_property := d.PropertyTable[_DecisionResult].(property.PropertyT)
	return resultlist_property.GetValue().(map[string]DecisionCandidate)
}

func (d *Decision) GetCandidate(p_name string) DecisionCandidate {
	candlist_property := d.PropertyTable[_CandidateList].(property.PropertyT)
	if candlist_property.GetValue() != nil {
		map_candlist := candlist_property.GetValue().(map[string]DecisionCandidate)
		return map_candlist[p_name]
	}
	return DecisionCandidate{}
}

//====动态添加方法
func (d *Decision) AddCandidate(p_candidate interface{}) {
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

func (d *Decision) RemoveCandidate(p_candidate interface{}) {
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

func (d *Decision) evaluateCandidate() {
	candlist_property := d.PropertyTable[_CandidateList].(property.PropertyT)
	if candlist_property.GetValue() != nil {
		for _, v_value := range candlist_property.GetValue().(map[string]DecisionCandidate) {
			v_value.SetContract(d.GetContract())
			if v_value.Eval() > 0 {
				d.AddDecisionResult(v_value)
			}
		}
	}
}

func (d *Decision) ResetDecisionResult() {
	resultlist_property := d.PropertyTable[_DecisionResult].(property.PropertyT)
	if resultlist_property.GetValue() != nil {
		resultlist_property.SetValue(make(map[string]DecisionCandidate, 0))
	}
	d.PropertyTable[_DecisionResult] = resultlist_property
}

func (d *Decision) AddDecisionResult(p_cand DecisionCandidate) {
	resultlist_property := d.PropertyTable[_DecisionResult].(property.PropertyT)
	if resultlist_property.GetValue() == nil {
		resultlist_property.SetValue(make(map[string]DecisionCandidate, 0))
	}
	map_resultlist := resultlist_property.GetValue().(map[string]DecisionCandidate)
	map_resultlist[p_cand.GetCname()] = p_cand
	resultlist_property.SetValue(map_resultlist)
	d.PropertyTable[_DecisionResult] = resultlist_property
}

func (d *Decision) RemoveDecisionResult(p_cand []DecisionCandidate) {
	resultlist_property := d.PropertyTable[_DecisionResult].(property.PropertyT)
	if p_cand != nil {
		map_resultlist := resultlist_property.GetValue().(map[string]DecisionCandidate)
		for _, v_cand := range p_cand {
			delete(map_resultlist, v_cand.GetCname())
		}
		resultlist_property.SetValue(map_resultlist)
		d.PropertyTable[_DecisionResult] = resultlist_property
	}
}
