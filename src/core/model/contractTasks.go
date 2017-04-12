package model

// table [contractTasks]
type ContractTasks struct {
	Id                 string `json:"id"`                  //合约执行步骤唯一标识ID
	ContractId         string `json:"contract_id"`         //合约ID
	ContractStep       string `json:"contract_step"`       //合约执行步骤描述
	ContractCondiction string `json:"contract_condiction"` //合约步骤执行条件
	ContractState      string `json:"contract_state"`      //合约步骤的执行状态
}
