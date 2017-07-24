package gRPCClient

import (
	"encoding/json"
	"fmt"
	"testing"

	"unicontract/src/core/engine"
)

func Test_Init(t *testing.T) {
	engine.Init()
	Init()
	t.Log(server)
	t.Log(port)
	t.Log(On)
}

func Test_FunctionRun(t *testing.T) {
	engine.Init()
	Init()
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
