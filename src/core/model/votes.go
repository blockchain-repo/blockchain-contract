package model

import (
	//"github.com/astaxie/beego"
	"unicontract/src/common"
)

type Vote struct {
	IsValid         bool   `json:"is_valid"`          //合约、合约交易投票结果，如true,false
	InvalidReason   string `json:"invalid_reason"`    //合约、合约交易投无效票原因
	VoteForContract string `json:"vote_for_contract"` //投票的合约、合约交易ID
	VoteType        string `json:"vote_type"`         //投票对象的类型，如CONTRACT，TRANSACTION等
	Timestamp       string `json:"timestamp"`         //节点投票时间戳
}

// table [vote]
type Votes struct {
	Id         string `json:"id"`          //投票唯一标识ID，最投票主体信息计算hash
	NodePubkey string `json:"node_pubkey"` //投票节点的公钥
	Vote       Vote   `json:"vote"`        //投票信息
	Signature  string `json:"signature"`   //投票节点签名
}

func (v *Votes) ContractElection() {

}

func (v *Votes) PartitionEligibleVotes() interface{} {
	return nil
}

func (v *Votes) CountVotes() interface{} {
	return nil
}

func (v *Votes) DecideVotes() interface{} {
	return nil
}

//  Verify the signature of a vote
func (v *Votes) VerifyVoteSignature() bool {
	signature := v.Signature
	pk_base58 := v.NodePubkey
	body := v.ToString()
	public_key := common.Sign(pk_base58, body)
	return common.Verify(public_key, body, signature)
}

func (v *Votes) VerifyVoteSchema() bool {
	return true
}

func (c *Votes) ToString() string {
	return common.Serialize(c.Vote)
}
