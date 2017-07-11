package expression

import (
	"encoding/json"
	"fmt"
	"strconv"
	"unicontract/src/common/uniledgerlog"
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

//序列化： 需要将运行态结构 序列化到 描述态中
func (la *LogicArgument) RunningToStatic() {
	la.Function.RunningToStatic()
	logicValue_property, ok := la.PropertyTable[_LogicValue].(property.PropertyT)
	if ok {
		la.LogicValue, _ = logicValue_property.GetValue().(int)
	}
}

func (la *LogicArgument) Serialize() (string, error) {
	la.RunningToStatic()
	if s_model, err := json.Marshal(la); err == nil {
		return string(s_model), err
	} else {
		uniledgerlog.Error("Contract Expression fail[" + err.Error() + "]")
		return "", err
	}
}

//===============运行态=====================
func (la *LogicArgument) InitLogicArgument() error {
	var err error = nil
	err = la.InitFunction()
	if err != nil {
		uniledgerlog.Error("InitLogicArgument fail[" + err.Error() + "]")
		return err
	}
	la.SetCtype(constdef.ComponentType[constdef.Component_Expression] + "." + constdef.ExpressionType[constdef.Expression_Condition])

	common.AddProperty(la, la.PropertyTable, _LogicValue, la.LogicValue)
	return err
}

//====属性Get方法
func (la *LogicArgument) GetLogicValue() int {
	la.Eval()
	loggicvalue_property, ok := la.PropertyTable[_LogicValue].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0
	}
	n, ok := loggicvalue_property.GetValue().(int)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0
	}
	return n
}

//====属性Set方法
func (la *LogicArgument) SetLogicValue(p_int interface{}) {
	if p_int == nil {
		uniledgerlog.Warn("[Param]p_int is nil，Check it!")
		return
	}
	ok := false
	la.LogicValue, ok = p_int.(int)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	loggicvalue_property, ok := la.PropertyTable[_LogicValue].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	loggicvalue_property.SetValue(la.LogicValue)
	la.PropertyTable[_LogicValue] = loggicvalue_property
}

func (la *LogicArgument) Eval() int {
	expression_property, ok := la.PropertyTable[_ExpressionStr].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0
	}
	v_expression, ok := expression_property.GetValue().(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return 0
	}
	r_flag, r_err := la.GetContract().EvaluateExpression(constdef.ExpressionType[constdef.Expression_Condition], v_expression)
	if r_err != nil {
		uniledgerlog.Warn("LogicArgument.Eval fail[" + r_err.Error() + "]")
		return la.LogicValue
	}
	var v_value int = 0
	b, ok := r_flag.(bool)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
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
