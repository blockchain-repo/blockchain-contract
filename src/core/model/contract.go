package model

import (
	"encoding/json"
	"unicontract/src/common"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/config"
	"unicontract/src/core/protos"
)

type ContractSignature struct {
	OwnerPubkey string `json:"OwnerPubkey"`
	Signature   string `json:"Signature"`
	// 13 bit
	SignTimestamp string `json:"SignTimestamp"`
}

type ContractAsset struct {
	AssetId     string            `json:"AssetId"`
	Name        string            `json:"Name"`
	Caption     string            `json:"Caption"`
	Description string            `json:"Description"`
	Unit        string            `json:"Unit"`
	Amount      float32           `json:"Amount"`
	MetaData    map[string]string `json:"MetaData"`
}

type ExpressionResult struct {
	Message string `json:"Message"`
	Code    int32  `json:"Code"`
	Data    string `json:"Data"`
	OutPut  string `json:"OutPut"`
}

type ComponentsExpression struct {
	Cname            string            `json:"Cname"`
	Ctype            string            `json:"Ctype"`
	Caption          string            `json:"Caption"`
	Description      string            `json:"Description"`
	ExpressionStr    string            `json:"ExpressionStr"`
	ExpressionResult *ExpressionResult `json:"ExpressionResult"`
	LogicValue       int32             `json:"LogicValue"`
	MetaAttribute    map[string]string `json:"MetaAttribute"`
}

type ComponentDataSub struct {
	Cname        string           `json:"Cname"`
	Ctype        string           `json:"Ctype"`
	Caption      string           `json:"Caption"`
	Description  string           `json:"Description"`
	ModifyDate   string           `json:"ModifyDate"`
	HardConvType string           `json:"HardConvType"`
	Category     []string         `json:"Category"`
	Mandatory    bool             `json:"Mandatory"`
	Unit         string           `json:"Unit"`
	Options      map[string]int32 `json:"Options"`
	Format       string           `json:"Format"`
	// Value interface{} int64
	ValueInt int32 `json:"ValueInt"`
	// Value interface{} unit64
	ValueUint uint32 `json:"ValueUint"`
	// Value interface{} float64
	ValueFloat float64 `json:"ValueFloat"`
	// Value interface{} string
	ValueString string `json:"ValueString"`
	// DefaultValue interface{} int64
	DefaultValueInt int32 `json:"DefaultValueInt"`
	// DefaultValue interface{} unit64
	DefaultValueUint uint32 `json:"DefaultValueUint"`
	// DefaultValue interface{} float64
	DefaultValueFloat float64 `json:"DefaultValueFloat"`
	// DefaultValueinterface{} string
	DefaultValueString string `json:"DefaultValueString"`
	// DataRange interface{} int64
	DataRangeInt []int32 `json:"DataRangeInt"`
	// DataRange interface{} unit64
	DataRangeUint []uint32 `json:"DataRangeUint"`
	// DataRange interface{} float64
	DataRangeFloat []float64 `json:"DataRangeFloat"`
}

type ComponentData struct {
	Cname        string           `json:"Cname"`
	Ctype        string           `json:"Ctype"`
	Caption      string           `json:"Caption"`
	Description  string           `json:"Description"`
	ModifyDate   string           `json:"ModifyDate"`
	HardConvType string           `json:"HardConvType"`
	Category     []string         `json:"Category"`
	Mandatory    bool             `json:"Mandatory"`
	Unit         string           `json:"Unit"`
	Options      map[string]int32 `json:"Options"`
	Format       string           `json:"Format"`
	// Value interface{} int64
	ValueInt int32 `json:"ValueInt"`
	// Value interface{} unit64
	ValueUint uint32 `json:"ValueUint"`
	// Value interface{} float64
	ValueFloat float64 `json:"ValueFloat"`
	// Value interface{} string
	ValueString string `json:"ValueString"`
	// DefaultValue interface{} int64
	DefaultValueInt int32 `json:"DefaultValueInt"`
	// DefaultValue interface{} unit64
	DefaultValueUint uint32 `json:"DefaultValueUint"`
	// DefaultValue interface{} float64
	DefaultValueFloat float64 `json:"DefaultValueFloat"`
	// DefaultValueinterface{} string
	DefaultValueString string `json:"DefaultValueString"`
	// DataRange interface{} int64
	DataRangeInt []int32 `json:"DataRangeInt"`
	// DataRange interface{} unit64
	DataRangeUint []uint32 `json:"DataRangeUint"`
	// DataRange interface{} float64
	DataRangeFloat []float64 `json:"DataRangeFloat"`
	// add 2017-06-26
	Value            string `json:"Value"`
	DefaultValue     string `json:"DefaultValue"`
	ValueBool        bool   `json:"ValueBool"`
	DefaultValueBool bool   `json:"DefaultValueBool"`
}

