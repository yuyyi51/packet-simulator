package log

import "go.uber.org/zap"

var originLogger *zap.Logger
var defaultLogger *zap.SugaredLogger

func init() {
	originLogger, _ = zap.NewDevelopment()
	defaultLogger = originLogger.Sugar()
}

func Debugf(template string, args ...interface{}) {
	defaultLogger.Debugf(template, args...)
	_ = originLogger.Sync()
}

func Infof(template string, args ...interface{}) {
	defaultLogger.Infof(template, args...)
	_ = originLogger.Sync()
}

func Warnf(template string, args ...interface{}) {
	defaultLogger.Warnf(template, args...)
	_ = originLogger.Sync()
}

func Errorf(template string, args ...interface{}) {
	defaultLogger.Errorf(template, args...)
	_ = originLogger.Sync()
}

func Panicf(template string, args ...interface{}) {
	defaultLogger.Panicf(template, args...)
	_ = originLogger.Sync()
}

func Fatalf(template string, args ...interface{}) {
	defaultLogger.Fatalf(template, args...)
	_ = originLogger.Sync()
}
