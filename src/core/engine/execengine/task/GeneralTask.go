package task
//描述态：属性为数组
//运行态：属性为map
//		描述态 =》运行态： 在Init中进行转化
//		运行态 =》描述态： 在反序列化中进行转化
import (
	"unicontract/src/core/engine/execengine/expression"
	"unicontract/src/core/engine/execengine/component"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
	"unicontract/src/core/engine/execengine/constdef"
	"fmt"
	"time"
	"unicontract/src/core/engine/common"
	"errors"
	"bytes"
)

type GeneralTask struct {
	component.GeneralComponent
	State string `json:"State"`
	//type:inf.IExpression
	PreCondition []interface{} `json:"PreCondition"`
	//type:inf.IExpression
	CompleteCondition []interface{} `json:"CompleteCondition"`
	//type:inf.IExpression
	DisgardCondition []interface{} `json:"DisgardCondition"`
	//type:inf.IData
	DataList []interface{}  `json:"DataList"`
	//type:inf.IExpression
	DataValueSetterExpressionList []interface{} `json:"DataValueSetterExpressionList"`
	NextTasks []string `json:"NextTasks"`
}

const (
	_State = "_State"
	_PreCondition = "_PreCondition"
	_CompleteCondition = "_CompleteCondition"
	_DisgardCondition = "_DisgardCondition"
	_NextTasks = "_NextTasks"
	_DataList = "_DataList"
	_DataValueSetterExpressionList = "_DataValueSetterExpressionList"
)

func NewGeneralTask() *GeneralTask{
	v_task := &GeneralTask{}
	return v_task
}
//===============接口实现===================
func (gt GeneralTask)SetContract(p_contract inf.ICognitiveContract){
	gt.GeneralComponent.SetContract(p_contract)
}

func (gt GeneralTask)GetContract() inf.ICognitiveContract{
	return gt.GeneralComponent.GetContract()
}

func (gt GeneralTask)GetName()string {
	return gt.GeneralComponent.GetCname()
}
func (gt GeneralTask)GetCtype()string{
	return gt.GeneralComponent.GetCtype()
}

func (gt GeneralTask)GetDescription()string {
	return gt.GeneralComponent.GetDescription()
}

func (gt *GeneralTask) GetState()string{
	if gt.PropertyTable[_State] == nil {
		return ""
	}
	state_property := gt.PropertyTable[_State].(property.PropertyT)
	return state_property.GetValue().(string)
}

func (gt *GeneralTask) SetState(p_state string){
	gt.State = p_state
	state_property := gt.PropertyTable[_State].(property.PropertyT)
	state_property.SetValue(p_state)
	gt.PropertyTable[_State] = state_property
}

func (gt *GeneralTask) GetNextTasks() []string{
	if gt.PropertyTable[_NextTasks] == nil {
		return nil
	}
	nexttask_property := gt.PropertyTable[_NextTasks].(property.PropertyT)
	return nexttask_property.GetValue().([]string)
}

//当前任务生命周期的执行：(根据任务状态选择相应的执行态方法进入)
//执行过程：PreProcess => Start or Complete or Digcard => PostProcess
//执行结果：
//    1. ret = -1：执行失败, 需要回滚
//    2. ret = 0 ：执行条件未达到
//    3. ret = 1 ：执行完成,转入后继任务
func (gt GeneralTask)UpdateState()(int8, error){
	var r_ret int8 = 0
	var r_err error = nil
	var r_str_error string = ""
	var r_flag int8 = -1
	if &gt == nil {
		r_ret = -1
		r_err = errors.New("Object pointer is null!")
		return r_ret,r_err
	}
	//预处理
	r_err = gt.PreProcess()
	if r_err != nil {
		//TODO log
		return r_ret,r_err
	}
	//处理中
	r_ret,r_err = gt.Start()
	if r_err != nil {
		r_str_error = r_str_error + "[Run_Error]:" + r_err.Error()
	}
	switch r_ret {
	case 1:
		//正常执行，转入下一任务
		r_flag = 1
	case -1:
		//轮询等待后，执行失败，则进行回滚
		r_flag = -1
	case 0:
		//轮询等待后，条件不成立，则暂时退出
		r_flag = 0
	}
	//后处理
	r_err = gt.PostProcess(r_flag)
	if r_err != nil {
		r_str_error = r_str_error + "[PostProcess_Error]" + r_err.Error()
	}
	r_err = errors.New(r_str_error)
	return r_ret,r_err
}
//===============描述态=====================
//反序列化实现运行态 map结构 到 数组结构的转化



