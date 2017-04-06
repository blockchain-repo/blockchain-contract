package model

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"testing"
	"time"
	"unicontract/src/common"
)

func Test_ContractPb(t *testing.T) {
	contractProto := ContractProto{}

	contractProto.Id = proto.String("2")
	contractProto.NodePubkey = proto.String("2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc")
	contractProto.MainPubkey = proto.String("93TEovPuYo6BQFm4ia9ta4qtL1TbAmnk9fV5kxmesAG5")
	contractProto.Signature = proto.String("2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc")
	contractProto.Voters = []string{"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet"}

	contractProto.Timestamp = proto.Int64(time.Now().Unix())
	contractProto.Version = proto.String("v1.0")

	contract := Contract{}
	contract.CreatorPubkey = proto.String("2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc")
	contract.CreateTimestamp = proto.Int64(time.Now().Unix())
	contract.Operation = proto.String("CREATE")

	contract_attributes := ContractAttributes{}
	contract_attributes.Name = proto.String("XXXXXX")
	contract_attributes.StartTimestamp = proto.Int64(time.Now().Unix())
	contract_attributes.EndTimestamp = proto.Int64(time.Now().Unix())

	contract.ContractAttributes = &contract_attributes

	contract.ContractOwners = []string{"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet"}

	contractSignatures := make([]*ContractSignature, 3)
	contractSignatures[0] = &ContractSignature{OwnerPubkey: proto.String("2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc"),
		Signature: proto.String("2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc"),
		Timestamp: proto.Int64(time.Now().Unix())}
	contractSignatures[1] = &ContractSignature{OwnerPubkey: proto.String("2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc"),
		Signature: proto.String("2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc"),
		Timestamp: proto.Int64(time.Now().Unix())}
	contractSignatures[2] = &ContractSignature{OwnerPubkey: proto.String("2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc"),
		Signature: proto.String("2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc"),
		Timestamp: proto.Int64(time.Now().Unix())}

	contract.ContractSignatures = contractSignatures

	contractAsserts := make([]*ContractAssert, 3)

	//json_data := `{"name":"start", "age":22}`
	//json_data = common.Serialize(json_data)

	contractAsserts[0] = &ContractAssert{Id: proto.String("111"), Name: proto.String("wx1"),
		Amount: proto.Float32(1000), Metadata: nil}
	contractAsserts[1] = &ContractAssert{Id: proto.String("113"), Name: proto.String("wx2"),
		Amount: proto.Float32(800)}
	contractAsserts[2] = &ContractAssert{Id: proto.String("112"), Name: proto.String("wx3"),
		Amount: proto.Float32(100)}

	contract.ContractAsserts = contractAsserts

	contractComponents := ContractComponents{}

	plans := make([]*Plan, 2)
	plans[0] = &Plan{Id: proto.String("ID_Axxxxxx"), Type: proto.String("PLAN"),
		State: proto.String("dormant"), Name: proto.String("N_Axxxxx"),
		Description: proto.String("xxxxx")}

	planConditions := make([]*PlanTaskCondition, 3)
	planConditions[0] = &PlanTaskCondition{Id: proto.String("1"), Type: proto.String("PreCondition"),
		Name: proto.String("XXXX"), Value: proto.String("XXXX"), Description: proto.String("xxxxx")}
	planConditions[1] = &PlanTaskCondition{Id: proto.String("2"), Type: proto.String("DisgardCondition"),
		Name: proto.String("XXXX"), Value: proto.String("XXXX"), Description: proto.String("xxxxx")}
	planConditions[2] = &PlanTaskCondition{Id: proto.String("3"), Type: proto.String("CompleteCondition"),
		Name: proto.String("XXXX"), Value: proto.String("XXXX"), Description: proto.String("xxxxx")}

	plans[0].Condition = planConditions
	plans[0].Level = proto.Int32(1)
	plans[0].ContractType = proto.String("RIGHT")
	plans[0].NextTask = []string{"1", "2"}

	plans[1] = &Plan{Id: proto.String("ID_Axxxxxx"), Type: proto.String("PLAN"),
		State: proto.String("dormant"), Name: proto.String("N_Axxxxx"),
		Description: proto.String("xxxxx")}
	planConditions2 := make([]*PlanTaskCondition, 3)
	planConditions2[0] = &PlanTaskCondition{Id: proto.String("1"), Type: proto.String("PreCondition"),
		Name: proto.String("XXXX"), Value: proto.String("XXXX"), Description: proto.String("xxxxx")}
	planConditions2[1] = &PlanTaskCondition{Id: proto.String("2"), Type: proto.String("DisgardCondition"),
		Name: proto.String("XXXX"), Value: proto.String("XXXX"), Description: proto.String("xxxxx")}
	planConditions2[2] = &PlanTaskCondition{Id: proto.String("3"), Type: proto.String("CompleteCondition"),
		Name: proto.String("XXXX"), Value: proto.String("XXXX"), Description: proto.String("xxxxx")}

	plans[1].Condition = planConditions2
	plans[1].Level = proto.Int32(1)
	plans[1].ContractType = proto.String("RIGHT")
	plans[1].NextTask = []string{"1", "2"}

	contractComponents.Plans = plans

	tasks := make([]*Task, 3)

	tasks[0] = &Task{Id: proto.String("ID_Cxxxxxx"), Type: proto.String("ENQUIRY"),
		State: proto.String("dormant"), Name: proto.String("Axxxxxx"),
		Description: proto.String("xxxxx")}
	taskConditions := make([]*PlanTaskCondition, 3)
	taskConditions[0] = &PlanTaskCondition{Id: proto.String("1"), Type: proto.String("PreCondition"),
		Name: proto.String("XXXX"), Value: proto.String("XXXX"), Description: proto.String("xxxxx")}
	taskConditions[1] = &PlanTaskCondition{Id: proto.String("2"), Type: proto.String("DisgardCondition"),
		Name: proto.String("XXXX"), Value: proto.String("XXXX"), Description: proto.String("xxxxx")}
	taskConditions[2] = &PlanTaskCondition{Id: proto.String("3"), Type: proto.String("CompleteCondition"),
		Name: proto.String("XXXX"), Value: proto.String("XXXX"), Description: proto.String("xxxxx")}

	tasks[0].Condition = taskConditions
	tasks[0].Level = proto.Int32(1)
	tasks[0].ContractType = proto.String("RIGHT")
	tasks[0].NextTask = []string{"Axxxxxx", "Bxxxxxx"}

	tasks[1] = &Task{Id: proto.String("ID_Cxxxxxx"), Type: proto.String("ENQUIRY"),
		State: proto.String("dormant"), Name: proto.String("Bxxxxxx"),
		Description: proto.String("xxxxx")}
	taskConditions1 := make([]*PlanTaskCondition, 3)
	taskConditions1[0] = &PlanTaskCondition{Id: proto.String("1"), Type: proto.String("PreCondition"),
		Name: proto.String("XXXX"), Value: proto.String("XXXX"), Description: proto.String("xxxxx")}
	taskConditions1[1] = &PlanTaskCondition{Id: proto.String("2"), Type: proto.String("DisgardCondition"),
		Name: proto.String("XXXX"), Value: proto.String("XXXX"), Description: proto.String("xxxxx")}
	taskConditions1[2] = &PlanTaskCondition{Id: proto.String("3"), Type: proto.String("CompleteCondition"),
		Name: proto.String("XXXX"), Value: proto.String("XXXX"), Description: proto.String("xxxxx")}

	tasks[1].Condition = taskConditions1
	tasks[1].Level = proto.Int32(1)
	tasks[1].ContractType = proto.String("RIGHT")
	tasks[1].NextTask = []string{"", ""}

	tasks[2] = &Task{Id: proto.String("ID_Cxxxxxx"), Type: proto.String("ACTION"),
		State: proto.String("dormant"), Name: proto.String("Cxxxxxx"),
		Description: proto.String("xxxxx")}
	taskConditions2 := make([]*PlanTaskCondition, 3)
	taskConditions2[0] = &PlanTaskCondition{Id: proto.String("1"), Type: proto.String("PreCondition"),
		Name: proto.String("XXXX"), Value: proto.String("XXXX"), Description: proto.String("xxxxx")}
	taskConditions2[1] = &PlanTaskCondition{Id: proto.String("2"), Type: proto.String("DisgardCondition"),
		Name: proto.String("XXXX"), Value: proto.String("XXXX"), Description: proto.String("xxxxx")}
	taskConditions2[2] = &PlanTaskCondition{Id: proto.String("3"), Type: proto.String("CompleteCondition"),
		Name: proto.String("XXXX"), Value: proto.String("XXXX"), Description: proto.String("xxxxx")}

	tasks[2].Condition = taskConditions2
	tasks[2].Level = proto.Int32(1)
	tasks[2].ContractType = proto.String("DUTY")
	tasks[2].NextTask = []string{"", ""}
	contractComponents.Tasks = tasks

	contract.ContractComponents = &contractComponents
	contractProto.Contract = &contract

	//fmt.Println(contractProto)
	fmt.Println(common.SerializePretty(contractProto))
}
