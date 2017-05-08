package contract

//UCVM：描述态 =》 运行态 =》 持久态
//      描述态： contract描述json文件文件 或 json串
//      运行态： 通过反序列化得到contract实例，然后调用Init方法，完成运行态的初始化
//      持久态： 执行结果 和 运行状态持久化到数据表中
import (
	"encoding/json"
	"fmt"
	"errors"

	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/table"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/component"
)

type CognitiveContract struct {
	component.GeneralComponent
	//根据实际业务场景增加的属性
	ContractId string `json:"ContractId"`
	ContractState string `json:"ContractState"`
	Creator string `json:"Creator"`
	CreatorTime string `json:"CreatorTime"`
	StartTime string `json:"StartTime"`
	EndTime string `json:"EndTime"`
    ContractOwners []string `json:"ContractOwners"`
	ContractAssets []ContractAsset `json:"ContractAssets"`
	ContractSignatures []ContractSignature `json:"ContractSignatures"`
	//type: Unknown, Data, Task, Expression
	ContractComponents []interface{} `json:"ContractComponents"`
	NextTasks []string `json:"NextTasks"`

	//TODO: need sort struct
	//type: map[string][]property.PropertyT   type:  Unknown, Data, Task, Expression
	ComponentTable table.ComponentTable  `json:"-"`
	//EventQueue event.EventQueue  `json:"-"`
	//EventHandlerPool event.EventHandlerPool  `json:"-"`
}

const (
	_ContractId = "_ContractId"
	_ContractState = "_ContractState"
	_Creator = "_Creator"
	_CreatorTime = "_CreatorTime"
	_StartTime = "_StartTime"
	_EndTime = "_EndTime"
	_ContractOwners = "_ContractOwners"
	_ContractAssets = "_ContractAssets"
	_ContractSignatures = "_ContractSignatures"
	_ContractComponents = "_ContractComponents"
	_NextTasks = "_NextTasks"
	_UCVM_Version = "_UCVM_Version"
	_UCVM_CopyRight = "_UCVM_CopyRight"
	_UCVM_Date = "_UCVM_Date"
)

func NewCognitiveContract() *CognitiveContract{
	cc := &CognitiveContract{}
	return cc
}

//===============接口实现===================
func (cc CognitiveContract) GetContractId()string {
	contractid_property := cc.PropertyTable[_ContractId].(property.PropertyT)
	return contractid_property.GetValue().(string)
}

func (cc CognitiveContract)GetVersion() string{
	return constdef.UCVM_Version
}

func (cc CognitiveContract) GetCopyRight() string {
	return constdef.UCVM_CopyRight
}

func (cc *CognitiveContract)GetTask(p_name string)interface{}{
	return cc.ComponentTable.GetComponent(p_name, constdef.ComponentType[constdef.Component_Task])
}

func (cc *CognitiveContract) AddComponent(p_component inf.IComponent) {
	if p_component != nil {
		var v_contract inf.ICognitiveContract = inf.ICognitiveContract(cc)
		p_component.SetContract(v_contract)
		cc.ComponentTable.AddComponent(p_component)
	}
}

func (cc CognitiveContract)EvaluateExpression(p_expression interface{})interface{}{
	//TODO
	//var str_expression string
	return p_expression.(bool)
}

//Description: Process the expression enclosed by <% %> in string
func (cc CognitiveContract) ProcessString(p_str string)string{
	//TODO
	return ""
}
//===============描述态=====================
//合约对象序列化
func (model *CognitiveContract)Serialize()(string, error){
	var err error = nil
	if model == nil {
		return "",err
	}
	//TODO：序列化时，ContratComponents的值转化(由ComponentTable得来)
	if s_model,err := json.Marshal(model);err == nil {
		return string(s_model),err
	}else {
		//TODO log
		fmt.Println(err)
		return "",err
	}
}

