package core

import (
	"testing"
	"unicontract/src/common"
	"unicontract/src/core/model"
	"unicontract/src/core/protos"
)

func TestWriteContract(t *testing.T) {
	contractModel := model.ContractModel{}
//	private_key := "5Pv7F7g9BvNDEMdb8HV5aLHpNTNkxVpNqnLTQ58Z5heC"
	// modify and set value for reference obj with &
	contractBody := &protos.ContractBody{}
	contractBody.ContractOwners = []string{
		"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
		"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
	}
	contractModel.ContractBody = contractBody
	contractModel.Id = common.GenTimestamp()
	WriteContract(contractModel)
}