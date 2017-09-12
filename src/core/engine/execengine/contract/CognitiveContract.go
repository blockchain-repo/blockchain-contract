package contract

//UCVM：描述态 =》 运行态 =》 持久态
//      描述态： contract描述json文件文件 或 json串
//      运行态： 通过反序列化得到contract实例，然后调用Init方法，完成运行态的初始化
//      持久态： 执行结果 和 运行状态持久化到数据表中
import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/expressionutils"
	"unicontract/src/core/engine/execengine/function"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
	"unicontract/src/core/engine/execengine/table"
)

type CognitiveContract struct {
	Id           string                `json:"id"`
	ContractHead CognitiveContractHead `json:"ContractHead"`
	ContractBody CognitiveContractBody `json:"ContractBody"`

	PropertyTable map[string]interface{} `json:"-"`
	//type: map[string][]property.PropertyT
	//      Unknown, Contract, Data, Task, Expression
	ComponentTable        *table.ComponentTable                  `json:"-"`
	ExpressionParseEngine *expressionutils.ExpressionParseEngine `json:"-"`
	FunctionParseEngine   *function.FunctionParseEngine          `json:"-"`
	//任务执行过程中需要的临时存储变量
	OrgId                string      `json:"-"`
	OrgTaskId            string      `json:"-"`
	OrgTaskExecuteIdx    int         `json:"-"`
	OutputId             string      `json:"-"`
	OutputTaskId         string      `json:"-"`
	OutputTaskExecuteIdx int         `json:"-"`
	OutputStruct         interface{} `json:"-"`
}

type CognitiveContractHead struct {
	AssignTime      string `json:"AssignTime"`
	MainPubkey      string `json:"MainPubkey"`
	OperateTime     string `json:"OperateTime"`
	Version         int    `json:"Version"`
	ConsensusResult int    `json:"ConsensusResult"`
}

type CognitiveContractBody struct {
	//合约默认属性
	ContractId         string              `json:"ContractId"`
	Cname              string              `json:"Cname"`
	Ctype              string              `json:"Ctype"`
	Caption            string              `json:"Caption"`
	Description        string              `json:"Description"`
	ContractState      string              `json:"ContractState"`
	Creator            string              `json:"Creator"`
	CreateTime         string              `json:"CreateTime"`
	StartTime          string              `json:"StartTime"`
	EndTime            string              `json:"EndTime"`
	ContractOwners     []string            `json:"ContractOwners"`
	ContractAssets     []ContractAsset     `json:"ContractAssets"`
	ContractSignatures []ContractSignature `json:"ContractSignatures"`
	ContractComponents []interface{}       `json:"ContractComponents"` //type: Unknown, Contract, Data, Task, Expression
	NextTasks          []string            `json:"NextTasks"`
	//合约自定义属性（根据实际业务场景增加）
	MetaAttribute map[string]string `json:"MetaAttribute"`
}

const (
	_Id           = "_Id"
	_ContractHead = "_ContractHead"
	_ContractBody = "_ContractBody"

	_MainPubkey      = "_MainPubkey"
	_AssignTime      = "_AssignTime"
	_OperateTime     = "_OperateTime"
	_Version         = "_Version"
	_ConsensusResult = "_ConsensusResult"

	_Cname              = "_Cname"
	_Ctype              = "_Ctype"
	_Caption            = "_Caption"
	_Description        = "_Description"
	_MetaAttribute      = "_MetaAttribute"
	_ContractId         = "_ContractId"
	_ContractState      = "_ContractState"
	_Creator            = "_Creator"
	_CreateTime         = "_CreateTime"
	_StartTime          = "_StartTime"
	_EndTime            = "_EndTime"
	_ContractOwners     = "_ContractOwners"
	_ContractAssets     = "_ContractAssets"
	_ContractSignatures = "_ContractSignatures"
	_ContractComponents = "_ContractComponents"
	_NextTasks          = "_NextTasks"
	_UCVM_Version       = "_UCVM_Version"
	_UCVM_CopyRight     = "_UCVM_CopyRight"
	_UCVM_Date          = "_UCVM_Date"

	_OrgId                = "_OrgId"
	_OrgTaskId            = "_OrgTaskId"
	_OrgTaskExecuteIdx    = "_OrgTaskExecuteIdx"
	_OutputId             = "_OutputId"
	_OutputTaskId         = "_OutputTaskId"
	_OutputTaskExecuteIdx = "_OutputTaskExecuteIdx"
	_OutputStruct         = "_OutputStruct"
)

func NewCognitiveContract() *CognitiveContract {
	cc := &CognitiveContract{}
	return cc
}

//===============接口实现===================
func (cc CognitiveContract) GetContractId() string {
	contractid_property, ok := cc.PropertyTable[_ContractId].(property.PropertyT)
	if !ok {
		return ""
	}
	value, ok := contractid_property.GetValue().(string)
	if !ok {
		return ""
	}
	return value
}

func (cc CognitiveContract) GetUCVMVersion() string {
	return constdef.UCVM_Version
}

func (cc CognitiveContract) GetUCVMCopyRight() string {
	return constdef.UCVM_CopyRight
}

func (cc *CognitiveContract) GetTask(p_name string) interface{} {
	return cc.ComponentTable.GetComponent(p_name, constdef.ComponentType[constdef.Component_Task])
}

func (cc *CognitiveContract) GetTaskByID(p_task_id string) interface{} {
	return cc.ComponentTable.GetTaskByID(p_task_id, constdef.ComponentType[constdef.Component_Task])
}

func (cc *CognitiveContract) GetComponentTtem(p_name string) interface{} {
	return cc.ComponentTable.GetComponent(p_name, "")
}

//Note:获取PropertyTable中的属性的值，为了保持统一的获取对象元素的方法
//Return: interface{}
func (cc *CognitiveContract) GetPropertyItem(p_name string) interface{} {
	if p_name != "" && cc.PropertyTable != nil {
		v_property, ok := cc.PropertyTable[p_name].(property.PropertyT)
		if !ok {
			return nil
		}
		return v_property.GetValue()
	}
	return nil
}

//将合约本身添加到ComponentTable中
func (cc *CognitiveContract) AddComponent(p_component inf.IComponent) {
	if p_component != nil {
		var v_contract inf.ICognitiveContract = inf.ICognitiveContract(cc)
		p_component.SetContract(v_contract)
		cc.ComponentTable.AddComponent(p_component)
	}
}

//Args: p_exprtype    => 表达式的类型（常量，变量，条件，函数，决策）
//      p_expression  => 表达式内容
func (cc CognitiveContract) EvaluateExpression(p_exprtype string, p_expression string) (interface{}, error) {
	v_ret, v_err := cc.ExpressionParseEngine.EvaluateExpressionValue(p_exprtype, p_expression)
	if v_err != nil {
		uniledgerlog.Error("EvaluateExpressionValue Fail[" + v_err.Error() + "]")
	}
	return v_ret, v_err
}

//Description: Process the expression enclosed by <% %> in string
//暂时不考虑：待补充
func (cc CognitiveContract) ProcessString(p_str string) string {
	//TODO
	//字符串表达式中<% %>代表引用变量值
	//1. 提取字符串中 <% %> 变量个数
	//2. 依次求提取变量表达式的值
	//3. 替换字符串中 <% %> 为真实值
	return p_str
}