//===============运行态=====================
//Init中实现描述态 数组格式 到 map结构的转化
func (gt *GeneralTask) InitGeneralTask()error {
	var err error = nil
	err = gt.InitGeneralComponent()
	if err != nil {
		//TODO log
		return err
	}
	gt.Ctype = common.TernaryOperator(gt.Ctype == "", constdef.ComponentType[constdef.Component_Task], gt.Ctype).(string)
	gt.SetCtype(gt.Ctype)
	// State default
	gt.State = common.TernaryOperator(gt.State == "", constdef.ComponentType[constdef.TaskState_Dormant], gt.State).(string)
	gt.AddProperty(gt, _State, gt.State)

	//PreCondition array to map
	if gt.PreCondition == nil {
		gt.PreCondition = make([]interface{}, 0)
	}
	map_precondition := make(map[string]inf.IExpression, 0)
	for _,p_precondition := range gt.PreCondition {
		//TODO 转化
		if p_precondition != nil {
			switch p_precondition.(type) {
			case inf.IExpression:
			case *inf.IExpression:
				tmp_precondition := p_precondition.(inf.IExpression)
				map_precondition[tmp_precondition.GetExpressionStr()] = tmp_precondition
			}
		}
	}
	gt.AddProperty(gt, _PreCondition, map_precondition)
	//CompleteCondition arrat to map
	if gt.CompleteCondition == nil {
		gt.CompleteCondition = make([]interface{}, 0)
	}
	map_completecondition := make(map[string]inf.IExpression, 0)
	for _,p_completecondition := range gt.CompleteCondition {
		if p_completecondition != nil {
			switch p_completecondition.(type) {
			case inf.IExpression:
			case *inf.IExpression:
				tmp_completecondition := p_completecondition.(inf.IExpression)
				map_completecondition[tmp_completecondition.GetExpressionStr()] = tmp_completecondition
			}
		}
	}
	gt.AddProperty(gt ,_CompleteCondition, map_completecondition)
	//DisgardCondition arrat to map
	if gt.DisgardCondition == nil {
		gt.DisgardCondition = make([]interface{}, 0)
	}
	map_digardcondition := make(map[string]inf.IExpression, 0)
	for _,p_digardcondition := range gt.DisgardCondition {
		if p_digardcondition != nil {
			switch p_digardcondition.(type) {
			case inf.IExpression:
			case *inf.IExpression:
				tmp_digardcondition := p_digardcondition.(inf.IExpression)
				map_digardcondition[tmp_digardcondition.GetExpressionStr()] = tmp_digardcondition
			}
		}
	}
	gt.AddProperty(gt, _DisgardCondition, map_digardcondition)
	//DataList arr to map
	if gt.DataList == nil {
		gt.DataList = make([]interface{}, 0)
	}
	map_datalist := make(map[string]inf.IData, 0)
	for _,p_data := range gt.DataList {
		if p_data != nil {
			switch p_data.(type) {
			case inf.IData:
			case *inf.IData:
				tmp_data := p_data.(inf.IData)
				map_datalist[tmp_data.GetName()] = tmp_data
			}
		}
	}
	gt.AddProperty(gt, _DataList, map_datalist)
	//DataValueSetterExpressionList arr to map
	if gt.DataValueSetterExpressionList == nil {
		gt.DataValueSetterExpressionList = make([]interface{}, 0)
	}
	map_dataexpressionlist := make(map[string]inf.IExpression, 0)
	for _,p_express := range gt.DataValueSetterExpressionList {
		if p_express != nil {
			switch p_express.(type) {
			case inf.IExpression:
			case *inf.IExpression:
				tmp_express := p_express.(inf.IExpression)
				map_dataexpressionlist[tmp_express.GetExpressionStr()] = tmp_express
			}
		}
	}
	gt.AddProperty(gt, _DataValueSetterExpressionList, map_dataexpressionlist)
	//nextTask array to map
	if gt.NextTasks == nil {
		gt.NextTasks = make([]string, 0)
	}
	gt.AddProperty(gt, _NextTasks, gt.NextTasks)

	return err
}
//====属性Get方法
func (gt *GeneralTask) GetPreCondition()map[string]inf.IExpression{
	if gt.PropertyTable[_PreCondition] == nil {
		return nil
	}
	precondition_property := gt.PropertyTable[_PreCondition].(property.PropertyT)
	return precondition_property.GetValue().(map[string]inf.IExpression)
}

