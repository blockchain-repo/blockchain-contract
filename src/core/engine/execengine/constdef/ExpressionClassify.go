package constdef

//---------------------------------------------------------------------------------------------
//表达式分类：
//    4大类：  1. 常量表达式  2. 变量表达式  3.条件表达式   4.函数表达式   5.决策表达式
//       1.常量表达式：1）.纯数字  2）.纯浮点数  3）.纯bool值  4）.纯字符串  5）.纯日期  6.纯数组串
//       2.变量表达式：component_table.property_table.attribute
//       3.条件表达式：1）纯bool值  2）函数bool值  3）逻辑bool值
//       4.函数表达式：FuncXXXXX()函数
//       5.决策表达式：[xxx, xxxx, xxxx]决策选择值
//---------------------------------------------------------------------------------------------
const (
	Expr_Num = iota + 1
	Expr_Float
	Expr_Bool
	Expr_String
	Expr_Date
	Expr_Array
	Expr_Condition
	Expr_Function
	Expr_Variable
	Expr_Candidate
)

var ExpressionClassify = map[int]string{
	Expr_Num :"Expr_Num",
	Expr_Float: "Expr_Float",
	Expr_Bool: "Expr_Bool",
	Expr_String: "Expr_String",
	Expr_Date: "Expr_Date",
	Expr_Array: "Expr_Array",
	Expr_Condition: "Expr_Condition",
	Expr_Function: "Expr_Function",
	Expr_Variable: "Expr_Variable",
	Expr_Candidate: "Expr_Candidate",
}