func (cc CognitiveContract) SetContract(p_contract inf.ICognitiveContract) {
	//为实现接口而设置的空方法
}

func (cc CognitiveContract) GetContract() inf.ICognitiveContract {
	return &cc
}

func (cc CognitiveContract) GetName() string {
	return cc.GetCname()
}
func (gc *CognitiveContract) GetCtype() string {
	if gc.PropertyTable[_Ctype] == nil {
		return ""
	}
	ctype_property, ok := gc.PropertyTable[_Ctype].(property.PropertyT)
	if !ok {
		return ""
	}
	value, ok := ctype_property.GetValue().(string)
	if !ok {
		return ""
	}
	return value
}

func (gc *CognitiveContract) GetId() string {
	if gc.PropertyTable[_Id] == nil {
		return ""
	}
	id_property, ok := gc.PropertyTable[_Id].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	str, ok := id_property.GetValue().(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	return str
}
func (gc *CognitiveContract) GetOrgTaskId() string {
	if gc.PropertyTable[_OrgTaskId] == nil {
		return ""
	}
	orgtaskid_property, ok := gc.PropertyTable[_OrgTaskId].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	str, ok := orgtaskid_property.GetValue().(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	return str
}

func (gc *CognitiveContract) GetOrgTaskExecuteIdx() int {
	if gc.PropertyTable[_OrgTaskExecuteIdx] == nil {
		return 0
	}
	orgtaskexecuteidx_property, ok := gc.PropertyTable[_OrgTaskExecuteIdx].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0
	}
	value, ok := orgtaskexecuteidx_property.GetValue().(int)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0
	}
	return value
}

func (gc *CognitiveContract) GetOutputId() string {
	if gc.PropertyTable[_OutputId] == nil {
		return ""
	}
	outputid_property, ok := gc.PropertyTable[_OutputId].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	str, ok := outputid_property.GetValue().(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	return str
}

func (gc *CognitiveContract) GetOutputTaskId() string {
	if gc.PropertyTable[_OutputTaskId] == nil {
		return ""
	}
	OutputTaskId_property, ok := gc.PropertyTable[_OutputTaskId].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	str, ok := OutputTaskId_property.GetValue().(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	return str
}

func (gc *CognitiveContract) GetOutputTaskExecuteIdx() int {
	if gc.PropertyTable[_OutputTaskExecuteIdx] == nil {
		return 0
	}
	OutputTaskExecuteIdx_property, ok := gc.PropertyTable[_OutputTaskExecuteIdx].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0
	}
	value, ok := OutputTaskExecuteIdx_property.GetValue().(int)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0
	}
	return value
}

