package model

import (
	"fmt"
	"testing"
	"unicontract/src/common"
	"unicontract/src/core/protos"
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
	contractModel.Id = contractModel.GenerateId()
	//fmt.Println("contract is : ", common.SerializePretty(contract))
	fmt.Println("private_key is : ", private_key)
	fmt.Println("contract is : ", common.Serialize(contract))
	fmt.Println("signatureContract isTest_Validate : ", signatureContract)
}

func Test_IsSignatureValid(t *testing.T) {
	//create new obj
	contractModel := ContractModel{}
	private_key := "Cnodz1gyhaNoFcPCr72G9brFrGFfNJQUPFchGXyL11Pt"
	// modify and set value for reference obj with &
	contract := &contractModel.Contract
	contractModel.MainPubkey = "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3"
	contract.CreatorPubkey = "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3"
	contract.Operation = "CREATE"
	contract.ContractOwners = []string{
		"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		//"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		//"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
	}
	contractSignatures := []*protos.ContractSignature{
		{
			OwnerPubkey: "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
			Signature:   "3XLffBVuFCZbZU1NcroQDSAgcdDtYQ2UK9ye9q9BzLaMiqjoHtJ3SirW5P9JJkjwAkC9CrguwKMRC36T2e769sqQ",
			Timestamp:   common.GenTimestamp(),
		},
		//{
		//	OwnerPubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		//	Signature:   "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		//	Timestamp:   common.GenTimestamp(),
		//},
	}
	contract.ContractSignatures = contractSignatures
	// sign for contract
	signatureContract := contractModel.Sign(private_key)

	fmt.Println("private_key is : ", private_key)
	fmt.Println("contract is : ", common.Serialize(contract))
	fmt.Println("signatureContract is : ", signatureContract)
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
	contractModel.MainPubkey = "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3"
	private_key := "Cnodz1gyhaNoFcPCr72G9brFrGFfNJQUPFchGXyL11Pt"
	//----------------------pub: HFStRMeL1aS824zsApWiEmZqn1r21FvkXQFkqGwJjb3d
	//----------------------pri: EUFo91bZzqtZEAMUZ4Mr5N69BBwfBkKexwXDydTncqqg
	// modify and set value for reference obj with &
	contract := &contractModel.Contract
	contract.CreatorPubkey = "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3"
	contract.Operation = "CREATE"
	contract.ContractOwners = []string{
		"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		//"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		//"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		//"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		//"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
	}
	// sign for contract
	signatureContract := contractModel.Sign(private_key)

	fmt.Println("private_key is : ", private_key)
	fmt.Println("contract is : ", common.Serialize(contract))
	fmt.Println("signatureContract is : ", signatureContract)
	contractSignatures := []*protos.ContractSignature{
		{
			OwnerPubkey: "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
			Signature:   "3XLffBVuFCZbZU1NcroQDSAgcdDtYQ2UK9ye9q9BzLaMiqjoHtJ3SirW5P9JJkjwAkC9CrguwKMRC36T2e769sqQ",
			Timestamp:   common.GenTimestamp(),
		},
		//{
		//	OwnerPubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		//	Signature:   "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		//	Timestamp:   common.GenTimestamp(),
		//},
	}
	contract.ContractSignatures = contractSignatures
	contractModel.Voters = []string{
		"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		//"3FyHdZVX4adfSSTg7rZDPMzqzM8k5fkpu43vbRLvEXLJ",
		//"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
	}
	contractModel.Id = contractModel.GenerateId()
	ok := contractModel.Validate()
	fmt.Println(ok)
}
