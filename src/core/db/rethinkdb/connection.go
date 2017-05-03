package rethinkdb

import (
	"github.com/astaxie/beego/logs"

	r "gopkg.in/gorethink/gorethink.v3"
)

func Connect() *r.Session { // FIXME: GetSession?
	/*
		conf := config.ReadConfig(config.DevelopmentEnv)
		session, err := r.Connect(r.ConnectOpts{
			Address:    conf.DatabaseUrl,
			Database:   conf.DatabaseName,
			InitialCap: conf.DatabaseInitialCap,
			MaxOpen:    conf.DatabaseMaxOpen,
		})
	*/

	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
	})

	if err != nil {
		logs.Error(err.Error())
	}
	return session
}

func ConnectDB(dbname string) *r.Session { // FIXME: GetSession?
	/*
	   conf := config.ReadConfig(config.DevelopmentEnv)
	   session, err := r.Connect(r.ConnectOpts{
	           Address:    conf.DatabaseUrl,
	           Database:   conf.DatabaseName,
	           InitialCap: conf.DatabaseInitialCap,
	           MaxOpen:    conf.DatabaseMaxOpen,
	   })
	*/

	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: dbname,
	})

	if err != nil {
		logs.Error(err.Error())
	}
	return session
}
