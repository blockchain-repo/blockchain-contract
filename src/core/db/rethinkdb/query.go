package rethinkdb

import (
	"errors"
	"fmt"

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

//---------------------------------------------------------------------------
func GetTaskSchedulesNoSend(strNodePubkey string) (string, error) {
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

//---------------------------------------------------------------------------
func GetTaskSchedulesSuccess() (string, error) {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("SuccessCount").Ge(1)).
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

//---------------------------------------------------------------------------
func _SetTaskScheduleFlag(strID string, alreadySend bool) error {
	if len(strID) == 0 {
		return fmt.Errorf("strID is null")
	}

	var strJSON string
	if alreadySend {
		strJSON = fmt.Sprintf("{\"SendFlag\":%d,\"LastExecuteTime\":\"%s\"}", 1, common.GenTimestamp())
	} else {
		strJSON = fmt.Sprintf("{\"SendFlag\":%d}", 0)
	}

	res := Update(DBNAME, TABLE_TASK_SCHEDULE, strID, strJSON)
	if res.Replaced|res.Unchanged >= 1 {
		return nil
	} else {
		return fmt.Errorf("update failed")
	}
}

func SetTaskScheduleSend(strID string) error {
	return _SetTaskScheduleFlag(strID, true)
}

func SetTaskScheduleNoSend(strID string) error {
	return _SetTaskScheduleFlag(strID, false)
}

//---------------------------------------------------------------------------
func _SetTaskScheduleCount(strID string, success bool) (int, error) {
	if len(strID) == 0 {
		return -1, fmt.Errorf("strID is null")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("id").Eq(strID)).Run(session)
	if err != nil {
		return -1, err
	}

	if res.IsNil() {
		return -1, fmt.Errorf("set is null")
	}

	var tasks map[string]interface{}
	err = res.One(&tasks)
	if err != nil {
		return -1, err
	}

	failedCount := tasks["FailedCount"].(float64)
	failedCount += 1

	successCount := tasks["SuccessCount"].(float64)
	successCount += 1

	var strJSON string
	if success {
		strJSON = fmt.Sprintf("{\"SuccessCount\":%f}", successCount)
	} else {
		strJSON = fmt.Sprintf("{\"FailedCount\":%f}", failedCount)
	}

	res1 := Update(DBNAME, TABLE_TASK_SCHEDULE, strID, strJSON)
	if res1.Replaced|res1.Unchanged >= 1 {
		if success {
			return int(successCount), nil
		} else {
			return int(failedCount), nil
		}

	} else {
		return -1, fmt.Errorf("update failed")
	}
}

func SetTaskScheduleFailedCount(strID string) (int, error) {
	return _SetTaskScheduleCount(strID, false)
}

func SetTaskScheduleSuccessCount(strID string) (int, error) {
	return _SetTaskScheduleCount(strID, true)
}

//---------------------------------------------------------------------------
func DeleteTaskSchedules(slID []string) (int, []error) {
	var nDeleteNum int
	var slerr []error
	for _, value := range slID {
		res := Delete(DBNAME, TABLE_TASK_SCHEDULE, value)
		if res.Deleted >= 1 {
			nDeleteNum += res.Deleted
		} else {
			slerr = append(slerr, fmt.Errorf(value))
		}
	}
	return nDeleteNum, slerr
}

/*----------------------------- TaskSchedule end---------------------------------------*/
