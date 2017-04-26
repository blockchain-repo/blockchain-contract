package model

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/protos"
)

// table [Contracts]
type ContractModel struct {
	protos.Contract //合约描述集合, (引用contract描述 for proto3)
}

// validate the contract
func (c *ContractModel) Validate() bool {
	// 1. validate contract.id
	idValid := c.Contract.Id == c.GenerateId() // Hash contractBody
	beego.Debug("model.contract.go Validate, contract.id:",c.Contract.Id,"generateId:", c.GenerateId())
	if !idValid {
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
		beego.Error("Unmarshal error ", err)
	}
	temp.ContractSignatures = nil
	contractBodySerialized := common.Serialize(temp)
	/*-------------module deep copy end --------------*/

	signatureContractBody := common.Sign(private_key, contractBodySerialized)
	return signatureContractBody
}

// Check the validity of a ContractBody's signature
func (c *ContractModel) IsSignatureValid() bool {

	/*-------------module deep copy start --------------*/
	var contractBodyClone = c.Contract.ContractBody

	// new obj
	var temp protos.ContractBody

	contractBodyCloneCloneBytes, _ := json.Marshal(contractBodyClone)
	err := json.Unmarshal(contractBodyCloneCloneBytes, &temp)
	if err != nil {
		beego.Error("model.contract.go IsSignatureValid error ", err)
	}
	temp.ContractSignatures = nil
	contractBody_serialized := common.Serialize(temp)
	/*-------------module deep copy end --------------*/

	contractOwners := c.ContractBody.ContractOwners
	contractSignatures := c.ContractBody.ContractSignatures

	contractOwners_len := len(contractOwners)
	if contractOwners_len != len(contractSignatures) {
		return false
	}

	for index, contractOwner := range contractOwners {
		if contractOwner == "" {
			return false
		}

		contractSignature := contractSignatures[index]
		if contractOwner != contractSignature.OwnerPubkey {
			return false
		}

		if contractSignature.Signature == "" {
			return false
		}

		// contract signature verify
		verifyFlag := common.Verify(contractOwner, contractBody_serialized, contractSignature.Signature)
		beego.Debug("contract verify[owner:", contractOwner, ",signature:",
			contractSignature.Signature, "contractBody_serialized", contractBody_serialized, "]")
		if !verifyFlag {
			return false
		}
	}

	return true
}

func (c *ContractModel) ToString() string {
	return common.Serialize(c)
}

// return the  id (hash generate)
func (c *ContractModel) GenerateId() string {
	contractBodySerialized := common.Serialize(c.Contract.ContractBody)
	beego.Warn("GenerateId ...", common.HashData(contractBodySerialized))
	return common.HashData(contractBodySerialized)
}

//Validate the contract header
func (c *ContractModel) validateContractHead() bool {

	pub_keys := config.GetAllPublicKey()
	pub_keysSet := common.StrArrayToHashSet(pub_keys)
	contractHead := c.Contract.ContractHead
	if contractHead.MainPubkey == "" {
		beego.Error("contract main_pubkey blank")
		return false
	}

	if !pub_keysSet.Has(contractHead.MainPubkey) {
		beego.Warn("main_pubkey ", contractHead.MainPubkey, " not in pubkeys")
		return false
	}
	return true
}
