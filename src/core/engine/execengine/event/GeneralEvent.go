package event

import "unicontract/src/core/engine/execengine/constdef"

type GeneralEvent struct {
	Source string `json:"Source"`
	Etype string `json:"Etype"`
	Urgency string `json:"Urgency"`
	Para string `json:"Para"`
}

func NewGeneralEvent(Source string, Etype string, Urgency string, Para string)*GeneralEvent{
	if Urgency == "" {
		Urgency = constdef.EventPriority[constdef.EventPriority_AfterEngineCycle]
	}
	e := &GeneralEvent{Source:Source, Etype:Etype, Urgency:Urgency, Para:Para}
	return e
}

func (ge *GeneralEvent) GetSource()string{
	return ge.Source
}

func (ge *GeneralEvent) GetEtype() string {
	return ge.Etype
}

func (ge *GeneralEvent) GetUrgency()string{
	return ge.Urgency
}

func (ge *GeneralEvent) GetPara() string{
	return ge.Para
}

func (ge *GeneralEvent) GetSignature() string{
	return ""
}