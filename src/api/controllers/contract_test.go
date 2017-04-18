package controllers

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"testing"
	"unicontract/src/core/protos"
	"unicontract/src/common"
)


// application content-type
const (
	APPLICATION_X_PROTOBUF   = "application/x-protobuf"
	APPLICATION_JSON         = "application/json"
	APPLICATION_OCTET_STREAM = "application/octet-stream"
)

func httpRequest(method string, urlStr string, body []byte, contentType string) ([]byte, error) {
	client := &http.Client{}
	req_body := bytes.NewReader(body)

	if method == "" {
		method = "POST"
	}
	req, err := http.NewRequest(method, urlStr, req_body)
	if err != nil {
		// handle error
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("RequestDataType", "proto")

	resp, err := client.Do(req)

	//fmt.Printf("response: code:%v, header:%v\n", resp.StatusCode, resp.Header)
	fmt.Printf("Request %s [%s] content-type=%s\n", urlStr, method, contentType)

	defer resp.Body.Close()
	//
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	if responseBody == nil || string(responseBody) == "null" {
		//fmt.Printf("Response: %v, body %v\n", resp.StatusCode, string(responseBody))
		//fmt.Printf("Response: %v, body %v\n", resp.StatusCode, string(responseBody))
	} else {
		fmt.Printf("Response: %v\n", resp.StatusCode)
		fmt.Printf("body:\n%s\n\n", responseBody)
	}
	//fmt.Println(string(responseBody))
	return responseBody, err
}

func Test_ContractProto(t *testing.T) {
	contract := protos.ContractProto{ // golang
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
		Version:   "v1.0",
		Contract: &protos.Contract{
			CreatorPubkey:   "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
			CreateTimestamp: common.GenTimestamp(),
			Operation:       "CREATE",
			ContractAttributes: &protos.ContractAttributes{
				Name:           "XXXXXX",
				StartTimestamp: common.GenTimestamp(),
				EndTimestamp:   common.GenTimestamp(),
			},
			//ContractOwners: []string{
			//	"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
			//	"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
			//	"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
			//	"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
			//},
			//ContractSignatures: []*ContractSignature{
			//	{
			//		OwnerPubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
			//		Signature:   "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
			//		Timestamp:   common.GenTimestamp(),
			//	},
			//	{
			//		OwnerPubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
			//		Signature:   "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
			//		Timestamp:   common.GenTimestamp(),
			//	},
			//	{
			//		OwnerPubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
			//		Signature:   "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
			//		Timestamp:   common.GenTimestamp(),
			//	},
			//},
			//	ContractAsserts: []*ContractAssert{
			//		{
			//			Id:       "111",
			//			Name:     "wx1",
			//			Amount:   1000,
			//			Metadata: nil,
			//		},
			//		{
			//			Id:     "113",
			//			Name:   "wx2",
			//			Amount: 800,
			//		},
			//		{
			//			Id:     "112",
			//			Name:   "wx3",
			//			Amount: 100,
			//		},
			//	},
			//	ContractComponents: &ContractComponents{
			//		Plans: []*Plan{
			//			{
			//				Id:          "ID_Axxxxxx",
			//				Type:        "PLAN",
			//				State:       "dormant",
			//				Name:        "N_Axxxxx",
			//				Description: "xxxxx",
			//				Condition: []*PlanTaskCondition{
			//					{
			//						Id:          "1",
			//						Type:        "PreCondition",
			//						Name:        "XXXX",
			//						Value:       "XXXX",
			//						Description: "xxxxx",
			//					},
			//					{
			//						Id:          "2",
			//						Type:        "DisgardCondition",
			//						Name:        "XXXX",
			//						Value:       "XXXX",
			//						Description: "xxxxx",
			//					},
			//					{
			//						Id:          "3",
			//						Type:        "CompleteCondition",
			//						Name:        "XXXX",
			//						Value:       "XXXX",
			//						Description: "xxxxx",
			//					},
			//				},
			//				Level:        1,
			//				ContractType: "RIGHT",
			//				NextTask:     []string{"1", "2"},
			//			},
			//			{
			//				Id:          "ID_Bxxxxxx",
			//				Type:        "PLAN",
			//				State:       "dormant",
			//				Name:        "N_Bxxxxx",
			//				Description: "xxxxx",
			//				Condition: []*PlanTaskCondition{
			//					{
			//						Id:          "1",
			//						Type:        "PreCondition",
			//						Name:        "XXXX",
			//						Value:       "XXXX",
			//						Description: "xxxxx",
			//					},
			//					{
			//						Id:          "2",
			//						Type:        "DisgardCondition",
			//						Name:        "XXXX",
			//						Value:       "XXXX",
			//						Description: "xxxxx",
			//					},
			//					{
			//						Id:          "3",
			//						Type:        "CompleteCondition",
			//						Name:        "XXXX",
			//						Value:       "XXXX",
			//						Description: "xxxxx",
			//					},
			//				},
			//				Level:        1,
			//				ContractType: "RIGHT",
			//				NextTask:     nil,
			//			},
			//		},
			//		Tasks: []*Task{
			//			{
			//				Id:          "ID_Cxxxxxx",
			//				Type:        "ENQUIRY",
			//				State:       "dormant",
			//				Name:        "Axxxxxx",
			//				Description: "xxxxx",
			//				Condition: []*PlanTaskCondition{
			//					{
			//						Id:          "1",
			//						Type:        "PreCondition",
			//						Name:        "XXXX",
			//						Value:       "XXXX",
			//						Description: "xxxxx",
			//					},
			//					{
			//						Id:          "2",
			//						Type:        "DisgardCondition",
			//						Name:        "XXXX",
			//						Value:       "XXXX",
			//						Description: "xxxxx",
			//					},
			//					{
			//						Id:          "3",
			//						Type:        "CompleteCondition",
			//						Name:        "XXXX",
			//						Value:       "XXXX",
			//						Description: "xxxxx",
			//					},
			//				},
			//				Level:        1,
			//				ContractType: "RIGHT",
			//				NextTask: []string{
			//					"Axxxxxx",
			//					"Bxxxxxx",
			//				},
			//			},
			//			{
			//				Id:          "ID_Cxxxxxx",
			//				Type:        "ENQUIRY",
			//				State:       "dormant",
			//				Name:        "Bxxxxxx",
			//				Description: "xxxxx",
			//				Condition: []*PlanTaskCondition{
			//					{
			//						Id:          "1",
			//						Type:        "PreCondition",
			//						Name:        "XXXX",
			//						Value:       "XXXX",
			//						Description: "xxxxx",
			//					},
			//					{
			//						Id:          "2",
			//						Type:        "DisgardCondition",
			//						Name:        "XXXX",
			//						Value:       "XXXX",
			//						Description: "xxxxx",
			//					},
			//					{
			//						Id:          "3",
			//						Type:        "CompleteCondition",
			//						Name:        "XXXX",
			//						Value:       "XXXX",
			//						Description: "xxxxx",
			//					},
			//				},
			//				Level:        1,
			//				ContractType: "RIGHT",
			//				NextTask:     []string{"", ""},
			//			},
			//			{Id: "ID_Cxxxxxx",
			//				Type:        "ACTION",
			//				State:       "dormant",
			//				Name:        "Cxxxxxx",
			//				Description: "xxxxx",
			//				Condition: []*PlanTaskCondition{
			//					{
			//						Id:          "1",
			//						Type:        "PreCondition",
			//						Name:        "XXXX",
			//						Value:       "XXXX",
			//						Description: "xxxxx"},
			//					{
			//						Id:          "2",
			//						Type:        "DisgardCondition",
			//						Name:        "XXXX",
			//						Value:       "XXXX",
			//						Description: "xxxxx",
			//					},
			//					{
			//						Id:          "3",
			//						Type:        "CompleteCondition",
			//						Name:        "XXXX",
			//						Value:       "XXXX",
			//						Description: "xxxxx",
			//					},
			//				},
			//				Level:        1,
			//				ContractType: "DUTY",
			//				NextTask:     []string{"", ""},
			//			},
			//		},
			//},
		},
	}

	//fmt.Println(common.Serialize(contract))
	//fmt.Println(common.SerializePretty(contract))
	data := protos.ContractData{
		Data: &contract,
		Token: "ZDNkM0xtWjFkSFZ5WlhabGNpNWpiMjA9",
	}
	result, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	fmt.Println("input result is ", result)

}


var default_url = "http://192.168.1.14:8088/v1/contract/"

func Test_AuthSignature(t *testing.T) {
	url := default_url + "authSignature"
	fmt.Println(url)
	contract := protos.ContractProto{ // proto-buf
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
		Version:   "v1.0",
	}

	data := protos.ContractData{
		Data: &contract,
		Token: "ZDNkM0xtWjFkSFZ5WlhabGNpNWpiMjA9",
	}
	requestBody, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	fmt.Println(requestBody)
	_, err = httpRequest("POST", url, requestBody, APPLICATION_X_PROTOBUF)
	if err != nil {
		// handle error
		fmt.Println("error ", err.Error())
	}
}

func Test_Creat(t *testing.T) {
	url := default_url + "create"
	fmt.Println(url)
	contract := protos.ContractProto{ // proto-buf
		//Id:         "2",
		//NodePubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		//MainPubkey: "93TEovPuYo6BQFm4ia9ta4qtL1TbAmnk9fV5kxmesAG5",
		//Signature:  "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		//Voters: []string{
		//	"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		//	"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		//	"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
		//},
		//Timestamp: common.GenTimestamp(),
		//Version:   "v1.0",
		Contract: &protos.Contract{
			CreatorPubkey:   "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
			CreateTimestamp: common.GenTimestamp(),
			Operation:       "CREATE",
			ContractAttributes: &protos.ContractAttributes{
				Name:           "XXXXXX",
				StartTimestamp: common.GenTimestamp(),
				EndTimestamp:   common.GenTimestamp(),
			},
		},
	}
	//requestBody, err := proto.Marshal(&contract)
	//if err != nil {
	//	fmt.Println("error ", err.Error())
	//}
	//fmt.Println(requestBody)
	fmt.Println(contract)

	data := protos.ContractData{
		Data: &contract,
		Token: "ZDNkM0xtWjFkSFZ5WlhabGNpNWpiMjA9",
	}
	requestBody, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	fmt.Println(requestBody)

	_, err = httpRequest("POST", url, requestBody, APPLICATION_X_PROTOBUF)
	//_, err = httpRequest("POST", url, requestBody, APPLICATION_X_PROTOBUF)
	if err != nil {
		// handle error
		fmt.Println("error ", err.Error())
	}

}

func Test_Signature(t *testing.T) {
	url := default_url + "signature"

	contract := protos.ContractProto{ // proto-buf
		Id:         "2",
		NodePubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		MainPubkey: "93TEovPuYo6BQFm4ia9ta4qtL1TbAmnk9fV5kxmesAG5",
		//Signature:  "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		Voters: []string{
			"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
			"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
			"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
		},
		Timestamp: common.GenTimestamp(),
		Version:   "v1.0",
	}

	data := protos.ContractData{
		Data: &contract,
		Token: "ZDNkM0xtWjFkSFZ5WlhabGNpNWpiMjA9",
	}

	requestBody, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("error ", err.Error())
	}

	_, err = httpRequest("POST", url, requestBody, APPLICATION_X_PROTOBUF)
	if err != nil {
		// handle error
		fmt.Println("error ", err.Error())
	}
	//fmt.Println("response is:", string(response))
}