type SelectBranchExpression struct {
	BranchExpressionStr   string `json:"BranchExpressionStr"`
	BranchExpressionValue string `json:"BranchExpressionValue"`
}

type ContractComponentSub struct {
	Cname                         string                  `json:"Cname"`
	Ctype                         string                  `json:"Ctype"`
	Caption                       string                  `json:"Caption"`
	Description                   string                  `json:"Description"`
	State                         string                  `json:"State"`
	PreCondition                  []*ComponentsExpression `json:"PreCondition"`
	CompleteCondition             []*ComponentsExpression `json:"CompleteCondition"`
	DiscardCondition              []*ComponentsExpression `json:"DiscardCondition"`
	NextTasks                     []string                `json:"NextTasks"`
	DataList                      []*ComponentData        `json:"DataList"`
	DataValueSetterExpressionList []*ComponentsExpression `json:"DataValueSetterExpressionList"`
	TaskList                      []string                `json:"TaskList"`
	SupportArguments              []string                `json:"SupportArguments"`
	AgainstArguments              []string                `json:"AgainstArguments"`
	Text                          []string                `json:"Text"`
	// add date: 2017-05-11 任务执行索引次数 int
	TaskExecuteIdx int32  `json:"TaskExecuteIdx"`
	TaskId         string `json:"TaskId"`
	// 2017-05-27 17:10:00 add
	SelectBranches []*SelectBranchExpression `json:"SelectBranches"`
	Result         int32                     `json:"Result"`
	SupportNum     int32                     `json:"SupportNum"`
	AgainstNum     int32                     `json:"AgainstNum"`
}

type ContractComponent struct {
	Cname                         string                  `json:"Cname"`
	Ctype                         string                  `json:"Ctype"`
	Caption                       string                  `json:"Caption"`
	Description                   string                  `json:"Description"`
	State                         string                  `json:"State"`
	PreCondition                  []*ComponentsExpression `json:"PreCondition"`
	CompleteCondition             []*ComponentsExpression `json:"CompleteCondition"`
	DiscardCondition              []*ComponentsExpression `json:"DiscardCondition"`
	NextTasks                     []string                `json:"NextTasks"`
	DataList                      []*ComponentData        `json:"DataList"`
	DataValueSetterExpressionList []*ComponentsExpression `json:"DataValueSetterExpressionList"`
	CandidateList                 []*ContractComponentSub `json:"CandidateList"`
	TaskList                      []string                `json:"TaskList"`
	// add date: 2017-05-11 任务执行索引次数 int
	TaskExecuteIdx int32  `json:"TaskExecuteIdx"`
	TaskId         string `json:"TaskId"`
	// 2017-05-27 17:10:00 add
	SelectBranches []*SelectBranchExpression `json:"SelectBranches"`
	MetaAttribute  map[string]string         `json:"MetaAttribute"`
}

type ContractBody struct {
	ContractId         string               `json:"ContractId"`
	Cname              string               `json:"Cname"`
	Ctype              string               `json:"Ctype"`
	Caption            string               `json:"Caption"`
	Description        string               `json:"Description"`
	ContractState      string               `json:"ContractState"`
	Creator            string               `json:"Creator"`
	CreateTime         string               `json:"CreateTime"`
	StartTime          string               `json:"StartTime"`
	EndTime            string               `json:"EndTime"`
	ContractOwners     []string             `json:"ContractOwners"`
	ContractAssets     []*ContractAsset     `json:"ContractAssets"`
	ContractSignatures []*ContractSignature `json:"ContractSignatures"`
	ContractComponents []*ContractComponent `json:"ContractComponents"`
	// add date: 2017-05-11 map[string]interface{} 合约属性MetaData
	//    bytes MetaAttribute = 15;
	MetaAttribute      map[string]string `json:"MetaAttribute"`
	NextTasks          []string          `json:"NextTasks"`
	ContractProductId  string            `json:"ContractProductId"`
	ContractTemplateId string            `json:"ContractTemplateId"`
}

type ContractHead struct {
	MainPubkey string `json:"MainPubkey"`
	Version    int32  `json:"Version"`
	// 指派处理时间 add 2017-05-27 17:10:0
	AssignTime string `json:"AssignTime"`
	// add date: 2017-05-11 合约执行时间戳
	// 操作时间,记录状态改变时间, Timestamp修改而来 2017-05-27 17:10:0
	OperateTime string `json:"OperateTime"`
	// add date: 2017-06-01 共识结果,是否共识等
	ConsensusResult int32 `json:"ConsensusResult"`
}

