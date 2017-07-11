package uniledgerlog

import (
	"fmt"
	"testing"
)

func Test_log(t *testing.T) {
	var logType ErrorType
	logType = API_ERROR
	var logErrorMsg string
	logErrorMsg = "Token is error!"
	Error(fmt.Sprintf("[%s][%s]", logType, logErrorMsg))

	logType = DEBUG_NO_ERROR
	logErrorMsg = "debug no error!"
	Debug(fmt.Sprintf("[%s][%s]", logType, logErrorMsg))
}