func (gc *CognitiveContract) GetOutputStruct() string {
	if gc.PropertyTable[_OutputStruct] == nil {
		return ""
	}
	outputstruct_property, ok := gc.PropertyTable[_OutputStruct].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	if outputstruct_property.GetValue() == nil {
		return ""
	}
	str, ok := outputstruct_property.GetValue().(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	return str
}

func (cc CognitiveContract) SetOrgId(p_OrgId string) {
	//Take case: Setter method need set value for gc.xxxxxx
	cc.OrgId = p_OrgId
	OrgId_property, ok := cc.PropertyTable[_OrgId].(property.PropertyT)
	if !ok {
		OrgId_property = *property.NewPropertyT(_OrgId)
	}
	OrgId_property.SetValue(p_OrgId)
	//Take case: Setter method need set value for gc.PropertyTable[xxxx]
	cc.PropertyTable[_OrgId] = OrgId_property
}
func (cc CognitiveContract) SetOrgTaskId(p_OrgTaskId string) {
	//Take case: Setter method need set value for gc.xxxxxx
	cc.OrgTaskId = p_OrgTaskId
	OrgTaskId_property, ok := cc.PropertyTable[_OrgTaskId].(property.PropertyT)
	if !ok {
		OrgTaskId_property = *property.NewPropertyT(_OrgTaskId)
	}
	OrgTaskId_property.SetValue(p_OrgTaskId)
	//Take case: Setter method need set value for gc.PropertyTable[xxxx]
	cc.PropertyTable[_OrgTaskId] = OrgTaskId_property
}

func (cc CognitiveContract) SetOrgTaskExecuteIdx(p_OrgTaskExecuteIdx int) {
	//Take case: Setter method need set value for gc.xxxxxx
	cc.OrgTaskExecuteIdx = p_OrgTaskExecuteIdx
	OrgTaskExecuteIdx_property, ok := cc.PropertyTable[_OrgTaskExecuteIdx].(property.PropertyT)
	if !ok {
		OrgTaskExecuteIdx_property = *property.NewPropertyT(_OrgTaskExecuteIdx)
	}
	OrgTaskExecuteIdx_property.SetValue(p_OrgTaskExecuteIdx)
	//Take case: Setter method need set value for gc.PropertyTable[xxxx]
	cc.PropertyTable[_OrgTaskExecuteIdx] = OrgTaskExecuteIdx_property
}
func (cc CognitiveContract) SetOutputId(p_outputId string) {
	//Take case: Setter method need set value for gc.xxxxxx
	cc.OutputId = p_outputId
	outputid_property, ok := cc.PropertyTable[_OutputId].(property.PropertyT)
	if !ok {
		outputid_property = *property.NewPropertyT(_OutputId)
	}
	outputid_property.SetValue(p_outputId)
	//Take case: Setter method need set value for gc.PropertyTable[xxxx]
	cc.PropertyTable[_OutputId] = outputid_property
}
func (cc CognitiveContract) SetOutputTaskId(p_OutputTaskId string) {
	//Take case: Setter method need set value for gc.xxxxxx
	cc.OutputTaskId = p_OutputTaskId
	OutputTaskId_property, ok := cc.PropertyTable[_OutputTaskId].(property.PropertyT)
	if !ok {
		OutputTaskId_property = *property.NewPropertyT(_OutputTaskId)
	}
	OutputTaskId_property.SetValue(p_OutputTaskId)
	//Take case: Setter method need set value for gc.PropertyTable[xxxx]
	cc.PropertyTable[_OutputTaskId] = OutputTaskId_property
}

func (cc CognitiveContract) SetOutputTaskExecuteIdx(p_OutputTaskExecuteIdx int) {
	//Take case: Setter method need set value for gc.xxxxxx
	cc.OutputTaskExecuteIdx = p_OutputTaskExecuteIdx
	OutputTaskExecuteIdx_property, ok := cc.PropertyTable[_OutputTaskExecuteIdx].(property.PropertyT)
	if !ok {
		OutputTaskExecuteIdx_property = *property.NewPropertyT(_OutputTaskExecuteIdx)
	}
	OutputTaskExecuteIdx_property.SetValue(p_OutputTaskExecuteIdx)
	//Take case: Setter method need set value for gc.PropertyTable[xxxx]
	cc.PropertyTable[_OutputTaskExecuteIdx] = OutputTaskExecuteIdx_property
}

func (cc CognitiveContract) SetOutputStruct(p_OutputStruct string) {
	//Take case: Setter method need set value for gc.xxxxxx
	cc.OutputStruct = p_OutputStruct
	OutputStruct_property, ok := cc.PropertyTable[_OutputStruct].(property.PropertyT)
	if !ok {
		OutputStruct_property = *property.NewPropertyT(_OutputStruct)
	}
	OutputStruct_property.SetValue(p_OutputStruct)
	//Take case: Setter method need set value for gc.PropertyTable[xxxx]
	cc.PropertyTable[_OutputStruct] = OutputStruct_property
}

func (gc *CognitiveContract) GetMainPubkey() string {
	if gc.PropertyTable[_MainPubkey] == nil {
		return ""
	}
	mainpubkey_property, ok := gc.PropertyTable[_MainPubkey].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	str, ok := mainpubkey_property.GetValue().(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	return str
}

func (cc CognitiveContract) SetMainPubkey(p_mainPubkey string) {
	//Take case: Setter method need set value for gc.xxxxxx
	cc.ContractHead.MainPubkey = p_mainPubkey
	mainpubkey_property, ok := cc.PropertyTable[_MainPubkey].(property.PropertyT)
	if !ok {
		mainpubkey_property = *property.NewPropertyT(_MainPubkey)
	}
	mainpubkey_property.SetValue(p_mainPubkey)
	//Take case: Setter method need set value for gc.PropertyTable[xxxx]
	cc.PropertyTable[_MainPubkey] = mainpubkey_property
}

//===============描述态=====================
//合约对象序列化
func (model *CognitiveContract) Serialize() (string, error) {
	var err error = nil
	if model == nil {
		return "", err
	}

	var task_count int = len(model.ComponentTable.CompTable[constdef.ComponentType[constdef.Component_Task]])
	component_array := model.ComponentTable.CompTable[constdef.ComponentType[constdef.Component_Task]]
	var new_contract_components []interface{} = make([]interface{}, task_count)
	for v_idx, _ := range component_array {
		if len(component_array[v_idx]) == 0 {
			err = fmt.Errorf("ComponentTable has nil task!")
			uniledgerlog.Error("Contract Serialize fail[" + err.Error() + "]")
			return "", err
		}
		//type: map[string]inf.ITask
		for v_key, _ := range component_array[v_idx] {
			//update data & expression in task
			ttask, ok := component_array[v_idx][v_key].(inf.ITask)
			if !ok {
				uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
				return "", fmt.Errorf("assert error")
			}
			new_task, err := ttask.UpdateStaticState()
			if err != nil {
				err = fmt.Errorf("Task.UpdateStaticState fail!")
				uniledgerlog.Error("Task.UpdateStaticState fail[" + err.Error() + "]")
				return "", err
			}
			task_map := common.Deserialize(common.Serialize(new_task))
			new_contract_components[v_idx] = task_map
			break
		}
	}
	model.ContractBody.ContractComponents = new_contract_components
	if s_model, err := json.Marshal(model); err == nil {
		return string(s_model), err
	} else {
		uniledgerlog.Error("Contract Serialize fail[" + err.Error() + "]")
		return "", err
	}
}

//合约对象反序列化
//return: *CognitiveContract
func (model *CognitiveContract) Deserialize(p_str string) (interface{}, error) {
	var err error = nil
	if p_str == "" || model == nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(p_str), &model); err != nil {
		uniledgerlog.Error("Contract Deserialize fail[" + err.Error() + "]")
		return nil, err
	}
	return model, err
}

//更新合约状态ContractState
//Args: p_state string
//return: true 更新成功； false 更新失败；
func (cc *CognitiveContract) UpdateContractState(p_state string) bool {
	var v_bool bool = true
	var v_contract_state string = cc.GetContractState()
	if p_state != "" {
		var v_exit_flag bool = false
		for _, v_value := range constdef.ContractState {
			if p_state == v_value {
				v_exit_flag = true
				break
			}
		}
		if !v_exit_flag {
			v_bool = false
			uniledgerlog.Error("UpdateContractState fail,[" + p_state + "]is not exist!")
			return v_bool
		} else {
			cc.SetContractState(p_state)
			uniledgerlog.Info("UpdateContractState success, contract_state change to " + p_state)
			return v_bool
		}
	}
	switch v_contract_state {
	case constdef.ContractState[constdef.Contract_Signature]:
		cc.SetContractState(constdef.ContractState[constdef.Contract_In_Process])
	case constdef.ContractState[constdef.Contract_Create]:
		cc.SetContractState(constdef.ContractState[constdef.Contract_In_Process])
	}
	return v_bool
}

//===============运行态=====================
//为防止包重复，本保内的属性，在该包内添加；公有的使用common中的
func (cc *CognitiveContract) AddProperty(object interface{}, str_name string, value interface{}) property.PropertyT {
	var pro_object property.PropertyT
	if value == nil {
		pro_object = *property.NewPropertyT(str_name)
		cc.PropertyTable[str_name] = pro_object
		return pro_object
	}
	switch value.(type) {
	case []ContractAsset:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.([]ContractAsset))
		cc.PropertyTable[str_name] = pro_object
		common.ReflectSetValue(object, str_name, value)
	case []ContractSignature:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.([]ContractSignature))
		cc.PropertyTable[str_name] = pro_object
		common.ReflectSetValue(object, str_name, value)
	default:
		//fmt.Println("[", str_name, ":", value, "]value type not support!!!")
		uniledgerlog.Error("[", str_name, ":", value, "]value type not support!!!")
	}
	return pro_object
}
func (cc *CognitiveContract) InitCognitiveContract() error {
	var err error = nil
	if cc.PropertyTable == nil {
		cc.PropertyTable = make(map[string]interface{}, 0)
	}
	//ID初始化
	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s) %s]",
		uniledgerlog.NO_ERROR, cc.ContractBody.ContractId, "struct init ID"))
	common.AddProperty(cc, cc.PropertyTable, _Id, cc.Id)
	//ContractHead初始化
	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s) %s]",
		uniledgerlog.NO_ERROR, cc.ContractBody.ContractId, "struct init ContractHead"))
	common.AddProperty(cc, cc.PropertyTable, _MainPubkey, cc.ContractHead.MainPubkey)
	common.AddProperty(cc, cc.PropertyTable, _AssignTime, cc.ContractHead.AssignTime)
	common.AddProperty(cc, cc.PropertyTable, _OperateTime, cc.ContractHead.OperateTime)
	common.AddProperty(cc, cc.PropertyTable, _Version, cc.ContractHead.Version)
	common.AddProperty(cc, cc.PropertyTable, _ConsensusResult, cc.ContractHead.ConsensusResult)
	//ContractBody初始化
	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s) %s]",
		uniledgerlog.NO_ERROR, cc.ContractBody.ContractId, "struct init ContractBody"))
	if cc.ContractBody.Cname == "" {
		uniledgerlog.Warn("Contract Need Cname!")
		err = errors.New("Contract Need Cname!")
		return err
	}
	if cc.ContractBody.Caption == "" {
		uniledgerlog.Warn("Contract Need Caption!")
		err = errors.New("Contract Need Caption!")
		return err
	}
	if cc.ContractBody.Description == "" {
		uniledgerlog.Warn("Contract Need Description!")
		err = errors.New("Contract Need Description!")
		return err
	}
	if cc.ContractBody.Creator == "" {
		uniledgerlog.Warn("Contract Need Creator!")
		err = errors.New("Contract Need Creator!")
		return err
	}
	if cc.ContractBody.CreateTime == "" {
		uniledgerlog.Warn("Contract Need CreateTime!")
		err = errors.New("Contract Need CreateTime!")
		return err
	}
	if cc.ContractBody.StartTime == "" {
		uniledgerlog.Warn("Contract Need StartTime!")
		err = errors.New("Contract Need StartTime!")
		return err
	}
	if cc.ContractBody.EndTime == "" {
		uniledgerlog.Warn("Contract Need EndTime!")
		err = errors.New("Contract Need EndTime!")
		return err
	}
	if cc.ContractBody.ContractOwners == nil || len(cc.ContractBody.ContractOwners) == 0 {
		uniledgerlog.Warn("Contract Need ContractOwners!")
		err = errors.New("Contract Need ContractOwners!")
		return err
	}
	if cc.ContractBody.ContractAssets == nil || len(cc.ContractBody.ContractAssets) == 0 {
		uniledgerlog.Warn("Contract Need ContractAssets!")
		err = errors.New("Contract Need ContractAssets!")
		return err
	}
	if cc.ContractBody.ContractSignatures == nil || len(cc.ContractBody.ContractSignatures) == 0 {
		uniledgerlog.Warn("Contract Need ContractOwners!")
		err = errors.New("Contract Need ContractOwners!")
		return err
	}
	if cc.ContractBody.MetaAttribute == nil {
		cc.ContractBody.MetaAttribute = make(map[string]string, 0)
	}
	//将描述态数据加载成运行态，因此value都是gc.xxxx(描述态的)
	common.AddProperty(cc, cc.PropertyTable, _Cname, cc.ContractBody.Cname)
	str, ok := common.TernaryOperator(cc.ContractBody.Ctype == "",
		constdef.ComponentType[constdef.Component_Contract], cc.ContractBody.Ctype).(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return fmt.Errorf("assert error")
	}
	cc.ContractBody.Ctype = str
	common.AddProperty(cc, cc.PropertyTable, _Ctype, cc.ContractBody.Ctype)
	common.AddProperty(cc, cc.PropertyTable, _Caption, cc.ContractBody.Caption)
	common.AddProperty(cc, cc.PropertyTable, _Description, cc.ContractBody.Description)
	common.AddProperty(cc, cc.PropertyTable, _MetaAttribute, cc.ContractBody.MetaAttribute)
	common.AddProperty(cc, cc.PropertyTable, _ContractId, cc.ContractBody.ContractId)
	str, ok = common.TernaryOperator(cc.ContractBody.ContractState == "",
		constdef.ContractState[constdef.Contract_Create], cc.ContractBody.ContractState).(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return fmt.Errorf("assert error")
	}
	cc.ContractBody.ContractState = str
	common.AddProperty(cc, cc.PropertyTable, _ContractState, cc.ContractBody.ContractState)
	common.AddProperty(cc, cc.PropertyTable, _Creator, cc.ContractBody.Creator)
	common.AddProperty(cc, cc.PropertyTable, _CreateTime, cc.ContractBody.CreateTime)
	common.AddProperty(cc, cc.PropertyTable, _StartTime, cc.ContractBody.StartTime)
	common.AddProperty(cc, cc.PropertyTable, _EndTime, cc.ContractBody.EndTime)
	common.AddProperty(cc, cc.PropertyTable, _ContractOwners, cc.ContractBody.ContractOwners)
	common.AddProperty(cc, cc.PropertyTable, _NextTasks, cc.ContractBody.NextTasks)
	//自有类型，自己解决添加
	cc.AddProperty(cc, _ContractAssets, cc.ContractBody.ContractAssets)
	cc.AddProperty(cc, _ContractSignatures, cc.ContractBody.ContractSignatures)

	//过程中的临时变量
	common.AddProperty(cc, cc.PropertyTable, _OrgId, cc.OrgId)
	common.AddProperty(cc, cc.PropertyTable, _OrgTaskId, cc.OrgTaskId)
	common.AddProperty(cc, cc.PropertyTable, _OrgTaskExecuteIdx, cc.OrgTaskExecuteIdx)
	common.AddProperty(cc, cc.PropertyTable, _OutputId, cc.OutputId)
	common.AddProperty(cc, cc.PropertyTable, _OutputTaskId, cc.OutputTaskId)
	common.AddProperty(cc, cc.PropertyTable, _OutputTaskExecuteIdx, cc.OutputTaskExecuteIdx)
	common.AddProperty(cc, cc.PropertyTable, _OutputStruct, cc.OutputStruct)

	var meta_map map[string]string = make(map[string]string, 0)
	meta_map[_UCVM_Version] = constdef.UCVM_Version
	meta_map[_UCVM_CopyRight] = constdef.UCVM_CopyRight
	meta_map[_UCVM_Date] = constdef.UCVM_Date
	cc.AddMetaAttribute(meta_map)

	//初始化指针变量
	cc.ComponentTable = table.NewComponentTable()
	cc.FunctionParseEngine = function.NewFunctionParseEngine()
	cc.ExpressionParseEngine = expressionutils.NewExpressionParseEngine()

	// TODO 以后应该去掉
	//合约预处理：初始化FunctionEngine & ExpressionEngine
	cc.loadBuildInFunctions()
	cc.loadExpressionParser()

	return err
}

