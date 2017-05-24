package protos

import (
	"fmt"
	"testing"
	"unicontract/src/common"
	"unicontract/src/config"
)

func Test_ContractPb(t *testing.T) {
	contract := Contract{
	//ContractHead:&ContractHead{
	//	MainPubkey:"123",
	//},
	}

	pubkey := config.Config.Keypair.PublicKey
	contract.Id = "e8d037d71b5dcdadcc90f8b59212b8705fb47369d4c6879f175594b63826fb53"
	contractHead := &ContractHead{}
	contractHead.MainPubkey = "93TEovPuYo6BQFm4ia9ta4qtL1TbAmnk9fV5kxmesAG5"
	contractHead.Version = 1

	contractBody := &ContractBody{}
	contractBody.ContractId = common.GenerateUUID()
	contractBody.Cname = "XXOOXX"
	contractBody.Ctype = "XXOOXX"
	contractBody.Caption = "XXOOXX"
	contractBody.Description = "XXOOXX"
	contractBody.Creator = "XINGSTAR"
	contractBody.CreateTime = common.GenTimestamp()
	contractBody.StartTime = common.GenTimestamp()
	contractBody.EndTime = common.GenTimestamp()
	contractBody.ContractOwners = []string{"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet"}
	contractAssets := make([]*ContractAsset, 3)

	contractAssets[0] = &ContractAsset{AssetId: "111", Name: "wx1", Description: "XXX",
		Amount: 1000, MetaData: nil}
	contractAssets[1] = &ContractAsset{AssetId: "113", Name: "wx2",
		Amount: 800}
	contractAssets[2] = &ContractAsset{AssetId: "112", Name: "wx3", Description: "XXX3",
		Amount: 100}

	contractBody.ContractAssets = contractAssets

	contractSignatures := make([]*ContractSignature, 3)
	contractSignatures[0] = &ContractSignature{OwnerPubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		Signature:     "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		SignTimestamp: common.GenTimestamp()}
	contractSignatures[1] = &ContractSignature{OwnerPubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		Signature:     "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		SignTimestamp: common.GenTimestamp()}
	contractSignatures[2] = &ContractSignature{OwnerPubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		Signature:     "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		SignTimestamp: common.GenTimestamp()}

	contractBody.ContractSignatures = contractSignatures

	contractComponents := []*ContractComponent{{
		Cname:       "AXXXXXX",
		Ctype:       "Task_Enquiry",
		Caption:     "",
		Description: "XXXXX",
		State:       "TaskState_Dormant",
		PreCondition: []*ComponentsExpression{
			{
				Cname:         "xxxx",
				Ctype:         "Expression_LogicArgument",
				Caption:       "",
				Description:   "",
				ExpressionStr: "xxxxxx",
				ExpressionResult: &ExpressionResult{
					Messsage: "xxx",
					Code:     "23434",
				},
				LogicValue: "xx0xx",
			},
		},
		CompleteCondition: []*ComponentsExpression{
			{
				Cname:         "xxxx",
				Ctype:         "Expression_LogicArgument",
				Caption:       "",
				Description:   "",
				ExpressionStr: "xxxxxx",
				ExpressionResult: &ExpressionResult{
					Messsage: "xxx",
					Code:     "23434",
				},
			},
		},
		DisgardCondition: []*ComponentsExpression{
			{
				Cname:         "xxxx",
				Ctype:         "Expression_LogicArgument",
				Caption:       "",
				Description:   "",
				ExpressionStr: "xxxxxx",
				ExpressionResult: &ExpressionResult{
					Messsage: "xxx",
					Code:     "23434",
				},
			},
		},
		NextTasks: []string{"Axxxx", "Bxxxx"},
		DataList: []*ComponentData{
			{
				Cname:        "xxxx",
				Ctype:        "Expression_LogicArgument",
				Caption:      "",
				Description:  "",
				ModifyDate:   common.GenTimestamp(),
				HardConvType: "xxxx000000xxxx",
				//Category:map[string]string{
				//	"name":"12323",
				//	"age":"wwwww",
				//},
				Parent:    nil,
				Mandatory: true,
				//DefaultValue
				Unit: "int64",
				// Value
				Options: map[string]int32{
					"name": 1212,
					"age":  3333,
				},
				// DataRange
				Format: "gogogog",
			},
		},
		DataValueSetterExpressionList: []*ComponentsExpression{
			{
				Cname:         "xxxx",
				Ctype:         "Expression_LogicArgument",
				Caption:       "",
				Description:   "",
				ExpressionStr: "xxxxxx",
				ExpressionResult: &ExpressionResult{
					Messsage: "xxx",
					Code:     "23434",
				},
			},
		},
	},
		{
			Cname:       "AXXXXXX",
			Ctype:       "Task_Action",
			Caption:     "",
			Description: "XXXXX",
			State:       "TaskState_Dormant",
			PreCondition: []*ComponentsExpression{
				{
					Cname:         "xxxx",
					Ctype:         "Expression_LogicArgument",
					Caption:       "",
					Description:   "",
					ExpressionStr: "xxxxxx",
					ExpressionResult: &ExpressionResult{
						Messsage: "xxx",
						Code:     "23434",
					},
				},
			},
			CompleteCondition: []*ComponentsExpression{
				{
					Cname:         "xxxx",
					Ctype:         "Expression_LogicArgument",
					Caption:       "",
					Description:   "",
					ExpressionStr: "xxxxxx",
					ExpressionResult: &ExpressionResult{
						Messsage: "xxx",
						Code:     "23434",
					},
				},
			},
			DisgardCondition: []*ComponentsExpression{
				{
					Cname:         "xxxx",
					Ctype:         "Expression_LogicArgument",
					Caption:       "",
					Description:   "",
					ExpressionStr: "xxxxxx",
					ExpressionResult: &ExpressionResult{
						Messsage: "xxx",
						Code:     "23434",
					},
				},
			},
			NextTasks: []string{"Axxxx", "Bxxxx"},
			DataList:  nil,
			DataValueSetterExpressionList: []*ComponentsExpression{
				{
					Cname:         "xxxx",
					Ctype:         "Expression_LogicArgument",
					Caption:       "",
					Description:   "",
					ExpressionStr: "xxxxxx",
					ExpressionResult: &ExpressionResult{
						Messsage: "xxx",
						Code:     "23434",
					},
				},
			},
		},
		{
			Cname:       "AXXXXXX",
			Ctype:       "Task_Decision",
			Caption:     "",
			Description: "XXXXX",
			State:       "TaskState_Dormant",
			PreCondition: []*ComponentsExpression{
				{
					Cname:         "xxxx",
					Ctype:         "Expression_LogicArgument",
					Caption:       "",
					Description:   "",
					ExpressionStr: "xxxxxx",
					ExpressionResult: &ExpressionResult{
						Messsage: "xxx",
						Code:     "23434",
					},
				},
			},
			CompleteCondition: []*ComponentsExpression{
				{
					Cname:         "xxxx",
					Ctype:         "Expression_LogicArgument",
					Caption:       "",
					Description:   "",
					ExpressionStr: "xxxxxx",
					ExpressionResult: &ExpressionResult{
						Messsage: "xxx",
						Code:     "23434",
					},
				},
			},
			DisgardCondition: []*ComponentsExpression{
				{
					Cname:         "xxxx",
					Ctype:         "Expression_LogicArgument",
					Caption:       "",
					Description:   "",
					ExpressionStr: "xxxxxx",
					ExpressionResult: &ExpressionResult{
						Messsage: "xxx",
						Code:     "23434",
					},
				},
			},
			NextTasks: []string{"Axxxx", "Bxxxx"},
			DataList:  nil,
			DataValueSetterExpressionList: []*ComponentsExpression{
				{
					Cname:         "xxxx",
					Ctype:         "Expression_LogicArgument",
					Caption:       "",
					Description:   "",
					ExpressionStr: "xxxxxx",
					ExpressionResult: &ExpressionResult{
						Messsage: "xxx",
						Code:     "23434",
					},
				},
			},
			// repeated ? CandidateList = 12; //todo
			// repeated ? DecisionResult = 13; //todo
			// repeated ? TaskList = 14; //todo
			TaskList:         []string{"1", "2", "5"},
			SupportArguments: []string{"int", "float"},
			Support:          2,
			Text:             []string{"1.go", "2.gogo"},
		},
	}

	contractBody.ContractComponents = contractComponents
	contract.ContractHead = contractHead
	contract.ContractBody = contractBody
	//fmt.Println(contractProto)
	fmt.Println(common.SerializePretty(contract))
	fmt.Println(pubkey)
}
