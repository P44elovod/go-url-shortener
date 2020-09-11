package apiserver

import (
	"flag"
	"go-url-shortener/helpers"

	"github.com/BurntSushi/toml"
)

type Config struct {
	BindPort    string `toml:"bind_port"`
	DatabaseURL string `toml:"db_url"`
}

func NewConfig() *Config {
	return &Config{
		BindPort: ":8080",
	}
}

func InitConfig(path string) (config *Config) {
	flag.Parse()

	config = NewConfig()
	_, err := toml.DecodeFile(path, config)
	helpers.FailOnError(err, "Failed at toml decode")

	return
}
