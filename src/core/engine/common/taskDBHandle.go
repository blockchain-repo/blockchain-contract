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
	beegoLog "github.com/astaxie/beego/logs"
)

import (
	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
)

//---------------------------------------------------------------------------
// 查询没有发送的task
func GetMonitorNoSendData(strNodePubkey string, nThreshold int) (string, error) {
	if len(strNodePubkey) == 0 {
		return "", fmt.Errorf("pubkey is null")
	}

	return rethinkdb.GetTaskSchedulesNoSend(strNodePubkey, nThreshold)
}

//---------------------------------------------------------------------------
// 查询失败次数已经超过阈值的task
func GetMonitorNoSuccessData(strNodePubkey string, nThreshold int, flag int) (string, error) {
	if len(strNodePubkey) == 0 {
		return "", fmt.Errorf("pubkey is null")
	}

	return rethinkdb.GetTaskSchedulesNoSuccess(strNodePubkey, nThreshold, flag)
}

//---------------------------------------------------------------------------
// 批量设置发送标志为“已发送”，在查询到失败任务时调用
func UpdateMonitorSendBatch(slID []interface{}) error {
	if len(slID) == 0 {
		return fmt.Errorf("id slice is null")
	}
	return rethinkdb.SetTaskScheduleFlagBatch(slID, true)
}

//---------------------------------------------------------------------------
// 设置发送标志为“已发送”，在将任务插入待执行队列后调用
func UpdateMonitorSend(strID string) error {
	if len(strID) == 0 {
		return fmt.Errorf("id is null")
	}
	return rethinkdb.SetTaskScheduleFlag(strID, true)
}

//---------------------------------------------------------------------------
// 执行失败：1.更新strContractID & strContractHashOldID的SendFlag = 0,
// FailedCount + 1, LastExecuteTime, strTaskId, TaskState, nTaskExecuteIndex
func UpdateMonitorFail(strContractID string,
	strContractHashID string,
	strTaskId string,
	strTaskState string,
	nTaskExecuteIndex int) error {
	strNodePubkey := config.Config.Keypair.PublicKey
	if len(strNodePubkey) == 0 ||
		len(strContractID) == 0 ||
		len(strContractHashID) == 0 {
		return fmt.Errorf("param is null")
	}

	strID, err := rethinkdb.GetID(strNodePubkey, strContractID,
		strContractHashID)
	if err != nil {
		return err
	}

	if len(strID) == 0 {
		return fmt.Errorf("not find")
	}

	err = rethinkdb.SetTaskScheduleFlag(strID, false)
	if err != nil {
		return err
	}

	err = rethinkdb.SetTaskScheduleCount(strID, 1)
	if err != nil {
		return err
	}
	return rethinkdb.SetTaskState(strID, strTaskId, strTaskState, nTaskExecuteIndex)
}

//---------------------------------------------------------------------------
// 执行条件不满足：1.更新strContractID & strContractHashOldID的SendFlag = 0,
// WaitCount + 1,LastExecuteTime, strTaskId, TaskState, nTaskExecuteIndex
func UpdateMonitorWait(strContractID string,
	strContractHashID string,
	strTaskId string,
	strTaskState string,
	nTaskExecuteIndex int) error {
	strNodePubkey := config.Config.Keypair.PublicKey
	if len(strNodePubkey) == 0 ||
		len(strContractID) == 0 ||
		len(strContractHashID) == 0 {
		return fmt.Errorf("param is null")
	}

	strID, err := rethinkdb.GetID(strNodePubkey, strContractID,
		strContractHashID)
	if err != nil {
		return err
	}

	if len(strID) == 0 {
		return fmt.Errorf("not find")
	}

	err = rethinkdb.SetTaskScheduleFlag(strID, false)
	if err != nil {
		return err
	}

	err = rethinkdb.SetTaskScheduleCount(strID, 2)
	if err != nil {
		return err
	}
	return rethinkdb.SetTaskState(strID, strTaskId, strTaskState, nTaskExecuteIndex)
}

//---------------------------------------------------------------------------
// 执行成功：1.更新strContractID & strContractHashOldID的的SendFlag=1,
//          SuccessCount + 1, LastExecuteTime, strTaskStateOld, TaskState, nTaskExecuteIndexOld
//        2.将strContractID & strContractHashNewID插入到扫描监控表中
func UpdateMonitorSucc(strContractID string,
	strContractHashIdOld string,
	strTaskStateOld string,
	strTaskIdOld string,
	nTaskExecuteIndexOld int,
	strContractHashIDNew string,
	strTaskIdNew string,
	strTaskStateNew string,
	nTaskExecuteIndexNew int) error {

	strNodePubkey := config.Config.Keypair.PublicKey
	if len(strNodePubkey) == 0 ||
		len(strContractID) == 0 ||
		len(strContractHashIdOld) == 0 ||
		len(strTaskIdOld) == 0 ||
		len(strTaskStateOld) == 0 ||
		len(strContractHashIDNew) == 0 ||
		len(strTaskIdNew) == 0 ||
		len(strTaskStateNew) == 0 {
		return fmt.Errorf("param is null")
	}

	strID, err := rethinkdb.GetID(strNodePubkey, strContractID, strContractHashIdOld)
	if err != nil {
		return err
	}

	if len(strID) == 0 {
		return fmt.Errorf("old contract id not find")
	}

	err = rethinkdb.SetTaskScheduleFlag(strID, true)
	if err != nil {
		return err
	}

	err = rethinkdb.SetTaskScheduleCount(strID, 0)
	if err != nil {
		return err
	}

	err = rethinkdb.SetTaskState(strID, strTaskIdOld, strTaskStateOld, nTaskExecuteIndexOld)
	if err != nil {
		return err
	}

	err = rethinkdb.SetTaskScheduleOverFlag(strID)
	if err != nil {
		return err
	}

	startTime, endTime, err := rethinkdb.GetValidTime(strID)
	if err != nil {
		return err
	}

	if len(startTime) == 0 || len(endTime) == 0 {
		return fmt.Errorf("old contract valid time not find")
	}

	var taskSchedule model.TaskSchedule
	taskSchedule.Id = common.GenerateUUID()
	taskSchedule.ContractId = strContractID
	taskSchedule.ContractHashId = strContractHashIDNew
	taskSchedule.TaskId = strTaskIdNew
	taskSchedule.TaskExecuteIndex = nTaskExecuteIndexNew
	taskSchedule.TaskState = strTaskStateNew
	taskSchedule.NodePubkey = strNodePubkey
	taskSchedule.StartTime = startTime
	taskSchedule.EndTime = endTime
	slJson, _ := json.Marshal(taskSchedule)
	return rethinkdb.InsertTaskSchedule(string(slJson))
}

//---------------------------------------------------------------------------
// 只供头节点调用，根据公钥环为每个节点插入待执行任务
func InsertTaskSchedules(taskScheduleBase model.TaskSchedule) error {
	var err error
	var slMapTaskSchedule []interface{}
	allPublicKeys := config.GetAllPublicKey()
	for index, _ := range allPublicKeys {
		var taskSchedule model.TaskSchedule
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

	nInsertCount, err := rethinkdb.InsertTaskSchedules(slMapTaskSchedule)
	beegoLog.Debug("insert taskScheduled count is %d, err is %v", nInsertCount, err)
	return err
}

//---------------------------------------------------------------------------
