package contract

type ContractSignature struct {
	OwnerPubkey string `json:"OwnerPubkey"`
	Signature string `json:"Signature"`
	SignTimestamp string  `json:"SignTimestamp"`
}

func NewContractSignature() *ContractSignature {
	cs := &ContractSignature{}
	return cs
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