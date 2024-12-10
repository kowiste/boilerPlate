package openob

import (
	"bytes"
	"context"
	"ddd/shared/logger"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Logger struct {
	serviceName   string
	environment   string
	endpoint      string
	headers       string
	orgID         string
	streamName    string
	httpClient    *http.Client
	minLevel      logger.Level
	consoleOutput bool
	enableTracing bool
}

type Config struct {
	ServiceName   string
	Environment   string
	Endpoint      string // base URL
	Headers       string
	OrgID         string // Usually "default"
	StreamName    string // Your stream name
	MinLevel      logger.Level
	ConsoleOutput bool
	EnableTracing bool
}

func NewLogger(cfg Config) (*Logger, error) {
	return &Logger{
		serviceName:   cfg.ServiceName,
		environment:   cfg.Environment,
		endpoint:      cfg.Endpoint,
		headers:       cfg.Headers,
		orgID:         cfg.OrgID,
		streamName:    cfg.StreamName,
		httpClient:    &http.Client{Timeout: 5 * time.Second},
		minLevel:      cfg.MinLevel,
		consoleOutput: cfg.ConsoleOutput,
		enableTracing: cfg.EnableTracing,
	}, nil
}

func (l *Logger) shouldLog(level logger.Level) bool {
	return level >= l.minLevel
}

type LogRecord struct {
	Stream string         `json:"stream"`
	Values map[string]any `json:"values"`
}

// convertToMap converts any struct or map to map[string]any
func convertToMap(v any) map[string]any {
	if v == nil {
		return nil
	}

	switch val := v.(type) {
	case map[string]any:
		return val
	default:
		// Convert struct to map using json marshaling
		data, err := json.Marshal(val)
		if err != nil {
			return nil
		}

		var result map[string]any
		if err := json.Unmarshal(data, &result); err != nil {
			return nil
		}
		return result
	}
}

func formatFields(fields map[string]any) string {
	if fields == nil {
		return ""
	}

	var fieldParts []string
	for k, v := range fields {
		fieldParts = append(fieldParts, fmt.Sprintf("%v:%v", k, v))
	}

	if len(fieldParts) == 0 {
		return ""
	}

	return fmt.Sprintf("> %s <", strings.Join(fieldParts, " "))
}

func (l *Logger) sendLog(ctx context.Context, level logger.Level, msg string, err error, fields any) {
	if !l.shouldLog(level) {
		return
	}

	// Convert fields to map
	fieldMap := convertToMap(fields)

	// Create log values
	values := map[string]any{
		"level":       level.String(),
		"message":     msg,
		"time":        time.Now().UTC().Format(time.RFC3339Nano),
		"service":     l.serviceName,
		"environment": l.environment,
	}

	if err != nil {
		values["error"] = err.Error()
	}

	// Add fields
	for k, v := range fieldMap {
		values[k] = v
	}

	// Console output if enabled
	if l.consoleOutput {
		logStr := fmt.Sprintf("[%s] %s [%s] %s",
			strings.ToUpper(level.String()),
			values["time"],
			l.serviceName,
			msg,
		)
		if err != nil {
			logStr += fmt.Sprintf(" | Error: %v", err)
		}
		if fieldMap != nil {
			logStr += fmt.Sprintf(" | Fields: %s", formatFields(fieldMap))
		}
		fmt.Println(logStr)
	}

	// If tracing is not enabled, return after console output
	if !l.enableTracing {
		return
	}

	// Create the log record
	records := []any{values}

	jsonData, err := json.Marshal(records)
	if err != nil {
		if l.consoleOutput {
			fmt.Printf("Error marshaling log data: %v\n", err)
		}
		return
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		fmt.Sprintf("http://%s/api/%s/%s/_json", l.endpoint, l.orgID, l.streamName),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		if l.consoleOutput {
			fmt.Printf("Error creating request: %v\n", err)
		}
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", l.headers))
	req.Header.Set("Content-Type", "application/json")

	resp, err := l.httpClient.Do(req)
	if err != nil {
		if l.consoleOutput {
			fmt.Printf("Error sending log: %v\n", err)
		}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && l.consoleOutput {
		fmt.Printf("Error response from server: %d\n", resp.StatusCode)
	}
}

func (l *Logger) Info(ctx context.Context, msg string, fields any) {
	l.sendLog(ctx, logger.InfoLevel, msg, nil, fields)
}

func (l *Logger) Error(ctx context.Context, err error, msg string, fields any) {
	l.sendLog(ctx, logger.ErrorLevel, msg, err, fields)
}

func (l *Logger) Debug(ctx context.Context, msg string, fields any) {
	l.sendLog(ctx, logger.DebugLevel, msg, nil, fields)
}

func (l *Logger) Warn(ctx context.Context, msg string, fields any) {
	l.sendLog(ctx, logger.WarnLevel, msg, nil, fields)
}
