package rethinkdb

import (
	"fmt"
	"log"

	r "gopkg.in/gorethink/gorethink.v3"
)

var Tables = [3] string{"Contract","Votes","ContractTasks"}

func CreateTable(db string ,name string) {
	session := ConnectDB(db)
	respo, err := r.TableCreate(name).RunWrite(session)
	if err != nil {
	    log.Fatalf("Error creating table: %s", err)
	}

	fmt.Printf("%d table created\n", respo.TablesCreated)
}

func CreateDatabase(name string) {
	session := Connect()
	resp, err := r.DBCreate(name).RunWrite(session)
        if err != nil {
            log.Fatalf("Error creating database: %s", err)
        }

        fmt.Printf("%d DB created\n", resp.DBsCreated)
}

func DropDatabase(name string) {
	session := Connect()
	resp, err := r.DBDrop(name).RunWrite(session)
	if err != nil {
		log.Fatalf("Error dropping database: %s", err)
	}

	fmt.Printf("%d DB dropped, %d tables dropped\n", resp.DBsDropped, resp.TablesDropped)
}

func InitDatabase() {
	dbname :="Unicontract"
	CreateDatabase(dbname)

	for _,x := range Tables {
		CreateTable(dbname,x)
	}
}
