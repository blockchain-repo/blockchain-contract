package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"testing"
	"unicontract/src/common"
	"unicontract/src/core/protos"
	"github.com/astaxie/beego"
	"unicontract/src/core/model"
	//"unicontract/src/core"
)

// application content-type
const (
	APPLICATION_X_PROTOBUF   = "application/x-protobuf"
	APPLICATION_JSON         = "application/json"
	APPLICATION_OCTET_STREAM = "application/octet-stream"
)

func httpRequest(method string, urlStr string, body []byte, requestHead map[string]string) ([]byte, error) {
	client := &http.Client{}
	req_body := bytes.NewReader(body)

	if method == "" {
		method = "POST"
	}
	req, err := http.NewRequest(method, urlStr, req_body)
	if err != nil {
		// handle error
	}
	contentType := requestHead["Content-Type"]
	if contentType == "" {
		contentType = APPLICATION_X_PROTOBUF
		//contentType = APPLICATION_JSON
	}
	requestDataType := requestHead["RequestDataType"]
	if requestDataType == "" {
		requestDataType = "proto"
	}
	token := requestHead["token"]
	if token == "" {
		token = "futurever"
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("RequestDataType", requestDataType)
	req.Header.Set("token", token)

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

func Test_Contract(t *testing.T) {
	contract := protos.Contract{ // golang
		//Id: "2",
		ContractHead: &protos.ContractHead{
			MainPubkey: "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			Version:    1,
		},
		ContractBody: &protos.ContractBody{
			Caption: "CREATE",
			Cname:   "futurever",
			ContractAssets: []*protos.ContractAsset{
				{
					AssetId:     "001",
					Name:        "futurever-1",
					Amount:      1000,
					Caption:     "futurever",
					Description: "",
					Unit:        "int32",
					MetaData:    nil,
				},
				{
					AssetId:     "003",
					Name:        "futurever-3",
					Amount:      452,
					Caption:     "futurever",
					Description: "",
					Unit:        "int32",
					MetaData:    nil,
				},
				{
					AssetId:     "002",
					Name:        "futurever-2",
					Amount:      99999,
					Caption:     "futurever",
					Description: "",
					Unit:        "int32",
					MetaData:    nil,
				},
			},
			ContractComponents: nil,
			ContractId:         common.GenerateUUID(),
			ContractOwners: []string{
				"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
				"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			},
			ContractSignatures: []*protos.ContractSignature{
				{
					OwnerPubkey:   "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
					Signature:     "65D27HW4uXYvkekGssAQB93D92onMyU1NVnCJnE1PgRKz2uFSPZ6aQvid4qZvkxys7G4r2Mf2KFn5BSQyEBhWs34",
					SignTimestamp: common.GenTimestamp(),
				},
				{
					OwnerPubkey:   "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
					Signature:     "5i5dTtQseQjWZ8UdchqQtgttyeeFmB3LDFYzNKafvV2YvTqwv4wZ9mFsH7qgysV9ow893D1h2Xnt1uCXLHtbKrkT",
					SignTimestamp: common.GenTimestamp(),
				},
			},
			ContractState: "",
			Creator:       "futurever",
			CreatorTime:   common.GenTimestamp(),
			Ctype:         "CONTRACT",
			Description:   "CREATE CONTRACT BY futurever [合约创建]",
			StartTime:     common.GenTimestamp(),
			EndTime:       common.GenTimestamp(),
		},

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
	}

	//fmt.Println(common.Serialize(contract))
	//fmt.Println(common.SerializePretty(contract))

	contract.Id = common.HashData(common.Serialize(contract.ContractBody))

	data := protos.ContractData{
		Data:  &contract,
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
	contract := protos.Contract{ // proto-buf
		Id: "2",
		ContractHead: &protos.ContractHead{
			MainPubkey: "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			Version:    1,
		},
		ContractBody: &protos.ContractBody{
			Caption: "CREATE",
			Cname:   "futurever",
			//ContractAssets: []*protos.ContractAsset{
			//	{
			//		AssetId:     "001",
			//		Name:        "futurever-1",
			//		Amount:      1000,
			//		Caption:     "futurever",
			//		Description: "",
			//		Unit:        "int32",
			//		MetaData:    nil,
			//	},
			//	{
			//		AssetId:     "003",
			//		Name:        "futurever-3",
			//		Amount:      452,
			//		Caption:     "futurever",
			//		Description: "",
			//		Unit:        "int32",
			//		MetaData:    nil,
			//	},
			//	{
			//		AssetId:     "002",
			//		Name:        "futurever-2",
			//		Amount:      99999,
			//		Caption:     "futurever",
			//		Description: "",
			//		Unit:        "int32",
			//		MetaData:    nil,
			//	},
			//},
			ContractComponents: nil,
			ContractId:         common.GenerateUUID(),
			ContractOwners: []string{
				"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
				"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			},
			ContractSignatures: []*protos.ContractSignature{
				{
					OwnerPubkey:   "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
					Signature:     "65D27HW4uXYvkekGssAQB93D92onMyU1NVnCJnE1PgRKz2uFSPZ6aQvid4qZvkxys7G4r2Mf2KFn5BSQyEBhWs34",
					SignTimestamp: common.GenTimestamp(),
				},
				{
					OwnerPubkey:   "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
					Signature:     "5i5dTtQseQjWZ8UdchqQtgttyeeFmB3LDFYzNKafvV2YvTqwv4wZ9mFsH7qgysV9ow893D1h2Xnt1uCXLHtbKrkT",
					SignTimestamp: common.GenTimestamp(),
				},
			},
			ContractState: "",
			Creator:       "futurever",
			CreatorTime:   common.GenTimestamp(),
			Ctype:         "CONTRACT",
			Description:   "CREATE CONTRACT BY futurever [合约创建]",
			StartTime:     common.GenTimestamp(),
			EndTime:       common.GenTimestamp(),
		},
	}

	data := protos.ContractData{
		Data:  &contract,
		Token: "ZDNkM0xtWjFkSFZ5WlhabGNpNWpiMjA9",
	}
	requestBody, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	fmt.Println(requestBody)
	requestHead := make(map[string]string)
	requestHead["Content-Type"] = APPLICATION_X_PROTOBUF

	_, err = httpRequest("POST", url, requestBody, requestHead)
	if err != nil {
		// handle error
		fmt.Println("error ", err.Error())
	}
}

func Test_CreatValid(t *testing.T) {
	url := default_url + "create"
	beego.Debug("请求地址", url)
	contractModel := model.ContractModel{}
	//contractAsset := []*protos.ContractAsset{}
	//contractComponent:=[]*protos.ContractComponent{}

	contractHead := &protos.ContractHead{"",1}

	contractOwners := []string{
		"BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
		//"4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		"9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
	}
	contractBody := &protos.ContractBody{
		ContractId:"UUID-1234-5678-90",
		Cname:"test contract output",
		Ctype:"CREATE",
		Caption:"购智能手机返话费合约产品协议",
		Description:"移动用户A花费500元购买移动运营商B的提供的合约智能手机C后",
		ContractState:"",
		Creator:common.GenTimestamp(),
		CreatorTime:common.GenTimestamp(),
		StartTime:common.GenTimestamp(),
		EndTime:common.GenTimestamp(),
		ContractOwners:contractOwners,
		ContractSignatures:nil,
		ContractAssets:nil,
		ContractComponents: nil,
	}

	contractModel.ContractHead = contractHead
	contractModel.ContractBody =contractBody

	contractSignatures := []*protos.ContractSignature{
		{
			OwnerPubkey:   "BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
			Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
			SignTimestamp: common.GenTimestamp(),
		},
		//{
		//	OwnerPubkey:   "4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		//	Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
		{
			OwnerPubkey:   "9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
			Signature:     contractModel.Sign("9647UfPdDSwBf5kw7tUrSe7cmYY5RvVX47GrGqSh4XVi"),
			SignTimestamp: common.GenTimestamp(),
		},
	}
	contractBody.ContractSignatures = contractSignatures
	contractModel.Id = contractModel.GenerateId()
	beego.Info(common.Serialize(contractModel))
	protoContract, _ := fromContractModelStrToContract(common.Serialize(contractModel))
	requestBody, err := proto.Marshal(&protoContract)
	if err != nil {
		beego.Error("proto.Marshal", err)
	}
	requestHead := make(map[string]string)
	requestHead["Content-Type"] = APPLICATION_X_PROTOBUF
	_, err = httpRequest("POST", url, requestBody, requestHead)
	if err != nil {
		// handle error
		fmt.Println("error ", err.Error())
		fmt.Println("handle error ", err.Error())
	}

}


func Test_Creat(t *testing.T) {
	url := default_url + "create"
	fmt.Println(url)
	contract := protos.Contract{ // proto-buf
		ContractHead: &protos.ContractHead{
			MainPubkey: "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			Version:    1,
		},
		ContractBody: &protos.ContractBody{
			Caption: "CREATE",
			Cname:   "futurever",
			ContractAssets: []*protos.ContractAsset{
				{
					AssetId:     "001",
					Name:        "futurever-1",
					Amount:      1000,
					Caption:     "futurever",
					Description: "",
					Unit:        "int32",
					MetaData:    nil,
				},
				{
					AssetId:     "003",
					Name:        "futurever-3",
					Amount:      452,
					Caption:     "futurever",
					Description: "",
					Unit:        "int32",
					MetaData:    nil,
				},
				{
					AssetId:     "002",
					Name:        "futurever-2",
					Amount:      99999,
					Caption:     "futurever",
					Description: "",
					Unit:        "int32",
					MetaData:    nil,
				},
			},
			ContractComponents: nil,
			ContractId:         common.GenerateUUID(),
			ContractOwners: []string{
				"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
				"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			},
			ContractSignatures: []*protos.ContractSignature{
				{
					OwnerPubkey:   "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
					Signature:     "65D27HW4uXYvkekGssAQB93D92onMyU1NVnCJnE1PgRKz2uFSPZ6aQvid4qZvkxys7G4r2Mf2KFn5BSQyEBhWs34",
					SignTimestamp: common.GenTimestamp(),
				},
				{
					OwnerPubkey:   "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
					Signature:     "5i5dTtQseQjWZ8UdchqQtgttyeeFmB3LDFYzNKafvV2YvTqwv4wZ9mFsH7qgysV9ow893D1h2Xnt1uCXLHtbKrkT",
					SignTimestamp: common.GenTimestamp(),
				},
			},
			ContractState: "",
			Creator:       "futurever",
			CreatorTime:   common.GenTimestamp(),
			Ctype:         "CONTRACT",
			Description:   "CREATE CONTRACT BY futurever [合约创建]",
			StartTime:     common.GenTimestamp(),
			EndTime:       common.GenTimestamp(),
		},
	}

	contract.Id = common.HashData(common.Serialize(contract.ContractBody))

	requestBody, err := proto.Marshal(&contract)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	//fmt.Println(requestBody)

	fmt.Println(requestBody)
	requestHead := make(map[string]string)
	requestHead["Content-Type"] = APPLICATION_X_PROTOBUF
	_, err = httpRequest("POST", url, requestBody, requestHead)
	if err != nil {
		// handle error
		fmt.Println("error ", err.Error())
		fmt.Println("handle error ", err.Error())
	}

}

func Test_Signature(t *testing.T) {
	url := default_url + "signature"

	contract := protos.Contract{ // proto-buf
		ContractHead: &protos.ContractHead{
			MainPubkey: "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			Version:    1,
		},
		ContractBody: &protos.ContractBody{
			Caption: "CREATE",
			Cname:   "futurever",
			ContractAssets: []*protos.ContractAsset{
				{
					AssetId:     "001",
					Name:        "futurever-1",
					Amount:      1000,
					Caption:     "futurever",
					Description: "",
					Unit:        "int32",
					MetaData:    nil,
				},
				{
					AssetId:     "003",
					Name:        "futurever-3",
					Amount:      452,
					Caption:     "futurever",
					Description: "",
					Unit:        "int32",
					MetaData:    nil,
				},
				{
					AssetId:     "002",
					Name:        "futurever-2",
					Amount:      99999,
					Caption:     "futurever",
					Description: "",
					Unit:        "int32",
					MetaData:    nil,
				},
			},
			ContractComponents: nil,
			ContractId:         common.GenerateUUID(),
			ContractOwners: []string{
				"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
				"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			},
			ContractSignatures: []*protos.ContractSignature{
				{
					OwnerPubkey:   "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
					Signature:     "65D27HW4uXYvkekGssAQB93D92onMyU1NVnCJnE1PgRKz2uFSPZ6aQvid4qZvkxys7G4r2Mf2KFn5BSQyEBhWs34",
					SignTimestamp: common.GenTimestamp(),
				},
				{
					OwnerPubkey:   "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
					Signature:     "5i5dTtQseQjWZ8UdchqQtgttyeeFmB3LDFYzNKafvV2YvTqwv4wZ9mFsH7qgysV9ow893D1h2Xnt1uCXLHtbKrkT",
					SignTimestamp: common.GenTimestamp(),
				},
			},
			ContractState: "",
			Creator:       "futurever",
			CreatorTime:   common.GenTimestamp(),
			Ctype:         "CONTRACT",
			Description:   "CREATE CONTRACT BY futurever [合约创建]",
			StartTime:     common.GenTimestamp(),
			EndTime:       common.GenTimestamp(),
		},
	}

	data := protos.ContractData{
		Data:  &contract,
		Token: "ZDNkM0xtWjFkSFZ5WlhabGNpNWpiMjA9",
	}

	requestBody, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	requestHead := make(map[string]string)
	requestHead["Content-Type"] = APPLICATION_X_PROTOBUF
	_, err = httpRequest("POST", url, requestBody, requestHead)
	if err != nil {
		// handle error
		fmt.Println("error ", err.Error())
	}
	//fmt.Println("response is:", string(response))
}

func Test_Terminate(t *testing.T) {
	url := default_url + "terminate"

	contract := protos.Contract{ // proto-buf
		ContractHead: &protos.ContractHead{
			MainPubkey: "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			Version:    1,
		},
		ContractBody: &protos.ContractBody{
			Caption: "CREATE",
			Cname:   "futurever",
			ContractAssets: []*protos.ContractAsset{
				{
					AssetId:     "001",
					Name:        "futurever-1",
					Amount:      1000,
					Caption:     "futurever",
					Description: "",
					Unit:        "int32",
					MetaData:    nil,
				},
				{
					AssetId:     "003",
					Name:        "futurever-3",
					Amount:      452,
					Caption:     "futurever",
					Description: "",
					Unit:        "int32",
					MetaData:    nil,
				},
				{
					AssetId:     "002",
					Name:        "futurever-2",
					Amount:      99999,
					Caption:     "futurever",
					Description: "",
					Unit:        "int32",
					MetaData:    nil,
				},
			},
			ContractComponents: nil,
			ContractId:         common.GenerateUUID(),
			ContractOwners: []string{
				"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
				"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			},
			ContractSignatures: []*protos.ContractSignature{
				{
					OwnerPubkey:   "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
					Signature:     "65D27HW4uXYvkekGssAQB93D92onMyU1NVnCJnE1PgRKz2uFSPZ6aQvid4qZvkxys7G4r2Mf2KFn5BSQyEBhWs34",
					SignTimestamp: common.GenTimestamp(),
				},
				{
					OwnerPubkey:   "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
					Signature:     "5i5dTtQseQjWZ8UdchqQtgttyeeFmB3LDFYzNKafvV2YvTqwv4wZ9mFsH7qgysV9ow893D1h2Xnt1uCXLHtbKrkT",
					SignTimestamp: common.GenTimestamp(),
				},
			},
			ContractState: "",
			Creator:       "futurever",
			CreatorTime:   common.GenTimestamp(),
			Ctype:         "CONTRACT",
			Description:   "CREATE CONTRACT BY futurever [合约创建]",
			StartTime:     common.GenTimestamp(),
			EndTime:       common.GenTimestamp(),
		},
	}
	data := protos.ContractData{
		Data:  &contract,
		Token: "ZDNkM0xtWjFkSFZ5WlhabGNpNWpiMjA9",
	}

	requestBody, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	requestHead := make(map[string]string)
	requestHead["Content-Type"] = APPLICATION_X_PROTOBUF
	_, err = httpRequest("POST", url, requestBody, requestHead)
	if err != nil {
		// handle error
		fmt.Println("error ", err.Error())
	}
	//fmt.Println("response is:", string(response))
}

func Test_Query(t *testing.T) {
	url := default_url + "query"

	contract := protos.Contract{ // proto-buf
		Id: "d881e3f1a9b79c1d3ef7f7c5cf03a1965362367bf6823a6550b6c417211e1889",
	}

	requestBody, err := proto.Marshal(&contract)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	requestHead := make(map[string]string)
	requestHead["Content-Type"] = APPLICATION_X_PROTOBUF
	responseData, err := httpRequest("POST", url, requestBody, requestHead)
	if err != nil {
		// handle error
		fmt.Println("error ", err.Error())
	}
	/*---------------------- response -----------------------*/
	var responseResult protos.Contract
	proto.Unmarshal(responseData, &responseResult)
	beego.Error("result is ",responseResult)


	var responseMap map[string]interface{}
	json.Unmarshal(responseData, &responseMap)
	fmt.Println("response is:", responseMap)
	responseDataBody := responseMap["data"]
	fmt.Println("responseDataBody is:", responseDataBody)
	ress, ok := responseDataBody.(string)

	if ok {
	}
	fmt.Println("ress is:", ress)
	rrrr, _ := json.Marshal(ress)
	fmt.Println("ress byte is:", rrrr)

	responseDataBodyByte, _ := json.Marshal(responseDataBody)
	fmt.Println(responseDataBodyByte)
	var contractData protos.ContractData
	proto.Unmarshal(responseDataBodyByte, &contractData)
	fmt.Println("contractData is:", contractData)

	//conotractData := responseData["data"]
	//contractData_str := common.Serialize(contractData.Data)
}

func Test_Track(t *testing.T) {
	url := default_url + "track"

	contract := protos.Contract{ // proto-buf
		Id: "2",
	}
	data := protos.ContractData{
		Data:  &contract,
		Token: "ZDNkM0xtWjFkSFZ5WlhabGNpNWpiMjA9",
	}
	requestBody, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	requestHead := make(map[string]string)
	requestHead["Content-Type"] = APPLICATION_X_PROTOBUF
	_, err = httpRequest("POST", url, requestBody, requestHead)
	if err != nil {
		// handle error
		fmt.Println("error ", err.Error())
	}
	//fmt.Println("response is:", string(response))
}

func Test_Update(t *testing.T) {
	url := default_url + "update"

	contract := protos.Contract{ // proto-buf
		ContractHead: &protos.ContractHead{
			MainPubkey: "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			Version:    1,
		},
		ContractBody: &protos.ContractBody{
			Caption: "CREATE",
			Cname:   "futurever",
			ContractAssets: []*protos.ContractAsset{
				{
					AssetId:     "001",
					Name:        "futurever-1",
					Amount:      1000,
					Caption:     "futurever",
					Description: "",
					Unit:        "int32",
					MetaData:    nil,
				},
				{
					AssetId:     "003",
					Name:        "futurever-3",
					Amount:      452,
					Caption:     "futurever",
					Description: "",
					Unit:        "int32",
					MetaData:    nil,
				},
				{
					AssetId:     "002",
					Name:        "futurever-2",
					Amount:      99999,
					Caption:     "futurever",
					Description: "",
					Unit:        "int32",
					MetaData:    nil,
				},
			},
			ContractComponents: nil,
			ContractId:         common.GenerateUUID(),
			ContractOwners: []string{
				"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
				"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			},
			ContractSignatures: []*protos.ContractSignature{
				{
					OwnerPubkey:   "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
					Signature:     "65D27HW4uXYvkekGssAQB93D92onMyU1NVnCJnE1PgRKz2uFSPZ6aQvid4qZvkxys7G4r2Mf2KFn5BSQyEBhWs34",
					SignTimestamp: common.GenTimestamp(),
				},
				{
					OwnerPubkey:   "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
					Signature:     "5i5dTtQseQjWZ8UdchqQtgttyeeFmB3LDFYzNKafvV2YvTqwv4wZ9mFsH7qgysV9ow893D1h2Xnt1uCXLHtbKrkT",
					SignTimestamp: common.GenTimestamp(),
				},
			},
			ContractState: "",
			Creator:       "futurever",
			CreatorTime:   common.GenTimestamp(),
			Ctype:         "CONTRACT",
			Description:   "CREATE CONTRACT BY futurever [合约创建]",
			StartTime:     common.GenTimestamp(),
			EndTime:       common.GenTimestamp(),
		},
	}
	data := protos.ContractData{
		Data:  &contract,
		Token: "ZDNkM0xtWjFkSFZ5WlhabGNpNWpiMjA9",
	}
	requestBody, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	requestHead := make(map[string]string)
	requestHead["Content-Type"] = APPLICATION_X_PROTOBUF
	_, err = httpRequest("POST", url, requestBody, requestHead)
	if err != nil {
		// handle error
		fmt.Println("error ", err.Error())
	}
	//fmt.Println("response is:", string(response))
}

func Test_Test(t *testing.T) {
	url := default_url + "test"

	contract := protos.Contract{ // proto-buf
		ContractHead: &protos.ContractHead{
			MainPubkey: "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			Version:    1,
		},
		ContractBody: &protos.ContractBody{
			Caption: "CREATE",
			Cname:   "futurever",
			ContractAssets: []*protos.ContractAsset{
				{
					AssetId:     "001",
					Name:        "futurever-1",
					Amount:      1000,
					Caption:     "futurever",
					Description: "",
					Unit:        "int32",
					MetaData:    nil,
				},
				{
					AssetId:     "003",
					Name:        "futurever-3",
					Amount:      452,
					Caption:     "futurever",
					Description: "",
					Unit:        "int32",
					MetaData:    nil,
				},
				{
					AssetId:     "002",
					Name:        "futurever-2",
					Amount:      99999,
					Caption:     "futurever",
					Description: "",
					Unit:        "int32",
					MetaData:    nil,
				},
			},
			ContractComponents: nil,
			ContractId:         common.GenerateUUID(),
			ContractOwners: []string{
				"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
				"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			},
			ContractSignatures: []*protos.ContractSignature{
				{
					OwnerPubkey:   "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
					Signature:     "65D27HW4uXYvkekGssAQB93D92onMyU1NVnCJnE1PgRKz2uFSPZ6aQvid4qZvkxys7G4r2Mf2KFn5BSQyEBhWs34",
					SignTimestamp: common.GenTimestamp(),
				},
				{
					OwnerPubkey:   "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
					Signature:     "5i5dTtQseQjWZ8UdchqQtgttyeeFmB3LDFYzNKafvV2YvTqwv4wZ9mFsH7qgysV9ow893D1h2Xnt1uCXLHtbKrkT",
					SignTimestamp: common.GenTimestamp(),
				},
			},
			ContractState: "",
			Creator:       "futurever",
			CreatorTime:   common.GenTimestamp(),
			Ctype:         "CONTRACT",
			Description:   "CREATE CONTRACT BY futurever [合约创建]",
			StartTime:     common.GenTimestamp(),
			EndTime:       common.GenTimestamp(),
		},
	}
	data := protos.ContractData{
		Data:  &contract,
		Token: "ZDNkM0xtWjFkSFZ5WlhabGNpNWpiMjA9",
	}
	requestBody, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	requestHead := make(map[string]string)
	requestHead["Content-Type"] = APPLICATION_X_PROTOBUF
	_, err = httpRequest("POST", url, requestBody, requestHead)
	if err != nil {
		// handle error
		fmt.Println("error ", err.Error())
	}
	//fmt.Println("response is:", string(response))
}
