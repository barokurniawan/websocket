package context

import (
	"github.com/BurntSushi/toml"
)

type AppConfig struct {
	Port           string
	Address        string
	AllowedOrigins []string
}

var Config AppConfig

func LoadConfig() {
	if _, err := toml.DecodeFile("appconfig.toml", &Config); err != nil {
		panic(err.Error())
	}
}
