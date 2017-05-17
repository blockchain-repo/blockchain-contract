package expression

import (
	"unicontract/src/core/engine/execengine/property"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/component"
	"unicontract/src/core/engine/common"
	"github.com/astaxie/beego/logs"
	"errors"
)

type GeneralExpression struct {
	component.GeneralComponent
	ExpressionStr string  `json:"ExpressionStr"`
	ExpressionResult common.OperateResult `json:"ExpressionResult"`
}

const (
	_ExpressionStr = "_ExpressionStr"
	_ExpressionResult = "_ExpressionResult"
)

func NewGeneralExpression(str_expression string) *GeneralExpression{
	v_expression := &GeneralExpression{}
    v_expression.ExpressionStr = str_expression
	return v_expression
}
//===============接口实现===================
func (ge GeneralExpression)SetContract(p_contract inf.ICognitiveContract) {
	ge.GeneralComponent.SetContract(p_contract)
}
func (ge GeneralExpression)GetContract() inf.ICognitiveContract {
	return ge.GeneralComponent.GetContract()
}
func (ge GeneralExpression)GetName()string{
	return ge.GeneralComponent.GetCname()
}
func (ge GeneralExpression)GetCtype()string{
	if ge.PropertyTable["_Ctype"] == nil {
		return ""
	}
	ctype_property := ge.PropertyTable["_Ctype"].(property.PropertyT)
	return ctype_property.GetValue().(string)
}

func (ge *GeneralExpression)SetExpressionResult(p_expresult common.OperateResult){
	ge.ExpressionResult = p_expresult
	result_property := ge.PropertyTable[_ExpressionResult].(property.PropertyT)
	result_property.SetValue(p_expresult)
	ge.PropertyTable[_ExpressionResult] = result_property
}
//===============描述态=====================
func (ge *GeneralExpression)ToString()string{
	return ge.GetCname() + ": " + ge.GetExpressionStr()
}
//===============运行态=====================
func (ge *GeneralExpression) InitExpression() error{
	var err error= nil
	if ge.ExpressionStr == "" {
		logs.Error("ExpressionStr is nil!")
		errors.New("Expression need ExpressionStr!")
		return err
	}
	err = ge.InitGeneralComponent()
	if err != nil {
		logs.Error("InitExpression fail["+err.Error()+"]")
		return err
	}
	ge.Ctype = common.TernaryOperator(ge.Ctype == "", constdef.ComponentType[constdef.Component_Expression], ge.Ctype).(string)
	ge.SetCtype(ge.Ctype)
	common.AddProperty(ge, ge.PropertyTable, _ExpressionStr, ge.ExpressionStr)
	common.AddProperty(ge, ge.PropertyTable, _ExpressionResult, ge.ExpressionResult)
	return err
}
//====属性Get方法
func (ge *GeneralExpression)GetExpressionStr()string {
	express_property := ge.PropertyTable[_ExpressionStr].(property.PropertyT)
	return express_property.GetValue().(string)
}

func (ge *GeneralExpression)GetExpressionResult()common.OperateResult{
	result_property := ge.PropertyTable[_ExpressionResult].(property.PropertyT)
	return result_property.GetValue().(common.OperateResult)
}
//====属性Set方法
func (ge *GeneralExpression)SetExpressionStr(p_expression string){
	ge.ExpressionStr = p_expression
	express_property := ge.PropertyTable[_ExpressionStr].(property.PropertyT)
	express_property.SetValue(p_expression)
	ge.PropertyTable[_ExpressionStr] = express_property
}
