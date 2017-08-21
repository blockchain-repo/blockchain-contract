package expression

import (
	"encoding/json"
	"errors"
	"fmt"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/component"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
)

type GeneralExpression struct {
	component.GeneralComponent
	ExpressionStr    string               `json:"ExpressionStr"`
	ExpressionResult common.OperateResult `json:"-"`
}

const (
	_ExpressionStr    = "_ExpressionStr"
	_ExpressionResult = "_ExpressionResult"
)

func NewGeneralExpression(str_name string, str_expression string) *GeneralExpression {
	v_expression := &GeneralExpression{}
	v_expression.Cname = str_name
	v_expression.ExpressionStr = str_expression
	return v_expression
}

//===============接口实现===================
func (ge GeneralExpression) SetContract(p_contract inf.ICognitiveContract) {
	ge.GeneralComponent.SetContract(p_contract)
}
func (ge GeneralExpression) GetContract() inf.ICognitiveContract {
	return ge.GeneralComponent.GetContract()
}
func (ge GeneralExpression) GetName() string {
	return ge.GeneralComponent.GetCname()
}
func (ge GeneralExpression) GetCtype() string {
	if ge.PropertyTable["_Ctype"] == nil {
		return ""
	}
	ctype_property, ok := ge.PropertyTable["_Ctype"].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	str, ok := ctype_property.GetValue().(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	return str
}

func (ge GeneralExpression) SetExpressionResult(p_expresult interface{}) {
	ok := false
	ge.ExpressionResult, ok = p_expresult.(common.OperateResult)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	result_property, ok := ge.PropertyTable[_ExpressionResult].(property.PropertyT)
	if !ok {
		result_property = *property.NewPropertyT(_ExpressionResult)
	}
	result_property.SetValue(p_expresult)
	ge.PropertyTable[_ExpressionResult] = result_property
}

func (ge GeneralExpression) CleanValueInProcess() {
	ge.SetExpressionResult(common.OperateResult{Code: 0, Message: "", Data: "", Output: ""})
}

//===============描述态=====================
func (ge *GeneralExpression) ToString() string {
	return ge.GetCname() + ": " + ge.GetExpressionStr()
}

//序列化： 需要将运行态结构 序列化到 描述态中
func (ge *GeneralExpression) RunningToStatic() {
	cname_property, ok := ge.PropertyTable["_Cname"].(property.PropertyT)
	if ok {
		ge.Cname, _ = cname_property.GetValue().(string)
	}
	ctype_property, ok := ge.PropertyTable["_Ctype"].(property.PropertyT)
	if ok {
		ge.Ctype, _ = ctype_property.GetValue().(string)
	}
	caption_property, ok := ge.PropertyTable["_Caption"].(property.PropertyT)
	if ok {
		ge.Caption, _ = caption_property.GetValue().(string)
	}
	description_property, ok := ge.PropertyTable["_Description"].(property.PropertyT)
	if ok {
		ge.Description, _ = description_property.GetValue().(string)
	}
	metaAttribute_property, ok := ge.PropertyTable["_MetaAttribute"].(property.PropertyT)
	if ok {
		ge.MetaAttribute, _ = metaAttribute_property.GetValue().(map[string]string)
	}
	expressionStr_property, ok := ge.PropertyTable[_ExpressionStr].(property.PropertyT)
	if ok {
		ge.ExpressionStr, _ = expressionStr_property.GetValue().(string)
	}
	expressionResult_property, ok := ge.PropertyTable[_ExpressionResult].(property.PropertyT)
	if ok {
		ge.ExpressionResult, _ = expressionResult_property.GetValue().(common.OperateResult)
	}
}

func (ge *GeneralExpression) Serialize() (string, error) {
	ge.RunningToStatic()
	if s_model, err := json.Marshal(ge); err == nil {
		return string(s_model), err
	} else {
		uniledgerlog.Error("Expression Serialize fail[" + err.Error() + "]")
		return "", err
	}
}

//===============运行态=====================
func (ge *GeneralExpression) InitExpression() error {
	var err error = nil
	if ge.ExpressionStr == "" {
		uniledgerlog.Error("ExpressionStr is nil!")
		errors.New("Expression need ExpressionStr!")
		return err
	}
	err = ge.InitGeneralComponent()
	if err != nil {
		uniledgerlog.Error("InitExpression fail[" + err.Error() + "]")
		return err
	}
	ge.SetCtype(constdef.ComponentType[constdef.Component_Expression])
	common.AddProperty(ge, ge.PropertyTable, _ExpressionStr, ge.ExpressionStr)
	common.AddProperty(ge, ge.PropertyTable, _ExpressionResult, ge.ExpressionResult)

	return err
}

//====属性Get方法
func (ge *GeneralExpression) GetExpressionStr() string {
	express_property, ok := ge.PropertyTable[_ExpressionStr].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	str, ok := express_property.GetValue().(string)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return ""
	}
	return str
}

func (ge *GeneralExpression) GetExpressionResult() common.OperateResult {
	var result common.OperateResult
	result_property, ok := ge.PropertyTable[_ExpressionResult].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return result
	}
	result, ok = result_property.GetValue().(common.OperateResult)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
	}
	return result
}

//====Set方法
func (ge *GeneralExpression) SetExpressionStr(p_expression string) {
	ge.ExpressionStr = p_expression
	express_property, ok := ge.PropertyTable[_ExpressionStr].(property.PropertyT)
	if !ok {
		express_property = *property.NewPropertyT(_ExpressionStr)
	}
	express_property.SetValue(p_expression)
	ge.PropertyTable[_ExpressionStr] = express_property
}
