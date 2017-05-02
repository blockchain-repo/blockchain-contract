package rethinkdb

import (
	"log"

	r "gopkg.in/gorethink/gorethink.v3"
)

func Reconfig(shards int,replicas int) *r.Cursor {
	dbname := DBNAME
	var opts r.ReconfigureOpts
	opts.Shards = shards
	opts.Replicas = replicas

	session := Connect()
	resp, err := r.DB(dbname).Reconfigure(opts).Run(session)
	if err != nil {
		log.Fatalf("Error reconfig database: %s", err)
	}
	return resp
}

