package main

import (
	"io"
	"log"
	"os"
)

type Loglevel int

const (
	LogLevelError = iota
	LogLevelWarning
	LogLevelInfo
)

type LogExtended struct {
	*log.Logger
	loglevel Loglevel
}

func NewLogExtended(out io.Writer, prefix string, flag int, lvl Loglevel) *LogExtended {
	return &LogExtended{
		Logger:   log.New(out, prefix, flag), //Тут надо писать имя типа без имени пакета!!
		loglevel: lvl,
	}
}

func (l *LogExtended) SetLogLevel(logLvl Loglevel) {
	l.loglevel = logLvl
}

func (l *LogExtended) Infoln(msg string) {
	if l.loglevel >= 2 {
		l.Println("INFO " + msg)
	}
}

func (l *LogExtended) Warnln(msg string) {
	if l.loglevel >= 1 {
		l.Println("WARN " + msg)
	}
}

func (l *LogExtended) Errorln(msg string) {
	if l.loglevel >= 0 {
		l.Println("ERROR " + msg)
	}
}

func (l *LogExtended) println(srcLogLvl Loglevel, prefix, msg string) {
	// игнорируем сообщения, если уровень логгера меньше scrLogLvl
	if srcLogLvl <= l.loglevel {
		l.Logger.Println(prefix + msg)
	}
}

func main() {
	logger := log.New(os.Stdout, "std logger: ", log.Ldate|log.Ltime)
	logger.Println("Test")
	logger.Println("Test2")

	logger2 := NewLogExtended(os.Stdout, "std logger: ", log.Ldate|log.Ltime, LogLevelInfo)
	logger2.Println("Test3")
	logger2.Infoln("Должно написаться") // loglevel = LogLevelInfo
	logger2.SetLogLevel(LogLevelError)
	logger2.Warnln("Не должно написаться") // loglevel = LogLevelError
	logger2.Errorln("Должно написаться")   // loglevel = LogLevelError
	logger2.println(0, "SECRET ", "должно написаться")
	logger2.println(1, "SECRET ", "не должно написаться")
}
