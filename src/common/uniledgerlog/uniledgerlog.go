package uniledgerlog

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

import (
	"github.com/astaxie/beego/logs"
)

const (
	LevelError = iota
	LevelWarn
	LevelInfo
	LevelDebug
)

var mapLevelKeys = map[int]string{
	LevelError: "ERROR",
	LevelWarn:  "WARN",
	LevelInfo:  "INFO",
	LevelDebug: "DEBUG",
}

func Init() {
	// 待补充
}

func _FormatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}

func _WriteLog(key int, format interface{}, v ...interface{}) {
	defer logs.SetLogFuncCall(true)
	logs.SetLogFuncCall(false)
	pc, file, line, _ := runtime.Caller(2)
	func_ := runtime.FuncForPC(pc)
	var f func(f interface{}, v ...interface{})
	switch key {
	case LevelError:
		f = logs.Error
	case LevelWarn:
		f = logs.Warn
	case LevelInfo:
		f = logs.Info
	case LevelDebug:
		f = logs.Debug
	}
	slStr := strings.Split(func_.Name(), ".")
	f("[%s] [%s : %d (%s)] %s", mapLevelKeys[key],
		slStr[0]+string(filepath.Separator)+filepath.Base(file), line, slStr[1], _FormatLog(format, v...))
}

func Error(f interface{}, v ...interface{}) {
	_WriteLog(LevelError, f, v...)
}

func Warn(f interface{}, v ...interface{}) {
	_WriteLog(LevelWarn, f, v...)
}

func Info(f interface{}, v ...interface{}) {
	_WriteLog(LevelInfo, f, v...)
}

func Debug(f interface{}, v ...interface{}) {
	_WriteLog(LevelDebug, f, v...)
}
