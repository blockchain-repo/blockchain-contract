package expressionutils

import (
	"testing"
)

func Test_IsSingleWord(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"",
		" ",
		"      ",
		"true",
		"abc",
		"1",
		"abc_1",
		" abc_1",
		"abc_1 ",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsSingleWord(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not SingleWord, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"abc||1 ",
		"abc==1 ",
		"abc>1 ",
		"+10",
		"-10",
		"10.23",
		"0.0001",
		"+10.23",
		"-0.0001",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsSingleWord(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is SingleWord, Check Error!", index, value)
		}
	}
}

func Test_IsExprNum(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"",
		" ",
		"100",
		"-100",
		"0",
		"+100",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsExprNum(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not Num, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"+100.00",
		"-100.23",
		".",
		"asdf",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsExprNum(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is Num, Check Error!", index, value)
		}
	}
}

func Test_IsExprFloat(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"",
		"-100.0",
		"-100.000001",
		"0.0",
		"10.20",
		"+100.0",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsExprFloat(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not Float, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"100.",
		"-100",
		"100",
		".04",
		"asdf",
		".",
		"+",
		"0",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsExprFloat(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is Float, Check Error!", index, value)
		}
	}
}

func Test_IsExprBool(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"",
		"true",
		"false",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsExprBool(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not Bool, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"TRUE",
		"FLASE",
		"True",
		"Flase",
		"asdf",
		"truetrue",
		"falsefalse",
		"truefalse",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsExprBool(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is Bool, Check Error!", index, value)
		}
	}
}

func Test_IsExprString(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"",
		"\"string\"",
		`'string'`,
		"`string`",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsExprString(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not String, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"string",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsExprString(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is String, Check Error!", index, value)
		}
	}
}

func Test_IsExprDate(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"",
		"2017-05-22 10:51:34",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsExprDate(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not Date, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"2017",
		"2017-05",
		"2017-05-22",
		"2017-05-22 10:51",
		"2017-",
		"2017-05-",
		"2017-05-22 ",
		"2017-05-22 10:51:9",
		"2017-5-22 10:51:09",
		"2017-05-22-10:51:09",
		"2017-35-22 34:51:34",
		"2017-05-32 34:51:34",
		"2017-02-30 34:51:34",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsExprDate(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is Date, Check Error!", index, value)
		}
	}
}

func Test_IsExprArray(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"",
		"[1,2]",
		"[1,2,a,4,1q1q1]",
		"[100,2,3,4,500]",
		"[100 ,2, 3,4,500]",
		"[a,b,c]", "[ a,b,c]", "[a ,b , c]", "[a, b,c ]", "[a, b, c ]",
		"[1,2,3,4.1]",
		"[_,_,_]",
		"[a.a.a,b.b,c.c.c.c]",

		"[.,.,.]", "[_.,._,.]", "[a.,b.,c.]", //TODO 此种正则处理不了类似这种
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsExprArray(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not Array, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"[]",
		"[,,,,1]",
		"[1 2 3 4,1]", "[1,2,3,4 1]", "[1,2,3,4|1]",
		"[a b c d,1]", "[a,x,e,3 1]", "[r,g,s,w|1]",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsExprArray(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is Array, Check Error!", index, value)
		}
	}
}

func Test_IsExprCondition(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"",
		"a&&b", "a&&b||c", "!a", "!a&&!b", "!a&&!b||!c",
		"a && b", "a&&b",
		"a || b", "a||b",
		"!a && !b", "!a &&!b", " !a &&!b", " ! a &&!b", "!a && ! b", " ! a && ! b ",
		"a>b", "a >b", "a> b",
		"a<b", "a <b", "a< b",
		"a==b", "a ==b", "a== b",
		"a>=b", "a >=b", "a>= b",
		"a<=b", "a <=b", "a<= b",
		"!a", " !a", "! a", "!a ", "!!!!!!!a",
		"a && b || c",

		"a. && b. || c.", ". && . || .", //TODO 此种正则处理不了类似这种
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsExprCondition(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not Condition, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"a&&", "a &&", " a &&", "a && ", " a && ",
		"a><=b",
		"a> =b",
		" !! !!!!!a",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsExprCondition(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is Condition, Check Error!", index, value)
		}
	}
}

func Test_IsExprFunction(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"",
		"Func1",
		"Funca",
		"Func_",
		"FuncA",
		"FuncFunc",
		`FuncIsConPutInUnichian("a90b93a2567a018afe52258f02c39c4de9b25e2e539b81778dbb897a3f88fc92")`,
		`FuncIsConPutInUnichian(a.b.c)`,
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsExprFunction(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not Function, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"Func",
		"func",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsExprFunction(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is Function, Check Error!", index, value)
		}
	}
}

