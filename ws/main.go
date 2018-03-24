package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"promise/common/app"
	"promise/common/object/constValue"
	"promise/ws/controller"
	"promise/ws/service"
)

func main() {
	app.Init("WSApp")

	go service.StartEventDispatcher()

	// ws namesapce.
	wsNS := beego.NewNamespace(
		app.RootURL+constValue.WSBaseURI,
		beego.NSRouter("/", &controller.RootController{}),
	)
	beego.AddNamespace(wsNS)

	// ws-sender namespace.
	wsSenderNS := beego.NewNamespace(
		app.RootURL+constValue.WSSenderBaseURI,
		beego.NSRouter("/", &controller.WsSenderRootController{}),
	)
	beego.AddNamespace(wsSenderNS)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.Run()
}
