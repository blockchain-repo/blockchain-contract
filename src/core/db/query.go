package db

import (
	"fmt"
	"unicontract/src/common"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
)

const dbname = "Unicontract"
const (
	table_contract         = "Contract"
	table_votes            = "Votes"
	table_contract_tasks   = "ContractTasks"
	table_consensus_fail   = "ConsensusFail"
	table_contract_outputs = "ContractOutputs"
)

// 根据合约[id]获取合约
func GetContractById(contractId string) string {
	res := rethinkdb.Get(dbname, table_contract, contractId)

	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		fmt.Printf("Error scanning database result: %s", err)
	}

	return common.Serialize(blo)
}

//根据合约[id]获取合约　处理主节点
func GetContractMainPubkeyById(contractId string) string {
	res := rethinkdb.Get(dbname, table_contract, contractId)

	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		fmt.Printf("Error scanning database result: %s", err)
	}

	if main_pubkey, ok := blo["main_pubkey"]; ok {
		pubkey, _ok := main_pubkey.(string)
		if _ok {
			return pubkey
		} else {
			return ""
		}
	} else {
		return ""
	}
}

// 根据合约[id]获取所有 vote
func GetVotesById(contractId string) string {
	res := rethinkdb.Get(dbname, table_votes, contractId)

	var blo map[string]interface{}
	err := res.All(&blo)
	if err != nil {
		fmt.Printf("Error scanning database result: %s", err)
	}

	return common.Serialize(blo)
}

func InsertContractStruct(contract model.ContractModel) bool {
	res := rethinkdb.Insert(dbname, table_contract, common.Serialize(contract))
	fmt.Printf("%d row inserted", res.Inserted)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

// contract serialize contract string
func InsertContract(contract string) bool {
	res := rethinkdb.Insert(dbname, table_contract, contract)
	fmt.Printf("%d row inserted", res.Inserted)
	if res.Inserted >= 1 {
		return true
	}
	return false
}
