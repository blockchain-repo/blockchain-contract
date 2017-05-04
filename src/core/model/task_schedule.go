// task_schedule
package model

// table [TaskSchedule]
type TaskSchedule struct {
	Id              string `json:"id"`
	ContractId      string
	NodePubkey      string
	StartTime       string
	EndTime         string
	LastExecuteTime string

	/*执行成功、失败次数记录字段*/
	SuccessCount int
	FailedCount  int

	// TODO time 最后一次执行的时间

	/*扫描标志。只有为0时，才会被送入队列等待执行
	  0 代表未进入执行队列，或者之前的执行失败，需要再次入列执行；
	  1 代表正在队列中等待执行或者已经执行成功。
	*/
	SendFlag int
}
