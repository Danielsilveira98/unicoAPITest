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

func TestLogger_print(t *testing.T) {
	pretty := []bool{false, true}
	ctx := context.Background()
	traceID := "myValue"
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	lvl := "test"
	msg := "Some error"
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

		json.Unmarshal(writerMock.writeInp, &gotL)

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
			t.Errorf("expect error with traceID as %v, got as %v", traceID, gotL.TraceID)
		}

		if diff := cmp.Diff(metaData, gotL.MetaData); diff != "" {
			t.Errorf("unexpected metaData (-want +got):\n%s", diff)
		}
	}
}

func TestLogget_getTraceID(t *testing.T) {
	ctx := context.Background()
	value := "myValue"
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, value)

	id := getTraceID(ctx)

	if id != value {
		t.Errorf("expect id %s, got %s", value, id)
	}
}

func TestLogger_Infof(t *testing.T) {
	ctx := context.Background()
	traceID := "myValue"
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	msg := "Some error"
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

	json.Unmarshal(writerMock.writeInp, &gotL)

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
		t.Errorf("expect error with traceID as %v, got as %v", traceID, gotL.TraceID)
	}

	if diff := cmp.Diff(metaData, gotL.MetaData); diff != "" {
		t.Errorf("unexpected metaData (-want +got):\n%s", diff)
	}
}

func TestLogger_Debugf(t *testing.T) {
	ctx := context.Background()
	traceID := "myValue"
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	msg := "Some error"
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

	json.Unmarshal(writerMock.writeInp, &gotL)

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
		t.Errorf("expect error with traceID as %v, got as %v", traceID, gotL.TraceID)
	}

	if diff := cmp.Diff(metaData, gotL.MetaData); diff != "" {
		t.Errorf("unexpected metaData (-want +got):\n%s", diff)
	}
}

func TestLogger_Warnf(t *testing.T) {
	ctx := context.Background()
	traceID := "myValue"
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	msg := "Some error"
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

	json.Unmarshal(writerMock.writeInp, &gotL)

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
		t.Errorf("expect error with traceID as %v, got as %v", traceID, gotL.TraceID)
	}

	if diff := cmp.Diff(metaData, gotL.MetaData); diff != "" {
		t.Errorf("unexpected metaData (-want +got):\n%s", diff)
	}
}

func TestLogger_Errorf(t *testing.T) {
	ctx := context.Background()
	traceID := "myValue"
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	msg := "Some error"
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

	json.Unmarshal(writerMock.writeInp, &gotL)

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
		t.Errorf("expect error with traceID as %v, got as %v", traceID, gotL.TraceID)
	}

	if diff := cmp.Diff(metaData, gotL.MetaData); diff != "" {
		t.Errorf("unexpected metaData (-want +got):\n%s", diff)
	}
}

func TestLogger_Info(t *testing.T) {
	ctx := context.Background()
	traceID := "myValue"
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	msg := "Some error"

	writerMock := &stubWriter{
		write: func(p []byte) (n int, err error) {
			return 1, nil
		},
	}

	lgg := NewLogger(writerMock, false)

	lgg.Info(ctx, msg)

	var gotL log

	json.Unmarshal(writerMock.writeInp, &gotL)

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
		t.Errorf("expect error with traceID as %v, got as %v", traceID, gotL.TraceID)
	}

}

func TestLogger_Debug(t *testing.T) {
	ctx := context.Background()
	traceID := "myValue"
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	msg := "Some error"

	writerMock := &stubWriter{
		write: func(p []byte) (n int, err error) {
			return 1, nil
		},
	}

	lgg := NewLogger(writerMock, false)

	lgg.Debug(ctx, msg)

	var gotL log

	json.Unmarshal(writerMock.writeInp, &gotL)

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
		t.Errorf("expect error with traceID as %v, got as %v", traceID, gotL.TraceID)
	}

}

func TestLogger_Warn(t *testing.T) {
	ctx := context.Background()
	traceID := "myValue"
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	msg := "Some error"

	writerMock := &stubWriter{
		write: func(p []byte) (n int, err error) {
			return 1, nil
		},
	}

	lgg := NewLogger(writerMock, false)

	lgg.Warn(ctx, msg)

	var gotL log

	json.Unmarshal(writerMock.writeInp, &gotL)

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
		t.Errorf("expect error with traceID as %v, got as %v", traceID, gotL.TraceID)
	}

}

func TestLogger_Error(t *testing.T) {
	ctx := context.Background()
	traceID := "myValue"
	ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceID)

	msg := "Some error"

	writerMock := &stubWriter{
		write: func(p []byte) (n int, err error) {
			return 1, nil
		},
	}

	lgg := NewLogger(writerMock, false)

	lgg.Error(ctx, msg)

	var gotL log

	json.Unmarshal(writerMock.writeInp, &gotL)

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
		t.Errorf("expect error with traceID as %v, got as %v", traceID, gotL.TraceID)
	}

}
