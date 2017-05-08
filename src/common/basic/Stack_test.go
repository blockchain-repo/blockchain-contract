package basic

import (
	"testing"
	"fmt"
	"strconv"
)

func TestNewStack(t *testing.T){
	var t_stack *Stack = NewStack()
	if t_stack == nil {
		t.Error("NewStack Error!")
	}
}

func TestStackPush(t *testing.T){
	var t_stack *Stack = NewStack()
	t_stack.Push(100)
	t_stack.Push("testing")
	if t_stack == nil || t_stack.Len() == 0 {
		t.Error("StackPush Error!")
	}
	fmt.Println("Stack Length: " + strconv.Itoa(t_stack.Len()))
}

func TestStackEmpty(t *testing.T){
	var t_stack *Stack = NewStack()
	if !t_stack.Empty() {
		t.Error("Stack Empty Error!")
	}
}

func TestStackPop(t *testing.T){
	var t_stack *Stack = NewStack()
	t_stack.Push(100)
	t_stack.Push(200)
	var t_element int = t_stack.Pop().(int)
	if t_stack.Len() != 1 || t_element != 200 {
		t.Error("Stack Pop Error!")
	}
}

func TestStackPeak(t *testing.T){
	var t_stack *Stack = NewStack()
	t_stack.Push(100)
	t_stack.Push(200)
	var t_element int = t_stack.Peak().(int)
	if t_stack.Len() != 2 || t_element != 200 {
		t.Error("Stack Peak Error!")
	}
}

