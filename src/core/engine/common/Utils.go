package common

import (
	"encoding/json"
	"log"
	"bytes"
	"time"
	"strconv"
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
	return strconv.FormatInt(millis, 10)
}

func Serialize(obj interface{}) string {
	str, err := json.Marshal(obj)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return string(str)
}

//only for selfTest, format json output
func SerializePretty(obj interface{}) string {
	input, err := json.Marshal(obj)
	if err != nil {
		log.Fatalf(err.Error())
	}
	var out bytes.Buffer
	err = json.Indent(&out, input, "", "\t")

	if err != nil {
		log.Fatalf(err.Error())
	}
	return string(out.String())
}

func Deserialize(jsonStr string) interface{} {
	var dat interface{}
	err := json.Unmarshal([]byte(jsonStr), &dat)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return dat
}

func TernaryOperator(p_cond bool, p_true_cond interface{}, p_false_cond interface{}) interface{} {
	if p_cond {
		return p_true_cond
	}
	return p_false_cond
}