package db

import (
	"fmt"
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
func (rethink *rethinkdb) _Insert(dbName, tableName, json string) (r.WriteResponse, error) {
	res, err := r.DB(dbName).Table(tableName).Insert(r.JSON(json)).RunWrite(rethink.session)
	if err != nil {
		return r.WriteResponse{}, err
	}
	return res, nil
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) _Delete(dbName, tableName, id string) (r.WriteResponse, error) {
	res, err := r.DB(dbName).Table(tableName).Get(id).Delete().RunWrite(rethink.session)
	if err != nil {
		return r.WriteResponse{}, err
	}
	return res, nil
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) _Update(dbName, tableName, id, json string) (r.WriteResponse, error) {
	res, err := r.DB(dbName).Table(tableName).Get(id).Update(r.JSON(json)).RunWrite(rethink.session)
	if err != nil {
		return r.WriteResponse{}, err
	}
	return res, nil
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) _Query(dbName, tableName, id string) (*r.Cursor, error) {
	res, err := r.DB(dbName).Table(tableName).Get(id).Run(rethink.session)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//---------------------------------------------------------------------------
// 插入一个nodepublickey的task方法
func (rethink *rethinkdb) InsertTaskSchedule(strTaskSchedule string) error {
	res, err := rethink._Insert(DATABASEB_NAME, TABLE_TASK_SCHEDULE, strTaskSchedule)
	if err != nil {
		return err
	}
	if res.Inserted >= 1 {
		return nil
	} else {
		return fmt.Errorf("insert failed")
	}
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
	res, err := r.DB(DATABASEB_NAME).
		Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("id").Eq(strID)).
		Run(rethink.session)
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

	startTime, ok := tasks["StartTime"].(string)
	if !ok {
		return "", "", fmt.Errorf("assert error")
	}

	endTime, ok := tasks["EndTime"].(string)
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
		res, err := rethink._Query(DATABASEB_NAME, TABLE_TASK_SCHEDULE, strID)
		if err != nil {
			return err
		}

		if res.IsNil() {
			return fmt.Errorf("null")
		}

		var task map[string]interface{}
		err = res.One(&task)
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

	res, err := rethink._Update(DATABASEB_NAME, TABLE_TASK_SCHEDULE, strID, strJSON)
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
// 设置OverFlag字段为1
func (rethink *rethinkdb) SetTaskScheduleOverFlag(strID string) error {
	strJSON := fmt.Sprintf("{\"OverFlag\":%d,\"LastExecuteTime\":\"%s\"}",
		1, common.GenTimestamp())

	res, err := rethink._Update(DATABASEB_NAME, TABLE_TASK_SCHEDULE, strID, strJSON)
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
// 设置TaskId,TaskState和TaskExecuteIndex字段的值
func (rethink *rethinkdb) SetTaskState(strID, strTaskId, strState string, nTaskExecuteIndex int) error {
	strJSON := fmt.Sprintf("{\"TaskId\":\"%s\",\"TaskState\":\"%s\",\"TaskExecuteIndex\":%d}",
		strTaskId, strState, nTaskExecuteIndex)

	res, err := rethink._Update(DATABASEB_NAME, TABLE_TASK_SCHEDULE, strID, strJSON)
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

	if res.Replaced|res.Unchanged >= 1 {

	} else {
		return fmt.Errorf("update failed")
	}

	strJSON := fmt.Sprintf("{\"LastExecuteTime\":\"%s\"}", common.GenTimestamp())

	res1, err := rethink._Update(DATABASEB_NAME, TABLE_TASK_SCHEDULE, strID, strJSON)
	if res1.Replaced|res1.Unchanged >= 1 {
		return nil
	} else {
		return fmt.Errorf("update failed")
	}
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
func (rethink *rethinkdb) GetTaskScheduleCount(stat string) (string, error) {
	res, err := r.DB(DATABASEB_NAME).
		Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field(stat).Ge(50)).
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