//====预处理方法
func (cc *CognitiveContract) loadBuildInFunctions() {
	cc.FunctionParseEngine.LoadFunctionsCommon()
}

func (cc *CognitiveContract) loadExpressionParser() {
	cc.ExpressionParseEngine.SetFunctionEngine(cc.FunctionParseEngine)
	var v_contract inf.ICognitiveContract = cc
	cc.ExpressionParseEngine.SetContract(v_contract)
}

//====动态增加
func (cc *CognitiveContract) AddContractWoner(p_owner string) {
	contractOwners_property, ok := cc.PropertyTable[_ContractOwners].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	if contractOwners_property.GetValue() == nil {
		contractOwners_property.SetValue(make([]string, 0))
	}
	if p_owner != "" {
		v_subject_list, ok := contractOwners_property.GetValue().([]string)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return
		}
		contractOwners_property.SetValue(append(v_subject_list, p_owner))
	}
	cc.PropertyTable[_ContractOwners] = contractOwners_property
	cc.ContractBody.ContractOwners, ok = contractOwners_property.GetValue().([]string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
}
func (cc *CognitiveContract) AddContractAsset(p_asset ContractAsset) {
	contractAssets_property, ok := cc.PropertyTable[_ContractAssets].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	if contractAssets_property.GetValue() != nil {
		contractAssets_property.SetValue(make([]ContractAsset, 0))
	}
	v_asset_list, ok := contractAssets_property.GetValue().([]ContractAsset)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	contractAssets_property.SetValue(append(v_asset_list, p_asset))

	cc.PropertyTable[_ContractAssets] = contractAssets_property
	cc.ContractBody.ContractAssets, ok = contractAssets_property.GetValue().([]ContractAsset)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
}
func (cc *CognitiveContract) AddSignature(p_signature ContractSignature) {
	contractSignature_property, ok := cc.PropertyTable[_ContractSignatures].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	if contractSignature_property.GetValue() != nil {
		contractSignature_property.SetValue(make([]ContractSignature, 0))
	}
	v_signature_list, ok := contractSignature_property.GetValue().([]ContractSignature)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	contractSignature_property.SetValue(append(v_signature_list, p_signature))

	cc.PropertyTable[_ContractSignatures] = contractSignature_property
	cc.ContractBody.ContractSignatures, ok = contractSignature_property.GetValue().([]ContractSignature)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
}
func (gc *CognitiveContract) AddMetaAttribute(metaProperty interface{}) {
	tmp, ok := metaProperty.(map[string]string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	if metaProperty != nil && len(tmp) != 0 {
		metaAttribute_property, ok := gc.PropertyTable[_MetaAttribute].(property.PropertyT)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return
		}
		tmp, ok := metaAttribute_property.GetValue().(map[string]string)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return
		}
		if metaAttribute_property.GetValue() == nil || len(tmp) == 0 {
			metaAttribute_property.SetValue(make(map[string]string, 0))
		}
		v_metaProperty, ok := metaProperty.(map[string]string)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return
		}
		for key, value := range v_metaProperty {
			// TODO ??????????
			metaAttribute_property.GetValue().(map[string]string)[key] = value
		}
		gc.PropertyTable[_MetaAttribute] = metaAttribute_property
		gc.ContractBody.MetaAttribute, ok = metaAttribute_property.GetValue().(map[string]string)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return
		}
	}
}

