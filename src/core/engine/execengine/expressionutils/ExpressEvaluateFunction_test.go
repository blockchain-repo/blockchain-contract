package expressionutils

import (
	"testing"
)

func Test_EvaluateExpressionConstant(t *testing.T) {
	ep := NewExpressionParseEngine()
	slTestString := []string{
		"",
		"123",
		"123.12",
		"true",
		"false",
		`"hello world"`,
		"2017-05-21 10:10:10",
		"", // TODO array ?
	}
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
	slTestString := []string{
		"", // TODO ?
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
	slTestString := []string{
		"",
		"true",
		"false",
		"0",
		"1",
		"Func1", // TODO ?
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

func Test_EvaluateExpressionFunction(t *testing.T) { // TODO ?
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

func Test_EvaluateExpressionCandidate(t *testing.T) { // TODO ?
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

func Test_EvaluateExpressionValue(t *testing.T) { // TODO ?
	slTestString := []string{
		`FuncIsConPutInUnichian("a90b93a2567a018afe52258f02c39c4de9b25e2e539b81778dbb897a3f88fc92")`,
	}

	ep := NewExpressionParseEngine()
	for _, value := range slTestString {
		interface_, err := ep.EvaluateExpressionValue("Expression_Function", value)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("%T, %+v\n", interface_, interface_)
		}
	}
}
