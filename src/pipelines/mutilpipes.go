package pipelines

import (
	"github.com/astaxie/beego/logs"
)

type Node struct {
	target    func(args ...interface{}) interface{}
	input     chan interface{}
	output    chan interface{}
	processes int
	name      string
}

func (n *Node) start() {
	for i := 0; i < n.processes; i++ {
		logs.Info("start Node run forver", n.name)
		go n.runForever()
	}
}

func (n *Node) runForever() {
	logs.Info("runforer")
	for {
		logs.Info("run")
		err := n.run()
		if err != nil {
			logs.Error(err)
			return
		}
	}
}

//t, ok := <-cvInput
//cvOutput <- common.Serialize(res)
func (n *Node) run() error {
	logs.Info("in run()")
	x, ok := <-n.input
	logs.Info("da")
	if !ok {
		logs.Info("null data")
		return nil
	}
	logs.Info(x)
	//TODO
	n.target(x)
	return nil
}

type Pipeline struct {
	nodes []Node
}

func (p *Pipeline) setup(indata Node) {
	nodes_copy := p.nodes[:]
	inNode := []Node{indata}
	nodes_copy = append(inNode, nodes_copy...)
	p.connect(nodes_copy)
}

func (p *Pipeline) connect(nodes []Node) (ch chan interface{}) {
	/*
        connect(items_copy, False)
	*/
	if len(nodes) == 0 {
		return nil
	}
	head := nodes[0]
	tail := nodes[1:]
	head.output = p.connect(tail)
	return head.input
}

func (p *Pipeline) start() {
	for _, node := range p.nodes {
		node.start()
	}

}
