/*************************************************************************
  > File Name: taskDBHandle.go
  > Module:
  > Function: 提供对TaskSchedule表的一些操作函数
  > Author: wangyp
  > Company:
  > Department:
  > Mail: wangyepeng87@163.com
  > Created Time: 2017.05.10
 ************************************************************************/
package common

import (
	"encoding/json"
	"fmt"
)

import (
	"unicontract/src/common"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/config"
	"unicontract/src/core/engine"
	"unicontract/src/core/engine/common/db"
)

//---------------------------------------------------------------------------
// FailStruct
type UpdateMonitorFailStruct struct {
	FstrContractID     string
	FstrContractHashID string
	FstrTaskId         string
	FstrTaskState      string
	FnTaskExecuteIndex int
}

// WaitStruct
type UpdateMonitorWaitStruct struct {
	WstrContractID     string
	WstrContractHashID string
	WstrTaskId         string
	WstrTaskState      string
	WnTaskExecuteIndex int
}

// SuccStruct
type UpdateMonitorSuccStruct struct {
	SstrContractID        string
	SstrContractHashIdOld string
	SstrTaskStateOld      string
	SstrTaskIdOld         string
	SnTaskExecuteIndexOld int
	SstrContractHashIDNew string
	SstrTaskIdNew         string
	SstrTaskStateNew      string
	SnTaskExecuteIndexNew int
	SnFlag                int
}

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
// 查询没有发送的task
func GetMonitorNoSendData(strNodePubkey string, nThreshold int) (string, error) {
	if len(strNodePubkey) == 0 {
		return "", fmt.Errorf("pubkey is null")
	}

	return DBInf.GetTaskSchedulesNoSend(strNodePubkey, nThreshold)
}

//---------------------------------------------------------------------------
// 查询失败次数已经超过阈值的task
func GetMonitorNoSuccessData(strNodePubkey string, nThreshold int, flag int) (string, error) {
	if len(strNodePubkey) == 0 {
		return "", fmt.Errorf("pubkey is null")
	}

	return DBInf.GetTaskSchedulesNoSuccess(strNodePubkey, nThreshold, flag)
}

//---------------------------------------------------------------------------
// 批量设置发送标志为“已发送”，在查询到失败任务时调用
func UpdateMonitorSendBatch(slID []interface{}) error {
	if len(slID) == 0 {
		return fmt.Errorf("id slice is null")
	}
	return DBInf.SetTaskScheduleFlagBatch(slID, true)
}

//---------------------------------------------------------------------------
// 设置发送标志为“已发送”，在将任务插入待执行队列后调用
func UpdateMonitorSend(strID string) error {
	if len(strID) == 0 {
		return fmt.Errorf("id is null")
	}
	return DBInf.SetTaskScheduleFlag(strID, true)
}

//---------------------------------------------------------------------------
// 执行失败：1.更新strContractID & strContractHashOldID的SendFlag = 0,
// FailedCount + 1, LastExecuteTime, strTaskId, TaskState, nTaskExecuteIndex
// information 是 UpdateMonitorFailStruct 的json序列化字符串
func UpdateMonitorFail(information string) error {
	var errMsg string

	if len(information) == 0 {
		return fmt.Errorf("param is null")
	}

	var failStruct UpdateMonitorFailStruct
	err := json.Unmarshal([]byte(information), &failStruct)
	if err != nil {
		return err
	}

	strNodePubkey := config.Config.Keypair.PublicKey
	if len(strNodePubkey) == 0 ||
		len(failStruct.FstrContractID) == 0 ||
		len(failStruct.FstrContractHashID) == 0 {
		if len(strNodePubkey) == 0 {
			errMsg = "[strNodePubkey is null]"
		}
		if len(failStruct.FstrContractID) == 0 {
			errMsg = "[strContractID is null]"
		}
		if len(failStruct.FstrContractHashID) == 0 {
			errMsg = "[strContractHashID is null]"
		}
		return fmt.Errorf("param is null, %s", errMsg)
	}

	strID, err := DBInf.GetID(strNodePubkey,
		failStruct.FstrContractID,
		failStruct.FstrContractHashID)
	if err != nil {
		return err
	}

	if len(strID) == 0 {
		return fmt.Errorf("not find")
	}

	err = DBInf.SetTaskScheduleFlag(strID, false)
	if err != nil {
		return err
	}

	err = DBInf.SetTaskScheduleCount(strID, 1)
	if err != nil {
		return err
	}
	return DBInf.SetTaskState(strID,
		failStruct.FstrTaskId,
		failStruct.FstrTaskState,
		failStruct.FnTaskExecuteIndex)
}

