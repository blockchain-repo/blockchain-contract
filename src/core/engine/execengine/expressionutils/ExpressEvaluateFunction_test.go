package expressionutils

import (
	"testing"
	"unicontract/src/core/engine/execengine/function"
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
	v_function := function.NewFunctionParseEngine()
	v_function.LoadFunctionsCommon()
	ep := NewExpressionParseEngine()
	ep.SetFunctionEngine(v_function)
	slTestString := []string{
		//"", // error
		"true",
		"false",
		"1==1&&2>1",
		"1==1&&2<1",
		"1 == 1 && 2 > 1",
		"(1 == 1) && (2 > 1)",
		"FuncTestMethod1()",
	}
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
	v_function := function.NewFunctionParseEngine()
	v_function.LoadFunctionsCommon()
	v_express_parse := NewExpressionParseEngine()
	v_express_parse.SetFunctionEngine(v_function)
	slTestString := []string{
		"FuncSleepTime(5)",
		"FuncGetNowDate()",
		"FuncTestMethod()",
		"FuncTestMethod(1)",
		"FuncTestMethod(1,2)",
	}
	for _, value := range slTestString {
		interface_, err := v_express_parse.EvaluateExpressionFunction(value)
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
	v_function := function.NewFunctionParseEngine()
	v_function.LoadFunctionsCommon()
	ep := NewExpressionParseEngine()
	ep.SetFunctionEngine(v_function)
	slTestString := map[string]string{
		"Expression_Constant":  "1111111",
		"Expression_Condition": "1 == 2 && 3 < 4",
		"Expression_Function":  "FuncTestMethod1()",
		"Expression_Variable":  "contract.contractbody", // TODO 需要外部整体测试，单元测试不可以
	}
	for key, value := range slTestString {
		interface_, err := ep.EvaluateExpressionValue(key, value)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("%T, %+v\n", interface_, interface_)
		}
	}
}
