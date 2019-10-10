package config

import (
	"github.com/kataras/iris"
	"github.com/spf13/viper"
)

type Configurator func(app *iris.Application)

type Configuration struct {
	*iris.Application
	Database *Database
	Log      *Log
}

func New(app *iris.Application) *Configuration {
	c := &Configuration{
		Application: app,
		Database: &Database{
			Connection: viper.GetString("database.connection"),
			Host:       viper.GetString("database.host"),
			Port:       viper.GetInt("database.port"),
			DBName:     viper.GetString("database.name"),
			Username:   viper.GetString("database.username"),
			Password:   viper.GetString("database.password"),
			Charset:    viper.GetString("database.charset"),
			Loc:        viper.GetString("database.loc"),
		},
		Log: &Log{},
	}

	return c
}

func (c *Configuration) SetupDatabase() {
	c.Database.Configure(c.Application)
}

func (c *Configuration) SetupLog() {
	c.Log.Configure(c.Application)
}

func (c *Configuration) Configure(cfgs ...Configurator) {
	for _, cfg := range cfgs {
		cfg(c.Application)
	}
}
