package expression

import (
	"testing"
	"unicontract/src/core/engine/execengine/constdef"
)

func TestInitGeneralExpression(t *testing.T){
	g_expression := NewGeneralExpression("1==1")
	err := g_expression.InitExpression()
	if err != nil {
		t.Error("InitExpression Error!")
	}
	g_expression.SetCname("Expression")
	g_expression.SetCaption("Expression")
	g_expression.SetDescription("Expression Test")
	g_expression.SetExpressionStr("a > b")

	if g_expression.GetCname() != "Expression" {
		t.Error("InitGeneralExpression Error, GetCname Error!")
	}
	if g_expression.GetCtype() != constdef.ComponentType[constdef.Component_Expression] {
		t.Error("InitGeneralExpression Error, GetCtype Error!")
	}
	if g_expression.GetCaption() != "Expression" {
		t.Error("InitGeneralExpression Error, GetCaption Error!")
	}
	if g_expression.GetDescription() != "Expression Test" {
		t.Error("InitGeneralExpression Error, GetDescription Error!")
	}
	if g_expression.GetExpressionStr() != "a > b" {
		t.Error("InitGeneralExpression Error, GetExpressionStr Error!")
	}

	//Test SetExpression
	g_expression.SetExpressionStr("1 + 1 = 2")
	if g_expression.GetExpressionStr() != "1 + 1 = 2" {
		t.Error("SetExpressionStr Error!")
	}

	//Test ToString
	if g_expression.ToString() != "Expression: 1 + 1 = 2" {
		t.Error("ToString Error!")
	}
}
