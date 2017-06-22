package variableutils

import (
	"github.com/astaxie/beego/logs"
	"strings"
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
			logs.Error("assert error")
		}
	} else if strings.HasPrefix(p_variable, "enquiry_") {
		parse_component, ok = p_component.(task.Enquiry)
		if !ok {
			logs.Error("assert error")
		}
	} else if strings.HasPrefix(p_variable, "action_") {
		parse_component, ok = p_component.(task.Action)
		if !ok {
			logs.Error("assert error")
		}
	} else if strings.HasPrefix(p_variable, "decision_") {
		parse_component, ok = p_component.(task.Decision)
		if !ok {
			logs.Error("assert error")
		}
	} else if strings.HasPrefix(p_variable, "plan_") {
		parse_component, ok = p_component.(task.Plan)
		if !ok {
			logs.Error("assert error")
		}
	} else if strings.HasPrefix(p_variable, "intdata_") {
		parse_component, ok = p_component.(data.IntData)
		if !ok {
			logs.Error("assert error")
		}
	} else if strings.HasPrefix(p_variable, "uintdata_") {
		parse_component, ok = p_component.(data.UintData)
		if !ok {
			logs.Error("assert error")
		}
	} else if strings.HasPrefix(p_variable, "float_") {
		parse_component, ok = p_component.(data.FloatData)
		if !ok {
			logs.Error("assert error")
		}
	} else if strings.HasPrefix(p_variable, "text_") {
		parse_component, ok = p_component.(data.TextData)
		if !ok {
			logs.Error("assert error")
		}
	} else if strings.HasPrefix(p_variable, "array_") {
		parse_component, ok = p_component.(data.ArrayData)
		if !ok {
			logs.Error("assert error")
		}
	} else if strings.HasPrefix(p_variable, "opresult_") {
		parse_component, ok = p_component.(data.OperateResultData)
		if !ok {
			logs.Error("assert error")
		}
	} else if strings.HasPrefix(p_variable, "function_") {
		parse_component, ok = p_component.(expression.Function)
		if !ok {
			logs.Error("assert error")
		}
	} else if strings.HasPrefix(p_variable, "logic_") {
		parse_component, ok = p_component.(expression.LogicArgument)
		if !ok {
			logs.Error("assert error")
		}
	}
	return parse_component
}
