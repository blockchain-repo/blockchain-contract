package model

import (
	"encoding/json"
	"github.com/astaxie/beego"
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

	/*-------------module deep copy--------------*/
	var transactionClone = c.Transaction

	// new obj
	var temp protos.Transaction

	transactionCloneBytes, _ := json.Marshal(transactionClone)
	err := json.Unmarshal(transactionCloneBytes, &temp)
	if err != nil {
		beego.Error("Unmarshal error ", err)
	}
	beego.Debug(temp)
	temp.Relaction.Signatures = nil
	beego.Debug(temp)
	contract_without_signatures_serialized := common.Serialize(temp)

	return common.HashData(contract_without_signatures_serialized)
}
