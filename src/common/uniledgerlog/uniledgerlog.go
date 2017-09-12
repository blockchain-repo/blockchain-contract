package uniledgerlog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

import (
	"unicontract/src/common/basic"
)

const (
	LevelError = iota
	LevelWarn
	LevelNotice
	LevelInfo
	LevelDebug
)

var mapLevelKeys = map[int]string{
	LevelError:  "ERROR",
	LevelWarn:   "WARN",
	LevelNotice: "NOTICE",
	LevelInfo:   "INFO",
	LevelDebug:  "DEBUG",
}

func _Serialize(obj interface{}, escapeHTML ...bool) string {
	setEscapeHTML := false
	if len(escapeHTML) >= 1 {
		setEscapeHTML = escapeHTML[0]
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	// disabled the HTMLEscape for &, <, and > to \u0026, \u003c, and \u003e in json string
	enc.SetEscapeHTML(setEscapeHTML)
	err := enc.Encode(obj)
	if err != nil {
		logs.Error(err.Error())
		return ""
	}
	return strings.TrimSpace(buf.String())
	//return strings.Replace(strings.TrimSpace(buf.String()), "\n", "", -1)
}

func Init() {
	logs.SetLogFuncCall(false)

	myBeegoLogAdapterMultiFile := &basic.MyBeegoLogAdapterMultiFile{}
	myBeegoLogAdapterMultiFile.FileName = beego.AppConfig.String("LogName")
	myBeegoLogAdapterMultiFile.Level, _ = beego.AppConfig.Int("LogSaveLevel")
	logMaxDays, _ := beego.AppConfig.Int("LogMaxDays")
	myBeegoLogAdapterMultiFile.MaxDays = int16(logMaxDays)
	myBeegoLogAdapterMultiFile.MaxLines, _ = beego.AppConfig.Int64("LogMaxLines")
	myBeegoLogAdapterMultiFile.MaxSize, _ = beego.AppConfig.Int64("LogMaxSize")
	myBeegoLogAdapterMultiFile.Rotate, _ = beego.AppConfig.Bool("LogRotate")
	myBeegoLogAdapterMultiFile.Daily, _ = beego.AppConfig.Bool("LogDaily")
	myBeegoLogAdapterMultiFile.Separate = beego.AppConfig.Strings("LogSeparate")

	log_config := basic.NewMyBeegoLogAdapterMultiFile(myBeegoLogAdapterMultiFile)
	log_config_str := _Serialize(log_config)
	//fmt.Println(log_config_str)

	// order 顺序必须按照
	// 1. logs.SetLevel(level)
	// 2. logs.SetLogger(logs.AdapterMultiFile, log_config_str)
	logLevel, _ := beego.AppConfig.Int("LogLevel")
	logs.SetLevel(logLevel)
	logs.SetLogger(logs.AdapterMultiFile, log_config_str)
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
	pc, file, line, _ := runtime.Caller(2)
	func_ := runtime.FuncForPC(pc)
	var f func(f interface{}, v ...interface{})
	switch key {
	case LevelError:
		f = logs.Error
	case LevelWarn:
		f = logs.Warn
	case LevelNotice:
		f = logs.Notice
	case LevelInfo:
		f = logs.Info
	case LevelDebug:
		f = logs.Debug
	}
	slStr := strings.Split(func_.Name(), ".")
	f("[%s] [%s : %d (%s)] %s", mapLevelKeys[key],
		slStr[0]+string(filepath.Separator)+filepath.Base(file), line, slStr[len(slStr)-1], _FormatLog(format, v...))
}

func Error(f interface{}, v ...interface{}) {
	_WriteLog(LevelError, f, v...)
}

func Warn(f interface{}, v ...interface{}) {
	_WriteLog(LevelWarn, f, v...)
}

func Notice(f interface{}, v ...interface{}) {
	_WriteLog(LevelNotice, f, v...)
}

func Info(f interface{}, v ...interface{}) {
	_WriteLog(LevelInfo, f, v...)
}

func Debug(f interface{}, v ...interface{}) {
	_WriteLog(LevelDebug, f, v...)
}
