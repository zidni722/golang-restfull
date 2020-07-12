package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

func GetHomeHandler(ctx iris.Context) {
	ctx.ViewData("Title", "Home")
	ctx.ViewData("AppName", viper.GetString("app.name"))
	ctx.ViewData("AppOwner", viper.GetString("app.owner"))
	ctx.ViewData("Message", "Welcome to golang-restfull!")
	ctx.View("index.html")
}
