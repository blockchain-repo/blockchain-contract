package transaction

import (
	"unicontract/src/common"
	"unicontract/src/core/model"
	"unicontract/src/config"
)

const (
	_GENESIS  = "GENESIS"
	_CREATE   = "CREATE"
	_TRANSFER = "TRANSFER"
	_CONTRACT = "CONTRACT"
	_FREEZEASSET   = "FREEZE"
	_UNFREEZEASSET = "UNFREEZE"
	_VERSION = 2
)

var ALLOWED_OPERATIONS = [4]string{_GENESIS, _CREATE, _TRANSFER, _CONTRACT}

//create asset
func Create(tx_signers []string, recipients [][2]interface{}, metadata model.Metadata, asset model.Asset,
	relation model.Relation, contract model.ContractModel) model.ContractOutput {
	isFeeze := false
	operation := _CREATE
	version := _VERSION
	timestamp := common.GenTimestamp()

	//generate outputs
	outputs := []*model.ConditionsItem{}
	for index, recipient := range recipients {
		pubkey := recipient[0].(string)
		amount := recipient[1].(int)
		outputs = append(outputs, model.GenerateOutput(index,isFeeze, pubkey, amount))
	}

	//generate inputs
	inputs := []*model.Fulfillment{}
	inputs = append(inputs, model.GenerateInput(tx_signers))
	contractOutput := model.GenerateConOutput(operation, asset, inputs, outputs, metadata, timestamp, version, relation, contract)
	return contractOutput
}

//transfer asset include:transfer/freeze/unfreeze
func Transfer(operation string, inputs []*model.Fulfillment, recipients [][2]interface{}, metadata model.Metadata, asset model.Asset,
	relation model.Relation, contract model.ContractModel) model.ContractOutput {
	isFeeze := false
	if operation == _FREEZEASSET {
		isFeeze = true
	}
	operation = _TRANSFER
	version := _VERSION
	timestamp := common.GenTimestamp()

	//generate outputs
	outputs := []*model.ConditionsItem{}
	for index, recipient := range recipients {
		pubkey := recipient[0].(string)
		amount := recipient[1].(int)
		outputs = append(outputs, model.GenerateOutput(index,isFeeze, pubkey, amount))
	}

	contractOutput := model.GenerateConOutput(operation, asset, inputs, outputs, metadata, timestamp, version, relation, contract)
	return contractOutput
}


func NodeSign(contractOutput model.ContractOutput)model.ContractOutput {
	vote1 := &model.Vote{}
	vote1.Id = common.GenerateUUID()
	vote1.NodePubkey = config.Config.Keypair.PublicKey
	vote1.VoteBody.Timestamp = common.GenTimestamp()
	vote1.VoteBody.InvalidReason = "resion"
	vote1.VoteBody.IsValid = true
	vote1.VoteBody.VoteFor = contractOutput.Id
	vote1.VoteBody.VoteType = "TRANSACTION"
	//note: contractoutput(transaction) node signatrue : use the contractOutput.id
	vote1.Signature = common.Sign(config.Config.Keypair.PrivateKey, contractOutput.Id)
	vote2 := &model.Vote{}
	vote2.Id = common.GenerateUUID()
	vote2.NodePubkey = config.Config.Keypair.PublicKey
	vote2.VoteBody.Timestamp = common.GenTimestamp()
	vote2.VoteBody.InvalidReason = "resion"
	vote2.VoteBody.IsValid = true
	vote2.VoteBody.VoteFor = contractOutput.Id
	vote2.VoteBody.VoteType = "TRANSACTION"
	vote2.Signature = common.Sign(config.Config.Keypair.PrivateKey, contractOutput.Id)
	vote3 := &model.Vote{}
	vote3.Id = common.GenerateUUID()
	vote3.NodePubkey = config.Config.Keypair.PublicKey
	vote3.VoteBody.Timestamp = common.GenTimestamp()
	vote3.VoteBody.InvalidReason = "resion"
	vote3.VoteBody.IsValid = true
	vote3.VoteBody.VoteFor = contractOutput.Id
	vote3.VoteBody.VoteType = "TRANSACTION"
	vote3.Signature = common.Sign(config.Config.Keypair.PrivateKey, contractOutput.Id)
	contractOutput.Transaction.Relation.Votes = []*model.Vote{
		vote1, vote2, vote3,
	}
	return contractOutput
}