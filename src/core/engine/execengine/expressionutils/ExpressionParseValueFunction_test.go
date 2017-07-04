package expressionutils

import (
	"testing"
	"unicontract/src/core/engine/execengine/function"
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
		"[1,2,3,4,5,6]",
		"contract_transfer.ContractBody.ContractOwners.1 == 'true'",
		`FuncIsConPutInUnichian(a.b.c)`,
		"asdf.asdf.asdf",
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
		"True",
		"TRUE",
		"1",
		"T",
		"t",
		"false",
		"False",
		"FALSE",
		"0",
		"F",
		"f",
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
		`"str ing"`,
		`'str ing'`,
		"`str ing`",
		"string",
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

func Test_ParseExprArrayValue(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()
	slTestRightStr := []string{
		"[1,2,3,4,5,6]",
		"[1a,2,3,4w,5,6]",
		"[1 ,2, 3 ,4,  5,6  ]",
		"[asdf.ddd,2,wwww.djdjd.jfu,4,5,6]",
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

func Test_ParseExprConditionValue(t *testing.T) {
	v_function := function.NewFunctionParseEngine()
	v_function.LoadFunctionsCommon()
	v_express_parse := NewExpressionParseEngine()
	v_express_parse.SetFunctionEngine(v_function)

	slTestRightStr := []string{
		//`FuncGetNowDate() == "2017-07-03 05:26:49"`, // not Condition
		//`FuncGetNowDay() == "4"`, // not Condition
		` 4 == 4`,
		`"4" == "4"`,
		`"4" == "3"`,
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

func Test_ParseExprFunctionValue(t *testing.T) {
	v_function := function.NewFunctionParseEngine()
	v_function.LoadFunctionsCommon()
	v_express_parse := NewExpressionParseEngine()
	v_express_parse.SetFunctionEngine(v_function)
	slTestRightStr := []string{
		//"FuncSleepTime(5)",
		//"FuncGetNowDate()",
		//"FuncTestMethod()",
		//"FuncTestMethod(1)",
		//"FuncTestMethod(1,2)",
		"FuncTestMethod(asdf.asdf)", // TODO 需要外部整体测试，单元测试不可以
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

func Test_ParseExprVariableValue(t *testing.T) {
	// TODO 需要外部整体测试，单元测试不可以

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

func Test_ParseVariablesInExprCondition(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()
	slTestRightStr := []string{
		//`FuncGetNowDate() == "2017-07-03 05:26:49"`, // not Condition
		//`"FuncGetNowDate()" == "2017-07-03 05:26:49"`, // not Condition
		` 4 == 4`,
		`"4" == "4"`,
		`"4" == "3"`,
		`contract1.contractbody.id == "3"`,
		`contract1.contractbody.id == contract2.contractbody.id`,
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
