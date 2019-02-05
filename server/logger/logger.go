package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

func New(writer io.Writer, format string, lvl string) *logrus.Logger {
	log := logrus.New()

	log.SetOutput(writer)

	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		log.Print("set level to info level")
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.SetLevel(level)
	}

	if format == "json" {
		log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		log.SetFormatter(&logrus.TextFormatter{})
	}

	return log
}
