package rethinkdb

import (
	"unicontract/src/common/uniledgerlog"

	r "gopkg.in/gorethink/gorethink.v3"
)

const DBNAME = "Unicontract"

const (
	TABLE_CONTRACTS            = "Contracts"
	TABLE_VOTES                = "Votes"
	TABLE_CONTRACT_TASKS       = "ContractTasks"
	TABLE_CONSENSUS_FAILURES   = "ConsensusFailures"
	TABLE_CONTRACT_OUTPUTS     = "ContractOutputs"
	TABLE_SEND_FAILING_RECORDS = "SendFailingRecords"

	//-------------------------------------------------------
	// 智能微网demo所需要数据表
	TABLE_ENERGYTRADINGDEMO_ROLE        = "EnergyTradingDemoRole"
	TABLE_ENERGYTRADINGDEMO_ENERGY      = "EnergyTradingDemoEnergy"
	TABLE_ENERGYTRADINGDEMO_TRANSACTION = "EnergyTradingDemoTransaction"
	TABLE_ENERGYTRADINGDEMO_BILL        = "EnergyTradingDemoBill"
	TABLE_ENERGYTRADINGDEMO_MSGNOTICE   = "EnergyTradingDemoMsgNotice"
	TABLE_ENERGYTRADINGDEMO_PRICE       = "EnergyTradingDemoPrice"
	//-------------------------------------------------------
	//tables for tianan
	TABLE_EARNINGS = "Earnings"
)

var Tables = []string{TABLE_CONTRACTS,
	TABLE_VOTES,
	TABLE_CONTRACT_TASKS,
	TABLE_CONSENSUS_FAILURES,
	TABLE_CONTRACT_OUTPUTS,
	TABLE_SEND_FAILING_RECORDS,
	//TABLE_ENERGYTRADINGDEMO_ROLE,
	//TABLE_ENERGYTRADINGDEMO_ENERGY,
	//TABLE_ENERGYTRADINGDEMO_TRANSACTION,
	//TABLE_ENERGYTRADINGDEMO_BILL,
	//TABLE_ENERGYTRADINGDEMO_MSGNOTICE,
	//TABLE_ENERGYTRADINGDEMO_PRICE,
	TABLE_EARNINGS,
}

func CreateTableIndex(db string) {
	session := ConnectDB(db)
	response, err := r.Table(TABLE_VOTES).IndexCreateFunc(
		"Vote_VoteFor",
		r.Row.Field("Vote").Field("VoteFor"),
		r.IndexCreateOpts{Multi: true},
	).RunWrite(session)
	if err != nil {
		uniledgerlog.Error("Error creating index:", err)
	}
	uniledgerlog.Info("%d index created", response.Created)
}

func CreateTable(db string, name string) {
	session := ConnectDB(db)
	respo, err := r.TableCreate(name).RunWrite(session)
	if err != nil {
		uniledgerlog.Error("Error creating table: %s", err)
	}

	uniledgerlog.Info("%d table created\n", respo.TablesCreated)
}

func CreateDatabase(name string) {
	session := Connect()
	resp, err := r.DBCreate(name).RunWrite(session)
	if err != nil {
		uniledgerlog.Error("Error creating database: %s", err)
	}

	uniledgerlog.Info("%d DB created\n", resp.DBsCreated)
}

func DropDatabase() {
	dbname := DBNAME
	session := Connect()
	resp, err := r.DBDrop(dbname).RunWrite(session)
	if err != nil {
		uniledgerlog.Error("Error dropping database: %s", err)
	}

	uniledgerlog.Info("%d DB dropped, %d tables dropped\n", resp.DBsDropped, resp.TablesDropped)
}

func InitDatabase() {
	dbname := DBNAME
	CreateDatabase(dbname)

	for _, x := range Tables {
		CreateTable(dbname, x)
	}

	CreateTableIndex(dbname)
}
