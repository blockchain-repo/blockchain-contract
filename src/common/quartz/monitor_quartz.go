package quartz

import (
	"github.com/astaxie/beego/logs"
	"time"
	"unicontract/src/common"
	"unicontract/src/common/monitor"
	"unicontract/src/core/db/rethinkdb"
)

func init() {
	go sendContractStatsToMonitor()
}

func sendContractStatsToMonitor() {
	timer := time.Tick(10 * time.Second)
	for now := range timer {
		logs.Info(now)
		Contracts_number, err := rethinkdb.GetContractsCount()
		if err != nil {
			logs.Error(err)
		}
		Contract_Create, err := rethinkdb.GetContractStatsCount("Contract_Create")
		if err != nil {
			logs.Error(err)
		}
		Contract_Signature, err := rethinkdb.GetContractStatsCount("Contract_Signature")
		if err != nil {
			logs.Error(err)
		}
		Contract_In_Process, err := rethinkdb.GetContractStatsCount("Contract_In_Process")
		if err != nil {
			logs.Error(err)
		}
		Contract_Discarded, err := rethinkdb.GetContractStatsCount("Contract_Discarded")
		if err != nil {
			logs.Error(err)
		}
		Contract_Completed, err := rethinkdb.GetContractStatsCount("Contract_Completed")
		if err != nil {
			logs.Error(err)
		}
		task_failed_Count, err := rethinkdb.GetTaskScheduleCount("FailedCount")
		if err != nil {
			logs.Error(err)
		}
		task_wait_Count, err := rethinkdb.GetTaskScheduleCount("WaitCount")
		if err != nil {
			logs.Error(err)
		}

		monitor.Monitor.Gauge("Contracts_number", common.StringToInt(Contracts_number))
		monitor.Monitor.Gauge("Contract_Create", common.StringToInt(Contract_Create))
		monitor.Monitor.Gauge("Contract_Signature", common.StringToInt(Contract_Signature))
		monitor.Monitor.Gauge("Contract_In_Process", common.StringToInt(Contract_In_Process))
		monitor.Monitor.Gauge("Contract_Discarded", common.StringToInt(Contract_Discarded))
		monitor.Monitor.Gauge("Contract_Completed", common.StringToInt(Contract_Completed))
		monitor.Monitor.Gauge("task_failed_Count", common.StringToInt(task_failed_Count))
		monitor.Monitor.Gauge("task_wait_Count", common.StringToInt(task_wait_Count))
	}
}
