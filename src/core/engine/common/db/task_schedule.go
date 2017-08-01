// task_schedule
package db

// table [TaskSchedule]
type TaskSchedule struct {
	Id               string `json:"id"`
	NodePubkey       string
	ContractId       string
	ContractHashId   string
	TaskId           string
	TaskState        string
	TaskExecuteIndex int
	SuccessCount     int
	FailedCount      int
	WaitCount        int
	SendFlag         int
	OverFlag         int
	StartTime        string
	EndTime          string
	LastExecuteTime  string
}
