package config

// ServerConfig defines the essential info for the emulator
// environment and linking with Nitro Client.
type ServerConfig struct {
	IP          string `mapstructure:"ip" default:"127.0.0.1"`            // IP address where Nitro will consume
	Port        uint16 `mapstructure:"port" default:"3000"`               // Port where Nitro will consume
	Environment string `mapstructure:"environment" default:"DEVELOPMENT"` // Environment of deployment
	PingRate    uint16 `mapstructure:"ping_rate" default:"5"`             // PingRate every time a ping packet should be sent. If 0 client must ping.
}

// DatabaseConfig defines the information to establish connection
// with the database for content persistence.
type DatabaseConfig struct {
	Host     string `mapstructure:"host" default:"127.0.0.1" security:"PRODUCTION"`    // Host of the SQL connection.
	Port     uint16 `mapstructure:"port" default:"3306" security:"PRODUCTION"`         // Port of the SQL connection.
	Database string `mapstructure:"database" default:"pixels" security:"PRODUCTION"`   // Database name of the SQL connection.
	User     string `mapstructure:"user" default:"pixels" security:"PRODUCTION"`       // User of the SQL connection.
	Password string `mapstructure:"password" default:"password" security:"PRODUCTION"` // Password of the SQL connection.
}

// LoggingConfig holds the configuration for the logging system.
type LoggingConfig struct {
	ConsoleColor bool   `mapstructure:"console_color" default:"true"` // ConsoleColor if console should be colored or not.
	JSON         bool   `mapstructure:"json" default:"false"`         // JSON if file logging should be in json for monitoring.
	Level        string `mapstructure:"level" default:"INFO"`         // Level of the logging.
}

// Config defines the complete model of configuration to
// be unmarshalled by a configuration provider.
type Config struct {
	Server   ServerConfig   `mapstructure:"server" default:""`   // Server base configuration.
	Database DatabaseConfig `mapstructure:"database" default:""` // Database connection configuration.
	Logging  LoggingConfig  `mapstructure:"logging" default:""`  // Logging configuration.
}
