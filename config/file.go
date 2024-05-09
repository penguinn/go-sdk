package config

import (
	"github.com/spf13/viper"
)

func FileInit(configFileType, configPath string) error {
	viper.SetConfigType(configFileType)
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
