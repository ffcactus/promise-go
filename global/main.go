package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"promise/base"
	"promise/global/controller"
)

func main() {
	base.Init("global")
	beego.SetLevel(beego.LevelDebug)

	ns := beego.NewNamespace(
		base.RootURL+"/global",
		beego.NSRouter("/", &controller.RootController{}),
	)

	beego.AddNamespace(ns)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.Run()
}