//合约对象反序列化
func (model *CognitiveContract)Deserialize(p_str string) (*CognitiveContract ,error){
	var err error = nil
	if p_str == "" || model == nil{
		return nil, err
	}
	if err := json.Unmarshal([]byte(p_str), &model); err != nil {
		//TODO log
		fmt.Println(err)
		return nil, err
	}
	return model, err
}
//===============运行态=====================
func (cc *CognitiveContract) AddProperty(object interface{}, str_name string, value interface{})property.PropertyT {
	var pro_object property.PropertyT
	if value == nil {
		pro_object = *property.NewPropertyT(str_name)
		cc.PropertyTable[str_name] = pro_object
		return pro_object
	}
	switch value.(type) {
	case string:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(string))
		cc.PropertyTable[str_name] = pro_object
		cc.ReflectSetValue(object, str_name, value)
	case []string:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.([]string))
		cc.PropertyTable[str_name] = pro_object
		cc.ReflectSetValue(object, str_name, value)
	case []ContractAsset:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.([]ContractAsset))
		cc.PropertyTable[str_name] = pro_object
		cc.ReflectSetValue(object, str_name, value)
	case []ContractSignature:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.([]ContractSignature))
		cc.PropertyTable[str_name] = pro_object
		cc.ReflectSetValue(object, str_name, value)
	case []interface{}:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.([]interface{}))
		cc.PropertyTable[str_name] = pro_object
		cc.ReflectSetValue(object, str_name, value)
	default:
		fmt.Println("[", str_name, ":", value, "]value type not support!!!")
	}
	return pro_object
}
func (cc *CognitiveContract)InitCognitiveContract()error{
	var err error = nil
	if cc.Cname == "" {
		//TODO log
		err = errors.New("Contract Need Cname!")
		return err
	}
	if cc.Caption == "" {
		//TODO log
		err = errors.New("Contract Need Caption!")
		return err
	}
	if cc.Description == "" {
		//TODO log
		err = errors.New("Contract Need Description!")
		return err
	}
	if cc.Creator == "" {
		//TODO log
		err = errors.New("Contract Need Creator!")
		return err
	}
	if cc.CreatorTime == "" {
		//TODO log
		err = errors.New("Contract Need CreatorTime!")
		return err
	}
	if cc.StartTime == "" {
		//TODO log
		err = errors.New("Contract Need StartTime!")
		return err
	}
	if cc.EndTime == "" {
		//TODO log
		err = errors.New("Contract Need EndTime!")
		return err
	}
	if cc.ContractOwners == nil || len(cc.ContractOwners) == 0 {
		//TODO log
		err = errors.New("Contract Need ContractOwners!")
		return err
	}
	if cc.ContractAssets == nil || len(cc.ContractAssets) == 0  {
		//TODO log
		err = errors.New("Contract Need ContractAssets!")
		return err
	}
	if cc.ContractSignatures == nil || len(cc.ContractSignatures) == 0  {
		//TODO log
		err = errors.New("Contract Need ContractOwners!")
		return err
	}
	err = cc.InitGeneralComponent()
	if err != nil {
		//TODO: log
		return err
	}
	cc.SetCtype(constdef.ComponentType[constdef.Component_Contract])
	cc.AddProperty(cc, _ContractId, cc.ContractId)
	cc.ContractState = common.TernaryOperator(cc.ContractState == "", constdef.ContractState[constdef.Contract_Create], cc.ContractState).(string)
	cc.AddProperty(cc, _ContractState, cc.ContractState)
	cc.AddProperty(cc, _Creator, cc.Creator)
	cc.AddProperty(cc, _CreatorTime, cc.CreatorTime)
	cc.AddProperty(cc, _StartTime, cc.StartTime)
	cc.AddProperty(cc, _EndTime, cc.EndTime)
	cc.AddProperty(cc, _ContractOwners, cc.ContractOwners)
	cc.AddProperty(cc, _ContractAssets, cc.ContractAssets)
	cc.AddProperty(cc, _ContractSignatures, cc.ContractSignatures)
	cc.AddProperty(cc, _NextTasks, cc.NextTasks)

	var meta_map map[string]string = make(map[string]string, 0)
	meta_map[_UCVM_Version] = constdef.UCVM_Version
	meta_map[_UCVM_CopyRight] = constdef.UCVM_CopyRight
	meta_map[_UCVM_Date] = constdef.UCVM_Date
    cc.GeneralComponent.AddMetaAttribute(meta_map)

	cc.ComponentTable = *new(table.ComponentTable)

	//cc.EventQueue = *new(event.EventQueue)
	//cc.EventHandlerPool = *new(event.EventHandlerPool)
	return err
}
//====预处理方法
func (cc *CognitiveContract) loadBuildInFunctions(){
	//TODO
}

