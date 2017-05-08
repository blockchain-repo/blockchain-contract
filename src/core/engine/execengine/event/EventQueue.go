package event

import (
	"unicontract/src/common/basic"
)

type EventQueue struct {
	EventList *basic.Queue
}

func (eq *EventQueue) ShowAllEvnet()*basic.Queue{
	return eq.EventList
}

func (eq *EventQueue) AddEvent(p_event interface{}){
	if eq.EventList.Len() == 0 {
		eq.EventList = basic.NewQueue()
	}
	eq.EventList.Push(p_event)
}

func (eq *EventQueue) popEvent() interface{}{
	var q_event interface{}
	if eq.EventList.Len() != 0 {
		q_event = eq.EventList.Pop()
	}
	return q_event
}

func (eq *EventQueue) FireTopHandler(){
	//var f_event interface{} = eq.popEvent()
	//TODO
}

func (eq *EventQueue) FireHandlers(){
	for {
		if eq.EventList == nil || eq.EventList.Len() == 0{
			break
		} else {
			//var f_event interface{} = eq.popEvent()
			//TODO
		}
	}
}


