package pipelines

import (
	"errors"

	"github.com/astaxie/beego/logs"
)

type Node struct {
	target     func(interface{}) interface{}
	input      chan interface{}
	output     chan interface{}
	routineNum int
	name       string
}

func (n *Node) start() {
	for i := 0; i < n.routineNum; i++ {
		go n.runForever()
	}
}

func (n *Node) runForever() {
	for {
		//logs.Info(n.name, ",in run forever")
		err := n.run()
		if err != nil {
			logs.Error(err)
			return
		}
	}
}

func (n *Node) run() error {
	x, ok := <-n.input
	if !ok {
		logs.Error(errors.New("read data from inputchannel error"))
		return nil
	}
	//TODO  not good enough, how to support multi params and returns
	if n.output == nil {
		return nil
	}
	n.output <- n.target(x)
	return nil
}

type Pipeline struct {
	nodes []*Node
}

func (p *Pipeline) setup(indata *Node) {
	inNode := []*Node{indata}
	nodes_all := append(inNode, p.nodes...)
	p.connect(nodes_all)
}

func (p *Pipeline) connect(nodes []*Node) (ch chan interface{}) {

	if len(nodes) == 0 {
		return nil
	}

	head := nodes[0]
	head.input = make(chan interface{}, 10)
	head.output = make(chan interface{}, 10)
	tail := nodes[1:]
	head.output = p.connect(tail)
	return head.input
}

func (p *Pipeline) start() {
	for index, _ := range p.nodes {
		p.nodes[index].start()
	}
}