func Test_IsExprVariable(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"",
		"asdf.asdf",
		"asdf.asdf.asdf",
		"_dddd",
		"_",
		"_22222",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsExprVariable(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not Variable, Check Error!", index, value)
		}
	}
	slTestErrorStr := []string{
		"asdf.",
		"_.",
		"__.",
		"3333",
		"3333.",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsExprVariable(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is Variable, Check Error!", index, value)
		}
	}
}

func Test_IsNameContract(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"contract_",
		"contract_sss",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsNameContract(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not NameContract, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"",
		"contract",
		"_contract_",
		"asdfasdfasdfcontract_",
		"contracteeeeeee",
		"contract_.",
		"contract_+",
		"contract_=",
		"contract__",
		"contract_ ",
		"contract_ kkkk.",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsNameContract(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is NameContract, Check Error!", index, value)
		}
	}
}

func Test_IsNameTaskEnquiry(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"task_enquiry_",
		"task_enquiry_sss",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsNameTaskEnquiry(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not NameTaskEnquiry, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"",
		"task_enquiry",
		"_task_enquiry_",
		"asdfasdfasdftask_enquiry_",
		"task_enquiryeeeeeee",
		"task_enquiry_.",
		"task_enquiry_+",
		"task_enquiry_=",
		"task_enquiry__",
		"task_enquiry_ ",
		"task_enquiry_ kkkk.",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsNameTaskEnquiry(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is NameTaskEnquiry, Check Error!", index, value)
		}
	}
}

func Test_IsNameTaskAction(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"task_action_",
		"task_action_sss",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsNameTaskAction(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not NameTaskAction, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"",
		"task_action",
		"_task_action_",
		"asdfasdfasdftask_action_",
		"task_actioneeeeeee",
		"task_action_.",
		"task_action_+",
		"task_action_=",
		"task_action__",
		"task_action_ ",
		"task_action_ kkkk.",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsNameTaskAction(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is NameTaskAction, Check Error!", index, value)
		}
	}
}

func Test_IsNameTaskDecision(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"task_decision_",
		"task_decision_sss",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsNameTaskDecision(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not NameTaskDecision, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"",
		"task_decision",
		"_task_decision_",
		"asdfasdfasdftask_decision_",
		"task_decisioneeeeeee",
		"task_decision_.",
		"task_decision_+",
		"task_decision_=",
		"task_decision__",
		"task_decision_ ",
		"task_decision_ kkkk.",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsNameTaskDecision(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is NameTaskDecision, Check Error!", index, value)
		}
	}
}

func Test_IsNameTaskPlan(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"task_plan_",
		"task_plan_sss",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsNameTaskPlan(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not NameTaskPlan, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"",
		"task_plan",
		"_task_plan_",
		"asdfasdfasdftask_plan_",
		"task_planeeeeeee",
		"task_plan_.",
		"task_plan_+",
		"task_plan_=",
		"task_plan__",
		"task_plan_ ",
		"task_plan_ kkkk.",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsNameTaskPlan(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is NameTaskPlan, Check Error!", index, value)
		}
	}
}

func Test_IsNameTaskCandidate(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"task_candidate_",
		"task_candidate_sss",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsNameTaskCandidate(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not NameTaskCandidate, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"",
		"task_candidate",
		"_task_candidate_",
		"asdfasdfasdftask_candidate_",
		"task_candidateeeeeeee",
		"task_candidate_.",
		"task_candidate_+",
		"task_candidate_=",
		"task_candidate__",
		"task_candidate_ ",
		"task_candidate_ kkkk.",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsNameTaskCandidate(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is NameTaskCandidate, Check Error!", index, value)
		}
	}
}

func Test_IsNameDataInt(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"data_intdata_",
		"data_intdata_sss",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsNameDataInt(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not NameDataInt, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"",
		"data_intdata",
		"_data_intdata_",
		"asdfasdfasdfdata_intdata_",
		"data_intdataeeeeeeee",
		"data_intdata_.",
		"data_intdata_+",
		"data_intdata_=",
		"data_intdata__",
		"data_intdata_ ",
		"data_intdata_ kkkk.",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsNameDataInt(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is NameDataInt, Check Error!", index, value)
		}
	}
}

func Test_IsNameDataUint(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"data_uintdata_",
		"data_uintdata_sss",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsNameDataUint(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not NameDataUint, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"",
		"data_uintdata",
		"_data_uintdata_",
		"asdfasdfasdfdata_uintdata_",
		"data_uintdataeeeeeeee",
		"data_uintdata_.",
		"data_uintdata_+",
		"data_uintdata_=",
		"data_uintdata__",
		"data_uintdata_ ",
		"data_uintdata_ kkkk.",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsNameDataUint(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is NameDataUint, Check Error!", index, value)
		}
	}
}

