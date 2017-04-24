package model

// table [consensusFail]
type ConsensusFail struct {
	Id              string //共识失败原因记录唯一标识ID
	ConsensusType   string //共识类型，如contract, transaction
	ConsensusId     string //共识对象ID，如contract_id  transaction_id
	ConsensusReason string //共识失败原因记录
	Timestamp       string //共识时间戳
}
