package db

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"unicontract/src/common"
	"unicontract/src/config"
)

var inf Datebase

func init() {
	config.Init()
	inf = GetInstance()
}

func Test_InsertTaskSchedule(t *testing.T) {
	var taskSchedule TaskSchedule
	taskSchedule.Id = common.GenerateUUID()
	taskSchedule.ContractHashId = common.GenerateUUID()
	taskSchedule.ContractId = common.GenerateUUID()
	taskSchedule.TaskId = "0"
	taskSchedule.TaskExecuteIndex = 1
	taskSchedule.NodePubkey = config.Config.Keypair.PublicKey
	taskSchedule.StartTime = common.GenTimestamp()
	taskSchedule.EndTime = strconv.FormatInt(time.Now().Add(time.Hour*24*5).UnixNano()/1000000, 10)

	slJson, _ := json.Marshal(taskSchedule)
	err := inf.InsertTaskSchedule(string(slJson))
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
	}
}

func Test_InsertTaskSchedules(t *testing.T) {
	var taskSchedule TaskSchedule
	taskSchedule.Id = common.GenerateUUID()
	taskSchedule.ContractId = common.GenerateUUID()
	taskSchedule.ContractHashId = common.GenerateUUID()
	taskSchedule.TaskId = "0"
	taskSchedule.TaskExecuteIndex = 1
	taskSchedule.NodePubkey = config.Config.Keypair.PublicKey
	taskSchedule.StartTime = common.GenTimestamp()
	taskSchedule.EndTime = strconv.FormatInt(time.Now().Add(time.Hour*24*5).UnixNano()/1000000, 10)
	taskSchedule.FailedCount = 50
	taskSchedule.WaitCount = 50

	mapObj1, _ := common.StructToMap(taskSchedule)
	taskSchedule.Id = common.GenerateUUID()
	mapObj2, _ := common.StructToMap(taskSchedule)
	taskSchedule.Id = common.GenerateUUID()
	mapObj3, _ := common.StructToMap(taskSchedule)
	taskSchedule.Id = common.GenerateUUID()
	mapObj4, _ := common.StructToMap(taskSchedule)
	taskSchedule.Id = common.GenerateUUID()
	mapObj5, _ := common.StructToMap(taskSchedule)
	taskSchedule.Id = common.GenerateUUID()

	var slMapTaskSchedule []interface{}
	slMapTaskSchedule = append(slMapTaskSchedule, mapObj1)
	slMapTaskSchedule = append(slMapTaskSchedule, mapObj2)
	slMapTaskSchedule = append(slMapTaskSchedule, mapObj3)
	slMapTaskSchedule = append(slMapTaskSchedule, mapObj4)
	slMapTaskSchedule = append(slMapTaskSchedule, mapObj5)

	insertCount, err := inf.InsertTaskSchedules(slMapTaskSchedule)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass. insertCount is %d\n", insertCount)
	}
}

func Test_GetID(t *testing.T) {
	strNodePubkey := config.Config.Keypair.PublicKey
	strContractID := "7b4ee68b-bece-4feb-b8c5-f3270b40af27"
	strContractHashId := "656b9c4f-f755-4e98-8700-5f409b833587"

	strID, err := inf.GetID(strNodePubkey, strContractID, strContractHashId)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass, id is \" %s \"\n", strID)
	}
}

func Test_GetValidTime(t *testing.T) {
	strID := "7d5e1a52-96d5-4866-ae1f-f752e3e27d75"
	startTime, endTime, err := inf.GetValidTime(strID)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass, startTime is \" %s \", endTime is \" %s \"\n", startTime, endTime)
	}
}

func Test_SetTaskScheduleFlagBatch(t *testing.T) {
	slID := make([]interface{}, 0)
	slID = append(slID, "40b9943a-dffc-443c-94ff-b5a550837325")
	slID = append(slID, "f0f69545-15be-4f23-871d-4eb80d5c3f60")
	slID = append(slID, "6bf5ba67-6b70-493c-b1ca-a80487d3bd7c")
	inf.SetTaskScheduleFlagBatch(slID, true)
}

func Test_SetTaskScheduleFlag(t *testing.T) {
	strID := "a2184e4c-260d-44c9-a6c4-cfa9962610ff"
	err := inf.SetTaskScheduleFlag(strID, false)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
	}
}

func Test_SetTaskScheduleOverFlag(t *testing.T) {
	strID := "a2184e4c-260d-44c9-a6c4-cfa9962610ff"
	err := inf.SetTaskScheduleOverFlag(strID)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
	}
}

