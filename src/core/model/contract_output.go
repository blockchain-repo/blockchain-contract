package model

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"unicontract/src/common"
	"unicontract/src/config"
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

// 合约&交易关系信息
type Relaction struct {
	ContractId string
	TaskId     string
	Voters     []string
	Votes      []*Vote
}

type Transaction struct {
	Asset         *Asset            `json:"asset"`
	Conditions    []*ConditionsItem `json:"conditions"`
	Fulfillments  []*Fulfillment    `json:"fulfillments"`
	Metadata      *Metadata         `json:"metadata"`
	Operation     string            `json:"operation"`
	Timestamp     string            `json:"timestamp"`
	Relaction     *Relaction
	ContractModel ContractModel `json:"Contract"` //合约描述集合, (引用contract描述 for proto3)
}

// table [ContractOutputs]
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
	//temp.Relaction.Signatures = nil
	contract_without_signatures_serialized := common.Serialize(temp)

	return common.HashData(contract_without_signatures_serialized)
}

// judge has enough votes for ContractOutput
func (c *ContractOutput) HasEnoughVotes() bool {
	voters := c.Transaction.Relaction.Voters
	votes := c.Transaction.Relaction.Votes
	voters_len := len(voters)

	if voters_len <= 0 {
		return false
	}

	if len(votes) != voters_len {
		return false
	}

	var invalid_signature_len int
	for index, voter := range voters {
		if voter == "" {
			return false
		}

		vote := votes[index]

		ContractOutputNodePubkey := vote.NodePubkey
		signature := vote.Signature

		if ContractOutputNodePubkey == "" {
			invalid_signature_len++
		} else {
			if voter != ContractOutputNodePubkey || signature == "" {
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
	votes := c.Transaction.Relaction.Votes
	voters_len := len(voters)

	/*----------------keyring----------------*/
	pub_keys := config.GetAllPublicKey()
	pub_keysSet := common.StrArrayToHashSet(pub_keys)

	validSignCount := 0
	for index, voter := range voters {
		if !pub_keysSet.Has(voter) {
			continue
		}
		vote := votes[index]
		nodePubkey := vote.NodePubkey
		if nodePubkey != voter {
			continue
		}
		nodeVoteSignature := vote.Signature
		if c.validateSignature(nodePubkey, nodeVoteSignature) {
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

	temp.Relaction.Votes = nil
	tempNoSign := common.Serialize(temp)
	return common.Verify(nodePubkey, tempNoSign, nodeSignature)
}
