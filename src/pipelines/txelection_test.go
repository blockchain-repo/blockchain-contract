package pipelines

import (
	"testing"
	"fmt"
	"strings"
)

func Test_pipStart(t *testing.T) {
	fmt.Printf("11")
	//starttxElection()
	str1 := "Go"
	str2 := "go"
	fmt.Println(strings.EqualFold(str1,str2))
	fmt.Print(str1==str2)
}
