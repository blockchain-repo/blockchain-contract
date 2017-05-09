package transaction

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/logs"
	"unicontract/src/chain"
	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/model"
)

const (
	_GENESIS  = "GENESIS"
	_CREATE   = "CREATE"
	_TRANSFER = "TRANSFER"
	_CONTRACT = "CONTRACT"
	_FREEZE   = "FREEZE"
	_UNFREEZ  = "UNFREEZE"
	_VERSION  = 2
)

var ALLOWED_OPERATIONS = [4]string{_GENESIS, _CREATE, _TRANSFER, _CONTRACT}

//create asset
func Create(tx_signers []string, recipients [][2]interface{}, metadata *model.Metadata, asset model.Asset,
	relation model.Relation, contract model.ContractModel) (model.ContractOutput, error) {

	isFeeze := false
	operation := _CREATE
	version := _VERSION
	timestamp := common.GenTimestamp()
	//generate outputs
	outputs := []*model.ConditionsItem{}
	for index, recipient := range recipients {
		pubkey := recipient[0].(string)
		amount := recipient[1].(int)
		output := &model.ConditionsItem{}
		output.GenerateOutput(index, isFeeze, pubkey, amount)
		outputs = append(outputs, output)
	}
	//generate inputs
	inputs := []*model.Fulfillment{}
	input := &model.Fulfillment{}
	input.GenerateInput(tx_signers)
	inputs = append(inputs, input)

	contractOutput := model.ContractOutput{}
	contractOutput.GenerateConOutput(operation, asset, inputs, outputs, metadata, timestamp, version, relation, contract)
	return contractOutput, nil
}

//transfer asset include:transfer/freeze/unfreeze
func Transfer(operation string, ownerbefore string, recipients [][2]interface{}, metadata *model.Metadata, asset model.Asset,
	relation model.Relation, contract model.ContractModel) (model.ContractOutput, error) {

	contractId := contract.ContractBody.ContractId
	isFeeze := false
	//generate inputs
	var inputs = []*model.Fulfillment{}
	var balance = 0
	var spentFlag = -1 //0:no asset was frozen;  1:the asset was frozen; 2:the frozen asset had transfer
	if operation == _FREEZE {
		isFeeze = true
		//generate inputs
		inputs, balance = GetUnfreezeUnspent(ownerbefore)
		//TODO check the owner ownerafter is the ownerbefore himself and only one ownerafter

	}

	if operation == _UNFREEZ {
		//note: I'm not sure whether I need to check the inputs is froozen or not

	}
	if operation == _TRANSFER {
		//generate inputs
		inputs, balance, spentFlag = GetFrozenUnspent(ownerbefore, contractId)
		if spentFlag == 0 {
			// TODO no asset was frozen, can't transfer asset

		} else if spentFlag == 2 {
			// TODO the frozen asset had transfer, no need to do this transfer

		}
	}

	//the operation in DB needs to be 'TRANSFER'
	operation = _TRANSFER
	version := _VERSION
	timestamp := common.GenTimestamp()

	//generate outputs
	outputs := []*model.ConditionsItem{}
	amounts := 0
	for index, recipient := range recipients {
		pubkey := recipient[0].(string)
		amount := recipient[1].(int)
		output := &model.ConditionsItem{}
		output.GenerateOutput(index, isFeeze, pubkey, amount)
		outputs = append(outputs, output)
		amounts += amount
	}

	if balance < amounts {
		err := errors.New("not enough asset to do the operation !!!")
		logs.Error(err)
		return model.ContractOutput{}, err
	} else if balance > amounts {
		pubkey := ownerbefore
		amount := balance - amounts
		output := &model.ConditionsItem{}
		output.GenerateOutput(1, false, pubkey, amount)
		outputs = append(outputs, output)
	}

	contractOutput := model.ContractOutput{}
	contractOutput.GenerateConOutput(operation, asset, inputs, outputs, metadata, timestamp, version, relation, contract)
	return contractOutput, nil
}

// the all unspent asset include 'freeze'/'unfreeze'
func GetAllUnspent(pubkey string, contractId string) {
	//TODO when it needed
}

