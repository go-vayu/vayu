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
)

// Get returns the configuration value.
func (k Key) Get() any {
	return viper.Get(string(k))
}

// GetString returns the configuration value as a string.
func (k Key) GetString() string {
	return viper.GetString(string(k))
}

// Set sets the configuration value.
func (k Key) Set(value any) {
	viper.SetDefault(string(k), value)
}

func initDefaultConfig() {
	OpenAQAPIKey.Set("")
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
