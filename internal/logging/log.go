package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)

var log = logrus.New()

func init() {
	log.SetReportCaller(true)
	log.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Println(args ...interface{}) {
	log.Println(args...)
}
func Error(args ...interface{}) {
	log.Error(args...)
}
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}
func Warn(args ...interface{}) {
	log.Warn(args...)
}
