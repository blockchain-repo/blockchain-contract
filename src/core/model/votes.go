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

type VotesWithoutId struct{
	NodePubkey string `json:"node_pubkey"` //投票节点的公钥
	Vote       Vote   `json:"vote"`        //投票信息
	Signature  string `json:"signature"`   //投票节点签名
}

// table [vote]
type Votes struct {
	Id         string `json:"id"`          //投票唯一标识ID，最投票主体信息计算hash
	VotesWithoutId
}

// Calculate the election status of a contract.
func (v *Votes) ContractElection() {

}

//  Filter votes from unknown nodes or nodes that are not listed on
//	block. This is the primary Sybill protection.
func (v *Votes) PartitionEligibleVotes(votes []Votes, eligible_voters []string) ([]Votes, []Votes) {
	//eligible := make([]Votes, len(votes))
	//ineligible := make([]Votes, len(votes))
	var eligible []Votes
	var ineligible []Votes
	for _, _votes := range votes {
		voter_eligible := make([]string, len(votes))
		for _, _voter_eligible := range eligible_voters {
			if _votes.NodePubkey == _voter_eligible {
				voter_eligible = append(voter_eligible[:], _voter_eligible)
			}
		}
		if voter_eligible != nil {
			if v.VerifyVoteSignature(_votes) {
				eligible = append(eligible[:], _votes)
				continue
			}
		}
		ineligible = append(ineligible[:], _votes)
	}
	return eligible, ineligible
}

func (v *Votes) CountVotes(eligible_votes []Votes) map[string]interface{} {

	// Group by pubkey to detect duplicate voting
	by_voter := make(map[Votes]bool)
	for _, votes := range eligible_votes {
		if !by_voter[votes] {
			by_voter[votes] = true
		}
	}
	n_valid := 0
	cheat := make([]Votes, len(by_voter))
	for votes, _ := range by_voter {
		cheat = append(cheat, votes)
		vote := votes.Vote
		if vote.IsValid {
			n_valid += 1
		}
	}

	resultMap := make(map[string]interface{})
	counts := make(map[string]int)
	counts["n_valid"] = n_valid
	counts["n_invalid"] = len(by_voter) - n_valid
	resultMap["cheat"] = cheat
	resultMap["counts"] = counts

	return resultMap
}

// TODO Decide on votes.
// TODO To return VALID there must be a clear majority that say VALID.
// TODO A tie on an even number of votes counts as INVALID.
func (v *Votes) DecideVotes(n_voters int, n_valid int, n_invalid int) string {
	if n_invalid*2 >= n_voters {
		return "INVALID"
	}

	if n_valid*2 > n_voters {
		return "VALID"
	}

	return "UNDECIDED"
}

// TODO Verify the signature of a vote
func (v *Votes) VerifyVoteSignature(vote Votes) bool {
	signature := vote.Signature
	pk_base58 := vote.NodePubkey
	body := vote.ToString()
	public_key := common.Sign(pk_base58, body)
	return common.Verify(public_key, body, signature)
}

func (v *Votes) VerifyVoteSchema() bool {
	return true
}

func (v *Votes) ToString() string {
	return common.Serialize(v)
}
