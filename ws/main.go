package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"promise/common/app"
	"promise/ws/controller"
	"promise/ws/service"
)

func main() {
	app.Init("WSApp")
	beego.SetLevel(beego.LevelDebug)
	go service.StartEventDispatcher()

	ns := beego.NewNamespace(
		app.RootURL+"/ws",
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
