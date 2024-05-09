package config

import (
	"strings"

	"github.com/spf13/viper"
)

func EnvInit() error {
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	return nil
}
