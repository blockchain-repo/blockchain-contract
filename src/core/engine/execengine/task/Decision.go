package task

import (
	"bytes"
	"fmt"

	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
)

type Decision struct {
	Enquiry
	CandidateList []DecisionCandidate `json:"CandidateList"` //决策结果都在CandidateList中体现；每个决策候选集，一个决策结果
}

const (
	_CandidateList = "_CandidateList"
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

//当task为Decision时，重写UpdateState(nBrotherNum int) (int8, error)方法
//当前任务生命周期的执行：（根据任务状态选择相应的执行态方法进入）
//入口时机：加载中的任务执行完成Discard,执行下一可执行任务Dormant
//执行过程：PreProcess => Start or Complete or Discard => PostProcess
//执行结果：
//    1. ret = -1：执行失败, 需要回滚
//    2. ret = 0 ：执行条件未达到
//    3. ret = 1 ：执行完成,转入后继任务
func (d Decision) UpdateState(nBrotherNum int) (int8, error) {
	var r_ret int8 = 0
	var r_err error = nil
	var r_str_error string = ""
	var r_flag int8 = -1
	if &d == nil {
		r_ret = -1
		r_err = fmt.Errorf("Object pointer is null!")
		return r_ret, r_err
	}

	//预处理
	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), begin to preprocess]",
		uniledgerlog.NO_ERROR, d.GetContract().GetContractId(), d.GetName(), d.GetTaskId()))
	r_err = d.PreProcess()
	if r_err != nil {
		uniledgerlog.Error("Task[" + d.GetName() + "] PreProcess fail!")
		return r_ret, r_err
	}

	//处理中
	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), begin to execute]",
		uniledgerlog.NO_ERROR, d.GetContract().GetContractId(), d.GetName(), d.GetTaskId()))
	r_ret, process_err := d.Start()
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
		uniledgerlog.NO_ERROR, d.GetContract().GetContractId(), d.GetName(), d.GetTaskId()))
	postProcess_err := d.PostProcess(r_flag, nBrotherNum)
	if postProcess_err != nil {
		r_str_error = r_str_error + "[PostProcess_Error]" + postProcess_err.Error()
	}
	if r_str_error != "" {
		r_err = fmt.Errorf(r_str_error)
	}
	return r_ret, r_err
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
	d.ResetCandidate()
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
		uniledgerlog.Error("InitDecision fail[" + err.Error() + "]")
		return err
	}

	d.SetCtype(constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Decision])

	//condidatelist arrar to map
	if d.CandidateList == nil {
		d.CandidateList = make([]DecisionCandidate, 0)
	}

	//map_candidatelist := make(map[string]DecisionCandidate, 0)
	map_candidatelist := make(map[string]interface{}, 0) // 实质上是 map[string]DecisionCandidate
	for _, p_cand := range d.CandidateList {
		p_cand.InitDecisionCandidate()
		map_candidatelist[p_cand.GetName()] = p_cand
	}

	common.AddProperty(d, d.PropertyTable, _CandidateList, map_candidatelist)
	return err
}

//====属性Get方法
//TODO： map本身是无序的，不需排序
func (d *Decision) GetCandidateList() map[string]DecisionCandidate {
	candlist_property, ok := d.PropertyTable[_CandidateList].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	candlist_value, ok := candlist_property.GetValue().(map[string]DecisionCandidate)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	return candlist_value
}

func (d *Decision) GetCandidate(p_name string) DecisionCandidate {
	candlist_property, ok := d.PropertyTable[_CandidateList].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return DecisionCandidate{}
	}
	if candlist_property.GetValue() != nil {
		candlist_map, ok := candlist_property.GetValue().(map[string]DecisionCandidate)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return DecisionCandidate{}
		}
		candlist_value, ok := candlist_map[p_name]
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.NULL_ERROR, ""))
			return DecisionCandidate{}
		}
		return candlist_value
	}
	return DecisionCandidate{}
}

//====动态添加方法
func (d *Decision) AddCandidate(p_candidate interface{}) {
	if p_candidate != nil {
		candlist_property, ok := d.PropertyTable[_CandidateList].(property.PropertyT)
		if !ok {
			candlist_property = *property.NewPropertyT(_CandidateList)
		}
		if candlist_property.GetValue() == nil {
			candlist_property.SetValue(make(map[string]DecisionCandidate, 0))
		}
		v_candidate, ok := p_candidate.(DecisionCandidate)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return
		}
		if d.GetContract() != nil {
			v_candidate.SetContract(d.GetContract())
		}
		map_candlist, ok := candlist_property.GetValue().(map[string]DecisionCandidate)
		if !ok {
			map_candlist = make(map[string]DecisionCandidate, 0)
		}
		map_candlist[v_candidate.GetCname()] = v_candidate
		candlist_property.SetValue(map_candlist)
		d.PropertyTable[_CandidateList] = candlist_property
	}
}

