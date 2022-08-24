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
	Time       *time.Time             `json:"time,omitempty"`
	Level      string                 `json:"level,omitempty"`
	Msg        string                 `json:"msg,omitempty"`
	TraceID    string                 `json:"trace_id,omitempty"`
	MetaData   map[string]interface{} `json:"meta_data,omitempty"`
	StackTrace map[string]string      `json:"stack_trace,omitempty"`
}

type Logger struct {
	writer io.Writer
	pretty bool
}

func NewLogger(writer io.Writer, pretty bool) *Logger {
	return &Logger{writer, pretty}
}

func (l *Logger) Info(ctx context.Context, msg string) {
	l.Infof(ctx, msg, nil)
}
func (l *Logger) Debug(ctx context.Context, msg string) {
	l.Debugf(ctx, msg, nil)
}
func (l *Logger) Warn(ctx context.Context, msg string) {
	l.Warnf(ctx, msg, nil)
}
func (l *Logger) Error(ctx context.Context, err domain.Error) {
	l.Errorf(ctx, err, nil)
}
func (l *Logger) Infof(ctx context.Context, msg string, md map[string]interface{}) {
	l.print(ctx, domain.LogLevelInfo, msg, md, nil)
}
func (l *Logger) Debugf(ctx context.Context, msg string, md map[string]interface{}) {
	l.print(ctx, domain.LogLevelDebug, msg, md, nil)
}
func (l *Logger) Warnf(ctx context.Context, msg string, md map[string]interface{}) {
	l.print(ctx, domain.LogLevelWarning, msg, md, nil)
}
func (l *Logger) Errorf(ctx context.Context, err domain.Error, md map[string]interface{}) {
	msg := err.Error()

	stackTrace := map[string]string{}
	prevErr := err.Previous
	for prevErr != nil {
		stackTrace[string(prevErr.Kind)] = prevErr.Error()
		prevErr = prevErr.Previous
	}

	l.print(ctx, domain.LogLevelError, msg, md, stackTrace)
}

func (l *Logger) print(
	ctx context.Context,
	lvl string,
	msg string,
	metaData map[string]interface{},
	stackTrace map[string]string,
) {
	now := time.Now()
	traceID := getTraceID(ctx)

	log := log{
		Time:       &now,
		Level:      lvl,
		Msg:        msg,
		TraceID:    traceID,
		MetaData:   metaData,
		StackTrace: stackTrace,
	}

	var err error
	var j []byte
	if l.pretty {
		j, err = json.MarshalIndent(log, "", "\t")
	} else {
		j, err = json.Marshal(log)
	}

	if err != nil {
		jErrBase := `{"time": "%v", "level": "%s", "msg" "Error on marshal log: %s", "trace_id": "%s"}`
		jErr := fmt.Sprintf(jErrBase, now, domain.LogLevelError, err.Error(), traceID)
		j = []byte(jErr)
	}

	_, err = l.writer.Write(append(j, []byte("\n")...))
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func getTraceID(ctx context.Context) string {
	v := ctx.Value(domain.TraceIDCtxKey)
	if v == nil {
		return ""
	}
	return v.(string)
}
