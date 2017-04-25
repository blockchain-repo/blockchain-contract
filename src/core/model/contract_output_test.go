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

	conditions := make([]*ConditionsItem, 1)
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

	Votes := []*Vote{
		{
			Id:         common.GenerateUUID(),
			NodePubkey: "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
			VoteBody: VoteBody{
				IsValid:         true,
				InvalidReason:   "",
				VoteForContract: "",
				VoteType:        "",
				Timestamp:       common.GenTimestamp(),
			},
			Signature: "65D27HW4uXYvkekGssAQB93D92onMyU1NVnCJnE1PgRKz2uFSPZ6aQvid4qZvkxys7G4r2Mf2KFn5BSQyEBhWs34",
		},
		{
			Id:         common.GenerateUUID(),
			NodePubkey: "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			VoteBody: VoteBody{
				IsValid:         true,
				InvalidReason:   "",
				VoteForContract: "",
				VoteType:        "",
				Timestamp:       common.GenTimestamp(),
			},
			Signature: "5i5dTtQseQjWZ8UdchqQtgttyeeFmB3LDFYzNKafvV2YvTqwv4wZ9mFsH7qgysV9ow893D1h2Xnt1uCXLHtbKrkT",
		},
	}

	relaction := &Relaction{
		ContractId: "3ea445410f608e6453cdcb7dbe42d57a89aca018993d7e87da85993cbccc6308",
		TaskId:     "123",
		Voters:     []string{"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3", "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo"},
		Votes:      Votes,
	}
	transaction.Relaction = relaction

	transaction.Timestamp = common.GenTimestamp()
	contractOutput.Transaction = transaction
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
