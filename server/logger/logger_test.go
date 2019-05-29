package logger

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestTextTypeLogger(t *testing.T) {
	New(os.Stdout, "text", "info")
	assert.Equal(t, &logrus.TextFormatter{}, Log.Formatter)
}

func TestJsonTypeLogger(t *testing.T) {
	New(os.Stdout, "json", "info")
	assert.Equal(t, &logrus.JSONFormatter{}, Log.Formatter)
}

func TestInfoLevelLogger(t *testing.T) {
	New(os.Stdout, "text", "info")
	assert.Equal(t, logrus.InfoLevel, Log.Level)
}

func TestWarnLevelLogger(t *testing.T) {
	New(os.Stdout, "text", "warn")
	assert.Equal(t, logrus.WarnLevel, Log.Level)
}

func TestDebugLevelLogger(t *testing.T) {
	New(os.Stdout, "text", "debug")
	assert.Equal(t, logrus.DebugLevel, Log.Level)
}

func TestErrorLevelLogger(t *testing.T) {
	New(os.Stdout, "text", "error")
	assert.Equal(t, logrus.ErrorLevel, Log.Level)
}

func TestFatalLevelLogger(t *testing.T) {
	New(os.Stdout, "text", "fatal")
	assert.Equal(t, logrus.FatalLevel, Log.Level)
}

func TestPanicLevelLogger(t *testing.T) {
	New(os.Stdout, "text", "panic")
	assert.Equal(t, logrus.PanicLevel, Log.Level)
}

func TestTraceLevelLogger(t *testing.T) {
	New(os.Stdout, "text", "trace")
	assert.Equal(t, logrus.TraceLevel, Log.Level)
}

func TestParseError(t *testing.T) {
	New(os.Stdout, "text", "asds")
	assert.Equal(t, logrus.InfoLevel, Log.Level)
}
