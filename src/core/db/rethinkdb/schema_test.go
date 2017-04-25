package rethinkdb

import (
	"testing"
)

func Test_InitDatabase(t *testing.T) {
	InitDatabase()
}

func Test_DropDatabase(t *testing.T) {
	DropDatabase(DBNAME)
}
