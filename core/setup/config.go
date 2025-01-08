package setup

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"pixels-emulator/core/config"
	"reflect"
	"strings"
)

// Config Creates and unmarshalls the configuration from viper,
// performing security checks like non-sensitive data via file.
func Config(path string, logger *zap.Logger) (*config.Config, error) {

	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("ini")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var tempCfg config.Config
	if err := v.Unmarshal(&tempCfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}
	CheckSecurityAlerts(&tempCfg, logger)

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.SetDefaults(v, "", config.Config{})

	var cfg config.Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &cfg, nil

}

// CheckSecurityAlerts recurses through the configuration structure to detect fields marked with 'security_alert' tags
// in the specified environment. If such a field is found, it logs a security alert using the zap logger.
func CheckSecurityAlerts(c *config.Config, logger *zap.Logger) {
	checkStruct(reflect.ValueOf(c), c.Server.Environment, logger)
}

func checkStruct(v reflect.Value, env string, logger *zap.Logger) {

	// Handle pointers by getting the actual value they point to
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Only proceed if the value is a struct
	if v.Kind() != reflect.Struct {
		return
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		securityAlertTag := fieldType.Tag.Get("security")
		if securityAlertTag != "" && securityAlertTag == env {
			// If the field is set (non-zero value), log a security alert
			if !field.IsZero() {
				logger.Info("Security Alert: Sensitive data detected in configuration",
					zap.String("field", fieldType.Name),
					zap.String("environment", env))
			}
		}

		// Recursively check nested structs
		if field.Kind() == reflect.Struct || field.Kind() == reflect.Ptr {
			checkStruct(field, env, logger)
		}
	}

}
