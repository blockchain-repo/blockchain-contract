package transaction

import (
	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/model"
	"unicontract/src/chain"
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"errors"
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
	return contractOutput
}

//transfer asset include:transfer/freeze/unfreeze
func Transfer(operation string, ownerbefore string, recipients [][2]interface{}, metadata model.Metadata, asset model.Asset,
	relation model.Relation, contract model.ContractModel) (model.ContractOutput, error) {
	isFeeze := false
	if operation == _FREEZE {
		isFeeze = true
		//TODO check the inputs is frozen or not ?
	}

	if operation == _UNFREEZ {
		//note: I'm not sure whether I need to check the inputs is froozen or not
	}
	if operation == _TRANSFER {
		//TODO `contract` can only transfer the asset which was frozzen
	}
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

	//generate inputs
	inputs, balance := GetUnspent(ownerbefore)

	if balance < amounts {
		err := errors.New("not enough asset to do the operation !!!")
		logs.Error(err)
		return nil, err
	}

	contractOutput := model.ContractOutput{}
	contractOutput.GenerateConOutput(operation, asset, inputs, outputs, metadata, timestamp, version, relation, contract)
	return contractOutput, nil
}

func GetUnspent(pubkey string) (inps []*model.Fulfillment, bal int) {
	param := "unspent=true&public_key=" + pubkey
	result, err := chain.GetUnspentTxs(param)
	if err != nil {
		logs.Error(err.Error())
		return nil, 0
	}
	inputs := []*model.Fulfillment{}
	var balance int
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
			Input:        inoutLink,
		}
		inputs = append(inputs, &input)
		balance += unspenStruct.Amount
		logs.Info("input====", common.StructSerialize(input))
	}
	if result.Code != 200 {
		logs.Error(errors.New("request send failed"))
		return nil, 0
	}
	return inputs, balance
}

func GetAsset(ownerbefore string) model.Asset {
	asset := model.Asset{}
	//TODO  get asset




	return asset
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

	//TODO update vote :find the index in voters.update the same place
	contractOutput.Transaction.Relation.Votes = []*model.Vote{
		vote,
	}
	return contractOutput
}
