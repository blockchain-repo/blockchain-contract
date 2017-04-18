package model

import (
	//"unicontract/src/common"
	"unicontract/src/core/protos"
)

// table [contract_output]
type ContractOutput struct {
	Id          string             `json:"id"`          //ContractOutput.Id
	Transaction protos.Transaction `json:"transaction"` //ContractOutput.Transaction
	Version     int16              `json:"version"`     //ContractOutput.Version
}