//====组件性操作
func (cc *CognitiveContract) GetComponentByType(c_type string) []map[string]interface{} {
	if c_type == "" {
		return nil
	}
	if _, ok := cc.ComponentTable.CompTable[c_type]; !ok {
		return nil
	}
	return cc.ComponentTable.CompTable[c_type]
}

func (cc *CognitiveContract) GetTasks() []map[string]interface{} {
	return cc.GetComponentByType(constdef.ComponentType[constdef.Component_Task])
}

func (cc *CognitiveContract) GetData(p_name string) interface{} {
	return cc.ComponentTable.GetComponent(p_name, constdef.ComponentType[constdef.Component_Data])
}

func (cc *CognitiveContract) GetExpression(p_name string) interface{} {
	return cc.ComponentTable.GetComponent(p_name, constdef.ComponentType[constdef.Component_Expression])
}

//获取PropertyTable
//return: map[string]property.propertyT
func (cc *CognitiveContract) GetPropertyTable() map[string]interface{} {
	return cc.PropertyTable
}

//Note: PropertyTable的key为属性变量名大写加_前缀，如：_NAME
//return: property.propertyT
func (cc *CognitiveContract) GetProperty(p_name string) interface{} {
	if p_name != "" && cc.PropertyTable != nil {
		value, ok := cc.PropertyTable[p_name].(property.PropertyT)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return nil
		}
		return value
	}
	return nil
}

//====属性Get方法	common.AddProperty(cc, cc.PropertyTable, _OrgId, cc.OrgId)
func (gc *CognitiveContract) GetCname() string {
	if gc.PropertyTable[_Cname] == nil {
		return ""
	}
	cname_property, ok := gc.PropertyTable[_Cname].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	str, ok := cname_property.GetValue().(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	return str
}

func (gc *CognitiveContract) GetCaption() string {
	var r_res string = ""
	if gc.PropertyTable[_Caption] == nil {
		r_res = ""
	} else {
		caption_property, ok := gc.PropertyTable[_Caption].(property.PropertyT)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return ""
		}
		r_res, ok = caption_property.GetValue().(string)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return ""
		}
	}
	return r_res
}

func (gc *CognitiveContract) GetDescription() string {
	var r_res string = ""
	if gc.PropertyTable[_Description] != nil {
		description_property, ok := gc.PropertyTable[_Description].(property.PropertyT)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return ""
		}
		r_res, ok = description_property.GetValue().(string)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return ""
		}
	}
	return r_res
}

