package model

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"unicontract/src/common"
)

// An Asset is a fungible unit to spend and lock with Transactions
type Asset struct {
	Data       interface{} `json:"data"`
	Id         string      `json:"id"`
	Divisible  bool        `json:"divisible"`
	Refillable bool        `json:"refillable"`
	Updatable  bool        `json:"updatable"`
}

type ConditionDetails struct {
	Bitmask   int32  `json:"bitmask"`
	PublicKey string `json:"public_key"`
	Signature string `json:"signature"`
	Type      string `json:"type"`
	TypeId    int32  `json:"type_id"`
}

type Condition struct {
	Details *ConditionDetails `json:"details"`
	Uri     string            `json:"uri"`
}

type ConditionsItem struct {
	Amount      int64      `json:"amount"`
	Cid         int32      `json:"cid"`
	Condition   *Condition `json:"condition"`
	OwnersAfter []string   `json:"owners_after"`
	Isfreeze    bool       `json:"isfreeze"`
}

type Fulfillment struct {
	Fid          int32       `json:"fid"`
	OwnersBefore []string    `json:"owners_before"`
	Fulfillment  string      `json:"fulfillment"`
	Input        interface{} `json:"input"`
}

type Metadata struct {
	Id   string      `json:"id"`
	Data interface{} `json:"data"`
}

type RelactionSignature struct {
	ContractNodePubkey string `json:"contract_node_pubkey"`
	Signature          string `json:"signature"`
}

// 合约&交易关系信息
type Relaction struct {
	ContractId string                `json:"contract_id"`
	TaskId     string                `json:"task_id"`
	Voters     []string              `json:"voters"`
	Signatures []*RelactionSignature `json:"signatures"`
}

type Transaction struct {
	Asset         *Asset            `json:"asset"`
	Conditions    []*ConditionsItem `json:"conditions"`
	Fulfillments  []*Fulfillment    `json:"fulfillments"`
	Metadata      *Metadata         `json:"metadata"`
	Operation     string            `json:"operation"`
	Timestamp     string            `json:"timestamp"`
	Relaction     *Relaction        `json:"relaction"`
	ContractModel ContractModel     `json:"contracts"` //合约描述集合, (引用contract描述 for proto3)
}

// table [contract_output]
type ContractOutput struct {
	Id          string      `json:"id"`          //ContractOutput.Id
	Transaction Transaction `json:"transaction"` //ContractOutput.Transaction
	Version     int16       `json:"version"`     //ContractOutput.Version
}

func (c *ContractOutput) ToString() string {
	return common.Serialize(c)
}

// return the id (hash generate)
func (c *ContractOutput) GenerateId() string {

	/*-------------module deep copy--------------*/
	var transactionClone = c.Transaction

	// new obj
	var temp Transaction

	transactionCloneBytes, _ := json.Marshal(transactionClone)
	err := json.Unmarshal(transactionCloneBytes, &temp)
	if err != nil {
		beego.Error("Unmarshal error ", err)
	}
	//TODO deal with the timestamps
	temp.Relaction.Signatures = nil
	contract_without_signatures_serialized := common.Serialize(temp)

	return common.HashData(contract_without_signatures_serialized)
}

// judge has enough votes for ContractOutput
func (c *ContractOutput) HasEnoughVotes() bool {
	voters := c.Transaction.Relaction.Voters
	signatures := c.Transaction.Relaction.Signatures
	voters_len := len(voters)

	if voters_len <= 0 {
		return false
	}

	if len(signatures) != voters_len {
		return false
	}

	var invalid_signature_len int
	for index, voter := range voters {
		if voter == "" {
			return false
		}

		_signature := signatures[index]

		contract_node_pubkey := _signature.ContractNodePubkey
		signature := _signature.Signature

		if contract_node_pubkey == "" {
			invalid_signature_len++
		} else {
			if voter != contract_node_pubkey || signature == "" {
				invalid_signature_len++
			}
		}
	}

	if invalid_signature_len >= (voters_len+1)/2 {
		return false
	}
	return true
}

// 判断hash，hash不包括voters的signatures
func (c *ContractOutput) ValidateHash() bool {
	hashId := c.Id
	rightId := c.GenerateId()
	if hashId != rightId {
		return false
	}
	return true
}

//  判断是否有>1/2的有效签名。 return bool
func (c *ContractOutput) ValidateContractOutput() bool {
	voters := c.Transaction.Relaction.Voters
	signatures := c.Transaction.Relaction.Signatures
	voters_len := len(voters)

	validSignCount := 0
	for index, voter := range voters {
		relationSignature := signatures[index]
		nodePubkey := relationSignature.ContractNodePubkey
		if nodePubkey != voter {
			continue
		}
		nodeSignature := relationSignature.Signature
		if c.validateSignature(nodePubkey, nodeSignature) {
			validSignCount++
		}
	}

	if validSignCount <= voters_len/2 {
		return false
	}
	return true
}

func (c *ContractOutput) validateSignature(nodePubkey string, nodeSignature string) bool {
	var ConOutTxClone = c.Transaction

	//deep copy
	var temp Transaction
	txCloneBytes, _ := json.Marshal(ConOutTxClone)
	err := json.Unmarshal(txCloneBytes, &temp)
	if err != nil {
		beego.Error("Unmarshal error ", err)
	}

	temp.Relaction.Signatures = nil
	tempNoSign := common.Serialize(temp)
	return common.Verify(nodePubkey, tempNoSign, nodeSignature)
}
