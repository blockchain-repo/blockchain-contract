package model

import (
	"fmt"
	"testing"
	"unicontract/src/common"
	"unicontract/src/config"
)

func Test_Votes(t *testing.T) {
	//create new obj
	vote := Vote{}
	vote.Id = common.GenerateUUID()
	vote.NodePubkey = config.Config.Keypair.PublicKey
	vote.VoteBody.Timestamp = common.GenTimestamp()
	vote.Signature = vote.SignVote()
	fmt.Println(vote)
	result := common.StructSerializePretty(vote)
	fmt.Println(result)
}
