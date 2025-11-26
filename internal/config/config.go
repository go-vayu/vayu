// Package config provides application configuration.
package config

import (
	"log/slog"
	"strings"

	"github.com/spf13/viper"
)

// Key represents a configuration key.
type Key string

// Configuration keys for the application.
const (
	// #nosec
	OpenAQAPIKey Key = `openaq.apikey`

	DatabaseHost                  Key = `database.host`
	DatabasePort                  Key = `database.port`
	DatabaseUser                  Key = `database.user`
	DatabasePassword              Key = `database.password`
	DatabaseDatabase              Key = `database.database`
	DatabaseMaxOpenConnections    Key = `database.maxopenconnections`
	DatabaseMaxIdleConnections    Key = `database.maxidleconnections`
	DatabaseMaxConnectionLifetime Key = `database.maxconnectionlifetime`
	DatabaseSslMode               Key = `database.sslmode`
)

// Get returns the configuration value.
func (k Key) Get() any {
	return viper.Get(string(k))
}

// GetString returns the configuration value as a string.
func (k Key) GetString() string {
	return viper.GetString(string(k))
}

// GetInt returns the configuration value as a int.
func (k Key) GetInt() int {
	return viper.GetInt(string(k))
}

func (k Key) setDefault(value any) {
	viper.SetDefault(string(k), value)
}

func initDefaultConfig() {
	OpenAQAPIKey.setDefault("")

	DatabaseHost.setDefault("localhost")
	DatabasePort.setDefault(5432)
	DatabaseUser.setDefault("vayu")
	DatabasePassword.setDefault("password")
	DatabaseDatabase.setDefault("vayu")
	DatabaseMaxOpenConnections.setDefault(100)
	DatabaseMaxIdleConnections.setDefault(50)
	DatabaseMaxConnectionLifetime.setDefault(10000)
	DatabaseSslMode.setDefault("disable")
}

// InitConfig initializes the configuration.
func InitConfig() {
	initDefaultConfig()

	viper.SetEnvPrefix("vayu")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config")

	err := viper.ReadInConfig()
	if err != nil {
		slog.Warn(err.Error())
		slog.Warn("Using default config.")
	}
}
