package variableutils

import (
	"strings"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/execengine/data"
	"unicontract/src/core/engine/execengine/expression"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/task"
)

func ReflectComponent(p_component interface{}, p_variable string) inf.IComponent {
	var parse_component inf.IComponent
	ok := false
	if strings.HasPrefix(p_variable, "contract_") {
		parse_component, ok = p_component.(inf.ICognitiveContract)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if strings.HasPrefix(p_variable, "enquiry_") {
		parse_component, ok = p_component.(task.Enquiry)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if strings.HasPrefix(p_variable, "action_") {
		parse_component, ok = p_component.(task.Action)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if strings.HasPrefix(p_variable, "decision_") {
		parse_component, ok = p_component.(task.Decision)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if strings.HasPrefix(p_variable, "plan_") {
		parse_component, ok = p_component.(task.Plan)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if strings.HasPrefix(p_variable, "intdata_") {
		parse_component, ok = p_component.(data.IntData)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if strings.HasPrefix(p_variable, "uintdata_") {
		parse_component, ok = p_component.(data.UintData)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if strings.HasPrefix(p_variable, "float_") {
		parse_component, ok = p_component.(data.FloatData)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if strings.HasPrefix(p_variable, "text_") {
		parse_component, ok = p_component.(data.TextData)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if strings.HasPrefix(p_variable, "array_") {
		parse_component, ok = p_component.(data.ArrayData)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if strings.HasPrefix(p_variable, "opresult_") {
		parse_component, ok = p_component.(data.OperateResultData)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if strings.HasPrefix(p_variable, "function_") {
		parse_component, ok = p_component.(expression.Function)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if strings.HasPrefix(p_variable, "logic_") {
		parse_component, ok = p_component.(expression.LogicArgument)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	}
	return parse_component
}
