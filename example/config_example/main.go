package main

import (
	"flag"
	"fmt"

	"github.com/penguinn/go-sdk/config"
	"github.com/penguinn/go-sdk/log"
	"github.com/spf13/viper"
)

type Config struct {
	Server      ServerConfig      `mapstructure:"Server"`
	UserManager UserManagerConfig `mapstructure:"UserManager"`
}

type ServerConfig struct {
	Port string `mapstructure:"Port"` // 启动端口
}

type UserManagerConfig struct {
	Method string `mapstructure:"Method"`
	URL    string `mapstructure:"URL"`
}

// GlobalConfig 实例化结构体，同时设置默认值
var GlobalConfig = Config{
	Server: ServerConfig{
		Port: ":8080",
	},
	UserManager: UserManagerConfig{
		Method: "GET",
	},
}

var configPath string
var configFileType string

func main() {
	flag.StringVar(&configFileType, "t", "ini", "config file type")
	flag.StringVar(&configPath, "f", "./config.ini", "config file path")
	flag.Parse()

	err := config.Init(configFileType, configPath, &GlobalConfig)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(viper.AllSettings())
}
