package db

import (
	"fmt"
)

import (
	"unicontract/src/common"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/config"
)

import (
	r "gopkg.in/gorethink/gorethink.v3"
)

const (
	DBNAME              = "Unicontract"
	TABLE_TASK_SCHEDULE = "TaskSchedule"
)

var Tables = []string{
	TABLE_TASK_SCHEDULE,
}

//---------------------------------------------------------------------------
type Rethinkdb struct {
}

//---------------------------------------------------------------------------
func (rethink Rethinkdb) Connect() (*r.Session, error) {
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
func (rethink Rethinkdb) ConnectDB(dbname string) (*r.Session, error) {
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
		Address:  ip + ":28015",
		Database: dbname,
	})

	if err != nil {
		return nil, err
	}
	return session, nil
}

//---------------------------------------------------------------------------
func (rethink Rethinkdb) InitDatabase() error {
	dbname := DBNAME
	rethink.CreateDatabase(dbname)

	for _, v := range Tables {
		rethink.CreateTable(dbname, v)
	}

	return nil
}

//---------------------------------------------------------------------------
func (rethink Rethinkdb) CreateDatabase(name string) error {
	session, err := rethink.Connect()
	if err != nil {
		return err
	}
	resp, err := r.DBCreate(name).RunWrite(session)
	if err != nil {
		return err
	}
	uniledgerlog.Info("%d DB created\n", resp.DBsCreated)
	return nil
}

//---------------------------------------------------------------------------
func (rethink Rethinkdb) CreateTable(db string, name string) error {
	session, err := rethink.ConnectDB(db)
	if err != nil {
		return err
	}
	resp, err := r.TableCreate(name).RunWrite(session)
	if err != nil {
		return err
	}
	uniledgerlog.Info("%d table created\n", resp.TablesCreated)
	return nil
}

//---------------------------------------------------------------------------
func (rethink Rethinkdb) DropDatabase() error {
	dbname := DBNAME
	session, err := rethink.Connect()
	if err != nil {
		return err
	}
	resp, err := r.DBDrop(dbname).RunWrite(session)
	if err != nil {
		return err
	}
	uniledgerlog.Info("%d DB dropped, %d tables dropped\n", resp.DBsDropped, resp.TablesDropped)
	return nil
}

//---------------------------------------------------------------------------
func (rethink Rethinkdb) _Insert(dbname, tablename, json string) (r.WriteResponse, error) {
	session, err := rethink.ConnectDB(dbname)
	if err != nil {
		return r.WriteResponse{}, err
	}
	res, err := r.Table(tablename).Insert(r.JSON(json)).RunWrite(session)
	if err != nil {
		return r.WriteResponse{}, err
	}
	return res, nil
}

//---------------------------------------------------------------------------
func (rethink Rethinkdb) _Delete(dbname, tablename, id string) (r.WriteResponse, error) {
	session, err := rethink.ConnectDB(dbname)
	if err != nil {
		return r.WriteResponse{}, err
	}
	res, err := r.Table(tablename).Get(id).Delete().RunWrite(session)
	if err != nil {
		return r.WriteResponse{}, err
	}
	return res, nil
}

//---------------------------------------------------------------------------
func (rethink Rethinkdb) _Update(dbname, tablename, id, json string) (r.WriteResponse, error) {
	session, err := rethink.ConnectDB(dbname)
	if err != nil {
		return r.WriteResponse{}, err
	}
	res, err := r.Table(tablename).Get(id).Update(r.JSON(json)).RunWrite(session)
	if err != nil {
		return r.WriteResponse{}, err
	}
	return res, nil
}

