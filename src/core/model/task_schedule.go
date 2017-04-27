// task_schedule
package model

// table [TaskSchedule]
type TaskSchedule struct {
	Id            string `json:"id"`
	ContractId    string
	NodePubkey    string
	StartTime     string
	EndTime       string
	ContractState string
	SendFlag      int
}
