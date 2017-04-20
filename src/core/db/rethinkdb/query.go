package rethinkdb

import (
	"log"

	"errors"
	"fmt"
	r "gopkg.in/gorethink/gorethink.v3"
	"unicontract/src/common"
)

func Get(db string, name string, id string) *r.Cursor {
	session := ConnectDB(db)
	res, err := r.Table(name).Get(id).Run(session)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return res
}

func Insert(db string, name string, jsonstr string) r.WriteResponse {
	session := ConnectDB(db)
	res, err := r.Table(name).Insert(r.JSON(jsonstr)).RunWrite(session)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return res
}

func Update(db string, name string, id string, jsonstr string) r.WriteResponse {
	session := ConnectDB(db)
	res, err := r.Table(name).Get(id).Update(r.JSON(jsonstr)).RunWrite(session)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return res
}

func Delete(db string, name string, id string) r.WriteResponse {
	session := ConnectDB(db)
	res, err := r.Table(name).Get(id).Delete().RunWrite(session)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return res
}

/*----------------------------unicontract ops-------------------------------------*/

// 根据合约[id]获取合约
func GetContractById(contractId string) (string, error) {
	res := Get(DBNAME, TABLE_CONTRACTS, contractId)

	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		fmt.Printf("Error scanning database result: %s", err)
		return "", errors.New(err.Error())
	}

	return common.Serialize(blo), nil
}

//根据合约[id]获取合约　处理主节点
func GetContractMainPubkeyById(contractId string) (string, error) {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACTS).Get(contractId).Field("main_pubkey").Run(session)
	if err != nil {
		log.Fatalf(err.Error())
		return "", errors.New(err.Error())
	}

	var blo string
	err = res.One(&blo)

	if err != nil {
		log.Fatalf(err.Error())
		//fmt.Printf("Error scanning database result: %s", err)
		return "", errors.New(err.Error())
	}
	return blo, nil
}

// 根据合约[id]获取所有 vote
func GetVotesByContractId(contractId string) (string, error) {

	if contractId == "" {
		return "", errors.New("contractId blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_VOTES).Filter(r.Row.Field("vote").Field("vote_for_contract").Eq(contractId)).Run(session)

	if err != nil {
		log.Fatalf(err.Error())
		return "", errors.New(err.Error())
	}

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		log.Fatalf(err.Error())
		//fmt.Printf("Error scanning database result: %s", err)
		return "", errors.New(err.Error())
	}
	return common.Serialize(blo), nil
}

// contract serialize contract string
func InsertContract(contract string) bool {
	if contract == "" {
		return false
	}

	res := Insert(DBNAME, TABLE_CONTRACTS, contract)
	//fmt.Printf("%d row inserted", res.Inserted)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

// vote serialize vote string
func InsertVote(vote string) bool {
	if vote == "" {
		return false
	}

	res := Insert(DBNAME, TABLE_VOTES, vote)
	//fmt.Printf("%d row inserted", res.Inserted)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

// contractOutput serialize contractOutput string
func InsertContractOutput(contractOutput string) bool {
	if contractOutput == "" {
		return false
	}

	res := Insert(DBNAME, TABLE_CONTRACT_OUTPUTS, contractOutput)
	//fmt.Printf("%d row inserted", res.Inserted)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

// vote serialize contract string
func InsertContractTask(vote string) bool {
	if vote == "" {
		return false
	}

	res := Insert(DBNAME, TABLE_VOTES, vote)
	//fmt.Printf("%d row inserted", res.Inserted)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

// vote serialize contract string
func InsertConsensusFail(vote string) bool {
	if vote == "" {
		return false
	}

	res := Insert(DBNAME, TABLE_VOTES, vote)
	//fmt.Printf("%d row inserted", res.Inserted)
	if res.Inserted >= 1 {
		return true
	}
	return false
}
