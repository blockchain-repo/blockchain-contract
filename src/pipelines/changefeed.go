package pipelines

import (
	"time"

	r "unicontract/src/core/db/rethinkdb"

	"github.com/astaxie/beego/logs"
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
	logs.Info("change feed run")
	var value interface{}
	res := r.Changefeed(c.db, c.table)
	for res.Next(&value) {
		m := value.(map[string]interface{})
		isInsert := (m["new_val"] == nil)
		isDelete := (m["old_val"] == nil)
		isUpdate := !isInsert && !isDelete
		if isInsert && ((c.operation & INSERT) != 0) {
			logs.Info(c.table, "Changefeed result : %s", m["new_val"])
			c.node.output <- m["new_val"]
		}
		if isDelete && ((c.operation & DELETE) != 0) {
			logs.Info(c.table, "Changefeed result : %s", m["old_val"])
			c.node.output <- m["old_val"]
		}
		if isUpdate && ((c.operation & UPDATE) != 0) {
			logs.Info(c.table, "Changefeed result : %s", m["new_val"])
			c.node.output <- m["new_val"]
		}
	}
	logs.Info("change feed out")
}

func (c *ChangeFeed) runForever() {
	for {
		c.runChangeFeed()
		time.Sleep(time.Second)
	}
}
