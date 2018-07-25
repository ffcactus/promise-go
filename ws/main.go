package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"promise/base"
	"promise/ws/controller"
	"promise/ws/service"
)

func main() {
	base.Init("WSApp")

	go service.StartDispatcher()

	// ws namesapce.
	wsNS := beego.NewNamespace(
		base.RootURL+base.WSBaseURI,
		beego.NSRouter("/", &controller.RootController{}),
	)
	beego.AddNamespace(wsNS)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.Run()
}
