package log

import (
	"serviceX/src/config"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLogger_Print(t *testing.T) {
	config.CreateInstance()
	// Create a new logger instance.
	CreateInstance(InfoLevel)

	// Create a test channel to capture log entries.
	testChannel := make(chan *LogEntry, 1)
	defer close(testChannel)

	// Add the test channel to the logger.
	Get().SetChannels(testChannel)

	// Create a message and log it.
	message := "Test log message"
	Get().Print(ErrorLevel, message)

	timeOut := time.NewTimer(5 * time.Second)
	// Check if the log entry was received on the test channel.
	select {
	case logEntry := <-testChannel:
		assert.NotNil(t, logEntry)
		//assert.Equal(t, InfoLevel, logEntry.)
		assert.Equal(t, message, logEntry.Message)
	case <-timeOut.C:
		t.Error("Expected log entry on the test channel, but got none.")
	}
}

func TestLogger_SetLocal(t *testing.T) {
	// Create a new logger instance.
	loggerInstance := &logger{}

	// Set local to true and then check it.
	loggerInstance.SetLocal(true)
	assert.True(t, loggerInstance.local)

	// Set local to false and then check it.
	loggerInstance.SetLocal(false)
	assert.False(t, loggerInstance.local)
}

func TestLogger_SetLevel(t *testing.T) {
	// Create a new logger instance.
	loggerInstance := &logger{}

	// Set the level to InfoLevel and then check it.
	loggerInstance.SetLevel(InfoLevel)
	assert.Equal(t, InfoLevel, loggerInstance.level)

	// Set the level to ErrorLevel and then check it.
	loggerInstance.SetLevel(ErrorLevel)
	assert.Equal(t, ErrorLevel, loggerInstance.level)
}
