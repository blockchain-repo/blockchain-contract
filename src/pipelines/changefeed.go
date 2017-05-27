package pipelines

import (
	r "unicontract/src/core/db/rethinkdb"

	"github.com/astaxie/beego/logs"
	"time"
)

type ChangeFeed struct {
	node      Node
	db        string
	table     string
	operation []string
}

func (c *ChangeFeed) runChangeFeed() {
	logs.Info("change feed run")
	var value interface{}
	res := r.Changefeed(c.db, c.table)
	for res.Next(&value) {
		m := value.(map[string]interface{})
		logs.Info(c.table, "Changefeed result : %s", m["new_val"])
		if m["new_val"] != nil {
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