func Test_IsNameDataFloat(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"data_float_",
		"data_float_sss",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsNameDataFloat(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not NameDataFloat, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"",
		"data_float",
		"_data_float_",
		"asdfasdfasdfdata_float_",
		"data_floateeeeeeee",
		"data_float_.",
		"data_float_+",
		"data_float_=",
		"data_float__",
		"data_float_ ",
		"data_float_ kkkk.",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsNameDataFloat(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is NameDataFloat, Check Error!", index, value)
		}
	}
}

func Test_IsNameDataText(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"data_text_",
		"data_text_sss",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsNameDataText(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not NameDataText, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"",
		"data_text",
		"_data_text_",
		"asdfasdfasdfdata_text_",
		"data_texteeeeeeee",
		"data_text_.",
		"data_text_+",
		"data_text_=",
		"data_text__",
		"data_text_ ",
		"data_text_ kkkk.",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsNameDataText(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is NameDataText, Check Error!", index, value)
		}
	}
}

func Test_IsNameDataDate(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"data_date_",
		"data_date_sss",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsNameDataDate(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not NameDataDate, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"",
		"data_date",
		"_data_date_",
		"asdfasdfasdfdata_date_",
		"data_dateeeeeeeee",
		"data_date_.",
		"data_date_+",
		"data_date_=",
		"data_date__",
		"data_date_ ",
		"data_date_ kkkk.",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsNameDataDate(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is NameDataDate, Check Error!", index, value)
		}
	}
}

func Test_IsNameDataArray(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"data_array_",
		"data_array_sss",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsNameDataArray(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not NameDataArray, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"",
		"data_array",
		"_data_array_",
		"asdfasdfasdfdata_array_",
		"data_arrayeeeeeeee",
		"data_array_.",
		"data_array_+",
		"data_array_=",
		"data_array__",
		"data_array_ ",
		"data_array_ kkkk.",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsNameDataArray(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is NameDataArray, Check Error!", index, value)
		}
	}
}

func Test_IsNameDataMatrix(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"data_matrix_",
		"data_matrix_sss",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsNameDataMatrix(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not NameDataMatrix, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"",
		"data_matrix",
		"_data_matrix_",
		"asdfasdfasdfdata_matrix_",
		"data_matrixeeeeeeee",
		"data_matrix_.",
		"data_matrix_+",
		"data_matrix_=",
		"data_matrix__",
		"data_matrix_ ",
		"data_matrix_ kkkk.",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsNameDataMatrix(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is NameDataMatrix, Check Error!", index, value)
		}
	}
}

func Test_IsNameDataCompound(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"data_compound_",
		"data_compound_sss",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsNameDataCompound(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not NameDataCompound, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"",
		"data_compound",
		"_data_compound_",
		"asdfasdfasdfdata_compound_",
		"data_compoundeeeeeeee",
		"data_compound_.",
		"data_compound_+",
		"data_compound_=",
		"data_compound__",
		"data_compound_ ",
		"data_compound_ kkkk.",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsNameDataCompound(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is NameDataCompound, Check Error!", index, value)
		}
	}
}

func Test_IsNameDataOperateResult(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"data_operateresult_",
		"data_operateresult_sss",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsNameDataOperateResult(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not NameDataOperateResult, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"",
		"data_operateresult",
		"_data_operateresult_",
		"asdfasdfasdfdata_operateresult_",
		"data_operateresult_eeeeeeee",
		"data_operateresult_.",
		"data_operateresult_+",
		"data_operateresult_=",
		"data_operateresult__",
		"data_operateresult_ ",
		"data_operateresult_ kkkk.",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsNameDataOperateResult(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is NameDataOperateResult, Check Error!", index, value)
		}
	}
}

func Test_IsNameExprFunc(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"expression_function_",
		"expression_function_sss",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsNameExprFunc(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not NameExprFunc, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"",
		"expression_function",
		"_expression_function_",
		"asdfasdfasdfexpression_function_",
		"expression_function_eeeeeeee",
		"expression_function_.",
		"expression_function_+",
		"expression_function_=",
		"expression_function__",
		"expression_function_ ",
		"expression_function_ kkkk.",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsNameExprFunc(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is NameExprFunc, Check Error!", index, value)
		}
	}
}

func Test_IsNameExprArgu(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()

	slTestRightStr := []string{
		"expression_logicargument_",
		"expression_logicargument_sss",
	}
	for index, value := range slTestRightStr {
		if !v_express_parse.IsNameExprArgu(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is not NameExprArgu, Check Error!", index, value)
		}
	}

	slTestErrorStr := []string{
		"",
		"expression_logicargument",
		"_expression_logicargument_",
		"asdfasdfasdfexpression_logicargument_",
		"expression_logicargument_eeeeeeee",
		"expression_logicargument_.",
		"expression_logicargument_+",
		"expression_logicargument_=",
		"expression_logicargument__",
		"expression_logicargument_ ",
		"expression_logicargument_ kkkk.",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsNameExprArgu(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is NameExprArgu, Check Error!", index, value)
		}
	}
}
