package main

import (
	"fmt"
	"github.com/zidni722/golang-restfull/bootstrap"
	"github.com/spf13/viper"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New(viper.GetString("app.name"), viper.GetString("app.owner"))
	app.Bootstrap()

	return app
}

func readConfig() {
	viper.SetConfigName("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found...")
	} else {
		viper.SetEnvPrefix("product-api")
		viper.AllowEmptyEnv(true)
		viper.AutomaticEnv()
	}
}

func setupRoute(app *bootstrap.Bootstrapper, cfg *config.Configuration) {
	route := routes.NewRoute(cfg)
	app.Configure(route.Configure)
}

func main() {
	readConfig()

	app := newApp()

	cfg := config.New(app.Application)
	cfg.SetupLog()
	cfg.SetupDatabase()

	setupRoute(app, cfg)

	app.Listen(":9090")

	defer cfg.Database.DB.Close()
}