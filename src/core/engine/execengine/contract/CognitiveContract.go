package contract

//UCVM：描述态 =》 运行态 =》 持久态
//      描述态： contract描述json文件文件 或 json串
//      运行态： 通过反序列化得到contract实例，然后调用Init方法，完成运行态的初始化
//      持久态： 执行结果 和 运行状态持久化到数据表中
import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/astaxie/beego/logs"

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
	OutputId     string                `json:"OutputId"`
	ContractHead CognitiveContractHead `json:"ContractHead"`
	ContractBody CognitiveContractBody `json:"ContractBody"`

	PropertyTable map[string]interface{} `json:"-"`
	//type: map[string][]property.PropertyT
	//      Unknown, Data, Task, Expression
	ComponentTable        *table.ComponentTable                  `json:"-"`
	ExpressionParseEngine *expressionutils.ExpressionParseEngine `json:"-"`
	FunctionParseEngine   *function.FunctionParseEngine          `json:"-"`
}

type CognitiveContractHead struct {
	MainPubkey string `json:"MainPubkey"`
	Timestamp  string `json:"Timestamp"`
	Version    int    `json:"Version"`
}

type CognitiveContractBody struct {
	//合约默认属性
	Cname              string              `json:"Cname"`
	Ctype              string              `json:"Ctype"`
	Caption            string              `json:"Caption"`
	Description        string              `json:"Description"`
	ContractId         string              `json:"ContractId"`
	ContractState      string              `json:"ContractState"`
	Creator            string              `json:"Creator"`
	CreatorTime        string              `json:"CreatorTime"`
	StartTime          string              `json:"StartTime"`
	EndTime            string              `json:"EndTime"`
	ContractOwners     []string            `json:"ContractOwners"`
	ContractAssets     []ContractAsset     `json:"ContractAssets"`
	ContractSignatures []ContractSignature `json:"ContractSignatures"`
	ContractComponents []interface{}       `json:"ContractComponents"` //type: Unknown, Data, Task, Expression
	NextTasks          []string            `json:"NextTasks"`
	//合约自定义属性（根据实际业务场景增加）
	MetaAttribute map[string]string `json:"MetaAttribute"`
}

const (
	_Id           = "_Id"
	_OutputId     = "_OutputId"
	_ContractHead = "_ContractHead"
	_ContractBody = "_ContractBody"

	_MainPubkey = "_MainPubkey"
	_Timestamp  = "_Timestamp"
	_Version    = "_Version"

	_Cname              = "_Cname"
	_Ctype              = "_Ctype"
	_Caption            = "_Caption"
	_Description        = "_Description"
	_MetaAttribute      = "_MetaAttribute"
	_ContractId         = "_ContractId"
	_ContractState      = "_ContractState"
	_Creator            = "_Creator"
	_CreatorTime        = "_CreatorTime"
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
)

func NewCognitiveContract() *CognitiveContract {
	cc := &CognitiveContract{}
	return cc
}

//===============接口实现===================
func (cc CognitiveContract) GetContractId() string {
	contractid_property := cc.PropertyTable[_ContractId].(property.PropertyT)
	return contractid_property.GetValue().(string)
}

func (cc CognitiveContract) GetVersion() string {
	return constdef.UCVM_Version
}

func (cc CognitiveContract) GetCopyRight() string {
	return constdef.UCVM_CopyRight
}

func (cc *CognitiveContract) GetTask(p_name string) interface{} {
	return cc.ComponentTable.GetComponent(p_name, constdef.ComponentType[constdef.Component_Task])
}

func (cc *CognitiveContract) GetComponentTtem(p_name string) interface{} {
	return cc.ComponentTable.GetComponent(p_name, "")
}

//Note:获取PropertyTable中的属性的值，为了保持统一的获取对象元素的方法
//Return: interface{}
func (cc *CognitiveContract) GetPropertyItem(p_name string) interface{} {
	if p_name != "" && cc.PropertyTable != nil {
		v_property := cc.PropertyTable[p_name].(property.PropertyT)
		return v_property.GetValue()
	}
	return nil
}

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
		logs.Error("EvaluateExpressionValue Fail[" + v_err.Error() + "]")
	}
	return v_ret, v_err
}

//Description: Process the expression enclosed by <% %> in string
//暂时不考虑：待补充
func (cc CognitiveContract) ProcessString(p_str string) string {
	//TODO
	return p_str
}

