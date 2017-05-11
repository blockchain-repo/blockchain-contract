package model

type ContractOutputLink struct {
	Cid  int    `json:"cid"`
	Txid string `json:"txid"`
}

type UnSpentOutput struct {
	Cid    int		`json:"cid"`
	Txid   string	`json:"txid"`
	Amount int		`json:"amount"`
}

type Fulfillment struct {
	Fid          int                 `json:"fid"`
	OwnersBefore []string            `json:"owners_before"`
	Fulfillment  interface{}         `json:"fulfillment"`
	Input        *ContractOutputLink `json:"input"`
}

//only used in the `CREATE` operation. when the operation is `TRANSFER`, the `inputs` need to get from the `unichain`
func (f *Fulfillment) GenerateInput(public_keys []string) {
	f.Fid = 0
	f.OwnersBefore = public_keys
	f.Fulfillment = "cf:4:RtTtCxNf1Bq7MFeIToEosMAa3v_jKtZUtqiWAXyFz1ejPMv-t7vT6DANcrYvKFHAsZblmZ1Xk03HQdJbGiMyb5CmQqGPHwlgKusNu9N_IDtPn7y16veJ1RBrUP-up4YD"
	//f.Input = nil
}
