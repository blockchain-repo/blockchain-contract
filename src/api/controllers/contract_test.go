package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/model"
	"unicontract/src/core/protos"

	"github.com/golang/protobuf/proto"
	"unicontract/src/core/engine/execengine/constdef"
)

// application content-type
const (
	APPLICATION_X_PROTOBUF   = "application/x-protobuf"
	APPLICATION_JSON         = "application/json"
	APPLICATION_OCTET_STREAM = "application/octet-stream"
)

func init() {
	config.Init()
}

func httpRequest(method string, urlStr string, body []byte, requestHead map[string]string) ([]byte, error) {
	client := &http.Client{}
	req_body := bytes.NewReader(body)

	if method == "" {
		method = "POST"
	}
	req, err := http.NewRequest(method, urlStr, req_body)
	if err != nil {
		fmt.Println("request error", err)
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
	fmt.Printf("Request %s [%s] content-type=%s\n", urlStr, method, contentType)

	if err == nil {
		responseBody, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("Response code: %v\n", resp.StatusCode)
		defer resp.Body.Close()
		return responseBody, err
	}
	return nil, err
}

func generatContractModel(produceValid bool, optArgs ...map[string]interface{}) (string, error) {
	contractOwnersLen := 3
	if tempLen, ok := optArgs[0]["contractOwnersLen"]; ok {
		contractOwnersLen, ok = tempLen.(int)
		if !ok {
			fmt.Println("optArgs type error for param contractOwnersLen")
			return "", nil
		}
	}
	// 生成的合约签名人个数
	contractSignaturesLen := contractOwnersLen
	if tempLen, ok := optArgs[0]["contractSignaturesLen"]; ok {
		contractSignaturesLen, ok = tempLen.(int)
		if !ok {
			fmt.Println("optArgs type error for param contractSignaturesLen")
			return "", nil
		}
	}

	if contractSignaturesLen >= contractOwnersLen || contractSignaturesLen <= 0 {
		contractSignaturesLen = contractOwnersLen
	}

	//generate contractOwnersLen keypair
	owners := make(map[string]string)
	ownersPubkeys := make([]string, contractOwnersLen)
	for i := 0; i < contractOwnersLen; i++ {
		publicKeyBase58, privateKeyBase58 := common.GenerateKeyPair()
		owners[publicKeyBase58] = privateKeyBase58
		ownersPubkeys[i] = publicKeyBase58
	}
	owners["3FyHdZVX4adfSSTg7rZDPMzqzM8k5fkpu43vbRLvEXLJ"] = "5Pv7F7g9BvNDEMdb8HV5aLHpNTNkxVpNqnLTQ58Z5heC"
	ownersPubkeys[0] = "3FyHdZVX4adfSSTg7rZDPMzqzM8k5fkpu43vbRLvEXLJ"
	owners["3JYvp8FrDvrhMVjGURoZvD5oVDnug3HZcotYG4AdvvSP"] = "2GY6Rif9NMSdQ7JPBSGqFPKggBi1Kaa2ru11wrCdiv6V"
	ownersPubkeys[1] = "3JYvp8FrDvrhMVjGURoZvD5oVDnug3HZcotYG4AdvvSP"
	//owners["Vn3aq7YcTWFjRXEUFKpzZvseeQuYMzAM6NZkGFzeBUY"] = "5cg11pBp193AkbEJ5CL36z5eDGgMtsaFeyVSaJwG7rSy"
	//ownersPubkeys[2] = "Vn3aq7YcTWFjRXEUFKpzZvseeQuYMzAM6NZkGFzeBUY"

	/*-------------------- generate contractModel ------------------*/
	contractModel := model.ContractModel{}

	//模拟用户发送的数据, mainpubkey 传入API 后,根据配置生成,此处请勿设置
	assignTime, err := common.GenSpecialTimestamp("2017-06-01 00:00:00")
	operateTime, err := common.GenSpecialTimestamp("2017-06-01 00:00:00")
	mainPubkey := config.Config.Keypair.PublicKey
	contractHead := &protos.ContractHead{mainPubkey, 1,
		assignTime, operateTime, 0}
	// random choose the creator
	randomCreator := ownersPubkeys[common.RandInt(0, contractOwnersLen)]
	//contractAsset := []*protos.ContractAsset{}
	//contractComponent:=[]*protos.ContractComponent{}

	startTime, err := common.GenSpecialTimestamp("2017-05-29 00:00:00")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	endTime, err := common.GenSpecialTimestamp("2017-10-06 07:00:00")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// random contractOwners 随机生成的合约拥有者数组
	contractOwners := ownersPubkeys
	createTime := common.GenTimestamp()
	contractBody := &protos.ContractBody{
		ContractId:  "UUID-1234-5678-90 wx1",
		Cname:       "test create contract",
		Ctype:       "CREATE",
		Caption:     "futurever",
		Description: "www.futurever.com",
		//ContractState: constdef.ContractState[constdef.Component_Unknown],
		//ContractState: constdef.ContractState[constdef.Contract_Create],
		ContractState: constdef.ContractState[constdef.Contract_Signature],
		//ContractState: constdef.ContractState[constdef.Contract_In_Process],
		//ContractState: constdef.ContractState[constdef.Contract_Completed],
		//ContractState:      constdef.ContractState[constdef.Contract_Disgarded],
		Creator:            randomCreator,
		CreateTime:         createTime,
		StartTime:          startTime,
		EndTime:            endTime,
		ContractOwners:     contractOwners,
		ContractSignatures: nil,
		ContractAssets: []*protos.ContractAsset{
			{
				AssetId:     "xxxxxxxxxxx",
				Name:        "asset_money",
				Caption:     "智能手机",
				Description: "智能手机，返还话费产品",
				Unit:        "台",
				Amount:      1,
				MetaData: map[string]string{
					"TestAsset1": "1111111111",
					"TestAsset2": "2222222222",
				},
			},
		},
		ContractComponents: []*protos.ContractComponent{
			{
				DataList: []*protos.ComponentData{
					{
						Cname:             "232",
						DefaultValueFloat: float64(33),
						DataRangeFloat:    []float64{32, 32},

						//DefaultValue:(*google_protobuf.Any{TypeUrl:"22", Value:2}),
						//Value:3,
					},
				},
			},
		},

		NextTasks: nil,
	}

	contractModel.ContractHead = contractHead
	contractModel.ContractBody = contractBody

	// 生成 签名
	contractSignatures := make([]*protos.ContractSignature, contractSignaturesLen)
	for i := 0; i < contractSignaturesLen; i++ {
		ownerPubkey := ownersPubkeys[i]
		privateKey := owners[ownerPubkey]
		contractSignatures[i] = &protos.ContractSignature{
			OwnerPubkey:   ownerPubkey,
			Signature:     contractModel.Sign(privateKey),
			SignTimestamp: common.GenTimestamp(),
		}
		//contractSignatures[i] = &protos.ContractSignature{}
		//contractSignatures[i].OwnerPubkey = ownerPubkey
		//contractSignatures[i].Signature = contractModel.Sign(privateKey)
		//contractSignatures[i].SignTimestamp = common.GenTimestamp()
	}

	contractBody.ContractSignatures = contractSignatures
	contractModel.Id = contractModel.GenerateId()
	if !produceValid {
		contractModel.ContractBody.Description = "generate error contract for test"
	}
	serializeContractModel := common.StructSerialize(contractModel)
	fmt.Println("produce the contractModel", serializeContractModel)
	fmt.Println("produce the contractModel", common.SerializePretty(contractModel))

	return serializeContractModel, nil

}

func generateProtoContract(produceValid bool, optArgs ...map[string]interface{}) ([]byte, error) {
	contractOwnersLen := 2
	if tempLen, ok := optArgs[0]["contractOwnersLen"]; ok {
		contractOwnersLen, ok = tempLen.(int)
		if !ok {
			fmt.Println("generateProtoContract optArgs error for  contractOwnersLen")
			optArgs[0]["contractOwnersLen"] = contractOwnersLen
		}
	}
	// 生成的合约签名人个数
	contractSignaturesLen := contractOwnersLen
	if tempLen, ok := optArgs[0]["contractSignaturesLen"]; ok {
		contractSignaturesLen, ok = tempLen.(int)
		if !ok {
			fmt.Println("optArgs type error for param contractSignaturesLen")
			optArgs[0]["contractSignaturesLen"] = contractSignaturesLen
		}
	}
	serializeContractModel, err := generatContractModel(produceValid, optArgs[0])
	if err != nil {
		return nil, err
	}
	protoContract, _ := fromContractModelStrToContract(serializeContractModel)
	requestBody, err := proto.Marshal(&protoContract)
	if err != nil {
		fmt.Println(requestBody, err)
		return nil, err
	}
	return requestBody, nil
}

//var default_url = "http://36.110.71.170:66/v1/contract/"

//var default_url = "http://192.168.1.14:8088/v1/contract/"
//var default_url = "http://36.110.71.170:66/v1/contract/"
//var default_url = "http://localhost:8088/v1/contract/"

//var default_url = "http://192.168.1.11:8088/v1/contract/"

//var default_url = "http://192.168.1.14:8088/v1/contract/"

var default_url = "http://localhost:8088/v1/contract/"

//var default_url = "http://www.wxcsdb88.com:8088/v1/contract/"

func Test_AuthSignature(t *testing.T) {
	url := default_url + "authSignature"
	produceValid := true
	extraAttr := make(map[string]interface{})
	extraAttr["contractOwnersLen"] = 2
	extraAttr["contractSignaturesLen"] = 2

	requestBody, err := generateProtoContract(produceValid, extraAttr)
	//requestBody, err := generateProtoContract(false, extraAttr)
	if err != nil {
		fmt.Println("generateProtoContract error ", err.Error())
		return
	}
	requestHead := make(map[string]string)
	requestHead["Content-Type"] = APPLICATION_X_PROTOBUF
	response, err := httpRequest("POST", url, requestBody, requestHead)
	if err != nil {
		fmt.Println("httpRequest error ", err.Error())
		return
	}
	//接受返回数据
	var responseData protos.ResponseData
	proto.Unmarshal(response, &responseData)
	fmt.Println(common.StructSerializePretty(responseData))
}

func Test_CreatContract(t *testing.T) {
	url := default_url + "create"
	produceValid := true
	extraAttr := make(map[string]interface{})
	extraAttr["contractOwnersLen"] = 2
	extraAttr["contractSignaturesLen"] = 2

	requestBody, err := generateProtoContract(produceValid, extraAttr)
	//requestBody, err := generateProtoContract(false, extraAttr)
	if err != nil {
		fmt.Println("generateProtoContract error ", err.Error())
		return
	}
	requestHead := make(map[string]string)
	requestHead["Content-Type"] = APPLICATION_X_PROTOBUF
	response, err := httpRequest("POST", url, requestBody, requestHead)
	//接受返回数据
	var responseData protos.ResponseData
	proto.Unmarshal(response, &responseData)
	fmt.Println(common.StructSerializePretty(responseData))
}

func Test_Signature(t *testing.T) {
	url := default_url + "signature"
	produceValid := true
	extraAttr := make(map[string]interface{})
	extraAttr["contractOwnersLen"] = 2
	extraAttr["contractSignaturesLen"] = 2

	requestBody, err := generateProtoContract(produceValid, extraAttr)
	//requestBody, err := generateProtoContract(false, extraAttr)
	if err != nil {
		fmt.Println("generateProtoContract error ", err.Error())
		return
	}

	requestHead := make(map[string]string)
	requestHead["Content-Type"] = APPLICATION_X_PROTOBUF
	response, err := httpRequest("POST", url, requestBody, requestHead)
	//接受返回数据
	var responseData protos.ResponseData
	proto.Unmarshal(response, &responseData)
	fmt.Println(common.StructSerializePretty(responseData))
}

func Test_Terminate(t *testing.T) {
	url := default_url + "terminate"
	produceValid := true
	extraAttr := make(map[string]interface{})
	extraAttr["contractOwnersLen"] = 2
	extraAttr["contractSignaturesLen"] = 2

	requestBody, err := generateProtoContract(produceValid, extraAttr)
	//requestBody, err := generateProtoContract(false, extraAttr)
	if err != nil {
		fmt.Println("generateProtoContract error ", err.Error())
		return
	}

	requestHead := make(map[string]string)
	requestHead["Content-Type"] = APPLICATION_X_PROTOBUF
	response, err := httpRequest("POST", url, requestBody, requestHead)
	//接受返回数据
	var responseData protos.ResponseData
	proto.Unmarshal(response, &responseData)
	fmt.Println(common.StructSerializePretty(responseData))
}

func Test_Query(t *testing.T) {
	url := default_url + "query"

	contract := protos.Contract{ // proto-buf
		Id: "64520eba60bde72f71b4646d6cc0872715e4717234ca6031c621d247e5c4553c",
	}

	requestBody, err := proto.Marshal(&contract)
	if err != nil {
		fmt.Println("proto.Marshal error ", err.Error())
		return
	}
	requestHead := make(map[string]string)
	requestHead["Content-Type"] = APPLICATION_X_PROTOBUF
	response, err := httpRequest("POST", url, requestBody, requestHead)
	if err != nil {
		fmt.Println("httpRequest error ", err.Error())
		return
	}

	/*---------------------- response 接受的响应数据-----------------------*/
	var responseData protos.ResponseData
	err = proto.Unmarshal(response, &responseData)
	if err != nil {
		fmt.Println("proto.Unmarshal protos.ResponseData error")
		return
	}
	fmt.Println("responseData content is: \n", common.StructSerializePretty(responseData))

	ok := responseData.Ok
	_ = ok
	msg := responseData.Msg
	_ = msg
	data := responseData.Data

	/*----------------- contract Unmarshal 数据库真实数据----------------------*/
	// 返回的数据是 字节数组->字符串 ,所以需要 字符串->字节数组
	// API response []byte -> string, resolve string -> []byte
	var contractProtoBytes = []byte(data)
	//fmt.Println(contractProtoBytes)

	var contractQueryData protos.Contract
	err = proto.Unmarshal(contractProtoBytes, &contractQueryData)
	if err != nil {
		fmt.Println("proto.Unmarshal protos.Contract error")
		return
	}
	//fmt.Println("query contract is:\n", contractQueryData)
	fmt.Println("Contract content is: \n", common.StructSerializePretty(contractQueryData))

}

func Test_Track(t *testing.T) {
	url := default_url + "track"

	contract := protos.Contract{ // proto-buf
		Id: "64520eba60bde72f71b4646d6cc0872715e4717234ca6031c621d247e5c4553c",
	}

	requestBody, err := proto.Marshal(&contract)
	if err != nil {
		fmt.Println("proto.Marshal error ", err.Error())
		return
	}
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	requestHead := make(map[string]string)
	requestHead["Content-Type"] = APPLICATION_X_PROTOBUF
	response, err := httpRequest("POST", url, requestBody, requestHead)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	/*---------------------- response 接受的响应数据-----------------------*/
	var responseData protos.ResponseData
	err = proto.Unmarshal(response, &responseData)
	if err != nil {
		fmt.Println("proto.Unmarshal protos.ResponseData error")
		return
	}
	fmt.Println("responseData content is: \n", common.StructSerializePretty(responseData))

	ok := responseData.Ok
	_ = ok
	msg := responseData.Msg
	_ = msg
	data := responseData.Data
	fmt.Println(data)
}

func Test_Update(t *testing.T) {
	url := default_url + "update"
	produceValid := true
	extraAttr := make(map[string]interface{})
	extraAttr["contractOwnersLen"] = 2
	extraAttr["contractSignaturesLen"] = 2

	requestBody, err := generateProtoContract(produceValid, extraAttr)
	//requestBody, err := generateProtoContract(false, extraAttr)
	if err != nil {
		fmt.Println("generateProtoContract error ", err.Error())
		return
	}

	requestHead := make(map[string]string)
	requestHead["Content-Type"] = APPLICATION_X_PROTOBUF
	response, err := httpRequest("POST", url, requestBody, requestHead)
	if err != nil {
		fmt.Println("httpRequest error ", err.Error())
		return
	}
	//接受返回数据
	var responseData protos.ResponseData
	proto.Unmarshal(response, &responseData)
	fmt.Println(common.StructSerializePretty(responseData))
}

func Test_Test(t *testing.T) {
	url := default_url + "test"
	produceValid := false
	extraAttr := make(map[string]interface{})
	extraAttr["contractOwnersLen"] = 2
	extraAttr["contractSignaturesLen"] = 2

	requestBody, err := generateProtoContract(produceValid, extraAttr)
	if err != nil {
		fmt.Println("generateProtoContract error ", err.Error())
		return
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
