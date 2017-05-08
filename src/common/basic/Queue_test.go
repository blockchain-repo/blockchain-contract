package basic

import (
	"testing"
	"fmt"
	"strconv"
)

func TestNewQueue(t *testing.T){
	var t_Queue *Queue = NewQueue()
	if t_Queue == nil {
		t.Error("NewQueue Error!")
	}
}

func TestQueuePush(t *testing.T){
	var t_Queue *Queue = NewQueue()
	t_Queue.Push(100)
	t_Queue.Push("testing")
	if t_Queue == nil || t_Queue.Len() == 0 {
		t.Error("QueuePush Error!")
	}
	fmt.Println("Queue Length: " + strconv.Itoa(t_Queue.Len()))
}

func TestQueueEmpty(t *testing.T){
	var t_Queue *Queue = NewQueue()
	if !t_Queue.Empty() {
		t.Error("Queue Empty Error!")
	}
}

func TestQueuePop(t *testing.T){
	var t_Queue *Queue = NewQueue()
	t_Queue.Push(100)
	t_Queue.Push(200)
	var t_element int = t_Queue.Pop().(int)
	if t_Queue.Len() != 1 || t_element != 100 {
		t.Error("Queue Pop Error!")
	}
}

func TestQueueFirst(t *testing.T){
	var t_Queue *Queue = NewQueue()
	t_Queue.Push(100)
	t_Queue.Push(200)
	var t_element int = t_Queue.First().(int)
	if t_Queue.Len() != 2 || t_element != 100 {
		t.Error("Queue Peak Error!")
	}
}