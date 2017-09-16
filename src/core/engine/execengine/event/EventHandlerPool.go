package event

type EventHandlerPool struct {
	eventPool map[string]interface{}
}

func NewEventHandlerPool() *EventHandlerPool {
	var ele map[string]interface{} = make(map[string]interface{}, 0)
	p := &EventHandlerPool{eventPool: ele}
	return p
}

func (ep *EventHandlerPool) AddHandler(handler interface{}) {
	if handler != nil && ep.eventPool != nil {
		switch handler.(type) {
		case EventHandler:
			v_handle := handler.(EventHandler)
			ep.eventPool[v_handle.GetSignature()] = handler
		case AttributeEvent:
			v_handle := handler.(AttributeEvent)
			ep.eventPool[v_handle.GetSignature()] = handler
		}
	}
}

func (ep *EventHandlerPool) RemoveHandler(handler interface{}) {
	if handler != nil && ep.eventPool != nil && len(ep.eventPool) > 0 {
		switch handler.(type) {
		case EventHandler:
			v_handle := handler.(EventHandler)
			delete(ep.eventPool, v_handle.GetSignature())
		case AttributeEvent:
			v_handle := handler.(AttributeEvent)
			delete(ep.eventPool, v_handle.GetSignature())
		}
	}
}

func (ep *EventHandlerPool) GetHandler(signature string) interface{} {
	if ep.eventPool != nil && len(ep.eventPool) > 0 {
		return ep.eventPool[signature]
	}
	return nil
}
