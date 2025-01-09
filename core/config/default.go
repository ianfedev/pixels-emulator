package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"reflect"
)

// SetDefaults sets the default values of a configuration struct in a specific Viper instance using reflection.
// It recursively sets defaults for each field in the struct, based on the "default" tag in the struct fields.
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

// CreateDefaultConfig creates a default config file if it does not exist.
// It checks if the config file exists at the given path, and if not, creates it with default values.
// It logs the creation of the file using the provided logger.
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
