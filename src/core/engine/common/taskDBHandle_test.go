// taskDBHandle_test
package common

import (
	"strconv"
	"testing"
	"time"

	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/model"
)

func init() {
	config.Init()
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
	strNodePubkey := config.Config.Keypair.PublicKey
	strContractID := "5a5ac312-9231-434c-8c0b-850e86dae9ef"
	failedCount, err := UpdateMonitorFail(strNodePubkey, strContractID)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass, failedCount is %d\n", failedCount)
	}
}

func Test_UpdateMonitorWait(t *testing.T) {
	strNodePubkey := config.Config.Keypair.PublicKey
	strContractID := "5a5ac312-9231-434c-8c0b-850e86dae9ef"
	err := UpdateMonitorWait(strNodePubkey, strContractID)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass \n")
	}
}

func Test_UpdateMonitorSucc(t *testing.T) {
	strNodePubkey := config.Config.Keypair.PublicKey
	strContractID := "5a5ac312-9231-434c-8c0b-850e86dae9ef"
	strContractIDnew := common.GenerateUUID()
	err := UpdateMonitorSucc(strNodePubkey, strContractID, strContractIDnew)
	if err != nil {
		t.Errorf("not pass, return err is \" %s \"\n", err.Error())
	} else {
		t.Logf("pass \n")
	}
}

func Test_InsertTaskSchedules(t *testing.T) {
	var taskSchedule model.TaskSchedule
	taskSchedule.ContractId = common.GenerateUUID()
	taskSchedule.StartTime = common.GenTimestamp()
	taskSchedule.EndTime = strconv.FormatInt(time.Now().Add(time.Hour*24*5).UnixNano()/1000000, 10)

	err := InsertTaskSchedules(taskSchedule)
	if err != nil {
		t.Error(err)
	}
}
