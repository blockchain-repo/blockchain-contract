package pipelines

import (
	r "unicontract/src/core/db/rethinkdb"
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"bytes"
)

type ChangeFeed struct {
	node      Node
	table     string
	operation []string
}

func (c *ChangeFeed)runForever(){
	c.runChangeFeed()
}

func (c *ChangeFeed) runChangeFeed() {
	logs.Info("change feed run")
	var value interface{}
	res := r.Changefeed("Unicontract", "ContractOutputs")
	for res.Next(&value) {
		logs.Info(" txElection step1 : txeChangefeed ")
		m := value.(map[string]interface{})
		v, err := json.Marshal(m["new_val"])
		if err != nil {
			logs.Error(err.Error())
			continue
		}
		if bytes.Equal(v, []byte("null")) {
			continue
		}
		logs.Info("txeChangefeed result : %s", v)
		c.node.output <- v
	}
}
