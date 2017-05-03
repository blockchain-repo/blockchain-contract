package rethinkdb

import (
	"fmt"
	"errors"

	"github.com/astaxie/beego/logs"
	r "gopkg.in/gorethink/gorethink.v3"

	"unicontract/src/common"
)

func Get(db string, name string, id string) *r.Cursor {
	session := ConnectDB(db)
	res, err := r.Table(name).Get(id).Run(session)
	if err != nil {
		logs.Error(err.Error())
	}
	return res
}

func Insert(db string, name string, jsonstr string) r.WriteResponse {
	session := ConnectDB(db)
	res, err := r.Table(name).Insert(r.JSON(jsonstr)).RunWrite(session)
	if err != nil {
		logs.Error(err.Error())
	}
	return res
}

func Update(db string, name string, id string, jsonstr string) r.WriteResponse {
	session := ConnectDB(db)
	res, err := r.Table(name).Get(id).Update(r.JSON(jsonstr)).RunWrite(session)
	if err != nil {
		logs.Error(err.Error())
	}
	return res
}

func Delete(db string, name string, id string) r.WriteResponse {
	session := ConnectDB(db)
	res, err := r.Table(name).Get(id).Delete().RunWrite(session)
	if err != nil {
		logs.Error(err.Error())
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
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		return "", err
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
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		return "", err
	}

	return common.Serialize(blo), nil
}

//根据 contract.id 获取合约处理主节点
func GetContractMainPubkeyByContract(id string) (string, error) {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACTS).Get(id).Count().Default(0).Run(session)
	if err != nil {
		return "", err
	}

	var blo int
	err = res.One(&blo)
	if err != nil {
		return "", err
	}
	if blo == 0 {
		return "", nil
	}

	// continue ...
	res, err = r.Table(TABLE_CONTRACTS).Get(id).Field("ContractHead").Field("MainPubkey").Run(session)
	if err != nil {
		return "", err
	}

	var blo2 string
	err = res.One(&blo2)
	if err != nil {
		return "", err
	}
	return blo2, nil
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
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

// 根据合约[id]获取所有 vote
func GetVotesByContractId(contractId string) (string, error) {

	if contractId == "" {
		return "", errors.New("contractId blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_VOTES).Filter(r.Row.Field("Vote").Field("VoteFor").Eq(contractId)).Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		return "", err
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

func GetContractOutputById(id string) (string, error) {
	if id == "" {
		return "", errors.New("id blank")
	}

	res := Get(DBNAME, TABLE_CONTRACT_OUTPUTS, id)
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

// 根据合约 整体的[id] transaction.contract.id 或者 relation.contractId 获取所有 contractOutput
func GetContractOutputByContractPrimaryId(id string) (string, error) {

	if id == "" {
		return "", errors.New("id blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACT_OUTPUTS).Filter(r.Row.Field("transaction").Field("Contract").Field("id").Eq(id)).Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err = res.One(&blo)
	if err != nil {
		return "", err
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

func GetContractTaskById(id string) (string, error) {
	if id == "" {
		return "", errors.New("id blank")
	}

	res := Get(DBNAME, TABLE_CONTRACT_TASKS, id)
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

func GetContractTasksByContractId(contractId string) (string, error) {
	if contractId == "" {
		return "", errors.New("contractId blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACT_TASKS).Filter(r.Row.Field("ContractId").Eq(contractId)).Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		logs.Error(err.Error())
		return "", err
	}
	return common.Serialize(blo), nil
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

func GetConsensusFailureById(id string) (string, error) {
	if id == "" {
		return "", errors.New("id blank")
	}

	res := Get(DBNAME, TABLE_CONSENSUS_FAILURES, id)
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

func GetConsensusFailuresByConsensusId(consensusId string) (string, error) {
	if consensusId == "" {
		return "", errors.New("consensusId blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONSENSUS_FAILURES).Filter(r.Row.Field("ConsensusId").Eq(consensusId)).Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		logs.Error(err.Error())
		return "", err
	}
	return common.Serialize(blo), nil
}

/*----------------------------- consensusFailures end---------------------------------------*/

/*----------------------------- SendFailingRecords start---------------------------------------*/
func GetAllRecords(db string, name string) ([]string, error) {
	session := ConnectDB(db)
	ids, err := r.Table(name).Field("id").Run(session)
	if err != nil {
		logs.Error(err.Error())
	}
	var idlist []string
	err = ids.All(&idlist)
	if err != nil {
		logs.Error(err.Error())
		return nil, errors.New(err.Error())
	}
	return idlist, nil
}

/*----------------------------- SendFailingRecords end---------------------------------------*/

/*----------------------------- TaskSchedule start---------------------------------------*/
func InsertTaskSchedule(strTaskSchedule string) error {
	if len(strTaskSchedule) == 0 {
		return fmt.Errorf("strTaskSchedule is null")
	}

	res := Insert(DBNAME, TABLE_TASK_SCHEDULE, strTaskSchedule)
	if res.Inserted >= 1 {
		return nil
	} else {
		return fmt.Errorf("insert failed")
	}
}

func GetTaskSchedules(strNodePubkey string) (string, error) {
	if len(strNodePubkey) == 0 {
		return "", fmt.Errorf("strNodePubkey is null")
	}

	now := common.GenTimestamp()
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("NodePubkey").Eq(strNodePubkey)).
		Filter(r.Row.Field("StartTime").Le(now)).
		Filter(r.Row.Field("EndTime").Ge(now)).
		Filter(r.Row.Field("SendFlag").Eq(0)).
		Run(session)
	if err != nil {
		return "", err
	}

	if res.IsNil() {
		return "", nil
	}

	var tasks []map[string]interface{}
	err = res.All(&tasks)
	if err != nil {
		return "", err
	}
	return common.Serialize(tasks), nil
}

func SetTaskScheduleSend(strID string) error {
	if len(strID) == 0 {
		return fmt.Errorf("strID is null")
	}

	res := Update(DBNAME, TABLE_TASK_SCHEDULE, strID, "{\"SendFlag\":1}")
	if res.Replaced|res.Unchanged >= 1 {
		return nil
	} else {
		return fmt.Errorf("update failed")
	}
}

func SetTaskScheduleNoSend(strID string) error {
	if len(strID) == 0 {
		return fmt.Errorf("strID is null")
	}

	res := Update(DBNAME, TABLE_TASK_SCHEDULE, strID, "{\"SendFlag\":0}")
	if res.Replaced|res.Unchanged >= 1 {
		return nil
	} else {
		return fmt.Errorf("update failed")
	}
}

/*----------------------------- TaskSchedule end---------------------------------------*/