func (cc CognitiveContract) GetName() string {
	return cc.GetCname()
}
func (gc *CognitiveContract) GetCtype() string {
	if gc.PropertyTable[_Ctype] == nil {
		return ""
	}
	ctype_property := gc.PropertyTable[_Ctype].(property.PropertyT)
	return ctype_property.GetValue().(string)
}
func (cc CognitiveContract) GetContract() inf.ICognitiveContract {
	return &cc
}
func (cc CognitiveContract) SetContract(p_contract inf.ICognitiveContract) {

}
func (gc *CognitiveContract) GetId() string {
	if gc.PropertyTable[_Id] == nil {
		return ""
	}
	id_property := gc.PropertyTable[_Id].(property.PropertyT)
	return id_property.GetValue().(string)
}
func (cc CognitiveContract) SetOutputId(p_outputId string) {
	//Take case: Setter method need set value for gc.xxxxxx
	cc.OutputId = p_outputId
	outputid_property := cc.PropertyTable[_OutputId].(property.PropertyT)
	outputid_property.SetValue(p_outputId)
	//Take case: Setter method need set value for gc.PropertyTable[xxxx]
	cc.PropertyTable[_OutputId] = outputid_property
}
func (gc *CognitiveContract) GetOutputId() string {
	if gc.PropertyTable[_OutputId] == nil {
		return ""
	}
	outputid_property := gc.PropertyTable[_OutputId].(property.PropertyT)
	return outputid_property.GetValue().(string)
}

//===============描述态=====================
//合约对象序列化
func (model *CognitiveContract) Serialize() (string, error) {
	var err error = nil
	if model == nil {
		return "", err
	}
	//TODO：序列化时，ContratComponents的值转化(由ComponentTable得来)
	if s_model, err := json.Marshal(model); err == nil {
		return string(s_model), err
	} else {
		logs.Error("Contract Serialize fail[" + err.Error() + "]")
		return "", err
	}
}

