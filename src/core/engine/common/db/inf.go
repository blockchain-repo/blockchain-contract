package db

type Datebase interface {
	Connect() (interface{}, error)
	ConnectDB(args ...interface{}) (interface{}, error)

	InitDatabase() error
	CreateDatabase(name string) error
	CreateTable(db string, name string) error
	DropDatabase() error

	Insert(args ...interface{}) (interface{}, error)
	Delete(args ...interface{}) (interface{}, error)
	Update(args ...interface{}) (interface{}, error)
	Query(args ...interface{}) (interface{}, error)

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
}