func Test_Terminate(t *testing.T) {
	url := default_url + "terminate"

	contract := protos.ContractProto{ // proto-buf
		Id:         "2",
		//NodePubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		//MainPubkey: "93TEovPuYo6BQFm4ia9ta4qtL1TbAmnk9fV5kxmesAG5",
		//Signature:  "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		//Voters: []string{
		//	"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		//	"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		//	"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
		//},
		//Timestamp: common.GenTimestamp(),
		//Version:   "v1.0",
	}
	data := protos.ContractData{
		Data: &contract,
		Token: "ZDNkM0xtWjFkSFZ5WlhabGNpNWpiMjA9",
	}

	requestBody, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("error ", err.Error())
	}

	_, err = httpRequest("POST", url, requestBody, APPLICATION_X_PROTOBUF)
	if err != nil {
		// handle error
		fmt.Println("error ", err.Error())
	}
	//fmt.Println("response is:", string(response))
}

func Test_Find(t *testing.T) {
	url := default_url + "find"

	contract := protos.ContractProto{ // proto-buf
		Id:         "2",
	}
	data := protos.ContractData{
		Data: &contract,
		Token: "ZDNkM0xtWjFkSFZ5WlhabGNpNWpiMjA9",
	}

	requestBody, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("error ", err.Error())
	}

	_, err = httpRequest("POST", url, requestBody, APPLICATION_X_PROTOBUF)
	if err != nil {
		// handle error
		fmt.Println("error ", err.Error())
	}
	//fmt.Println("response is:", string(response))
}

