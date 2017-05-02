package rethinkdb

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/model"
	"unicontract/src/core/protos"

	"github.com/astaxie/beego"
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

/*----------------------------Unicontract ops-------------------------------------*/

/*----------------------------- contracts start---------------------------------------*/
func Test_InsertContractStruct(t *testing.T) {
	config.Init()
	//create new obj
	contractModel := model.ContractModel{}
	//TODO

	private_key := "C6WdHVbHAErN7KLoWs9VCBESbAXQG6PxRtKktWzoKytR"
	// modify and set value for reference obj with &
	contract := &contractModel.Contract
	contractHead := &protos.ContractHead{}
	contractBody := &protos.ContractBody{}

	contractHead.MainPubkey = config.Config.Keypair.PublicKey
	contractHead.Version = 1
	contractBody.Cname = "star"
	contractBody.ContractOwners = []string{
		"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
		//"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
	}
	contractBody.CreatorTime = common.GenTimestamp()
	contractBody.Creator = "wangxin"
	contractBody.Caption = "CREATOR"
	contractBody.Description = "合约创建"
	contractBody.ContractId = common.GenerateUUID() //contractId
	// sign for contract
	signatureContract := contractModel.Sign(private_key)
	contractModel.ContractHead = contractHead
	contractModel.ContractBody = contractBody

	fmt.Println("private_key is : ", private_key)
	fmt.Println("contract is : ", common.Serialize(contract))
	fmt.Println("signatureContract is : ", signatureContract)

	contractModel.Id = contractModel.GenerateId()
	isTrue := InsertContract(common.Serialize(contractModel))
	fmt.Println(isTrue)
}

func Test_GetContractById(t *testing.T) {
	id := "9a596c277e80c59b4e70e6a1b53520ba3ccdd954d4e4d1e52c31de61dfbc3c75"
	/*-------------------examples:------------------*/
	contractStr, err := GetContractById(id)
	if err != nil {
		beego.Debug(err)
	}
	if contractStr == "" {
		beego.Debug("Test_GetContractById result is blank")
		return
	}

	var contract model.ContractModel
	json.Unmarshal([]byte(contractStr), &contract)

	if err != nil {
		beego.Debug("Error Test_GetContractById Unmarshal")
		return
	}
	fmt.Println(contract)
	fmt.Println(common.SerializePretty(contract))
}

func Test_GetContractsByContractId(t *testing.T) {
	contractId := "aa3caafc-205d-480e-be39-b6e9e3213059"
	/*-------------------examples:------------------*/
	contractStr, err := GetContractsByContractId(contractId)
	var contracts []model.ContractModel
	json.Unmarshal([]byte(contractStr), &contracts)

	if err != nil {
		fmt.Println("error Test_GetContractsByContractId")
	}
	fmt.Println("records count is ", len(contracts))
	fmt.Println(contracts)
	fmt.Println(common.SerializePretty(contracts))
}

func Test_GetContractMainPubkeyByContract(t *testing.T) {
	//contractId := "834fbab3-9118-45a5-b6d4-31d7baad5e13x"
	id := "ecd4200f171d4be58e3e428b1c104045c7c9fdd367ea6a112c57cd9069eb6720"
	main_pubkey, err := GetContractMainPubkeyByContract(id)
	if err != nil {
		fmt.Println("error Test_GetContractMainPubkeyById")
	}
	fmt.Println("222", main_pubkey)
}

/*----------------------------- contracts end---------------------------------------*/

/*----------------------------- votes start---------------------------------------*/
func Test_InsertVote(t *testing.T) {
	config.Init()
	vote := model.Vote{}

	vote.NodePubkey = config.Config.Keypair.PublicKey
	voteBody := &vote.VoteBody
	voteBody.Timestamp = common.GenTimestamp()
	/*-------------- random false and reason------------------*/
	random_n := common.RandInt(0, 10)
	if random_n <= 6 {
		voteBody.IsValid = true
		voteBody.InvalidReason = ""
	} else {
		voteBody.IsValid = false
		voteBody.InvalidReason = "random false[random is " + strconv.Itoa(random_n) + "]"
	}

	voteBody.VoteType = "CONTRACT"
	voteBody.VoteFor = "ba934f88cf20a3d440efdcc92e93d9722925502c268513187f6e5805e60c0e42"
	vote.Signature = "3FyHdZVX4adfSSTg7rZDPMzqzM8k5fkpu43vbRLvEXLJ"
	vote.Id = common.GenerateUUID()
	isTrue := InsertVote(common.Serialize(vote))
	if isTrue {
		fmt.Println("insert vote success! ", random_n)
	}
}

