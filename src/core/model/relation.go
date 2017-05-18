package model

import (
	"sort"
	"unicontract/src/config"
)

// 合约&交易关系信息
type Relation struct {
	ContractId     string // 合约ID string（存储内容改为第一次创建合约，合约描述态的ID）
	ContractHashId string // 合约HashID(原ContractId中的内容) string
	TaskExecuteIdx int    // 合约任务执行索引   int
	TaskId         string
	Voters         []string
	Votes          []*Vote
}

func (r *Relation) GenerateRelation(contracHashId string, contractid string, taskid string, taskExecuteIdx int) {
	r.ContractHashId = contracHashId
	r.ContractId = contractid
	r.TaskId = taskid
	r.TaskExecuteIdx = taskExecuteIdx
	r.Voters = sort.StringSlice(config.GetAllPublicKey())
}
