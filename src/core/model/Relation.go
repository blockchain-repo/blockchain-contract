package model

import (
	"unicontract/src/config"
)

// 合约&交易关系信息
type Relation struct {
	ContractId string
	TaskId     string
	Voters     []string
	Votes      []*Vote
}


func (r *Relation)GenerateRelation(contractid string,taskid string){
	r.ContractId = contractid
	r.TaskId = taskid
	r.Voters = config.GetAllPublicKey()
}