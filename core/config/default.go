package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"reflect"
)

// SetDefaults recursively sets default values for the fields in the provided config structure.
// It uses the mapstructure and default tags to determine the key-value pairs for default values.
// The function also supports nested structs by calling itself recursively for struct fields.
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

// CreateDefaultConfig creates a default configuration file if it doesn't exist.
// It initializes the Viper configuration, sets defaults using the SetDefaults function,
// and writes the configuration to a file if it doesn't already exist.
func CreateDefaultConfig(path string, logger *zap.Logger) error {

	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("ini")

	cfg := Config{}
	SetDefaults(v, "", cfg)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := v.WriteConfigAs(path); err != nil {
			return fmt.Errorf("error creating default config file: %w", err)
		}
		logger.Info("Config file not found. Created default config file.")
	}

	return nil
}
