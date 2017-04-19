package db

import (
	"fmt"
	"testing"
	"unicontract/src/common"
	"unicontract/src/core/model"
)

func Test_InsertContractStruct(t *testing.T) {
	//create new obj
	contractModel := model.ContractModel{}
	//TODO

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
	contractModel.MainPubkey = "AZfjdKxEr9G3NwdAkco22nN8PfgQvCr5TDPK1tqsGZrk"
	fmt.Println("node_pubkey is : ", contractModel.NodePubkey)

	contractModel.Id = contractModel.GetId()
	isTrue := InsertContractStruct(contractModel)
	fmt.Println(isTrue)
}

func Test_GetContractById(t *testing.T) {
	contractId := "459695c44aec47091b7920bb34f391365c970780d9394451d66593797e092811"
	contract := GetContractById(contractId)
	fmt.Println(contract)
}

func Test_GetContractMainPubkeyById(t *testing.T) {
	contractId := "459695c44aec47091b7920bb34f391365c970780d9394451d66593797e092811"
	main_pubkey := GetContractMainPubkeyById(contractId)
	fmt.Println(main_pubkey)
}
