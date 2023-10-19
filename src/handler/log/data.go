package log

import (
	"encoding/json"
	"fmt"
	"path"
	"runtime"
	"serviceX/src/config"
	"strconv"
	"time"
)

type LogEntry struct {
	Service   string
	Caller    string
	Message   string
	Timestamp time.Time
}

func NewLog(message string) *LogEntry {
	_, file, line, _ := runtime.Caller(2)
	return &LogEntry{
		Service:   config.Get().Name,
		Caller:    path.Base(file) + ":" + strconv.Itoa(line),
		Message:   message,
		Timestamp: time.Now(),
	}
}

func (l *LogEntry) String() string {
	return fmt.Sprintf("[%s] Service: %s, CodeLine: %s , Message: %s", l.Timestamp.Format("2006-01-02T15:04:05Z"), l.Service, l.Caller, l.Message)
}
func (l *LogEntry) Marshal() (data []byte) {
	data, _ = json.Marshal(l)
	return
}