//合约对象反序列化
func (model *CognitiveContract) Deserialize(p_str string) (*CognitiveContract, error) {
	var err error = nil
	if p_str == "" || model == nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(p_str), &model); err != nil {
		logs.Error("Contract Deserialize fail[" + err.Error() + "]")
		return nil, err
	}
	return model, err
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
		fmt.Println("[", str_name, ":", value, "]value type not support!!!")
	}
	return pro_object
}
func (cc *CognitiveContract) InitCognitiveContract() error {
	var err error = nil
	if cc.PropertyTable == nil {
		cc.PropertyTable = make(map[string]interface{}, 0)
	}
	//ID初始化
	common.AddProperty(cc, cc.PropertyTable, _Id, cc.Id)
	common.AddProperty(cc, cc.PropertyTable, _OutputId, cc.OutputId)
	//ContractHead初始化
	common.AddProperty(cc, cc.PropertyTable, _MainPubkey, cc.ContractHead.MainPubkey)
	common.AddProperty(cc, cc.PropertyTable, _Timestamp, cc.ContractHead.Timestamp)
	common.AddProperty(cc, cc.PropertyTable, _Version, cc.ContractHead.Version)
	//ContractBody初始化
	if cc.ContractBody.Cname == "" {
		logs.Warning("Contract Need Cname!")
		err = errors.New("Contract Need Cname!")
		return err
	}
	if cc.ContractBody.Caption == "" {
		logs.Warning("Contract Need Caption!")
		err = errors.New("Contract Need Caption!")
		return err
	}
	if cc.ContractBody.Description == "" {
		logs.Warning("Contract Need Description!")
		err = errors.New("Contract Need Description!")
		return err
	}
	if cc.ContractBody.Creator == "" {
		logs.Warning("Contract Need Creator!")
		err = errors.New("Contract Need Creator!")
		return err
	}
	if cc.ContractBody.CreatorTime == "" {
		logs.Warning("Contract Need CreatorTime!")
		err = errors.New("Contract Need CreatorTime!")
		return err
	}
	if cc.ContractBody.StartTime == "" {
		logs.Warning("Contract Need StartTime!")
		err = errors.New("Contract Need StartTime!")
		return err
	}
	if cc.ContractBody.EndTime == "" {
		logs.Warning("Contract Need EndTime!")
		err = errors.New("Contract Need EndTime!")
		return err
	}
	if cc.ContractBody.ContractOwners == nil || len(cc.ContractBody.ContractOwners) == 0 {
		logs.Warning("Contract Need ContractOwners!")
		err = errors.New("Contract Need ContractOwners!")
		return err
	}
	if cc.ContractBody.ContractAssets == nil || len(cc.ContractBody.ContractAssets) == 0 {
		logs.Warning("Contract Need ContractAssets!")
		err = errors.New("Contract Need ContractAssets!")
		return err
	}
	if cc.ContractBody.ContractSignatures == nil || len(cc.ContractBody.ContractSignatures) == 0 {
		logs.Warning("Contract Need ContractOwners!")
		err = errors.New("Contract Need ContractOwners!")
		return err
	}
	if cc.ContractBody.MetaAttribute == nil {
		cc.ContractBody.MetaAttribute = make(map[string]string, 0)
	}
	//将描述态数据加载成运行态，因此value都是gc.xxxx(描述态的)
	common.AddProperty(cc, cc.PropertyTable, _Cname, cc.ContractBody.Cname)
	cc.ContractBody.Ctype = common.TernaryOperator(cc.ContractBody.Ctype == "", constdef.ComponentType[constdef.Component_Contract], cc.ContractBody.Ctype).(string)
	common.AddProperty(cc, cc.PropertyTable, _Ctype, cc.ContractBody.Ctype)
	common.AddProperty(cc, cc.PropertyTable, _Caption, cc.ContractBody.Caption)
	common.AddProperty(cc, cc.PropertyTable, _Description, cc.ContractBody.Description)
	common.AddProperty(cc, cc.PropertyTable, _MetaAttribute, cc.ContractBody.MetaAttribute)
	common.AddProperty(cc, cc.PropertyTable, _ContractId, cc.ContractBody.ContractId)
	cc.ContractBody.ContractState = common.TernaryOperator(cc.ContractBody.ContractState == "", constdef.ContractState[constdef.Contract_Create], cc.ContractBody.ContractState).(string)
	common.AddProperty(cc, cc.PropertyTable, _ContractState, cc.ContractBody.ContractState)
	common.AddProperty(cc, cc.PropertyTable, _Creator, cc.ContractBody.Creator)
	common.AddProperty(cc, cc.PropertyTable, _CreatorTime, cc.ContractBody.CreatorTime)
	common.AddProperty(cc, cc.PropertyTable, _StartTime, cc.ContractBody.StartTime)
	common.AddProperty(cc, cc.PropertyTable, _EndTime, cc.ContractBody.EndTime)
	common.AddProperty(cc, cc.PropertyTable, _ContractOwners, cc.ContractBody.ContractOwners)
	common.AddProperty(cc, cc.PropertyTable, _NextTasks, cc.ContractBody.NextTasks)
	//自有类型，自己解决添加
	cc.AddProperty(cc, _ContractAssets, cc.ContractBody.ContractAssets)
	cc.AddProperty(cc, _ContractSignatures, cc.ContractBody.ContractSignatures)

	var meta_map map[string]string = make(map[string]string, 0)
	meta_map[_UCVM_Version] = constdef.UCVM_Version
	meta_map[_UCVM_CopyRight] = constdef.UCVM_CopyRight
	meta_map[_UCVM_Date] = constdef.UCVM_Date
	cc.AddMetaAttribute(meta_map)

	//初始化指针变量
	cc.ComponentTable = table.NewComponentTable()
	cc.FunctionParseEngine = function.NewFunctionParseEngine()
	cc.ExpressionParseEngine = expressionutils.NewExpressionParseEngine()

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
	contractOwners_property := cc.PropertyTable[_ContractOwners].(property.PropertyT)
	if contractOwners_property.GetValue() == nil {
		contractOwners_property.SetValue(make([]string, 0))
	}
	if p_owner != "" {
		v_subject_list := contractOwners_property.GetValue().([]string)
		contractOwners_property.SetValue(append(v_subject_list, p_owner))
	}
	cc.PropertyTable[_ContractOwners] = contractOwners_property
	cc.ContractBody.ContractOwners = contractOwners_property.GetValue().([]string)
}
func (cc *CognitiveContract) AddContractAsset(p_asset ContractAsset) {
	contractAssets_property := cc.PropertyTable[_ContractAssets].(property.PropertyT)
	if contractAssets_property.GetValue() != nil {
		contractAssets_property.SetValue(make([]ContractAsset, 0))
	}
	v_asset_list := contractAssets_property.GetValue().([]ContractAsset)
	contractAssets_property.SetValue(append(v_asset_list, p_asset))

	cc.PropertyTable[_ContractAssets] = contractAssets_property
	cc.ContractBody.ContractAssets = contractAssets_property.GetValue().([]ContractAsset)
}
func (cc *CognitiveContract) AddSignature(p_signature ContractSignature) {
	contractSignature_property := cc.PropertyTable[_ContractSignatures].(property.PropertyT)
	if contractSignature_property.GetValue() != nil {
		contractSignature_property.SetValue(make([]ContractSignature, 0))
	}
	v_signature_list := contractSignature_property.GetValue().([]ContractSignature)
	contractSignature_property.SetValue(append(v_signature_list, p_signature))

	cc.PropertyTable[_ContractSignatures] = contractSignature_property
	cc.ContractBody.ContractSignatures = contractSignature_property.GetValue().([]ContractSignature)
}
func (gc *CognitiveContract) AddMetaAttribute(metaProperty interface{}) {
	if metaProperty != nil && len(metaProperty.(map[string]string)) != 0 {
		metaAttribute_property := gc.PropertyTable[_MetaAttribute].(property.PropertyT)
		if metaAttribute_property.GetValue() == nil || len(metaAttribute_property.GetValue().(map[string]string)) == 0 {
			metaAttribute_property.SetValue(make(map[string]string, 0))
		}
		v_metaProperty := metaProperty.(map[string]string)
		for key, value := range v_metaProperty {
			metaAttribute_property.GetValue().(map[string]string)[key] = value
		}
		gc.PropertyTable[_MetaAttribute] = metaAttribute_property
		gc.ContractBody.MetaAttribute = metaAttribute_property.GetValue().(map[string]string)
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
		return cc.PropertyTable[p_name].(property.PropertyT)
	}
	return nil
}

