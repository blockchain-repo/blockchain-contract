package model

import (
	"fmt"
	"testing"
	"unicontract/src/common"
)

func Test_Votes(t *testing.T) {
	//create new obj
	vote := Vote{}
	vote.NodePubkey = "2123123"
	vote.NodePubkey = "123123"
	fmt.Println(vote)
	result := common.SerializePretty(vote)
	fmt.Println(result)

	//var str interface{}
	//str = "334"
	//a, ok := str.(string)
	//fmt.Println(ok)
	//fmt.Println(a)
}
