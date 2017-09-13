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
	GetUCVMVersion() string
	GetUCVMCopyRight() string
	GetTask(string) interface{}
	GetData(string) interface{}
	GetExpression(string) interface{}
	GetTaskByID(string) interface{}
	GetComponentTtem(string) interface{}
	GetPropertyItem(string) interface{}
	AddComponent(IComponent)
	EvaluateExpression(string, string) (interface{}, error)
	ProcessString(string) string
	UpdateComponentRunningState(string, string, interface{}) error

	SetContract(ICognitiveContract)
	GetContract() ICognitiveContract
	GetName() string
	GetCtype() string
	GetContractState() string

	GetId() string
	GetOrgTaskId() string
	GetOrgTaskExecuteIdx() int
	GetOutputId() string
	GetOutputTaskId() string
	GetOutputTaskExecuteIdx() int
	GetOutputStruct() string

	SetOrgId(string)
	SetOrgTaskId(string)
	SetOrgTaskExecuteIdx(int)
	SetOutputId(p_outputId string)
	SetOutputTaskId(string)
	SetOutputTaskExecuteIdx(int)
	SetOutputStruct(string)

	GetMainPubkey() string
	SetMainPubkey(string)

	UpdateContractState(string) bool

	Serialize() (string, error)
	Deserialize(p_str string) (interface{}, error)
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
	CleanValueInProcess()
	Serialize() (string, error)
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
	UpdateState(nBrotherNum int) (int8, error)
	GetTaskId() string
	GetTaskExecuteIdx() int
	SetTaskId(string)
	SetTaskExecuteIdx(int)
	CleanValueInProcess()
	UpdateStaticState() (interface{}, error)
}

//expression interface
type IExpression interface {
	SetContract(constract ICognitiveContract)
	GetContract() ICognitiveContract
	GetName() string
	GetCtype() string
	GetExpressionStr() string
	SetExpressionResult(p_expresult interface{})
	CleanValueInProcess()
	Serialize() (string, error)
}
