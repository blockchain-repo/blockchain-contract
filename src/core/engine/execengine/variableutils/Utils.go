package variableutils

import (
	"unicontract/src/core/engine/execengine/task"
	"strings"
	"unicontract/src/core/engine/execengine/data"
	"unicontract/src/core/engine/execengine/expression"
	"unicontract/src/core/engine/execengine/inf"
)

func ReflectComponent(p_component interface{}, p_variable string) inf.IComponent {
	var parse_component inf.IComponent
	if strings.HasPrefix(p_variable, "contract_") {
		parse_component = p_component.(inf.ICognitiveContract)
	} else if strings.HasPrefix(p_variable, "enquiry_") {
		parse_component = p_component.(task.Enquiry)
	} else if strings.HasPrefix(p_variable, "action_") {
		parse_component = p_component.(task.Action)
	} else if strings.HasPrefix(p_variable, "decision_") {
		parse_component = p_component.(task.Decision)
	} else if strings.HasPrefix(p_variable, "plan_") {
		parse_component = p_component.(task.Plan)
	} else if strings.HasPrefix(p_variable, "intdata_") {
		parse_component = p_component.(data.IntData)
	} else if strings.HasPrefix(p_variable, "uintdata_") {
		parse_component = p_component.(data.UintData)
	} else if strings.HasPrefix(p_variable, "float_") {
		parse_component = p_component.(data.FloatData)
	} else if strings.HasPrefix(p_variable, "text_") {
		parse_component = p_component.(data.TextData)
	} else if strings.HasPrefix(p_variable, "array_") {
		parse_component = p_component.(data.ArrayData)
	} else if strings.HasPrefix(p_variable, "opresult_") {
		parse_component = p_component.(data.OperateResultData)
	} else if strings.HasPrefix(p_variable, "function_") {
		parse_component = p_component.(expression.Function)
	} else if strings.HasPrefix(p_variable, "logic_") {
		parse_component = p_component.(expression.LogicArgument)
	}
	return parse_component
}