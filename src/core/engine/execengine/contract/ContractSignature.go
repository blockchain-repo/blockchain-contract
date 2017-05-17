package contract

import "reflect"

type ContractSignature struct {
	OwnerPubkey string `json:"OwnerPubkey"`
	Signature string `json:"Signature"`
	SignTimestamp string  `json:"SignTimestamp"`
}

func NewContractSignature() *ContractSignature {
	cs := &ContractSignature{}
	return cs
}

func (ca *ContractSignature) GetItem(p_propertyname string)interface{}{
	var r_result interface{}
	//Get Value By reflect
	v_refl_object := reflect.ValueOf(ca).Elem()
	v_refl_field := v_refl_object.FieldByName(p_propertyname)
	r_result = v_refl_field.Interface()
	return r_result
}
func (cs *ContractSignature) GetOwnerPubkey()string {
	return cs.OwnerPubkey
}

func (cs *ContractSignature) GetSignature()string {
	return cs.Signature
}

func (cs *ContractSignature) GetSignTimestamp()string {
	return cs.SignTimestamp
}

func (cs *ContractSignature)SetOwnerPubkey(p_key string) {
	cs.OwnerPubkey = p_key
}

func (cs *ContractSignature)SetSignature(p_Signature string) {
	cs.Signature = p_Signature
}

func (cs *ContractSignature)SetSignTimestamp(p_SignTimestamp string) {
	cs.SignTimestamp = p_SignTimestamp
}