package pipelines

import (
	"errors"
	"time"

	"unicontract/src/common/uniledgerlog"

	"github.com/astaxie/beego"
)

var (
	NodeInputChannelSize  int
	NodeOutputChannelSize int
)

type Node struct {
	target     func(interface{}) interface{}
	input      chan interface{}
	output     chan interface{}
	routineNum int
	name       string
	timeout    int64
}

func (n *Node) start() {
	for i := 0; i < n.routineNum; i++ {
		go n.runForever()
	}
}

func (n *Node) runForever() {
	for {
		//uniledgerlog.Info(n.name, ",in run forever")
		err := n.run()
		if err != nil {
			uniledgerlog.Error(err)
			return
		}
	}
}

func (n *Node) run() error {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * time.Duration(n.timeout)) //等待10秒钟
		if n.timeout != 0 {
			timeout <- true
		}
	}()
	select {
	case x, ok := <-n.input:
		//从ch中读到数据
		if !ok {
			uniledgerlog.Error(errors.New("read data from inputchannel error"))
			return nil
		}
		//TODO  not good enough, how to support multi params and returns
		out := n.target(x)
		if n.output == nil || out == nil {
			return nil
		}
		n.output <- out
	case <-timeout:
		//一直没有从ch中读取到数据，但从timeout中读取到数据
		//uniledgerlog.Info("read data timeout")
		return nil
	}
	return nil
}

type Pipeline struct {
	nodes []*Node
}

func (p *Pipeline) setup(indata *Node) {
	inNode := []*Node{indata}
	nodes_all := append(inNode, p.nodes...)
	var err error
	NodeInputChannelSize, err = beego.AppConfig.Int("PipelineNodeInputChannelSize")
	if err != nil {
		uniledgerlog.Error(err)
		NodeInputChannelSize = 10
	}
	NodeOutputChannelSize, err = beego.AppConfig.Int("PipelineNodeOutputChannelSize")
	if err != nil {
		uniledgerlog.Error(err)
		NodeOutputChannelSize = 10
	}
	p.connect(nodes_all)
}

func (p *Pipeline) connect(nodes []*Node) (ch chan interface{}) {

	if len(nodes) == 0 {
		return nil
	}

	head := nodes[0]
	head.input = make(chan interface{}, NodeInputChannelSize)
	head.output = make(chan interface{}, NodeOutputChannelSize)
	tail := nodes[1:]
	head.output = p.connect(tail)
	return head.input
}

func (p *Pipeline) start() {
	for index, _ := range p.nodes {
		p.nodes[index].start()
	}
}
