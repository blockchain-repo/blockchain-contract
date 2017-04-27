package model

import (
	"unicontract/src/common"
)

// table [contractTasks]
type ContractTask struct {
	Id                 string  `json:"id"`//合约执行步骤唯一标识ID, uuid
	ContractId         string //合约ID
	ContractStep       string //合约执行步骤描述
	ContractCondiction string //合约步骤执行条件
	ContractState      string //合约步骤的执行状态
}

func (c *ContractTask) ToString() string {
	return common.StructSerialize(c)
}
