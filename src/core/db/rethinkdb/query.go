package rethinkdb

import (
//	"fmt"
	"log"

	r "gopkg.in/gorethink/gorethink.v3"
)

func Get(db string,name string,id string) *r.Cursor {
	session := ConnectDB(db)
	res, err := r.Table(name).Get(id).Run(session)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return res
}

func Insert(db string,name string,jsonstr string) r.WriteResponse {
	session := ConnectDB(db)
	res, err :=r.Table(name).Insert(r.JSON(jsonstr)).RunWrite(session)
        if err != nil {
                log.Fatalf(err.Error())
        }
        return res
}

func Update(db string,name string,id string,jsonstr string) r.WriteResponse {
	session := ConnectDB(db)
	res, err :=r.Table(name).Get(id).Update(r.JSON(jsonstr)).RunWrite(session)
	if err != nil {
                log.Fatalf(err.Error())
        }
        return res
}

func Delete(db string,name string,id string) r.WriteResponse {
	session := ConnectDB(db)
	res, err :=r.Table(name).Get(id).Delete().RunWrite(session)
        if err != nil {
                log.Fatalf(err.Error())
        }
        return res
}
