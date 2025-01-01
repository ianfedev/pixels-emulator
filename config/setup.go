package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"strings"
)

// CreateConfig Creates and unmarshalls the configuration from viper,
// performing security checks like non-sensitive data via file.
func CreateConfig(path string, logger *zap.Logger) (*Config, error) {

	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("ini")

	/*
		if err := v.ReadInConfig(); err != nil {
				return nil, fmt.Errorf("error reading config file: %w", err)
			}
	*/

	// TODO: Read for security alerts.

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	SetDefaults(v, "", Config{})

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &cfg, nil

}
