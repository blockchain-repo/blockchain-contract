package constdef

const (
	Expression_Unknown = iota
	Expression_Expression
	Expression_Function
	Expression_LogicArgument
)

var ExpressionType = map[int]string{
	Expression_Unknown : "Expression_Unknown",
	Expression_Expression: "Expression_Expression",
	Expression_Function: "Expression_Function",
	Expression_LogicArgument: "Expression_LogicArgument",
}