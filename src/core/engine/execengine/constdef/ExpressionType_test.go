package constdef

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestExpressionType(t *testing.T) {
	var ctype0 int = Expression_Unknown
	if ctype0 != 0 || ExpressionType[Expression_Unknown] != "Expression_Unknown" {
		t.Error("Expression_Unknown value Error!")
	}
	var ctype1 int = Expression_Constant
	if ctype1 != 1 || ExpressionType[Expression_Constant] != "Expression_Constant" {
		t.Error("Expression_Constant value Error!")
	}
	var ctype2 int = Expression_Variable
	if ctype2 != 2 || ExpressionType[Expression_Variable] != "Expression_Variable" {
		t.Error("Expression_Variable value Error!")
	}
	var ctype3 int = Expression_Condition
	if ctype3 != 3 || ExpressionType[Expression_Condition] != "Expression_Condition" {
		t.Error("Expression_Condition value Error!")
	}
	var ctype4 int = Expression_Function
	if ctype4 != 4 || ExpressionType[Expression_Function] != "Expression_Function" {
		t.Error("Expression_Function value Error!")
	}
	var ctype5 int = Expression_Candidate
	if ctype5 != 5 || ExpressionType[Expression_Candidate] != "Expression_Candidate" {
		t.Error("Expression_Candidate value Error!")
	}
}

func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

func TestSplit(t *testing.T) {
	p_expression := "FuncGetNowDate() == 'aaaaaaaa' > bbb || cc && dd ! f != a > a < b >= aa <= bb + ee - ff * gg / hh % hh"
	reg := regexp.MustCompile(ExpressionTagString) //六位连续的数字
	fmt.Println("------FindAll------")
	fmt.Println(ExpressionTagString)
	dataSlice := reg.FindAllIndex([]byte(p_expression), -1)
	var int_begin int = 0
	var int_end int = 0
	var int_temp int = 0
	for idx, v := range dataSlice {
		fmt.Println(v)
		if idx == 0 {
			int_begin = 0
		} else {
			int_begin = int_temp
		}
		int_end = v[0]
		int_temp = v[1]
		fmt.Println("=====", int_begin, " ", int_end)
		fmt.Println(Substr2(p_expression, int_begin, int_end))
	}
	fmt.Println(Substr2(p_expression, int_temp, len(p_expression)))
}