func Test_SetTaskScheduleCount(t *testing.T) {
	strID := "ea626dd3-313a-4ab9-998b-79a4117a2944"
	err := inf.SetTaskScheduleCount(strID, 0)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
	}
}

func Test_SetTaskState(t *testing.T) {
	strID := "ea626dd3-313a-4ab9-998b-79a4117a2944"
	strTaskId := "1"
	strStat := "asdfasdfasdfasdfasdf"
	nTaskExecuteIndex := 12121
	err := inf.SetTaskState(strID, strTaskId, strStat, nTaskExecuteIndex)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
	}
}

func Test_GetTaskSchedulesNoSend(t *testing.T) {
	strNodePubkey := config.Config.Keypair.PublicKey
	retStr, err := inf.GetTaskSchedulesNoSend(strNodePubkey, 500)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
		if len(retStr) != 0 {
			//var slTask []model.TaskSchedule
			var slTask []map[string]interface{}
			json.Unmarshal([]byte(retStr), &slTask)
			t.Logf("slTask count is %d, %+v\n", len(slTask), slTask)

			t.Logf("Id type is %T\n", slTask[0]["id"])
			t.Logf("ContractId type is %T\n", slTask[0]["ContractId"])
			t.Logf("NodePubkey type is %T\n", slTask[0]["NodePubkey"])
			t.Logf("SendFlag type is %T\n", slTask[0]["SendFlag"])
			t.Logf("StartTime type is %T\n", slTask[0]["StartTime"])
			t.Logf("EndTime type is %T\n", slTask[0]["EndTime"])
			t.Logf("FailedCount type is %T\n", slTask[0]["FailedCount"])
			t.Logf("SuccessCount type is %T\n", slTask[0]["SuccessCount"])
			t.Logf("LastExecuteTime type is %T\n", slTask[0]["LastExecuteTime"])
		}
	}
}

func Test_GetTaskSchedulesNoSuccess(t *testing.T) {
	strNodePubkey := config.Config.Keypair.PublicKey
	retStr, err := inf.GetTaskSchedulesNoSuccess(strNodePubkey, 40, 1)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
		if len(retStr) != 0 {
			//var slTask []model.TaskSchedule
			var slTask []map[string]interface{}
			json.Unmarshal([]byte(retStr), &slTask)
			t.Logf("slTask count is %d, %+v\n", len(slTask), slTask)

			t.Logf("Id type is %T\n", slTask[0]["id"])
			t.Logf("ContractId type is %T\n", slTask[0]["ContractId"])
			t.Logf("NodePubkey type is %T\n", slTask[0]["NodePubkey"])
			t.Logf("SendFlag type is %T\n", slTask[0]["SendFlag"])
			t.Logf("StartTime type is %T\n", slTask[0]["StartTime"])
			t.Logf("EndTime type is %T\n", slTask[0]["EndTime"])
			t.Logf("FailedCount type is %T\n", slTask[0]["FailedCount"])
			t.Logf("SuccessCount type is %T\n", slTask[0]["SuccessCount"])
			t.Logf("LastExecuteTime type is %T\n", slTask[0]["LastExecuteTime"])
		}
	}
}

func Test_GetTaskSchedulesSuccess(t *testing.T) {
	str, err := inf.GetTaskSchedulesSuccess(config.Config.Keypair.PublicKey)
	if err != nil {
		t.Error(err)
	}

	if len(str) == 0 {
		t.Logf("is null\n")
	} else {
		var slTask []TaskSchedule
		json.Unmarshal([]byte(str), &slTask)
		t.Logf("slTask count is %d, %+v\n", len(slTask), slTask)
	}
}

func Test_DeleteTaskSchedules(t *testing.T) {
	slID := make([]interface{}, 0)
	slID = append(slID, "40b9943a-dffc-443c-94ff-b5a550837325")
	slID = append(slID, "f0f69545-15be-4f23-871d-4eb80d5c3f60")
	slID = append(slID, "6bf5ba67-6b70-493c-b1ca-a80487d3bd7c")

	deleteNum, err := inf.DeleteTaskSchedules(slID)
	t.Logf("deleteNum is %d\n", deleteNum)
	t.Logf("err is %+v\n", err)
}

func Test_GetTaskScheduleCount(t *testing.T) {
	count, err := inf.GetTaskScheduleCount("WaitCount", 50)
	if err != nil {
		t.Error(err)
	}
	t.Error(count)
	t.Logf("deleteNum is %d\n", count)
	t.Logf("err is %+v\n", err)
}
