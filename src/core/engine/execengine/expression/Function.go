package expression

import (
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/property"
)

type Function struct {
	GeneralExpression
	//type: inf.IExpression
	SelectBranchs []string `json:"SelectBranchs"`
}
const(
	_SelectBranchs = "_SelectBranchs"
)

func NewFunction() *Function{
	nf := &Function{}
	return nf
}
//===============接口实现===================
func (f Function)SetContract(p_contract inf.ICognitiveContract) {
	f.GeneralExpression.SetContract(p_contract)
}
func (f Function)GetContract() inf.ICognitiveContract {
	return f.GeneralExpression.GetContract()
}
func (f Function)GetName()string{
	return f.GeneralExpression.GetCname()
}
func (f Function)GetCtype()string{
	return f.GeneralExpression.GetCtype()
}
func (f Function)GetExpressionStr()string{
	return f.GeneralExpression.GetExpressionStr()
}
func (f Function)SetExpressionResult(p_expresult common.OperateResult){
	f.GeneralExpression.SetExpressionResult(p_expresult)
}
//===============描述态=====================
func (f *Function)ToString()string{
	return f.GetCname()
}
//===============运行态=====================
func (f *Function) InitFunction()error{
	var err error = nil
	err = f.InitExpression()
	if err != nil {
		//TODO log
		return err
	}
	f.Ctype = common.TernaryOperator(f.Ctype == "", constdef.ComponentType[constdef.Component_Expression] + "." + constdef.ExpressionType[constdef.Expression_Function], f.Ctype).(string)
	f.SetCtype(f.Ctype)

	if f.SelectBranchs == nil {
		f.SelectBranchs = make([]string, 0)
	}
	common.AddProperty(f, f.PropertyTable, _SelectBranchs, f.SelectBranchs)
	return err
}

//====Get方法
func (f *Function) GetSelectBranchs()[]string{
	selectbranch_property := f.PropertyTable[_LogicValue].(property.PropertyT)
	return selectbranch_property.GetValue().([]string)
}
//====Set方法
func (f *Function) SetSelectBranchs(p_selectbranchs []string){
	f.SelectBranchs = p_selectbranchs
	selectbranch_property := f.PropertyTable[_SelectBranchs].(property.PropertyT)
	selectbranch_property.SetValue(p_selectbranchs)
	f.PropertyTable[_SelectBranchs] = selectbranch_property
}