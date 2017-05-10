// taskDBHandle
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
// 设置发送标志为“已发送”，在将任务插入待执行队列后调用
func UpdateMonitorSend(strID string) error {
	if len(strID) == 0 {
		return fmt.Errorf("id is null")
	}
	return rethinkdb.SetTaskScheduleFlag(strID, true)
}

//---------------------------------------------------------------------------
// 执行失败：1.更新strContractID 的SendFlag = 0, FailedCount + 1, LastExecuteTime
// 返回FailedCount(或者SuccessCount)和error
func UpdateMonitorFail(strNodePubkey, strContractID string) (int, error) {
	if len(strNodePubkey) == 0 || len(strContractID) == 0 {
		return -1, fmt.Errorf("pubkey or contractid is null")
	}

	strID, err := rethinkdb.GetID(strNodePubkey, strContractID)
	if err != nil {
		return -1, err
	}

	if len(strID) == 0 {
		return -1, fmt.Errorf("not find")
	}

	err = rethinkdb.SetTaskScheduleFlag(strID, false)
	if err != nil {
		return -1, err
	}

	return rethinkdb.SetTaskScheduleCount(strID, false)
}

//---------------------------------------------------------------------------
// 执行条件不满足：1.更新strNodePubkey  的SendFlag = 0, LastExecuteTime
func UpdateMonitorWait(strNodePubkey, strContractID string) error {
	if len(strNodePubkey) == 0 || len(strContractID) == 0 {
		return fmt.Errorf("pubkey or contractid is null")
	}

	strID, err := rethinkdb.GetID(strNodePubkey, strContractID)
	if err != nil {
		return err
	}

	if len(strID) == 0 {
		return fmt.Errorf("not find")
	}

	return rethinkdb.SetTaskScheduleFlag(strID, false)
}

//---------------------------------------------------------------------------
// 执行成功：1.更新strContractIDold 的SendFlag=1, SuccessCount + 1, LastExecuteTime
//         2.将strContractIDnew 插入到扫描监控表中
func UpdateMonitorSucc(strNodePubkey, strContractIDold, strContractIDnew string) error {
	if len(strNodePubkey) == 0 ||
		len(strContractIDold) == 0 ||
		len(strContractIDnew) == 0 {
		return fmt.Errorf("pubkey or contractid is null")
	}

	strID, err := rethinkdb.GetID(strNodePubkey, strContractIDold)
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

	_, err = rethinkdb.SetTaskScheduleCount(strID, true)
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
	taskSchedule.ContractId = strContractIDnew
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
	allPublicKeys := config.GetAllPublicKey()
	for index, _ := range allPublicKeys {
		var taskSchedule model.TaskSchedule
		taskSchedule.Id = common.GenerateUUID()
		taskSchedule.ContractId = taskScheduleBase.ContractId
		taskSchedule.NodePubkey = allPublicKeys[index]
		taskSchedule.StartTime = taskScheduleBase.StartTime
		taskSchedule.EndTime = taskScheduleBase.EndTime
		taskSchedule.FailedCount = 0
		taskSchedule.SendFlag = 0

		slJson, _ := json.Marshal(taskSchedule)
		err = rethinkdb.InsertTaskSchedule(string(slJson))
		if err != nil {
			beegoLog.Error("insert [%s] TaskSchedule is error, error is %s",
				taskScheduleBase.ContractId, err.Error())
			break
		}
	}
	return err
}

//---------------------------------------------------------------------------
