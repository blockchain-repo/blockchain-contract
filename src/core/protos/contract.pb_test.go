package protos

import (
	"fmt"
	"testing"
	"unicontract/src/common"
)

func Test_ContractPb(t *testing.T) {
	contractProto := ContractProto{}

	contractProto.Id = "2"
	contractProto.NodePubkey = "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc"
	contractProto.MainPubkey = "93TEovPuYo6BQFm4ia9ta4qtL1TbAmnk9fV5kxmesAG5"
	contractProto.Signature = "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc"
	contractProto.Voters = []string{"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet"}

	contractProto.Timestamp = common.GenTimestamp()
	contractProto.Version = 1

	contract := Contract{}
	contract.CreatorPubkey = "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc"
	contract.CreateTimestamp = common.GenTimestamp()
	contract.Operation = "CREATE"

	contract_attributes := ContractAttributes{}
	contract_attributes.Name = "XXXXXX"
	contract_attributes.StartTimestamp = common.GenTimestamp()
	contract_attributes.EndTimestamp = common.GenTimestamp()

	contract.ContractAttributes = &contract_attributes

	contract.ContractOwners = []string{"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet"}

	contractSignatures := make([]*ContractSignature, 3)
	contractSignatures[0] = &ContractSignature{OwnerPubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		Signature: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		Timestamp: common.GenTimestamp()}
	contractSignatures[1] = &ContractSignature{OwnerPubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		Signature: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		Timestamp: common.GenTimestamp()}
	contractSignatures[2] = &ContractSignature{OwnerPubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		Signature: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		Timestamp: common.GenTimestamp()}

	contract.ContractSignatures = contractSignatures

	contractAsserts := make([]*ContractAssert, 3)

	//json_data := `{"name":"start", "age":22}`
	//json_data = common.Serialize(json_data)

	contractAsserts[0] = &ContractAssert{Id: "111", Name:"wx1",
		Amount: 1000, Metadata: nil}
	contractAsserts[1] = &ContractAssert{Id: "113", Name: "wx2",
		Amount: 800}
	contractAsserts[2] = &ContractAssert{Id: "112", Name:"wx3",
		Amount: 100}

	contract.ContractAsserts = contractAsserts

	contractComponents := ContractComponents{}

	plans := make([]*Plan, 2)
	plans[0] = &Plan{Id:"ID_Axxxxxx", Type: "PLAN",
		State: "dormant", Name: "N_Axxxxx",
		Description: "xxxxx"}

	planConditions := make([]*PlanTaskCondition, 3)
	planConditions[0] = &PlanTaskCondition{Id: "1", Type: "PreCondition",
		Name: "XXXX", Value: "XXXX", Description: "xxxxx"}
	planConditions[1] = &PlanTaskCondition{Id: "2", Type:"DisgardCondition",
		Name: "XXXX", Value: "XXXX", Description: "xxxxx"}
	planConditions[2] = &PlanTaskCondition{Id: "3", Type: "CompleteCondition",
		Name: "XXXX", Value: "XXXX", Description: "xxxxx"}

	plans[0].Condition = planConditions
	plans[0].Level = 1
	plans[0].ContractType = "RIGHT"
	plans[0].NextTask = []string{"1", "2"}

	plans[1] = &Plan{Id: "ID_Axxxxxx", Type: "PLAN",
		State: "dormant", Name: "N_Axxxxx",
		Description: "xxxxx"}
	planConditions2 := make([]*PlanTaskCondition, 3)
	planConditions2[0] = &PlanTaskCondition{Id: "1", Type: "PreCondition",
		Name: "XXXX", Value: "XXXX", Description: "xxxxx"}
	planConditions2[1] = &PlanTaskCondition{Id: "2", Type: "DisgardCondition",
		Name: "XXXX", Value: "XXXX", Description: "xxxxx"}
	planConditions2[2] = &PlanTaskCondition{Id: "3", Type: "CompleteCondition",
		Name: "XXXX", Value: "XXXX", Description: "xxxxx"}

	plans[1].Condition = planConditions2
	plans[1].Level = 1
	plans[1].ContractType = "RIGHT"
	plans[1].NextTask = []string{"1", "2"}

	contractComponents.Plans = plans

	tasks := make([]*Task, 3)

	tasks[0] = &Task{Id: "ID_Cxxxxxx", Type: "ENQUIRY",
		State: "dormant", Name: "Axxxxxx",
		Description: "xxxxx"}
	taskConditions := make([]*PlanTaskCondition, 3)
	taskConditions[0] = &PlanTaskCondition{Id: "1", Type: "PreCondition",
		Name: "XXXX", Value: "XXXX", Description: "xxxxx"}
	taskConditions[1] = &PlanTaskCondition{Id: "2", Type: "DisgardCondition",
		Name: "XXXX", Value: "XXXX", Description: "xxxxx"}
	taskConditions[2] = &PlanTaskCondition{Id: "3", Type: "CompleteCondition",
		Name: "XXXX", Value: "XXXX", Description: "xxxxx"}

	tasks[0].Condition = taskConditions
	tasks[0].Level = 1
	tasks[0].ContractType = "RIGHT"
	tasks[0].NextTask = []string{"Axxxxxx", "Bxxxxxx"}

	tasks[1] = &Task{Id: "ID_Cxxxxxx", Type: "ENQUIRY",
		State: "dormant", Name: "Bxxxxxx",
		Description: "xxxxx"}
	taskConditions1 := make([]*PlanTaskCondition, 3)
	taskConditions1[0] = &PlanTaskCondition{Id: "1", Type: "PreCondition",
		Name: "XXXX", Value: "XXXX", Description: "xxxxx"}
	taskConditions1[1] = &PlanTaskCondition{Id: "2", Type: "DisgardCondition",
		Name: "XXXX", Value: "XXXX", Description: "xxxxx"}
	taskConditions1[2] = &PlanTaskCondition{Id: "3", Type: "CompleteCondition",
		Name: "XXXX", Value: "XXXX", Description: "xxxxx"}

	tasks[1].Condition = taskConditions1
	tasks[1].Level = (1)
	tasks[1].ContractType = "RIGHT"
	tasks[1].NextTask = []string{"", ""}

	tasks[2] = &Task{Id: "ID_Cxxxxxx", Type: "ACTION",
		State: "dormant", Name: "Cxxxxxx",
		Description: "xxxxx"}
	taskConditions2 := make([]*PlanTaskCondition, 3)
	taskConditions2[0] = &PlanTaskCondition{Id: "1", Type: "PreCondition",
		Name: "XXXX", Value: "XXXX", Description: "xxxxx"}
	taskConditions2[1] = &PlanTaskCondition{Id: "2", Type: "DisgardCondition",
		Name: "XXXX", Value: "XXXX", Description: "xxxxx"}
	taskConditions2[2] = &PlanTaskCondition{Id: "3", Type: "CompleteCondition",
		Name: "XXXX", Value: "XXXX", Description: "xxxxx"}

	tasks[2].Condition = taskConditions2
	tasks[2].Level = (1)
	tasks[2].ContractType = "DUTY"
	tasks[2].NextTask = []string{"", ""}
	contractComponents.Tasks = tasks

	contract.ContractComponents = &contractComponents
	contractProto.Contract = &contract

	//fmt.Println(contractProto)
	fmt.Println(common.SerializePretty(contractProto))
}
