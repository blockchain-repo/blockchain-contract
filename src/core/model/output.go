package model

type ConditionDetails struct {
	Bitmask   int32  `json:"bitmask"`
	PublicKey string `json:"public_key"`
	Signature string `json:"signature"`
	Type      string `json:"type"`
	TypeId    int32  `json:"type_id"`
}

type Condition struct {
	Details *ConditionDetails `json:"details"`
	Uri     string            `json:"uri"`
}

type ConditionsItem struct {
	Amount      int64      `json:"amount"`
	Cid         int32      `json:"cid"`
	Condition   *Condition `json:"condition"`
	OwnersAfter []string   `json:"owners_after"`
	Isfreeze    bool       `json:"isfreeze"`
}


func GenerateOutput(isFeeze bool,pub string, amount int) *ConditionsItem{

	return &ConditionsItem{}
}