package model

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"testing"
	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/protos"
)

// API receive and transfer it to contractModel
func fromContractToContractModel(contract protos.Contract) ContractModel {
	var contractModel ContractModel
	contractModel.Contract = contract
	return contractModel
}

// go rethink get contractModel string and transfer it to contract
func fromContractModelStrToContract(contractModelStr string) (protos.Contract, error) {
	// 1. to contractModel
	var contractModel ContractModel
	err := json.Unmarshal([]byte(contractModelStr), &contractModel)
	// 2. to contract
	contract := contractModel.Contract
	if err != nil {
		logs.Error("error fromContractModelStrToContract", err)
		return contract, err
	}

	return contract, nil
}

func generatContractModel(produceValid bool, optArgs ...map[string]interface{}) (string, error) {
	contractOwnersLen := 1
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

	/*-------------------- generate contractModel ------------------*/
	contractModel := ContractModel{}

	//模拟用户发送的数据, mainpubkey 传入API 后,根据配置生成,此处请勿设置
	mainPubkey := config.Config.Keypair.PublicKey
	contractHead := &protos.ContractHead{mainPubkey, 1, common.GenTimestamp(), common.GenTimestamp(), 1}

	// random choose the creator
	randomCreator := ownersPubkeys[common.RandInt(0, contractOwnersLen)]
	//contractAsset := []*protos.ContractAsset{}
	//contractComponent:=[]*protos.ContractComponent{}

	startTime, err := common.GenSpecialTimestamp("2017-04-29 00:00:00")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	endTime, err := common.GenSpecialTimestamp("2017-05-06 07:00:00")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// random contractOwners 随机生成的合约拥有者数组
	contractOwners := ownersPubkeys

	contractBody := &protos.ContractBody{
		ContractId:         "UUID-1234-5678-90",
		Cname:              "test create contract ",
		Ctype:              "CREATE",
		Caption:            "futurever",
		Description:        "www.futurever.com",
		ContractState:      "",
		Creator:            randomCreator,
		CreateTime:         common.GenTimestamp(),
		StartTime:          startTime,
		EndTime:            endTime,
		ContractOwners:     contractOwners,
		ContractSignatures: nil,
		ContractAssets:     nil,
		ContractComponents: nil,
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

	return serializeContractModel, nil

}

func generateProtoContract(produceValid bool, optArgs ...map[string]interface{}) ([]byte, error) {
	contractOwnersLen := 1
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

func Test_Sign(t *testing.T) {
	contractOwnersLen := 2
	contractSignaturesLen := contractOwnersLen
	produceValid := true
	optArgs := make(map[string]interface{})
	optArgs["contractOwnersLen"] = contractOwnersLen
	// 生成的合约签名人个数
	optArgs["contractSignaturesLen"] = contractSignaturesLen

	serializeContractModel, err := generatContractModel(produceValid, optArgs)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	var contract ContractModel
	json.Unmarshal([]byte(serializeContractModel), &contract)
	fmt.Println("Generate", produceValid, "contractModel")
	fmt.Printf("[Id=%s, contractSignatures=%v]", contract.Id, contract.ContractBody.ContractSignatures)
	fmt.Println(common.StructSerializePretty(contract))
}

func Test_IsSignatureValid(t *testing.T) {
	//create new obj
	contractModel := &ContractModel{}
	contractHead := &protos.ContractHead{}
	contractBody := &protos.ContractBody{}

	private_key := "Cnodz1gyhaNoFcPCr72G9brFrGFfNJQUPFchGXyL11Pt"
	//private_key := "GEXZrZShFsHKZB94mpRmaDBBtjWCDz6TgK17R9DXUwex"
	// modify and set value for reference obj with &
	contractHead.MainPubkey = "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3"
	//contractHead.MainPubkey = "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo"
	contractBody.Cname = "star"
	contractBody.ContractOwners = []string{
		"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
	}
	contractSignatures := []*protos.ContractSignature{
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
	}
	contractBody.ContractSignatures = contractSignatures

	contractModel.ContractHead = contractHead
	contractModel.ContractBody = contractBody
	// sign for contract
	signatureContract := contractModel.Sign(private_key)

	_ = private_key
	_ = signatureContract

	//fmt.Println("private_key is : ", private_key)
	//fmt.Println("contract is : ", common.StructSerializePretty(contract))
	//fmt.Println("signatureContract is : ", signatureContract)
	fmt.Println("contractModel is : ", contractModel)
	isSignatureValid := contractModel.IsSignatureValid()
	if isSignatureValid {
		t.Log("contract 签名有效")
	} else {
		t.Error("contract 签名无效")
	}
}

func Test_Validate(t *testing.T) {
	//create new obj
	contractModel := ContractModel{}
	contractHead := &protos.ContractHead{}
	contractBody := &protos.ContractBody{}

	private_key := "GEXZrZShFsHKZB94mpRmaDBBtjWCDz6TgK17R9DXUwex"
	//private_key := "Cnodz1gyhaNoFcPCr72G9brFrGFfNJQUPFchGXyL11Pt"
	contractHead.MainPubkey = "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo"
	//contractHead.MainPubkey = "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3"
	//----------------------pub: HFStRMeL1aS824zsApWiEmZqn1r21FvkXQFkqGwJjb3d
	//----------------------pri: EUFo91bZzqtZEAMUZ4Mr5N69BBwfBkKexwXDydTncqqg
	// modify and set value for reference obj with &
	contractBody.Cname = "star"
	contractBody.ContractOwners = []string{
		"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
		//"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
	}
	contractSignatures := []*protos.ContractSignature{
		{
			OwnerPubkey:   "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			Signature:     "4JmRZ2A1Dqf4sGQVS7Jo6nNdR17XxdYddSC3fE4bv6ov48J9CCSMSKmx9AUtkaqJLpsLEGepzjpTbZrXCpbohVeU",
			SignTimestamp: common.GenTimestamp(),
		},
		//{
		//	OwnerPubkey: "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		//	Signature:   "65D27HW4uXYvkekGssAQB93D92onMyU1NVnCJnE1PgRKz2uFSPZ6aQvid4qZvkxys7G4r2Mf2KFn5BSQyEBhWs34",
		//	//Signature:   "3XLffBVuFCZbZU1NcroQDSAgcdDtYQ2UK9ye9q9BzLaMiqjoHtJ3SirW5P9JJkjwAkC9CrguwKMRC36T2e769sqQ",
		//	SignTimestamp:   common.GenTimestamp(),
		//},
	}
	contractModel.ContractHead = contractHead
	contractModel.ContractBody = contractBody
	contractBody.ContractSignatures = contractSignatures
	// sign for contract
	signatureContract := contractModel.Sign(private_key)

	fmt.Println("private_key is : ", private_key)
	fmt.Println("contract is : ", common.StructSerialize(contractBody))
	fmt.Println("signatureContract is : ", signatureContract)

	//contractModel.ContractBody = contractBody
	contractModel.Id = contractModel.GenerateId()
	ok := contractModel.Validate()
	fmt.Println(ok)
}

func Test_HashDataForDisabledHTMLEscape(t *testing.T) {
	jsonByte, _ := ioutil.ReadFile("./test1.json")
	// convert json to interface, order the json and serialize
	//var contractData map[string]interface{}
	//json.Unmarshal(jsonByte, &contractData)
	//fmt.Println(string(jsonByte))

	var contractModel ContractModel
	json.Unmarshal(jsonByte, &contractModel)
	fmt.Println(contractModel)
	fmt.Println(contractModel.ContractBody)

	// disabled the HTMLEscape for &, <, and > to \u0026, \u003c, and \u003e in json string
	contractModelSerializeStr := common.StructSerialize(contractModel.ContractBody)
	fmt.Println(contractModelSerializeStr)
	hashId := common.HashData(contractModelSerializeStr)

	fmt.Println("----------------------hashId:", hashId)
}
