package constdef

const (
	Tag_Or = iota
	Tag_And
	Tag_Not
	Tag_Eq
	Tag_Ne
	Tag_Gt
	Tag_Lt
	Tag_Ge
	Tag_Le
	Tag_Add
	Tag_Sub
	Tag_Mul
	Tag_Div
	Tag_Mod
	Tag_Lb //left bracket
	Tag_Rb //right bracket
)

var ExpressionTag = map[int]string{
	Tag_Or:  "\\|\\|",
	Tag_And: "&&",
	Tag_Eq:  "==",
	Tag_Ne:  "!=",
	Tag_Ge:  ">=",
	Tag_Le:  "<=",
	Tag_Not: "!",
	Tag_Gt:  ">",
	Tag_Lt:  "<",
	Tag_Add: "\\+",
	Tag_Sub: " - ",
	Tag_Mul: "\\*",
	Tag_Div: "/",
	Tag_Mod: "%",
	Tag_Lb:  "(",
	Tag_Rb:  ")",
}

var ExpressionTagString string = ExpressionTag[Tag_Or] + "|" +
	ExpressionTag[Tag_And] + "|" +
	ExpressionTag[Tag_Ne] + "|" +
	ExpressionTag[Tag_Ge] + "|" +
	ExpressionTag[Tag_Le] + "|" +
	ExpressionTag[Tag_Eq] + "|" +
	ExpressionTag[Tag_Not] + "|" +
	ExpressionTag[Tag_Gt] + "|" +
	ExpressionTag[Tag_Lt] + "|" +
	ExpressionTag[Tag_Add] + "|" +
	ExpressionTag[Tag_Sub] + "|" +
	ExpressionTag[Tag_Mul] + "|" +
	ExpressionTag[Tag_Div] + "|" +
	ExpressionTag[Tag_Mod]
