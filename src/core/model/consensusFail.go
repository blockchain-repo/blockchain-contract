package model

// table [consensusFail]
type ConsensusFail struct {
	Id              string `json:"id"`               //共识失败原因记录唯一标识ID
	ConsensusType   string `json:"consensus_type"`   //共识类型，如contract, transaction
	ConsensusId     string `json:"consensus_id"`     //共识对象ID，如contract_id  transaction_id
	ConsensusReason string `json:"consensus_reason"` //共识失败原因记录
	Timestamp       string `json:"timestamp"`        //共识时间戳
}
