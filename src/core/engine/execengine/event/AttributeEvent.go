package event

import (
	"strings"
	"unicontract/src/core/engine/execengine/constdef"
)

type AttributeEvent struct {
	GeneralEvent
	OldValue string `json:"OldValue"`
	NewValue string `json:"NewValue"`
	Attribute string `json:"Attribute"`
}

func NewAttributeEvent(source string, Attribute string, OldValue string, NewValue string,
	para string, urgency string) *AttributeEvent{
	if urgency == "" {
		urgency = constdef.EventPriority[constdef.EventPriority_AfterEngineCycle]
	}
	e := &AttributeEvent{GeneralEvent{Source:source, Etype:constdef.EventType[constdef.EventType_Attribute], Urgency:urgency, Para:para},
			OldValue, NewValue, Attribute}
	return e
}

func (ae *AttributeEvent) GetAttribute() string{
	return ae.Attribute
}

func (ae *AttributeEvent) GetOldValue() string{
	return ae.OldValue
}

func (ae *AttributeEvent) GetNewValue() string{
	return ae.NewValue
}

func (ae *AttributeEvent) GetSignature() string{
	var r_str string = strings.Join([]string{ae.GetSource(), ae.GetAttribute(), ae.GetOldValue(), ae.GetNewValue()}, "")
	return r_str
}