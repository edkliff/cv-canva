package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func InitLog(level int) *logrus.Entry {
	l := logrus.New()
	l.Out = os.Stdout

	l.SetReportCaller(true)
	l.SetFormatter(newLogFormatter())

	switch level {
	case 6:
		l.SetLevel(logrus.PanicLevel)
	case 5:
		l.SetLevel(logrus.FatalLevel)
	case 4:
		l.SetLevel(logrus.ErrorLevel)
	case 3:
		l.SetLevel(logrus.WarnLevel)
	case 2:
		l.SetLevel(logrus.InfoLevel)
	case 1:
		l.SetLevel(logrus.DebugLevel)
	case 0:
		l.SetLevel(logrus.TraceLevel)
	default:
		l.SetLevel(logrus.InfoLevel)
	}

	return logrus.NewEntry(l)
}

func newLogFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	if len(arr) == 0 {
		return "unknown file"
	}
	return arr[len(arr)-1]
}
