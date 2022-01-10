package logger

import (
	"github.com/sirupsen/logrus"
	// nested "github.com/antonfisher/nested-logrus-formatter"
	"io"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Logger interface {
	SetReportCaller(bool)
	Errorln(...interface{})
	Warningln(...interface{})
	Infoln(...interface{})
	Debugln(...interface{})
	Fatalln(...interface{})
	Println(...interface{})

	Errorf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Writer() *io.PipeWriter
	SetOutput(writer io.Writer)
}
type logger struct {
	*logrus.Logger
}

//func formatted() *logrus.Entry {
//	pc, file, line, ok := runtime.Caller(1)
//	if !ok {
//		panic("Could not get context info for logger!")
//	}
//
//	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
//	funcname := runtime.FuncForPC(pc).Name()
//	fn := funcname[strings.LastIndex(funcname, ".")+1:]
//	return logrus.WithField("file", filename).WithField("function", fn)
//}

func NewLogger() Logger {
	log := logrus.New()
	log.SetReportCaller(true)
	log.SetFormatter(customFormatter())
	log.SetReportCaller(true)
	return &logger{log}
}

func customFormatter() *logrus.TextFormatter {
	Formatter := new(logrus.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	Formatter.CallerPrettyfier = func(f *runtime.Frame) (string, string) {
		s := strings.Split(f.Function, ".")
		funcname := s[len(s)-1]
		_, filename := path.Split(f.File)
		return funcname, " " + filename + ":" + strconv.Itoa(f.Line)
	}
	return Formatter
}
func (l *logger) format() string {
	return time.Now().Format("2006/01/02 15:04:05")
}