func Test_Track(t *testing.T) {
	url := default_url + "track"

	contract := protos.ContractProto{ // proto-buf
		Id:         "2",
	}
	data := protos.ContractData{
		Data: &contract,
		Token: "ZDNkM0xtWjFkSFZ5WlhabGNpNWpiMjA9",
	}
	requestBody, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("error ", err.Error())
	}

	_, err = httpRequest("POST", url, requestBody, APPLICATION_X_PROTOBUF)
	if err != nil {
		// handle error
		fmt.Println("error ", err.Error())
	}
	//fmt.Println("response is:", string(response))
}

func Test_Update(t *testing.T) {
	url := default_url + "update"

	contract := protos.ContractProto{ // proto-buf
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
		Version:   "v1.0",
	}
	data := protos.ContractData{
		Data: &contract,
		Token: "ZDNkM0xtWjFkSFZ5WlhabGNpNWpiMjA9",
	}
	requestBody, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("error ", err.Error())
	}

	_, err = httpRequest("POST", url, requestBody, APPLICATION_X_PROTOBUF)
	if err != nil {
		// handle error
		fmt.Println("error ", err.Error())
	}
	//fmt.Println("response is:", string(response))
}

func Test_Test(t *testing.T) {
	url := default_url + "test"

	contract := protos.ContractProto{ // proto-buf
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
		Version:   "v1.0",
	}
	data := protos.ContractData{
		Data: &contract,
		Token: "ZDNkM0xtWjFkSFZ5WlhabGNpNWpiMjA9",
	}
	requestBody, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("error ", err.Error())
	}

	_, err = httpRequest("POST", url, requestBody, APPLICATION_X_PROTOBUF)
	if err != nil {
		// handle error
		fmt.Println("error ", err.Error())
	}
	//fmt.Println("response is:", string(response))
}
