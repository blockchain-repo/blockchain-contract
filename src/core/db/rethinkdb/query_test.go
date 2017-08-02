package rethinkdb

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"testing"
	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/model"

	"reflect"

	"github.com/astaxie/beego"
	"unicontract/src/common/uniledgerlog"
)

func init() {
	config.Init()
}

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
	//create new obj
	contractModel := model.ContractModel{}
	//TODO

	private_key := "C6WdHVbHAErN7KLoWs9VCBESbAXQG6PxRtKktWzoKytR"
	// modify and set value for reference obj with &
	//contract := &contractModel.Contract
	contractHead := &model.ContractHead{}
	contractBody := &model.ContractBody{}

	contractHead.MainPubkey = config.Config.Keypair.PublicKey
	contractHead.Version = 1
	contractBody.Cname = "star"
	contractBody.ContractOwners = []string{
		"qC5zpgJBqUdqi3Gd6ENfGzc5ZM9wrmqmiPX37M9gjq3",
		"J2rSKoCuoZE1MKkXGAvETp757ZuARveRvJYAzJxqEjoo",
		//"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
	}
	contractBody.CreateTime = common.GenTimestamp()
	contractBody.Creator = "wangxin"
	contractBody.Caption = "CREATOR"
	contractBody.Description = "合约创建"
	contractBody.ContractId = common.GenerateUUID() //contractId
	// sign for contract
	signatureContract := contractModel.Sign(private_key)
	contractModel.ContractHead = contractHead
	contractModel.ContractBody = contractBody
	contractModel.ContractBody.ContractState = "Contract_Signature"

	fmt.Println("private_key is : ", private_key)
	fmt.Println("contract is : ", common.Serialize(contractModel))
	fmt.Println("signatureContract is : ", signatureContract)

	contractModel.Id = contractModel.GenerateId()
	isTrue := InsertContract(common.Serialize(contractModel))
	fmt.Println(isTrue)
}

func Test_GetContractById(t *testing.T) {
	id := "813410e5e448924010c3b5574beb2f6449bf2dd49ae0d4faea62030c37b23a2"
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

func Test_SetContractConsensusResultById(t *testing.T) {
	id := "4591d2c4c88c1ca6ea763956cf64070fc8ef3ad14c4f98277205819efe66b4c4"
	err := SetContractConsensusResultById(id, common.GenTimestamp(), 1)
	if err != nil {
		t.Error(err)
	}
}

func Test_DeleteContract(t *testing.T) {
	id := "1663c124ba5f28c5e0a030da52646144e69156f3ad2b311d4929d66291d2b4fe"
	strContract, err := GetContractById(id)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(strContract)
	}

	var contract_ model.ContractModel
	err = json.Unmarshal([]byte(strContract), &contract_)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v\n", contract_)
	}

	t.Logf("delete contract return is [ %t ]\n", DeleteContract(contract_.Id))

	contract_.ContractHead.MainPubkey = "ASDFASDFASDFASDFASDF"

	sldata, err := json.Marshal(contract_)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(InsertContract(string(sldata)))
	}
}

func Test_GetNoConsensusContracts(t *testing.T) {
	timestamp := common.GenTimestamp()
	strContracts, err := GetNoConsensusContracts(timestamp, 2)
	if err != nil {
		t.Error(err)
	}

	if len(strContracts) == 0 {
		t.Log("is null")
	}

	t.Log(strContracts)
}

func Test_GetContractsCount(t *testing.T) {

	count, err := GetContractsCount()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(count)

	number, err := strconv.Atoi(count)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(number)
}

func Test_GetContractStatsCount(t *testing.T) {

	count, err := GetContractStatsCount("Contract_In_Process")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(count)

	//number, err := strconv.Atoi(count)
	//if err != nil {
	//	t.Error(err)
	//}
	//fmt.Println(number)
}

func Test_QueryOutput(t *testing.T) {
	output, err := QueryOutput("170614121259039057")
	if err != nil {
		t.Error(err)
	}
	t.Log(output)
}

func Test_QueryOutputNum(t *testing.T) {
	count, err := QueryOutputNum("170614121259039057")
	if err != nil {
		t.Error(err)
	}
	t.Log(count)
}

func Test_QueryContractStartTime(t *testing.T) {
	startTime, err := QueryContractStartTime("170614121259039057")
	if err != nil {
		t.Error(err)
	}
	t.Log(startTime)
}

/*----------------------------- contracts end---------------------------------------*/

