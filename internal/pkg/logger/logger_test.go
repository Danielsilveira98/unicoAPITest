package logger

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
	"github.com/google/go-cmp/cmp"
)

type stubWriter struct {
	writeInp []byte
	write    func(p []byte) (n int, err error)
}

func (s *stubWriter) Write(p []byte) (n int, err error) {
	s.writeInp = p
	return s.write(p)
}

const (
	traceID = "tracing"
	msg     = "Log message"
)

func TestLogger_print(t *testing.T) {
	pretty := []bool{false, true}
	ctx := context.Background()
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	lvl := "test"
	metaData := map[string]interface{}{
		"key": "value",
	}

	writerMock := &stubWriter{
		write: func(p []byte) (n int, err error) {
			return 1, nil
		},
	}

	for _, prt := range pretty {
		lgg := NewLogger(writerMock, prt)

		lgg.print(ctx, lvl, msg, metaData)

		var gotL log

		err := json.Unmarshal(writerMock.writeInp, &gotL)
		if err != nil {
			t.Fatal(err)
		}
		if err != nil {
			t.Fatal(err)
		}

		if gotL.Time == nil {
			t.Error("expect error with time not nil")
		}

		if lvl != gotL.Level {
			t.Errorf("expect error with level as %v, got as %v", lvl, gotL.Level)
		}

		if msg != gotL.Msg {
			t.Errorf("expect error with msg as %v, got as %v", msg, gotL.Msg)
		}

		if traceID != gotL.TraceID {
			t.Errorf("expect error with traceID as %v, got as %v", domain.TraceIDCtxKey, gotL.TraceID)
		}

		if diff := cmp.Diff(metaData, gotL.MetaData); diff != "" {
			t.Errorf("unexpected metaData (-want +got):\n%s", diff)
		}
	}
}

func TestLogget_getTraceID(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	id := getTraceID(ctx)

	if id != traceID {
		t.Errorf("expect id %s, got %s", traceID, id)
	}
}

func TestLogger_Infof(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	msg := "msg Infof"
	metaData := map[string]interface{}{
		"key": "value",
	}

	writerMock := &stubWriter{
		write: func(p []byte) (n int, err error) {
			return 1, nil
		},
	}

	lgg := NewLogger(writerMock, false)

	lgg.Infof(ctx, msg, metaData)

	var gotL log

	err := json.Unmarshal(writerMock.writeInp, &gotL)
	if err != nil {
		t.Fatal(err)
	}

	if gotL.Time == nil {
		t.Error("expect error with time not nil")
	}

	if domain.LogLevelInfo != gotL.Level {
		t.Errorf("expect error with level as %v, got as %v", domain.LogLevelInfo, gotL.Level)
	}

	if msg != gotL.Msg {
		t.Errorf("expect error with msg as %v, got as %v", msg, gotL.Msg)
	}

	if traceID != gotL.TraceID {
		t.Errorf("expect error with traceID as %v, got as %v", domain.TraceIDCtxKey, gotL.TraceID)
	}

	if diff := cmp.Diff(metaData, gotL.MetaData); diff != "" {
		t.Errorf("unexpected metaData (-want +got):\n%s", diff)
	}
}

func TestLogger_Debugf(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	msg := "msg Debugf"
	metaData := map[string]interface{}{
		"key": "value",
	}

	writerMock := &stubWriter{
		write: func(p []byte) (n int, err error) {
			return 1, nil
		},
	}

	lgg := NewLogger(writerMock, false)

	lgg.Debugf(ctx, msg, metaData)

	var gotL log

	err := json.Unmarshal(writerMock.writeInp, &gotL)
	if err != nil {
		t.Fatal(err)
	}

	if gotL.Time == nil {
		t.Error("expect error with time not nil")
	}

	if domain.LogLevelDebug != gotL.Level {
		t.Errorf("expect error with level as %v, got as %v", domain.LogLevelDebug, gotL.Level)
	}

	if msg != gotL.Msg {
		t.Errorf("expect error with msg as %v, got as %v", msg, gotL.Msg)
	}

	if traceID != gotL.TraceID {
		t.Errorf("expect error with traceID as %v, got as %v", domain.TraceIDCtxKey, gotL.TraceID)
	}

	if diff := cmp.Diff(metaData, gotL.MetaData); diff != "" {
		t.Errorf("unexpected metaData (-want +got):\n%s", diff)
	}
}

func TestLogger_Warnf(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	msg := "msg Warnf"
	metaData := map[string]interface{}{
		"key": "value",
	}

	writerMock := &stubWriter{
		write: func(p []byte) (n int, err error) {
			return 1, nil
		},
	}

	lgg := NewLogger(writerMock, false)

	lgg.Warnf(ctx, msg, metaData)

	var gotL log

	err := json.Unmarshal(writerMock.writeInp, &gotL)
	if err != nil {
		t.Fatal(err)
	}

	if gotL.Time == nil {
		t.Error("expect error with time not nil")
	}

	if domain.LogLevelWarning != gotL.Level {
		t.Errorf("expect error with level as %v, got as %v", domain.LogLevelWarning, gotL.Level)
	}

	if msg != gotL.Msg {
		t.Errorf("expect error with msg as %v, got as %v", msg, gotL.Msg)
	}

	if traceID != gotL.TraceID {
		t.Errorf("expect error with traceID as %v, got as %v", domain.TraceIDCtxKey, gotL.TraceID)
	}

	if diff := cmp.Diff(metaData, gotL.MetaData); diff != "" {
		t.Errorf("unexpected metaData (-want +got):\n%s", diff)
	}
}

