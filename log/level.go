package log

import (
	"github.com/sirupsen/logrus"
)

func Tracef(format string, args ...interface{}) {
	defaultLogger.Logf(logrus.TraceLevel, format, args...)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Logf(logrus.DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Logf(logrus.InfoLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.Logf(logrus.WarnLevel, format, args...)
}

func Warningf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Logf(logrus.ErrorLevel, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	defaultLogger.Logf(logrus.FatalLevel, format, args...)
	defaultLogger.Exit(1)
}

func Panicf(format string, args ...interface{}) {
	defaultLogger.Logf(logrus.PanicLevel, format, args...)
}

func Trace(args ...interface{}) {
	defaultLogger.Log(logrus.TraceLevel, args...)
}

func Debug(args ...interface{}) {
	defaultLogger.Log(logrus.DebugLevel, args...)
}

func Info(args ...interface{}) {
	defaultLogger.Log(logrus.InfoLevel, args...)
}

func Warn(args ...interface{}) {
	defaultLogger.Log(logrus.WarnLevel, args...)
}

func Warning(args ...interface{}) {
	defaultLogger.Warn(args...)
}

func Error(args ...interface{}) {
	defaultLogger.Log(logrus.ErrorLevel, args...)
}

func Fatal(args ...interface{}) {
	defaultLogger.Log(logrus.FatalLevel, args...)
	defaultLogger.Exit(1)
}

func Panic(args ...interface{}) {
	defaultLogger.Log(logrus.PanicLevel, args...)
}

func TraceFn(fn logrus.LogFunction) {
	defaultLogger.LogFn(logrus.TraceLevel, fn)
}

func DebugFn(fn logrus.LogFunction) {
	defaultLogger.LogFn(logrus.DebugLevel, fn)
}

func InfoFn(fn logrus.LogFunction) {
	defaultLogger.LogFn(logrus.InfoLevel, fn)
}

func WarnFn(fn logrus.LogFunction) {
	defaultLogger.LogFn(logrus.WarnLevel, fn)
}

func WarningFn(fn logrus.LogFunction) {
	defaultLogger.WarnFn(fn)
}

func ErrorFn(fn logrus.LogFunction) {
	defaultLogger.LogFn(logrus.ErrorLevel, fn)
}

func FatalFn(fn logrus.LogFunction) {
	defaultLogger.LogFn(logrus.FatalLevel, fn)
	defaultLogger.Exit(1)
}

func PanicFn(fn logrus.LogFunction) {
	defaultLogger.LogFn(logrus.PanicLevel, fn)
}

func Traceln(args ...interface{}) {
	defaultLogger.Logln(logrus.TraceLevel, args...)
}

func Debugln(args ...interface{}) {
	defaultLogger.Logln(logrus.DebugLevel, args...)
}

func Infoln(args ...interface{}) {
	defaultLogger.Logln(logrus.InfoLevel, args...)
}

func Warnln(args ...interface{}) {
	defaultLogger.Logln(logrus.WarnLevel, args...)
}

func Warningln(args ...interface{}) {
	defaultLogger.Warnln(args...)
}

func Errorln(args ...interface{}) {
	defaultLogger.Logln(logrus.ErrorLevel, args...)
}

func Fatalln(args ...interface{}) {
	defaultLogger.Logln(logrus.FatalLevel, args...)
	defaultLogger.Exit(1)
}

func Panicln(args ...interface{}) {
	defaultLogger.Logln(logrus.PanicLevel, args...)
}
