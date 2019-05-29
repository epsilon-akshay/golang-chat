package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func New(writer io.Writer, format string, lvl string) {
	logger := logrus.New()

	logger.SetOutput(writer)

	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(level)
	}

	if format == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{})
	}

	Log = logger
}
