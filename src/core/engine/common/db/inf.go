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
	InsertBatch(dbName, tableName string, slTaskSchedule []interface{}) (int, error)
	Delete(dbName, tableName, id string) (bool, error)
	DeleteBatch(dbName, tableName string, slID []interface{}) (int, error)
	Update(dbName, tableName, id, json string) (bool, error)
	UpdateBatch(dbName, tableName, json string, slID []interface{}) (bool, error)
	UpdateToAdd(strDBName, strTableName, strID, strField string, num int) (bool, error)
	Query(dbName, tableName, id string) (map[string]interface{}, error)

	QueryByContractId(strDBName, strTableName, strNodePubkey, strContractID string) ([]map[string]interface{}, error)
	QueryByIdAndHashId(strDBName, strTableName, strNodePubkey, strContractID, strContractHashId string) (map[string]interface{}, error)
	GetTaskSchedulesNoSend(strDBName, strTableName, strNowTime, strNodePubkey string, nThreshold int) ([]map[string]interface{}, error)
	GetTaskSchedulesNoSuccess(strDBName, strTableName, strCount, strNodePubkey string, nThreshold int, flag int) ([]map[string]interface{}, error)
	GetTaskSchedulesSuccess(strDBName, strTableName, strNodePubkey string) ([]map[string]interface{}, error)

	GetTaskScheduleCount(strDBName, strTableName, stat string, num int) (string, error)
	GetTaskSendFlagCount(strDBName, strTableName string, stat int) (string, error)
}
