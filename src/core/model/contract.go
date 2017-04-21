package model

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"unicontract/src/common"
	"unicontract/src/core/protos"
	"unicontract/src/config"
)

// table [contract]
type ContractModel struct {
	Id         string `json:"id"`          //合约唯一标识ID，对合约主体信息计算hash
	Version    int32   `json:"version"`     //合约描述结构版本号
	MainPubkey string `json:"main_pubkey"` //合约处理主节点公钥
	Timestamp  string `json:"timestamp"`   //合约运行跟踪时间戳（以合约执行层输出结果时间为准）

	//合约运行节点投票公钥环
	//“voters”:[voter_node1_pubkey, voter_node2_pubkey, voter_node3_pubkey],
	Voters   []string        `json:"voters"`
	Contract protos.Contract `json:"contract"` //合约描述集合, (引用contract描述 for proto3)
}

// validate the contract
func (c *ContractModel) Validate() bool {
	headerValid := c.validateContractHeader()
	if !headerValid {
		return false
	}

	idValid := c.Id == c.GenerateId()
	if !idValid {
		return false
	}

	signatureValid := c.IsSignatureValid()
	if !signatureValid {
		return false
	}

	return true
}

//Create a signature for the contract
func (c *ContractModel) Sign(private_key string) string {
	/*-------------module deep copy start --------------*/
	var contractClone = c.Contract

	// new obj
	var temp protos.Contract

	contractCloneBytes, _ := json.Marshal(contractClone)
	err := json.Unmarshal(contractCloneBytes, &temp)
	if err != nil {
		beego.Error("Unmarshal error ", err)
	}
	temp.ContractSignatures = nil
	contract_serialized := common.Serialize(temp)
	/*-------------module deep copy end --------------*/

	signatureContract := common.Sign(private_key, contract_serialized)
	return signatureContract
}

// Check the validity of a Contract's signature
func (c *ContractModel) IsSignatureValid() bool {

	/*-------------module deep copy start --------------*/
	var contractClone = c.Contract

	// new obj
	var temp protos.Contract

	contractCloneBytes, _ := json.Marshal(contractClone)
	err := json.Unmarshal(contractCloneBytes, &temp)
	if err != nil {
		beego.Error("Unmarshal error ", err)
	}
	temp.ContractSignatures = nil
	contract_serialized := common.Serialize(temp)
	/*-------------module deep copy end --------------*/

	contractOwners := c.Contract.ContractOwners
	contractSignatures := c.Contract.ContractSignatures

	contractOwners_len := len(contractOwners)
	if contractOwners_len != len(contractSignatures) {
		return false
	}

	inValidSignatureCount := 0
	for index, contractOwner := range contractOwners {
		if inValidSignatureCount >= (contractOwners_len+1)/2 {
			return false
		}

		if contractOwner == "" {
			inValidSignatureCount++
			continue
		}

		contractSignature := contractSignatures[index]
		if contractOwner != contractSignature.OwnerPubkey{
			inValidSignatureCount++
			continue
		}

		if contractSignature.Signature == "" {
			inValidSignatureCount++
			continue
		}

		// contract signature verify
		verifyFlag := common.Verify(contractOwner, contract_serialized, contractSignature.Signature)
		fmt.Println(contractOwner)
		fmt.Println(contract_serialized)
		fmt.Println(contractSignature.Signature)
		if !verifyFlag {
			inValidSignatureCount++
			continue
		}
	}

	if inValidSignatureCount >= (contractOwners_len+1)/2 {
		return false
	}
	return true
}

func (c *ContractModel) ToString() string {
	return common.Serialize(c)
}

// return the  id (hash generate)
func (c *ContractModel) GenerateId() string {
	/*-------------module deep copy--------------*/
	var contractClone = c.Contract

	// new obj
	var temp protos.Contract

	transactionCloneBytes, _ := json.Marshal(contractClone)
	err := json.Unmarshal(transactionCloneBytes, &temp)
	if err != nil {
		beego.Error("Unmarshal error ", err)
	}
	//TODO deal with the timestamps
	temp.ContractSignatures = nil
	contract_without_signatures_serialized := common.Serialize(temp)

	return common.HashData(contract_without_signatures_serialized)
}

//Validate the contract header
func (c *ContractModel) validateContractHeader() bool {

	pub_keys := config.GetAllPublicKey()
	pub_keysSet := common.StrArrayToHashSet(pub_keys)

	//todo voters in keyring or not ?
	if c.MainPubkey == "" {
		beego.Error("contract main_pubkey blank")
		return false
	}

	if !pub_keysSet.Has(c.MainPubkey) {
		beego.Warn("main_pubkey ", c.MainPubkey," not in pubkeys")
		return false
	}

	_contract := &c.Contract
	if _contract == nil {
		beego.Error("Empty contract is not allowed")
		return false
	}

	if _contract.Operation != "CREATE" {
		return false
		//beego.Error("missing validate the contract [creator_pubkey]")
		//beego.Error("missing validate the contract [create_timestamp]")
		//beego.Error("missing validate the contract [contract_attributes]")
		//beego.Error("missing validate the contract [contract_owners]")
		//beego.Error("missing validate the contract [contract_signatures]")
		//beego.Error("missing validate the contract [contract_asserts]")
		//beego.Error("missing validate the contract [contract_components]")
	}

	return true
}
