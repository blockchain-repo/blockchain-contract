package expression

import (
	"testing"
	"unicontract/src/core/engine/execengine/constdef"
)

func TestFunction(t *testing.T) {
	g_function := &Function{}
	g_function.InitFunction()

	g_function.SetCname("Function")
	g_function.SetCaption("Function")
	g_function.SetDescription("Function Test")
	g_function.SetExpressionStr("InitFunction")

	if g_function.GetCname() != "Function" {
		t.Error("InitFunction Error,GetCname Error!")
	}
	if g_function.GetCtype() != constdef.ComponentType[constdef.Component_Expression] + "." + constdef.ExpressionType[constdef.Expression_Function] {
		t.Error("InitFunction Error, GetCtype Error!")
	}
	if g_function.GetCaption() != "Function" {
		t.Error("InitFunction Error,GetCaption Error!")
	}
	if g_function.GetDescription() != "Function Test" {
		t.Error("InitFunction Error,GetDescription Error!")
	}

	if g_function.GetExpressionStr() != "InitFunction" {
		t.Error("InitFunction Error,GetExpressionStr Error!")
	}
	//Test ToString
	if g_function.ToString() != "Function" {
		t.Error("InitFunction Error,ToString Error!")
	}
}