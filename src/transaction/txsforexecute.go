package transaction

import (
	"encoding/json"
	"fmt"
	"sort"
	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"

	"time"
	"unicontract/src/common/uniledgerlog"
)

func ExecuteCreate(tx_signers string, recipients [][2]interface{}, metadataStr string,
	relationStr string, contractStr string) (outputStr string, err error) {

	asset := GetAsset(tx_signers)
	metadata, relation, contract, err := GenModelByExecStr(metadataStr, relationStr, contractStr)
	ownerbefore := append([]string{}, tx_signers)
	output, _ := Create(ownerbefore, recipients, &metadata, asset, relation, contract)
	output, vote, index := NodeSign(output)
	b := MergeContractOutput(common.StructSerialize(output), output.Id, vote, index)
	_ = b
	//uniledgerlog.Info(b)
	return common.StructSerialize(common.Serialize(output)), err
}

func ExecuteFreeze(operation string, ownerbefore string, recipients [][2]interface{},
	metadataStr string, relationStr string, contractStr string) (outputStr string, err error) {
	asset := GetAsset(ownerbefore)
	metadata, relation, contract, err := GenModelByExecStr(metadataStr, relationStr, contractStr)
	//uniledgerlog.Info("==after: ", contract)

	output, err := Transfer(operation, ownerbefore, recipients, &metadata, asset, relation, contract)
	if err != nil {
		return "", err
	}
	//uniledgerlog.Info(err)
	uniledgerlog.Info(output)
	output, vote, index := NodeSign(output)
	b := MergeContractOutput(common.StructSerialize(output), output.Id, vote, index)
	uniledgerlog.Info(b)
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
	contractModel, vote, index := NodeSign(contractModel)

	b := MergeContractOutput(common.StructSerialize(contractModel), contractModel.Id, vote, index)
	uniledgerlog.Info(b)
	return common.StructSerialize(contractModel), err
}

func ExecuteUnfreeze(operation string, ownerbefore string, recipients [][2]interface{},
	metadataStr string, relationStr string, contractStr string) (outputStr string, err error) {
	asset := GetAsset(ownerbefore)
	metadata, relation, contract, err := GenModelByExecStr(metadataStr, relationStr, contractStr)

	output, err := Transfer(operation, ownerbefore, recipients, &metadata, asset, relation, contract)
	//uniledgerlog.Info(err)
	uniledgerlog.Info(output)
	if err != nil {
		return "", err
	}
	output, vote, index := NodeSign(output)
	b := MergeContractOutput(common.StructSerialize(output), output.Id, vote, index)
	uniledgerlog.Info(b)
	return common.StructSerialize(common.Serialize(output)), err
}

func ExecuteInterim(metadataStr string, relationStr string, contractStr string) (outputStr string, err error) {
	metadata, relation, contract, err := GenModelByExecStr(metadataStr, relationStr, contractStr)
	output, _ := Interim(&metadata, relation, contract)
	return common.StructSerialize(output), err
}

func ExecuteInterimComplete(contractOutPut string, taskStatus string, contractState string) (outputStr string, err error) {
	var contractModel model.ContractOutput
	err = json.Unmarshal([]byte(contractOutPut), &contractModel)
	taskId := contractModel.Transaction.Relation.TaskId

	UpdateContractState(&contractModel, contractState)
	UpdateTaskStauts(&contractModel, taskId, taskStatus)

	contractModel.Id = contractModel.GenerateId()
	contractModel, vote, index := NodeSign(contractModel)

	b := MergeContractOutput(common.StructSerialize(contractModel), contractModel.Id, vote, index)
	uniledgerlog.Info(b)
	if !b {
		err = fmt.Errorf("ExecuteInterimComplete fail!")
	}
	return common.StructSerialize(contractModel), err
}

func GenerateRelation(contractHashId string, contractId string, taskId string, taskIndex int) string {
	voters := config.GetAllPublicKey()
	sort.Strings(voters)
	uniledgerlog.Info(voters)
	relation := model.Relation{
		ContractId:     contractId,
		ContractHashId: contractHashId,
		TaskId:         taskId,
		TaskExecuteIdx: taskIndex,
		Voters:         voters,
	}
	return common.Serialize(relation)
}
func UpdateContractState(contractModel *model.ContractOutput, contractState string) {
	contractModel.Transaction.ContractModel.ContractBody.ContractState = contractState
	contractModel.Transaction.ContractModel.Id = common.HashData(common.StructSerialize(contractModel.Transaction.ContractModel.ContractBody))
	contractModel.Transaction.Relation.ContractHashId = contractModel.Transaction.ContractModel.Id
	contractModel.Id = common.HashData(common.StructSerialize(contractModel))
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

func GetPurchaseAmount(pubkey string) float64 {
	res, _ := rethinkdb.GetInfoByUser(pubkey)
	purchaseAmount, ok := res["purchaseAmount"].(float64)
	if !ok {
		purchaseAmount = float64(0)
	}
	yield, ok := res["yield"].(float64)
	if !ok {
		yield = float64(0)
	}
	countAmount := purchaseAmount + yield
	return countAmount
}

func GetInfoByUser(pubkey string) map[string]interface{} {
	res, _ := rethinkdb.GetInfoByUser(pubkey)
	uniledgerlog.Info(res)
	return res
}

//
func MergeContractOutput(contractModelStr string, id string, vote *model.Vote, index int) bool {
	b := rethinkdb.InsertContractOutput(contractModelStr)
	if !b {
		voteMap, err := common.StructToMap(vote)
		if err != nil {
			return false
		}
		b = rethinkdb.UpdateContractOutVote(id, voteMap, index)
		if !b {
			uniledgerlog.Error("Merge ContractOut Vote fail!")
		}
	}
	return true
}

//
func GetInterestCount(pubkey string) float64 {
	res, _ := rethinkdb.GetLastInterest(pubkey)
	ok := false
	countInterest := 0.0
	countYeild := 0.0
	firstPurchaseAmount := 0.0
	for _, r := range res {
		firstPurchaseAmount, ok = r["firstPurchaseAmount"].(float64)
		if !ok {
			firstPurchaseAmount = float64(0)
		}
		tmp, ok := r["yield"].(float64)
		if !ok {
			tmp = float64(0)
		}
		countInterest = countInterest + tmp
		countYeild = countYeild + tmp
	}
	uniledgerlog.Info(countInterest + countYeild + firstPurchaseAmount)

	return countInterest + countYeild + firstPurchaseAmount
}

func SaveEarnings(pubkey string, isRaise bool, rate float64, firstPurchaseAmount float64, interest float64, purchaseAmount float64, yeild float64) bool {
	res, _ := rethinkdb.GetInfoByUser(pubkey)
	delete(res, "id")
	format := "2006-01-02 15:04:05"
	//now, _ := time.Parse(format, "2017-06-18 15:19:58")
	now, _ := time.Parse(format, time.Now().Format(format))
	res["date"] = now
	res["isRaise"] = isRaise
	res["rate"] = rate
	res["firstPurchaseAmount"] = firstPurchaseAmount
	res["interest"] = interest
	res["purchaseAmount"] = purchaseAmount
	res["yeild"] = yeild

	uniledgerlog.Info(res)
	str := common.Serialize(res)
	uniledgerlog.Info(str)
	result := rethinkdb.InsertInterestCount(str)

	return result
}
