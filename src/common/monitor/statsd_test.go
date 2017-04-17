package monitor

import (
	"testing"
	"fmt"
	"time"
)

/**
 * function : 测试获取配置文件信息
 */
func TestGetMonitorConfig(t *testing.T) {

	monitorConfig := _GetMonitorConfig()
	fmt.Println(monitorConfig)
}

/**
 * function : 测试 将yaml中信息取出来
 */
func TestGetMonitorParam(t *testing.T){

	address,prefix,simpleRate,flushTime := _GetMonitorParam()
	fmt.Println(address)
	fmt.Println(prefix)
	fmt.Println(simpleRate)
	fmt.Println(flushTime)
}

/**
 * function : 测试 将yaml中信息取出来
 */
func TestGetMonitor(t *testing.T) {
	monitor := _GetMonitor()
	fmt.Println(monitor)
	monitor1 := _GetMonitor()
	fmt.Println(monitor1)
	monitor2 := _GetMonitor()
	fmt.Println(monitor2)
	monitor3 := _GetMonitor()
	fmt.Println(monitor3)
}

/**
 * function : 测试 单例模式
 */
func TestGetMonitorClient(t *testing.T) {
	monitor := _GetMonitorClient()
	fmt.Println(monitor)

	monitor1 := _GetMonitorClient()
	fmt.Println(monitor1)

	monitor2 := _GetMonitorClient()
	fmt.Println(monitor2)
}

/**
 * function : 测试 init模式
 */
func TestInit(t *testing.T) {
	fmt.Println(Monitor)
	fmt.Println(Monitor)
	fmt.Println(Monitor)
	fmt.Println(Monitor)
}

/**
 * function : 测试 Monitor方法
 */
func TestMonitor(t *testing.T) {
	Monitor.Gauge("GagueTest",1)
	Monitor.Count("CountTest",1)
	Monitor.Histogram("Histogram",1)
	Monitor.Increment("INcrementTest")
	Monitor.NewTiming().Send("NewTime")
	Monitor.Timing("TimingTest",2)
	Monitor.Unique("UniqueTest","ss")
}

/**
 * function : 测试 Gauge方法
 */
func TestGague(t *testing.T){
	fmt.Println("begin:",time.Now())
	for i := 1; i <= 100000; i++ {
		Monitor.Gauge("GagueTest",i)
		Monitor.Gauge("GagueTest2",100000-i)
	}
	fmt.Println("end:",time.Now())
}

/**
 * function : 测试 Count方法
 */
func TestCount(t *testing.T){
	fmt.Println("begin:",time.Now())
	for i := 1; i <= 100000; i++ {
		Monitor.Count("CountTest3",i)
		Monitor.Count("CountTest4",100000-i)
	}
	fmt.Println("end:",time.Now())
}

/**
 * function : 测试 Count方法
 */
func TestIncrement(t *testing.T){
	fmt.Println("begin:",time.Now())
	for i := 1; i <= 100000; i++ {
		Monitor.Increment("IncrementTest")
	}
	fmt.Println("end:",time.Now())
}

/**
 * function : 测试 方法
 */
func TestTest(t *testing.T){
	fmt.Println("begin:",time.Now())
	for i := 1; i <= 10000000; i++ {
		Monitor.Gauge("Twa",i)
	}
	fmt.Println("end:",time.Now())

}












