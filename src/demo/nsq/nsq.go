package nsq

import (
	"github.com/nsqio/go-nsq"
	logger "unicontract/src/common/seelog"
	"fmt"
)

func NewProducer() *nsq.Producer{
	p,err :=nsq.NewProducer("192.168.80.39:4150",nsq.NewConfig())
	if err != nil{
		logger.Error(err)
	}
	return p
}

func Publish(){
	p :=NewProducer()
	err :=p.Publish("unichain",[]byte("ssss"))

	if err != nil {
		logger.Error(err)
	}
	p1 :=NewProducer()
	err1 :=p1.Publish("wsp",[]byte("wspwsp"))

	if err1 != nil {
		logger.Error(err)
	}
}




type ConsumerT struct {

}

func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	//fmt.Println(msg.Attempts)
	//fmt.Println(msg.ID)
	//fmt.Println(msg.NSQDAddress)
	//fmt.Println(msg.Timestamp)
	//fmt.Println(string(msg.Body))
	fmt.Println(string(msg.Body))
	//msg.WriteTo(os.Stdout)

	return nil
}

type ConsumerAAA struct {

}

func (*ConsumerAAA) HandleMessage(msg *nsq.Message) error {
	fmt.Println(string(msg.Body))
	p1,err1 :=nsq.NewProducer("192.168.80.39:4150",nsq.NewConfig())
	if err1 != nil{
		logger.Error(err1)
	}
	err4 := p1.Publish("secnsss",[]byte("secnssss"))
	if err4 != nil{
		logger.Error(err4)
	}

	return nil
}

type ConsumerBBB struct {

}

func (*ConsumerBBB) HandleMessage(msg *nsq.Message) error {
	//fmt.Println(msg.Attempts)
	//fmt.Println(msg.ID)
	//fmt.Println(msg.NSQDAddress)
	//fmt.Println(msg.Timestamp)
	//fmt.Println(string(msg.Body))
	fmt.Println(string(msg.Body))
	//msg.WriteTo(os.Stdout)

	return nil
}

type ConsumerCCC struct {

}

func (*ConsumerCCC) HandleMessage(msg *nsq.Message) error {
	//fmt.Println(msg.Attempts)
	//fmt.Println(msg.ID)
	//fmt.Println(msg.NSQDAddress)
	//fmt.Println(msg.Timestamp)
	//fmt.Println(string(msg.Body))
	fmt.Println(string(msg.Body))
	//msg.WriteTo(os.Stdout)

	return nil
}


func NewConsumer(topic string,channel string) *nsq.Consumer{
	c, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	if err != nil{
		logger.Error(err)
	}
	return c
}



func HandlerMessage(){
	c := NewConsumer("unichain","aaa")
	c.AddHandler(&ConsumerT{})
	err := c.ConnectToNSQD("127.0.0.1:4150")
	if err != nil{
		logger.Error(err)
	}
}

func MutiChannel(){
	c := NewConsumer("unichain","11111")
	c.AddHandler(&ConsumerAAA{})
	err := c.ConnectToNSQD("127.0.0.1:4150")
	if err != nil{
		logger.Error(err)
	}

	c1 := NewConsumer("unichain","22222")
	c1.AddHandler(&ConsumerBBB{})
	err1 := c1.ConnectToNSQD("127.0.0.1:4150")
	if err1 != nil{
		logger.Error(err1)
	}

	c2 := NewConsumer("unichain","33333")
	c2.AddHandler(&ConsumerCCC{})
	err2 := c2.ConnectToNSQD("127.0.0.1:4150")
	if err2 != nil{
		logger.Error(err2)
	}

	c3 := NewConsumer("unichain","33333")
	c3.AddHandler(&ConsumerCCC{})
	err3 := c3.ConnectToNSQD("127.0.0.1:4150")
	if err3 != nil{
		logger.Error(err2)
	}

}

func AddHandler(){
	c := NewConsumer("unichain","ssssss")
	c.AddHandler(&ConsumerT{})
	err := c.ConnectToNSQD("127.0.0.1:4150")
	if err != nil{
		logger.Error(err)
	}
	c.ConnectToNSQLookupd("")
	c.AddHandler(
		nsq.HandlerFunc(
			func (m * nsq.Message) error{
				fmt.Println("a")
				return nil
			},
		))
}


func Test3(){
	p,err1 :=nsq.NewProducer("127.0.0.1:4150",nsq.NewConfig())
	if err1 != nil{
		logger.Error(err1)
	}
	err2 :=p.Publish("secn",[]byte("secn"))

	if err2 != nil {
		logger.Error(err2)
	}


	c, err := nsq.NewConsumer("secn", "aa", nsq.NewConfig())
	if err != nil{
		logger.Error(err)
	}

	c.AddHandler(&ConsumerAAA{})
	err3 := c.ConnectToNSQD("127.0.0.1:4150")
	if err3 != nil{
		logger.Error(err3)
	}

}
func Test4(){
	c1, err := nsq.NewConsumer("secnsss", "dd", nsq.NewConfig())
	if err != nil{
		logger.Error(err)
	}

	c1.AddHandler(nsq.HandlerFunc(func (message *nsq.Message) error{

		fmt.Println(string(message.Body))

		return nil
	}))
	err3 := c1.ConnectToNSQD("127.0.0.1:4150")
	if err3 != nil{
		logger.Error(err3)
	}
	logger.Debug("ssssssssssssssss")
}

func Test5(){
	Test4()
	logger.Debug("aaaaaaaaaaaaaaaaa")
}

func Test6(){
	Test5()
}


