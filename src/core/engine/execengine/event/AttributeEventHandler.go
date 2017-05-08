package event

type AttributeEventHandler struct {
	EventHandler
	attribute string
}

func NewAttributeEventHandle(attribute string)*AttributeEventHandler{
	e := &AttributeEventHandler{attribute:attribute}
	return e
}

func (aeh *AttributeEventHandler) GetAttribute()string{
	return aeh.attribute
}

func (aeh *AttributeEventHandler) SetAttribute(attribute string){
	aeh.attribute = attribute
}

func (aeh *AttributeEventHandler) GetSignature()string{
	return aeh.EventHandler.GetSignature()
}