/*----------------------------- votes start---------------------------------------*/
func Test_InsertVote(t *testing.T) {
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
	voteBody.VoteFor = "676c5244facc65629dcfab324fbf8499724e6b685cad20a90db63ba47eddaf78"
	vote.Signature = vote.SignVote() //"3FyHdZVX4adfSSTg7rZDPMzqzM8k5fkpu43vbRLvEXLJ"
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

func Test_DeleteVotes(t *testing.T) {
	slID := make([]interface{}, 0)
	slID = append(slID, "5e3f5fc2-1e1f-487f-a673-a0cebe30aca3")
	slID = append(slID, "10b75956-0087-43db-80f4-ca7bdb478002")
	slID = append(slID, "8903222a-0824-4fad-8a75-2c4f1902ac47")

	deleteNum, err := DeleteVotes(slID)
	t.Logf("deleteNum is %d\n", deleteNum)
	t.Logf("err is %+v\n", err)
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
	contractHead := &model.ContractHead{}
	contractBody := &model.ContractBody{}
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

func Test_GetConsensusFailuresCount(t *testing.T) {
	count, _ := GetConsensusFailuresCount()
	fmt.Println(count)
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

func TestSession(t *testing.T) {
	session := ConnectDB(DBNAME)

	fmt.Println(reflect.TypeOf(session))
}

/*智能微网demo start---------------------------------------------------------*/
func Test_InsertEnergyTradingDemoRole(t *testing.T) {
	var slString []string

	strPublicKey, _ := common.GenerateKeyPair()
	person1 := model.DemoRole{
		Id:          common.GenerateUUID(),
		Name:        "个人",
		PublicKey:   strPublicKey,
		Infermation: "",
		Type:        0,
	}
	sldata, _ := json.Marshal(person1)
	slString = append(slString, string(sldata))

	strPublicKey, _ = common.GenerateKeyPair()
	mapInformation := make(map[string]string)
	mapInformation["ownerPublicKey"] = person1.PublicKey
	slInformation, _ := json.Marshal(mapInformation)
	electricityMeter1 := model.DemoRole{
		Id:          common.GenerateUUID(),
		Name:        "个人电表",
		PublicKey:   strPublicKey,
		Infermation: string(slInformation),
		Type:        1,
	}
	sldata, _ = json.Marshal(electricityMeter1)
	slString = append(slString, string(sldata))

	strPublicKey, _ = common.GenerateKeyPair()
	operator1 := model.DemoRole{
		Id:          common.GenerateUUID(),
		Name:        "运营商",
		PublicKey:   strPublicKey,
		Infermation: "",
		Type:        2,
	}
	sldata, _ = json.Marshal(operator1)
	slString = append(slString, string(sldata))

	strPublicKey, _ = common.GenerateKeyPair()
	wind1 := model.DemoRole{
		Id:          common.GenerateUUID(),
		Name:        "风电",
		PublicKey:   strPublicKey,
		Infermation: "",
		Type:        3,
	}
	sldata, _ = json.Marshal(wind1)
	slString = append(slString, string(sldata))

	strPublicKey, _ = common.GenerateKeyPair()
	light1 := model.DemoRole{
		Id:          common.GenerateUUID(),
		Name:        "光电",
		PublicKey:   strPublicKey,
		Infermation: "",
		Type:        4,
	}
	sldata, _ = json.Marshal(light1)
	slString = append(slString, string(sldata))

	strPublicKey, _ = common.GenerateKeyPair()
	fire1 := model.DemoRole{
		Id:          common.GenerateUUID(),
		Name:        "火电",
		PublicKey:   strPublicKey,
		Infermation: "",
		Type:        5,
	}
	sldata, _ = json.Marshal(fire1)
	slString = append(slString, string(sldata))

	strPublicKey, _ = common.GenerateKeyPair()
	guowang1 := model.DemoRole{
		Id:          common.GenerateUUID(),
		Name:        "国网",
		PublicKey:   strPublicKey,
		Infermation: "",
		Type:        6,
	}
	sldata, _ = json.Marshal(guowang1)
	slString = append(slString, string(sldata))

	for _, v := range slString {
		err := InsertEnergyTradingDemoRole(v)
		if err != nil {
			t.Error(err)
		}
	}
}

func Test_InsertEnergyTradingDemoEnergy(t *testing.T) {
	// 模拟采集电表耗电
	electricityMeter1 := model.DemoEnergy{
		Id:               common.GenerateUUID(),
		PublicKey:        "GpNgHgCAvaVXHwqyPCNR6yKd92Xshde87RejFaH3v3oB",
		Timestamp:        common.GenTimestamp(),
		Electricity:      580,
		TotalElectricity: 280,
		Money:            110,
		Type:             0,
	}
	sldata, _ := json.Marshal(electricityMeter1)

	err := InsertEnergyTradingDemoEnergy(string(sldata))
	if err != nil {
		t.Error(err)
	}

	/*
			type DemoEnergy struct {
			Id               string `json:"id"`
			PublicKey        string
			Timestamp        string  // 时间戳
			Electricity      float64 // 当前总电量（电）
			TotalElectricity float64 // 只当为电表时有效，当月总耗电量（电）
			Money            float64 // 只当为电表时有效，代表当前表内余额（钱）
			Type             int     // 0：电表 1：风电 2：光电 3：火电 4：国网
		}
	*/
	// 模拟采集发电厂发电
	/*slKey := []string{
		"9Vqg4tSk9ocLfhwj2eeNgKgNR65oSV7WF9kYDu1HiwdM", // 风
		"3XmEh9ZtvDAcxtgiFL11cw9YAppCqhQaWQ6mrKxWhbom", // 光
		"H7tMDKFPMGsG2pV4Lpcic5MQiN1fKkqVaj6A15MMgNTQ", // 火
		"4nkFyWhLrUAGZxr1Ku5NreywPPA6HEKkpqV2hDgr1kLU", // 国网
	}
	var m float64
	for i, v := range slKey {
		if i == 0 {
			m = float64(55)
		} else if i == 1 {
			m = float64(45)
		} else if i == 2 {
			m = float64(65)
		} else {
			m = float64(60)
		}
		electricityPowerPlant1 := model.DemoEnergy{
			Id:          common.GenerateUUID(),
			PublicKey:   v,
			Timestamp:   common.GenTimestamp(),
			Electricity: m,
			Type:        i + 1,
		}
		sldata, _ := json.Marshal(electricityPowerPlant1)

		err := InsertEnergyTradingDemoEnergy(string(sldata))
		if err != nil {
			t.Error(err)
		}
	}*/
}

func Test_InsertTransaction_Bill(t *testing.T) {
	// bill
	/*
			type DemoBill struct {
			Id        string `json:"id"`
			PublicKey string
			Timestamp string
			Type      int // 0：用户账户充值 1：用户购电充值 2：分张
		}
	*/
	strPublicKey, _ := common.GenerateKeyPair()
	bill1 := model.DemoBill{
		Id:        common.GenerateUUID(),
		PublicKey: strPublicKey,
		Timestamp: common.GenTimestamp(),
		Type:      1,
	}
	sldata, _ := json.Marshal(bill1)
	err := InsertEnergyTradingDemoBill(string(sldata))
	if err != nil {
		t.Error(err)
		return
	}

	// transaction
	/*
			type DemoTransaction struct {
			Id            string  `json:"id"`
			BillId        string  // 对应的票据表id
			Timestamp     string  // 交易时间戳
			FromPublicKey string  // 付款方
			ToPublicKey   string  // 收款方
			Money         float64 // 金额
			Type          int     // 0：用户账户充值 1：用户购电充值 2：分张
		}
	*/
	transaction1 := model.DemoTransaction{
		Id:            common.GenerateUUID(),
		BillId:        bill1.Id,
		Timestamp:     common.GenTimestamp(),
		FromPublicKey: "DwRyavLipw2mtd531bhFZMLyKWGTeGgCKJbcQitQj4Ca",
		ToPublicKey:   "DmVKFKGsrKXqBMwEW7kL1t3cENHMF9t41QNunJYBnf1k",
		Money:         50,
		Type:          1,
	}
	sldata, _ = json.Marshal(transaction1)

	err = InsertEnergyTradingDemoTransaction(string(sldata))
	if err != nil {
		t.Error(err)
	}
}

func Test_InsertEnergyTradingDemoPrice(t *testing.T) {
	/*
			type DemoPrice struct {
			Id          string  `json:"id"`
			Level       int     // 阶梯等级
			One         float64 // 波峰电价
			Two         float64 // 波平电价
			Three       float64 // 波谷电价
			Description string  // 阶梯电价描述
		}
	*/
	price1 := model.DemoPrice{
		Id:          common.GenerateUUID(),
		Level:       1,
		Low:         1,
		High:        240,
		One:         1,
		Two:         2,
		Three:       3,
		Description: "第一阶梯",
	}
	sldata, _ := json.Marshal(price1)
	err := InsertEnergyTradingDemoPrice(string(sldata))
	if err != nil {
		t.Error(err)
	}

	price2 := model.DemoPrice{
		Id:          common.GenerateUUID(),
		Level:       2,
		Low:         241,
		High:        400,
		One:         2,
		Two:         3,
		Three:       4,
		Description: "第二阶梯",
	}
	sldata, _ = json.Marshal(price2)
	err = InsertEnergyTradingDemoPrice(string(sldata))
	if err != nil {
		t.Error(err)
	}

	price3 := model.DemoPrice{
		Id:          common.GenerateUUID(),
		Level:       3,
		Low:         401,
		High:        math.MaxFloat64,
		One:         3,
		Two:         4,
		Three:       5,
		Description: "第三阶梯",
	}
	sldata, _ = json.Marshal(price3)
	err = InsertEnergyTradingDemoPrice(string(sldata))
	if err != nil {
		t.Error(err)
	}
}

func Test_GetMeterKeyByUserKey(t *testing.T) {
	key := "64mDgEqY9KGp3NCfJPrrjiruL9hmuYiimmaD2234UYWd"
	meterKey, err := GetMeterKeyByUserKey(key)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(meterKey)
}

func Test_GetPrice(t *testing.T) {
	price, err := GetPrice()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(price)
}

func Test_GetMeterQueryLastTime(t *testing.T) {
	timestamp, err := GetMeterQueryLastTime("5x1hxnPWpHRpvwR3tdo7ygPZ77sSUkywY56VhGhaLpUm")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(timestamp)
}

func Test_UpdateMeterQueryLastTime(t *testing.T) {
	err := UpdateMeterQueryLastTime("5x1hxnPWpHRpvwR3tdo7ygPZ77sSUkywY56VhGhaLpUm", "1497595656727")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("ok")
}

func Test_GetMeterinforFromEnergy(t *testing.T) {
	item, err := GetMeterinforFromEnergy("5x1hxnPWpHRpvwR3tdo7ygPZ77sSUkywY56VhGhaLpUm", "1497595656727", false)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(item)

	item, err = GetMeterinforFromEnergy("5x1hxnPWpHRpvwR3tdo7ygPZ77sSUkywY56VhGhaLpUm", common.GenTimestamp(), true)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(item)
}

func Test_GetPowerPlantEnergy(t *testing.T) {
	slKey := []string{
		"9Vqg4tSk9ocLfhwj2eeNgKgNR65oSV7WF9kYDu1HiwdM", // 风
		"3XmEh9ZtvDAcxtgiFL11cw9YAppCqhQaWQ6mrKxWhbom", // 光
		"H7tMDKFPMGsG2pV4Lpcic5MQiN1fKkqVaj6A15MMgNTQ", // 火
		"4nkFyWhLrUAGZxr1Ku5NreywPPA6HEKkpqV2hDgr1kLU", // 国网
	}

	m, err := GetPowerPlantEnergy(slKey, "1497679290000", "1497682890000")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(m)
}

func Test_GetRolePublicKey(t *testing.T) {
	// 电表
	meterkeys, err := GetRolePublicKey(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(meterkeys)

	// 发电厂
	var powerplantskeys []string
	type_ := []int{2, 3, 4, 5, 6}
	for _, e := range type_ {
		keys, err := GetRolePublicKey(e)
		if err != nil {
			t.Error(err)
			return
		}
		powerplantskeys = append(powerplantskeys, keys...)
	}

	t.Log(powerplantskeys)
}

func Test_GetTransactionRecords(t *testing.T) {
	str, err := GetTransactionRecords()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(str)
}

/*智能微网demo end---------------------------------------------------------*/

func Test_InsertInterestCount(t *testing.T) {
	str := `{"username":"alice","pubkey":"key...","contractid":"conid","raiseStart":"date","raiseEnd":"date","raiseAmount":20000,"firstPurchaseAmount":20,"balanceAmount":21,"purchaseAmount":20,"date":"2","isRaise":false,"rate":0.02,"yield":2,"timestamp":123147}`
	InsertInterestCount(str)
}

func Test_GetInfoByUser(t *testing.T) {
	res, _ := GetInfoByUser("key...")
	uniledgerlog.Info(res)
}

func Test_GetLastInterest(t *testing.T) {
	res, _ := GetLastInterest("key...")
	uniledgerlog.Info(res)
}

func TestUpdateContractOutVote(t *testing.T) {
	vote := model.Vote{Id: "123456789"}
	voteMap, _ := common.StructToMap(vote)
	//5c3e37d71557e57fa7e3946fa5fc76b1ffacd35cf16e240ba74b14b14fa2ee89
	b := UpdateContractOutVote("5c3e37d71557e57fa7e3946fa5fc76b1ffacd35cf16e240ba74b14b14fa2ee89", voteMap, 1)
	uniledgerlog.Info(b)
}