// table [Contracts]
type ContractModel struct {
	Id           string        `json:"id"`
	ContractHead *ContractHead `json:"ContractHead"`
	ContractBody *ContractBody `json:"ContractBody"`
}

//// table [Contracts]
//type ContractModel struct {
//	protos.Contract //合约描述集合, (引用contract描述 for proto3)
//}

// validate the contract
func (c *ContractModel) Validate() bool {
	// 1. validate contract.id
	idValid := c.Id == c.GenerateId() // Hash contractBody
	uniledgerlog.Warn("valid id is: " + c.GenerateId())
	if !idValid {
		uniledgerlog.Error("Validate idValid false")
		return false
	}

	signatureValid := c.IsSignatureValid()
	if !signatureValid {
		return false
	}

	return true
}

//Create a signature for the ContractBody
func (c *ContractModel) Sign(private_key string) string {
	/*-------------module deep copy start --------------*/
	var contractBodyClone = c.ContractBody

	// new obj
	var temp protos.ContractBody

	contractBodyCloneBytes, _ := json.Marshal(contractBodyClone)
	err := json.Unmarshal(contractBodyCloneBytes, &temp)
	if err != nil {
		uniledgerlog.Error("Unmarshal error ", err)
	}
	temp.ContractSignatures = nil
	contractBodySerialized := common.StructSerialize(temp)
	/*-------------module deep copy end --------------*/

	signatureContractBody := common.Sign(private_key, contractBodySerialized)
	return signatureContractBody
}

// Check the validity of a ContractBody's signature
func (c *ContractModel) IsSignatureValid() bool {

	/*-------------module deep copy start --------------*/
	var contractBodyClone = c.ContractBody

	// new obj
	var temp protos.ContractBody

	contractBodyCloneCloneBytes, _ := json.Marshal(contractBodyClone)
	err := json.Unmarshal(contractBodyCloneCloneBytes, &temp)
	if err != nil {
		uniledgerlog.Error("[module-model]IsSignatureValid error ", err)
	}
	temp.ContractSignatures = nil
	contractBody_serialized := common.StructSerialize(temp)
	/*-------------module deep copy end --------------*/
	contractOwners := c.ContractBody.ContractOwners
	// 合约 owners 不能存在重复的
	len_contractOwners := len(contractOwners)
	if len_contractOwners == 0 {
		uniledgerlog.Error("IsSignatureValid len_contractOwners 长度不能为0")
		return false
	}
	contractOwnersSet := common.StrArrayToHashSet(c.ContractBody.ContractOwners)
	if len_contractOwners != contractOwnersSet.Len() {
		uniledgerlog.Error("IsSignatureValid contractOwners 存在重复项")
		return false
	}
	contractSignatures := c.ContractBody.ContractSignatures
	for _, contractSignature := range contractSignatures {

		ownerPubkey := contractSignature.OwnerPubkey
		if !contractOwnersSet.Has(ownerPubkey) {
			uniledgerlog.Error("IsSignatureValid contractOwner ", ownerPubkey, " 不存在于", contractOwners)
			return false
		}
		if contractSignature.SignTimestamp == "" {
			uniledgerlog.Error("IsSignatureValid SignTimestamp is blank")
			return false
		}
		signature := contractSignature.Signature
		if signature == "" {
			uniledgerlog.Error("IsSignatureValid signature is blank")
			return false
		}
		// contract signature verify
		verifyFlag := common.Verify(ownerPubkey, contractBody_serialized, signature)
		//uniledgerlog.Debug("contract verify[owner:", ownerPubkey, ",signature:", signature, "contractBody", contractBody_serialized, "]\n", verifyFlag)
		if !verifyFlag {
			uniledgerlog.Error("IsSignatureValid contract signature verify fail")
			return false
		}
	}

	return true
}

func (c *ContractModel) ToString() string {
	return common.StructSerialize(c)
}

// return the  id (hash generate)
func (c *ContractModel) GenerateId() string {
	contractBodySerialized := common.StructSerialize(c.ContractBody)
	//uniledgerlog.Warn("contractBodySerialized:\n", contractBodySerialized)
	return common.HashData(contractBodySerialized)
}

//Validate the contract header
func (c *ContractModel) validateContractHead() bool {

	pub_keys := config.GetAllPublicKey()
	pub_keysSet := common.StrArrayToHashSet(pub_keys)
	contractHead := c.ContractHead
	if contractHead.MainPubkey == "" {
		uniledgerlog.Error("contract main_pubkey blank")
		return false
	}

	if !pub_keysSet.Has(contractHead.MainPubkey) {
		uniledgerlog.Warn("main_pubkey ", contractHead.MainPubkey, " not in pubkeys")
		return false
	}
	return true
}
