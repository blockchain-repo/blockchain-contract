package constdef

// 表达式分为4大类：  1. 常量表达式  2. 变量表达式  3.条件表达式   4.函数表达式   5.决策表达式
const (
	Expression_Unknown = iota
	Expression_Constant
	Expression_Variable
	Expression_Condition
	Expression_Function
	Expression_Candidate
)

var ExpressionType = map[int]string{
	Expression_Unknown:   "Expression_Unknown",
	Expression_Constant:  "Expression_Constant",
	Expression_Variable:  "Expression_Variable",
	Expression_Condition: "Expression_Condition",
	Expression_Function:  "Expression_Function",
	Expression_Candidate: "Expression_Candidate",
}
