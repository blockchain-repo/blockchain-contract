package rethinkdb

import (
	"sync"
	"time"

	"unicontract/src/common/uniledgerlog"
	"unicontract/src/config"

	"github.com/astaxie/beego"
	r "gopkg.in/gorethink/gorethink.v3"
)

var (
	one                     sync.Once
	session                 *r.Session
	rethinkDBInitialCap     int
	rethinkDBMaxOpen        int
	rethinkDBReconnectCount int
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
	port := config.Config.Port
	session, err := r.Connect(r.ConnectOpts{
		Address: ip + ":" + port,
	})

	if err != nil {
		uniledgerlog.Error(err.Error())
	}
	return session
}

func ConnectDB(dbname string) *r.Session { // FIXME: GetSession?
	var err error
	one.Do(func() {
		var err error
		rethinkDBInitialCap, err = beego.AppConfig.Int("RethinkDBInitialCap")
		if err != nil {
			uniledgerlog.Error(err)
			rethinkDBInitialCap = 50
		}

		rethinkDBMaxOpen, err = beego.AppConfig.Int("RethinkDBMaxOpen")
		if err != nil {
			uniledgerlog.Error(err)
			rethinkDBMaxOpen = 100
		}

		rethinkDBReconnectCount, err = beego.AppConfig.Int("RethinkDBReconnectCount")
		if err != nil {
			uniledgerlog.Error(err)
			rethinkDBReconnectCount = 5
		}

		ip := config.Config.LocalIp
		port := config.Config.Port
		session, err = r.Connect(r.ConnectOpts{
			Address:    ip + ":" + port,
			Database:   dbname,
			InitialCap: rethinkDBInitialCap,
			MaxOpen:    rethinkDBMaxOpen,
		})
		if err != nil {
			uniledgerlog.Error(err.Error())
		}
	})

	count := rethinkDBReconnectCount
	for count > 0 {
		if !session.IsConnected() {
			count--
			err = session.Reconnect()
			if err != nil {
				uniledgerlog.Error(err.Error())
			}
		} else {
			break
		}
		time.Sleep(time.Millisecond * 500)
	}

	return session
}
