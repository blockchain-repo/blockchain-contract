package rethinkdb

import (
	"math/rand"
	"sync"
	"time"

	"unicontract/src/common/uniledgerlog"
	"unicontract/src/config"

	"github.com/astaxie/beego"
	r "gopkg.in/gorethink/gorethink.v3"
)

var (
	one              sync.Once
	slSession        []*r.Session
	rethinkDBMaxOpen int
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
		slSession = make([]*r.Session, 0)
		rethinkDBMaxOpen, err = beego.AppConfig.Int("RethinkDBMaxOpen")
		if err != nil {
			uniledgerlog.Error(err)
			rethinkDBMaxOpen = 100
		}
	})
	if len(slSession) < rethinkDBMaxOpen {
		ip := config.Config.LocalIp
		port := config.Config.Port
		session, err := r.Connect(r.ConnectOpts{
			Address:  ip + ":" + port,
			Database: dbname,
		})

		if err != nil {
			uniledgerlog.Error(err.Error())
		}
		slSession = append(slSession, session)
		return session
	}
	rand.Seed(time.Now().UnixNano())
	return slSession[rand.Intn(len(slSession))]
}
