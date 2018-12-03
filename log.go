package wst

import (
	"strings"

	"github.com/wsterr/wst/logs"
)

const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

func Emergency(v ...interface{}) {
	logs.Emergency(strings.Repeat("%v ", len(v)), v)
}

func Alert(v ...interface{}) {

}

func Critical(v ...interface{}) {

}

func Error(v ...interface{}) {

}

func Warn(v ...interface{}) {

}

func Notice(v ...interface{}) {

}

func Info(v ...interface{}) {

}

func Debug(v ...interface{}) {

}
