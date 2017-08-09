package db

import (
	"fmt"
	"strconv"
	"sync"
)

import (
	"unicontract/src/common"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/config"
)

import (
	r "gopkg.in/gorethink/gorethink.v3"
)

//---------------------------------------------------------------------------
type rethinkdb struct {
	session *r.Session
}

//---------------------------------------------------------------------------
const (
	DATABASEB_NAME      = "Unicontract"
	TABLE_TASK_SCHEDULE = "TaskSchedule"
)

var (
	rethinkInstance *rethinkdb
	once            sync.Once

	Tables = []string{
		TABLE_TASK_SCHEDULE,
	}
)

//---------------------------------------------------------------------------
func GetInstance() *rethinkdb {
	once.Do(func() {
		rethinkInstance = &rethinkdb{session: nil}
		rethinkInstance.session, _ = rethinkInstance.connect()
	})
	return rethinkInstance
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) connect() (*r.Session, error) {
	/*
		conf := config.ReadConfig(config.DevelopmentEnv)
		session, err := r.Connect(r.ConnectOpts{
			Address:    conf.DatabaseUrl,
			Database:   conf.DatabaseName,
			InitialCap: conf.DatabaseInitialCap,
			MaxOpen:    conf.DatabaseMaxOpen,
		})
	*/
	ip := config.Config.LocalIp
	session, err := r.Connect(r.ConnectOpts{
		Address: ip + ":28015",
	})

	if err != nil {
		return nil, err
	}
	return session, nil
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) InitDatabase() error {
	rethink.CreateDatabase(DATABASEB_NAME)

	for _, v := range Tables {
		rethink.CreateTable(DATABASEB_NAME, v)
	}

	return nil
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) CreateDatabase(dbName string) error {
	resp, err := r.DBCreate(dbName).RunWrite(rethink.session)
	if err != nil {
		return err
	}
	uniledgerlog.Info("%d DB created\n", resp.DBsCreated)
	return nil
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) CreateTable(dbName string, tableName string) error {
	resp, err := r.DB(dbName).TableCreate(tableName).RunWrite(rethink.session)
	if err != nil {
		return err
	}
	uniledgerlog.Info("%d table created\n", resp.TablesCreated)
	return nil
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) DropDatabase(dbName string) error {
	resp, err := r.DBDrop(dbName).RunWrite(rethink.session)
	if err != nil {
		return err
	}
	uniledgerlog.Info("%d DB dropped, %d tables dropped\n", resp.DBsDropped, resp.TablesDropped)
	return nil
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) Insert(dbName, tableName, json string) (bool, error) {
	res, err := r.DB(dbName).Table(tableName).Insert(r.JSON(json)).RunWrite(rethink.session)
	if err != nil {
		return false, err
	}
	if res.Inserted >= 1 {
		return true, err
	} else {
		return false, fmt.Errorf("insert failed")
	}
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) Delete(dbName, tableName, id string) (bool, error) {
	res, err := r.DB(dbName).Table(tableName).Get(id).Delete().RunWrite(rethink.session)
	if err != nil {
		return false, err
	}
	if res.Deleted >= 1 {
		return true, err
	} else {
		return false, fmt.Errorf("delete failed")
	}
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) Update(dbName, tableName, id, json string) (bool, error) {
	res, err := r.DB(dbName).Table(tableName).Get(id).Update(r.JSON(json)).RunWrite(rethink.session)
	if err != nil {
		return false, err
	}
	if res.Replaced >= 1 || res.Unchanged >= 1 {
		return true, err
	} else {
		return false, fmt.Errorf("update failed")
	}
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) Query(dbName, tableName, id string) (map[string]interface{}, error) {
	res, err := r.DB(dbName).Table(tableName).Get(id).Run(rethink.session)
	if err != nil {
		return nil, err
	}
	if res.IsNil() {
		return nil, fmt.Errorf("query result is null")
	}
	var task map[string]interface{}
	err = res.One(&task)
	if err != nil {
		return nil, err
	}
	return task, err
}

//---------------------------------------------------------------------------
// 插入一个nodepublickey的task方法
func (rethink *rethinkdb) InsertTaskSchedule(strTaskSchedule string) error {
	_, err := rethink.Insert(DATABASEB_NAME, TABLE_TASK_SCHEDULE, strTaskSchedule)
	return err
}

//---------------------------------------------------------------------------
// 插入task方法
func (rethink *rethinkdb) InsertTaskSchedules(slTaskSchedule []interface{}) (int, error) {
	res, err := r.DB(DATABASEB_NAME).Table(TABLE_TASK_SCHEDULE).Insert(slTaskSchedule).RunWrite(rethink.session)
	return res.Inserted, err
}

