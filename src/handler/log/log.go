package log

import (
	"log"
	"os"
	"sync"
)

const (
	InfoLevel Level = iota
	ErrorLevel
)

var lock = &sync.Mutex{}
var singleInstance *logger

type Level int
type logger struct {
	level Level
	ch    []chan *LogEntry
}

// CreateInstance
func CreateInstance(level Level) {
	lock.Lock()
	defer lock.Unlock()
	if singleInstance == nil {
		singleInstance = &logger{level: ErrorLevel}
		log.SetFlags(0)
		log.SetOutput(os.Stderr) //Set default log to terminal
	}
}
func Get() *logger {
	return singleInstance
}

// SetLevel set level of
func (l *logger) SetLevel(level Level) {
	lock.Lock()
	defer lock.Unlock()
	l.level = level
}

// SetChannels set the channels where stream the logs
func (l *logger) SetChannels(channels ...chan *LogEntry) {
	lock.Lock()
	defer lock.Unlock()
	l.ch = append(l.ch, channels...)
}

// Print  send the message to the terminal and all the set channels
func (l *logger) Print(level Level, message string) {
	lock.Lock()
	defer lock.Unlock()

	if level < l.level {
		return
	}
	outData := NewLog(message)
	log.Println(outData)      //print in terminal
	for _, ch := range l.ch { //stream to others loggers
		ch <- outData
	}

}
