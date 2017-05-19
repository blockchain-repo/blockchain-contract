// taskDBHandle_test
package common

import (
	"strconv"
	"testing"
	"time"

	"encoding/json"
	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/model"
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

	var slTasks []model.TaskSchedule
	json.Unmarshal([]byte(str), &slTasks)
	t.Logf("slTask count is %d, %+v\n", len(slTasks), slTasks)
}

func Test_GetMonitorFailedData(t *testing.T) {
	strNodePubkey := config.Config.Keypair.PublicKey
	str, err := GetMonitorFailedData(strNodePubkey, 50)
	if err != nil {
		t.Error(err)
	}

	if len(str) == 0 {
		t.Log("no send data")
	}

	var slTasks []model.TaskSchedule
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
	strContractID := "e38b0a4c-374b-4ad8-b27e-93eb6c0d633c"
	strContractHashID := "053d7d59-60be-4c18-8efd-afd0b3385e74"
	err := UpdateMonitorFail(strContractID, strContractHashID)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
	}
}

func Test_UpdateMonitorWait(t *testing.T) {
	strContractID := "e212353c-36cd-4c3c-ad8a-239767d53b40"
	strContractHashID := "94059f17-6dbe-4901-b958-c3758b1e6ecb"
	err := UpdateMonitorWait(strContractID, strContractHashID)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass \n")
	}
}

func Test_UpdateMonitorSucc(t *testing.T) {
	strContractID := "e212353c-36cd-4c3c-ad8a-239767d53b40"
	strContractHashOldID := "94059f17-6dbe-4901-b958-c3758b1e6ecb"
	strContractHashNewID := common.GenerateUUID()
	err := UpdateMonitorSucc(strContractID, strContractHashOldID, strContractHashNewID)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass \n")
	}
}

func Test_InsertTaskSchedules(t *testing.T) {
	var taskSchedule model.TaskSchedule
	taskSchedule.ContractId = common.GenerateUUID()
	taskSchedule.ContractHashId = common.GenerateUUID()
	taskSchedule.StartTime = common.GenTimestamp()
	taskSchedule.EndTime = strconv.FormatInt(time.Now().Add(time.Hour*24*5).UnixNano()/1000000, 10)

	err := InsertTaskSchedules(taskSchedule)
	if err != nil {
		t.Error(err)
	}
}
