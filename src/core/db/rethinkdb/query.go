package rethinkdb

import (
	"log"

	"errors"
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

/*----------------------------- contracts start---------------------------------------*/
// contract serialize contract string
func InsertContract(contract string) bool {
	if contract == "" {
		return false
	}

	res := Insert(DBNAME, TABLE_CONTRACTS, contract)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

func GetContractById(id string) (string, error) {
	if id == "" {
		return "", errors.New("id blank")
	}
	res := Get(DBNAME, TABLE_CONTRACTS, id)
	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return common.Serialize(blo), nil
}

// 根据合约[id]获取具有相同contractId的合约
func GetContractsByContractId(contractId string) (string, error) {
	if contractId == "" {
		return "", errors.New("contractId blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACTS).Filter(r.Row.Field("ContractBody").Field("ContractId").Eq(contractId)).Run(session)

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return common.Serialize(blo), nil
}

//根据合约[id]获取合约　处理主节点
func GetContractMainPubkeyByContractId(contractId string) (string, error) {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACTS).Filter(r.Row.Field("ContractBody").Field("ContractId").Eq(contractId)).
		Field("ContractHead").Field("MainPubkey").Run(session)
	if err != nil {
		log.Fatalf(err.Error())
		return "", errors.New(err.Error())
	}

	var blo string
	err = res.One(&blo)

	if err != nil {
		log.Fatalf(err.Error())
		return "", errors.New(err.Error())
	}
	return blo, nil
}

/*----------------------------- contracts end---------------------------------------*/

/*----------------------------- votes start---------------------------------------*/
// vote serialize vote string
func InsertVote(vote string) bool {
	if vote == "" {
		return false
	}

	res := Insert(DBNAME, TABLE_VOTES, vote)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

func GetVoteById(id string) (string, error) {
	if id == "" {
		return "", errors.New("id blank")
	}
	res := Get(DBNAME, TABLE_VOTES, id)
	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return common.Serialize(blo), nil
}

// 根据合约[id]获取所有 vote
func GetVotesByContractId(contractId string) (string, error) {

	if contractId == "" {
		return "", errors.New("contractId blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_VOTES).Filter(r.Row.Field("Vote").Field("VoteForContract").Eq(contractId)).Run(session)

	if err != nil {
		log.Fatalf(err.Error())
		return "", errors.New(err.Error())
	}

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		log.Fatalf(err.Error())
		return "", errors.New(err.Error())
	}
	return common.Serialize(blo), nil
}

/*----------------------------- votes end---------------------------------------*/

/*----------------------------- contractOutputs start---------------------------------------*/
// contractOutput serialize contractOutput string
func InsertContractOutput(contractOutput string) bool {
	if contractOutput == "" {
		return false
	}

	res := Insert(DBNAME, TABLE_CONTRACT_OUTPUTS, contractOutput)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

// 根据合约[id]获取所有 contractOutput
func GetContractOutputByContractId(contractId string) (string, error) {

	if contractId == "" {
		return "", errors.New("contractId blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACT_OUTPUTS).Filter(r.Row.Field("transaction").Field("contracts").Field("id").Eq(contractId)).Run(session)

	if err != nil {
		log.Fatalf(err.Error())
		return "", errors.New(err.Error())
	}

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		log.Fatalf(err.Error())
		return "", errors.New(err.Error())
	}
	return common.Serialize(blo), nil
}

/*----------------------------- contractOutputs end---------------------------------------*/

/*----------------------------- contractTask start---------------------------------------*/
// contractTask serialize contractTask string
func InsertContractTask(contractTask string) bool {
	if contractTask == "" {
		return false
	}

	res := Insert(DBNAME, TABLE_CONTRACT_TASKS, contractTask)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

/*----------------------------- contractTask end---------------------------------------*/

/*----------------------------- consensusFailures start---------------------------------------*/
// consensusFailures serialize consensusFailures string
func InsertConsensusFailure(consensusFailures string) bool {
	if consensusFailures == "" {
		return false
	}

	res := Insert(DBNAME, TABLE_CONSENSUS_FAILURES, consensusFailures)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

func GetConsensusFailureById(id string) (string, error){
	if id == "" {
		return "", errors.New("id blank")
	}

	res := Get(DBNAME, TABLE_CONSENSUS_FAILURES, id)
	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return common.Serialize(blo), nil
}

func GetConsensusFailuresByConsensusId(consensusId string) (string, error){
	if consensusId == "" {
		return "", errors.New("consensusId blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONSENSUS_FAILURES).Filter(r.Row.Field("ConsensusId").Eq(consensusId)).Run(session)

	if err != nil {
		log.Fatalf(err.Error())
		return "", errors.New(err.Error())
	}

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		log.Fatalf(err.Error())
		return "", errors.New(err.Error())
	}
	return common.Serialize(blo), nil
}

/*----------------------------- consensusFailures end---------------------------------------*/

/*----------------------------- SendFailingRecords start---------------------------------------*/
func GetAllRecords(db string, name string) ([]string, error) {
	session := ConnectDB(db)
	ids, err := r.Table(name).Field("id").Run(session)
	if err != nil {
		log.Fatalf(err.Error())
	}
	var idlist []string
	err = ids.All(&idlist)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, errors.New(err.Error())
	}
	return idlist, nil
}

/*----------------------------- SendFailingRecords end---------------------------------------*/
