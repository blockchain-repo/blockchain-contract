package expressionutils

import (
	"testing"
)

func Test_ParseExpressionClassify(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()
	slTestRightStr := []string{
		"",
		"122",
		"+122",
		"-122",
		"122.0",
		"122.33",
		"+122.33",
		"-122.33",
		"true",
		"false",
		"`string`",
		`"string"`,
		`'string'`,
		"2017-05-22 10:10:10",
		"", // TODO 6.纯数组串
		"", // TODO 7.条件表达式
		"", // TODO 8.函数表达式
		"", // TODO 9.变量表达式
	}
	for index, value := range slTestRightStr {
		strRet := v_express_parse.ParseExpressionClassify(value)
		t.Logf("[ %d ] ParseExpressionClassify ret string is %s", index, strRet)
	}
}

func Test_ParseExprNumValue(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()
	slTestRightStr := []string{
		"122",
		"+122",
		"-122",
	}
	for _, value := range slTestRightStr {
		nRet, err := v_express_parse.ParseExprNumValue(value)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("ParseExprNumValue ret is %d", nRet)
		}
	}
}

func Test_ParseExprFloatValue(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()
	slTestRightStr := []string{
		"122.0",
		"+122.0",
		"-122.0",
		"122",
		"+122",
		"-122",
	}
	for _, value := range slTestRightStr {
		nRet, err := v_express_parse.ParseExprFloatValue(value)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("ParseExprFloatValue ret is %f", nRet)
		}
	}
}

func Test_ParseExprBoolValue(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()
	slTestRightStr := []string{
		"true",
		"false",
	}
	for _, value := range slTestRightStr {
		b, err := v_express_parse.ParseExprBoolValue(value)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("ParseExprBoolValue ret is %t", b)
		}
	}
}

func Test_ParseExprStringValue(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()
	slTestRightStr := []string{
		`"string"`,
		`'string'`,
		"`string`",
	}
	for _, value := range slTestRightStr {
		strRet, err := v_express_parse.ParseExprStringValue(value)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("ParseExprStringValue ret is %s", strRet)
		}
	}
}

func Test_ParseExprDateValue(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()
	slTestRightStr := []string{
		"2017-05-22 10:10:10",
	}
	for _, value := range slTestRightStr {
		strRet, err := v_express_parse.ParseExprDateValue(value)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("ParseExprDateValue ret is %s", strRet)
		}
	}
}

func Test_ParseExprArrayValue(t *testing.T) { //TODO ?
	v_express_parse := NewExpressionParseEngine()
	slTestRightStr := []string{
		"",
	}
	for _, value := range slTestRightStr {
		ret, err := v_express_parse.ParseExprArrayValue(value)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("ParseExprArrayValue ret is %v", ret)
		}
	}
}

func Test_ParseExprConditionValue(t *testing.T) { //TODO ?
	v_express_parse := NewExpressionParseEngine()
	slTestRightStr := []string{
		"FuncGetNowDate() == '2017-06-23 17:56:00'",
	}
	for _, value := range slTestRightStr {
		b, err := v_express_parse.ParseExprConditionValue(value)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("ParseExprConditionValue ret is %t", b)
		}
	}
}

func Test_ParseExprFunctionValue(t *testing.T) { //TODO ?
	v_express_parse := NewExpressionParseEngine()
	slTestRightStr := []string{
		"",
	}
	for _, value := range slTestRightStr {
		ret, err := v_express_parse.ParseExprFunctionValue(value)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("ParseExprFunctionValue ret is %v", ret)
		}
	}
}

func Test_ParseExprVariableValue(t *testing.T) { //TODO ?
	v_express_parse := NewExpressionParseEngine()
	slTestRightStr := []string{
		"",
	}
	for _, value := range slTestRightStr {
		ret, err := v_express_parse.ParseExprVariableValue(value)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("ParseExprVariableValue ret is %v", ret)
		}
	}
}

func Test_ParseVariablesInExprCondition(t *testing.T) { //TODO ?
	v_express_parse := NewExpressionParseEngine()
	slTestRightStr := []string{
		"",
	}
	for _, value := range slTestRightStr {
		ret, err := v_express_parse.ParseVariablesInExprCondition(value)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("ParseVariablesInExprCondition ret is --- %v", ret)
		}
	}
}
