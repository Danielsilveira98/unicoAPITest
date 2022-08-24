package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
)

type log struct {
	Time     *time.Time             `json:"time,omitempty"`
	Level    string                 `json:"level,omitempty"`
	Msg      string                 `json:"msg,omitempty"`
	TraceID  string                 `json:"trace_id,omitempty"`
	MetaData map[string]interface{} `json:"meta_data,omitempty"`
}

var jErrBase = `{"time": "%v", "level": "%s", "msg" "Error on marshal log: %s", "trace_id": "%s"}`

type Logger struct {
	writer io.Writer
	pretty bool
}

func NewLogger(writer io.Writer, pretty bool) *Logger {
	return &Logger{writer, pretty}
}

func (l *Logger) Info(ctx context.Context, msg string) {
	l.print(ctx, domain.LogLevelInfo, msg, nil)
}
func (l *Logger) Debug(ctx context.Context, msg string) {
	l.print(ctx, domain.LogLevelDebug, msg, nil)
}
func (l *Logger) Warn(ctx context.Context, msg string) {
	l.print(ctx, domain.LogLevelWarning, msg, nil)
}
func (l *Logger) Error(ctx context.Context, msg string) {
	l.print(ctx, domain.LogLevelError, msg, nil) // TODO write with an other writer
}
func (l *Logger) Infof(ctx context.Context, msg string, md map[string]interface{}) {
	l.print(ctx, domain.LogLevelInfo, msg, md)
}
func (l *Logger) Debugf(ctx context.Context, msg string, md map[string]interface{}) {
	l.print(ctx, domain.LogLevelDebug, msg, md)
}
func (l *Logger) Warnf(ctx context.Context, msg string, md map[string]interface{}) {
	l.print(ctx, domain.LogLevelWarning, msg, md)
}
func (l *Logger) Errorf(ctx context.Context, msg string, md map[string]interface{}) {
	l.print(ctx, domain.LogLevelError, msg, md) // TODO write with an other writer
}

func (l *Logger) print(ctx context.Context, lvl string, msg string, metaData map[string]interface{}) {
	now := time.Now()
	traceID := getTraceID(ctx)

	log := log{
		Time:     &now,
		Level:    lvl,
		Msg:      msg,
		TraceID:  getTraceID(ctx),
		MetaData: metaData,
	}

	var err error
	var j []byte
	if l.pretty {
		j, err = json.MarshalIndent(log, "", "\t")
	} else {
		j, err = json.Marshal(log)
	}

	if err != nil {
		jErr := fmt.Sprintf(jErrBase, now, domain.LogLevelError, err.Error(), traceID)
		j = []byte(jErr)
	}

	l.writer.Write(append(j, []byte("\n")...))
}

func getTraceID(ctx context.Context) string {
	return ctx.Value(domain.TraceIDCtxKey).(string)
}
