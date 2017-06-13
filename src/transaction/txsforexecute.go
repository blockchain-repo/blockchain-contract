package transaction

import (
	"encoding/json"
	"fmt"
	"sort"
	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"

	"github.com/astaxie/beego/logs"
)

func ExecuteCreate(tx_signers string, recipients [][2]interface{}, metadataStr string,
	relationStr string, contractStr string) (outputStr string, err error) {

	asset := GetAsset(tx_signers)
	metadata, relation, contract, err := GenModelByExecStr(metadataStr, relationStr, contractStr)
	ownerbefore := append([]string{}, tx_signers)
	output, _ := Create(ownerbefore, recipients, &metadata, asset, relation, contract)
	output = NodeSign(output)
	b := rethinkdb.InsertContractOutput(common.StructSerialize(output))
	logs.Info(b)
	return common.StructSerialize(common.Serialize(output)), err
}

func ExecuteFreeze(operation string, ownerbefore string, recipients [][2]interface{},
	metadataStr string, relationStr string, contractStr string) (outputStr string, err error) {
	asset := GetAsset(ownerbefore)
	metadata, relation, contract, err := GenModelByExecStr(metadataStr, relationStr, contractStr)

	output, err := Transfer(operation, ownerbefore, recipients, &metadata, asset, relation, contract)
	if err != nil {
		return "", err
	}
	logs.Info(err)
	logs.Info(output)
	output = NodeSign(output)
	b := rethinkdb.InsertContractOutput(common.StructSerialize(output))
	logs.Info(b)
	return common.StructSerialize(common.Serialize(output)), err
}

func ExecuteTransfer(operation string, ownerbefore string, recipients [][2]interface{},
	metadataStr string, relationStr string, contractStr string) (outputStr string, err error) {
	asset := GetAsset(ownerbefore)
	metadata, relation, contract, err := GenModelByExecStr(metadataStr, relationStr, contractStr)

	output, err := Transfer(operation, ownerbefore, recipients, &metadata, asset, relation, contract)
	return common.StructSerialize(output), err
}

func ExecuteTransferComplete(contractOutPut string, taskStatus string) (outputStr string, err error) {
	var contractModel model.ContractOutput
	err = json.Unmarshal([]byte(contractOutPut), &contractModel)
	taskId := contractModel.Transaction.Relation.TaskId

	UpdateTaskStauts(&contractModel, taskId, taskStatus)

	contractModel.Id = contractModel.GenerateId()
	contractModel = NodeSign(contractModel)

	b := rethinkdb.InsertContractOutput(common.StructSerialize(contractModel))
	logs.Info(b)
	return common.StructSerialize(contractModel), err
}

func ExecuteUnfreeze(operation string, ownerbefore string, recipients [][2]interface{},
	metadataStr string, relationStr string, contractStr string) (outputStr string, err error) {
	asset := GetAsset(ownerbefore)
	metadata, relation, contract, err := GenModelByExecStr(metadataStr, relationStr, contractStr)

	output, err := Transfer(operation, ownerbefore, recipients, &metadata, asset, relation, contract)
	logs.Info(err)
	logs.Info(output)
	if err != nil {
		return "", err
	}
	output = NodeSign(output)
	b := rethinkdb.InsertContractOutput(common.StructSerialize(output))
	logs.Info(b)
	return common.StructSerialize(common.Serialize(output)), err
}

func ExecuteInterim(metadataStr string, relationStr string, contractStr string) (outputStr string, err error) {
	metadata, relation, contract, err := GenModelByExecStr(metadataStr, relationStr, contractStr)
	output, _ := Interim(&metadata, relation, contract)
	return common.StructSerialize(output), err
}

func ExecuteInterimComplete(contractOutPut string, taskStatus string) (outputStr string, err error) {
	var contractModel model.ContractOutput
	err = json.Unmarshal([]byte(contractOutPut), &contractModel)
	taskId := contractModel.Transaction.Relation.TaskId

	UpdateTaskStauts(&contractModel, taskId, taskStatus)

	contractModel.Id = contractModel.GenerateId()
	contractModel = NodeSign(contractModel)

	b := rethinkdb.InsertContractOutput(common.StructSerialize(contractModel))
	logs.Info(b)
	if !b {
		err = fmt.Errorf("ExecuteInterimComplete fail!")
	}
	return common.StructSerialize(contractModel), err
}

func GenerateRelation(contractHashId string, contractId string, taskId string, taskIndex int) string {
	voters := config.GetAllPublicKey()
	sort.Strings(voters)
	logs.Info(voters)
	relation := model.Relation{
		ContractId:     contractId,
		ContractHashId: contractHashId,
		TaskId:         taskId,
		TaskExecuteIdx: taskIndex,
		Voters:         voters,
	}
	return common.Serialize(relation)
}

func UpdateTaskStauts(contractModel *model.ContractOutput, taskId string, taskStatus string) {
	task := contractModel.Transaction.ContractModel.ContractBody.ContractComponents
	for index, component := range task {
		if component.TaskId == taskId {
			task[index].State = taskStatus
		}
	}
	contractModel.Transaction.ContractModel.Id = common.HashData(common.StructSerialize(contractModel.Transaction.ContractModel.ContractBody))
	contractModel.Transaction.Relation.ContractHashId = contractModel.Transaction.ContractModel.Id
	contractModel.Id = common.HashData(common.StructSerialize(contractModel))
}

func ExecuteGetContract(contractId string) (string, error) {
	con, err := GetContractFromUnichain(contractId)
	return common.StructSerialize(con), err
}
