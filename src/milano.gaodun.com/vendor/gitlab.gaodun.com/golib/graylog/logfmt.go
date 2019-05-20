package graylog

import (
	"io"
	"log"
	"os"
)


type MangoLogger struct {
	Debug  bool
	out    io.Writer
	addLog func(info ...interface{})
	*log.Logger
	Info interface{}
}

func (nl *MangoLogger) Write(p []byte) (int, error) {
	if nl.Debug {
		return nl.out.Write(p)
	}

	go nl.addLog(string(p))
	return 0, nil
}

func (nl *MangoLogger) SetMangoOutput(f func(info ...interface{})) {
	nl.addLog = f
}


func New(debug bool, f func(info ...interface{})) *MangoLogger {
	l := log.New(os.Stderr, "", log.LstdFlags|log.Llongfile)
	nl := MangoLogger{out: os.Stderr, Debug: debug}
	nl.SetMangoOutput(f)
	l.SetOutput(&nl)
	nl.Logger = l
	return &nl
}

func NewSQL(debug bool, f func(info ...interface{})) MangoLogger {
	nl := MangoLogger{out: os.Stderr, Debug: debug}
	nl.SetMangoOutput(f)
	return nl
}

var mangoLog = New(false, AddGrayLog)
