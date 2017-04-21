package model

import (
	"fmt"
	"testing"
	"unicontract/src/common"
)

func Test_ContractOutput(t *testing.T) {
	contractOutput := ContractOutput{}
	contractOutput.Id = "1"
	//transaction := contractOutput.Transaction

	conditions := make([]*ConditionsItem,1)
	conditions[0] = &ConditionsItem{
		Amount: 14213,
		Cid:    0,
		Condition: &Condition{
			Details: nil,
			Uri:     "dd-dsd-qwq-ddd-aa",
		},
		OwnersAfter: []string{"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc"},
	}
	contractOutput.Transaction.Conditions = conditions
	contractOutput.Transaction.Timestamp = common.GenTimestamp()


	result := common.Serialize(contractOutput)
	//result := common.SerializePretty(contractOutput)
	fmt.Println(result)
}

func Test_GenerateId(t *testing.T) {
	contractOutput := ContractOutput{}
	contractOutput.Id = "1"
	transaction := contractOutput.Transaction
	conditions := make([]*ConditionsItem, 1)
	conditions[0] = &ConditionsItem{
		Amount: 14213,
		Cid:    0,
		Condition: &Condition{
			Details: nil,
			Uri:     "dd-dsd-qwq-ddd-aa-ww",
		},
		OwnersAfter: []string{"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc"},
	}
	contractOutput.Transaction.Conditions = conditions

	//relation_signatures := make([]*protos.RelactionSignature, 2)
	//relation_signatures[0] = &protos.RelactionSignature{
	//	ContractNodePubkey:"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
	//	Signature:"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
	//}
	//relation_signatures[1] = &protos.RelactionSignature{
	//	ContractNodePubkey:"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
	//	Signature:"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
	//}

	relation_signatures := []*RelactionSignature{
		{
			ContractNodePubkey: "JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
			Signature:          "EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
		},
		{
			ContractNodePubkey: "JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
			Signature:          "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		},
	}
	relaction := Relaction{
		Signatures: relation_signatures,
		ContractId: "11",
		TaskId:     "1212",
	}
	transaction.Relaction = &relaction

	transaction.Timestamp = common.GenTimestamp()
	contractOutput.Transaction = transaction
	//fmt.Printf("1111%+v\n", transaction)
	fmt.Println(transaction)
	//result := common.Serialize(contractOutput)
	contract_output_id := contractOutput.GenerateId()
	fmt.Println("contract_output_id= " + contract_output_id)

	//transaction.Relaction.Signatures = nil
	fmt.Println(transaction)
	//fmt.Printf("2222%+v\n", transaction)
	fmt.Println("contract_output_id2= " + contract_output_id)

	expected_id := "6834f1539cb01247dd4b3e1b789b09e4a3a7477d8d55b69555e1311de6130ab0"
	//expected_id := "ea9d67ef76162ff69b5a4911e2acf8d8b088a968031f3af8168eee71ca1fdc01"
	//expected_id := "2b2cd0b17a07407e83155d0f5999eb107b9127eee1b876124322e5b17c9b01c4"
	//expected_id := "ea9d67ef76162ff69b5a4911e2acf8d8b088a968031f3af8168eee71ca1fdc012"
	if contract_output_id != expected_id {
		t.Error()
	}

}
