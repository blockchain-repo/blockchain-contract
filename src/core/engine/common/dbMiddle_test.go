package common

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/engine/common/db"
)

func init() {
	config.Init()
	DBInf = db.GetInstance()
}

func Test_InsertTaskSchedule(t *testing.T) {
	var taskSchedule db.TaskSchedule
	taskSchedule.Id = common.GenerateUUID()
	taskSchedule.ContractHashId = common.GenerateUUID()
	taskSchedule.ContractId = common.GenerateUUID()
	taskSchedule.TaskId = "0"
	taskSchedule.TaskExecuteIndex = 1
	taskSchedule.NodePubkey = config.Config.Keypair.PublicKey
	taskSchedule.StartTime = common.GenTimestamp()
	taskSchedule.EndTime = strconv.FormatInt(time.Now().Add(time.Hour*24*5).UnixNano()/1000000, 10)

	slJson, _ := json.Marshal(taskSchedule)
	err := InsertTaskSchedule(string(slJson))
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
	}
}

func Test_InsertTaskSchedules_(t *testing.T) {
	var taskSchedule db.TaskSchedule
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

	insertCount, err := InsertTaskSchedules_(slMapTaskSchedule)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass. insertCount is %d\n", insertCount)
	}
}

func Test_DeleteTaskSchedules(t *testing.T) {
	slID := make([]interface{}, 0)
	slID = append(slID, "32a92ae4-a524-4771-b95f-b64d323433e4")
	slID = append(slID, "fcfb3db5-faa1-4477-822e-e1b2a8bf5d38")
	slID = append(slID, "fc776c28-5dc7-4124-9c0e-3abab2b6b88b")

	deleteNum, err := DeleteTaskSchedules(slID)
	t.Logf("deleteNum is %d\n", deleteNum)
	t.Logf("err is %+v\n", err)
}

func Test_GetID(t *testing.T) {
	strNodePubkey := config.Config.Keypair.PublicKey
	strContractID := "03a1c78d-67c0-45c3-9ec0-6de9bb9c288e"
	strContractHashId := "875ef66a-12be-4489-87bd-6174b8448225"

	strID, err := GetID(strNodePubkey, strContractID, strContractHashId)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass, id is \" %s \"\n", strID)
	}
}

func Test_GetValidTime(t *testing.T) {
	strID := "9be998c1-dda6-44b9-971f-2e7e906f1098"
	startTime, endTime, err := GetValidTime(strID)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass, startTime is \" %s \", endTime is \" %s \"\n", startTime, endTime)
	}
}

func Test_GetTaskSchedulesNoSend(t *testing.T) {
	strNodePubkey := config.Config.Keypair.PublicKey
	retStr, err := GetTaskSchedulesNoSend(strNodePubkey, 500)
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
	retStr, err := GetTaskSchedulesNoSuccess(strNodePubkey, 80, 1)
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
	str, err := GetTaskSchedulesSuccess(config.Config.Keypair.PublicKey)
	if err != nil {
		t.Error(err)
	}

	if len(str) == 0 {
		t.Logf("is null\n")
	} else {
		var slTask []db.TaskSchedule
		json.Unmarshal([]byte(str), &slTask)
		t.Logf("slTask count is %d, %+v\n", len(slTask), slTask)
	}
}

func Test_SetTaskScheduleFlag(t *testing.T) {
	strID := "9be998c1-dda6-44b9-971f-2e7e906f1098"
	err := SetTaskScheduleFlag(strID, true)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
	}
}

func Test_SetTaskScheduleFlagBatch(t *testing.T) {
	slID := make([]interface{}, 0)
	slID = append(slID, "39465922-34fc-40fa-85d1-29ae9dd289aa")
	slID = append(slID, "5c4b5b76-b6b9-47f0-a996-7e411a931601")
	slID = append(slID, "2c6e063e-2448-4f54-a340-7c2cc8d341e6")
	SetTaskScheduleFlagBatch(slID, true)
}

func Test_SetTaskScheduleOverFlag(t *testing.T) {
	strID := "9be998c1-dda6-44b9-971f-2e7e906f1098"
	err := SetTaskScheduleOverFlag(strID)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
	}
}

func Test_SetTaskState(t *testing.T) {
	strID := "9be998c1-dda6-44b9-971f-2e7e906f1098"
	strTaskId := "1"
	strStat := "asdfasdfasdfasdfasdf"
	nTaskExecuteIndex := 12121
	err := SetTaskState(strID, strTaskId, strStat, nTaskExecuteIndex)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
	}
}

func Test_SetTaskScheduleCount(t *testing.T) {
	strID := "9be998c1-dda6-44b9-971f-2e7e906f1098"
	err := SetTaskScheduleCount(strID, 0)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
	}
}
