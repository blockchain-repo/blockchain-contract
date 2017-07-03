package expressionutils

import (
	"testing"
)

func Test_EvaluateExpressionConstant(t *testing.T) {
	slTestString := []string{
		//"",                    // error
		"123",                 // int64
		"123.12",              // float64
		"true",                // bool
		"True",                // default->string
		"false",               // bool
		"False",               // default->string
		`"hello world"`,       // string
		`'hello world'`,       // string
		"`hello world`",       // string
		"hello world",         // default->string
		"hello world\\",       // default->string
		"2017-05-21 10:10:10", // date->string
		"2017-5-21 10:10:10",  // default->string
		"2017-5-21 10:5:10",   // default->string
		"[1,2,3,4,5,6]",       // array
		"[1, 2 ,3,4  ,5, 6]",  // array
		"[aa,2a,3a,d,5,6]",    // array
	}
	ep := NewExpressionParseEngine()
	for _, value := range slTestString {
		interface_, err := ep.EvaluateExpressionConstant(value)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("%T, %+v\n", interface_, interface_)
		}
	}
}

func Test_EvaluateExpressionVariable(t *testing.T) {
	// TODO 需要外部整体测试，单元测试不可以

	slTestString := []string{
		"", // error
	}

	ep := NewExpressionParseEngine()
	for _, value := range slTestString {
		interface_, err := ep.EvaluateExpressionVariable(value)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("%T, %+v\n", interface_, interface_)
		}
	}
}

func Test_EvaluateExpressionCondition(t *testing.T) {
	// TODO 需要外部整体测试，单元测试不可以

	slTestString := []string{
		//"", // error
		"true",
		"false",
	}

	ep := NewExpressionParseEngine()
	for _, value := range slTestString {
		interface_, err := ep.EvaluateExpressionCondition(value)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("%T, %+v\n", interface_, interface_)
		}
	}
}

func Test_EvaluateExpressionFunction(t *testing.T) {
	// TODO 需要外部整体测试，单元测试不可以

	slTestString := []string{
		"",
	}

	ep := NewExpressionParseEngine()
	for _, value := range slTestString {
		interface_, err := ep.EvaluateExpressionFunction(value)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("%T, %+v\n", interface_, interface_)
		}
	}
}

func Test_EvaluateExpressionCandidate(t *testing.T) {
	slTestString := []string{
		"",
	}

	ep := NewExpressionParseEngine()
	for _, value := range slTestString {
		interface_, err := ep.EvaluateExpressionCandidate(value)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("%T, %+v\n", interface_, interface_)
		}
	}
}

func Test_EvaluateExpressionValue(t *testing.T) {
	// TODO 需要外部整体测试，单元测试不可以

	slTestString := []string{
		//`FuncIsConPutInUnichian("a90b93a2567a018afe52258f02c39c4de9b25e2e539b81778dbb897a3f88fc92")`,
		`FuncGetNowDate() == FuncGetNowDate()`,
	}

	ep := NewExpressionParseEngine()
	for _, value := range slTestString {
		interface_, err := ep.EvaluateExpressionValue("Expression_Condition", value)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("%T, %+v\n", interface_, interface_)
		}
	}
}
