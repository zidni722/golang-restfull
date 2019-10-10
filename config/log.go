package config

import (
	"os"
	"time"

	"github.com/kataras/iris"
)

type Log struct{}

func (l *Log) TodayFileName() string {
	today := time.Now().Format("storage/logs/Jan-02-2006")
	return today + ".log"
}

func (l *Log) NewLogFile() *os.File {
	filename := l.TodayFileName()

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}

func (l *Log) Configure(app *iris.Application) {
	app.Logger().SetOutput(l.NewLogFile())
}