func (cc *CognitiveContract) loadExpressionParser(){
	//TODO
}
//====动态增加
func (cc *CognitiveContract)AddContractWoner(p_owner string){
	contractOwners_property := cc.PropertyTable[_ContractOwners].(property.PropertyT)
	if contractOwners_property.GetValue() == nil{
		contractOwners_property.SetValue(make([]string, 0))
	}
	if p_owner != "" {
		v_subject_list := contractOwners_property.GetValue().([]string)
		contractOwners_property.SetValue(append(v_subject_list, p_owner))
	}
	cc.PropertyTable[_ContractOwners] = contractOwners_property
	cc.ContractOwners = contractOwners_property.GetValue().([]string)
}
func (cc *CognitiveContract)AddContractAsset(p_asset ContractAsset){
	contractAssets_property := cc.PropertyTable[_ContractAssets].(property.PropertyT)
	if contractAssets_property.GetValue() != nil {
		contractAssets_property.SetValue(make([]ContractAsset, 0))
	}
	v_asset_list := contractAssets_property.GetValue().([]ContractAsset)
	contractAssets_property.SetValue(append(v_asset_list, p_asset))

	cc.PropertyTable[_ContractAssets] = contractAssets_property
	cc.ContractAssets = contractAssets_property.GetValue().([]ContractAsset)
}
func (cc *CognitiveContract)AddSignature(p_signature ContractSignature){
	contractSignature_property := cc.PropertyTable[_ContractSignatures].(property.PropertyT)
	if contractSignature_property.GetValue() != nil {
		contractSignature_property.SetValue(make([]ContractSignature, 0))
	}
	v_signature_list := contractSignature_property.GetValue().([]ContractSignature)
	contractSignature_property.SetValue(append(v_signature_list, p_signature))

	cc.PropertyTable[_ContractSignatures] = contractSignature_property
	cc.ContractSignatures = contractSignature_property.GetValue().([]ContractSignature)
}

//====组件性操作
func (cc *CognitiveContract)GetComponentByType(c_type string)[]map[string]interface{}{
	if c_type == "" {
		return nil
	}
	if _, ok := cc.ComponentTable.CompTable[c_type]; !ok {
		return nil
	}
	return cc.ComponentTable.CompTable[c_type]
}

func (cc *CognitiveContract)GetTasks()[]map[string]interface{}{
	return cc.GetComponentByType(constdef.ComponentType[constdef.Component_Task])
}

func (cc *CognitiveContract)GetData(p_name string)interface{}{
	return cc.ComponentTable.GetComponent(p_name, constdef.ComponentType[constdef.Component_Data])
}

func (cc *CognitiveContract)GetExpression(p_name string)interface{}{
	return cc.ComponentTable.GetComponent(p_name, constdef.ComponentType[constdef.Component_Expression])
}

func (cc *CognitiveContract)GetTtem(p_name string)interface{}{
	return cc.ComponentTable.GetComponent(p_name, "")
}

