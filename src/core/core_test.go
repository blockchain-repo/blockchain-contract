package core

import (
	"fmt"
	"testing"
	"unicontract/src/common"
	"unicontract/src/core/model"
	"unicontract/src/core/protos"
)

func TestWriteContract(t *testing.T) {
	contractModel := model.ContractModel{}
	//

	contractHead := &protos.ContractHead{"", 1}

	contractOwners := []string{
		"BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
		//"4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		"9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
	}
	contractBody := &protos.ContractBody{
		ContractId:         "UUID-1234-5678-90",
		Cname:              "test contract output",
		Ctype:              "CREATE",
		Caption:            "购智能手机返话费合约产品协议",
		Description:        "移动用户A花费500元购买移动运营商B的提供的合约智能手机C后",
		ContractState:      "",
		Creator:            common.GenTimestamp(),
		CreatorTime:        common.GenTimestamp(),
		StartTime:          common.GenTimestamp(),
		EndTime:            common.GenTimestamp(),
		ContractOwners:     contractOwners,
		ContractSignatures: nil,
		ContractAssets:     nil,
		ContractComponents: nil,
	}

	contractModel.ContractHead = contractHead
	contractModel.ContractBody = contractBody

	contractSignatures := []*protos.ContractSignature{
		{
			OwnerPubkey:   "BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
			Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
			SignTimestamp: common.GenTimestamp(),
		},
		//{
		//	OwnerPubkey:   "4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		//	Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
		{
			OwnerPubkey:   "9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
			Signature:     contractModel.Sign("9647UfPdDSwBf5kw7tUrSe7cmYY5RvVX47GrGqSh4XVi"),
			SignTimestamp: common.GenTimestamp(),
		},
	}
	contractBody.ContractSignatures = contractSignatures

	contractModel.Id = contractModel.GenerateId()
	fmt.Print(common.Serialize(contractModel))
	WriteContract(contractModel)
}
