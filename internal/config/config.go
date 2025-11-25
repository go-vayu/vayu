package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Key string

const (
	OPENAQ_API_KEY Key = `openaq.apikey`
)

func (k Key) Get() any {
	return viper.Get(string(k))
}

func (k Key) GetString() string {
	return viper.GetString(string(k))
}

func (k Key) Set(value any) {
	viper.SetDefault(string(k), value)
}

func initDefaultConfig() {
	OPENAQ_API_KEY.Set("")
}

func InitConfig() {
	initDefaultConfig()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	viper.SetEnvPrefix("vayu")
	viper.EnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}
