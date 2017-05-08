package event

import (
	"unicontract/src/core/engine/execengine/component"
)

type EventHandler struct {
	component.GeneralComponent
	eventType int
	signature string
	handler []string
}

func NewEventHandler(signature string, eventtype int)*EventHandler{
	h := &EventHandler{signature:signature, eventType:eventtype}
	return h
}

func (eh *EventHandler) GetEventType()int{
	return eh.eventType
}

func (eh *EventHandler) GetSignature() string{
	return eh.signature
}

func (eh *EventHandler) AddHandler(p_handler string){
	if eh.handler != nil {
		eh.handler = make([]string, 0)
	}else{
		eh.handler = append(eh.handler, p_handler)
	}
}

func (eh *EventHandler) Fire(){
	//for _,value := range eh.handler{
		//TODO
		//contract express
		//eh.contract.EvaluateExpression(value)
	//}
}

