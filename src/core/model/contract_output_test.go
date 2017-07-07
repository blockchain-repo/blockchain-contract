package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"unicontract/src/common/uniledgerlog"
	"strings"
	"testing"
	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/protos"
)

func init() {
	config.Init()
}
func GenerateOutputTest() string {

	contractOutput := ContractOutput{}
	contractOutput.Version = 2

	transaction := Transaction{}
	transaction.Asset = &Asset{}                 //todo
	transaction.Conditions = []*ConditionsItem{} //todo
	transaction.Fulfillments = []*Fulfillment{}
	//fulfillment := &Fulfillment{
	//	Fid:          0,
	//	OwnersBefore: []string{config.Config.Keypair.PublicKey},
	//}
	//transaction.Fulfillments = []*Fulfillment{
	//	fulfillment,
	//}

	//{
	//
	//	"fid": 0 ,
	//	"fulfillment": "cf:4:RtTtCxNf1Bq7MFeIToEosMAa3v_jKtZUtqiWAXyFz1ejPMv-t7vT6DANcrYvKFHAsZblmZ1Xk03HQdJbGiMyb5CmQqGPHwlgKusNu9N_IDtPn7y16veJ1RBrUP-up4YD" ,
	//	"input": null ,
	//	"owners_before": [
	//	"5mVrPtqUzXwKYL2JeZo4cQq2spt8qfGVx3qE2V7NqgyU"
	//	]
	//
	//}
	tempMap := make(map[string]interface{})
	tempMap["a"] = "1"
	tempMap["c"] = "3"
	tempMap["b"] = "2"
	tempMap["A"] = "4"
	tempMap["6"] = 5

	transaction.Metadata = &Metadata{
		Id:   "meta-data-id",
		Data: tempMap,
	}
	transaction.Operation = "CONTRACT"
	//transaction.Timestamp = ""

	//--------------------contract-------------------------
	contractAsset := []*protos.ContractAsset{}
	contractComponent := []*protos.ContractComponent{}
	contractHead := &protos.ContractHead{config.Config.Keypair.PublicKey, 1,
		common.GenTimestamp(), common.GenTimestamp(), 0}

	contractOwners := []string{
		"BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
		"4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
		"9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
	}
	contractBody := &protos.ContractBody{
		ContractId:         "feca0672-4ad7-4d9a-ad57-83d48db2269b",
		Cname:              "test contract output",
		Ctype:              "CREATE",
		Caption:            "购智能手机返话费合约产品协议",
		Description:        "移动用户A花费500元购买移动运营商B的提供的合约智能手机C后",
		ContractState:      "",
		Creator:            common.GenTimestamp(),
		CreateTime:         "1493111926720",
		StartTime:          "1493111926730",
		EndTime:            "1493111926740",
		ContractOwners:     contractOwners,
		ContractSignatures: nil,
		ContractAssets:     contractAsset,
		ContractComponents: contractComponent,
		MetaAttribute:      nil,
	}
	transaction.ContractModel.ContractHead = nil
	transaction.ContractModel.ContractBody = contractBody

	contractSignatures := []*protos.ContractSignature{
		{
			OwnerPubkey:   "BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
			Signature:     transaction.ContractModel.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
			SignTimestamp: "1493111926751",
		},
		{
			OwnerPubkey:   "4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
			Signature:     transaction.ContractModel.Sign("AnV4aa3KCpsNF8bEqQ8qF8T97iW4KnhMmPKwaFWgKhRo"),
			SignTimestamp: "1493111926752",
		},
		{
			OwnerPubkey:   "9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
			Signature:     transaction.ContractModel.Sign("9647UfPdDSwBf5kw7tUrSe7cmYY5RvVX47GrGqSh4XVi"),
			SignTimestamp: "1493111926753",
		},
	}
	contractBody.ContractSignatures = contractSignatures

	transaction.ContractModel.Id = common.HashData(common.StructSerialize(contractBody))

	//--------------------relaction-------------------------
	transaction.Relation = &Relation{
		ContractHashId: transaction.ContractModel.Id,
		ContractId:     transaction.ContractModel.ContractBody.ContractId,
		TaskId:         "task-id-123456789",
		TaskExecuteIdx: 0,
		Voters: []string{
			config.Config.Keypair.PublicKey, config.Config.Keypair.PublicKey, config.Config.Keypair.PublicKey,
		},
	}

	contractOutput.Version = 2
	contractOutput.Transaction = transaction
	fmt.Println("hash-pre: ", common.StructSerialize(contractOutput))
	contractOutput.Id = common.HashData(common.StructSerialize(contractOutput))
	//fulfillment.Fulfillment = "cf:4:RtTtCxNf1Bq7MFeIToEosMAa3v_jKtZUtqiWAXyFz1ejPMv-t7vT6DANcrYvKFHAsZblmZ1Xk03HQdJbGiMyb5CmQqGPHwlgKusNu9N_IDtPn7y16veJ1RBrUP-up4YD"

	//operation:transfer
	//vote1 := &Vote{}
	//vote1.Id = common.GenerateUUID()
	//vote1.NodePubkey = config.Config.Keypair.PublicKey
	//vote1.VoteBody.Timestamp = common.GenTimestamp()
	//vote1.VoteBody.InvalidReason = "resion"
	//vote1.VoteBody.IsValid = true
	//vote1.VoteBody.VoteFor = contractOutput.Id
	//vote1.VoteBody.VoteType = "TRANSACTION"
	////note: contractoutput(transaction) node signatrue : use the contractOutput.id
	//vote1.Signature = common.Sign(config.Config.Keypair.PrivateKey, contractOutput.Id)
	//vote2 := &Vote{}
	//vote2.Id = common.GenerateUUID()
	//vote2.NodePubkey = config.Config.Keypair.PublicKey
	//vote2.VoteBody.Timestamp = common.GenTimestamp()
	//vote2.VoteBody.InvalidReason = "resion"
	//vote2.VoteBody.IsValid = true
	//vote2.VoteBody.VoteFor = contractOutput.Id
	//vote2.VoteBody.VoteType = "TRANSACTION"
	//vote2.Signature = common.Sign(config.Config.Keypair.PrivateKey, contractOutput.Id)
	//vote3 := &Vote{}
	//vote3.Id = common.GenerateUUID()
	//vote3.NodePubkey = config.Config.Keypair.PublicKey
	//vote3.VoteBody.Timestamp = common.GenTimestamp()
	//vote3.VoteBody.InvalidReason = "resion"
	//vote3.VoteBody.IsValid = true
	//vote3.VoteBody.VoteFor = contractOutput.Id
	//vote3.VoteBody.VoteType = "TRANSACTION"
	//vote3.Signature = common.Sign(config.Config.Keypair.PrivateKey, contractOutput.Id)
	//transaction.Relaction.Votes = []*Vote{
	//	vote1, vote2, vote3,
	//}
	//operation:contract
	vote4 := &Vote{}
	vote4.Id = common.GenerateUUID()
	vote4.NodePubkey = config.Config.Keypair.PublicKey
	vote4.VoteBody.Timestamp = common.GenTimestamp()
	vote4.VoteBody.InvalidReason = "resion"
	vote4.VoteBody.IsValid = true
	vote4.VoteBody.VoteFor = transaction.ContractModel.Id
	vote4.VoteBody.VoteType = "CONTRACT"
	//note:contractoutput(contract) node signature : use the vote data
	//uniledgerlog.Info("voteSign: ",common.Serialize(vote4.VoteBody))
	vote4.Signature = vote4.SignVote()
	vote5 := &Vote{}
	vote5.Id = common.GenerateUUID()
	vote5.NodePubkey = config.Config.Keypair.PublicKey
	vote5.VoteBody.Timestamp = common.GenTimestamp()
	vote5.VoteBody.InvalidReason = "resion"
	vote5.VoteBody.IsValid = true
	vote5.VoteBody.VoteFor = transaction.ContractModel.Id
	vote5.VoteBody.VoteType = "CONTRACT"
	vote5.Signature = vote5.SignVote()
	vote6 := &Vote{}
	vote6.Id = common.GenerateUUID()
	vote6.NodePubkey = config.Config.Keypair.PublicKey
	vote6.VoteBody.Timestamp = common.GenTimestamp()
	vote6.VoteBody.InvalidReason = "resion"
	vote6.VoteBody.IsValid = true
	vote6.VoteBody.VoteFor = transaction.ContractModel.Id
	vote6.VoteBody.VoteType = "CONTRACT"
	vote6.Signature = vote6.SignVote()
	transaction.Relation.Votes = []*Vote{
		vote4, vote5, vote6,
	}

	//--------------------contract-out-put-------------------------
	//contractOutput.Transaction.Timestamp = common.GenTimestamp()
	contractOutput.Transaction.ContractModel.ContractHead = contractHead

	fmt.Println(common.Serialize(contractOutput))
	return common.Serialize(contractOutput)
}

