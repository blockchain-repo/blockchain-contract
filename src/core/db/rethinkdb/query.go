package rethinkdb

import (
	"errors"
	"fmt"

	"unicontract/src/common"

	"github.com/astaxie/beego/logs"
	r "gopkg.in/gorethink/gorethink.v3"
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

func DeleteContract(id string) bool {
	if len(id) == 0 {
		return false
	}

	res := Delete(DBNAME, TABLE_CONTRACTS, id)
	if res.Deleted >= 1 {
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

//根据传入条件查询 execute 合约 仅取出一条 , ContractState = Contract_Signature
func GetOneContractByMapCondition(conditions map[string]interface{}) (string, error) {
	contractId, _ := conditions["contractId"].(string)
	if contractId == "" {
		return "", errors.New("contractId blank")
	}
	owner, _ := conditions["owner"].(string)
	//if owner == "" {
	//	return "", errors.New("owner blank")
	//}
	//status, ok := conditions["status"].(string)
	//logs.Warn(ok)
	session := ConnectDB(DBNAME)
	var res *r.Cursor
	var err error
	if owner != "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq("Contract_Signature")).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId").Eq(contractId)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Run(session)
	} else {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq("Contract_Signature")).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId").Eq(contractId)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Run(session)
	}

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

// 根据传入条件查询 execute 合约, ContractState = Contract_Signature
func GetContractsByMapCondition(conditions map[string]interface{}) (string, error) {
	//status, _ := conditions["status"].(string)
	contractId, _ := conditions["contractId"].(string)
	owner, _ := conditions["owner"].(string)
	//if owner == "" {
	//	return "", errors.New("owner blank")
	//}
	//contractName, _ := conditions["contractName"].(string)
	//contractName = strings.Trim(contractName, " ")
	contractState := "Contract_Signature"
	session := ConnectDB(DBNAME)
	var res *r.Cursor
	var err error

	if contractId == "" && owner == "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId").Eq(contractId)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Run(session)
	} else if contractId == "" && owner != "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Run(session)
	} else if contractId != "" && owner == "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId").Eq(contractId)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Run(session)
	} else {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Run(session)
	}

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
	logs.Warn(blo)
	return common.Serialize(blo), nil
}

// 根据传入条件查询合约
func GetContractsLogByMapCondition(conditions map[string]interface{}) (string, error) {
	contractId, _ := conditions["contractId"].(string)
	//owner, _ := conditions["owner"].(string)
	//if owner == "" {
	//	return "", errors.New("owner blank")
	//}
	if contractId == "" {
		return "", errors.New("contractId blank")
	}
	session := ConnectDB(DBNAME)
	var res *r.Cursor
	var err error
	if contractId != "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			//Filter(r.Row.Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId").Eq(contractId)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Run(session)

	}
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

func GetNoConsensusContracts(time string, state int) (string, error) {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACTS).
		Filter(r.Row.Field("ContractHead").Field("ConsensusResult").Eq(state)).
		Filter(r.Row.Field("ContractHead").Field("AssignTime").Le(time)).
		Run(session)
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

func SetContractConsensusResultById(id string, time string, ConsensusResult int) error {
	strJSON := fmt.Sprintf("{\"ContractHead\":{\"AssignTime\":\"%s\",\"ConsensusResult\":%d}}",
		time, ConsensusResult)

	res := Update(DBNAME, TABLE_CONTRACTS, id, strJSON)
	if res.Replaced|res.Unchanged >= 1 {
		return nil
	} else {
		return fmt.Errorf("update failed")
	}
}

// 获取合约个数
func GetContractsCount() (string, error) {

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACTS).Count().Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo string
	err = res.One(&blo)
	if err != nil {
		return "", err
	}
	return blo, nil
}

