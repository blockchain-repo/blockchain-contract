package rethinkdb

import (
	"unicontract/src/common/uniledgerlog"

	r "gopkg.in/gorethink/gorethink.v3"
	"unicontract/src/config"
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
	ip := config.Config.LocalIp
	session, err := r.Connect(r.ConnectOpts{
		Address: ip + ":28015",
	})

	if err != nil {
		uniledgerlog.Error(err.Error())
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
	ip := config.Config.LocalIp
	session, err := r.Connect(r.ConnectOpts{
		Address: ip + ":28015",
		Database: dbname,
	})

	if err != nil {
		uniledgerlog.Error(err.Error())
	}
	return session
}