package expressionutils

import "testing"

//Test Function:
//    IsSingleWord
//    IsExprNum
//    IsExprFloat
//    IsExprBool
//    IsExprString
//    IsExprDate
//    IsExprArray
//    IsExprCondition
//    IsExprFunction
//    IsExprVariable
//    IsNameContract
//    IsNameTaskEnquiry
//    IsNameTaskAction
//    IsNameTaskDecision
//    IsNameTaskPlan
//    IsNameTaskCandidate
//    IsNameDataInt
//    IsNameDataUint
//    IsNameDataFloat
//    IsNameDataText
//    IsNameDataDate
//    IsNameDataArray
//    IsNameDataMatrix
//    IsNameDataCompound
//    IsNameDataOperateResult
//    IsNameExprFunc
//    IsNameExprArgu

func TestIsSingleWord(t *testing.T) {
	var v_express_parse *ExpressionParseEngine = NewExpressionParseEngine()
	var v_test_str1 string = ""
	if !v_express_parse.IsSingleWord(v_test_str1) {
		t.Error("[" + v_test_str1 + "] is SingleWord, Check Error!")
	}
	var v_test_str2 string = "true"
	if !v_express_parse.IsSingleWord(v_test_str2) {
		t.Error("[" + v_test_str2 + "] is SingleWord, Check Error!")
	}

	var v_test_str3 string = "abc"
	if !v_express_parse.IsSingleWord(v_test_str3) {
		t.Error("[" + v_test_str3 + "] is SingleWord, Check Error!")
	}

	var v_test_str4 string = "1"
	if !v_express_parse.IsSingleWord(v_test_str4) {
		t.Error("[" + v_test_str4 + "] is SingleWord, Check Error!")
	}

	var v_test_str5 string = "abc_1"
	if !v_express_parse.IsSingleWord(v_test_str5) {
		t.Error("[" + v_test_str5 + "] is SingleWord, Check Error!")
	}

	var v_test_str6 string = " abc_1"
	if !v_express_parse.IsSingleWord(v_test_str6) {
		t.Error("[" + v_test_str6 + "] is not SingleWord, Check Error!")
	}

	var v_test_str7 string = "abc_1 "
	if !v_express_parse.IsSingleWord(v_test_str7) {
		t.Error("[" + v_test_str7 + "] is not SingleWord, Check Error!")
	}

	var v_test_str8 string = "abc 1 "
	if v_express_parse.IsSingleWord(v_test_str8) {
		t.Error("[" + v_test_str8 + "] is not SingleWord, Check Error!")
	}

	var v_test_str9 string = "abc||1 "
	if v_express_parse.IsSingleWord(v_test_str9) {
		t.Error("[" + v_test_str9 + "] is not SingleWord, Check Error!")
	}
	var v_test_str10 string = "abc==1 "
	if v_express_parse.IsSingleWord(v_test_str10) {
		t.Error("[" + v_test_str10 + "] is not SingleWord, Check Error!")
	}
	var v_test_str11 string = "abc>1 "
	if v_express_parse.IsSingleWord(v_test_str11) {
		t.Error("[" + v_test_str11 + "] is not SingleWord, Check Error!")
	}
	var v_test_str12 string = "+10"
	if v_express_parse.IsSingleWord(v_test_str12) {
		t.Error("[" + v_test_str12 + "] is not SingleWord, Check Error!")
	}
	var v_test_str13 string = "-10"
	if v_express_parse.IsSingleWord(v_test_str13) {
		t.Error("[" + v_test_str13 + "] is not SingleWord, Check Error!")
	}
	var v_test_str14 string = "10.23"
	if v_express_parse.IsSingleWord(v_test_str14) {
		t.Error("[" + v_test_str14 + "] is not SingleWord, Check Error!")
	}
	var v_test_str15 string = "0.0001"
	if v_express_parse.IsSingleWord(v_test_str15) {
		t.Error("[" + v_test_str15 + "] is not SingleWord, Check Error!")
	}
	var v_test_str16 string = "+10.23"
	if v_express_parse.IsSingleWord(v_test_str16) {
		t.Error("[" + v_test_str16 + "] is not SingleWord, Check Error!")
	}
	var v_test_str17 string = "-0.0001"
	if v_express_parse.IsSingleWord(v_test_str17) {
		t.Error("[" + v_test_str17 + "] is not SingleWord, Check Error!")
	}
}

func TestIsExprNum(t *testing.T) {
	var v_express_parse *ExpressionParseEngine = NewExpressionParseEngine()

	var v_test_num_1 string = "-100"
	if !v_express_parse.IsExprNum(v_test_num_1) {
		t.Error("[" + v_test_num_1 + "] is not Num, Check Error!")
	}
	var v_test_num_2 string = "0"
	if !v_express_parse.IsExprNum(v_test_num_2) {
		t.Error("[" + v_test_num_2 + "] is not Num, Check Error!")
	}
	var v_test_num_3 string = "100"
	if !v_express_parse.IsExprNum(v_test_num_3) {
		t.Error("[" + v_test_num_3 + "] is not Num, Check Error!")
	}
	var v_test_num_4 string = "+100"
	if !v_express_parse.IsExprNum(v_test_num_4) {
		t.Error("[" + v_test_num_4 + "] is not Num, Check Error!")
	}
	var v_test_num_5 string = "-100.02"
	if v_express_parse.IsExprNum(v_test_num_5) {
		t.Error("[" + v_test_num_5 + "] is not Num, Check Error!")
	}
	var v_test_num_6 string = "+100.00"
	if v_express_parse.IsExprNum(v_test_num_6) {
		t.Error("[" + v_test_num_6 + "] is not Num, Check Error!")
	}
}

func TestIsExprFloat(t *testing.T) {
	var v_express_parse *ExpressionParseEngine = NewExpressionParseEngine()
	var v_test_num_1 string = "-100.0"
	if !v_express_parse.IsExprFloat(v_test_num_1) {
		t.Error("[" + v_test_num_1 + "] is not Num, Check Error!")
	}
	var v_test_num_2 string = "-100.0001"
	if !v_express_parse.IsExprFloat(v_test_num_2) {
		t.Error("[" + v_test_num_2 + "] is not Num, Check Error!")
	}
	var v_test_num_3 string = "0.0"
	if !v_express_parse.IsExprFloat(v_test_num_3) {
		t.Error("[" + v_test_num_3 + "] is not Num, Check Error!")
	}
	var v_test_num_4 string = "10.20"
	if !v_express_parse.IsExprFloat(v_test_num_4) {
		t.Error("[" + v_test_num_4 + "] is not Num, Check Error!")
	}
	var v_test_num_5 string = "+100.00"
	if !v_express_parse.IsExprFloat(v_test_num_5) {
		t.Error("[" + v_test_num_5 + "] is not Num, Check Error!")
	}
	var v_test_num_6 string = "100"
	if v_express_parse.IsExprFloat(v_test_num_6) {
		t.Error("[" + v_test_num_6 + "] is not Num, Check Error!")
	}
	var v_test_num_7 string = "100."
	if v_express_parse.IsExprFloat(v_test_num_7) {
		t.Error("[" + v_test_num_7 + "] is not Num, Check Error!")
	}
	var v_test_num_8 string = ".04"
	if v_express_parse.IsExprFloat(v_test_num_8) {
		t.Error("[" + v_test_num_8 + "] is not Num, Check Error!")
	}
}
