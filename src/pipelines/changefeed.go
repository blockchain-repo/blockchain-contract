package pipelines

import (
	"time"

	r "unicontract/src/core/db/rethinkdb"

	"unicontract/src/common/uniledgerlog"
)

const (
	INSERT = 1
	DELETE = 2
	UPDATE = 4
)

type ChangeFeed struct {
	node      Node
	db        string
	table     string
	operation int
}

func (c *ChangeFeed) runChangeFeed() {
	uniledgerlog.Info("change feed run")
	var value interface{}
	res := r.Changefeed(c.db, c.table)
	for res.Next(&value) {
		m := value.(map[string]interface{})
		//uniledgerlog.Info(m)
		isInsert := (m["old_val"] == nil)
		isDelete := (m["new_val"] == nil)
		isUpdate := !isInsert && !isDelete
		//uniledgerlog.Info(isInsert, isDelete, isUpdate)
		if isInsert && ((c.operation & INSERT) != 0) {
			uniledgerlog.Info(c.table, "insert Changefeed result : %s", m["new_val"])
			c.node.output <- m["new_val"]
		}
		if isDelete && ((c.operation & DELETE) != 0) {
			uniledgerlog.Info(c.table, "delete Changefeed result : %s", m["old_val"])
			c.node.output <- m["old_val"]
		}
		if isUpdate && ((c.operation & UPDATE) != 0) {
			uniledgerlog.Info(c.table, "update Changefeed result : %s", m["new_val"])
			c.node.output <- m["new_val"]
		}
	}
	uniledgerlog.Info("change feed out")
}

func (c *ChangeFeed) runForever() {
	for {
		c.runChangeFeed()
		time.Sleep(time.Second)
	}
}
