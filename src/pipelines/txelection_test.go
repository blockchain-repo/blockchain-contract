package pipelines

import (
	"testing"
	"time"

	"unicontract/src/common"
	"unicontract/src/config"

	"unicontract/src/common/uniledgerlog"
	"unicontract/src/chain"
)

func init() {
	config.Init()
}

func pip3(arg interface{}) interface{} {
	uniledgerlog.Info("P3 param:", arg)
	s := common.Serialize(arg)
	uniledgerlog.Info("P3 return:===", s)
	return s
}

func pip4(arg interface{}) interface{} {
	uniledgerlog.Info("P4 param:===", arg)
	s := "return pip4"
	time.Sleep(time.Second * 5)
	uniledgerlog.Info("P4 return:===", s)
	return s
}

func TestStart(t *testing.T) {
	startTxElection()
	time.Sleep(time.Second * 2)
	uniledgerlog.Info("down")
}

func TestChannel(t *testing.T) {
	chan1 := make(chan interface{}, 10)
	chan2 := make(chan interface{}, 10)
	//chan1 <-"1"
	//uniledgerlog.Info(<-chan2)

	uniledgerlog.Info("%v", chan1)
	uniledgerlog.Info("%v", chan2)
	chan2 <- "2"
	chan1 <- "1"
	chan1 = chan2

	uniledgerlog.Info(&chan1)
	uniledgerlog.Info(&chan2)

	uniledgerlog.Info(<-chan1)
	uniledgerlog.Info(<-chan1)
	uniledgerlog.Info(<-chan2)
	uniledgerlog.Info(<-chan2)
}

func TestTail(t *testing.T) {
	txNodeSlice := make([]Node, 0)
	txNodeSlice = append(txNodeSlice, Node{target: pip3, routineNum: 1, name: "pip2", input: make(chan interface{}, 10), output: make(chan interface{}, 10)})
	txNodeSlice = append(txNodeSlice, Node{target: pip3, routineNum: 1, name: "pip3", input: make(chan interface{}, 10), output: make(chan interface{}, 10)})
	txNodeSlice = append(txNodeSlice, Node{target: pip4, routineNum: 2, name: "pip4", input: make(chan interface{}, 10), output: make(chan interface{}, 10)})

	txNodeSlice[0].name = "change"
	head := txNodeSlice[0]
	uniledgerlog.Info(txNodeSlice[0].name)
	uniledgerlog.Info(head.name)

	tail := txNodeSlice[1:]
	tail[0].name = "change"
	uniledgerlog.Info(txNodeSlice[1].name)

	//for _, node := range txNodeSlice {
	//	uniledgerlog.Info("pip in range:", node.name)
	//	//go node.start()
	//	node.name = "change"
	//}
	uniledgerlog.Info(txNodeSlice[0].name)
	uniledgerlog.Info(txNodeSlice[1].name)
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
	uniledgerlog.Info(nodes_all[0].name, "----", node1.name)
}

func TestTxQueryEists(t *testing.T) {
	tx_id := "1235be2235f4514bc24c26c43e2d93d59ced6ab5470f4fafc9187d43469cfc66"
	jsonBody := `{"tx_id":"` + tx_id + `"}`
	result, _ := chain.GetContractTx(jsonBody)
	uniledgerlog.Info(result)
	uniledgerlog.Info(len(result.Data.([]interface{})))
}
