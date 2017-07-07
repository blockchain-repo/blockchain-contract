package quartz

import (
	"time"
	"unicontract/src/common"
	"unicontract/src/common/monitor"
	"unicontract/src/core/db/rethinkdb"

	"unicontract/src/common/uniledgerlog"
)

func init() {
	go sendContractStatsToMonitor()
}

func sendContractStatsToMonitor() {
	timer := time.Tick(20 * time.Second)
	for now := range timer {
		uniledgerlog.Info(now)
		Contracts_number, err := rethinkdb.GetContractsCount()
		if err != nil {
			uniledgerlog.Error(err)
		}
		Contract_Create, err := rethinkdb.GetContractStatsCount("Contract_Create")
		if err != nil {
			uniledgerlog.Error(err)
		}
		Contract_Signature, err := rethinkdb.GetContractStatsCount("Contract_Signature")
		if err != nil {
			uniledgerlog.Error(err)
		}
		Contract_In_Process, err := rethinkdb.GetContractStatsCount("Contract_In_Process")
		if err != nil {
			uniledgerlog.Error(err)
		}
		Contract_Discarded, err := rethinkdb.GetContractStatsCount("Contract_Discarded")
		if err != nil {
			uniledgerlog.Error(err)
		}
		Contract_Completed, err := rethinkdb.GetContractStatsCount("Contract_Completed")
		if err != nil {
			uniledgerlog.Error(err)
		}
		task_send_flag_success, err := rethinkdb.GetTaskSendFlagCount(1)
		if err != nil {
			uniledgerlog.Error(err)
		}
		task_send_flag_fail, err := rethinkdb.GetTaskSendFlagCount(0)
		if err != nil {
			uniledgerlog.Error(err)
		}
		task_failed_Count, err := rethinkdb.GetTaskScheduleCount("FailedCount")
		if err != nil {
			uniledgerlog.Error(err)
		}
		task_wait_Count, err := rethinkdb.GetTaskScheduleCount("WaitCount")
		if err != nil {
			uniledgerlog.Error(err)
		}

		monitor.Monitor.Gauge("Contracts_number", common.StringToInt(Contracts_number))
		monitor.Monitor.Gauge("Contract_Create", common.StringToInt(Contract_Create))
		monitor.Monitor.Gauge("Contract_Signature", common.StringToInt(Contract_Signature))
		monitor.Monitor.Gauge("Contract_In_Process", common.StringToInt(Contract_In_Process))
		monitor.Monitor.Gauge("Contract_Discarded", common.StringToInt(Contract_Discarded))
		monitor.Monitor.Gauge("Contract_Completed", common.StringToInt(Contract_Completed))
		monitor.Monitor.Gauge("task_send_flag_success", common.StringToInt(task_send_flag_success))
		monitor.Monitor.Gauge("task_send_flag_fail", common.StringToInt(task_send_flag_fail))
		monitor.Monitor.Gauge("task_failed_count", common.StringToInt(task_failed_Count))
		monitor.Monitor.Gauge("task_wait_count", common.StringToInt(task_wait_Count))
		contract_number_temp := 1
		if common.StringToInt(Contracts_number) != 0 {
			contract_number_temp = common.StringToInt(Contracts_number)
		}
		monitor.Monitor.Gauge("contract_decrease_rate", common.StringToInt(task_send_flag_success)/contract_number_temp)
	}
}