func (gt *GeneralTask) GetCompleteCondition()map[string]inf.IExpression{
	if gt.PropertyTable[_CompleteCondition] == nil {
		return nil
	}
	completecondition_property := gt.PropertyTable[_CompleteCondition].(property.PropertyT)
	return completecondition_property.GetValue().(map[string]inf.IExpression)
}

func (gt *GeneralTask) GetDisgardCondition()map[string]inf.IExpression{
	if gt.PropertyTable[_DisgardCondition] == nil {
		return nil
	}
	disgardcondition_property := gt.PropertyTable[_DisgardCondition].(property.PropertyT)
	return disgardcondition_property.GetValue().(map[string]inf.IExpression)
}

func (gt *GeneralTask)GetDataList() map[string]inf.IData{
	datalist_property := gt.PropertyTable[_DataList].(property.PropertyT)
	return datalist_property.GetValue().(map[string]inf.IData)
}

func (gt *GeneralTask)GetDataValueSetterExpressionList()map[string]inf.IExpression{
	dataexpress_property := gt.PropertyTable[_DataValueSetterExpressionList].(property.PropertyT)
	return dataexpress_property.GetValue().(map[string]inf.IExpression)
}

//====属性动态初始化
func (gt *GeneralTask)ReSet(){
	gt.SetState(constdef.TaskState[constdef.TaskState_Dormant])
}

func (gt *GeneralTask)AddNextTasks(task string){
	nexttask_property := gt.PropertyTable[_NextTasks].(property.PropertyT)
	if nexttask_property.GetValue() == nil{
		nexttask_property.SetValue(make([]string, 0))
	}
	if task != ""{
		arr_nexttasks := nexttask_property.GetValue().([]string)
		arr_nexttasks = append(arr_nexttasks, task)
		nexttask_property.SetValue(arr_nexttasks)
		gt.PropertyTable[_NextTasks] = nexttask_property
	}
}

func (gt *GeneralTask)AddPreCondition(p_condition string){
	precondition_property := gt.PropertyTable[_PreCondition].(property.PropertyT)
	if precondition_property.GetValue() == nil {
		precondition_property.SetValue(make(map[string]inf.IExpression, 0))
	}
	map_precondition := precondition_property.GetValue().(map[string]inf.IExpression)
	map_precondition[p_condition] = expression.NewGeneralExpression(p_condition)

	precondition_property.SetValue(map_precondition)
	gt.PropertyTable[_PreCondition] = precondition_property
}

func (gt *GeneralTask)AddCompleteCondition(p_condition string){
	completecondition_property := gt.PropertyTable[_CompleteCondition].(property.PropertyT)
	if completecondition_property.GetValue() == nil {
		completecondition_property.SetValue(make(map[string]inf.IExpression, 0))
	}
	map_completecondition := completecondition_property.GetValue().(map[string]inf.IExpression)
	map_completecondition[p_condition] = expression.NewGeneralExpression(p_condition)

	completecondition_property.SetValue(map_completecondition)
	gt.PropertyTable[_CompleteCondition] = completecondition_property
}

