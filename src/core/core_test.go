package core

import (
	"testing"
	"unicontract/src/core/model"
)

func TestWriteContract(t *testing.T) {
	contractModel := model.ContractModel{}
	contractModel.Id ="sssssssssssssssssssssssssssssssssssssssssssssss"
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
	_ = signatureContract
	WriteContract(contractModel)
}