package rethinkdb

import (
	"github.com/astaxie/beego/logs"
	r "gopkg.in/gorethink/gorethink.v3"
)

func Changefeed(db string, name string) *r.Cursor {
	session := ConnectDB(db)
	res, err := r.Table(name).Changes().Run(session)
	if err != nil {
		logs.Error(err.Error())
	}
	return res
}
