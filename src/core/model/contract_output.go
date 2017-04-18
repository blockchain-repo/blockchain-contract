package model

import (
	"unicontract/src/common"
	"unicontract/src/core/protos"
)

type ContractOutputWithoutId struct {

}

// table [contract_output]
type ContractOutput struct {
	Id          string             `json:"id"`          //ContractOutput.Id
	Transaction protos.Transaction `json:"transaction"` //ContractOutput.Transaction
	Version     int16              `json:"version"`     //ContractOutput.Version
}


func (c *ContractOutput) ToString() string {
	return common.Serialize(c)
}

// return the id (hash generate)
func (c *ContractOutput) GetId() string {
	contract_serialized := common.Serialize(c.Transaction)
	return common.HashData(contract_serialized)
}