func Test_GetVoteById(t *testing.T) {
	id := "032af183-5ffb-4091-bfe0-d4aae1af4b5c"
	/*-------------------examples:------------------*/
	voteStr, err := GetVoteById(id)
	fmt.Println(voteStr == "")
	var vote model.Vote
	json.Unmarshal([]byte(voteStr), &vote)

	if err != nil {
		fmt.Println("error Test_GetVoteById")
	} else {
		fmt.Println(common.SerializePretty(vote))
	}
}

func Test_GetVotesByContractId(t *testing.T) {
	contractId := "ecd4200f171d4be58e3e428b1c104045c7c9fdd367ea6a112c57cd9069eb6720"

	/*-------------------examples:------------------*/
	votesStr, err := GetVotesByContractId(contractId)
	fmt.Println(votesStr)
	var votes []model.Vote
	json.Unmarshal([]byte(votesStr), &votes)

	if err != nil {
		fmt.Println("GetVotesByContractId fail!")
	} else {
		fmt.Println("records count is ", len(votes))
		if len(votes) > 0 {
			fmt.Println(common.SerializePretty(votes))
		}
	}
}

/*----------------------------- votes end---------------------------------------*/

/*----------------------------- contractOutputs start---------------------------------------*/
func Test_InsertContractOutput(t *testing.T) {
	conotractOutput := model.ContractOutput{}
	transaction := &conotractOutput.Transaction
	conotractOutput.Version = 1
	transaction.Asset = nil
	transaction.Conditions = nil
	transaction.Fulfillments = nil

	tempMap := make(map[string]interface{})
	tempMap["a"] = "1"
	tempMap["c"] = "3"
	tempMap["b"] = "2"
	transaction.Metadata = &model.Metadata{
		Id:   "meta-data-id",
		Data: tempMap,
	}
	transaction.Operation = "OUTPUT"
	transaction.Timestamp = common.GenTimestamp()

	relaction := &model.Relation{
		ContractId: "834fbab3-9118-45a5-b6d4-31d7baad5e13",
		Voters: []string{
			"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
			"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
		},
	}

	Votes := []*model.Vote{
		{
			Id:         common.GenerateUUID(),
			NodePubkey: "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
			VoteBody: model.VoteBody{
				IsValid:       true,
				InvalidReason: "",
				//IsValid:         false,
				//InvalidReason:   "random false",
				VoteFor:   "7fb5daf3548c2d0d9b71ce25ee962d164cbb87d82078d7361b8424a95c7c4b94",
				VoteType:  "None",
				Timestamp: common.GenTimestamp(),
			},
			Signature: "65D27HW4uXYvkekGssAQB93D92onMyU1NVnCJnE1PgRKz2uFSPZ6aQvid4qZvkxys7G4r2Mf2KFn5BSQyEBhWs34",
		},
		{
			Id:         common.GenerateUUID(),
			NodePubkey: "J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
			VoteBody: model.VoteBody{
				IsValid:       true,
				InvalidReason: "",
				VoteFor:       "7fb5daf3548c2d0d9b71ce25ee962d164cbb87d82078d7361b8424a95c7c4b94",
				VoteType:      "None",
				Timestamp:     common.GenTimestamp(),
			},
			Signature: "5i5dTtQseQjWZ8UdchqQtgttyeeFmB3LDFYzNKafvV2YvTqwv4wZ9mFsH7qgysV9ow893D1h2Xnt1uCXLHtbKrkT",
		},
	}
	relaction.Votes = Votes
	transaction.Relation = relaction
	//create new obj
	contract := model.ContractModel{}
	// modify and set value for reference obj with &
	contractHead := &protos.ContractHead{}
	contractBody := &protos.ContractBody{}
	contractHead.MainPubkey = "qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3"
	contractHead.Version = 1
	contractBody.Cname = "star"
	contractBody.ContractOwners = []string{
		"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
		//"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
	}
	contract.Id = contract.GenerateId()
	contract.ContractHead = contractHead
	contract.ContractBody = contractBody
	transaction.ContractModel = contract
	// sign for contract
	conotractOutput.Id = conotractOutput.GenerateId()
	fmt.Println(common.StructSerialize(conotractOutput))
	fmt.Println(common.StructSerializePretty(conotractOutput))
	isTrue := InsertContractOutput(common.StructSerialize(conotractOutput))
	if isTrue {
		fmt.Println("insert conotractOutput success!")
	}
}

