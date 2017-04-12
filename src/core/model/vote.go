package model

type Vote struct {
	IsValid         bool   `json:"is_valid"`          //合约、合约交易投票结果，如true,false
	InvalidReason   string `json:"invalid_reason"`    //合约、合约交易投无效票原因
	VoteType        string `json:"vote_type"`         //投票对象的类型，如CONTRACT，TRANSACTION等
	VoteForContract string `json:"vote_for_contract"` //投票的合约、合约交易ID
}

// table [vote]
type Votes struct {
	Id         string `json:"id"`          //投票唯一标识ID，最投票主体信息计算hash
	NodePubkey string `json:"node_pubkey"` //投票节点的公钥
	Timestamp  string `json:"timestamp"`   //节点投票时间戳
	Vote       Vote   `json:"vote"`        //投票信息
}


func (v *Votes) getVote() interface{} {
	return nil
}
