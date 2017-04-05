package common

import (
	"encoding/json"
	"strconv"
	"time"
)

func GenDate() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05 PM")
}

func GenTimestamp() string {
	t := time.Now()
	nanos := t.UnixNano()
	millis := nanos / 1000000 //ms len=13
	return strconv.FormatInt(millis,10)
}

func Serialize(dat map[string]interface{}) string {
	str, err := json.Marshal(dat)
	if err != nil {
		panic(err)
	}
	return string(str)
}

func Deserialize(str string) map[string]interface{} {
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(str), &dat)
        if  err != nil {
		panic(err)
	}
	return dat
}
