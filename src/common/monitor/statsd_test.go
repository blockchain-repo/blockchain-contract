package monitor

import (
	"testing"
	"fmt"
	"time"
	"math/rand"
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

	address,prefix,maxPacketSize := _GetMonitorParam()
	fmt.Println(address)
	fmt.Println(prefix)
	fmt.Println(maxPacketSize)
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
func Test_GetMonitorClient(t *testing.T) {
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
		Monitor.Count("CountTest1",i)
		Monitor.Count("CountTest2",100000-i)
	}
	fmt.Println("end:",time.Now())
}

/**
 * function : 测试 Count方法
 */
func TestIncrement(t *testing.T){
	fmt.Println("begin:",time.Now())
	for i := 1; i <= 10000; i++ {
		Monitor.Increment("IncrementTest")
	}
	fmt.Println("end:",time.Now())
}

/**
 * function : 测试 方法
 */
func TestTest(t *testing.T){
	fmt.Println("begin:",time.Now())
	for i := 1; i <= 1000; i++ {
		Monitor.Gauge("Twa",20000)
		Monitor.Timing("TimeTest",11)
	}
	fmt.Println("end:",time.Now())

}

/**
 * function : 测试 方法
 */
func TestTiming(t *testing.T){
	for i := 1;i<=1000;i++ {
		begin := time.Now().UnixNano()/1000000
		end := time.Now().UnixNano()/1000000
		time := end - begin
		Monitor.Timing("TimeTest",time)
	}
}

/**
 * function : 测试 方法
 */
func TestNewTiming(t *testing.T){
	for i:=1;i<=100;i++ {
		Monitor.Increment("WJT")
		Monitor.Flush()
	}
}

/**
 * function : 测试 方法
 */
func TestGaGUE(t *testing.T){

	for i:=1;i<=6000;i++ {
		if i == 6000{
			fmt.Println(time.Now())
			Monitor.Gauge("unichain",rand.Int())
		}
	}

}

/**
 * function : 测试 方法
 */
func TestGague1(t *testing.T){
	begin := time.Now().UnixNano()/1000000
	fmt.Println("begin:",time.Now().UnixNano()/1000000)
	for i:=1;i<=10000;i++ {
		Monitor.Increment("ssss")
		Monitor.Gauge("zzz",i)
	}
	fmt.Println("end:",time.Now().UnixNano()/1000000)
	end := time.Now().UnixNano()/1000000
	fmt.Println(end -begin)
}

/**
 * function : 测试 方法
 */
func TestTime2(t *testing.T){

	monitorFunc := func(){
		defer Monitor.NewTiming().Send("AA")
	}

	monitorFunc()
}


/**
 * function : 测试 方法
 */
func TestTime3(t *testing.T){
	monitor1 := _GetMonitorClient()
	//begin := time.Now().UnixNano()/1000000
	//end := time.Now().UnixNano()/1000000
	//timeT := end - begin
	monitor1.Timing("aaaaaaaaaaaaa",1)
	monitor1.Timing("aaaaaaaaaaaaa",2)
	monitor1.Timing("aaaaaaaaaaaaa",3)
	monitor1.Flush()
	monitor1.Timing("aaaaaaaaaaaaa",4)
	monitor1.Timing("aaaaaaaaaaaaa",5)
	monitor1.Timing("aaaaaaaaaaaaa",6)
	monitor1.Flush()
	fmt.Println(time.Now())
}

/**
 * function : 测试 方法
 */
func TestTime(t *testing.T){

}

/**
 * function : 测试 方法
 */
func TestTime5(t *testing.T){
	monitor1 := _GetMonitorClient()
	for i:=1;i<=5;i++ {

		monitor1.Count("CCCCCCCCCCCCCCCCCC",-2)
	}

	monitor1.Flush()

}

/**
 * function : 测试 方法
 */
func TestTime6(t *testing.T){
	monitor1 := _GetMonitorClient()
	for i:=1;i<=100000;i++ {
		monitor1.Unique("AAA","SSSSSSSSS")
	}
	monitor1.Flush()

}

func TestTest1(t *testing.T){
	monitor1 := _GetMonitorClient()
	for i:=1;i<=1000;i++ {
		monitor1.Gauge("HAHAHAHAHAHAHAHAHAHA",1)
		monitor1.Gauge("HEHEHEHEHEHHEHEHEHHE",2)
	}
	monitor1.Flush()

}




