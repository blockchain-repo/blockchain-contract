package controllers

import (
	"bytes"
	//"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"testing"
	"unicontract/src/common"
	"unicontract/src/core/model"
	"unicontract/src/core/protos"
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

func test_CreateContract(t *testing.T) {

}

//var default_url = "http://192.168.1.14:8088/v1/contract/"
var default_url = "http://localhost:8088/v1/contract/"

func Test_AuthSignature(t *testing.T) {
	url := default_url + "authSignature"
	contractModel := model.ContractModel{}
	//contractAsset := []*protos.ContractAsset{}
	//contractComponent:=[]*protos.ContractComponent{}

	contractHead := &protos.ContractHead{"", 1}

	contractOwners := []string{
		"BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
		//"4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		//"9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
	}
	contractBody := &protos.ContractBody{
		ContractId:    "UUID-1234-5678-90",
		Cname:         "test contract output",
		Ctype:         "CREATE",
		Caption:       "购智能手机返话费合约产品协议",
		Description:   "移动用户A花费500元购买移动运营商B的提供的合约智能手机C后",
		ContractState: "",
		Creator:       common.GenTimestamp(),
		CreatorTime:   common.GenTimestamp(),
		EndTime:       common.GenTimestamp(),
		StartTime:     common.GenTimestamp(),

		ContractOwners:     contractOwners,
		ContractSignatures: nil,
		ContractAssets:     nil,
		ContractComponents: nil,
	}

	contractModel.ContractHead = contractHead
	contractModel.ContractBody = contractBody

	contractSignatures := []*protos.ContractSignature{
		{
			OwnerPubkey:   "BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
			Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
			SignTimestamp: common.GenTimestamp(),
		},
		//{
		//	OwnerPubkey:   "BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
		//	Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
		//{
		//	OwnerPubkey:   "4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		//	Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
		//{Create
		//	OwnerPubkey:   "9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
		//	Signature:     contractModel.Sign("9647UfPdDSwBf5kw7tUrSe7cmYY5RvVX47GrGqSh4XVi"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
	}
	contractBody.ContractSignatures = contractSignatures
	contractModel.Id = contractModel.GenerateId()
	serializeContractModel := common.StructSerialize(contractModel)
	fmt.Println("produce the contractModel", serializeContractModel)

	protoContract, _ := fromContractModelStrToContract(serializeContractModel)
	requestBody, err := proto.Marshal(&protoContract)
	if err != nil {
		fmt.Println("proto.Marshal", err)
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

func Test_CreatValidContract(t *testing.T) {
	url := default_url + "create"
	contractModel := model.ContractModel{}
	//contractAsset := []*protos.ContractAsset{}
	//contractComponent:=[]*protos.ContractComponent{}

	contractHead := &protos.ContractHead{"", 1}

	contractOwners := []string{
		"BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
		//"4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		//"9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
	}
	contractBody := &protos.ContractBody{
		ContractId:    "UUID-1234-5678-90",
		Cname:         "test contract output",
		Ctype:         "CREATE",
		Caption:       "购智能手机返话费合约产品协议",
		Description:   "移动用户A花费500元购买移动运营商B的提供的合约智能手机C后",
		ContractState: "",
		Creator:       common.GenTimestamp(),
		CreatorTime:   common.GenTimestamp(),
		EndTime:       common.GenTimestamp(),
		StartTime:     common.GenTimestamp(),

		ContractOwners:     contractOwners,
		ContractSignatures: nil,
		ContractAssets:     nil,
		ContractComponents: nil,
	}

	contractModel.ContractHead = contractHead
	contractModel.ContractBody = contractBody

	contractSignatures := []*protos.ContractSignature{
		{
			OwnerPubkey:   "BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
			Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
			SignTimestamp: common.GenTimestamp(),
		},
		//{
		//	OwnerPubkey:   "BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
		//	Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
		//{
		//	OwnerPubkey:   "4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		//	Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
		//{Create
		//	OwnerPubkey:   "9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
		//	Signature:     contractModel.Sign("9647UfPdDSwBf5kw7tUrSe7cmYY5RvVX47GrGqSh4XVi"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
	}
	contractBody.ContractSignatures = contractSignatures
	contractModel.Id = contractModel.GenerateId()
	serializeContractModel := common.StructSerialize(contractModel)
	fmt.Println("produce the contractModel", serializeContractModel)

	protoContract, _ := fromContractModelStrToContract(serializeContractModel)
	requestBody, err := proto.Marshal(&protoContract)
	if err != nil {
		fmt.Println("proto.Marshal", err)
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
	contractModel := model.ContractModel{}

	contractHead := &protos.ContractHead{"", 1}

	contractOwners := []string{
		"BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
		//"4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		//"9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
	}
	contractBody := &protos.ContractBody{
		ContractId:    "UUID-1234-5678-90",
		Cname:         "test contract output",
		Ctype:         "CREATE",
		Caption:       "购智能手机返话费合约产品协议",
		Description:   "移动用户A花费500元购买移动运营商B的提供的合约智能手机C后",
		ContractState: "",
		Creator:       common.GenTimestamp(),
		CreatorTime:   common.GenTimestamp(),
		EndTime:       common.GenTimestamp(),
		StartTime:     common.GenTimestamp(),

		ContractOwners:     contractOwners,
		ContractSignatures: nil,
		ContractAssets:     nil,
		ContractComponents: nil,
		// 下面的用法是错误的,proto处理后,再次转换回来[] 会变成 null, 而不是json的 [], hash会出错
		//ContractAssets:     []*protos.ContractAsset{},
		//ContractComponents: []*protos.ContractComponent{},
	}

	contractModel.ContractHead = contractHead
	contractModel.ContractBody = contractBody

	contractSignatures := []*protos.ContractSignature{
		{
			OwnerPubkey:   "BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
			Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
			SignTimestamp: common.GenTimestamp(),
		},
		//{
		//	OwnerPubkey:   "BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
		//	Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
		//{
		//	OwnerPubkey:   "4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		//	Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
		//{Create
		//	OwnerPubkey:   "9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
		//	Signature:     contractModel.Sign("9647UfPdDSwBf5kw7tUrSe7cmYY5RvVX47GrGqSh4XVi"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
	}
	contractBody.ContractSignatures = contractSignatures
	contractModel.Id = contractModel.GenerateId()
	serializeContractModel := common.StructSerialize(contractModel)
	fmt.Println("produce the contractModel", serializeContractModel)

	protoContract, _ := fromContractModelStrToContract(serializeContractModel)
	requestBody, err := proto.Marshal(&protoContract)
	if err != nil {
		fmt.Println("proto.Marshal", err)
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

	contractModel := model.ContractModel{}
	//contractAsset := []*protos.ContractAsset{}
	//contractComponent:=[]*protos.ContractComponent{}

	contractHead := &protos.ContractHead{"", 1}

	contractOwners := []string{
		"BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
		//"4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		//"9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
	}
	contractBody := &protos.ContractBody{
		ContractId:    "UUID-1234-5678-90",
		Cname:         "test contract output",
		Ctype:         "CREATE",
		Caption:       "购智能手机返话费合约产品协议",
		Description:   "移动用户A花费500元购买移动运营商B的提供的合约智能手机C后",
		ContractState: "",
		Creator:       common.GenTimestamp(),
		CreatorTime:   common.GenTimestamp(),
		EndTime:       common.GenTimestamp(),
		StartTime:     common.GenTimestamp(),

		ContractOwners:     contractOwners,
		ContractSignatures: nil,
		ContractAssets:     nil,
		ContractComponents: nil,
	}

	contractModel.ContractHead = contractHead
	contractModel.ContractBody = contractBody

	contractSignatures := []*protos.ContractSignature{
		{
			OwnerPubkey:   "BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
			Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
			SignTimestamp: common.GenTimestamp(),
		},
		//{
		//	OwnerPubkey:   "BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
		//	Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
		//{
		//	OwnerPubkey:   "4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		//	Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
		//{Create
		//	OwnerPubkey:   "9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
		//	Signature:     contractModel.Sign("9647UfPdDSwBf5kw7tUrSe7cmYY5RvVX47GrGqSh4XVi"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
	}
	contractBody.ContractSignatures = contractSignatures
	contractModel.Id = contractModel.GenerateId()
	serializeContractModel := common.StructSerialize(contractModel)
	fmt.Println("produce the contractModel", serializeContractModel)

	protoContract, _ := fromContractModelStrToContract(serializeContractModel)
	requestBody, err := proto.Marshal(&protoContract)
	if err != nil {
		fmt.Println("proto.Marshal", err)
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

	contractModel := model.ContractModel{}
	//contractAsset := []*protos.ContractAsset{}
	//contractComponent:=[]*protos.ContractComponent{}

	contractHead := &protos.ContractHead{"", 1}

	contractOwners := []string{
		"BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
		//"4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		//"9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
	}
	contractBody := &protos.ContractBody{
		ContractId:    "UUID-1234-5678-90",
		Cname:         "test contract output",
		Ctype:         "CREATE",
		Caption:       "购智能手机返话费合约产品协议",
		Description:   "移动用户A花费500元购买移动运营商B的提供的合约智能手机C后",
		ContractState: "",
		Creator:       common.GenTimestamp(),
		CreatorTime:   common.GenTimestamp(),
		EndTime:       common.GenTimestamp(),
		StartTime:     common.GenTimestamp(),

		ContractOwners:     contractOwners,
		ContractSignatures: nil,
		ContractAssets:     nil,
		ContractComponents: nil,
	}

	contractModel.ContractHead = contractHead
	contractModel.ContractBody = contractBody

	contractSignatures := []*protos.ContractSignature{
		{
			OwnerPubkey:   "BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
			Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
			SignTimestamp: common.GenTimestamp(),
		},
		//{
		//	OwnerPubkey:   "BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
		//	Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
		//{
		//	OwnerPubkey:   "4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		//	Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
		//{Create
		//	OwnerPubkey:   "9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
		//	Signature:     contractModel.Sign("9647UfPdDSwBf5kw7tUrSe7cmYY5RvVX47GrGqSh4XVi"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
	}
	contractBody.ContractSignatures = contractSignatures
	contractModel.Id = contractModel.GenerateId()
	serializeContractModel := common.StructSerialize(contractModel)
	fmt.Println("produce the contractModel", serializeContractModel)

	protoContract, _ := fromContractModelStrToContract(serializeContractModel)
	requestBody, err := proto.Marshal(&protoContract)
	if err != nil {
		fmt.Println("proto.Marshal", err)
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
	contractModel := model.ContractModel{}
	//contractAsset := []*protos.ContractAsset{}
	//contractComponent:=[]*protos.ContractComponent{}

	contractHead := &protos.ContractHead{"", 1}

	contractOwners := []string{
		"BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
		//"4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		//"9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
	}
	contractBody := &protos.ContractBody{
		ContractId:    "UUID-1234-5678-90",
		Cname:         "test contract output",
		Ctype:         "CREATE",
		Caption:       "购智能手机返话费合约产品协议",
		Description:   "移动用户A花费500元购买移动运营商B的提供的合约智能手机C后",
		ContractState: "",
		Creator:       common.GenTimestamp(),
		CreatorTime:   common.GenTimestamp(),
		EndTime:       common.GenTimestamp(),
		StartTime:     common.GenTimestamp(),

		ContractOwners:     contractOwners,
		ContractSignatures: nil,
		ContractAssets:     nil,
		ContractComponents: nil,
	}

	contractModel.ContractHead = contractHead
	contractModel.ContractBody = contractBody

	contractSignatures := []*protos.ContractSignature{
		{
			OwnerPubkey:   "BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
			Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
			SignTimestamp: common.GenTimestamp(),
		},
		//{
		//	OwnerPubkey:   "BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
		//	Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
		//{
		//	OwnerPubkey:   "4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		//	Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
		//{Create
		//	OwnerPubkey:   "9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
		//	Signature:     contractModel.Sign("9647UfPdDSwBf5kw7tUrSe7cmYY5RvVX47GrGqSh4XVi"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
	}
	contractBody.ContractSignatures = contractSignatures
	contractModel.Id = contractModel.GenerateId()
	serializeContractModel := common.StructSerialize(contractModel)
	fmt.Println("produce the contractModel", serializeContractModel)

	protoContract, _ := fromContractModelStrToContract(serializeContractModel)
	requestBody, err := proto.Marshal(&protoContract)
	if err != nil {
		fmt.Println("proto.Marshal", err)
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

	contractModel := model.ContractModel{}
	//contractAsset := []*protos.ContractAsset{}
	//contractComponent:=[]*protos.ContractComponent{}

	contractHead := &protos.ContractHead{"", 1}

	contractOwners := []string{
		"BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
		//"4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		//"9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
	}
	contractBody := &protos.ContractBody{
		ContractId:    "UUID-1234-5678-90",
		Cname:         "test contract output",
		Ctype:         "CREATE",
		Caption:       "购智能手机返话费合约产品协议",
		Description:   "移动用户A花费500元购买移动运营商B的提供的合约智能手机C后",
		ContractState: "",
		Creator:       common.GenTimestamp(),
		CreatorTime:   common.GenTimestamp(),
		EndTime:       common.GenTimestamp(),
		StartTime:     common.GenTimestamp(),

		ContractOwners:     contractOwners,
		ContractSignatures: nil,
		ContractAssets:     nil,
		ContractComponents: nil,
	}

	contractModel.ContractHead = contractHead
	contractModel.ContractBody = contractBody

	contractSignatures := []*protos.ContractSignature{
		{
			OwnerPubkey:   "BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
			Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
			SignTimestamp: common.GenTimestamp(),
		},
		//{
		//	OwnerPubkey:   "BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
		//	Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
		//{
		//	OwnerPubkey:   "4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		//	Signature:     contractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
		//{Create
		//	OwnerPubkey:   "9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
		//	Signature:     contractModel.Sign("9647UfPdDSwBf5kw7tUrSe7cmYY5RvVX47GrGqSh4XVi"),
		//	SignTimestamp: common.GenTimestamp(),
		//},
	}
	contractBody.ContractSignatures = contractSignatures
	contractModel.Id = contractModel.GenerateId()
	serializeContractModel := common.StructSerialize(contractModel)
	fmt.Println("produce the contractModel", serializeContractModel)

	protoContract, _ := fromContractModelStrToContract(serializeContractModel)
	requestBody, err := proto.Marshal(&protoContract)
	if err != nil {
		fmt.Println("proto.Marshal", err)
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
