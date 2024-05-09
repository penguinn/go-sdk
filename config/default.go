package config

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

func RecurSetDefault(key string, t reflect.Type, v reflect.Value) {
	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i)
		fileValue := v.Field(i)

		if fieldType.Type.Kind() != reflect.Struct {
			viper.SetDefault(fmt.Sprintf("%s.%s", key, fieldType.Name), fileValue.Interface())
		} else if fieldType.Type.Kind() == reflect.Struct && key == "" {
			RecurSetDefault(fieldType.Name, fieldType.Type, fileValue)
		} else {
			RecurSetDefault(fmt.Sprintf("%s.%s", key, fieldType.Name), fieldType.Type, fileValue)
		}
	}
}
