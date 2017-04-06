package models

//import "encoding/json"

type Head struct {
	Id          string   `json:"id"`          //根据合约Body计算的hash值
	Node_pubkey string   `json:"node_pubkey"` //合约集群中，运行该合约的节点公钥
	Main_pubkey string   `json:"main_pubkey"` //合约集群中，交易处理&共识处理主节点公钥
	Signature   string   `json:"signature"`   //合约集群中，运行该合约的节点签名
	Voters      []string `json:"voters"`      //合约集群中，对该合约进行共识投票的节点公钥环
	Timestamp   int64    `json:"timestamp"`
	Version     string   `json:"version"` //合约描述结构版本号
}

type ContractAttributes struct {
	//属性 Name, StartDate,EndDate,
	Name      string `json:"name"`
	StartDate int64  `json:"start_timestamp"`
	EndDate   int64  `json:"end_timestamp"`
}

type ContractSignature struct {
	Owner_pubkey string `json:"owner_pubkey"`
	Signature    string `json:"signature"`
	Timestamp    int64  `json:"timestamp"`
}

type ContractAssert struct {
	// 资产 Assert1{id,name,amount,metadata},Assert2,Assert3……
	Id       string                 `json:"id"`
	Name     string                 `json:"name"`
	Amount   float64                `json:"amount"`
	Metadata map[string]interface{} `json:"metadata"`
}

//contract_components
type PlanTaskCondition struct {
	Id            string `json:"id"`
	ConditionType string `json:"type"`
	Name          string `json:"name"`
	Value         string `json:"value"`
	Description   string `json:"description"`
}

type Plan struct {
	Id           string              `json:"id"`
	PlanType     string              `json:"type"`
	State        string              `json:"state"`
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	Condition    []PlanTaskCondition `json:"condition"`
	Level        int16               `json:"level"`
	ContractType string              `json:"contract_type"`
	NextTask     []string            `json:"next_task"` //plan id

}

type Task struct {
	Id           string              `json:"id"`
	PlanType     string              `json:"type"`
	State        string              `json:"state"`
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	Condition    []PlanTaskCondition `json:"condition"`
	Level        int16               `json:"level"`
	ContractType string              `json:"contract_type"`
	NextTask     []string            `json:"next_task"` //task id

}

type ContractComponents struct {
	Plans []Plan `json:"plans"`
	Tasks []Task `json:"tasks"`
}

type Other struct {
	Creator_pubkey   string `json:"creator_pubkey"`
	Create_timestamp int64  `json:"create_timestamp"`
	Operation        string `json:"operation"`
}

type Body struct {
	Other

	ContractAttributes `json:"contract_attributes"`

	//主体 Owner1,Owner2,Owner3……
	ContractOwners []string `json:"contract_owners"`
	//ContractSignature array
	ContractSignatures []ContractSignature `json:"contract_signatures"`

	//资产 Assert1{id,name,amount,metadata},Assert2,Assert3……
	ContractAsserts    []ContractAssert `json:"contract_asserts"`
	ContractComponents `json:"contract_components"`
}

type Contract struct {
	Head
	Body `json:"contract"`
}
