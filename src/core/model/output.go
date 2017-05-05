package model

type ConditionDetails struct {
	Bitmask   int         `json:"bitmask"`
	PublicKey string      `json:"public_key"`
	Signature interface{} `json:"signature"`
	Type      string      `json:"type"`
	TypeId    int         `json:"type_id"`
}

type Condition struct {
	Details *ConditionDetails `json:"details"`
	Uri     string            `json:"uri"`
}

type ConditionsItem struct {
	Amount      int        `json:"amount"`
	Cid         int        `json:"cid"`
	Condition   *Condition `json:"condition"`
	OwnersAfter []string   `json:"owners_after"`
	Isfreeze    bool       `json:"isfreeze"`
}

func (c *ConditionsItem)GenerateOutput(cid int, isFeeze bool, pub string, amount int) {
	condetails := ConditionDetails{
		Bitmask:   32,
		PublicKey: pub,
		//Signature: nil,
		Type:   "fulfillment",
		TypeId: 4,
	}
	cond := Condition{
		Details: &condetails,
		Uri:     "",
	}
	//cc:4:20:RtTtCxNf1Bq7MFeIToEosMAa3v_jKtZUtqiWAXyFz1c:96
	c.Amount=amount
	c.Cid=cid
	c.Condition=&cond
	c.OwnersAfter=[]string{pub}
	c.Isfreeze=isFeeze

}