//---------------------------------------------------------------------------
func (rethink Rethinkdb) _Query(dbname, tablename, id string) (*r.Cursor, error) {
	session, err := rethink.ConnectDB(dbname)
	if err != nil {
		return nil, err
	}
	res, err := r.Table(tablename).Get(id).Run(session)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//---------------------------------------------------------------------------
// 插入一个nodepublickey的task方法
func (rethink Rethinkdb) InsertTaskSchedule(strTaskSchedule string) error {
	res, err := rethink._Insert(DBNAME, TABLE_TASK_SCHEDULE, strTaskSchedule)
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
func (rethink Rethinkdb) InsertTaskSchedules(slTaskSchedule []interface{}) (int, error) {
	session, err := rethink.ConnectDB(DBNAME)
	if err != nil {
		return 0, err
	}
	res, err := r.Table(TABLE_TASK_SCHEDULE).Insert(slTaskSchedule).RunWrite(session)
	return res.Inserted, err
}

//---------------------------------------------------------------------------
// 根据nodePubkey和contractID获得表内ID
func (rethink Rethinkdb) GetID(strNodePubkey, strContractID string, strContractHashId string) (string, error) {
	session, err := rethink.ConnectDB(DBNAME)
	if err != nil {
		return "", err
	}
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

	id, ok := tasks["id"].(string)
	if !ok {
		return "", fmt.Errorf("assert error")
	}

	return id, nil
}

//---------------------------------------------------------------------------
// 根据ID获取starttime和endtime
func (rethink Rethinkdb) GetValidTime(strID string) (string, string, error) {
	session, err := rethink.ConnectDB(DBNAME)
	if err != nil {
		return "", "", err
	}
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
func (rethink Rethinkdb) SetTaskScheduleFlagBatch(slID []interface{}, alreadySend bool) error {
	var strJSON string
	if alreadySend {
		strJSON = fmt.Sprintf("{\"SendFlag\":%d,\"OverFlag\":%d,\"LastExecuteTime\":\"%s\"}",
			1, 1, common.GenTimestamp())
	} else {
		strJSON = fmt.Sprintf("{\"SendFlag\":%d,\"LastExecuteTime\":\"%s\"}",
			0, common.GenTimestamp())
	}

	session, err := rethink.ConnectDB(DBNAME)
	if err != nil {
		return err
	}
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
func (rethink Rethinkdb) SetTaskScheduleFlag(strID string, alreadySend bool) error {
	var sendflag int
	if alreadySend {
		sendflag = 1
	} else {
		res, err := rethink._Query(DBNAME, TABLE_TASK_SCHEDULE, strID)
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

	res, err := rethink._Update(DBNAME, TABLE_TASK_SCHEDULE, strID, strJSON)
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
func (rethink Rethinkdb) SetTaskScheduleOverFlag(strID string) error {
	strJSON := fmt.Sprintf("{\"OverFlag\":%d,\"LastExecuteTime\":\"%s\"}",
		1, common.GenTimestamp())

	res, err := rethink._Update(DBNAME, TABLE_TASK_SCHEDULE, strID, strJSON)
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
func (rethink Rethinkdb) SetTaskState(strID, strTaskId, strState string, nTaskExecuteIndex int) error {
	strJSON := fmt.Sprintf("{\"TaskId\":\"%s\",\"TaskState\":\"%s\",\"TaskExecuteIndex\":%d}",
		strTaskId, strState, nTaskExecuteIndex)

	res, err := rethink._Update(DBNAME, TABLE_TASK_SCHEDULE, strID, strJSON)
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
func (rethink Rethinkdb) SetTaskScheduleCount(strID string, flag int) error {
	var strFSW string
	if flag == 0 {
		strFSW = "SuccessCount"
	} else if flag == 1 {
		strFSW = "FailedCount"
	} else {
		strFSW = "WaitCount"
	}

	session, err := rethink.ConnectDB(DBNAME)
	if err != nil {
		return err
	}
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

	res1, err := rethink._Update(DBNAME, TABLE_TASK_SCHEDULE, strID, strJSON)
	if res1.Replaced|res1.Unchanged >= 1 {
		return nil
	} else {
		return fmt.Errorf("update failed")
	}
}

//---------------------------------------------------------------------------
// 获取所有未发送的任务，用于放在待执行队列中
func (rethink Rethinkdb) GetTaskSchedulesNoSend(strNodePubkey string, nThreshold int) (string, error) {
	now := common.GenTimestamp()
	session, err := rethink.ConnectDB(DBNAME)
	if err != nil {
		return "", err
	}
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
func (rethink Rethinkdb) GetTaskSchedulesNoSuccess(strNodePubkey string, nThreshold int, flag int) (string, error) {
	var strCount string
	if flag == 0 {
		strCount = "FailedCount"
	} else if flag == 1 {
		strCount = "WaitCount"
	}

	session, err := rethink.ConnectDB(DBNAME)
	if err != nil {
		return "", err
	}
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
func (rethink Rethinkdb) GetTaskSchedulesSuccess(strNodePubkey string) (string, error) {
	if len(strNodePubkey) == 0 {
		return "", fmt.Errorf("pubkey is null")
	}

	session, err := rethink.ConnectDB(DBNAME)
	if err != nil {
		return "", err
	}
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
func (rethink Rethinkdb) DeleteTaskSchedules(slID []interface{}) (int, error) {
	session, err := rethink.ConnectDB(DBNAME)
	if err != nil {
		return 0, err
	}
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		GetAll(slID...).Delete().RunWrite(session)
	return res.Deleted, err
}

//---------------------------------------------------------------------------
