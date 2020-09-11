package main

import (
	"flag"
	"go-url-shortener/app/apiserver"
	"go-url-shortener/helpers"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "config file path")
}
func main() {
	config := apiserver.InitConfig(configPath)
	helpers.FailOnError(apiserver.Start(config), "Failed on starting API server")

}
