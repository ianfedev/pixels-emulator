package config

import (
	"github.com/spf13/viper"
	"reflect"
)

// SetDefaults sets the default values of a configuration struct in a specific Viper instance using reflection.
func SetDefaults(v *viper.Viper, prefix string, config interface{}) {
	value := reflect.Indirect(reflect.ValueOf(config))
	typeOf := value.Type()

	if value.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := typeOf.Field(i)

		msTag := fieldType.Tag.Get("mapstructure")
		defaultTag := fieldType.Tag.Get("default")

		if msTag == "" {
			continue
		}
		key := msTag
		if prefix != "" {
			key = prefix + "." + msTag
		}

		if field.Kind() == reflect.Struct {
			SetDefaults(v, key, field.Interface())
			continue
		}

		if defaultTag != "" {
			v.SetDefault(key, defaultTag)
		}
	}
}
