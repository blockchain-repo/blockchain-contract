package rethinkdb

import (
	"log"

	r "gopkg.in/gorethink/gorethink.v3"
)

func Changefeed(db string, name string) *r.Cursor {
	session := ConnectDB(db)
	res, err := r.Table(name).Changes().Run(session)
	if err != nil {
		log.Fatal(err.Error())
	}
	return res
}
