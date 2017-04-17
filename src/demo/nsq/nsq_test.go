package nsq

import (
	"time"
	"testing"
)

func TestNewProducer(t *testing.T) {
	Publish()
	//HandlerMessage()
	time.Sleep(time.Second * 3)
}
func TestHandlerMessage(t *testing.T) {
	//Publish()
	HandlerMessage()
	time.Sleep(time.Second * 3)
}

func TestMutiChannel(t *testing.T) {
	//Publish()
	MutiChannel()
	time.Sleep(time.Second * 3)
}
func TestTest3(t *testing.T) {
	Test3()

	Test4()
}
func TestTest4(t *testing.T) {
	Test4()
	Test5()
	Test6()

}

func TestAddHandler(t *testing.T) {
	AddHandler()
}

//func TestPublish(t *testing.T) {
//	p := NewProducer("127.0.0.1")
//	Publish(p,"wsp","wsp")
//	Publish(p,"hcy","hcy")
//
//	p1 := NewProducer("192.168.80.40")
//
//	Publish(p1,"wsp1","wsp1")
//	Publish(p1,"hcy1","hcy1")
//}
//func TestTest1(t *testing.T) {
//	Test1()
//}