func (gt *GeneralTask)AddDisgardCondition(p_condition string){
	disgardcondition_property := gt.PropertyTable[_DisgardCondition].(property.PropertyT)
	if disgardcondition_property.GetValue() == nil {
		disgardcondition_property.SetValue(make([]inf.IExpression, 0))
	}
	map_disgardcondition := disgardcondition_property.GetValue().(map[string]inf.IExpression)
	map_disgardcondition[p_condition] = expression.NewGeneralExpression(p_condition)

	disgardcondition_property.SetValue(map_disgardcondition)
	gt.PropertyTable[_DisgardCondition] = disgardcondition_property
}
//====属性Set方法
//TODO: 缺少Compounddata考虑
func (gt *GeneralTask)GetData(p_name string)(interface{}, error){
	var err error = nil
	datalist_property := gt.PropertyTable[_DataList].(property.PropertyT)
	if datalist_property.GetValue() != nil {
		var data_map map[string]inf.IData = datalist_property.GetValue().(map[string]inf.IData)
		r_data,ok := data_map[p_name]
		if !ok{
			err = errors.New("Find data[" + p_name + "] Error!")
		}
		return r_data, err
	} else {
		err = errors.New("DataList is nil,find data[" + p_name + "] Error!")
		return nil,err
	}
}

//TODO: 缺少Compounddata考虑
func (gt *GeneralTask)GetDataExpression(p_name string)(interface{}, error){
	var err error = nil
	dataexpressionlist_property := gt.PropertyTable[_DataValueSetterExpressionList].(property.PropertyT)
	if dataexpressionlist_property.GetValue() != nil {
		var dataexpression_map map[string]inf.IExpression = dataexpressionlist_property.GetValue().(map[string]inf.IExpression)
		r_data,ok := dataexpression_map[p_name]
		if !ok{
			err = errors.New("Find dataExpression[" + p_name + "] Error!")
		}
		return r_data, err
	} else {
		err = errors.New("DataValueSetterExpressionList is nil,find dataExpression[" + p_name + "] Error!")
		return nil,err
	}
}

func (gt *GeneralTask)AddData(p_data inf.IData, p_dataSetterExpresstionStr string){
	if gt.PropertyTable[_DataList] == nil {
		return
	}
	datalist_property := gt.PropertyTable[_DataList].(property.PropertyT)
	dataexpressionlist_property := gt.PropertyTable[_DataValueSetterExpressionList].(property.PropertyT)
	if datalist_property.GetValue() == nil {
		map_datalist := make(map[string]inf.IData,0)
		datalist_property.SetValue(map_datalist)
		map_dataexpressionlist := make(map[string]inf.IExpression, 0)
		dataexpressionlist_property.SetValue(map_dataexpressionlist)
	}

	map_datalist := datalist_property.GetValue().(map[string]inf.IData)
	map_datalist[p_data.GetName()] = p_data
	datalist_property.SetValue(map_datalist)
	gt.PropertyTable[_DataList] = datalist_property
	//TODO: contract.component_table add component
	if p_dataSetterExpresstionStr != "" {
		map_dataexpresslist := dataexpressionlist_property.GetValue().(map[string]inf.IExpression)
		map_dataexpresslist[p_data.GetName()] = expression.NewGeneralExpression(p_dataSetterExpresstionStr)
		dataexpressionlist_property.SetValue(map_dataexpresslist)
		gt.PropertyTable[_DataValueSetterExpressionList] = dataexpressionlist_property
	}
}

func (gt *GeneralTask)RemoveData(p_name string){
	if gt.PropertyTable[_DataList] == nil {
		return
	}
	datalist_property := gt.PropertyTable[_DataList].(property.PropertyT)
	if datalist_property.GetValue() != nil {
		map_datalist := datalist_property.GetValue().(map[string]inf.IData)
		delete(map_datalist, p_name)
		datalist_property.SetValue(map_datalist)
		gt.PropertyTable[_DataList] = datalist_property
	}//TODO: contract.component_table delete component
	dataExpression_property := gt.PropertyTable[_DataValueSetterExpressionList].(property.PropertyT)
	if dataExpression_property.GetValue() != nil {
		map_dataExpression := dataExpression_property.GetValue().(map[string]inf.IExpression)
		delete(map_dataExpression, p_name)
		dataExpression_property.SetValue(map_dataExpression)
		gt.PropertyTable[_DataValueSetterExpressionList] = dataExpression_property
	}
}
//====运行条件判断
func (gt *GeneralTask)testCompleteCondition()bool{
	var r_flag bool = false
	if len(gt.GetPreCondition()) == 0{
		r_flag = true
	}
	for _,value := range gt.GetPreCondition(){
		v_contract := gt.GetContract()
		if v_contract.EvaluateExpression(value.GetExpressionStr()).(bool){
			r_flag = false
		}
	}
	return r_flag
}

