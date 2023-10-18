package log

import (
	"io"
	"log"
	"runtime"
	"sync"
)

const (
	ErrorLevel Level = iota
	InfoLevel
)

var lock = &sync.Mutex{}
var singleInstance *logger

type Level int
type logger struct {
	level Level
}

func CreateInstance(level Level, outputs ...io.Writer) {
	lock.Lock()
	defer lock.Unlock()
	if singleInstance == nil {
		singleInstance = &logger{level: ErrorLevel}
		log.SetFlags(0)
		log.SetOutput(io.MultiWriter(outputs...))
	}
}
func Get() *logger {
	return singleInstance
}

func (l *logger) SetLevel(level Level) {
	lock.Lock()
	defer lock.Unlock()
	l.level = level
}

// SetOutputs 
func (l *logger) SetOutputs(output ...io.Writer) {
	lock.Lock()
	defer lock.Unlock()
	log.SetOutput(io.MultiWriter(output...))
}

func (l *logger) Print(level Level, message string) {
	lock.Lock()
	defer lock.Unlock()

	if level < l.level {
		return
	}
	outData := new(logEntry)
	outData.Fill(message)
	log.Println(outData)
}
func getCaller() (file string, line int) {
	_, file, line, _ = runtime.Caller(3)
	return
}
