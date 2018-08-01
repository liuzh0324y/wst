package logs

import (
	"os"
	"time"
)

type consoleWriter struct {
	lw    *logWriter
	Level int
}

func NewConsole() Logger {
	return &consoleWriter{
		lw:    newLogWriter(os.Stdout),
		Level: LevelDebug,
	}
}

func (this *consoleWriter) Init(conf string) error {
	return nil
}

func (this *consoleWriter) WriteMsg(when time.Time, msg string, level int) error {
	return nil
}

func (this *consoleWriter) Destroy() {

}

func (this *consoleWriter) Flush() {

}
