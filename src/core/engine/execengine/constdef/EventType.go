package constdef

const (
	EventType_Unknow = iota
	EventType_Engine
	EventType_Attribute
	EventType_Component
)

var EventType = map[int]string{
	EventType_Unknow : "EventType_Unknow",
	EventType_Engine: "EventType_Engine",
	EventType_Attribute: "EventType_Attribute",
	EventType_Component: "EventType_Component",
}