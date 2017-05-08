package constdef

import (
	"testing"
)

func TestExpressionType(t *testing.T){
	var ctype0 int = Expression_Unknown
	if ctype0 != 0 || ExpressionType[Expression_Unknown] != "Expression_Unknown"{
		t.Error("Expression_Unknown value Error!")
	}
	var ctype1 int = Expression_Expression
	if ctype1 != 1 || ExpressionType[Expression_Expression] != "Expression_Expression"{
		t.Error("Expression_Expression value Error!")
	}
	var ctype2 int = Expression_Function
	if ctype2 != 2 || ExpressionType[Expression_Function] != "Expression_Function"{
		t.Error("Expression_Function value Error!")
	}
	var ctype3 int = Expression_LogicArgument
	if ctype3 != 3 || ExpressionType[Expression_LogicArgument] != "Expression_LogicArgument"{
		t.Error("Expression_LogicArgument value Error!")
	}
}