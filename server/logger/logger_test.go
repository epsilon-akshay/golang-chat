package logger

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestTextTypeLogger(t *testing.T) {
	New(os.Stdout, "text", "info")
	assert.Equal(t, &logrus.TextFormatter{}, log.Formatter)
}

func TestJsonTypeLogger(t *testing.T) {
	New(os.Stdout, "json", "info")
	assert.Equal(t, &logrus.JSONFormatter{}, log.Formatter)
}

func TestInfoLevelLogger(t *testing.T) {
	New(os.Stdout, "text", "info")
	assert.Equal(t, logrus.InfoLevel, log.Level)
}

func TestWarnLevelLogger(t *testing.T) {
	New(os.Stdout, "text", "warn")
	assert.Equal(t, logrus.WarnLevel, log.Level)
}

func TestDebugLevelLogger(t *testing.T) {
	New(os.Stdout, "text", "debug")
	assert.Equal(t, logrus.DebugLevel, log.Level)
}

func TestErrorLevelLogger(t *testing.T) {
	New(os.Stdout, "text", "error")
	assert.Equal(t, logrus.ErrorLevel, log.Level)
}

func TestFatalLevelLogger(t *testing.T) {
	New(os.Stdout, "text", "fatal")
	assert.Equal(t, logrus.FatalLevel, log.Level)
}

func TestPanicLevelLogger(t *testing.T) {
	New(os.Stdout, "text", "panic")
	assert.Equal(t, logrus.PanicLevel, log.Level)
}

func TestTraceLevelLogger(t *testing.T) {
	New(os.Stdout, "text", "trace")
	assert.Equal(t, logrus.TraceLevel, log.Level)
}

func TestParseError(t *testing.T) {
	New(os.Stdout, "text", "asds")
	assert.Equal(t, logrus.InfoLevel, log.Level)
}
