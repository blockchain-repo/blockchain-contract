package inf

//contract interface
type ICognitiveContract interface {
	/*
		SetData(IData)
		SetTask(ITask)
		SetExpression(IExpression)
		SetEvent(IEvent)
	*/
	GetContractId() string
	GetVersion() string
	GetCopyRight() string
	GetTask(string) interface{}
	GetComponentTtem(p_name string) interface{}
	GetPropertyItem(p_name string) interface{}
	AddComponent(p_component IComponent)
	EvaluateExpression(p_exprtype string, p_expression string) (interface{}, error)
	ProcessString(string) string

	SetContract(constract ICognitiveContract)
	GetContract() ICognitiveContract
	GetName() string
	GetCtype() string
	GetId() string
	SetOutputId(p_outputId string)
	GetOutputId() string
}

//component
type IComponent interface {
	SetContract(constract ICognitiveContract)
	GetContract() ICognitiveContract
	GetName() string
	GetCtype() string
}

//data interface
type IData interface {
	SetContract(constract ICognitiveContract)
	GetContract() ICognitiveContract
	GetName() string
	GetCtype() string
	GetValue() interface{}
	SetValue(interface{})
}

//task interface
type ITask interface {
	SetContract(constract ICognitiveContract)
	GetContract() ICognitiveContract
	GetName() string
	GetCtype() string
	GetDescription() string
	GetState() string
	SetState(string)
	GetNextTasks() []string
	UpdateState() (int8, error)
}

//expression interface
type IExpression interface {
	SetContract(constract ICognitiveContract)
	GetContract() ICognitiveContract
	GetName() string
	GetCtype() string
	GetExpressionStr() string
	SetExpressionResult(p_expresult interface{})
}
