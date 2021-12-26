package main

import (
	application "pandudpn/api/app"

	"github.com/spf13/viper"
)

func main() {
	initConfig()

	app := application.NewApp()
	app.Register()
}

func initConfig() {
	viper.AutomaticEnv()
}
