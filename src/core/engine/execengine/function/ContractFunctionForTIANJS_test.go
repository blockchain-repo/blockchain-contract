package function

import (
	"unicontract/src/common/uniledgerlog"
	"math/rand"
	"testing"
	"time"
)

func TestFuncGetNowDate(t *testing.T) {
	//format := "2006-01-02 15:04:05"
	//now := time.Now() //2017-06-17 15:19:30.930240121 +0800 CST
	/*
		now, _ := time.Parse(format, time.Now().Format(format)) //2017-06-17 15:19:58 +0000 UTC
		uniledgerlog.Info(now)
		a, _ := time.Parse(format, "2017-06-17 11:00:00")
		b, _ := time.Parse(format, "2017-06-17 16:00:00")
		fmt.Println(now.Weekday())
		fmt.Println(now.Format(format), a.Format(format), b.Format(format))
		fmt.Println(now.After(a))
		fmt.Println(now.Before(a))
		fmt.Println(now.After(b))
		fmt.Println(now.Before(b))
		fmt.Println(a.After(b))
		fmt.Println(a.Before(b))
		fmt.Println(now.Format(format), a.Format(format), b.Format(format))
		fmt.Println(now.Unix(), a.Unix(), b.Unix())
	*/

	var RaisePeriodFrom string = "2017-06-18 11:00:00"
	var RaisePeriodTo string = "2017-06-17 16:00:00"
	format := "2006-01-02 15:04:05"
	start, _ := time.Parse(format, RaisePeriodFrom)
	end, _ := time.Parse(format, RaisePeriodTo)
	now, _ := time.Parse(format, "2017-06-18 15:19:58")
	//now, _ := time.Parse(format, time.Now().Format(format))
	d, _ := time.ParseDuration("-24h")
	var dateCheck time.Time = now.Add(d)
	uniledgerlog.Info(dateCheck)
	var week = dateCheck.Weekday()
	uniledgerlog.Info(week)
	var sunday time.Weekday = 0
	var sturday time.Weekday = 6
	if sunday == week {
		uniledgerlog.Info(sunday)
		d, _ := time.ParseDuration("-48h")
		dateCheck = dateCheck.Add(d)
	}
	if sturday == week {
		uniledgerlog.Info(sturday)
		d, _ := time.ParseDuration("-24h")
		dateCheck = dateCheck.Add(d)
	}
	var isIn bool = dateCheck.After(start) && dateCheck.Before(end)
	uniledgerlog.Info(dateCheck, isIn)
}

func TestRand(t *testing.T) {
	var rate float64 = 0.03
	//var realrate = rand.Float64() * rate
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	realrate := (r.Float64() + 0.6) * rate
	uniledgerlog.Info(realrate)
}

//0.6046602879796196
//0.6046602879796196
