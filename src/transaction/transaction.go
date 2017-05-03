package transaction

import (
	"unicontract/src/common"
	"unicontract/src/core/model"
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
	for _, recipient := range recipients {
		pubkey := recipient[0].(string)
		amount := recipient[1].(int)
		outputs = append(outputs, model.GenerateOutput(isFeeze, pubkey, amount))
	}

	//generate inputs
	inputs := []*model.Fulfillment{}
	inputs = append(inputs, model.GenerateInput(tx_signers))

	contractOutput := model.GenerateConOutput(operation, asset, inputs, outputs, metadata, timestamp, version, relation, contract)
	return contractOutput
}

//transfer asset
func Transfer(operation string, inputs []model.Fulfillment, recipients [][2]interface{}, metadata model.Metadata, asset model.Asset,
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
	for _, recipient := range recipients {
		pubkey := recipient[0].(string)
		amount := recipient[1].(int)
		outputs = append(outputs, model.GenerateOutput(isFeeze, pubkey, amount))
	}

	contractOutput := model.GenerateConOutput(operation, asset, inputs, outputs, metadata, timestamp, version, relation, contract)
	return contractOutput
}
