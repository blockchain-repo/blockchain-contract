package pipelines

import (
	"github.com/astaxie/beego/logs"
	"testing"
	"time"
	"unicontract/src/common"
	"unicontract/src/config"
)

func init() {
	config.Init()
}

func pip3(arg interface{}) interface{} {
	logs.Info("P3 param:", arg)
	s := common.Serialize(arg)
	logs.Info("P3 return:===", s)
	return s
}

func pip4(arg interface{}) interface{} {
	logs.Info("P4 param:===", arg)
	s := "return pip4"
	time.Sleep(time.Second * 5)
	logs.Info("P4 return:===", s)
	return s
}

func TestStart(t *testing.T) {
	startTxElection()
	time.Sleep(time.Second * 2)
	logs.Info("down")
}

func TestChannel(t *testing.T) {
	chan1 := make(chan interface{}, 10)
	chan2 := make(chan interface{}, 10)
	//chan1 <-"1"
	//logs.Info(<-chan2)

	logs.Info("%v", chan1)
	logs.Info("%v", chan2)
	chan2 <- "2"
	chan1 <- "1"
	chan1 = chan2

	logs.Info(&chan1)
	logs.Info(&chan2)

	logs.Info(<-chan1)
	logs.Info(<-chan1)
	logs.Info(<-chan2)
	logs.Info(<-chan2)
}

func TestTail(t *testing.T) {
	txNodeSlice := make([]Node, 0)
	txNodeSlice = append(txNodeSlice, Node{target: pip3, routineNum: 1, name: "pip2", input: make(chan interface{}, 10), output: make(chan interface{}, 10)})
	txNodeSlice = append(txNodeSlice, Node{target: pip3, routineNum: 1, name: "pip3", input: make(chan interface{}, 10), output: make(chan interface{}, 10)})
	txNodeSlice = append(txNodeSlice, Node{target: pip4, routineNum: 2, name: "pip4", input: make(chan interface{}, 10), output: make(chan interface{}, 10)})

	txNodeSlice[0].name = "change"
	head := txNodeSlice[0]
	logs.Info(txNodeSlice[0].name)
	logs.Info(head.name)

	tail := txNodeSlice[1:]
	tail[0].name = "change"
	logs.Info(txNodeSlice[1].name)

	//for _, node := range txNodeSlice {
	//	logs.Info("pip in range:", node.name)
	//	//go node.start()
	//	node.name = "change"
	//}
	logs.Info(txNodeSlice[0].name)
	logs.Info(txNodeSlice[1].name)
}

func TestNode(t *testing.T) {
	node1 := Node{target: pip3, routineNum: 1, name: "pip2", input: make(chan interface{}, 10), output: make(chan interface{}, 10)}
	node2 := Node{target: pip3, routineNum: 1, name: "pip3", input: make(chan interface{}, 10), output: make(chan interface{}, 10)}
	node3 := Node{target: pip3, routineNum: 1, name: "pip4", input: make(chan interface{}, 10), output: make(chan interface{}, 10)}

	txNodeSlice := make([]*Node, 0, 5)
	txNodeSlice = append(txNodeSlice, &node2)
	txNodeSlice = append(txNodeSlice, &node3)
	txPip := Pipeline{
		nodes: txNodeSlice,
	}

	inNode := []*Node{&node1}
	nodes_all := append(inNode, txPip.nodes...)

	nodes_all[0].name = "change1"
	nodes_all[1].name = "change2"
	nodes_all[2].name = "change3"
	logs.Info(nodes_all[0].name, "----", node1.name)
}
