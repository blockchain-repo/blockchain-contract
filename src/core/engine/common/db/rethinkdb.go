package db

import (
	"fmt"
	"sync"
	//"time"
)

import (
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/config"
	"unicontract/src/core/engine"
)

import (
	r "gopkg.in/gorethink/gorethink.v3"
)

var scanEngineConf map[interface{}]interface{}

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
	scanEngineConf = engine.UCVMConf["ScanEngine"].(map[interface{}]interface{})
	rethinkdbInitialCap, _ := scanEngineConf["rethinkdbInitialCap"].(int)
	rethinkdbMaxOpen, _ := scanEngineConf["rethinkdbMaxOpen"].(int)
	ip := config.Config.LocalIp
	port := config.Config.Port
	session, err := r.Connect(r.ConnectOpts{
		Address:    ip + ":" + port,
		InitialCap: rethinkdbInitialCap,
		MaxOpen:    rethinkdbMaxOpen,
	})
	if err != nil {
		return nil, err
	}

	//count := 5
	//for count > 0 {
	//	if !session.IsConnected() {
	//		count--
	//		err = session.Reconnect()
	//		if err != nil {
	//			return nil, err
	//		}
	//	} else {
	//		break
	//	}
	//	time.Sleep(time.Millisecond * 500)
	//}

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
	return err
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) CreateTable(dbName string, tableName string) error {
	resp, err := r.DB(dbName).TableCreate(tableName).RunWrite(rethink.session)
	if err != nil {
		return err
	}
	uniledgerlog.Info("%d table created\n", resp.TablesCreated)
	return err
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) DropDatabase(dbName string) error {
	resp, err := r.DBDrop(dbName).RunWrite(rethink.session)
	if err != nil {
		return err
	}
	uniledgerlog.Info("%d DB dropped, %d tables dropped\n", resp.DBsDropped, resp.TablesDropped)
	return err
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
// 插入多个task方法
func (rethink *rethinkdb) InsertBatch(dbName, tableName string, slTaskSchedule []interface{}) (int, error) {
	res, err := r.DB(dbName).
		Table(tableName).
		Insert(slTaskSchedule).
		RunWrite(rethink.session)
	return res.Inserted, err
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
// 删除多个id的任务
func (rethink *rethinkdb) DeleteBatch(dbName, tableName string, slID []interface{}) (int, error) {
	res, err := r.DB(dbName).
		Table(tableName).
		GetAll(slID...).
		Delete().
		RunWrite(rethink.session)
	return res.Deleted, err
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
func (rethink *rethinkdb) UpdateBatch(dbName, tableName, json string, slID []interface{}) (bool, error) {
	res, err := r.DB(DATABASEB_NAME).
		Table(TABLE_TASK_SCHEDULE).
		GetAll(slID...).
		Update(r.JSON(json)).
		RunWrite(rethink.session)
	if err != nil {
		return false, err
	}
	if res.Replaced >= 1 || res.Unchanged >= 1 {
		return true, err
	} else {
		return false, fmt.Errorf("updateBatch failed")
	}
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) UpdateToAdd(strDBName, strTableName, strID, strField string, num int) (bool, error) {
	res, err := r.DB(strDBName).
		Table(strTableName).
		Get(strID).
		Update(map[string]interface{}{strField: r.Row.Field(strField).Add(num)}).
		RunWrite(rethink.session)

	if err != nil {
		return false, err
	}

	if res.Replaced >= 1 || res.Unchanged >= 1 {
		return true, err
	} else {
		return false, fmt.Errorf("updateToAdd failed")
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
func (rethink *rethinkdb) QueryByContractId(strDBName, strTableName, strNodePubkey, strContractID string) ([]map[string]interface{}, error) {
	res, err := r.DB(strDBName).
		Table(strTableName).
		Filter(r.Row.Field("ContractId").Eq(strContractID)).
		Filter(r.Row.Field("NodePubkey").Eq(strNodePubkey)).
		Run(rethink.session)
	if err != nil {
		return nil, err
	}

	if res.IsNil() {
		return nil, fmt.Errorf("query result is null")
	}

	var tasks []map[string]interface{}
	err = res.All(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, err
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) QueryByIdAndHashId(strDBName, strTableName, strNodePubkey, strContractID, strContractHashId string) (map[string]interface{}, error) {
	res, err := r.DB(strDBName).
		Table(strTableName).
		Filter(r.Row.Field("ContractHashId").Eq(strContractHashId)).
		Filter(r.Row.Field("ContractId").Eq(strContractID)).
		Filter(r.Row.Field("NodePubkey").Eq(strNodePubkey)).
		Run(rethink.session)
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
// 获取所有未发送的任务，用于放在待执行队列中
func (rethink *rethinkdb) GetTaskSchedulesNoSend(strDBName, strTableName, strNowTime, strNodePubkey string, nThreshold int) ([]map[string]interface{}, error) {
	res, err := r.DB(strDBName).
		Table(strTableName).
		Filter(r.Row.Field("NodePubkey").Eq(strNodePubkey)).
		Filter(r.Row.Field("StartTime").Le(strNowTime)).
		Filter(r.Row.Field("EndTime").Ge(strNowTime)).
		Filter(r.Row.Field("FailedCount").Lt(nThreshold)).
		Filter(r.Row.Field("SendFlag").Eq(0)).
		Run(rethink.session)
	if err != nil || res.IsNil() {
		return nil, err
	}

	var tasks []map[string]interface{}
	err = res.All(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

//---------------------------------------------------------------------------
// 获取所有失败次数(等待次数)超过阈值的task
func (rethink *rethinkdb) GetTaskSchedulesNoSuccess(strDBName, strTableName, strCount, strNodePubkey string, nThreshold int, flag int) ([]map[string]interface{}, error) {
	res, err := r.DB(strDBName).
		Table(strTableName).
		Filter(r.Row.Field("NodePubkey").Eq(strNodePubkey)).
		Filter(r.Row.Field(strCount).Ge(nThreshold)).
		Filter(r.Row.Field("SendFlag").Eq(0)).
		Run(rethink.session)
	if err != nil || res.IsNil() {
		return nil, err
	}

	var tasks []map[string]interface{}
	err = res.All(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, err
}

//---------------------------------------------------------------------------
// 获取已经执行成功后的任务，用于清理数据
func (rethink *rethinkdb) GetTaskSchedulesSuccess(strDBName, strTableName, strNodePubkey string) ([]map[string]interface{}, error) {
	res, err := r.DB(strDBName).
		Table(strTableName).
		Filter(r.Row.Field("SuccessCount").Ge(1)).
		Filter(r.Row.Field("NodePubkey").Eq(strNodePubkey)).
		Run(rethink.session)
	if err != nil || res.IsNil() {
		return nil, err
	}

	var tasks []map[string]interface{}
	err = res.All(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) GetTaskScheduleCount(strDBName, strTableName, stat string, num int) (string, error) {
	res, err := r.DB(strDBName).
		Table(strTableName).
		Filter(r.Row.Field(stat).Ge(num)).
		Count().
		Run(rethink.session)
	if err != nil || res.IsNil() {
		return "", err
	}

	var blo string
	err = res.One(&blo)
	if err != nil {
		return "", err
	}
	return blo, nil
}

//---------------------------------------------------------------------------
func (rethink *rethinkdb) GetTaskSendFlagCount(strDBName, strTableName string, stat int) (string, error) {
	res, err := r.DB(strDBName).
		Table(strTableName).
		Filter(r.Row.Field("SendFlag").Eq(stat)).
		Count().
		Run(rethink.session)
	if err != nil || res.IsNil() {
		return "", err
	}

	var blo string
	err = res.One(&blo)
	if err != nil {
		return "", err
	}
	return blo, nil
}

//---------------------------------------------------------------------------
