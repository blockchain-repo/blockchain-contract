package model

type Fulfillment struct {
	Fid          int32       `json:"fid"`
	OwnersBefore []string    `json:"owners_before"`
	Fulfillment  interface{} `json:"fulfillment"`
	Input        interface{} `json:"input"`
}


func GenerateInput(public_keys []string) *Fulfillment{

	return &Fulfillment{}
}