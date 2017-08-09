package db

type RunState int

const (
	NORMAL RunState = iota
	ALREADY_RUN_SUCCESS
	NO_ARRIVAL_TIME
	OVER_TIME
	FAILED_TIMES_BEYOND
	WAIT_TIMES_BEYOND
	WAIT_FOR_RUN
)

func (this RunState) String() string {
	switch this {
	case NORMAL:
		return "normal"
	case ALREADY_RUN_SUCCESS:
		return "already run success"
	case NO_ARRIVAL_TIME:
		return "no arrival run time"
	case OVER_TIME:
		return "over run time"
	case FAILED_TIMES_BEYOND:
		return "failed times beyond"
	case WAIT_TIMES_BEYOND:
		return "wait times beyond"
	case WAIT_FOR_RUN:
		return "waiting for run"
	default:
		return "Unknow"
	}
}

type Datebase interface {
	InitDatabase() error
	CreateDatabase(dbName string) error
	CreateTable(dbName string, tableName string) error
	DropDatabase(dbName string) error

	Insert(dbName, tableName, json string) (bool, error)
	Delete(dbName, tableName, id string) (bool, error)
	Update(dbName, tableName, id, json string) (bool, error)
	Query(dbName, tableName, id string) (map[string]interface{}, error)

	InsertTaskSchedule(strTaskSchedule string) error
	InsertTaskSchedules(slTaskSchedule []interface{}) (int, error)
	GetID(strNodePubkey, strContractID string, strContractHashId string) (string, error)
	GetValidTime(strID string) (string, string, error)
	SetTaskScheduleFlag(strID string, alreadySend bool) error
	SetTaskScheduleFlagBatch(slID []interface{}, alreadySend bool) error
	SetTaskScheduleOverFlag(strID string) error
	SetTaskState(strID, strTaskId, strState string, nTaskExecuteIndex int) error
	SetTaskScheduleCount(strID string, flag int) error
	GetTaskSchedulesNoSend(strNodePubkey string, nThreshold int) (string, error)
	GetTaskSchedulesNoSuccess(strNodePubkey string, nThreshold int, flag int) (string, error)
	GetTaskSchedulesSuccess(strNodePubkey string) (string, error)
	DeleteTaskSchedules(slID []interface{}) (int, error)

	GetTaskScheduleState(strContractID, strContractHashId string, failedThreshold, waitThreshold int) (RunState, error)
	GetTaskScheduleCount(stat string) (string, error)
	GetTaskSendFlagCount(stat int) (string, error)
}
