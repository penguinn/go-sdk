package main

import (
	"github.com/penguinn/go-sdk/log"
)

func main() {
	// 简单测试样例
	log.Info("program start")
	log.With(map[string]interface{}{"example": true}).Info("program end")
	log.NewLogger().WithFields(map[string]interface{}{"project": "easydata"})

	// 设置有默认标签的logger
	logger := log.NewDefaultFieldsLogger(nil, map[string]interface{}{"a": true, "b": true})
	logger.Error("hello")
	logger.Warn("world")

	// 设置文件和stdout 双输出
	log.SetFileOutput("./log/stdout.log")
	log.Info("test file and stdout")
}
