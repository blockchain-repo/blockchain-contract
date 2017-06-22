package constdef

const (
	//值表达式 正则
	Regexp_True = iota
	Regexp_False
	Regexp_Nil
	//单个字符 正则
	Regexp_Signal
	//表达式类型 正则
	Regexp_Num
	Regexp_Float
	Regexp_Bool
	Regexp_String
	Regexp_Date
	Regexp_Array
	Regexp_Condition
	Regexp_Func
	Regexp_Name
	//变量名称正则
	Regexp_Name_Contract
	Regexp_Name_Task_Enquiry
	Regexp_Name_Task_Action
	Regexp_Name_Task_Decision
	Regexp_Name_Task_Plan
	Regexp_Name_Task_Candidate
	Regexp_Name_Data_Int
	Regexp_Name_Data_Uint
	Regexp_Name_Data_Float
	Regexp_Name_Data_Text
	Regexp_Name_Data_Date
	Regexp_Name_Data_Array
	Regexp_Name_Data_Matrix
	Regexp_Name_Data_Compound
	Regexp_Name_Data_OperateResult
	Regexp_Name_Expr_Func
	Regexp_Name_Expr_Argu
)

var ExpressionRegexp = map[int]string{
	Regexp_True:  "true",
	Regexp_False: "false",
	Regexp_Nil:   "nil",

	Regexp_Signal: `\w+`,

	Regexp_Num:       "[+-]*[0-9]+",
	Regexp_Float:     "[+-]*[0-9]+[.][0-9]+",
	Regexp_Bool:      "true|false",
	Regexp_String:    "[\"\\'\\`][a-zA-Z0-9_]+[\"\\'\\`]",
	Regexp_Date:      `([0-9]{4}-((0[1-9]|1[0-2])-(0[1-9]|1[0-9]|2[0-8])|(0[13-9]|1[0-2])-(29|30)|(0[13578]|1[02])-31)|([0-9]{2}(0[48]|[2468][048]|[13579][26])|(0[48]|[2468][048]|[13579][26])00)-02-29) ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]`,
	Regexp_Array:     `(\[)([ ]*[a-zA-Z0-9_.]+[ ]*,[ ]*)*([ ]*[a-zA-Z0-9_.]+[ ]*)(\])`,
	Regexp_Condition: `[ ]*[!]*[ ]*[0-9a-zA-z_.]+([ ]*(>|<|==|!=|>=|<=|&&|\|\|)[ ]*[!]*[ ]*[0-9a-zA-z_.]+[ ]*)*`,
	Regexp_Func:      `Func[a-zA-Z0-9_()",]+`,
	Regexp_Name:      `[_a-zA-Z][a-zA-Z0-9_]*(.[a-zA-Z0-9_]+)*`,

	Regexp_Name_Contract:           "contract_",
	Regexp_Name_Task_Enquiry:       "task_enquiry_",
	Regexp_Name_Task_Action:        "task_action_",
	Regexp_Name_Task_Decision:      "task_decision_",
	Regexp_Name_Task_Plan:          "task_plan_",
	Regexp_Name_Task_Candidate:     "task_candidate_",
	Regexp_Name_Data_Int:           "data_intdata_",
	Regexp_Name_Data_Uint:          "data_uintdata_",
	Regexp_Name_Data_Float:         "data_float_",
	Regexp_Name_Data_Text:          "data_text_",
	Regexp_Name_Data_Date:          "data_date_",
	Regexp_Name_Data_Array:         "data_array_",
	Regexp_Name_Data_Matrix:        "data_matrix_",
	Regexp_Name_Data_Compound:      "data_compound_",
	Regexp_Name_Data_OperateResult: "data_operateresult_",
	Regexp_Name_Expr_Func:          "expression_function_",
	Regexp_Name_Expr_Argu:          "expression_condition_",
}
