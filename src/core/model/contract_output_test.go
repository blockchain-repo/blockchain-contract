package model

import (
	"fmt"
	"testing"
	"unicontract/src/common"
)

func Test_ContractOutput(t *testing.T) {
	contractOutput := ContractOutput{}
	contractOutput.Id = "1"
	result := common.Serialize(contractOutput)
	fmt.Println(result)
}