func TestLogger_Errorf(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	msg := "msg Errorf"
	metaData := map[string]interface{}{
		"key": "value",
	}

	writerMock := &stubWriter{
		write: func(p []byte) (n int, err error) {
			return 1, nil
		},
	}

	lgg := NewLogger(writerMock, false)

	lgg.Errorf(ctx, msg, metaData)

	var gotL log

	err := json.Unmarshal(writerMock.writeInp, &gotL)
	if err != nil {
		t.Fatal(err)
	}

	if gotL.Time == nil {
		t.Error("expect error with time not nil")
	}

	if domain.LogLevelError != gotL.Level {
		t.Errorf("expect error with level as %v, got as %v", domain.LogLevelError, gotL.Level)
	}

	if msg != gotL.Msg {
		t.Errorf("expect error with msg as %v, got as %v", msg, gotL.Msg)
	}

	if traceID != gotL.TraceID {
		t.Errorf("expect error with traceID as %v, got as %v", domain.TraceIDCtxKey, gotL.TraceID)
	}

	if diff := cmp.Diff(metaData, gotL.MetaData); diff != "" {
		t.Errorf("unexpected metaData (-want +got):\n%s", diff)
	}
}

func TestLogger_Info(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	msg := "msg Info"

	writerMock := &stubWriter{
		write: func(p []byte) (n int, err error) {
			return 1, nil
		},
	}

	lgg := NewLogger(writerMock, false)

	lgg.Info(ctx, msg)

	var gotL log

	err := json.Unmarshal(writerMock.writeInp, &gotL)
	if err != nil {
		t.Fatal(err)
	}

	if gotL.Time == nil {
		t.Error("expect error with time not nil")
	}

	if domain.LogLevelInfo != gotL.Level {
		t.Errorf("expect error with level as %v, got as %v", domain.LogLevelInfo, gotL.Level)
	}

	if msg != gotL.Msg {
		t.Errorf("expect error with msg as %v, got as %v", msg, gotL.Msg)
	}

	if traceID != gotL.TraceID {
		t.Errorf("expect error with traceID as %v, got as %v", domain.TraceIDCtxKey, gotL.TraceID)
	}
}

func TestLogger_Debug(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	msg := "msg Debug"

	writerMock := &stubWriter{
		write: func(p []byte) (n int, err error) {
			return 1, nil
		},
	}

	lgg := NewLogger(writerMock, false)

	lgg.Debug(ctx, msg)

	var gotL log

	err := json.Unmarshal(writerMock.writeInp, &gotL)
	if err != nil {
		t.Fatal(err)
	}

	if gotL.Time == nil {
		t.Error("expect error with time not nil")
	}

	if domain.LogLevelDebug != gotL.Level {
		t.Errorf("expect error with level as %v, got as %v", domain.LogLevelDebug, gotL.Level)
	}

	if msg != gotL.Msg {
		t.Errorf("expect error with msg as %v, got as %v", msg, gotL.Msg)
	}

	if traceID != gotL.TraceID {
		t.Errorf("expect error with traceID as %v, got as %v", domain.TraceIDCtxKey, gotL.TraceID)
	}
}

func TestLogger_Warn(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	msg := "msg Warn"

	writerMock := &stubWriter{
		write: func(p []byte) (n int, err error) {
			return 1, nil
		},
	}

	lgg := NewLogger(writerMock, false)

	lgg.Warn(ctx, msg)

	var gotL log

	err := json.Unmarshal(writerMock.writeInp, &gotL)
	if err != nil {
		t.Fatal(err)
	}

	if gotL.Time == nil {
		t.Error("expect error with time not nil")
	}

	if domain.LogLevelWarning != gotL.Level {
		t.Errorf("expect error with level as %v, got as %v", domain.LogLevelWarning, gotL.Level)
	}

	if msg != gotL.Msg {
		t.Errorf("expect error with msg as %v, got as %v", msg, gotL.Msg)
	}

	if traceID != gotL.TraceID {
		t.Errorf("expect error with traceID as %v, got as %v", domain.TraceIDCtxKey, gotL.TraceID)
	}
}

func TestLogger_Error(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	msg := "msg Error"

	writerMock := &stubWriter{
		write: func(p []byte) (n int, err error) {
			return 1, nil
		},
	}

	lgg := NewLogger(writerMock, false)

	lgg.Error(ctx, msg)

	var gotL log

	err := json.Unmarshal(writerMock.writeInp, &gotL)
	if err != nil {
		t.Fatal(err)
	}

	if gotL.Time == nil {
		t.Error("expect error with time not nil")
	}

	if domain.LogLevelError != gotL.Level {
		t.Errorf("expect error with level as %v, got as %v", domain.LogLevelError, gotL.Level)
	}

	if msg != gotL.Msg {
		t.Errorf("expect error with msg as %v, got as %v", msg, gotL.Msg)
	}

	if traceID != gotL.TraceID {
		t.Errorf("expect error with traceID as %v, got as %v", domain.TraceIDCtxKey, gotL.TraceID)
	}
}
