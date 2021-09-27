package loger

import (
	// "errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

type LogLevel uint16

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
)

type Loger struct {
	// 设置等级则输出对应等级日志
	// 例如:
	// log := loger.NewLoger()
	// log.Level = loger.WARNING
	// 输出高于warning的日志
	Level  LogLevel  //设置打印日志等级
	ToFile bool      //是否写入文件
	OutPut io.Writer //终端打印 os.Stdout  或输入文件
	FileLoger
}

func NewLoger() *Loger {
	return &Loger{}
}

func GetFotmatTime() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}

func getLogDetails() (funcName, file string, line int) {
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		fmt.Println("Runtime.caller faild ")
		return
	}

	funcName = runtime.FuncForPC(pc).Name()
	return
}

func (l *Loger) log(logLevel LogLevel, msg string, a ...interface{}) {
	msg = fmt.Sprintf(msg, a...)
	funcName, fileName, line := getLogDetails()

	if l.ToFile {
		l.OutPut = l.FileLoger.fileObj
		l.FileLoger.NewFileLoger()
	} else {
		l.OutPut = os.Stdout
	}
	if l.Level <= logLevel {
		fmt.Fprintf(l.OutPut, "[%s] [%s] [%s:%s:%d] %s\n", GetFotmatTime(), ParseLogLevel(logLevel), funcName, fileName, line, msg)

	}
}

func (l *Loger) Debug(msg string, a ...interface{}) {
	l.log(DEBUG, msg, a...)
}

func (l *Loger) Info(msg string, a ...interface{}) {
	l.log(INFO, msg, a...)
}

func (l *Loger) Warning(msg string, a ...interface{}) {
	l.log(WARNING, msg, a...)
}

func (l *Loger) Error(msg string, a ...interface{}) {
	l.log(ERROR, msg, a...)
}

func (l *Loger) Fatal(msg string, a ...interface{}) {
	l.log(FATAL, msg, a...)
}
