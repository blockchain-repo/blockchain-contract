package gRPCClient

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_FunctionRun(t *testing.T) {
	params := make(map[string]interface{})
	count := 1
	params[fmt.Sprintf("Param%02d", count)] = 100
	count++
	params[fmt.Sprintf("Param%02d", count)] = 200
	slData, err := json.Marshal(params)
	if err != nil {
		t.Error(err)
		return
	}
	result, err := FunctionRun("123", "FuncAdd", string(slData))
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v", result)
	}
}