//---------------------------------------------------------------------------
// 执行条件不满足：1.更新strContractID & strContractHashOldID的SendFlag = 0,
// WaitCount + 1,LastExecuteTime, strTaskId, TaskState, nTaskExecuteIndex
// information 是 UpdateMonitorWaitStruct 的json序列化字符串
func UpdateMonitorWait(information string) error {
	var errMsg string

	if len(information) == 0 {
		return fmt.Errorf("param is null")
	}

	var waitStruct UpdateMonitorWaitStruct
	err := json.Unmarshal([]byte(information), &waitStruct)
	if err != nil {
		return err
	}

	strNodePubkey := config.Config.Keypair.PublicKey
	if len(strNodePubkey) == 0 ||
		len(waitStruct.WstrContractID) == 0 ||
		len(waitStruct.WstrContractHashID) == 0 {
		if len(strNodePubkey) == 0 {
			errMsg = "[strNodePubkey is null]"
		}
		if len(waitStruct.WstrContractID) == 0 {
			errMsg = "[strContractID is null]"
		}
		if len(waitStruct.WstrContractHashID) == 0 {
			errMsg = "[strContractHashID is null]"
		}
		return fmt.Errorf("param is null, %s", errMsg)
	}

	strID, err := DBInf.GetID(strNodePubkey,
		waitStruct.WstrContractID,
		waitStruct.WstrContractHashID)
	if err != nil {
		return err
	}

	if len(strID) == 0 {
		return fmt.Errorf("not find")
	}

	err = DBInf.SetTaskScheduleFlag(strID, false)
	if err != nil {
		return err
	}

	err = DBInf.SetTaskScheduleCount(strID, 2)
	if err != nil {
		return err
	}
	return DBInf.SetTaskState(strID,
		waitStruct.WstrTaskId,
		waitStruct.WstrTaskState,
		waitStruct.WnTaskExecuteIndex)
}

