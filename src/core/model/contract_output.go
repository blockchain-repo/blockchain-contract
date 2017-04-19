package model

import (
	"encoding/json"
	"fmt"
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
func (c ContractOutput) GetId() string {
	contract_serialized := common.Serialize(c.Transaction)
	_ = contract_serialized

	/*-------------deep copy--------------*/
	var transactionClone interface{} = c.Transaction
	fmt.Println(transactionClone)

	var temp protos.Transaction
	transactionCloneBytes, _ := json.Marshal(transactionClone)
	err := json.Unmarshal(transactionCloneBytes, &temp)
	if err != nil {
		fmt.Println(err)
		beego.Error("contract output getId error ", err)
	}

	temp.Relaction.Signatures = nil
	contract_without_signatures_serialized := common.Serialize(temp)
	fmt.Println(temp)
	return common.HashData(contract_without_signatures_serialized)

	//var transaction protos.Transaction
	//fmt.Println(&(c.Transaction))
	//transaction = c.Transaction
	//transaction.Relaction.Signatures = nil
	//transaction.Relaction.TaskId = "2312312"
	//fmt.Println(transaction)
	//contract_without_signatures_serialized := common.Serialize(transaction)
	//fmt.Println(c.Transaction)

}
