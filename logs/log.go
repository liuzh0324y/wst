package logs

import "time"

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

type Logger interface {
	Init(config string) error
	WriteMsg(when time.Time, msg string, level int) error
	Destroy()
	Flush()
}

func Emergency(f interface{}, v ...interface{}) {

}

func Alert(f interface{}, v ...interface{}) {

}

func Critical(f interface{}, v ...interface{}) {

}

func Error(f interface{}, v ...interface{}) {

}

func Warn(f interface{}, v ...interface{}) {

}

func Notice(f interface{}, v ...interface{}) {

}

func Info(f interface{}, v ...interface{}) {

}

func Debug(f interface{}, v ...interface{}) {

}
