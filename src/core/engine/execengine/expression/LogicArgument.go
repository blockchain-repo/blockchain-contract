package expression

import (
	"unicontract/src/core/engine/execengine/property"
	"unicontract/src/core/engine/execengine/inf"
	"strconv"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/common"
)

type LogicArgument struct {
	Function
	LogicValue int `json:"LogicValue"`
}

const (
	_LogicValue = "_LogicValue"
)
//===============接口实现===================
func (la LogicArgument)SetContract(p_contract inf.ICognitiveContract) {
	la.Function.SetContract(p_contract)
}
func (la LogicArgument)GetContract() inf.ICognitiveContract {
	return la.Function.GetContract()
}
func (la LogicArgument)GetName()string{
	return la.Function.GetCname()
}
func (la LogicArgument)GetCtype()string{
	return la.Function.GetCtype()
}
func (la LogicArgument)GetExpressionStr()string{
	return la.Function.GetExpressionStr()
}
//===============描述态=====================
func (la *LogicArgument)ToString() string{
	return la.GetCname() + ":" + strconv.Itoa(la.GetValue())
}
//===============运行态=====================
func (la *LogicArgument)InitLogicArgument()error{
	var err error = nil
	err = la.InitFunction()
	if err != nil {
		//TODO log
		return err
	}
	la.Ctype = common.TernaryOperator(la.Ctype == "", constdef.ComponentType[constdef.Component_Expression] + "." + constdef.ExpressionType[constdef.Expression_LogicArgument], la.Ctype).(string)
	la.SetCtype(la.Ctype)

	la.AddProperty(la, _LogicValue,la.LogicValue)
	return err
}

//====属性Get方法
func (la *LogicArgument)GetValue()int{
	la.Eval()
	loggicvalue_property := la.PropertyTable[_LogicValue].(property.PropertyT)
	return loggicvalue_property.GetValue().(int)
}

func (la *LogicArgument)Eval() int{
	expression_property := la.PropertyTable[_ExpressionStr].(property.PropertyT)
	var v_expression string = expression_property.GetValue().(string)
	var r_flag bool = la.GetContract().EvaluateExpression(v_expression).(bool)
	loggicvalue_property := la.PropertyTable[_LogicValue].(property.PropertyT)
	if r_flag {
		la.LogicValue = 1
	} else {
		la.LogicValue = 0
	}
	loggicvalue_property.SetValue(la.LogicValue)
	la.PropertyTable[_LogicValue] = loggicvalue_property
	return la.LogicValue
}