//====属性Get方法
func (cc *CognitiveContract) GetContractState()string {
	state_property := cc.PropertyTable[_ContractState].(property.PropertyT)
	return state_property.GetValue().(string)
}
func (cc *CognitiveContract)GetCreator()string {
	creator_property := cc.PropertyTable[_Creator].(property.PropertyT)
	return creator_property.GetValue().(string)
}
func (cc *CognitiveContract)GetCreatorTime()string{
	creatorTime_property := cc.PropertyTable[_CreatorTime].(property.PropertyT)
	return creatorTime_property.GetValue().(string)
}
func (cc *CognitiveContract)GetStartTime()string{
	startTime_property := cc.PropertyTable[_StartTime].(property.PropertyT)
	return startTime_property.GetValue().(string)
}
func (cc *CognitiveContract)GetEndTime()string{
	endTime_property := cc.PropertyTable[_EndTime].(property.PropertyT)
	return endTime_property.GetValue().(string)
}
func (cc *CognitiveContract)GetContractOwners()interface{}{
	contractOwners_property := cc.PropertyTable[_ContractOwners].(property.PropertyT)
	return contractOwners_property.GetValue().([]string)
}
func (cc *CognitiveContract)GetContractAssets()interface{}{
	contractAssets_property := cc.PropertyTable[_ContractAssets].(property.PropertyT)
	return contractAssets_property.GetValue().([]ContractAsset)
}
func (cc *CognitiveContract)GetContractSignatures()interface{}{
	contractSignatures := cc.PropertyTable[_ContractSignatures].(property.PropertyT)
	return contractSignatures.GetValue().([]ContractSignature)
}
func (cc *CognitiveContract) GetContractComponents()[]interface{}{
	return cc.ContractComponents
}
func (cc *CognitiveContract) GetNextTasks()[]string{
	nexttasks_property := cc.PropertyTable[_NextTasks].(property.PropertyT)
	return nexttasks_property.GetValue().([]string)
}
//====属性Set方法
func (cc *CognitiveContract)SetContractId(p_ConstractId  string){
	cc.ContractId = p_ConstractId
	contractid_property := cc.PropertyTable[_ContractId].(property.PropertyT)
	contractid_property.SetValue(p_ConstractId)
	cc.PropertyTable[_ContractId] = contractid_property
}
func (cc *CognitiveContract)SetContractState(p_State  string){
	cc.ContractState = p_State
	state_property := cc.PropertyTable[_ContractState].(property.PropertyT)
	state_property.SetValue(p_State)
	cc.PropertyTable[_ContractState] = state_property
}
func (cc *CognitiveContract)SetCreator(p_Creator string){
	cc.Creator = p_Creator
	creator_property := cc.PropertyTable[_Creator].(property.PropertyT)
	creator_property.SetValue(p_Creator)
	cc.PropertyTable[_Creator] = creator_property
}
func (cc *CognitiveContract)SetCreatorTime(p_CreatorTime string){
	cc.CreatorTime = p_CreatorTime
	creatortime_property := cc.PropertyTable[_CreatorTime].(property.PropertyT)
	creatortime_property.SetValue(p_CreatorTime)
	cc.PropertyTable[_CreatorTime] = creatortime_property
}
func (cc *CognitiveContract)SetStartTime(p_StartTime string){
	cc.StartTime = p_StartTime
	starttime_property := cc.PropertyTable[_StartTime].(property.PropertyT)
	starttime_property.SetValue(p_StartTime)
	cc.PropertyTable[_StartTime] = starttime_property
}
func (cc *CognitiveContract)SetEndTime(p_EndTime string){
	cc.EndTime = p_EndTime
	endtime_property := cc.PropertyTable[_EndTime].(property.PropertyT)
	endtime_property.SetValue(p_EndTime)
	cc.PropertyTable[_EndTime] = endtime_property
}
func (cc *CognitiveContract)SetContractOwners(p_ContractOwners []string){
	cc.ContractOwners = p_ContractOwners
	contractowners_property := cc.PropertyTable[_ContractOwners].(property.PropertyT)
	contractowners_property.SetValue(p_ContractOwners)
	cc.PropertyTable[_ContractOwners] = contractowners_property
}
func (cc *CognitiveContract)SetContractAssets (p_ContractAssets []ContractAsset){
	cc.ContractAssets = p_ContractAssets
	contractassets_property := cc.PropertyTable[_ContractAssets].(property.PropertyT)
	contractassets_property.SetValue(p_ContractAssets)
	cc.PropertyTable[_ContractAssets] = contractassets_property
}
func (cc *CognitiveContract)SetContractSignatures (p_ContractSignatures []ContractSignature){
	cc.ContractSignatures = p_ContractSignatures
	contractsignatures_property := cc.PropertyTable[_ContractSignatures].(property.PropertyT)
	contractsignatures_property.SetValue(p_ContractSignatures)
	cc.PropertyTable[_ContractSignatures] = contractsignatures_property
}
func (cc *CognitiveContract)SetNextTasks (p_NextTasks []string){
	cc.NextTasks = p_NextTasks
	nexttasks_property := cc.PropertyTable[_NextTasks].(property.PropertyT)
	nexttasks_property.SetValue(p_NextTasks)
	cc.PropertyTable[_NextTasks] = nexttasks_property
}
//====事件处理
/*
func (cc *CognitiveContract)RegisterEvent(p_event event.GeneralEvent){
	cc.EventQueue.AddEvent(p_event)
	if p_event.GetUrgency() == constdef.EventPriority[constdef.EventPriority_Immediate] {
		cc.EventQueue.FireTopHandler()
	}
}

func (cc *CognitiveContract)ProcessEventQueue(){
	cc.EventQueue.FireHandlers()
}
*/
//====任务执行
// 1.从start节点的后继任务开始运行，将后续任务如队列
// 2.轮询判断队列中的任务(寻找应当执行的任务)
// 3.    后继任务中都是dromant state的任务，将后继任务重新入队，进入6 轮询判断；
// 4.    后继任务中有digcard的任务，则将同级任务跳过；将该任务的后继任务加入队列；调回2 重新判定
// 5.    后继任务中有inprocess 或 completed 的任务，将该同级任务跳过；将该任务加入队列；进入6 轮询判断
// 6.轮询判断队列中的任务（执行任务）
// 7.    任务是inprocess 或 completed state, 执行执行
// 8.    任务是dromant state,需要轮询队列中的任务，是否可以执行；
// 9.          不满足运行条件，继续判断同级任务
// 10.         满足运行条件，则执行该任务，跳过队列中的其他同级任务
func (cc *CognitiveContract)UpdateTasksState() (int8, error){
	var r_ret int8 = -1
	var r_err error = nil
	var next_tasks []string = cc.GetNextTasks()
	fmt.Println(next_tasks)
	if next_tasks == nil || len(next_tasks) == 0 {
		r_err = errors.New("contract has no start tasks!")
		return r_ret,r_err
	}
	var r_task_queue *common.Queue = common.NewQueue()
	for _,v_task := range next_tasks {
		r_task_queue.Push(v_task)
	}
	//判断后继任务是否有执行过(state_digcard)的：
	//     有(state_digcard)，则清空队列，将该任务后继任务入队，继续判断；
	//     有(state_inprocess or state_completed),则清空队列，将该任务入队，跳出判断，进入下一判断
	//     无(且队列不空时)，继续判断
	//     无(且队列为空时)，则将当前轮询的后继任务入队，跳出循环，进入下一判断
	for !r_task_queue.Empty() {
		tmp_str_task := r_task_queue.Pop()
		f_f_task := cc.GetTask(tmp_str_task.(string))
		if f_f_task == nil {
			r_ret = -1
			r_err = errors.New("GetTask is null!")
			return r_ret,r_err
		}
		if f_f_task.(inf.ITask).GetState() == constdef.TaskState[constdef.TaskState_Disgarded] {
			for r_task_queue.Len() != 0 {
				r_task_queue.Pop()
			}
			next_tasks = f_f_task.(inf.ITask).GetNextTasks()
			for _, t_task := range next_tasks {
				r_task_queue.Push(t_task)
			}
			continue
		} else if f_f_task.(inf.ITask).GetState() == constdef.TaskState[constdef.TaskState_In_Progress]  || f_f_task.(inf.ITask).GetState() == constdef.TaskState[constdef.TaskState_Completed] {
			for r_task_queue.Len() != 0 {
				r_task_queue.Pop()
			}
			r_task_queue.Push(f_f_task.(inf.ITask).GetName())
			break
	    }else if r_task_queue.Len() != 0 {
			continue
		} else {
			for _,v_task := range next_tasks {
				r_task_queue.Push(v_task)
			}
			break
		}
	}
	//执行任务流，任务执行返回的状态：
	//       -1: 任务状态流转过程中，在某一状态时，执行失败，返回 -1; State=Inprocess, Completed
	//       0 : 任务状态流转过程中，在某一状态时，达不到执行条件 返回0; State=Dromant, Inprocess, Completed
	//       1 : 任务状态流转完成，才会返回 1; State=Digcard
	//注：此处只代表单个任务的执行结果，每次执行只能执行一个任务
    var f_err error = nil
	for !r_task_queue.Empty() {
		tmp_str_task := r_task_queue.Pop()
		f_s_task := cc.GetTask(tmp_str_task.(string))
		if f_s_task == nil {
			r_ret = -1
			r_err = errors.New("GetTask is null!")
			return r_ret,r_err
		}
		r_ret,f_err = f_s_task.(inf.ITask).UpdateState()
		switch r_ret{
		case 1://执行成功后，合约退出， 注意：后续任务不入队列了，等待共识成功后初始化到扫描监控表中，下次加载再执行
			break
			/*
			for r_task_queue.Len() != 0 {
				r_task_queue.Pop()
			}
			next_tasks = f_s_task.(inf.ITask).GetNextTasks()
			for _, t_task := range next_tasks {
				//注意：解决循环执行任务问题，当后继任务入队时，需要将后继任务更新为Dromant状态
				//      通过循环执行次数条件,退出循环执行
				tmp_next_task := cc.GetTask(t_task)
				if tmp_next_task != nil {
					tmp_next_task.(inf.ITask).SetState(constdef.TaskState[constdef.TaskState_Dormant])
					r_task_queue.Push(t_task)
				}
			}
			*/
		case 0://执行条件不成立
			if f_s_task.(inf.ITask).GetState() == constdef.TaskState[constdef.TaskState_Dormant] {//继续判断同级中的下一任务
				continue
			} else {//合约退出
				r_err = errors.New("task["+ f_s_task.(inf.ITask).GetName()+"] condition not fullfill!")
				break
			}
		case -1://执行失败后，合约退出
			r_err = errors.New("task["+ f_s_task.(inf.ITask).GetName()+"] execute fail!")
			break
		}
		if f_err != nil {
			//TODO log
			fmt.Println("Error: ["+f_err.Error()+"]")
		}
	}
	return r_ret,r_err
}