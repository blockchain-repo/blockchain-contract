package chain

import (
	"testing"
	"os"
	"fmt"
	"unicontract/src/common/seelog"
)

/**
 * function:
 * param :
 * return nil:
 */
func TestGetValue(t *testing.T) {
	CreatTransaction()
}

func TestWeather(t *testing.T) {
	Weather()
}
func TestToday(t *testing.T) {
	Today()
}

func TestTest(t *testing.T) {
	logxmlpath := os.Getenv("CONFIGPATH")
	logxmlpath = logxmlpath + "/seelog.xml"
	fmt.Println(logxmlpath)
	requestPath := os.Getenv("CONFIGPATH")
	requestPath = requestPath + "/requestConfig.yaml"
	fmt.Println(requestPath)
	seelog.Debug("ssssssssssssssssss")
	Test()
}