func (gc *CognitiveContract) GetMetaAttribute() map[string]string {
	if gc.PropertyTable[_MetaAttribute] == nil {
		return nil
	}
	metaattribute_property, ok := gc.PropertyTable[_MetaAttribute].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	v, ok := metaattribute_property.GetValue().(map[string]string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	return v
}

//属性Set方法
func (cc CognitiveContract) SetId(p_Id string) {
	//Take case: Setter method need set value for gc.xxxxxx
	cc.Id = p_Id
	id_property, ok := cc.PropertyTable[_Id].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	id_property.SetValue(p_Id)
	//Take case: Setter method need set value for gc.PropertyTable[xxxx]
	cc.PropertyTable[_Id] = id_property
}
func (gc *CognitiveContract) SetCname(str_name string) {
	//Take case: Setter method need set value for gc.xxxxxx
	gc.ContractBody.Cname = str_name
	cname_property, ok := gc.PropertyTable[_Cname].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	cname_property.SetValue(str_name)
	//Take case: Setter method need set value for gc.PropertyTable[xxxx]
	gc.PropertyTable[_Cname] = cname_property
}

func (cc *CognitiveContract) GetContractState() string {
	state_property, ok := cc.PropertyTable[_ContractState].(property.PropertyT)
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
func (cc *CognitiveContract) GetCreator() string {
	creator_property, ok := cc.PropertyTable[_Creator].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	str, ok := creator_property.GetValue().(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	return str
}
func (cc *CognitiveContract) GetCreateTime() string {
	CreateTime_property, ok := cc.PropertyTable[_CreateTime].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	str, ok := CreateTime_property.GetValue().(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	return str
}
func (cc *CognitiveContract) GetStartTime() string {
	startTime_property, ok := cc.PropertyTable[_StartTime].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	str, ok := startTime_property.GetValue().(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	return str
}
func (cc *CognitiveContract) GetEndTime() string {
	endTime_property, ok := cc.PropertyTable[_EndTime].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	str, ok := endTime_property.GetValue().(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	return str
}
func (cc *CognitiveContract) GetContractOwners() interface{} {
	contractOwners_property, ok := cc.PropertyTable[_ContractOwners].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	str, ok := contractOwners_property.GetValue().([]string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	return str
}
func (cc *CognitiveContract) GetContractAssets() interface{} {
	contractAssets_property, ok := cc.PropertyTable[_ContractAssets].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	str, ok := contractAssets_property.GetValue().([]ContractAsset)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	return str
}
func (cc *CognitiveContract) GetContractSignatures() interface{} {
	contractSignatures, ok := cc.PropertyTable[_ContractSignatures].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	str, ok := contractSignatures.GetValue().([]ContractSignature)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	return str
}
func (cc *CognitiveContract) GetContractComponents() []interface{} {
	return cc.ContractBody.ContractComponents
}
func (cc *CognitiveContract) GetNextTasks() []string {
	nexttasks_property, ok := cc.PropertyTable[_NextTasks].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	str, ok := nexttasks_property.GetValue().([]string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	return str
}

//====属性Set方法
func (gc *CognitiveContract) SetCtype(str_type string) {
	gc.ContractBody.Ctype = str_type
	ctype_property, ok := gc.PropertyTable[_Ctype].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	ctype_property.SetValue(str_type)
	gc.PropertyTable[_Ctype] = ctype_property
}

func (gc *CognitiveContract) SetCaption(str_Caption string) {
	gc.ContractBody.Caption = str_Caption
	caption_property, ok := gc.PropertyTable[_Caption].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	caption_property.SetValue(str_Caption)
	gc.PropertyTable[_Caption] = caption_property
}

func (gc *CognitiveContract) SetDescription(str_Description string) {
	gc.ContractBody.Description = str_Description
	description_property, ok := gc.PropertyTable[_Description].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	description_property.SetValue(str_Description)
	gc.PropertyTable[_Description] = description_property
}

func (gc *CognitiveContract) SetMetaAttribute(p_metaAttribute map[string]string) {
	gc.ContractBody.MetaAttribute = p_metaAttribute
	metaAttribute_property, ok := gc.PropertyTable[_MetaAttribute].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	metaAttribute_property.SetValue(p_metaAttribute)
	gc.PropertyTable[_MetaAttribute] = metaAttribute_property
}

func (cc *CognitiveContract) SetContractId(p_ConstractId string) {
	cc.ContractBody.ContractId = p_ConstractId
	contractid_property, ok := cc.PropertyTable[_ContractId].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	contractid_property.SetValue(p_ConstractId)
	cc.PropertyTable[_ContractId] = contractid_property
}
func (cc *CognitiveContract) SetContractState(p_State string) {
	cc.ContractBody.ContractState = p_State
	state_property, ok := cc.PropertyTable[_ContractState].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	state_property.SetValue(p_State)
	cc.PropertyTable[_ContractState] = state_property
}
func (cc *CognitiveContract) SetCreator(p_Creator string) {
	cc.ContractBody.Creator = p_Creator
	creator_property, ok := cc.PropertyTable[_Creator].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	creator_property.SetValue(p_Creator)
	cc.PropertyTable[_Creator] = creator_property
}
func (cc *CognitiveContract) SetCreateTime(p_CreateTime string) {
	cc.ContractBody.CreateTime = p_CreateTime
	CreateTime_property, ok := cc.PropertyTable[_CreateTime].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	CreateTime_property.SetValue(p_CreateTime)
	cc.PropertyTable[_CreateTime] = CreateTime_property
}
func (cc *CognitiveContract) SetStartTime(p_StartTime string) {
	cc.ContractBody.StartTime = p_StartTime
	starttime_property, ok := cc.PropertyTable[_StartTime].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	starttime_property.SetValue(p_StartTime)
	cc.PropertyTable[_StartTime] = starttime_property
}
func (cc *CognitiveContract) SetEndTime(p_EndTime string) {
	cc.ContractBody.EndTime = p_EndTime
	endtime_property, ok := cc.PropertyTable[_EndTime].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	endtime_property.SetValue(p_EndTime)
	cc.PropertyTable[_EndTime] = endtime_property
}
func (cc *CognitiveContract) SetContractOwners(p_ContractOwners []string) {
	cc.ContractBody.ContractOwners = p_ContractOwners
	contractowners_property, ok := cc.PropertyTable[_ContractOwners].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	contractowners_property.SetValue(p_ContractOwners)
	cc.PropertyTable[_ContractOwners] = contractowners_property
}
func (cc *CognitiveContract) SetContractAssets(p_ContractAssets []ContractAsset) {
	cc.ContractBody.ContractAssets = p_ContractAssets
	contractassets_property, ok := cc.PropertyTable[_ContractAssets].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	contractassets_property.SetValue(p_ContractAssets)
	cc.PropertyTable[_ContractAssets] = contractassets_property
}
func (cc *CognitiveContract) SetContractSignatures(p_ContractSignatures []ContractSignature) {
	cc.ContractBody.ContractSignatures = p_ContractSignatures
	contractsignatures_property, ok := cc.PropertyTable[_ContractSignatures].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	contractsignatures_property.SetValue(p_ContractSignatures)
	cc.PropertyTable[_ContractSignatures] = contractsignatures_property
}
func (cc *CognitiveContract) SetNextTasks(p_NextTasks []string) {
	cc.ContractBody.NextTasks = p_NextTasks
	nexttasks_property, ok := cc.PropertyTable[_NextTasks].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	nexttasks_property.SetValue(p_NextTasks)
	cc.PropertyTable[_NextTasks] = nexttasks_property
}
func (cc *CognitiveContract) SetContractComponents(p_components []interface{}) {
	cc.ContractBody.ContractComponents = p_components
}

//====任务执行
// 1.从start节点的后继任务开始运行，将后续任务如队列
// 2.轮询判断队列中的任务(寻找应当执行的任务)
// 3.    后继任务中都是dromant state的任务，将后继任务重新入队，进入6 轮询判断；
// 4.    后继任务中有digcard的任务，则将同级任务跳过；将该任务的后继任务加入队列；调回2 重新判定
// 5.    后继任务中有inprocess/complete的任务，将该同级任务跳过；将该任务加入队列；进入6 轮询判断
// 6.轮询判断队列中的任务（执行任务）
// 7.    任务是inprocess state, 执行执行
// 8.    任务是dromant state,需要轮询队列中的任务，是否可以执行；
// 9.          不满足运行条件，继续判断同级任务
// 10.         满足运行条件，则执行该任务，跳过队列中的其他同级任务
func (cc *CognitiveContract) UpdateTasksState() (int8, error) {
	ok := false
	var r_ret int8 = -1
	var r_err error = nil
	var next_tasks []string = cc.GetNextTasks()
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Contract Executeing....:")
	r_buf.WriteString("[ContractID]: " + cc.GetContractId() + ";")
	r_buf.WriteString("[ContractHashID]: " + cc.GetId() + ";")
	if next_tasks == nil || len(next_tasks) == 0 {
		r_err = errors.New("contract has no start tasks!")
		r_buf.WriteString("[Result]: UpdateTasksState fail;")
		r_buf.WriteString("[Error]: " + r_err.Error() + ";")
		uniledgerlog.Warn(r_buf.String())
		return r_ret, r_err
	}
	var r_task_queue *common.Queue = common.NewQueue()
	for _, v_task := range next_tasks {
		r_task_queue.Push(v_task)
	}
	//根据合约中记录的当前执行任务ID（OrgTaskId）获取任务名称
	var do_process_task inf.ITask
	if cc.GetOrgTaskId() != "" {
		do_process_task, ok = cc.GetTaskByID(cc.GetOrgTaskId()).(inf.ITask)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return r_ret, fmt.Errorf("assert error")
		}
	}
	//判断后继任务是否有执行过(state_discard 或 state_completed)的：
	//     有(state_discard 或 state_completed)，则清空队列，将该任务后继任务入队，继续判断；
	//     有(state_inprocess),则清空队列，将该任务入队，跳出判断，进入下一判断
	//     无(且队列不空时)，继续判断
	//     无(且队列为空时)，则将当前轮询的后继任务入队，跳出循环，进入下一判断
	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s) %s]",
		uniledgerlog.NO_ERROR, cc.GetContractId(), "load current execute task"))
	for !r_task_queue.Empty() {
		tmp_str_task := r_task_queue.Pop()
		str, ok := tmp_str_task.(string)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return r_ret, fmt.Errorf("assert error")
		}
		f_f_task := cc.GetTask(str)
		if f_f_task == nil {
			r_ret = -1
			r_err = errors.New("Judge Task, GetTask is null!")
			r_buf.WriteString("[Result]: UpdateTasksState fail;")
			str, ok := tmp_str_task.(string)
			if !ok {
				uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
				return r_ret, fmt.Errorf("assert error")
			}
			r_buf.WriteString("[Error]: " + str + "," + r_err.Error() + ";")
			uniledgerlog.Warn(r_buf.String())
			return r_ret, r_err
		}
		ttask, ok := f_f_task.(inf.ITask)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return r_ret, fmt.Errorf("assert error")
		}
		if ttask.GetState() == constdef.TaskState[constdef.TaskState_Discard] ||
			ttask.GetState() == constdef.TaskState[constdef.TaskState_Completed] {
			for r_task_queue.Len() != 0 {
				r_task_queue.Pop()
			}
			//通过合约中记录的当前执行任务，则直接对后继任务进行重置，解决循环执行问题
			next_tasks = ttask.GetNextTasks()
			doTask, ok := do_process_task.(inf.ITask)
			if !ok {
				uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
				return r_ret, fmt.Errorf("assert error")
			}
			if ttask.GetName() == doTask.GetName() {
				uniledgerlog.Info("=========== NowProcessTask:" + doTask.GetName())
				uniledgerlog.Info("=========== NextProcessTask:" + ttask.GetName())
				for _, t_task := range next_tasks {
					//注意：解决循环执行任务问题，当后继任务入队时，需要将后继任务更新为Dromant状态
					//      通过循环执行次数条件,退出循环执行
					r_err = cc.UpdateLoopExecuteTask(t_task) // TODO ???
					if r_err != nil {
						r_ret = -1
						r_err = errors.New("UpdateLoopExecuteTask fai!")
						r_buf.WriteString("[Result]: UpdateLoopExecuteTask fail;")
						r_buf.WriteString("[Error]: " + t_task + "," + r_err.Error() + ";")
						uniledgerlog.Warn(r_buf.String())
						return r_ret, r_err
					}
					r_task_queue.Push(t_task)
				}
			} else {
				uniledgerlog.Error("----------- NowProcessTask:" + doTask.GetName())
				uniledgerlog.Error("----------- NextProcessTask:" + ttask.GetName())
				for _, t_task := range next_tasks {
					r_task_queue.Push(t_task)
				}
				continue
			}
		} else if ttask.GetState() == constdef.TaskState[constdef.TaskState_In_Progress] {
			for r_task_queue.Len() != 0 {
				r_task_queue.Pop()
			}
			r_task_queue.Push(ttask.GetName())
			break
		} else if r_task_queue.Len() != 0 {
			continue
		} else {
			for _, v_task := range next_tasks {
				r_task_queue.Push(v_task)
			}
			break
		}
	}
	//执行任务流，任务执行返回的状态：
	//       -1: 任务状态流转过程中，在某一状态时，执行失败，返回 -1; State=Dormaant, Inprocess
	//       0 : 任务状态流转过程中，在某一状态时，达不到执行条件 返回0; State=Dromant, Inprocess, Completed
	//       1 : 任务状态流转完成，才会返回 1; State=Digcard
	//注：此处只代表单个任务的执行结果，每次执行只能执行一个任务
	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s) %s]",
		uniledgerlog.NO_ERROR, cc.GetContractId(), "begin to execute current task"))
	var f_err error = nil
	for !r_task_queue.Empty() {
		tmp_str_task := r_task_queue.Pop()
		str, ok := tmp_str_task.(string)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return r_ret, fmt.Errorf("assert error")
		}
		f_s_task := cc.GetTask(str)
		if f_s_task == nil {
			r_ret = -1
			r_err = errors.New("Execute Task, GetTask is null!")
			r_buf.WriteString("[Result]: UpdateTasksState fail;")
			r_buf.WriteString("[Error]: " + r_err.Error() + ";")
			uniledgerlog.Warn(r_buf.String())
			return r_ret, r_err
		}
		aaaaa, ok := f_s_task.(inf.ITask)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return r_ret, fmt.Errorf("assert error")
		}
		uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s) current task name is (%s), id is (%s), ready to perform]",
			uniledgerlog.NO_ERROR, cc.GetContractId(), aaaaa.GetName(), aaaaa.GetTaskId()))
		r_ret, f_err = aaaaa.UpdateState()
		switch r_ret {
		case 1: //执行成功后，跳转到下一合约任务；
			// 注意：后续任务不入队列了，等待共识成功后初始化到扫描监控表中，下次加载再执行
			for r_task_queue.Len() != 0 {
				r_task_queue.Pop()
			}
			/*
				next_tasks = aaaaa.GetNextTasks()
				for _, t_task := range next_tasks {
					//注意：解决循环执行任务问题，当后继任务入队时，需要将后继任务更新为Dromant状态
					//      通过循环执行次数条件,退出循环执行
					r_err = cc.UpdateLoopExecuteTask(t_task)
					if r_err != nil {
						r_ret = -1
						r_err = errors.New("UpdateLoopExecuteTask fai!")
						r_buf.WriteString("[Result]: UpdateLoopExecuteTask fail;")
						r_buf.WriteString("[Error]: " + t_task + "," + r_err.Error() + ";")
						uniledgerlog.Warn(r_buf.String())
						return r_ret, r_err
					}
				}
				//避免重复执行，将当前contractHashID转化为outputID,保证执行的是下一个任务
				cc.SetId(cc.GetOutputId())
				cc.SetOrgId(cc.GetOutputId())
				cc.SetOrgTaskId(cc.GetOutputTaskId())
				cc.SetOrgTaskExecuteIdx(cc.GetOutputTaskExecuteIdx())
			*/
		case 0: //执行条件不成立
			if aaaaa.GetState() == constdef.TaskState[constdef.TaskState_Dormant] { //继续判断同级中的下一任务
				continue
			} else { //合约退出
				r_err = errors.New("task[" + aaaaa.GetName() + "] condition not fullfill!")
				break
			}
		case -1: //执行失败后，合约退出
			r_err = errors.New("task[" + aaaaa.GetName() + "] execute fail!")
			break
		}
		if f_err != nil {
			uniledgerlog.Error("Contract Task Execute has error" + f_err.Error())
		}
	}
	return r_ret, r_err
}

