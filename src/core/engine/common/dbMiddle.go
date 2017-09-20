package common

import (
	"fmt"
	"strconv"
)

import (
	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/engine"
	"unicontract/src/core/engine/common/db"
)

var scanEngineConf map[interface{}]interface{}
var DBInf db.Datebase

//---------------------------------------------------------------------------
func Init() {
	scanEngineConf = engine.UCVMConf["ScanEngine"].(map[interface{}]interface{})
	dbname, _ := scanEngineConf["db"].(string)
	if dbname == "rethinkdb" {
		DBInf = db.GetInstance()
	}
}

//---------------------------------------------------------------------------
func InitDatabase() {
	DBInf.InitDatabase()
}

//---------------------------------------------------------------------------
func DropDatabase() {
	DBInf.DropDatabase(db.DATABASEB_NAME)
}

//---------------------------------------------------------------------------
// 插入一个task方法
func InsertTaskSchedule(strTaskSchedule string) error {
	_, err := DBInf.Insert(db.DATABASEB_NAME, db.TABLE_TASK_SCHEDULE, strTaskSchedule)
	return err
}

//---------------------------------------------------------------------------
// 插入多个task方法
func InsertTaskSchedules_(slID []interface{}) (int, error) {
	return DBInf.InsertBatch(db.DATABASEB_NAME, db.TABLE_TASK_SCHEDULE, slID)
}

//---------------------------------------------------------------------------
// 删除一系列id的任务
func DeleteTaskSchedules(slID []interface{}) (int, error) {
	return DBInf.DeleteBatch(db.DATABASEB_NAME, db.TABLE_TASK_SCHEDULE, slID)
}

//---------------------------------------------------------------------------
// 根据nodePubkeyco、ntractID和strContractHashId获得表内ID
func GetID(strNodePubkey, strContractID, strContractHashId string) (string, error) {
	task, err := DBInf.QueryByIdAndHashId(db.DATABASEB_NAME, db.TABLE_TASK_SCHEDULE,
		strNodePubkey, strContractID, strContractHashId)
	if err != nil {
		return "", err
	}

	id, ok := task["id"].(string)
	if !ok {
		return "", fmt.Errorf("assert error")
	}

	return id, nil
}

//---------------------------------------------------------------------------
// 根据nodePubkey和contractID获得表内ID
func GetIDs(strNodePubkey, strContractID string) ([]string, error) {
	tasks, err := DBInf.QueryByContractId(db.DATABASEB_NAME, db.TABLE_TASK_SCHEDULE, strNodePubkey, strContractID)
	if err != nil {
		return nil, err
	}

	var slID []string
	for index := range tasks {
		id, ok := tasks[index]["id"].(string)
		if !ok {
			return nil, fmt.Errorf("assert error")
		}
		slID = append(slID, id)
	}

	return slID, nil
}