func Test_InserContractOutput(t *testing.T) {
	str := GenerateOutputTest()
	b := rethinkdb.InsertContractOutput(str)
	fmt.Println(b)
}

func Test_ContractOutput(t *testing.T) {
	contractOutput := ContractOutput{}
	contractOutput.Id = "1"
	//transaction := contractOutput.Transaction

	conditions := make([]*ConditionsItem, 1)
	conditions[0] = &ConditionsItem{
		Amount: 14213,
		Cid:    0,
		Condition: &Condition{
			Details: nil,
			Uri:     "dd-dsd-qwq-ddd-aa",
		},
		OwnersAfter: []string{"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc"},
	}
	contractOutput.Transaction.Conditions = conditions
	contractOutput.Transaction.Timestamp = common.GenTimestamp()

	result := common.Serialize(contractOutput)
	//result := common.SerializePretty(contractOutput)
	fmt.Println(result)
}

func Test_GenerateId(t *testing.T) {
	contractOutput := ContractOutput{}
	contractOutput.Id = "1"
	transaction := contractOutput.Transaction
	conditions := make([]*ConditionsItem, 1)
	conditions[0] = &ConditionsItem{
		Amount: 14213,
		Cid:    0,
		Condition: &Condition{
			Details: nil,
			Uri:     "dd-dsd-qwq-ddd-aa-ww",
		},
		OwnersAfter: []string{"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc"},
	}
	contractOutput.Transaction.Conditions = conditions

	Votes := []*Vote{
		{
			Id:         common.GenerateUUID(),
			NodePubkey: "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
			VoteBody: VoteBody{
				IsValid:       true,
				InvalidReason: "",
				VoteFor:       "",
				VoteType:      "",
				Timestamp:     common.GenTimestamp(),
			},
			Signature: "65D27HW4uXYvkekGssAQB93D92onMyU1NVnCJnE1PgRKz2uFSPZ6aQvid4qZvkxys7G4r2Mf2KFn5BSQyEBhWs34",
		},
		{
			Id:         common.GenerateUUID(),
			NodePubkey: "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			VoteBody: VoteBody{
				IsValid:       true,
				InvalidReason: "",
				VoteFor:       "",
				VoteType:      "",
				Timestamp:     common.GenTimestamp(),
			},
			Signature: "5i5dTtQseQjWZ8UdchqQtgttyeeFmB3LDFYzNKafvV2YvTqwv4wZ9mFsH7qgysV9ow893D1h2Xnt1uCXLHtbKrkT",
		},
	}

	relaction := &Relation{
		ContractId: "3ea445410f608e6453cdcb7dbe42d57a89aca018993d7e87da85993cbccc6308",
		TaskId:     "123",
		Voters:     []string{"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3", "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo"},
		Votes:      Votes,
	}
	transaction.Relation = relaction

	transaction.Timestamp = common.GenTimestamp()
	contractOutput.Transaction = transaction
	fmt.Println(transaction)
	//result := common.Serialize(contractOutput)
	contract_output_id := contractOutput.GenerateId()
	fmt.Println("contract_output_id= " + contract_output_id)

	//transaction.Relaction.Signatures = nil
	fmt.Println(transaction)
	//fmt.Printf("2222%+v\n", transaction)
	fmt.Println("contract_output_id2= " + contract_output_id)

	expected_id := "6834f1539cb01247dd4b3e1b789b09e4a3a7477d8d55b69555e1311de6130ab0"
	//expected_id := "ea9d67ef76162ff69b5a4911e2acf8d8b088a968031f3af8168eee71ca1fdc01"
	//expected_id := "2b2cd0b17a07407e83155d0f5999eb107b9127eee1b876124322e5b17c9b01c4"
	//expected_id := "ea9d67ef76162ff69b5a4911e2acf8d8b088a968031f3af8168eee71ca1fdc012"
	if contract_output_id != expected_id {
		t.Error()
	}

}