func Test_GetContractOutputById(t *testing.T) {
	id := "4e162ce9d44e057b8e972dd316748bb61543c8f9c6075ff44c448a620557c13ax"
	//contractId := "a888c9204173537aec1949dc8d5ecac718cadcc68966017d9e0ab6d62a5675692"

	/*-------------------examples:------------------*/
	contractOutputStr, err := GetContractOutputById(id)
	fmt.Println(contractOutputStr == "")
	var contractOutput model.ContractOutput
	json.Unmarshal([]byte(contractOutputStr), &contractOutput)

	if err != nil {
		fmt.Println("Test_GetContractOutputById fail!")
	}
	fmt.Println(common.SerializePretty(contractOutput))
}

func Test_GetContractOutputByContractPrimaryId(t *testing.T) {
	contract_Id := "99120e82996f17f6ff5a33c6a7fd0d84491a5653500e136fc14876c956435489"
	//contractId := "a888c9204173537aec1949dc8d5ecac718cadcc68966017d9e0ab6d62a5675692"

	/*-------------------examples:------------------*/
	contractOutputStr, err := GetContractOutputByContractPrimaryId(contract_Id)
	var contractOutput model.ContractOutput
	json.Unmarshal([]byte(contractOutputStr), &contractOutput)

	if err != nil {
		fmt.Println("Test_GetContractOutputByContractPrimaryId fail!")
	}
	fmt.Println(common.SerializePretty(contractOutput))
}

/*----------------------------- contractOutputs end---------------------------------------*/

func Test_GetAllRecords(t *testing.T) {
	idList, _ := GetAllRecords("Unicontract", "SendFailingRecords")
	for _, value := range idList {
		fmt.Println(value)
	}
}

/*----------------------------- consensusFailures start---------------------------------------*/
func Test_InsertConsensusFailure(t *testing.T) {
	/*-------------------examples:------------------*/
	consensusFailure := &model.ConsensusFailure{}
	consensusFailure.Id = common.GenerateUUID()
	consensusFailure.Timestamp = common.GenTimestamp()
	consensusFailure.ConsensusId = "3ea445410f608e6453cdcb7dbe42d57a89aca018993d7e87da85993cbccc6308"

	consensusFailure.ConsensusReason = "random " + strconv.Itoa(common.RandInt(0, 10))
	consensusFailure.ConsensusType = "CONTRACT"
	//consensusFailure.ConsensusType = "TRANSACTION"

	ok := InsertConsensusFailure(common.Serialize(consensusFailure))

	if ok {
		fmt.Println("InsertConsensusFailure success")
	}
}

func Test_GetConsensusFailureById(t *testing.T) {
	id := "0a4957ed-b074-4326-879d-6a26b44843b2"
	//id := "5c63f2c4-a578-450e-8714-66e99c1ad364"
	/*-------------------examples:------------------*/
	consensusFailureStr, err := GetConsensusFailureById(id)
	fmt.Println(consensusFailureStr == "")
	var consensusFailure model.ConsensusFailure
	json.Unmarshal([]byte(consensusFailureStr), &consensusFailure)

	if err != nil {
		fmt.Println("error Test_GetConsensusFailureById")
	}
	fmt.Println(consensusFailure)
	fmt.Println(common.SerializePretty(consensusFailure))
}

func Test_GetConsensusFailuresByConsensusId(t *testing.T) {
	//consensusId := "834fbab3-9118-45a5-b6d4-31d7baad5e13"
	consensusId := "3ea445410f608e6453cdcb7dbe42d57a89aca018993d7e87da85993cbccc6308"
	//contractId := "a888c9204173537aec1949dc8d5ecac718cadcc68966017d9e0ab6d62a5675692"

	/*-------------------examples:------------------*/
	consensusFailuresStr, err := GetConsensusFailuresByConsensusId(consensusId)
	var consensusFailures []model.ConsensusFailure
	json.Unmarshal([]byte(consensusFailuresStr), &consensusFailures)

	if err != nil {
		fmt.Println("Test_GetConsensusFailuresByConsensusId fail!")
	}
	fmt.Println("records count is ", len(consensusFailures))
	//fmt.Println(consensusFailures)
	fmt.Println(common.SerializePretty(consensusFailures))
}

/*----------------------------- consensusFailures end---------------------------------------*/

