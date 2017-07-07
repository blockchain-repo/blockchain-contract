package model

import (
	"encoding/json"
	"unicontract/src/common/uniledgerlog"
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
	var contractBodyClone = c.Contract.ContractBody

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
		uniledgerlog.Debug("contract verify[owner:", ownerPubkey, ",signature:", signature, "contractBody", contractBody_serialized, "]\n", verifyFlag)
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
	contractBodySerialized := common.StructSerialize(c.Contract.ContractBody)
	uniledgerlog.Warn("contractBodySerialized:\n", contractBodySerialized)
	return common.HashData(contractBodySerialized)
}

//Validate the contract header
func (c *ContractModel) validateContractHead() bool {

	pub_keys := config.GetAllPublicKey()
	pub_keysSet := common.StrArrayToHashSet(pub_keys)
	contractHead := c.Contract.ContractHead
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

//transfer contract(proto) to contractModel
func FromContractToContractModel(contract protos.Contract) ContractModel {
	var contractModel ContractModel
	contractModel.Contract = contract
	return contractModel
}

// transfer contractModel string to contract(proto)
func FromContractModelStrToContract(contractModelStr string) (protos.Contract, error) {
	// 1. to contractModel
	var contractModel ContractModel
	err := json.Unmarshal([]byte(contractModelStr), &contractModel)
	// 2. to contract
	contract := contractModel.Contract
	if err != nil {
		uniledgerlog.Error("error fromContractModelStrToContract", err)
		return contract, err
	}

	return contract, nil
}

// transfer contractModel to contract(proto)
func FromContractModelToContract(contractModel ContractModel) protos.Contract {
	/*-------------module deep copy start --------------*/
	var contractModelClone = contractModel

	// new obj
	var temp protos.Contract

	contractModelCloneBytes, _ := json.Marshal(contractModelClone)
	err := json.Unmarshal(contractModelCloneBytes, &temp)
	if err != nil {
		uniledgerlog.Error("[module-model]FromContractModelToContract error ", err)
	}
	contract := contractModel.Contract
	return contract
}
