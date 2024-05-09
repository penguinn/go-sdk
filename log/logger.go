package log

import (
	"io"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	logrus.Logger
	entryPool sync.Pool
}

var defaultLogger *Logger

func init() {
	defaultLogger = &Logger{
		Logger: logrus.Logger{
			Out:          os.Stdout,
			Formatter:    TextFormat,
			Hooks:        make(logrus.LevelHooks),
			Level:        logrus.InfoLevel,
			ExitFunc:     os.Exit,
			ReportCaller: true,
		},
	}
}

func NewLogger() *Logger {
	return &Logger{
		Logger: logrus.Logger{
			Out:          os.Stdout,
			Formatter:    TextFormat,
			Hooks:        make(logrus.LevelHooks),
			Level:        logrus.InfoLevel,
			ExitFunc:     os.Exit,
			ReportCaller: true,
		},
	}
}

func NewDefaultFieldsLogger(logger *Logger, fields map[string]interface{}) *logrus.Entry {
	if logger == nil {
		logger = defaultLogger
	}

	return logger.WithFields(fields)
}

// 设置无锁，增加并发
func SetNoLock() {
	defaultLogger.SetNoLock()
}

// 设置日志显示级别
func SetLevel(level logrus.Level) {
	defaultLogger.SetLevel(level)
}

// 设置日志格式
func SetFormatter(formatter logrus.Formatter) {
	defaultLogger.SetFormatter(formatter)
}

// 设置日志输出位置
func SetOutput(output io.Writer) {
	defaultLogger.SetOutput(output)
}

// 设置每条日志是否打印出日志输出地
func SetReportCaller(reportCaller bool) {
	defaultLogger.SetReportCaller(reportCaller)
}

// 增加日志回调，打印更多信息
func AddHook(hook logrus.Hook) {
	defaultLogger.AddHook(hook)
}

// 增加滚动日志，默认一个日志最大100m, 文件最多数目20，文件最多保留60天
func SetFileOutput(logFilePath string) {
	lumberjackLogrotate := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxBackups: 20,
		MaxAge:     60,
		Compress:   true,
	}
	logMultiWriter := io.MultiWriter(os.Stdout, lumberjackLogrotate)
	defaultLogger.SetOutput(logMultiWriter)
}
