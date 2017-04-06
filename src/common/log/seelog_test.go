package log

import (
	"testing"
	"fmt"
	"time"
)

func TestTrace(t *testing.T) {
	Trace("trace test....")
}

func TestTracef(t *testing.T) {
	Tracef("测试%s","trace")
}


func TestDebuge(t *testing.T) {
	Debug("debuge test....")
}

func TestDebugef(t *testing.T) {
	Debugf("测试%s","debuge")
}


func TestInfo(t *testing.T) {
	for ; ; {
		Info("Info test....")
		Debug("debuge test....")
		time.Sleep(time.Second * 1)
	}
}

func TestInfof(t *testing.T) {
	Infof("测试%s","Infof")
}


func TestWarn(t *testing.T) {
	Warn("Warn test....")
}

func TestWarnf(t *testing.T) {
	Warnf("测试%s","Warnf")
}


func TestError(t *testing.T) {
	Error("Error test....")
}

func TestErrorf(t *testing.T) {
	Errorf("测试%s","Errorf")
}


func TestCritical(t *testing.T) {
	Critical("Critical test....")
}

func TestCriticalf(t *testing.T) {
	Criticalf("测试%s","Criticalf")
}

func TestClosed(t *testing.T) {
	Info("begin")
	fmt.Println("关闭了吗?",Closed())
	Close()
	fmt.Println("关闭了吗?",Closed())
	Info("end")
}

func TestClose(t *testing.T) {
	Info("SSSS")
	Close()
	Info("CCCC")
}