func (d *Decision) RemoveCandidate(p_candidate interface{}) {
	if p_candidate != nil {
		candlist_property, ok := d.PropertyTable[_CandidateList].(property.PropertyT)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return
		}
		if candlist_property.GetValue() != nil {
			v_candidate, ok := p_candidate.(DecisionCandidate)
			if !ok {
				uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
				return
			}
			map_candlist, ok := candlist_property.GetValue().(map[string]DecisionCandidate)
			if !ok {
				uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
				return
			}
			delete(map_candlist, v_candidate.GetCname())
			candlist_property.SetValue(map_candlist)
			d.PropertyTable[_CandidateList] = candlist_property
		}
	}
}

func (d *Decision) evaluateCandidate() error {
	var err error

	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), Decision get CandidateList]",
		uniledgerlog.NO_ERROR, d.GetContract().GetContractId(), d.GetName(), d.GetTaskId()))
	candlist_property, ok := d.PropertyTable[_CandidateList].(property.PropertyT)
	if !ok {
		err = fmt.Errorf("[%s][%s]", uniledgerlog.ASSERT_ERROR, "d.PropertyTable[_CandidateList].(property.PropertyT)")
		uniledgerlog.Error(err.Error())
		return err
	}
	if candlist_property.GetValue() != nil {
		//candlist_map, ok := candlist_property.GetValue().(map[string]DecisionCandidate)
		candlist_map, ok := candlist_property.GetValue().(map[string]interface{})
		if !ok {
			err = fmt.Errorf("[%s][%s]", uniledgerlog.ASSERT_ERROR, "candlist_property.GetValue().(map[string]DecisionCandidate)")
			uniledgerlog.Error(err.Error())
			return err
		}

		uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), Decision begin to evaluate]",
			uniledgerlog.NO_ERROR, d.GetContract().GetContractId(), d.GetName(), d.GetTaskId()))
		for v_key, v_value := range candlist_map {
			tmp_DecisionCandidate, ok := v_value.(DecisionCandidate)
			if !ok {
				err = fmt.Errorf("[%s][%s]", uniledgerlog.ASSERT_ERROR, "candlist_property.GetValue().(map[string]DecisionCandidate)")
				uniledgerlog.Error(err.Error())
				return err
			}
			uniledgerlog.Error("%+v", tmp_DecisionCandidate)
			tmp_DecisionCandidate.SetContract(d.GetContract())
			tmp_DecisionCandidate.ResetDecisionCandidate()
			err := tmp_DecisionCandidate.Eval()
			if err != nil {
				uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.EXECUTE_ERROR, ""))
				return err
			}
			candlist_map[v_key] = tmp_DecisionCandidate
		}
		candlist_property.SetValue(candlist_map)
		d.PropertyTable[_CandidateList] = candlist_property
	}
	return err
}

func (d *Decision) ResetCandidate() {
	resultlist_property, ok := d.PropertyTable[_CandidateList].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	if resultlist_property.GetValue() != nil {
		resultlist_property.SetValue(make(map[string]DecisionCandidate, 0))
	}
	d.PropertyTable[_CandidateList] = resultlist_property
}

//针对决策单独进行Start操作
func (d *Decision) Start() (int8, error) {
	var r_ret int8 = 0
	var r_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Task Process Runing:Dormant State.")
	r_buf.WriteString("[ContractID]: " + d.GetContract().GetContractId() + ";")
	r_buf.WriteString("[ContractHashID]: " + d.GetContract().GetId() + ";")
	r_buf.WriteString("[TaskName]: " + d.GetName() + ";")
	uniledgerlog.Info(r_buf.String(), "Start begin....")

	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), check start precondition]",
		uniledgerlog.NO_ERROR, d.GetContract().GetContractId(), d.GetName(), d.GetTaskId()))
	if d.IsDormant() && d.testPreCondition() {
		r_err = d.evaluateCandidate()
		//执行失败，返回 -1
		if r_err != nil {
			r_ret = -1
			r_buf.WriteString("[Result]: Task execute fail;")
			r_buf.WriteString("[Error]: " + r_err.Error() + ";")
			r_buf.WriteString("Start fail....")
			uniledgerlog.Error(r_buf.String())
			return r_ret, r_err
		}

		r_buf.WriteString("[Result]: Task execute success;")
		uniledgerlog.Info(r_buf.String(), " Dormant to Inprocess....")
		d.SetState(constdef.TaskState[constdef.TaskState_In_Progress])
	} else if d.IsDormant() && !d.testPreCondition() { //未达到执行条件，返回 0
		r_ret = 0
		r_buf.WriteString("[Result]: preCondition not true;")
		uniledgerlog.Warn(r_buf.String(), " exit....")
		return r_ret, r_err
	}

	//执行完动作后需要等待执行完成
	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s), task name is (%s), id is (%s), begin to complete]",
		uniledgerlog.NO_ERROR, d.GetContract().GetContractId(), d.GetName(), d.GetTaskId()))
	r_ret, r_err = d.Complete()
	return r_ret, r_err
}
