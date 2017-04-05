package models

import (
	"fmt"
	"testing"
	//"reflect"
	"bytes"
	"encoding/json"
	"github.com/apex/log"
	"time"
	//"os"
)

func Test_Contract(t *testing.T) {

	//var cc = new(Contract)
	//Contract.Head.Id = "2"
	var contract = new(Contract)
	contract.Head.Id = "2"
	contract.Head.Node_pubkey = "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc"
	contract.Head.Main_pubkey = "93TEovPuYo6BQFm4ia9ta4qtL1TbAmnk9fV5kxmesAG5"
	contract.Head.Signature = "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc"
	contract.Head.Voters = []string{"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet"}

	contract.Head.Version = "v1.0"


	contract.Body.Other.Creator_pubkey = "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc"
	contract.Body.Other.Create_timestamp = time.Now().Unix()
	contract.Body.Other.Operation = "CREATE"
	contract.Body.ContractAttributes.Name = "XXXXXX"
	contract.Body.ContractAttributes.StartDate = time.Now().Unix()
	contract.Body.ContractAttributes.EndDate = time.Now().Unix()
	contract.Body.ContractOwners = []string{"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
		"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet"}

	contractSignatures := make([]ContractSignature, 3)
	contractSignatures[0] = ContractSignature{Owner_pubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
	Signature:"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc", Timestamp:time.Now().Unix()}
	contractSignatures[1] = ContractSignature{Owner_pubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
	Signature:"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc", Timestamp:time.Now().Unix()}
	contractSignatures[2] = ContractSignature{Owner_pubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
	Signature:"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc", Timestamp:time.Now().Unix()}

	contract.Body.ContractSignatures = contractSignatures

	contractAsserts := make([]ContractAssert, 3)
	contractAsserts[0] = ContractAssert{Id: "111", Name: "wx1", Amount:1000}
	contractAsserts[1] = ContractAssert{Id: "113", Name: "wx2", Amount:800}
	contractAsserts[2] = ContractAssert{Id: "112", Name: "wx3", Amount:100}
	contract.Body.ContractAsserts = contractAsserts


	plans := make([]Plan, 2)

	planConditions := make([]PlanTaskCondition, 3)
	plans[0] = Plan{Id:"ID_Axxxxxx", PlanType:"PLAN", State:"dormant", Name:"N_Axxxxx",Description:"xxxxx"}

	planConditions[0] = PlanTaskCondition{Id:"1", ConditionType:"PreCondition", Name:"XXXX",Value:"XXXX",Description:"xxxxx"}
	planConditions[1] = PlanTaskCondition{Id:"2", ConditionType:"DisgardCondition", Name:"XXXX",Value:"XXXX",Description:"xxxxx"}
	planConditions[2] = PlanTaskCondition{Id:"3", ConditionType:"CompleteCondition", Name:"XXXX",Value:"XXXX",Description:"xxxxx"}
	plans[0].Condition = planConditions
	plans[0].Level = 1
	plans[0].ContractType ="RIGHT"
	plans[0].NextTask = []string{"1","2"}

	planConditions2 := make([]PlanTaskCondition, 3)
	plans[1] = Plan{Id:"ID_Axxxxxx", PlanType:"PLAN", State:"dormant", Name:"N_Axxxxx",Description:"xxxxx"}

	planConditions2[0] = PlanTaskCondition{Id:"1", ConditionType:"PreCondition", Name:"XXXX",Value:"XXXX",Description:"xxxxx"}
	planConditions2[1] = PlanTaskCondition{Id:"2", ConditionType:"DisgardCondition", Name:"XXXX",Value:"XXXX",Description:"xxxxx"}
	planConditions2[2] = PlanTaskCondition{Id:"3", ConditionType:"CompleteCondition", Name:"XXXX",Value:"XXXX",Description:"xxxxx"}
	plans[1].Condition = planConditions2
	plans[1].Level = 1
	plans[1].ContractType ="RIGHT"
	plans[1].NextTask = []string{"1","2"}



	tasks := make([]Task, 3)

	tasks[0] = Task{Id:"ID_Cxxxxxx", PlanType:"ENQUIRY", State:"dormant", Name:"Axxxxxx",Description:"xxxxx"}
	taskConditions := make([]PlanTaskCondition, 3)

	taskConditions[0] = PlanTaskCondition{Id:"1", ConditionType:"PreCondition", Name:"XXXX",Value:"XXXX",Description:"xxxxx"}
	taskConditions[1] = PlanTaskCondition{Id:"2", ConditionType:"DisgardCondition", Name:"XXXX",Value:"XXXX",Description:"xxxxx"}
	taskConditions[2] = PlanTaskCondition{Id:"3", ConditionType:"CompleteCondition", Name:"XXXX",Value:"XXXX",Description:"xxxxx"}
	tasks[0].Condition = taskConditions
	tasks[0].Level = 1
	tasks[0].ContractType ="RIGHT"
	tasks[0].NextTask = []string{"Axxxxxx","Bxxxxxx"}

	tasks[1] = Task{Id:"ID_Cxxxxxx", PlanType:"ENQUIRY", State:"dormant", Name:"Bxxxxxx",Description:"xxxxx"}
	taskConditions2 := make([]PlanTaskCondition, 3)

	taskConditions2[0] = PlanTaskCondition{Id:"1", ConditionType:"PreCondition", Name:"XXXX",Value:"XXXX",Description:"xxxxx"}
	taskConditions2[1] = PlanTaskCondition{Id:"2", ConditionType:"DisgardCondition", Name:"XXXX",Value:"XXXX",Description:"xxxxx"}
	taskConditions2[2] = PlanTaskCondition{Id:"3", ConditionType:"CompleteCondition", Name:"XXXX",Value:"XXXX",Description:"xxxxx"}
	tasks[1].Condition = taskConditions2
	tasks[1].Level = 2
	tasks[1].ContractType ="RIGHT"
	tasks[1].NextTask = []string{"",""}

	tasks[2] = Task{Id:"ID_Cxxxxxx", PlanType:"ACTION", State:"dormant", Name:"Cxxxxxx",Description:"xxxxx"}
	taskConditions3 := make([]PlanTaskCondition, 3)

	taskConditions3[0] = PlanTaskCondition{Id:"1", ConditionType:"PreCondition", Name:"XXXX",Value:"XXXX",Description:"xxxxx"}
	taskConditions3[1] = PlanTaskCondition{Id:"2", ConditionType:"DisgardCondition", Name:"XXXX",Value:"XXXX",Description:"xxxxx"}
	taskConditions3[2] = PlanTaskCondition{Id:"3", ConditionType:"CompleteCondition", Name:"XXXX",Value:"XXXX",Description:"xxxxx"}
	tasks[2].Condition = taskConditions3
	tasks[2].Level = 2
	tasks[2].ContractType ="DUTY"
	tasks[2].NextTask = []string{"",""}


	contract.Body.ContractComponents.Plans = plans
	contract.Body.ContractComponents.Tasks = tasks

	//fmt.Println(ToJson(contract))
	fmt.Println(ToJsonFormat(contract))

	//a := map[string]string{"a":"123","name":"123123"}
	//fmt.Println(ToJson(a))
	//fmt.Println(ToJsonFormat(a))
	//fmt.Println(Deserialize(ToJson(a)))
	//fmt.Println(Deserialize(ToJsonFormat(a)))
	//
	//b := [5]int{1,2,3}
	//fmt.Println(ToJson(b))
	//fmt.Println(ToJsonFormat(b))
	//fmt.Println(Deserialize(ToJson(b)))
	//fmt.Println(Deserialize(ToJsonFormat(b)))
	//
	//c := "asddfad"
	//fmt.Println(ToJson(c))
	//fmt.Println(ToJsonFormat(c))
	//fmt.Println(Deserialize(ToJson(c)))
	//fmt.Println(Deserialize(ToJsonFormat(c)))
	//
	//d := `json:{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`
	//fmt.Println(ToJson(d))
	//fmt.Println(ToJsonFormat(d))
	//fmt.Println(Deserialize(ToJson(d)))
	//fmt.Println(Deserialize(ToJsonFormat(d)))

}

func ToJsonFormat(obj interface{}) string {
	fmt.Println("== struct convert to pretty json str ==")
	input, err := json.Marshal(obj)
	if err != nil {
		log.Error(err.Error())
	}

	var out bytes.Buffer
	err = json.Indent(&out, input, "", "\t")

	if err != nil {
		log.Error(err.Error())
	}
	//out.WriteTo(os.Stdout)
	//fmt.Println(string(out.String()))
	return string(out.String())
}

func ToJson(obj interface{}) string {
	fmt.Println("== struct convert to normal json str ==")
	str, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(str))
	return string(str)
}

func Deserialize(str string) interface{} {

	fmt.Println("== json str to obj ==")
	var obj interface{}
	err := json.Unmarshal([]byte(str), &obj)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)
	return obj
}