// 获取合约不同状态个数
func GetContractStatsCount(stat string) (string, error) {

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACTS).
		Filter(r.Row.Field("ContractBody").Field("ContractState").Eq(stat)).
		Count().Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo string
	err = res.One(&blo)
	if err != nil {
		return "", err
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

// 删除一系列id的vote
func DeleteVotes(slID []interface{}) (int, error) {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_VOTES).
		GetAll(slID...).Delete().RunWrite(session)
	return res.Deleted, err
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

func GetConsensusFailuresCount() (int, error) {
	session := ConnectDB(DBNAME)
	count, err := r.Table(TABLE_CONSENSUS_FAILURES).Count().Run(session)
	if err != nil {
		return -1, err
	}
	if count.IsNil() {
		return -1, nil
	}

	var blo int
	err = count.One(&blo)
	if err != nil {
		logs.Error(err.Error())
		return -1, err
	}
	return blo, nil
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

func GetSendFailingRecordsCount() (int, error) {
	session := ConnectDB(DBNAME)
	count, err := r.Table(TABLE_SEND_FAILING_RECORDS).Count().Run(session)
	if err != nil {
		return -1, err
	}
	if count.IsNil() {
		return -1, nil
	}

	var blo int
	err = count.One(&blo)
	if err != nil {
		logs.Error(err.Error())
		return -1, err
	}
	return blo, nil
}

/*----------------------------- SendFailingRecords end---------------------------------------*/

/*TaskSchedule start-------------------------------------------------------*/
// 插入一个nodepublickey的task方法
func InsertTaskSchedule(strTaskSchedule string) error {
	res := Insert(DBNAME, TABLE_TASK_SCHEDULE, strTaskSchedule)
	if res.Inserted >= 1 {
		return nil
	} else {
		return fmt.Errorf("insert failed")
	}
}

//---------------------------------------------------------------------------
// 插入task方法
func InsertTaskSchedules(slTaskSchedule []interface{}) (int, error) {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).Insert(slTaskSchedule).RunWrite(session)
	return res.Inserted, err
}

//---------------------------------------------------------------------------
// 根据nodePubkey和contractID获得表内ID
func GetID(strNodePubkey, strContractID string, strContractHashId string) (string, error) {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("ContractHashId").Eq(strContractHashId)).
		Filter(r.Row.Field("ContractId").Eq(strContractID)).
		Filter(r.Row.Field("NodePubkey").Eq(strNodePubkey)).
		Run(session)
	if err != nil {
		return "", err
	}

	if res.IsNil() {
		return "", nil
	}

	var tasks map[string]interface{}
	err = res.One(&tasks)
	if err != nil {
		return "", err
	}

	return tasks["id"].(string), nil
}

//---------------------------------------------------------------------------
// 根据ID获取starttime和endtime
func GetValidTime(strID string) (string, string, error) {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("id").Eq(strID)).
		Run(session)
	if err != nil {
		return "", "", err
	}

	if res.IsNil() {
		return "", "", nil
	}

	var tasks map[string]interface{}
	err = res.One(&tasks)
	if err != nil {
		return "", "", err
	}

	return tasks["StartTime"].(string), tasks["EndTime"].(string), nil
}

