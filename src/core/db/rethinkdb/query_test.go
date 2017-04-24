package rethinkdb

import (
	"encoding/json"
	"fmt"
	"testing"
	"unicontract/src/common"
	"unicontract/src/core/model"
	"unicontract/src/core/protos"
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
	contract.CreatorPubkey = "6prCcrjvCz5YwmiraCJko8niFpNQDv9296WoMeDo5FMo"
	contract.Operation = "CREATE"
	contract.ContractOwners = []string{
		//"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		//"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		"6prCcrjvCz5YwmiraCJko8niFpNQDv9296WoMeDo5FMo",
	}
	// sign for contract
	signatureContract := contractModel.Sign(private_key)

	fmt.Println("private_key is : ", private_key)
	fmt.Println("contract is : ", common.Serialize(contract))
	fmt.Println("signatureContract is : ", signatureContract)
	contractModel.Voters = []string{
		"6prCcrjvCz5YwmiraCJko8niFpNQDv9296WoMeDo5FMo",
		//"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
	}
	contractModel.Timestamp = common.GenTimestamp()
	//contractModel.NodePubkey = "AZfjdKxEr9G3NwdAkco22nN8PfgQvCr5TDPK1tqsGZrk"
	contractModel.MainPubkey = "6prCcrjvCz5YwmiraCJko8niFpNQDv9296WoMeDo5FMo"

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

	vote.NodePubkey = "3FyHdZVX4adfSSTg7rZDPMzqzM8k5fkpu43vbRLvEXLJ"
	voteBody := &vote.VoteBody
	voteBody.Timestamp = common.GenTimestamp()
	voteBody.IsValid = true
	voteBody.InvalidReason = ""
	voteBody.VoteType = "CONTRACT"
	voteBody.VoteForContract = "1c54ec6c1c362beba4480b417e6646c0565ae322b0894f6c83f460d78b992711"
	vote.Signature = "3FyHdZVX4adfSSTg7rZDPMzqzM8k5fkpu43vbRLvEXLJ"
	vote.Id = common.GenerateUUID()
	isTrue := InsertVote(common.Serialize(vote))
	if isTrue {
		fmt.Println("insert vote success!")
	}
}

func Test_GetVotesByContractId(t *testing.T) {
	contractId := "a888c9204173537aec1949dc8d5ecac718cadcc68966017d9e0ab6d62a567569"
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
	contract := protos.Contract{}
	// modify and set value for reference obj with &
	contract.CreatorPubkey = "5Pv7F7g9BvNDEMdb8HV5aLHpNTNkxVpNqnLTQ58Z5heC"
	contract.Operation = "CREATE"
	contract.ContractOwners = []string{
		"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		//"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
	}
	// sign for contract

	conotractOutput.Id = conotractOutput.GenerateId()

	isTrue := InsertContractOutput(common.Serialize(conotractOutput))
	if isTrue {
		fmt.Println("insert conotractOutput success!")
	}
}

func Test_ContractOutput(t *testing.T) {
	contractId := "aa70b6ea73c5b9f3b563c4242e3a625c89197ffc8c031cd3c4e1bc278e006645"
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