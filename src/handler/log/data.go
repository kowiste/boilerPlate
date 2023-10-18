package log

import (
	"fmt"
	"path"
	"serviceX/src/config"
	"strconv"
	"time"
)

type logEntry struct {
	Service   string
	Caller    string
	Message   string
	Timestamp time.Time
}

func (l *logEntry) Fill(data string) {
	file, line := getCaller()
	l.Message = data
	l.Service = config.Get().Name
	//removing anything behind src
	l.Caller = path.Base(file) + ":" + strconv.Itoa(line)
	// Set the date field.
	l.Timestamp = time.Now()
}

func (l *logEntry) String() string {
	return fmt.Sprintf("[%s] Service: %s, CodeLine: %s , Message: %s", l.Timestamp.Format("2006-01-02T15:04:05Z"), l.Service, l.Caller, l.Message)
}