func (gt *GeneralTask)testDisgardCondition()bool{
	var r_flag bool = false
	if len(gt.GetDisgardCondition()) == 0{
		r_flag = true
	}
	for _,value := range gt.GetDisgardCondition(){
		v_contract := gt.GetContract()
		if v_contract.EvaluateExpression(value.GetExpressionStr()).(bool){
			r_flag = false
		}
	}
	return r_flag
}

func (gt *GeneralTask)testPreCondition()bool{
	var r_flag bool = false
	if len(gt.GetPreCondition()) == 0{
		r_flag = true
	}
	for _,value := range gt.GetPreCondition(){
		v_contract := gt.GetContract()
		if v_contract.EvaluateExpression(value.GetExpressionStr()).(bool){
			r_flag = false
		}
	}
	return r_flag
}
//====运行状态控制
func (gt *GeneralTask)IsDormant()bool{
	return gt.GetState() == constdef.TaskState[constdef.TaskState_Dormant]
}

func (gt *GeneralTask)IsInProgress()bool{
	return gt.GetState() == constdef.TaskState[constdef.TaskState_In_Progress]
}

func (gt *GeneralTask)IsCompleted()bool{
	return gt.GetState() == constdef.TaskState[constdef.TaskState_Completed]
}

func (gt *GeneralTask)IsDisgarded()bool{
	return gt.GetState() == constdef.TaskState[constdef.TaskState_Disgarded]
}

//任务运行前进行的预处理
func (gt *GeneralTask)PreProcess()error{
	var r_err error = nil

	return r_err
}

//用于执行回滚操作，回滚后将任务状态改为dromant
func (gt *GeneralTask)Dromant() (int8, error) {
	//TODO log
	fmt.Println("Task[", gt.GetName(), "] State[Dromant] begin....")
	var r_ret int8 = 0
	var r_err error = nil
	if gt.IsInProgress() || gt.IsCompleted(){

		//TODO
		fmt.Println("Task[", gt.GetName(), "] State[Start to Dromant]....")
		gt.SetState(constdef.TaskState[constdef.TaskState_Dormant])
		r_ret = -1
		r_err = errors.New("Task execute fail, complete roclback!")
	}
	return r_ret,r_err
}

func (gt *GeneralTask)Start() (int8, error) {
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Contract Runing:Dromant State.")
	r_buf.WriteString("[ContractID]: " + gt.GetContract().GetContractId() + ";")
	r_buf.WriteString("[TaskName]: " + gt.GetName() + ";")
	//TODO log
	fmt.Println(r_buf.String(), " begin....")
	var r_ret int8 = 0
	var r_err error = nil
	if gt.IsDormant() && gt.testPreCondition() {
		// TODO：执行结果统一结构描述
		//r_resoponse := gt.TestMethod() 比如查询操作，资产创建，资产转移中的冻结资产，非资产动作
		//if 执行失败{
		//    r_ret = -1
		//    r_err = "执行失败"
		//    执行结果Set到DataValueSetterExpressionList中的result中
		//TODO log
		//    fmt.Println(r_buf.String(), " fail....")
		//    return r_ret,r_err
		//}
        //r_response 中的data set 到 Data中， 执行结果Set到DataValueSetterExpressionList中的result结果中
		//TODO log
		fmt.Println(r_buf.String(), " Dormant to Inprocess....")
		gt.SetState(constdef.TaskState[constdef.TaskState_In_Progress])
	} else if gt.IsDormant() && !gt.testPreCondition(){
		r_ret = 0
		//TODO log
		fmt.Println(r_buf.String(), " exit....")
		return r_ret,r_err
	}
	//执行完动作后需要等待执行完成
	var v_exit_flag int8 = 0
	for v_exit_flag == 0 {
		r_ret,r_err = gt.Complete()
		if r_ret == 0{
			continue
		}else {
			break
		}
	}
	return r_ret,r_err
}

