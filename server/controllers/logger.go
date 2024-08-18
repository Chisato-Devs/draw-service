package controllers

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var logger = Logger()

type myFormatter struct {
	logrus.TextFormatter
}

func (f *myFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int

	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = 31 // gray
	case logrus.WarnLevel:
		levelColor = 33 // yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 34 // blue
	}
	return []byte(fmt.Sprintf("[%s] - \x1b[%dm%s\x1b[0m - %s\n", entry.Time.Format(f.TimestampFormat), levelColor, strings.ToUpper(entry.Level.String()), entry.Message)), nil
}

func Logger() *logrus.Logger {
	f, _ := os.OpenFile("logrus.txt", os.O_CREATE|os.O_WRONLY, 0777)
	logger := &logrus.Logger{
		Out:   io.MultiWriter(os.Stderr, f),
		Level: logrus.InfoLevel,
		Formatter: &myFormatter{
			logrus.TextFormatter{
				FullTimestamp:          true,
				TimestampFormat:        "2006-01-02 15:04:05",
				ForceColors:            true,
				DisableLevelTruncation: true,
			},
		},
	}
	return logger
}
