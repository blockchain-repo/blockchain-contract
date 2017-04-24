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
	//private_key := "Cnodz1gyhaNoFcPCr72G9brFrGFfNJQUPFchGXyL11Pt"
	private_key := "GEXZrZShFsHKZB94mpRmaDBBtjWCDz6TgK17R9DXUwex"

	// modify and set value for reference obj with &
	contract := &contractModel.Contract
	contractHead := &protos.ContractHead{}
	contractBody := &protos.ContractBody{}
	// modify and set value for reference obj with &
	//contractHead.MainPubkey = "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3"
	contractHead.MainPubkey = "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo"
	contractBody.Cname = "star"
	contractBody.ContractOwners = []string{
		"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
		//"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
	}
	contractModel.ContractHead = contractHead
	contractModel.ContractBody = contractBody
	fmt.Println("contract is : ", common.Serialize(contract))
	// sign for contract
	signatureContract := contractModel.Sign(private_key)
	contractModel.Id = contractModel.GenerateId()
	//fmt.Println("contract is : ", common.SerializePretty(contract))
	fmt.Println("private_key is : ", private_key)
	fmt.Println("contract is : ", common.Serialize(contract))
	fmt.Println("signatureContract isTest_Validate : ", signatureContract)
	// 65D27HW4uXYvkekGssAQB93D92onMyU1NVnCJnE1PgRKz2uFSPZ6aQvid4qZvkxys7G4r2Mf2KFn5BSQyEBhWs34
	// 5i5dTtQseQjWZ8UdchqQtgttyeeFmB3LDFYzNKafvV2YvTqwv4wZ9mFsH7qgysV9ow893D1h2Xnt1uCXLHtbKrkT
}

func Test_IsSignatureValid(t *testing.T) {
	//create new obj
	contractModel := &ContractModel{}
	contractHead := &protos.ContractHead{}
	contractBody := &protos.ContractBody{}

	private_key := "Cnodz1gyhaNoFcPCr72G9brFrGFfNJQUPFchGXyL11Pt"
	//private_key := "GEXZrZShFsHKZB94mpRmaDBBtjWCDz6TgK17R9DXUwex"
	// modify and set value for reference obj with &
	contractHead.MainPubkey = "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3"
	//contractHead.MainPubkey = "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo"
	contractBody.Cname = "star"
	contractBody.ContractOwners = []string{
		"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
	}
	contractSignatures := []*protos.ContractSignature{
		{
			OwnerPubkey: "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
			Signature:   "65D27HW4uXYvkekGssAQB93D92onMyU1NVnCJnE1PgRKz2uFSPZ6aQvid4qZvkxys7G4r2Mf2KFn5BSQyEBhWs34",
			SignTimestamp:   common.GenTimestamp(),
		},
		{
			OwnerPubkey: "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			Signature:   "5i5dTtQseQjWZ8UdchqQtgttyeeFmB3LDFYzNKafvV2YvTqwv4wZ9mFsH7qgysV9ow893D1h2Xnt1uCXLHtbKrkT",
			SignTimestamp:   common.GenTimestamp(),
		},
	}
	contractBody.ContractSignatures = contractSignatures

	contractModel.ContractHead = contractHead
	contractModel.ContractBody = contractBody
	// sign for contract
	signatureContract := contractModel.Sign(private_key)

	_ = private_key
	_ = signatureContract

	//fmt.Println("private_key is : ", private_key)
	//fmt.Println("contract is : ", common.Serialize(contract))
	//fmt.Println("signatureContract is : ", signatureContract)
	fmt.Println("contractModel is : ", contractModel)
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
	contractHead := &protos.ContractHead{}
	contractBody := &protos.ContractBody{}

	private_key := "GEXZrZShFsHKZB94mpRmaDBBtjWCDz6TgK17R9DXUwex"
	//private_key := "Cnodz1gyhaNoFcPCr72G9brFrGFfNJQUPFchGXyL11Pt"
	contractHead.MainPubkey = "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo"
	//contractHead.MainPubkey = "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3"
	//----------------------pub: HFStRMeL1aS824zsApWiEmZqn1r21FvkXQFkqGwJjb3d
	//----------------------pri: EUFo91bZzqtZEAMUZ4Mr5N69BBwfBkKexwXDydTncqqg
	// modify and set value for reference obj with &
	contractBody.Cname = "star"
	contractBody.ContractOwners = []string{
		"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
		//"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
	}
	contractSignatures := []*protos.ContractSignature{
		{
			OwnerPubkey: "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			Signature:   "4JmRZ2A1Dqf4sGQVS7Jo6nNdR17XxdYddSC3fE4bv6ov48J9CCSMSKmx9AUtkaqJLpsLEGepzjpTbZrXCpbohVeU",
			SignTimestamp:   common.GenTimestamp(),
		},
		//{
		//	OwnerPubkey: "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		//	Signature:   "65D27HW4uXYvkekGssAQB93D92onMyU1NVnCJnE1PgRKz2uFSPZ6aQvid4qZvkxys7G4r2Mf2KFn5BSQyEBhWs34",
		//	//Signature:   "3XLffBVuFCZbZU1NcroQDSAgcdDtYQ2UK9ye9q9BzLaMiqjoHtJ3SirW5P9JJkjwAkC9CrguwKMRC36T2e769sqQ",
		//	SignTimestamp:   common.GenTimestamp(),
		//},
	}
	contractModel.ContractHead = contractHead
	contractModel.ContractBody = contractBody
	contractBody.ContractSignatures = contractSignatures
	// sign for contract
	signatureContract := contractModel.Sign(private_key)

	fmt.Println("private_key is : ", private_key)
	fmt.Println("contract is : ", common.Serialize(contractBody))
	fmt.Println("signatureContract is : ", signatureContract)

	//contractModel.ContractBody = contractBody
	contractModel.Id = contractModel.GenerateId()
	ok := contractModel.Validate()
	fmt.Println(ok)
}
