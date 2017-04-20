package model

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"unicontract/src/common"
	"unicontract/src/core/protos"
)

// table [contract_output]
type ContractOutput struct {
	Id          string             `json:"id"`          //ContractOutput.Id
	Transaction protos.Transaction `json:"transaction"` //ContractOutput.Transaction
	Version     int16              `json:"version"`     //ContractOutput.Version
}

func (c *ContractOutput) ToString() string {
	return common.Serialize(c)
}

// return the id (hash generate)
func (c *ContractOutput) GetId() string {

	/*-------------module deep copy--------------*/
	var transactionClone = c.Transaction

	// new obj
	var temp protos.Transaction

	transactionCloneBytes, _ := json.Marshal(transactionClone)
	err := json.Unmarshal(transactionCloneBytes, &temp)
	if err != nil {
		beego.Error("Unmarshal error ", err)
	}
	beego.Debug(temp)
	temp.Relaction.Signatures = nil
	beego.Debug(temp)
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
		_ = signature

		if contract_node_pubkey == "" {
			invalid_signature_len += 1
		} else {
			if voter != contract_node_pubkey || signature == "" {
				invalid_signature_len += 1
			}
		}

	}

	if invalid_signature_len >= voters_len/2 {
		return false
	}
	return true
}

// 判断hash，hash不包括voters的signatures
func (c *ContractOutput) ValidateHash() bool {
	hashId := c.Id
	rightId := c.GetId()
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
	for index, voter := range voters{
		relationSignature := signatures[index]
		nodePubkey := relationSignature.ContractNodePubkey
		if nodePubkey != voter {
			continue
		}
		nodeSignature := relationSignature.Signature
		if c.validateSignature(nodePubkey,nodeSignature) {
			validSignCount++
		}
	}

	if validSignCount <= voters_len/2 {
		return false
	}
	return true
}

func (c *ContractOutput) validateSignature(nodePubkey string,nodeSignature string) bool{
	var ConOutTxClone = c.Transaction

	//deep copy
	var temp protos.Transaction
	txCloneBytes, _ := json.Marshal(ConOutTxClone)
	err := json.Unmarshal(txCloneBytes, &temp)
	if err != nil {
		beego.Error("Unmarshal error ", err)
	}

	temp.Relaction.Signatures = nil
	tempNoSign := common.Serialize(temp)
	return common.Verify(nodePubkey,tempNoSign,nodeSignature)
}