// the unspent asset only include 'unfreeze'
func GetUnfreezeUnspent(pubkey string) (inps []*model.Fulfillment, bal int) {
	param := "unspent=true&public_key=" + pubkey
	result, err := chain.GetUnspentTxs(param)
	if err != nil {
		logs.Error(err.Error())
		return nil, 0
	}
	inputs := []*model.Fulfillment{}
	var balance int
	//logs.Info(result.Code)
	//logs.Info(result.Data)
	for index, unspend := range result.Data.([]interface{}) {
		unspenStruct := model.UnSpentOutput{}
		mapObjBytes, _ := json.Marshal(unspend)
		json.Unmarshal(mapObjBytes, &unspenStruct)
		logs.Info("unspend====", common.StructSerialize(unspenStruct))
		//generate input
		inoutLink := model.ContractOutputLink{
			Cid:  unspenStruct.Cid,
			Txid: unspenStruct.Txid,
		}
		ownerbefore := []string{pubkey}
		input := model.Fulfillment{
			Fid:          index,
			OwnersBefore: ownerbefore,
			Fulfillment:  "cf:4:RtTtCxNf1Bq7MFeIToEosMAa3v_jKtZUtqiWAXyFz1ejPMv-t7vT6DANcrYvKFHAsZblmZ1Xk03HQdJbGiMyb5CmQqGPHwlgKusNu9N_IDtPn7y16veJ1RBrUP-up4YD",
			Input:        &inoutLink,
		}
		inputs = append(inputs, &input)
		balance += unspenStruct.Amount
		//logs.Info("input====", common.StructSerialize(input))
	}
	if result.Code != 200 {
		logs.Error(errors.New("request send failed"))
		return nil, 0
	}
	return inputs, balance
}

//the unspent asset only include 'freeze'
func GetFrozenUnspent(pubkey string, contractId string) (inps []*model.Fulfillment, bal int, flag int) {
	param := "unspent=true&public_key=" + pubkey + "&contract_id=" + contractId
	result, err := chain.GetFreezeUnspentTxs(param)
	if err != nil {
		logs.Error(err.Error())
		return nil, 0, -1
	}
	inputs := []*model.Fulfillment{}
	var balance int
	//logs.Info(result.Code)
	//logs.Info(result.Data)
	for index, unspend := range result.Data.([]interface{}) {
		unspenStruct := model.UnSpentOutput{}
		mapObjBytes, _ := json.Marshal(unspend)
		json.Unmarshal(mapObjBytes, &unspenStruct)
		logs.Info("unspend====", common.StructSerialize(unspenStruct))
		//generate input
		inoutLink := model.ContractOutputLink{
			Cid:  unspenStruct.Cid,
			Txid: unspenStruct.Txid,
		}
		ownerbefore := []string{pubkey}
		input := model.Fulfillment{
			Fid:          index,
			OwnersBefore: ownerbefore,
			Fulfillment:  "cf:4:RtTtCxNf1Bq7MFeIToEosMAa3v_jKtZUtqiWAXyFz1ejPMv-t7vT6DANcrYvKFHAsZblmZ1Xk03HQdJbGiMyb5CmQqGPHwlgKusNu9N_IDtPn7y16veJ1RBrUP-up4YD",
			Input:        &inoutLink,
		}
		inputs = append(inputs, &input)
		balance += unspenStruct.Amount
		logs.Info("input====", common.StructSerialize(input))
	}
	if result.Code != 200 {
		logs.Error(errors.New("request send failed"))
		return nil, 0, -1
	}
	return inputs, balance, flag
}

func GetAsset(ownerbefore string) model.Asset {
	asset := model.Asset{}
	//TODO  get asset

	return asset
}

func GetContractFromUnichain(contractId string) model.ContractModel {
	param := `{"contract_id":"` + contractId + `"}`
	result, err := chain.GetContractById(param)
	if err != nil {
		logs.Error(err.Error())
		return model.ContractModel{}
	}
	contractStruct := model.ContractModel{}

	for index, contract := range result.Data.([]interface{}) {
		contractStruct = model.ContractModel{}
		contractBytes, _ := json.Marshal(contract)
		json.Unmarshal(contractBytes, &contractStruct)
		logs.Info("index=", index, "----contract=", common.StructSerialize(contractStruct))
	}
	return contractStruct
}

func NodeSign(contractOutput model.ContractOutput) model.ContractOutput {
	vote := &model.Vote{}
	vote.Id = common.GenerateUUID()
	vote.NodePubkey = config.Config.Keypair.PublicKey
	vote.VoteBody.Timestamp = common.GenTimestamp()
	vote.VoteBody.InvalidReason = "valid"
	vote.VoteBody.IsValid = true
	vote.VoteBody.VoteFor = contractOutput.Id
	vote.VoteBody.VoteType = "TRANSACTION"

	// note: contractoutput(transaction) node signatrue : use the contractOutput.id
	// TODO change the sign data
	vote.Signature = common.Sign(config.Config.Keypair.PrivateKey, contractOutput.Id)
	logs.Info(contractOutput.Id)
	logs.Info(config.Config.Keypair.PrivateKey)
	logs.Info(vote.Signature)
	//TODO update vote :find the index in voters.update the same place
	contractOutput.Transaction.Relation.Votes = []*model.Vote{
		vote,
	}
	return contractOutput
}
