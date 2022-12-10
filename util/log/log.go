package log

import (
	"fmt"
	"time"
)

type Log interface {
	Info(logInfo string)
	Debug(logInfo string)
	Error(logInfo string)
	Infof(format string, item ...interface{})
	Debugf(format string, item ...interface{})
	Errorf(format string, item ...interface{})
}

type LocalLog struct {
}

func (g *LocalLog) Info(logInfo string) {
	fmt.Println(getCurrentTime() + "|info|" + logInfo)
}

func (g *LocalLog) Debug(logInfo string) {
	fmt.Println(getCurrentTime() + "|debug|" + logInfo)
}

func (g *LocalLog) Error(logInfo string) {
	fmt.Println(getCurrentTime() + "|error|" + logInfo)
}

func (g *LocalLog) Infof(format string, item ...interface{}) {
	_format := getCurrentTime() + "|info|" + format + "\n"
	fmt.Printf(_format, item...)
}

func (g *LocalLog) Debugf(format string, item ...interface{}) {
	_format := getCurrentTime() + "|debug|" + format + "\n"
	fmt.Printf(_format, item...)
}

func (g *LocalLog) Errorf(format string, item ...interface{}) {
	_format := getCurrentTime() + "|error|" + format + "\n"
	fmt.Printf(_format, item...)
}

func getCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
