package common

type SelectBranchExpression struct {
	BranchExpressionStr   string      `json:"BranchExpressionStr"`
	BranchExpressionValue interface{} `json:"BranchExpressionValue"`
}

func NewSelectBranchExpression() *SelectBranchExpression {
	nbe := &SelectBranchExpression{}
	return nbe
}

func (nbe *SelectBranchExpression) GetBranchExpressionStr() string {
	return nbe.BranchExpressionStr
}

func (nbe *SelectBranchExpression) GetBranchExpressionValue() interface{} {
	return nbe.BranchExpressionValue
}

func (nbe *SelectBranchExpression) SetBranchExpressionStr(p_BranchExpressionStr string) {
	nbe.BranchExpressionStr = p_BranchExpressionStr
}

func (nbe *SelectBranchExpression) SetBranchExpressionValue(p_BranchExpressionValue interface{}) {
	nbe.BranchExpressionValue = p_BranchExpressionValue
}
