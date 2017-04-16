package rethinkdb

import (
	"fmt"
	"testing"
)

func Test_Reconfig(t *testing.T) {
	res :=Reconfig("bigchain",5,1)
	fmt.Printf("%s\n",res)
}
