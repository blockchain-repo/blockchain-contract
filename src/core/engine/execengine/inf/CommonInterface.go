package inf

//contract interface
type ICognitiveContract interface {
	/*
	SetData(IData)
	SetTask(ITask)
	SetExpression(IExpression)
	SetEvent(IEvent)
	*/
	GetContractId()string
	GetVersion()string
	GetCopyRight()string
	GetTask(string)interface{}
	AddComponent(IComponent)
	EvaluateExpression(interface{})interface{}
	ProcessString(string)string
}

//component
type IComponent interface{
	SetContract(constract ICognitiveContract)
	GetContract()ICognitiveContract
	GetName() string
	GetCtype() string
}

//data interface
type IData interface {
	SetContract(constract ICognitiveContract)
	GetContract() ICognitiveContract
	GetName()string
	GetCtype() string
	GetValue()interface{}
}

//task interface
type ITask interface {
	SetContract(constract ICognitiveContract)
	GetContract() ICognitiveContract
	GetName()string
	GetCtype() string
	GetDescription()string
	GetState() string
	SetState(string)
	GetNextTasks() []string
	UpdateState() (int8, error)
}

//expression interface
type IExpression interface {
	SetContract(constract ICognitiveContract)
	GetContract() ICognitiveContract
	GetName()string
	GetCtype() string
	GetExpressionStr()string
}

//event interface
type IEvent interface {
	SetContract(constract ICognitiveContract)
	GetContract() ICognitiveContract
}