//====属性Get方法
func (gc *CognitiveContract) GetCname() string {
	if gc.PropertyTable[_Cname] == nil {
		return ""
	}
	cname_property := gc.PropertyTable[_Cname].(property.PropertyT)
	return cname_property.GetValue().(string)
}

func (gc *CognitiveContract) GetCaption() string {
	var r_res string = ""
	if gc.PropertyTable[_Caption] == nil {
		r_res = ""
	} else {
		caption_property := gc.PropertyTable[_Caption].(property.PropertyT)
		r_res = caption_property.GetValue().(string)
	}
	return r_res
}

func (gc *CognitiveContract) GetDescription() string {
	var r_res string = ""
	if gc.PropertyTable[_Description] != nil {
		description_property := gc.PropertyTable[_Description].(property.PropertyT)
		r_res = description_property.GetValue().(string)
	}
	return r_res
}

func (gc *CognitiveContract) GetMetaAttribute() map[string]string {
	if gc.PropertyTable[_MetaAttribute] == nil {
		return nil
	}
	metaattribute_property := gc.PropertyTable[_MetaAttribute].(property.PropertyT)
	return metaattribute_property.GetValue().(map[string]string)
}

//属性Set方法
func (cc CognitiveContract) SetId(p_Id string) {
	//Take case: Setter method need set value for gc.xxxxxx
	cc.Id = p_Id
	id_property := cc.PropertyTable[_Id].(property.PropertyT)
	id_property.SetValue(p_Id)
	//Take case: Setter method need set value for gc.PropertyTable[xxxx]
	cc.PropertyTable[_Id] = id_property
}
func (gc *CognitiveContract) SetCname(str_name string) {
	//Take case: Setter method need set value for gc.xxxxxx
	gc.ContractBody.Cname = str_name
	cname_property := gc.PropertyTable[_Cname].(property.PropertyT)
	cname_property.SetValue(str_name)
	//Take case: Setter method need set value for gc.PropertyTable[xxxx]
	gc.PropertyTable[_Cname] = cname_property
}