func (gt *GeneralTask)Complete() (int8, error) {
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Contract Runing:Inprogress State.")
	r_buf.WriteString("[ContractID]: " + gt.GetContract().GetContractId() + ";")
	r_buf.WriteString("[TaskName]: " + gt.GetName() + ";")
	//TODO log
	fmt.Println(r_buf.String(), " begin....")
	var r_ret int8 = 0
	var r_err error = nil
	// CompleteCondition 需要包含单节点任务执行结果条件
	//   任务执行成功，继续往下执行
	//   任务执行失败，该任务需要重新执行
	if gt.IsInProgress() && gt.testCompleteCondition() {
		//TODO log
		fmt.Println(r_buf.String(), " Inprocess to Complete....")
		gt.SetState(constdef.TaskState[constdef.TaskState_Completed])
		//TODO
		//1.判断状态变化，需要写入产出
		//1).根据任务类型构建合约运行产出（合约产出交易、合约运行状态）
		//2).对运行产出签名
		//3).运行产出存入产出表
		//r_resoponse := gt.TestMethod()
		//if 执行失败{
		//    r_ret = -1
		//    r_err = "执行失败"
		//TODO log
		//    fmt.Println(r_buf.String(), " fail....")
		//    return r_ret,r_err
		//}
	} else if gt.IsInProgress() && !gt.testCompleteCondition() {
		r_ret = 0
		return r_ret,r_err
	}
	//保证顺利执行，给执行方法留下执行时间，需要多次sleep等待执行
	var v_sleep_num uint8 = 60
	for v_sleep_num > 0 {
		v_sleep_num = v_sleep_num - 1
		r_ret,r_err = gt.Disgard()
		if r_ret == 0{
			time.Sleep(60)
		} else {
			break
		}
	}
	return r_ret,r_err
}


func (gt *GeneralTask)Disgard()(int8, error){
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Contract Runing:Complete State.")
	r_buf.WriteString("[ContractID]: " + gt.GetContract().GetContractId() + ";")
	r_buf.WriteString("[TaskName]: " + gt.GetName() + ";")
	fmt.Println(r_buf.String(), " begin....")
	var r_ret int8 = 0
	var r_err error = nil
	// DisgardCondition 需要包含多节点共识结果标识
	//   任务执行结果共识通过后，继续往下执行；
	//   任务执行结果共识不通过，该任务需要重新执行；
	if gt.IsCompleted() && gt.testDisgardCondition(){
		//DisgardCondition中需要默认添加任务执行结果入链判断条件
		//TODO log
		fmt.Println(r_buf.String(), " Complete to Digcard....")
		gt.SetState(constdef.TaskState[constdef.TaskState_Disgarded])
		r_ret = 1
	}
	return r_ret,r_err
}

//任务运行后进行的后处理
func (gt *GeneralTask)PostProcess(p_flag int8) error{
	var r_err error = nil
	//TODO: 获取当前机器公钥，当前合约ID，新建合约ID
	switch p_flag{
	case -1:
		//执行失败：1.更新contractID1 的flag=0, failNum+1, timestamp
		//TODO: 业鹏提供接口 UpdateMonitorFail(node_pubkey, contractID)
	case 0:
		//执行条件不满足：1.更新contractID1 的flag=0，timestamp
		//TODO: 业鹏提供接口 UpdateMonitorWait(node_pubkey, contractID)
	case 1:
		//执行成功：1 更新contractID1 的flag=1, succNum+1, timestamp, 2.将contractID2插入到扫描监控表中
		//TODO: 业鹏提供接口 UpdateMonitorSucc(node_pubkey, contractID1, contractID2)
	}
	return r_err
}

func (gt *GeneralTask)TestMethod() error{
	var r_err error = nil
	return r_err
}