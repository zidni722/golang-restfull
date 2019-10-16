package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/zidni722/golang-restfull/bootstrap"
	"github.com/zidni722/golang-restfull/config"
	"github.com/zidni722/golang-restfull/routes"
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
	port := viper.GetString("app.server_port")

	cfg := config.New(app.Application)
	cfg.SetupLog()
	cfg.SetupDatabase()

	setupRoute(app, cfg)

	fmt.Print("Application is running on port :" + port)

	app.Listen(":" + port)

	defer cfg.Database.DB.Close()
}