func (cc *CognitiveContract) GetContractState() string {
	state_property := cc.PropertyTable[_ContractState].(property.PropertyT)
	return state_property.GetValue().(string)
}
func (cc *CognitiveContract) GetCreator() string {
	creator_property := cc.PropertyTable[_Creator].(property.PropertyT)
	return creator_property.GetValue().(string)
}
func (cc *CognitiveContract) GetCreatorTime() string {
	creatorTime_property := cc.PropertyTable[_CreatorTime].(property.PropertyT)
	return creatorTime_property.GetValue().(string)
}
func (cc *CognitiveContract) GetStartTime() string {
	startTime_property := cc.PropertyTable[_StartTime].(property.PropertyT)
	return startTime_property.GetValue().(string)
}
func (cc *CognitiveContract) GetEndTime() string {
	endTime_property := cc.PropertyTable[_EndTime].(property.PropertyT)
	return endTime_property.GetValue().(string)
}
func (cc *CognitiveContract) GetContractOwners() interface{} {
	contractOwners_property := cc.PropertyTable[_ContractOwners].(property.PropertyT)
	return contractOwners_property.GetValue().([]string)
}
func (cc *CognitiveContract) GetContractAssets() interface{} {
	contractAssets_property := cc.PropertyTable[_ContractAssets].(property.PropertyT)
	return contractAssets_property.GetValue().([]ContractAsset)
}
func (cc *CognitiveContract) GetContractSignatures() interface{} {
	contractSignatures := cc.PropertyTable[_ContractSignatures].(property.PropertyT)
	return contractSignatures.GetValue().([]ContractSignature)
}
func (cc *CognitiveContract) GetContractComponents() []interface{} {
	return cc.ContractBody.ContractComponents
}
func (cc *CognitiveContract) GetNextTasks() []string {
	nexttasks_property := cc.PropertyTable[_NextTasks].(property.PropertyT)
	return nexttasks_property.GetValue().([]string)
}

//====属性Set方法
func (gc *CognitiveContract) SetCtype(str_type string) {
	gc.ContractBody.Ctype = str_type
	ctype_property := gc.PropertyTable[_Ctype].(property.PropertyT)
	ctype_property.SetValue(str_type)
	gc.PropertyTable[_Ctype] = ctype_property
}

func (gc *CognitiveContract) SetCaption(str_Caption string) {
	gc.ContractBody.Caption = str_Caption
	caption_property := gc.PropertyTable[_Caption].(property.PropertyT)
	caption_property.SetValue(str_Caption)
	gc.PropertyTable[_Caption] = caption_property
}

func (gc *CognitiveContract) SetDescription(str_Description string) {
	gc.ContractBody.Description = str_Description
	description_property := gc.PropertyTable[_Description].(property.PropertyT)
	description_property.SetValue(str_Description)
	gc.PropertyTable[_Description] = description_property
}

func (gc *CognitiveContract) SetMetaAttribute(p_metaAttribute map[string]string) {
	gc.ContractBody.MetaAttribute = p_metaAttribute
	metaAttribute_property := gc.PropertyTable[_MetaAttribute].(property.PropertyT)
	metaAttribute_property.SetValue(p_metaAttribute)
	gc.PropertyTable[_MetaAttribute] = metaAttribute_property
}

