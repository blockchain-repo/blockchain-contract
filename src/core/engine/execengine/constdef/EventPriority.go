package constdef

const (
	EventPriority_Unknow = iota
	EventPriority_Immediate
	EventPriority_AfterEngineCycle
)

var EventPriority = map[int]string{
	EventPriority_Unknow : "EventPriority_Unknow",
	EventPriority_Immediate: "EventPriority_Immediate",
	EventPriority_AfterEngineCycle: "EventPriority_AfterEngineCycle",
}