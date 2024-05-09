package config

import (
	"errors"
	"reflect"

	"github.com/spf13/viper"
)

func Init(configFileType, configPath string, s interface{}) error {
	// 设置默认值
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() == reflect.Ptr {
		RecurSetDefault("", t.Elem(), v.Elem())
	} else if t.Kind() == reflect.Struct {
		RecurSetDefault("", t, v)
	} else {
		return errors.New("invalid param structure")
	}

	// 初始化文件配置
	if configPath != "" {
		err := FileInit(configFileType, configPath)
		if err != nil {
			return err
		}
	}

	// 初始化环境配置
	err := EnvInit()
	if err != nil {
		return err
	}

	// 读取所有配置后，解析到最终结构体中
	err = viper.Unmarshal(s)
	if err != nil {
		return err
	}

	return nil
}