/*----------------------------- contractTask start---------------------------------------*/
func Test_InsertContractTask(t *testing.T) {
	/*-------------------examples:------------------*/
	contractTask := &model.ContractTask{}
	contractTask.Id = common.GenerateUUID()
	contractTask.ContractStep = "contractTask step..."
	contractTask.ContractCondiction = "contractTask condition..."
	contractTask.ContractId = "123"
	//contractTask.ContractState = "contractTask state..."
	contractTask.ContractState = "contractTask state..." + strconv.Itoa(common.RandInt(0, 10))

	ok := InsertContractTask(common.Serialize(contractTask))

	if ok {
		fmt.Println("Test_InsertContractTask success")
	}
}

func Test_GetContractTaskById(t *testing.T) {
	id := "de70dfdc-0d12-466e-94a7-a1c5cfed0e0e"
	//id := "5c63f2c4-a578-450e-8714-66e99c1ad364"
	/*-------------------examples:------------------*/
	contractTaskStr, err := GetContractTaskById(id)
	var contractTask model.ContractTask
	json.Unmarshal([]byte(contractTaskStr), &contractTask)

	if err != nil {
		fmt.Println("error Test_GetContractTaskById")
	}
	fmt.Println(contractTask)
	fmt.Println(common.SerializePretty(contractTask))
}

func Test_GetContractTasksByContractId(t *testing.T) {
	contractId := "123"

	/*-------------------examples:------------------*/
	contractTasksStr, err := GetContractTasksByContractId(contractId)
	var contractTasks []model.ContractTask
	json.Unmarshal([]byte(contractTasksStr), &contractTasks)

	if err != nil {
		fmt.Println("Test_GetContractTasksByContractId fail!")
	}
	fmt.Println("records count is ", len(contractTasks))
	//fmt.Println(consensusFailures)
	fmt.Println(common.SerializePretty(contractTasks))
}

/*----------------------------- contractTask end---------------------------------------*/

/*----------------------------- TaskSchedule start---------------------------------------*/
func Test_InsertTaskSchedule(t *testing.T) {
	var strTaskSchedule string

	//test 1
	t.Logf("test strTaskSchedule is null\n")
	err := InsertTaskSchedule(strTaskSchedule)
	if err != nil {
		t.Logf("pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Errorf("not pass\n")
	}

	//test 2
	t.Logf("test strTaskSchedule is not null\n")
	var taskSchedule model.TaskSchedule
	taskSchedule.Id = common.GenerateUUID()
	taskSchedule.ContractId = common.GenerateUUID()
	taskSchedule.NodePubkey = config.Config.Keypair.PublicKey
	taskSchedule.StartTime = common.GenTimestamp()
	taskSchedule.EndTime = "1593281188956" //common.GenTimestamp()

	slJson, _ := json.Marshal(taskSchedule)
	strTaskSchedule = string(slJson)
	err = InsertTaskSchedule(strTaskSchedule)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
	}
}

func Test_GetTaskSchedules(t *testing.T) {
	var strNodePubkey string

	//test 1
	t.Logf("test strNodePubkey is null\n")
	_, err := GetTaskSchedules(strNodePubkey)
	if err != nil {
		t.Logf("pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Errorf("not pass\n")
	}

	//test 2
	t.Logf("test strNodePubkey is not null\n")
	strNodePubkey = config.Config.Keypair.PublicKey
	retStr, err := GetTaskSchedules(strNodePubkey)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
		var slTask []model.TaskSchedule
		json.Unmarshal([]byte(retStr), &slTask)
		t.Logf("%+v\n", slTask)
	}
}

func Test_SetTaskScheduleSend(t *testing.T) {
	var strID string

	//test 1
	t.Logf("test strID is null\n")
	err := SetTaskScheduleSend(strID)
	if err != nil {
		t.Logf("pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Errorf("not pass\n")
	}

	//test 2
	t.Logf("test strID is not null\n")
	strID = "25291a3d-4082-4c83-998a-bc2db59dcd82"
	err = SetTaskScheduleSend(strID)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
	}
}

func Test_SetTaskScheduleNoSend(t *testing.T) {
	var strID string

	//test 1
	t.Logf("test strID is null\n")
	err := SetTaskScheduleSend(strID)
	if err != nil {
		t.Logf("pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Errorf("not pass\n")
	}

	//test 2
	t.Logf("test strID is not null\n")
	strID = "25291a3d-4082-4c83-998a-bc2db59dcd82"
	err = SetTaskScheduleNoSend(strID)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
	}
}

/*----------------------------- TaskSchedule end---------------------------------------*/
