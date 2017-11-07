package main

import (
	godiscovery "github.com/fananchong/go-discovery"
	"github.com/golang/glog"
)

type GLog struct {
}

func NewGLog() *GLog {
	return &GLog{}
}

func (this *GLog) Print(args ...interface{}) {
	glog.Error(args)
}

func (this *GLog) Printf(format string, args ...interface{}) {
	glog.Errorf(format, args)
}

func (this *GLog) Println(args ...interface{}) {
	glog.Errorln(args)
}

func (this *GLog) Info(args ...interface{}) {
	glog.Info(args)
}

func (this *GLog) Infof(format string, args ...interface{}) {
	glog.Infof(format, args)
}

func (this *GLog) Infoln(args ...interface{}) {
	glog.Infoln(args)
}

func (this *GLog) Warning(args ...interface{}) {
	glog.Warning(args)
}

func (this *GLog) Warningln(args ...interface{}) {
	glog.Warningln(args)
}

func (this *GLog) Warningf(format string, args ...interface{}) {
	glog.Warningf(format, args)
}

func (this *GLog) Error(args ...interface{}) {
	glog.Error(args)
}

func (this *GLog) Errorf(format string, args ...interface{}) {
	glog.Errorf(format, args)
}

func (this *GLog) Errorln(args ...interface{}) {
	glog.Errorln(args)
}

func (this *GLog) Fatal(args ...interface{}) {
	glog.Fatal(args)
}

func (this *GLog) Fatalln(args ...interface{}) {
	glog.Fatalln(args)
}

func (this *GLog) Fatalf(format string, args ...interface{}) {
	glog.Fatalf(format, args)
}

func (this *GLog) Flush() {
	glog.Flush()
}

var (
	xlog godiscovery.ILogger = NewGLog()
)

func SetLogger(log godiscovery.ILogger) {
	xlog = log
}