//---------------------------------------------------------------------------
// 执行成功：1.更新strContractID & strContractHashOldID的的SendFlag=1,
//          SuccessCount + 1, LastExecuteTime, strTaskStateOld, TaskState, nTaskExecuteIndexOld
//         2.将strContractID & strContractHashNewID插入到扫描监控表中
// information 是 UpdateMonitorSuccStruct 的json序列化字符串
func UpdateMonitorSucc(information string) error {
	var errMsg string

	if len(information) == 0 {
		return fmt.Errorf("param is null")
	}

	var succStruct UpdateMonitorSuccStruct
	err := json.Unmarshal([]byte(information), &succStruct)
	if err != nil {
		return err
	}

	strNodePubkey := config.Config.Keypair.PublicKey
	if len(strNodePubkey) == 0 ||
		len(succStruct.SstrContractID) == 0 ||
		len(succStruct.SstrContractHashIdOld) == 0 ||
		len(succStruct.SstrContractHashIDNew) == 0 ||
		len(succStruct.SstrTaskIdNew) == 0 ||
		len(succStruct.SstrTaskStateNew) == 0 {
		if len(strNodePubkey) == 0 {
			errMsg = "[strNodePubkey is null]"
		}
		if len(succStruct.SstrContractID) == 0 {
			errMsg = "[strContractID is null]"
		}
		if len(succStruct.SstrContractHashIdOld) == 0 {
			errMsg = "[strContractHashIdOld is null]"
		}
		if len(succStruct.SstrContractHashIDNew) == 0 {
			errMsg = "[strContractHashIDNew is null]"
		}
		if len(succStruct.SstrTaskIdNew) == 0 {
			errMsg = "[strTaskIdNew is null]"
		}
		if len(succStruct.SstrTaskStateNew) == 0 {
			errMsg = "[strTaskStateNew is null]"
		}
		return fmt.Errorf("param is null, %s", errMsg)
	}

	strID, err := DBInf.GetID(strNodePubkey,
		succStruct.SstrContractID,
		succStruct.SstrContractHashIdOld)
	if err != nil {
		return err
	}

	if len(strID) == 0 {
		return fmt.Errorf("old contract id not find")
	}

	err = DBInf.SetTaskScheduleFlag(strID, true)
	if err != nil {
		return err
	}

	err = DBInf.SetTaskScheduleCount(strID, 0)
	if err != nil {
		return err
	}

	err = DBInf.SetTaskState(strID,
		succStruct.SstrTaskIdOld,
		succStruct.SstrTaskStateOld,
		succStruct.SnTaskExecuteIndexOld)
	if err != nil {
		return err
	}

	err = DBInf.SetTaskScheduleOverFlag(strID)
	if err != nil {
		return err
	}

	startTime, endTime, err := DBInf.GetValidTime(strID)
	if err != nil {
		return err
	}

	if len(startTime) == 0 || len(endTime) == 0 {
		return fmt.Errorf("old contract valid time not find")
	}

	var taskSchedule db.TaskSchedule
	taskSchedule.SendFlag = succStruct.SnFlag
	taskSchedule.Id = common.GenerateUUID()
	taskSchedule.ContractId = succStruct.SstrContractID
	taskSchedule.ContractHashId = succStruct.SstrContractHashIDNew
	taskSchedule.TaskId = succStruct.SstrTaskIdNew
	taskSchedule.TaskExecuteIndex = succStruct.SnTaskExecuteIndexNew
	taskSchedule.TaskState = succStruct.SstrTaskStateNew
	taskSchedule.NodePubkey = strNodePubkey
	taskSchedule.StartTime = startTime
	taskSchedule.EndTime = endTime
	slJson, _ := json.Marshal(taskSchedule)
	return DBInf.InsertTaskSchedule(string(slJson))
}

//---------------------------------------------------------------------------
// 直接把task干死
func UpdateMonitorDeal(strContractID string, strContractHashID string) error {
	strNodePubkey := config.Config.Keypair.PublicKey
	if len(strNodePubkey) == 0 ||
		len(strContractID) == 0 ||
		len(strContractHashID) == 0 {
		return fmt.Errorf("param is null")
	}

	strID, err := DBInf.GetID(strNodePubkey, strContractID,
		strContractHashID)
	if err != nil {
		return err
	}

	if len(strID) == 0 {
		return fmt.Errorf("not find")
	}

	return DBInf.SetTaskScheduleOverFlag(strID)
}

//---------------------------------------------------------------------------
// 只供头节点调用，根据公钥环为每个节点插入待执行任务
func InsertTaskSchedules(taskScheduleBase db.TaskSchedule) error {
	var err error
	var slMapTaskSchedule []interface{}
	allPublicKeys := config.GetAllPublicKey()
	for index, _ := range allPublicKeys {
		var taskSchedule db.TaskSchedule
		taskSchedule.Id = common.GenerateUUID()
		taskSchedule.ContractHashId = taskScheduleBase.ContractHashId
		taskSchedule.ContractId = taskScheduleBase.ContractId
		taskSchedule.TaskId = "0"
		taskSchedule.TaskExecuteIndex = 1
		taskSchedule.NodePubkey = allPublicKeys[index]
		taskSchedule.StartTime = taskScheduleBase.StartTime
		taskSchedule.EndTime = taskScheduleBase.EndTime

		mapObj, _ := common.StructToMap(taskSchedule)
		slMapTaskSchedule = append(slMapTaskSchedule, mapObj)
	}

	nInsertCount, err := DBInf.InsertTaskSchedules(slMapTaskSchedule)
	uniledgerlog.Debug("insert taskScheduled count is %d, err is %v", nInsertCount, err)
	return err
}

//---------------------------------------------------------------------------
