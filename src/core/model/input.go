package model

type Fulfillment struct {
	Fid          int32       `json:"fid"`
	OwnersBefore []string    `json:"owners_before"`
	Fulfillment  interface{} `json:"fulfillment"`
	Input        interface{} `json:"input"`
}

//only used in the `CREATE` operation. when the operation is `TRANSFER`, the `inputs` need to get from the `unichain`
func GenerateInput(public_keys []string) *Fulfillment{
	fulfill := &Fulfillment{
		Fid:0,
		OwnersBefore:public_keys,
		Fulfillment:"cf:4:RtTtCxNf1Bq7MFeIToEosMAa3v_jKtZUtqiWAXyFz1ejPMv-t7vT6DANcrYvKFHAsZblmZ1Xk03HQdJbGiMyb5CmQqGPHwlgKusNu9N_IDtPn7y16veJ1RBrUP-up4YD",
		Input:nil,
	}
	return fulfill
}