func (cc *CognitiveContract) SetContractId(p_ConstractId string) {
	cc.ContractBody.ContractId = p_ConstractId
	contractid_property := cc.PropertyTable[_ContractId].(property.PropertyT)
	contractid_property.SetValue(p_ConstractId)
	cc.PropertyTable[_ContractId] = contractid_property
}
func (cc *CognitiveContract) SetContractState(p_State string) {
	cc.ContractBody.ContractState = p_State
	state_property := cc.PropertyTable[_ContractState].(property.PropertyT)
	state_property.SetValue(p_State)
	cc.PropertyTable[_ContractState] = state_property
}
func (cc *CognitiveContract) SetCreator(p_Creator string) {
	cc.ContractBody.Creator = p_Creator
	creator_property := cc.PropertyTable[_Creator].(property.PropertyT)
	creator_property.SetValue(p_Creator)
	cc.PropertyTable[_Creator] = creator_property
}
func (cc *CognitiveContract) SetCreatorTime(p_CreatorTime string) {
	cc.ContractBody.CreatorTime = p_CreatorTime
	creatortime_property := cc.PropertyTable[_CreatorTime].(property.PropertyT)
	creatortime_property.SetValue(p_CreatorTime)
	cc.PropertyTable[_CreatorTime] = creatortime_property
}
func (cc *CognitiveContract) SetStartTime(p_StartTime string) {
	cc.ContractBody.StartTime = p_StartTime
	starttime_property := cc.PropertyTable[_StartTime].(property.PropertyT)
	starttime_property.SetValue(p_StartTime)
	cc.PropertyTable[_StartTime] = starttime_property
}
func (cc *CognitiveContract) SetEndTime(p_EndTime string) {
	cc.ContractBody.EndTime = p_EndTime
	endtime_property := cc.PropertyTable[_EndTime].(property.PropertyT)
	endtime_property.SetValue(p_EndTime)
	cc.PropertyTable[_EndTime] = endtime_property
}
func (cc *CognitiveContract) SetContractOwners(p_ContractOwners []string) {
	cc.ContractBody.ContractOwners = p_ContractOwners
	contractowners_property := cc.PropertyTable[_ContractOwners].(property.PropertyT)
	contractowners_property.SetValue(p_ContractOwners)
	cc.PropertyTable[_ContractOwners] = contractowners_property
}
func (cc *CognitiveContract) SetContractAssets(p_ContractAssets []ContractAsset) {
	cc.ContractBody.ContractAssets = p_ContractAssets
	contractassets_property := cc.PropertyTable[_ContractAssets].(property.PropertyT)
	contractassets_property.SetValue(p_ContractAssets)
	cc.PropertyTable[_ContractAssets] = contractassets_property
}
func (cc *CognitiveContract) SetContractSignatures(p_ContractSignatures []ContractSignature) {
	cc.ContractBody.ContractSignatures = p_ContractSignatures
	contractsignatures_property := cc.PropertyTable[_ContractSignatures].(property.PropertyT)
	contractsignatures_property.SetValue(p_ContractSignatures)
	cc.PropertyTable[_ContractSignatures] = contractsignatures_property
}
func (cc *CognitiveContract) SetNextTasks(p_NextTasks []string) {
	cc.ContractBody.NextTasks = p_NextTasks
	nexttasks_property := cc.PropertyTable[_NextTasks].(property.PropertyT)
	nexttasks_property.SetValue(p_NextTasks)
	cc.PropertyTable[_NextTasks] = nexttasks_property
}

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
func (cc *CognitiveContract) UpdateTasksState() (int8, error) {
	var r_ret int8 = -1
	var r_err error = nil
	var next_tasks []string = cc.GetNextTasks()
	fmt.Println(next_tasks)
	if next_tasks == nil || len(next_tasks) == 0 {
		r_err = errors.New("contract has no start tasks!")
		return r_ret, r_err
	}
	var r_task_queue *common.Queue = common.NewQueue()
	for _, v_task := range next_tasks {
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
			return r_ret, r_err
		}
		if f_f_task.(inf.ITask).GetState() == constdef.TaskState[constdef.TaskState_Discard] {
			for r_task_queue.Len() != 0 {
				r_task_queue.Pop()
			}
			next_tasks = f_f_task.(inf.ITask).GetNextTasks()
			for _, t_task := range next_tasks {
				r_task_queue.Push(t_task)
			}
			continue
		} else if f_f_task.(inf.ITask).GetState() == constdef.TaskState[constdef.TaskState_In_Progress] || f_f_task.(inf.ITask).GetState() == constdef.TaskState[constdef.TaskState_Completed] {
			for r_task_queue.Len() != 0 {
				r_task_queue.Pop()
			}
			r_task_queue.Push(f_f_task.(inf.ITask).GetName())
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
			return r_ret, r_err
		}
		r_ret, f_err = f_s_task.(inf.ITask).UpdateState()
		switch r_ret {
		case 1: //执行成功后，合约退出， 注意：后续任务不入队列了，等待共识成功后初始化到扫描监控表中，下次加载再执行
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
		case 0: //执行条件不成立
			if f_s_task.(inf.ITask).GetState() == constdef.TaskState[constdef.TaskState_Dormant] { //继续判断同级中的下一任务
				continue
			} else { //合约退出
				r_err = errors.New("task[" + f_s_task.(inf.ITask).GetName() + "] condition not fullfill!")
				break
			}
		case -1: //执行失败后，合约退出
			r_err = errors.New("task[" + f_s_task.(inf.ITask).GetName() + "] execute fail!")
			break
		}
		if f_err != nil {
			//TODO log
			fmt.Println("Error: [" + f_err.Error() + "]")
		}
	}
	return r_ret, r_err
}
