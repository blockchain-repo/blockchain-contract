package constdef

const (
	Component_Unknown = iota
	Component_Contract
	Component_Task
	Component_Data
	Component_Expression
)

var ComponentType = map[int]string{
	Component_Unknown : "Component_Unknown",
	Component_Contract: "Component_Contract",
	Component_Task: "Component_Task",
	Component_Data: "Component_Data",
	Component_Expression: "Component_Expression",
}