//---------------------------------------------------------------------------
// 批量设置SendFlag字段，发送为1,未发送为0
func SetTaskScheduleFlagBatch(slID []interface{}, alreadySend bool) error {
	var strJSON string
	if alreadySend {
		strJSON = fmt.Sprintf("{\"SendFlag\":%d,\"OverFlag\":%d,\"LastExecuteTime\":\"%s\"}",
			1, 1, common.GenTimestamp())
	} else {
		strJSON = fmt.Sprintf("{\"SendFlag\":%d,\"LastExecuteTime\":\"%s\"}",
			0, common.GenTimestamp())
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		GetAll(slID...).Update(r.JSON(strJSON)).RunWrite(session)
	if err != nil {
		return err
	}
	if res.Replaced|res.Unchanged >= 1 {
		return nil
	} else {
		return fmt.Errorf("update failed")
	}
}

//---------------------------------------------------------------------------
// 设置SendFlag字段，发送为1,未发送为0
func SetTaskScheduleFlag(strID string, alreadySend bool) error {
	var sendflag int
	if alreadySend {
		sendflag = 1
	} else {
		res := Get(DBNAME, TABLE_TASK_SCHEDULE, strID)
		if res.IsNil() {
			return fmt.Errorf("null")
		}

		var task map[string]interface{}
		err := res.One(&task)
		if err != nil {
			return err
		}

		overFlag := task["OverFlag"].(float64)
		if overFlag != 1 {
			sendflag = 0
		} else {
			return nil
		}
	}

	strJSON := fmt.Sprintf("{\"SendFlag\":%d,\"LastExecuteTime\":\"%s\"}",
		sendflag, common.GenTimestamp())

	res := Update(DBNAME, TABLE_TASK_SCHEDULE, strID, strJSON)
	if res.Replaced|res.Unchanged >= 1 {
		return nil
	} else {
		return fmt.Errorf("update failed")
	}
}

//---------------------------------------------------------------------------
// 设置OverFlag字段为1
func SetTaskScheduleOverFlag(strID string) error {
	strJSON := fmt.Sprintf("{\"OverFlag\":%d,\"LastExecuteTime\":\"%s\"}",
		1, common.GenTimestamp())

	res := Update(DBNAME, TABLE_TASK_SCHEDULE, strID, strJSON)
	if res.Replaced|res.Unchanged >= 1 {
		return nil
	} else {
		return fmt.Errorf("update failed")
	}
}

//---------------------------------------------------------------------------
// 设置TaskId,TaskState和TaskExecuteIndex字段的值
func SetTaskState(strID, strTaskId, strState string, nTaskExecuteIndex int) error {
	strJSON := fmt.Sprintf("{\"TaskId\":\"%s\",\"TaskState\":\"%s\",\"TaskExecuteIndex\":%d}",
		strTaskId, strState, nTaskExecuteIndex)

	res := Update(DBNAME, TABLE_TASK_SCHEDULE, strID, strJSON)
	if res.Replaced|res.Unchanged >= 1 {
		return nil
	} else {
		return fmt.Errorf("update failed")
	}
}

//---------------------------------------------------------------------------
// 设置FailedCount\SuccessCount\WaitCount字段加一
func SetTaskScheduleCount(strID string, flag int) error {
	var strFSW string
	if flag == 0 {
		strFSW = "SuccessCount"
	} else if flag == 1 {
		strFSW = "FailedCount"
	} else {
		strFSW = "WaitCount"
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		Get(strID).
		Update(map[string]interface{}{strFSW: r.Row.Field(strFSW).Add(1)}).
		RunWrite(session)

	if err != nil {
		return err
	}

	if res.Replaced|res.Unchanged >= 1 {

	} else {
		return fmt.Errorf("update failed")
	}

	strJSON := fmt.Sprintf("{\"LastExecuteTime\":\"%s\"}", common.GenTimestamp())

	res = Update(DBNAME, TABLE_TASK_SCHEDULE, strID, strJSON)
	if res.Replaced|res.Unchanged >= 1 {
		return nil
	} else {
		return fmt.Errorf("update failed")
	}
}

//---------------------------------------------------------------------------
// 获取所有未发送的任务，用于放在待执行队列中
func GetTaskSchedulesNoSend(strNodePubkey string, nThreshold int) (string, error) {
	now := common.GenTimestamp()
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("NodePubkey").Eq(strNodePubkey)).
		Filter(r.Row.Field("StartTime").Le(now)).
		Filter(r.Row.Field("EndTime").Ge(now)).
		Filter(r.Row.Field("FailedCount").Lt(nThreshold)).
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
// 获取所有失败次数(等待次数)超过阈值的task
func GetTaskSchedulesNoSuccess(strNodePubkey string, nThreshold int, flag int) (string, error) {
	var strCount string
	if flag == 0 {
		strCount = "FailedCount"
	} else if flag == 1 {
		strCount = "WaitCount"
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("NodePubkey").Eq(strNodePubkey)).
		Filter(r.Row.Field(strCount).Ge(nThreshold)).
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
// 获取已经执行成功后的任务，用于清理数据
func GetTaskSchedulesSuccess(strNodePubkey string) (string, error) {
	if len(strNodePubkey) == 0 {
		return "", fmt.Errorf("pubkey is null")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("SuccessCount").Ge(1)).
		Filter(r.Row.Field("NodePubkey").Eq(strNodePubkey)).
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
// 删除一系列id的任务
func DeleteTaskSchedules(slID []interface{}) (int, error) {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		GetAll(slID...).Delete().RunWrite(session)
	return res.Deleted, err
}

//---------------------------------------------------------------------------
// 删除一系列id的任务
func GetTaskScheduleCount(stat string) (string, error) {

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field(stat).Ge(50)).
		Count().Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo string
	err = res.One(&blo)
	if err != nil {
		return "", err
	}
	return blo, nil
}

// 删除一系列id的任务
func GetTaskSendFlagCount(stat int) (string, error) {

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("SendFlag").Eq(stat)).
		Count().Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo string
	err = res.One(&blo)
	if err != nil {
		return "", err
	}
	return blo, nil
}

/*TaskSchedule end---------------------------------------------------------*/
