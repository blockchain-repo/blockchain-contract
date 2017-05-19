package constdef

import (
	"testing"
)

func TestExpressionType(t *testing.T){
	var ctype0 int = Expression_Unknown
	if ctype0 != 0 || ExpressionType[Expression_Unknown] != "Expression_Unknown"{
		t.Error("Expression_Unknown value Error!")
	}
	var ctype1 int = Expression_Constant
	if ctype1 != 1 || ExpressionType[Expression_Constant] != "Expression_Constant"{
		t.Error("Expression_Constant value Error!")
	}
	var ctype2 int = Expression_Variable
	if ctype2 != 2 || ExpressionType[Expression_Variable] != "Expression_Variable"{
		t.Error("Expression_Variable value Error!")
	}
	var ctype3 int = Expression_Condition
	if ctype3 != 3 || ExpressionType[Expression_Condition] != "Expression_Condition"{
		t.Error("Expression_Condition value Error!")
	}
	var ctype4 int = Expression_Function
	if ctype4 != 4 || ExpressionType[Expression_Function] != "Expression_Function"{
		t.Error("Expression_Function value Error!")
	}
	var ctype5 int = Expression_Candidate
	if ctype5 != 5 || ExpressionType[Expression_Candidate] != "Expression_Candidate"{
		t.Error("Expression_Candidate value Error!")
	}
}