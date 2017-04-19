package model

import (
	"fmt"
	"testing"
	"unicontract/src/common"
)

func Test_Sign(t *testing.T) {
	//create new obj
	contractModel := ContractModel{}
	private_key := "5Pv7F7g9BvNDEMdb8HV5aLHpNTNkxVpNqnLTQ58Z5heC"
	// modify and set value for reference obj with &
	contract := &contractModel.Contract
	contract.CreatorPubkey = "5Pv7F7g9BvNDEMdb8HV5aLHpNTNkxVpNqnLTQ58Z5heC"
	contract.Operation = "CREATE"
	contract.ContractOwners = []string{
		"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		//"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
	}
	// sign for contract
	signatureContract := contractModel.Sign(private_key)
	//fmt.Println("contract is : ", common.SerializePretty(contract))
	fmt.Println("private_key is : ", private_key)
	fmt.Println("contract is : ", common.Serialize(contract))
	fmt.Println("signatureContract isTest_Validate : ", signatureContract)
}

func Test_IsSignatureValid(t *testing.T) {
	//create new obj
	contractModel := ContractModel{}
	private_key := "5Pv7F7g9BvNDEMdb8HV5aLHpNTNkxVpNqnLTQ58Z5heC"
	// modify and set value for reference obj with &
	contract := &contractModel.Contract
	contract.CreatorPubkey = "5Pv7F7g9BvNDEMdb8HV5aLHpNTNkxVpNqnLTQ58Z5heC"
	contract.Operation = "CREATE"
	contract.ContractOwners = []string{
		"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		//"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
	}
	// sign for contract
	signatureContract := contractModel.Sign(private_key)
	contractModel.Signature = signatureContract

	fmt.Println("private_key is : ", private_key)
	fmt.Println("contract is : ", common.Serialize(contract))
	fmt.Println("signatureContract is : ", signatureContract)
	contractModel.NodePubkey = "3FyHdZVX4adfSSTg7rZDPMzqzM8k5fkpu43vbRLvEXLJ"
	//contractModel.NodePubkey = "AZfjdKxEr9G3NwdAkco22nN8PfgQvCr5TDPK1tqsGZrk"
	fmt.Println("node_pubkey is : ", contractModel.NodePubkey)
	isSignatureValid := contractModel.IsSignatureValid()
	if isSignatureValid {
		t.Log("contract 签名有效")
	} else {
		t.Error("contract 签名无效")
	}
}

func Test_Validate(t *testing.T) {
	//create new obj
	contractModel := ContractModel{}
	private_key := "5Pv7F7g9BvNDEMdb8HV5aLHpNTNkxVpNqnLTQ58Z5heC"
	// modify and set value for reference obj with &
	contract := &contractModel.Contract
	contract.CreatorPubkey = "5Pv7F7g9BvNDEMdb8HV5aLHpNTNkxVpNqnLTQ58Z5heC"
	contract.Operation = "CREATE"
	contract.ContractOwners = []string{
		"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		//"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
	}
	// sign for contract
	signatureContract := contractModel.Sign(private_key)
	contractModel.Signature = signatureContract

	fmt.Println("private_key is : ", private_key)
	fmt.Println("contract is : ", common.Serialize(contract))
	fmt.Println("signatureContract is : ", signatureContract)
	contractModel.NodePubkey = "3FyHdZVX4adfSSTg7rZDPMzqzM8k5fkpu43vbRLvEXLJ"
	contractModel.Voters = []string{
		"3FyHdZVX4adfSSTg7rZDPMzqzM8k5fkpu43vbRLvEXLJ",
		"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
	}
	//contractModel.NodePubkey = "AZfjdKxEr9G3NwdAkco22nN8PfgQvCr5TDPK1tqsGZrk"
	fmt.Println("node_pubkey is : ", contractModel.NodePubkey)
	contractModel.Validate()
}
