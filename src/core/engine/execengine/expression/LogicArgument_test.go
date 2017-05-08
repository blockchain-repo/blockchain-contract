package expression

import (
	"testing"
	"unicontract/src/core/engine/execengine/constdef"
)

func TestLogicArgument(t *testing.T) {
	g_logicargument := LogicArgument{}
	g_logicargument.InitLogicArgument()

	g_logicargument.SetCname("LogicArgument")
	g_logicargument.SetCaption("LogicArgument")
	g_logicargument.SetDescription("LogicArgument Test")
	g_logicargument.SetExpressionStr( "a > b")

	if g_logicargument.GetCname() != "LogicArgument" {
		t.Error("InitLogicArgument Error, GetCname Error!")
	}
	if g_logicargument.GetCtype() != constdef.ComponentType[constdef.Component_Expression] + "." + constdef.ExpressionType[constdef.Expression_LogicArgument] {
		t.Error("InitLogicArgument Error, GetCtype Error!")
	}
	if g_logicargument.GetCaption() != "LogicArgument" {
		t.Error("InitLogicArgument Error, GetCaption Error!")
	}
	if g_logicargument.GetDescription() != "LogicArgument Test" {
		t.Error("InitLogicArgument Error, GetDescription Error!")
	}
	if g_logicargument.GetExpressionStr() != "a > b" {
		t.Error("InitLogicArgument Error, GetExpressionStr Error!")
	}
}