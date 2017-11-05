package godiscovery

import (
	"github.com/golang/glog"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (this *Logger) Fatal(args ...interface{}) {
	glog.Fatal(args)
}

func (this *Logger) Fatalf(format string, args ...interface{}) {
	glog.Fatalf(format, args)
}

func (this *Logger) Fatalln(args ...interface{}) {
	glog.Fatalln(args)
}

func (this *Logger) Print(args ...interface{}) {
	glog.Info(args)
}

func (this *Logger) Printf(format string, args ...interface{}) {
	glog.Infof(format, args)
}

func (this *Logger) Println(args ...interface{}) {
	glog.Infoln(args)
}
