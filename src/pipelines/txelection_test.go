package pipelines

import (
	"testing"
	"fmt"
)

func Test_pipStart(t *testing.T) {
	//starttxElection()
	//str1 := "Go"
	//str2 := "go"
	//fmt.Println(strings.EqualFold(str1,str2))
	//fmt.Print(str1==str2)
	validSignCount := 5
	voters_len := 9
	fmt.Print(voters_len/2 )
	if validSignCount <= voters_len/2 {
		fmt.Print("1")
	}
	fmt.Print("2")

}