//---------------------------------------------------------------------------
// 根据ID获取starttime和endtime
func GetValidTime(strID string) (string, string, error) {
	task, err := DBInf.Query(db.DATABASEB_NAME, db.TABLE_TASK_SCHEDULE, strID)
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
// 获取所有未发送的任务，用于放在待执行队列中
func GetTaskSchedulesNoSend(strNodePubkey string, nThreshold int) (string, error) {
	now := common.GenTimestamp()
	tasks, err := DBInf.GetTaskSchedulesNoSend(db.DATABASEB_NAME, db.TABLE_TASK_SCHEDULE, now, strNodePubkey, nThreshold)
	if err != nil || tasks == nil {
		return "", err
	}

	return common.Serialize(tasks), err
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

	tasks, err := DBInf.GetTaskSchedulesNoSuccess(db.DATABASEB_NAME, db.TABLE_TASK_SCHEDULE, strCount, strNodePubkey, nThreshold, flag)
	if err != nil || tasks == nil {
		return "", err
	}

	return common.Serialize(tasks), err
}

//---------------------------------------------------------------------------
// 获取已经执行成功后的任务，用于清理数据
func GetTaskSchedulesSuccess(strNodePubkey string) (string, error) {
	tasks, err := DBInf.GetTaskSchedulesSuccess(db.DATABASEB_NAME, db.TABLE_TASK_SCHEDULE, strNodePubkey)
	if err != nil || tasks == nil {
		return "", err
	}

	return common.Serialize(tasks), err
}

//---------------------------------------------------------------------------
// 获得task状态
func GetTaskScheduleState(strContractID, strContractHashId string, failedThreshold, waitThreshold int) (db.RunState, error) {
	var state db.RunState
	var err error
	strPublicKey := config.Config.Keypair.PublicKey
	if len(strPublicKey) != 0 {
		// 根据contractid和hashid查询对应的id
		strID, err := GetID(strPublicKey, strContractID, strContractHashId)
		if err == nil {
			// 根据id查询出那条记录
			task, err := DBInf.Query(db.DATABASEB_NAME, db.TABLE_TASK_SCHEDULE, strID)
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
					state = db.WAIT_FOR_RUN
					break
				case overFlag == 0 && testTime(startTime, endTime) == 0:
					state = db.NORMAL
					break
				case overFlag == 1 && successCount == 1:
					state = db.ALREADY_RUN_SUCCESS
					break
				case overFlag == 0 && testTime(startTime, endTime) == -1:
					state = db.NO_ARRIVAL_TIME
					break
				case overFlag == 0 && testTime(startTime, endTime) == 1:
					state = db.OVER_TIME
					break
				case overFlag == 1 && failedCount > float64(failedThreshold):
					state = db.FAILED_TIMES_BEYOND
					break
				case overFlag == 1 && waitCount > float64(waitThreshold):
					state = db.WAIT_TIMES_BEYOND
					break
				default:
					state = db.NORMAL
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
// 设置SendFlag字段，发送为1,未发送为0
func SetTaskScheduleFlag(strID string, alreadySend bool) error {
	var sendflag int
	if alreadySend {
		sendflag = 1
	} else {
		task, err := DBInf.Query(db.DATABASEB_NAME, db.TABLE_TASK_SCHEDULE, strID)
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

	_, err := DBInf.Update(db.DATABASEB_NAME, db.TABLE_TASK_SCHEDULE, strID, strJSON)
	return err
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

	_, err := DBInf.UpdateBatch(db.DATABASEB_NAME, db.TABLE_TASK_SCHEDULE, strJSON, slID)
	return err
}

//---------------------------------------------------------------------------
// 设置OverFlag字段为1
func SetTaskScheduleOverFlag(strID string) error {
	strJSON := fmt.Sprintf("{\"OverFlag\":%d,\"LastExecuteTime\":\"%s\"}",
		1, common.GenTimestamp())

	_, err := DBInf.Update(db.DATABASEB_NAME, db.TABLE_TASK_SCHEDULE, strID, strJSON)
	return err
}

//---------------------------------------------------------------------------
// 设置TaskId,TaskState和TaskExecuteIndex字段的值
func SetTaskState(strID, strTaskId, strState string, nTaskExecuteIndex int) error {
	strJSON := fmt.Sprintf("{\"TaskId\":\"%s\",\"TaskState\":\"%s\",\"TaskExecuteIndex\":%d}",
		strTaskId, strState, nTaskExecuteIndex)

	_, err := DBInf.Update(db.DATABASEB_NAME, db.TABLE_TASK_SCHEDULE, strID, strJSON)
	return err
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

	_, err := DBInf.UpdateToAdd(db.DATABASEB_NAME, db.TABLE_TASK_SCHEDULE, strID, strFSW, 1)
	if err != nil {
		return err
	}

	strJSON := fmt.Sprintf("{\"LastExecuteTime\":\"%s\"}", common.GenTimestamp())

	_, err = DBInf.Update(db.DATABASEB_NAME, db.TABLE_TASK_SCHEDULE, strID, strJSON)
	return err
}

//---------------------------------------------------------------------------
// 批量设置SendFlag字段，发送为1,未发送为0
func SetContractTerminateBatch(slID []interface{}) error {
	strJSON := fmt.Sprintf("{\"OverFlag\":%d,\"LastExecuteTime\":\"%s\"}",
		1, common.GenTimestamp())

	_, err := DBInf.UpdateBatch(db.DATABASEB_NAME, db.TABLE_TASK_SCHEDULE, strJSON, slID)
	return err
}

//---------------------------------------------------------------------------