//check合约是否可执行（1. 合约签名齐全， 2. 合约起始日期达到）
//return :  true 可执行；  false 不可执行；
func (cc *CognitiveContract) CanExecute() bool {
	var v_bool bool = true
	tmp, ok := cc.GetContractOwners().([]string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return false
	}
	var v_owner_count int = len(tmp)
	tmp1, ok := cc.GetContractSignatures().([]ContractSignature)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return false
	}
	var v_signature_count int = len(tmp1)
	var v_contract_state string = cc.GetContractState()
	if v_contract_state == constdef.ContractState[constdef.Contract_Completed] || v_contract_state == constdef.ContractState[constdef.Contract_Discarded] {
		uniledgerlog.Warn("ContractState is Completed or Discarded, contract can't execute!")
		//此时强制更新扫描表合约执行状态，防止再次被扫描加载
		err := common.UpdateMonitorDeal(cc.GetContractId(), cc.GetId())
		if err != nil {
			uniledgerlog.Error("Contract Completed or Discarded, Update Task Monitor fail[" + err.Error() + "]")
		}
		v_bool = false
		return v_bool
	}
	//constract_state: Create or Signature need check signatures
	if v_contract_state == constdef.ContractState[constdef.Contract_Signature] || v_contract_state == constdef.ContractState[constdef.Contract_Create] {
		//check owners signature count
		if v_signature_count < v_owner_count {
			v_bool = false
			uniledgerlog.Error("contract signatures count not equals contract owners!")
			return v_bool
		}
		//check owners signaature content
		var v_idx int = 0
		sl, ok := cc.GetContractOwners().([]string)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return false
		}
		for _, v_owner := range sl {
			sl1, ok := cc.GetContractSignatures().([]ContractSignature)
			if !ok {
				uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
				return false
			}
			for v_idx, v_contract_signature := range sl1 {
				if v_owner == v_contract_signature.GetOwnerPubkey() {
					break
				}
				v_idx = v_idx + 1
			}
			if v_idx >= v_signature_count {
				v_bool = false
				uniledgerlog.Error("contract signatures content not all contract owners!")
				return v_bool
			}
		}
	}
	//check contract begin_time & end_time
	var now_date string = common.GenTimestamp()
	var contract_starttime string = cc.GetStartTime()
	var contract_endtime string = cc.GetEndTime()
	var v_err error = nil
	var now_date_int int64
	var starttime_int int64
	var endtime_int int64
	now_date_int, v_err = strconv.ParseInt(now_date, 10, 64)
	if v_err != nil {
		uniledgerlog.Error("Now_date ParseInt Error(" + v_err.Error() + ")!")
		v_bool = false
		return v_bool
	}
	starttime_int, v_err = strconv.ParseInt(contract_starttime, 10, 64)
	if v_err != nil {
		uniledgerlog.Error("Start_time ParseInt Error(" + v_err.Error() + ")!")
		v_bool = false
		return v_bool
	}
	endtime_int, v_err = strconv.ParseInt(contract_endtime, 10, 64)
	if v_err != nil {
		uniledgerlog.Error("End_time ParseInt Error(" + v_err.Error() + ")!")
		v_bool = false
		return v_bool
	}
	if now_date_int < starttime_int {
		uniledgerlog.Warn("Now_date not gt StartTime, can't execute contract!")
		v_bool = false
		return v_bool
	}
	if now_date_int > endtime_int {
		//合约超过截止时间，合约状态更新为：丢弃
		cc.SetContractState(constdef.ContractState[constdef.Contract_Discarded])
		uniledgerlog.Error("Now_date gt EndTime, contract can't execute!")
		v_bool = false
		return v_bool
	}
	return v_bool
}

