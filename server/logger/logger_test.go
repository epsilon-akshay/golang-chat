package logger

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestTextTypeLogger(t *testing.T) {
	lg := New(os.Stdout, "text", "info")
	assert.Equal(t, &logrus.TextFormatter{}, lg.Formatter)
}

func TestJsonTypeLogger(t *testing.T) {
	lg := New(os.Stdout, "json", "info")
	assert.Equal(t, &logrus.JSONFormatter{}, lg.Formatter)
}

func TestInfoLevelLogger(t *testing.T) {
	lg := New(os.Stdout, "text", "info")
	assert.Equal(t, logrus.InfoLevel, lg.Level)
}

func TestWarnLevelLogger(t *testing.T) {
	lg := New(os.Stdout, "text", "warn")
	assert.Equal(t, logrus.WarnLevel, lg.Level)
}

func TestDebugLevelLogger(t *testing.T) {
	lg := New(os.Stdout, "text", "debug")
	assert.Equal(t, logrus.DebugLevel, lg.Level)
}

func TestErrorLevelLogger(t *testing.T) {
	lg := New(os.Stdout, "text", "error")
	assert.Equal(t, logrus.ErrorLevel, lg.Level)
}

func TestFatalLevelLogger(t *testing.T) {
	lg := New(os.Stdout, "text", "fatal")
	assert.Equal(t, logrus.FatalLevel, lg.Level)
}

func TestPanicLevelLogger(t *testing.T) {
	lg := New(os.Stdout, "text", "panic")
	assert.Equal(t, logrus.PanicLevel, lg.Level)
}

func TestTraceLevelLogger(t *testing.T) {
	lg := New(os.Stdout, "text", "trace")
	assert.Equal(t, logrus.TraceLevel, lg.Level)
}

func TestParseError(t *testing.T) {
	lg := New(os.Stdout, "text", "asds")
	assert.Equal(t, logrus.InfoLevel, lg.Level)
}
