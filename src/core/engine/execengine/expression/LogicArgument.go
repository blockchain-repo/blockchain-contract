package expression

import (
	"github.com/astaxie/beego/logs"
	"strconv"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
)

type LogicArgument struct {
	Function
	LogicValue int `json:"LogicValue"`
}

func NewLogicArgument() *LogicArgument {
	la := &LogicArgument{}
	return la
}

const (
	_LogicValue = "_LogicValue"
)

//===============接口实现===================
func (la LogicArgument) SetContract(p_contract inf.ICognitiveContract) {
	la.Function.SetContract(p_contract)
}
func (la LogicArgument) GetContract() inf.ICognitiveContract {
	return la.Function.GetContract()
}
func (la LogicArgument) GetName() string {
	return la.Function.GetCname()
}
func (la LogicArgument) GetCtype() string {
	return la.Function.GetCtype()
}
func (la LogicArgument) GetExpressionStr() string {
	return la.Function.GetExpressionStr()
}
func (la LogicArgument) SetExpressionResult(p_expresult interface{}) {
	la.GeneralExpression.SetExpressionResult(p_expresult)
}

func (la LogicArgument) CleanValueInProcess() {
	la.GeneralExpression.CleanValueInProcess()
	la.SetLogicValue(0)
}

//===============描述态=====================
func (la *LogicArgument) ToString() string {
	return la.GetCname() + ":" + strconv.Itoa(la.GetLogicValue())
}

//===============运行态=====================
func (la *LogicArgument) InitLogicArgument() error {
	var err error = nil
	err = la.InitFunction()
	if err != nil {
		logs.Error("InitLogicArgument fail[" + err.Error() + "]")
		return err
	}
	la.SetCtype(constdef.ComponentType[constdef.Component_Expression] + "." + constdef.ExpressionType[constdef.Expression_Condition])

	common.AddProperty(la, la.PropertyTable, _LogicValue, la.LogicValue)
	return err
}

//====属性Get方法
func (la *LogicArgument) GetLogicValue() int {
	la.Eval()
	loggicvalue_property,ok := la.PropertyTable[_LogicValue].(property.PropertyT)
	if !ok{
		logs.Error("assert error")
		return 0
	}
	n,ok:=loggicvalue_property.GetValue().(int)
	if !ok{
		logs.Error("assert error")
		return 0
	}
	return n
}

//====属性Set方法
func (la *LogicArgument) SetLogicValue(p_int interface{}) {
	if p_int == nil {
		logs.Warning("[Param]p_int is nil，Check it!")
		return
	}
	ok:=false
	la.LogicValue,ok = p_int.(int)
	if !ok{
		logs.Error("assert error")
		return
	}
	loggicvalue_property,ok := la.PropertyTable[_LogicValue].(property.PropertyT)
	if !ok{
		logs.Error("assert error")
		return
	}
	loggicvalue_property.SetValue(la.LogicValue)
	la.PropertyTable[_LogicValue] = loggicvalue_property
}

func (la *LogicArgument) Eval() int {
	expression_property,ok := la.PropertyTable[_ExpressionStr].(property.PropertyT)
	if !ok{
		logs.Error("assert error")
		return 0
	}
	v_expression,ok := expression_property.GetValue().(string)
	if !ok{
		logs.Error("assert error")
		return 0
	}
	r_flag, r_err := la.GetContract().EvaluateExpression(constdef.ExpressionType[constdef.Expression_Condition], v_expression)
	if r_err != nil {
		logs.Warning("LogicArgument.Eval fail[" + r_err.Error() + "]")
		return la.LogicValue
	}
	var v_value int = 0
	b,ok:=r_flag.(bool)
	if !ok{
		logs.Error("assert error")
		return 0
	}
	if b {
		v_value = 1
	} else {
		v_value = 0
	}
	la.SetLogicValue(v_value)
	return la.LogicValue
}