//---------------------------------------------------------------------------
// 根据nodePubkey和contractID获得表内ID
func (rethink *rethinkdb) GetID(strNodePubkey, strContractID, strContractHashId string) (string, error) {
	res, err := r.DB(DATABASEB_NAME).
		Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("ContractHashId").Eq(strContractHashId)).
		Filter(r.Row.Field("ContractId").Eq(strContractID)).
		Filter(r.Row.Field("NodePubkey").Eq(strNodePubkey)).
		Run(rethink.session)
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

	id, ok := tasks["id"].(string)
	if !ok {
		return "", fmt.Errorf("assert error")
	}

	return id, nil
}

//---------------------------------------------------------------------------
// 根据ID获取starttime和endtime
func (rethink *rethinkdb) GetValidTime(strID string) (string, string, error) {
	task, err := rethink.Query(DATABASEB_NAME, TABLE_TASK_SCHEDULE, strID)
	if err != nil {
		return "", "", err
	}

	startTime, ok := task["StartTime"].(string)
	if !ok {
		return "", "", fmt.Errorf("assert error")
	}

	endTime, ok := task["EndTime"].(string)
	if !ok {
		return "", "", fmt.Errorf("assert error")
	}

	return startTime, endTime, nil
}

//---------------------------------------------------------------------------
// 批量设置SendFlag字段，发送为1,未发送为0
func (rethink *rethinkdb) SetTaskScheduleFlagBatch(slID []interface{}, alreadySend bool) error {
	var strJSON string
	if alreadySend {
		strJSON = fmt.Sprintf("{\"SendFlag\":%d,\"OverFlag\":%d,\"LastExecuteTime\":\"%s\"}",
			1, 1, common.GenTimestamp())
	} else {
		strJSON = fmt.Sprintf("{\"SendFlag\":%d,\"LastExecuteTime\":\"%s\"}",
			0, common.GenTimestamp())
	}

	res, err := r.DB(DATABASEB_NAME).
		Table(TABLE_TASK_SCHEDULE).
		GetAll(slID...).
		Update(r.JSON(strJSON)).
		RunWrite(rethink.session)
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
func (rethink *rethinkdb) SetTaskScheduleFlag(strID string, alreadySend bool) error {
	var sendflag int
	if alreadySend {
		sendflag = 1
	} else {
		task, err := rethink.Query(DATABASEB_NAME, TABLE_TASK_SCHEDULE, strID)
		if err != nil {
			return err
		}

		overFlag, ok := task["OverFlag"].(float64)
		if !ok {
			return fmt.Errorf("assert error")
		}
		if overFlag != 1 {
			sendflag = 0
		} else {
			return nil
		}
	}

	strJSON := fmt.Sprintf("{\"SendFlag\":%d,\"LastExecuteTime\":\"%s\"}",
		sendflag, common.GenTimestamp())

	_, err := rethink.Update(DATABASEB_NAME, TABLE_TASK_SCHEDULE, strID, strJSON)
	return err
}

//---------------------------------------------------------------------------
// 设置OverFlag字段为1
func (rethink *rethinkdb) SetTaskScheduleOverFlag(strID string) error {
	strJSON := fmt.Sprintf("{\"OverFlag\":%d,\"LastExecuteTime\":\"%s\"}",
		1, common.GenTimestamp())

	_, err := rethink.Update(DATABASEB_NAME, TABLE_TASK_SCHEDULE, strID, strJSON)
	return err
}

//---------------------------------------------------------------------------
// 设置TaskId,TaskState和TaskExecuteIndex字段的值
func (rethink *rethinkdb) SetTaskState(strID, strTaskId, strState string, nTaskExecuteIndex int) error {
	strJSON := fmt.Sprintf("{\"TaskId\":\"%s\",\"TaskState\":\"%s\",\"TaskExecuteIndex\":%d}",
		strTaskId, strState, nTaskExecuteIndex)

	_, err := rethink.Update(DATABASEB_NAME, TABLE_TASK_SCHEDULE, strID, strJSON)
	return err
}

//---------------------------------------------------------------------------
// 设置FailedCount\SuccessCount\WaitCount字段加一
func (rethink *rethinkdb) SetTaskScheduleCount(strID string, flag int) error {
	var strFSW string
	if flag == 0 {
		strFSW = "SuccessCount"
	} else if flag == 1 {
		strFSW = "FailedCount"
	} else {
		strFSW = "WaitCount"
	}

	res, err := r.DB(DATABASEB_NAME).
		Table(TABLE_TASK_SCHEDULE).
		Get(strID).
		Update(map[string]interface{}{strFSW: r.Row.Field(strFSW).Add(1)}).
		RunWrite(rethink.session)

	if err != nil {
		return err
	}

	if res.Replaced >= 1 || res.Unchanged >= 1 {

	} else {
		return fmt.Errorf("update failed")
	}

	strJSON := fmt.Sprintf("{\"LastExecuteTime\":\"%s\"}", common.GenTimestamp())

	_, err = rethink.Update(DATABASEB_NAME, TABLE_TASK_SCHEDULE, strID, strJSON)
	return err
}

//---------------------------------------------------------------------------
// 获取所有未发送的任务，用于放在待执行队列中
func (rethink *rethinkdb) GetTaskSchedulesNoSend(strNodePubkey string, nThreshold int) (string, error) {
	now := common.GenTimestamp()
	res, err := r.DB(DATABASEB_NAME).
		Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("NodePubkey").Eq(strNodePubkey)).
		Filter(r.Row.Field("StartTime").Le(now)).
		Filter(r.Row.Field("EndTime").Ge(now)).
		Filter(r.Row.Field("FailedCount").Lt(nThreshold)).
		Filter(r.Row.Field("SendFlag").Eq(0)).
		Run(rethink.session)
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
func (rethink *rethinkdb) GetTaskSchedulesNoSuccess(strNodePubkey string, nThreshold int, flag int) (string, error) {
	var strCount string
	if flag == 0 {
		strCount = "FailedCount"
	} else if flag == 1 {
		strCount = "WaitCount"
	}

	res, err := r.DB(DATABASEB_NAME).
		Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("NodePubkey").Eq(strNodePubkey)).
		Filter(r.Row.Field(strCount).Ge(nThreshold)).
		Filter(r.Row.Field("SendFlag").Eq(0)).
		Run(rethink.session)
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
func (rethink *rethinkdb) GetTaskSchedulesSuccess(strNodePubkey string) (string, error) {
	if len(strNodePubkey) == 0 {
		return "", fmt.Errorf("pubkey is null")
	}

	res, err := r.DB(DATABASEB_NAME).
		Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("SuccessCount").Ge(1)).
		Filter(r.Row.Field("NodePubkey").Eq(strNodePubkey)).
		Run(rethink.session)
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
func (rethink *rethinkdb) DeleteTaskSchedules(slID []interface{}) (int, error) {
	res, err := r.DB(DATABASEB_NAME).
		Table(TABLE_TASK_SCHEDULE).
		GetAll(slID...).
		Delete().
		RunWrite(rethink.session)
	return res.Deleted, err
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) GetTaskScheduleCount(stat string, num int) (string, error) {
	res, err := r.DB(DATABASEB_NAME).
		Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field(stat).Ge(num)).
		Count().
		Run(rethink.session)
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

//---------------------------------------------------------------------------
func (rethink *rethinkdb) GetTaskSendFlagCount(stat int) (string, error) {
	res, err := r.DB(DATABASEB_NAME).
		Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("SendFlag").Eq(stat)).
		Count().
		Run(rethink.session)
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

//---------------------------------------------------------------------------
func (rethink *rethinkdb) GetTaskScheduleState(strContractID, strContractHashId string, failedThreshold, waitThreshold int) (RunState, error) {
	var state RunState
	var err error
	strPublicKey := config.Config.Keypair.PublicKey
	if len(strPublicKey) != 0 {
		// 根据contractid和hashid查询对应的id
		strID, err := rethink.GetID(strPublicKey, strContractID, strContractHashId)
		if err == nil {
			// 根据id查询出那条记录
			task, err := rethink.Query(DATABASEB_NAME, TABLE_TASK_SCHEDULE, strID)
			if err == nil {
				overFlag, _ := task["OverFlag"].(float64)
				sendFlag, _ := task["SendFlag"].(float64)
				startTime, _ := task["StartTime"].(string)
				endTime, _ := task["EndTime"].(string)
				successCount, _ := task["SuccessCount"].(float64)
				failedCount, _ := task["FailedCount"].(float64)
				waitCount, _ := task["WaitCount"].(float64)

				switch {
				case overFlag == 0 && testTime(startTime, endTime) == 0 && sendFlag == 1:
					state = WAIT_FOR_RUN
					break
				case overFlag == 0 && testTime(startTime, endTime) == 0:
					state = NORMAL
					break
				case overFlag == 1 && successCount == 1:
					state = ALREADY_RUN_SUCCESS
					break
				case overFlag == 0 && testTime(startTime, endTime) == -1:
					state = NO_ARRIVAL_TIME
					break
				case overFlag == 0 && testTime(startTime, endTime) == 1:
					state = OVER_TIME
					break
				case overFlag == 1 && failedCount > float64(failedThreshold):
					state = FAILED_TIMES_BEYOND
					break
				case overFlag == 1 && waitCount > float64(waitThreshold):
					state = WAIT_TIMES_BEYOND
					break
				default:
					state = NORMAL
					break
				}
			}
		}
	} else {
		err = fmt.Errorf("strPublicKey is not exist")
	}
	return state, err
}

func testTime(start, end string) int {
	var state int
	startTimeStamp, _ := strconv.Atoi(start)
	endTimeStamp, _ := strconv.Atoi(end)
	nowTimeStamp, _ := strconv.Atoi(common.GenTimestamp())

	switch {
	case nowTimeStamp < startTimeStamp:
		state = -1
		break
	case nowTimeStamp < startTimeStamp && nowTimeStamp > endTimeStamp:
		state = 0
		break
	case nowTimeStamp > endTimeStamp:
		state = 1
		break
	default:
		state = 0
	}
	return state
}

//---------------------------------------------------------------------------