//根据合约交易完整结构体中的Relation部分提取当前运行的任务信息，给OrgTaskinfo赋值
//Args: p_relation_map => ContractOutput结构中relation map结构
func (cc *CognitiveContract) SetOrgTaskInfo(p_relation_map map[string]interface{}) error {
	var v_err error = nil
	if p_relation_map == nil {
		v_err = errors.New("Param[p_relation_json] is nil!")
		uniledgerlog.Warn("SetOrgTaskInfo fail, Error[" + v_err.Error() + "]")
		return v_err
	}

	//提取ContractHashID
	v_contractHashID, ok := p_relation_map["ContractHashId"].(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return fmt.Errorf("assert error")
	}
	cc.SetOrgId(v_contractHashID)
	//提取TaskID
	v_taskID, ok := p_relation_map["TaskId"].(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return fmt.Errorf("assert error")
	}
	cc.SetOrgTaskId(v_taskID)
	//提取TaskIndexID
	f, ok := p_relation_map["TaskExecuteIdx"].(float64)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return fmt.Errorf("assert error")
	}
	var v_taskIndexID int = int(f)
	cc.SetOrgTaskExecuteIdx(v_taskIndexID)

	return v_err
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//更新执行态中的组件信息 【接口方法】
//Args; p_ctype     string       类型(task, data, expression)
//      p_name      string       名称
//      p_component interface{}  组件
func (cc *CognitiveContract) UpdateComponentRunningState(p_ctype string, p_name string, p_component interface{}) error {
	componentTable, err := cc.ComponentTable.UpdateComponent(p_ctype, p_name, p_component)
	cc.ComponentTable = componentTable
	if err != nil {
		uniledgerlog.Error("UpdateComponentRunningState fail, Ctype: " + p_ctype + ", Name: " + p_name)
	}
	return err
}

//更新合约中指定的任务信息内容
//Args: p_task_name  string   待修改任务的名称
func (cc *CognitiveContract) UpdateLoopExecuteTask(p_task_name string) error {
	var err error = nil
	if p_task_name == "" {
		err = fmt.Errorf("Param[p_task_name] is nil!")
		return err
	}
	//update task value
	task_component := cc.GetTask(p_task_name)
	if task_component != nil {
		v_nexttask_object, ok := task_component.(inf.ITask)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return fmt.Errorf("assert error")
		}
		if v_nexttask_object.GetState() == constdef.TaskState[constdef.TaskState_Completed] ||
			v_nexttask_object.GetState() == constdef.TaskState[constdef.TaskState_Discard] {
			//待循环执行的任务需要修改：State, TaskExecuteIdx；清空结果值
			v_nexttask_object.SetState(constdef.TaskState[constdef.TaskState_Dormant])
			v_nexttask_object.SetTaskExecuteIdx(v_nexttask_object.GetTaskExecuteIdx() + 1)
			v_nexttask_object.CleanValueInProcess()
		}
	}
	//update component property
	err = cc.UpdateContractComponents(task_component)
	if err != nil {
		return err
	}
	//update component_table
	err = cc.UpdateComponentRunningState(constdef.ComponentType[constdef.Component_Task], p_task_name, task_component)
	if err != nil {
		return err
	}
	return err
}

//更新合约属性ContractComponents，任务组件有更新，需要将整个结构体进行修改
//Args: p_task_component interface{} 待更新的合约任务组件
func (cc *CognitiveContract) UpdateContractComponents(p_task_component interface{}) error {
	var err error = nil
	var task_components []interface{} = cc.GetContractComponents()
	var new_task_components []interface{} = make([]interface{}, 0)
	for _, v_component := range task_components {
		if v_component == nil {
			continue
		}
		//序列化回来的component是map结构
		map_component, ok := v_component.(map[string]interface{})
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return fmt.Errorf("assert error")
		}
		//识别map中的中任务名称
		ttask, ok := p_task_component.(inf.ITask)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return fmt.Errorf("assert error")
		}
		if map_component["TaskId"] == ttask.GetTaskId() {
			//构造map结构的Task
			var str_json string = common.Serialize(p_task_component)
			var map_component interface{} = common.Deserialize(str_json)
			new_task_components = append(new_task_components, map_component)
		} else {
			new_task_components = append(new_task_components, v_component)
		}
	}
	cc.SetContractComponents(new_task_components)
	return err
}
