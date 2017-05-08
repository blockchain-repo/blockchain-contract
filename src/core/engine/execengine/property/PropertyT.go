package property

import (
	"strings"
)

type PropertyT struct {
	Name string
	Value interface{}
	OldValue interface{}
	NewValue interface{}
}

func NewPropertyT(name string) *PropertyT{
	p := &PropertyT{Name:name}
	return p
}

func (p *PropertyT) GetName()string{
	return p.Name
}

func (p *PropertyT) GetValue() interface{}{
	return p.Value
}

func (p *PropertyT) GetOldValue() interface{}{
	return p.OldValue
}

func (p *PropertyT) GetNewValue() interface{}{
	return p.NewValue
}

func (p *PropertyT) SetValue(value interface{}){
	p.Value = value
}

func (p *PropertyT) SetOldValue(oldValue interface{}){
	p.OldValue = oldValue
}

func (p *PropertyT) SetNewValue(newValue interface{}){
	p.NewValue = newValue
}

func (p *PropertyT) ToString() string{
	var p_str string
	p_str = strings.Join([]string{"name:", p.Name}, "")
	return p_str
}
