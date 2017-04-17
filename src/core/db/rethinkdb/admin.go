package rethinkdb

import (
	"log"

	r "gopkg.in/gorethink/gorethink.v3"
)

func Reconfig(db string,shards int,replicas int) *r.Cursor {
	var opts r.ReconfigureOpts
	opts.Shards = shards
	opts.Replicas = replicas

	session := Connect()
	resp, err := r.DB(db).Reconfigure(opts).Run(session)
	if err != nil {
		log.Fatalf("Error reconfig database: %s", err)
	}

	return resp
}

