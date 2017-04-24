package rethinkdb

import (
	"encoding/json"
	"fmt"
	"testing"
	"unicontract/src/common"
	"unicontract/src/core/model"
	"unicontract/src/core/protos"
	"unicontract/src/config"
)

func Test_Get(t *testing.T) {
	res := Get("Unicontract", "Contracts", "123151f1ddassd")
	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		fmt.Printf("Error scanning database result: %s", err)
	}
	str := common.Serialize(blo)
	fmt.Printf("blo:%s\n", str)

}

func Test_Insert(t *testing.T) {
	res := Insert("bigchain", "votes", "{\"back\":\"jihhh\"}")
	fmt.Printf("%d row inserted", res.Inserted)
}

func Test_Update(t *testing.T) {
	res := Update("bigchain", "votes", "37adc1b6-e22a-4d39-bc99-f1f44608a15b", "{\"1111back\":\"j111111111111ihhh\"}")
	fmt.Printf("%d row replaced", res.Replaced)
}

func Test_Delete(t *testing.T) {
	res := Delete("bigchain", "votes", "37adc1b6-e22a-4d39-bc99-f1f44608a15b")
	fmt.Printf("%d row deleted", res.Deleted)
}

/*----------------------------unicontract ops-------------------------------------*/

func Test_InsertContractStruct(t *testing.T) {
	//create new obj
	contractModel := model.ContractModel{}
	//TODO

	private_key := "C6WdHVbHAErN7KLoWs9VCBESbAXQG6PxRtKktWzoKytR"
	// modify and set value for reference obj with &
	contract := &contractModel.Contract
	contractHead := &protos.ContractHead{}
	contractBody := &protos.ContractBody{}

	contractHead.MainPubkey = "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo"
	contractBody.Cname = "star"
	contractBody.ContractOwners = []string{
		"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
		//"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
	}
	// sign for contract
	signatureContract := contractModel.Sign(private_key)

	fmt.Println("private_key is : ", private_key)
	fmt.Println("contract is : ", common.Serialize(contract))
	fmt.Println("signatureContract is : ", signatureContract)

	contractModel.Id = contractModel.GenerateId()
	isTrue := InsertContract(common.Serialize(contractModel))
	fmt.Println(isTrue)
}

func Test_GetContractById(t *testing.T) {
	contractId := "a888c9204173537aec1949dc8d5ecac718cadcc68966017d9e0ab6d62a567569"
	/*-------------------examples:------------------*/
	contractStr, err := GetContractById(contractId)
	var contract model.ContractModel
	json.Unmarshal([]byte(contractStr), &contract)

	if err != nil {
		fmt.Println("error Test_GetContractById")
	}
	fmt.Println(contract)
	fmt.Println(common.SerializePretty(contract))
}

func Test_GetContractMainPubkeyById(t *testing.T) {
	contractId := "a888c9204173537aec1949dc8d5ecac718cadcc68966017d9e0ab6d62a567569"
	main_pubkey, err := GetContractMainPubkeyById(contractId)
	if err != nil {
		fmt.Println("error Test_GetContractMainPubkeyById")
	}
	fmt.Println(main_pubkey)
}

func Test_InsertVote(t *testing.T) {
	vote := model.Vote{}

	vote.NodePubkey = config.Config.Keypair.PublicKey
	voteBody := &vote.VoteBody
	voteBody.Timestamp = common.GenTimestamp()
	voteBody.IsValid = true
	voteBody.InvalidReason = ""
	voteBody.VoteType = "CONTRACT"
	voteBody.VoteForContract = "73544efa921145090f521672169dbdbd8ffa684f16988a40d1817b3a50d717d6"
	vote.Signature = "3FyHdZVX4adfSSTg7rZDPMzqzM8k5fkpu43vbRLvEXLJ"
	vote.Id = common.GenerateUUID()
	isTrue := InsertVote(common.Serialize(vote))
	if isTrue {
		fmt.Println("insert vote success!")
	}
}

func Test_GetVotesByContractId(t *testing.T) {
	contractId := "73544efa921145090f521672169dbdbd8ffa684f16988a40d1817b3a50d717d6"
	//contractId := "a888c9204173537aec1949dc8d5ecac718cadcc68966017d9e0ab6d62a5675692"

	/*-------------------examples:------------------*/
	votesStr, err := GetVotesByContractId(contractId)
	var votes []model.Vote
	json.Unmarshal([]byte(votesStr), &votes)

	if err != nil {
		fmt.Println("GetVotesByContractId fail!")
	}
	fmt.Println(votes)
	fmt.Println(common.SerializePretty(votes))
}

func Test_InsertContractOutput(t *testing.T) {
	conotractOutput := model.ContractOutput{}
	transaction := &conotractOutput.Transaction
	transaction.Asset = nil
	transaction.Conditions = nil
	transaction.Fulfillments = nil
	transaction.Metadata = nil
	transaction.Operation = "OUTPUT"
	transaction.Timestamp = common.GenTimestamp()

	relaction := model.Relaction{
		ContractId: "a888c9204173537aec1949dc8d5ecac718cadcc68966017d9e0ab6d62a567569",
		Voters: []string{
			"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
			"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		},
	}

	signatures := []*model.RelactionSignature{
		{
			ContractNodePubkey: "JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
			Signature:          "EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
		},
		{
			ContractNodePubkey: "JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
			Signature:          "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		},
	}
	relaction.Signatures = signatures

	//create new obj
	contract := model.ContractModel{}
	// modify and set value for reference obj with &
	contractHead := &protos.ContractHead{}
	contractBody := &protos.ContractBody{}
	contractHead.MainPubkey = "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo"
	contractBody.Cname = "star"
	contractBody.ContractOwners = []string{
		"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
		//"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
	}
	contract.ContractHead = contractHead
	contract.ContractBody = contractBody
	transaction.ContractModel = contract
	// sign for contract
	conotractOutput.Id = conotractOutput.GenerateId()

	isTrue := InsertContractOutput(common.Serialize(conotractOutput))
	if isTrue {
		fmt.Println("insert conotractOutput success!")
	}
}

func Test_ContractOutput(t *testing.T) {
	contractId := "3af8ef59ccbab0510d4afaace6729bc8110d1614a6380eefe87ee3cd1859cc15"
	//contractId := "a888c9204173537aec1949dc8d5ecac718cadcc68966017d9e0ab6d62a5675692"

	/*-------------------examples:------------------*/
	votesStr, err := GetVotesByContractId(contractId)
	var votes []model.Vote
	json.Unmarshal([]byte(votesStr), &votes)

	if err != nil {
		fmt.Println("GetVotesByContractId fail!")
	}
	//fmt.Println(votes)
	fmt.Println(common.SerializePretty(votes))
}

func Test_GetAllRecords(t *testing.T){
	idList,_ :=GetAllRecords("Unicontract","SendFailingRecords")
	for _,value := range idList {
		fmt.Println(value)
	}
}