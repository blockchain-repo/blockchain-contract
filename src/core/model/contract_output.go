package model

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
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

type Metadata struct {
	Id   string                 `json:"id"`
	Data map[string]interface{} `json:"data"`
}


type Transaction struct {
	Asset         *Asset            `json:"asset"`
	Conditions    []*ConditionsItem `json:"conditions"`
	Fulfillments  []*Fulfillment    `json:"fulfillments"`
	Metadata      *Metadata         `json:"metadata"`
	Operation     string            `json:"operation"`
	Timestamp     string            `json:"timestamp"`
	Relation      *Relation         `json:"Relation"`
	ContractModel ContractModel     `json:"Contract"` //合约描述集合, (引用contract描述 for proto3)
}

// table [ContractOutputs]
type ContractOutput struct {
	Id          string      `json:"id,omitempty"` //ContractOutput.Id
	Transaction Transaction `json:"transaction"`  //ContractOutput.Transaction
	Version     int         `json:"version"`      //ContractOutput.Version
}

func (c *ContractOutput) ToString() string {
	return common.Serialize(c)
}

// return the id (hash generate)
func (c *ContractOutput) GenerateId() string {

	/*-------------module deep copy--------------*/
	var contractOutput = c

	// new obj
	var temp ContractOutput

	transactionCloneBytes, _ := json.Marshal(contractOutput)
	err := json.Unmarshal(transactionCloneBytes, &temp)
	if err != nil {
		logs.Error("Unmarshal error ", err)
		return ""
	}
	//logs.Info(common.Serialize(temp))
	//operation := c.Transaction.Operation
	temp.Id = ""
	temp.Transaction.Relation.Votes = nil
	temp.Transaction.ContractModel.ContractHead = nil
	temp.Transaction.Timestamp = ""
	temp.RemoveSignature()
	serializeStr := common.StructSerialize(temp)
	logs.Info("before-sign--",serializeStr)
	return common.HashData(serializeStr)
}

func (c *ContractOutput) RemoveSignature() {
	for _, fulfillment := range c.Transaction.Fulfillments {
		fulfillment.Fulfillment = nil
	}
}

// judge has enough votes for ContractOutput
func (c *ContractOutput) HasEnoughVotes() bool {
	voters := c.Transaction.Relation.Voters
	votes := c.Transaction.Relation.Votes
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
	//operation := c.Transaction.Operation
	var hashId string
	//if operation == "CONTRACT" {
	//	hashId = c.Transaction.ContractModel.Id
	//} else {
	//
	//}
	hashId = c.Id
	rightId := c.GenerateId()

	if hashId != rightId {
		return false
	}
	return true
}

//  判断是否有>1/2的有效签名。 return bool
func (c *ContractOutput) ValidateContractOutput() bool {
	voters := c.Transaction.Relation.Voters
	votes := c.Transaction.Relation.Votes
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
		var signData string
		operation := c.Transaction.Operation
		if operation == "CONTRACT" {
			signData = common.StructSerialize(vote.VoteBody)
		} else {
			signData = vote.VoteBody.VoteFor
		}
		//logs.Info("signData:",signData)
		if common.Verify(nodePubkey, signData, nodeVoteSignature) {
			validSignCount++
		}
	}

	if validSignCount <= voters_len/2 {
		return false
	}
	return true
}

/*
type Transaction struct {
	Asset         *Asset            `json:"asset"`
	Conditions    []*ConditionsItem `json:"conditions"`
	Fulfillments  []*Fulfillment    `json:"fulfillments"`
	Metadata      *Metadata         `json:"metadata"`
	Operation     string            `json:"operation"`
	Timestamp     string            `json:"timestamp"`
	Relation     *Relation		`json:"Relation"`
	ContractModel ContractModel `json:"Contract"` //合约描述集合, (引用contract描述 for proto3)
}
*/

func (c *ContractOutput)GenerateConOutput(operation string, asset Asset, inputs []*Fulfillment, outputs []*ConditionsItem, metadata Metadata, timestamp string, version int, relation Relation, contract ContractModel) {
	tx := Transaction{
		Asset:         &asset,
		Conditions:    outputs,
		Fulfillments:  inputs,
		Metadata:      &metadata,
		Operation:     operation,
		Timestamp:     timestamp,
		Relation:      &relation,
		ContractModel: contract,
	}
	c.Transaction = tx
	c.Version = version
	conOutId := c.GenerateId()
	c.Id = conOutId
}
