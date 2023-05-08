package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	l := New()
	assert.NotNil(t, l)
}

func Test_logger_With(t *testing.T) {
	l, observer := NewTest()
	l = l.With("requestID", "123")
	l.Info("Logging with request ID field appended")
	logs := observer.All()
	entry := logs[0]
	fields := entry.Context
	assert.Equal(t, "requestID", fields[0].Key)
	assert.Equal(t, "123", fields[0].String)
}

func Test_logger_WithoutCaller(t *testing.T) {
	l, observer := NewTest()
	l = l.WithoutCaller()
	l.Info("Logging without caller")
	logs := observer.All()
	entry := logs[0]
	assert.Equal(t, false, entry.Caller.Defined)
}
