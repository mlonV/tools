package loger

import (
	// "errors"
	"fmt"
	// "strings"
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
	Level LogLevel
}

func NewLoger() *Loger {
	return &Loger{}
}

func GetFotmatTime() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}

// func ParseLogLevel(s string) (LogLevel, error) {
// 	s = strings.ToLower(s)
// 	switch s {
// 	case "debug":
// 		return DEBUG, nil
// 	case "info":
// 		return INFO, nil
// 	case "warning":
// 		return WARNING, nil
// 	case "error":
// 		return ERROR, nil
// 	case "fatal":
// 		return FATAL, nil
// 	default:
// 		err := errors.New("unknown log level, default log level warning")
// 		return WARNING, err
// 	}
// }

func (l *Loger) Debug(str string) {
	if l.Level <= DEBUG {
		fmt.Printf("[%s] [DEBUG] %s\n", GetFotmatTime(), str)
	}
}

func (l *Loger) Info(str string) {
	if l.Level <= INFO {
		fmt.Printf("[%s] [Info] %s\n", GetFotmatTime(), str)
	}
}

func (l *Loger) Warning(str string) {
	if l.Level <= WARNING {
		fmt.Printf("[%s] [Warning] %s\n", GetFotmatTime(), str)
	}
}

func (l *Loger) Error(str string) {
	if l.Level <= ERROR {
		fmt.Printf("[%s] [Error] %s\n", GetFotmatTime(), str)
	}
}

func (l *Loger) Fatal(str string) {
	if l.Level <= FATAL {
		fmt.Printf("[%s] [Fatal] %s\n", GetFotmatTime(), str)
	}
}
