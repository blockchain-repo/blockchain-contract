// taskDBHandle_test
package common

import (
	"strconv"
	"testing"
	"time"

	"encoding/json"
	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/engine/common/db"
)

func init() {
	config.Init()
}

func Test_GetMonitorNoSendData(t *testing.T) {
	strNodePubkey := config.Config.Keypair.PublicKey
	str, err := GetMonitorNoSendData(strNodePubkey, 50)
	if err != nil {
		t.Error(err)
	}

	if len(str) == 0 {
		t.Log("no send data")
	}

	var slTasks []db.TaskSchedule
	json.Unmarshal([]byte(str), &slTasks)
	t.Logf("slTask count is %d, %+v\n", len(slTasks), slTasks)
}

func Test_GetMonitorNoSuccessData(t *testing.T) {
	strNodePubkey := config.Config.Keypair.PublicKey
	str, err := GetMonitorNoSuccessData(strNodePubkey, 50, 0)
	if err != nil {
		t.Error(err)
	}

	if len(str) == 0 {
		t.Log("no send data")
	}

	var slTasks []db.TaskSchedule
	json.Unmarshal([]byte(str), &slTasks)
	t.Logf("slTask count is %d, %+v\n", len(slTasks), slTasks)
}

func Test_UpdateMonitorSendBatch(t *testing.T) {
	var slID []interface{}
	slID = append(slID, "172a6bd7-f502-46fd-aba9-a6c098a9ee28")
	slID = append(slID, "28f0b597-4403-4082-a9a1-cd765099faa6")
	slID = append(slID, "d5501c6f-3f74-47d7-bcaa-1f7050aa8196")
	err := UpdateMonitorSendBatch(slID)
	if err != nil {
		t.Error(err)
	}
}

func Test_UpdateMonitorSend(t *testing.T) {
	strID := "d5501c6f-3f74-47d7-bcaa-1f7050aa8196"
	err := UpdateMonitorSend(strID)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
	}
}

func Test_UpdateMonitorFail(t *testing.T) {
	var failStruct UpdateMonitorFailStruct
	failStruct.FstrContractID = "9d3be6de-4fb1-4bd0-867b-b83e18f80203"
	failStruct.FstrContractHashID = "54f21b37-601d-42c8-93f5-f3acf41c19c4"
	failStruct.FstrTaskId = "asdfasdfasdfasfasdf"
	failStruct.FstrTaskState = "asdfasdfasdf"
	failStruct.FnTaskExecuteIndex = 56785
	slFailData, _ := json.Marshal(failStruct)
	err := UpdateMonitorFail(string(slFailData))
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
	}
}

func Test_UpdateMonitorWait(t *testing.T) {
	var waitStruct UpdateMonitorWaitStruct
	waitStruct.WstrContractID = "9d3be6de-4fb1-4bd0-867b-b83e18f80203"
	waitStruct.WstrContractHashID = "54f21b37-601d-42c8-93f5-f3acf41c19c4"
	waitStruct.WstrTaskId = "1234123412341234"
	waitStruct.WstrTaskState = "asdfasdfasdf"
	waitStruct.WnTaskExecuteIndex = 2222
	slWaitData, _ := json.Marshal(waitStruct)
	err := UpdateMonitorWait(string(slWaitData))
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass \n")
	}
}

func Test_UpdateMonitorSucc(t *testing.T) {
	var succStruct UpdateMonitorSuccStruct
	succStruct.SstrContractID = "72bdda0a-f8e6-4fa5-89e5-f93a5b470159"
	succStruct.SstrContractHashIdOld = "1da2972e-a40d-45f7-a4ec-c19c3a9f7a02"
	succStruct.SstrTaskStateOld = "old"
	succStruct.SstrTaskIdOld = "999999"
	succStruct.SnTaskExecuteIndexOld = 22
	succStruct.SstrContractHashIDNew = common.GenerateUUID()
	succStruct.SstrTaskIdNew = "1000000"
	succStruct.SstrTaskStateNew = "new"
	succStruct.SnTaskExecuteIndexNew = 23
	succStruct.SnFlag = 0
	slSuccData, _ := json.Marshal(succStruct)
	err := UpdateMonitorSucc(string(slSuccData))
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass \n")
	}
}

func Test_InsertTaskSchedules(t *testing.T) {
	var taskSchedule db.TaskSchedule
	taskSchedule.ContractId = common.GenerateUUID()
	taskSchedule.ContractHashId = common.GenerateUUID()
	taskSchedule.StartTime = common.GenTimestamp()
	taskSchedule.EndTime = strconv.FormatInt(time.Now().Add(time.Hour*24*5).UnixNano()/1000000, 10)

	err := InsertTaskSchedules(taskSchedule)
	if err != nil {
		t.Error(err)
	}
}

func Test_UpdateMonitorDeal(t *testing.T) {
	err := UpdateMonitorDeal("170627144008249104", "bee43b640f8af8fe340a59b651bafef24156d71de650f35b6ae8f0d2f9e69e8e")
	if err != nil {
		t.Error(err)
	}
}

func Test_GetTaskScheduleCount(t *testing.T) {
	count, err := GetTaskScheduleCount("WaitCount", 50)
	if err != nil {
		t.Error(err)
	}
	t.Log(count)
	t.Logf("deleteNum is %s\n", count)
	t.Logf("err is %+v\n", err)
}