func TestHash(t *testing.T) {
	//temp := "hello unich>ain 2017"
	ddd := `{"Caption":"购智能手机返话费合约产品协议","Cname":"contract_mobilecallback","ContractId":"contract test from js->java->go-python","ContractOwners":["3FyHdZVX4adfSSTg7rZDPMzqzM8k5fkpu43vbRLvEXLJ","BbfY4Dc5s8dTP1Z1yixnetezRKYREHqwbP8GQGh3WyVS"],"StartTime":"1496535120000"}`
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(ddd)

	ser := strings.TrimSpace(buf.String())
	uniledgerlog.Info(ser)
	hash := common.HashData(ser)
	uniledgerlog.Info(hash)
}

//c2219c1a23b637f9a1dac5cbea7625b2de8b100a6e65e5a92f555b7ab8e37424
//b5ce10fe914d4195fcb8e9072e3cdaaded517c2a2913759978ffc46ab5024db7

//408c606f08e709778a68a5f6adc793caf7907f28f9d41772c013e305195cd343

//fd6264142534a7a68431f5bee5fab7b74609ccfceff8a67ff5af10b3d200e2b0	hello unich>ain 2017
//27045bf50089651906c5cc43510332c33e9df035428bf6989ed58a1bb7ea1832  hello unichain 2017
