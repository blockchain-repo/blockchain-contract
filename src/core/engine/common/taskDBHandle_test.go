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
	t.Logf("%+v\n", slTasks)
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
	t.Logf("%+v\n", slTasks)
}

func Test_UpdateMonitorSendBatch(t *testing.T) {
	var slID []interface{}
	slID = append(slID, "95971bd5-189c-4feb-9bcb-60f8d24594a9")
	err := UpdateMonitorSendBatch(slID)
	if err != nil {
		t.Error(err)
	}
}

func Test_UpdateMonitorSend(t *testing.T) {
	strID := "12667eff-6bff-4cb1-983d-3958c3c5d6a2"
	err := UpdateMonitorSend(strID)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
	}
}

func Test_UpdateMonitorFail(t *testing.T) {
	strContractID := "5a5ac312-9231-434c-8c0b-850e86dae9ef"
	strContractHashID := ""
	err := UpdateMonitorFail(strContractID, strContractHashID)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass\n")
	}
}

func Test_UpdateMonitorWait(t *testing.T) {
	strContractID := "5a5ac312-9231-434c-8c0b-850e86dae9ef"
	strContractHashID := ""
	err := UpdateMonitorWait(strContractID, strContractHashID)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass \n")
	}
}

func Test_UpdateMonitorSucc(t *testing.T) {
	strContractID := "5a5ac312-9231-434c-8c0b-850e86dae9ef"
	strContractHashOldID := ""
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
