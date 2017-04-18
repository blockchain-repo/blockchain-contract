package model

import (
	"fmt"
	"testing"
	"unicontract/src/common"
)

func Test_Votes(t *testing.T) {
	//create new obj
	votes := Votes{}
	votes.VotesWithoutId.NodePubkey = "2123123"
	votes.NodePubkey = "123123"
	fmt.Println(votes)
	result := common.SerializePretty(votes)
	fmt.Println(result)
	result2 := common.SerializePretty(votes.VotesWithoutId)
	fmt.Println(result2)

	//var str interface{}
	//str = "334"
	//a, ok := str.(string)
	//fmt.Println(ok)
	//fmt.Println(a)
}
