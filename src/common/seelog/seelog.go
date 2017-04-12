package seelog

import (
	"github.com/cihub/seelog"
	"os"
)

var logger seelog.LoggerInterface

func init(){
	logxmlpath := os.Getenv("REQUEST")
	logxmlpath = logxmlpath + "conf/seelog.xml"
	newLogger, err := seelog.LoggerFromConfigAsFile(logxmlpath)

	if err != nil {
		seelog.Critical("err parsing config log file", err)
		return
	}
	logger = newLogger
	logger.SetAdditionalStackDepth(1)
}

func Trace(v ...interface{}){
	defer logger.Flush()
	logger.Trace(v)
}

func Tracef(format string, params ...interface{}){
	defer logger.Flush()
	logger.Tracef(format,params)
}

func Debug(v ...interface{}){
	defer logger.Flush()
	logger.Debug(v)
}

func Debugf(format string, params ...interface{}){
	defer logger.Flush()
	logger.Debugf(format,params)
}

func Info(v ...interface{}){
	defer logger.Flush()
	logger.Info(v)
}

func Infof(format string, params ...interface{}){
	defer logger.Flush()
	logger.Infof(format,params )
}


func Warn(v ...interface{}){
	defer logger.Flush()
	logger.Warn(v)
}

func Warnf(format string, params ...interface{}){
	defer logger.Flush()
	logger.Warnf(format,params)
}

func Error(v ...interface{}){
	defer logger.Flush()
	logger.Error(v)
}

func Errorf(format string, params ...interface{}){
	defer logger.Flush()
	logger.Errorf(format,params)
}

func Critical(v ...interface{}){
	defer logger.Flush()
	logger.Critical(v)
}

func Criticalf(format string, params ...interface{}){
	defer logger.Flush()
	logger.Criticalf(format,params)
}

func Close(){
	logger.Close()
}

func Closed()  bool{
	return logger.Closed()
}
