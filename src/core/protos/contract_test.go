package protos

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"testing"
	"unicontract/src/common"
)

func Test_ContractProto(t *testing.T) {
	contract := ContractProto{ // golang
		//contract := &ContractProto{ // proto-buf
		Id:         "2",
		NodePubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		MainPubkey: "93TEovPuYo6BQFm4ia9ta4qtL1TbAmnk9fV5kxmesAG5",
		Signature:  "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		Voters: []string{
			"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
			"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
			"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
		},
		Timestamp: common.GenTimestamp(),
		Version:   1,
		Contract: &Contract{
			CreatorPubkey:   "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
			CreateTimestamp: common.GenTimestamp(),
			Operation:       "CREATE",
			ContractAttributes: &ContractAttributes{
				Name:           "XXXXXX",
				StartTimestamp: common.GenTimestamp(),
				EndTimestamp:   common.GenTimestamp(),
			},
			ContractOwners: []string{
				"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
				"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
				"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
				"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
			},
			ContractSignatures: []*ContractSignature{
				{
					OwnerPubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
					Signature:   "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
					Timestamp:   string(common.GenTimestamp()),
				},
				{
					OwnerPubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
					Signature:   "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
					Timestamp:  string(common.GenTimestamp()),
				},
				{
					OwnerPubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
					Signature:   "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
					Timestamp:   string(common.GenTimestamp()),
				},
			},
			ContractAsserts: []*ContractAssert{
				{
					Id:       "111",
					Name:     "wx1",
					Amount:   1000,
					Metadata: nil,
				},
				{
					Id:     "113",
					Name:   "wx2",
					Amount: 800,
				},
				{
					Id:     "112",
					Name:   "wx3",
					Amount: 100,
				},
			},
			ContractComponents: &ContractComponents{
				Plans: []*Plan{
					{
						Id:          "ID_Axxxxxx",
						Type:        "PLAN",
						State:       "dormant",
						Name:        "N_Axxxxx",
						Description: "xxxxx",
						Condition: []*PlanTaskCondition{
							{
								Id:          "1",
								Type:        "PreCondition",
								Name:        "XXXX",
								Value:       "XXXX",
								Description: "xxxxx",
							},
							{
								Id:          "2",
								Type:        "DisgardCondition",
								Name:        "XXXX",
								Value:       "XXXX",
								Description: "xxxxx",
							},
							{
								Id:          "3",
								Type:        "CompleteCondition",
								Name:        "XXXXprotos.",
								Value:       "XXXX",
								Description: "xxxxx",
							},
						},
						Level:        1,
						ContractType: "RIGHT",
						NextTask:     []string{"1", "2"},
					},
					{
						Id:          "ID_Bxxxxxx",
						Type:        "PLAN",
						State:       "dormant",
						Name:        "N_Bxxxxx",
						Description: "xxxxx",
						Condition: []*PlanTaskCondition{
							{
								Id:          "1",
								Type:        "PreCondition",
								Name:        "XXXX",
								Value:       "XXXX",
								Description: "xxxxx",
							},
							{
								Id:          "2",
								Type:        "DisgardCondition",
								Name:        "XXXX",
								Value:       "XXXX",
								Description: "xxxxx",
							},
							{
								Id:          "3",
								Type:        "CompleteCondition",
								Name:        "XXXX",
								Value:       "XXXX",
								Description: "xxxxx",
							},
						},
						Level:        1,
						ContractType: "RIGHT",
						NextTask:     nil,
					},
				},
				Tasks: []*Task{
					{
						Id:          "ID_Cxxxxxx",
						Type:        "ENQUIRY",
						State:       "dormant",
						Name:        "Axxxxxx",
						Description: "xxxxx",
						Condition: []*PlanTaskCondition{
							{
								Id:          "1",
								Type:        "PreCondition",
								Name:        "XXXX",
								Value:       "XXXX",
								Description: "xxxxx",
							},
							{
								Id:          "2",
								Type:        "DisgardCondition",
								Name:        "XXXX",
								Value:       "XXXX",
								Description: "xxxxx",
							},
							{
								Id:          "3",
								Type:        "CompleteCondition",
								Name:        "XXXX",
								Value:       "XXXX",
								Description: "xxxxx",
							},
						},
						Level:        1,
						ContractType: "RIGHT",
						NextTask: []string{
							"Axxxxxx",
							"Bxxxxxx",
						},
					},
					{
						Id:          "ID_Cxxxxxx",
						Type:        "ENQUIRY",
						State:       "dormant",
						Name:        "Bxxxxxx",
						Description: "xxxxx",
						Condition: []*PlanTaskCondition{
							{
								Id:          "1",
								Type:        "PreCondition",
								Name:        "XXXX",
								Value:       "XXXX",
								Description: "xxxxx",
							},
							{
								Id:          "2",
								Type:        "DisgardCondition",
								Name:        "XXXX",
								Value:       "XXXX",
								Description: "xxxxx",
							},
							{
								Id:          "3",
								Type:        "CompleteCondition",
								Name:        "XXXX",
								Value:       "XXXX",
								Description: "xxxxx",
							},
						},
						Level:        1,
						ContractType: "RIGHT",
						NextTask:     []string{"", ""},
					},
					{Id: "ID_Cxxxxxx",
						Type:        "ACTION",
						State:       "dormant",
						Name:        "Cxxxxxx",
						Description: "xxxxx",
						Condition: []*PlanTaskCondition{
							{
								Id:          "1",
								Type:        "PreCondition",
								Name:        "XXXX",
								Value:       "XXXX",
								Description: "xxxxx"},
							{
								Id:          "2",
								Type:        "DisgardCondition",
								Name:        "XXXX",
								Value:       "XXXX",
								Description: "xxxxx",
							},
							{
								Id:          "3",
								Type:        "CompleteCondition",
								Name:        "XXXX",
								Value:       "XXXX",
								Description: "xxxxx",
							},
						},
						Level:        1,
						ContractType: "DUTY",
						NextTask:     []string{"", ""},
					},
				},
			},
		},
	}

	//fmt.Println(common.Serialize(contract))
	//fmt.Println(common.SerializePretty(contract))
	result, err := proto.Marshal(&contract)
	if err != nil {
		fmt.Println("error ", err.Error())
	}

	//test---------ContractData Start--------------
	data := ContractData{
		Data: &contract,
		Token: "ZDNkM0xtWjFkSFZ5WlhabGNpNWpiMjA9",
	}
	fmt.Println("data is\n", data)
	requestBody, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	fmt.Println(requestBody)
	result, err = proto.Marshal(&data)
	if err != nil {
		fmt.Println("error %s", err.Error())
	}
	var contractData ContractData
	proto.Unmarshal(result, &contractData)
	fmt.Printf("proto deserialize contract content is %s\n", contractData)
	contractData_str := common.Serialize(contractData)

	fmt.Printf("contract json len is %d, pretty json len is %d, proto-buf len is %d and origin content is \n%v",
		1, 1, len(contractData_str), contractData_str)
	//return
	//test-------------ContractData End----------

	contract_json_len := len(common.Serialize(contract))
	contract_pretty_json_len := len(common.SerializePretty(contract))

	fmt.Printf("contract json len is %d, pretty json len is %d, proto-buf len is %d and content is %v\n",
		contract_json_len, contract_pretty_json_len, len(result), result)
	fmt.Printf("contract json len is %d, pretty json len is %d, proto-buf len is %d and content is %v\n",
		contract_json_len, contract_pretty_json_len, len(result), string(result))

	var origin_contract ContractProto
	proto.Unmarshal(result, &origin_contract)
	fmt.Printf("proto deserialize contract content is %s\n", origin_contract)
	origin_contract_str := common.Serialize(origin_contract)

	fmt.Printf("contract json len is %d, pretty json len is %d, proto-buf len is %d and origin content is \n%v",
		contract_json_len, contract_pretty_json_len, len(origin_contract_str), origin_contract_str)

	//origin_contract_pretty_str := common.SerializePretty(origin_contract)
	//fmt.Printf("contract json len is %d, pretty json len is %d, proto-buf len is %d and origin content is \n%v",
	//	contract_json_len, contract_pretty_json_len, len(origin_contract_pretty_str), origin_contract_pretty_str)

}
