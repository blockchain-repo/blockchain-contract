package models

import (
	"fmt"
	"testing"
)

func Test_Contract2(t *testing.T) {
	contract := Contract{}
	contract.Id = "2"
	contract.GetId()
	fmt.Println("contract id is ", contract)
	fmt.Println("contract id is ", contract)
}
