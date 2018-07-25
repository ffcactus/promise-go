package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"promise/base"
	"promise/director/controller"
)

func main() {
	base.Init("Director")
	// Add namespace for Node.
	beego.AddNamespace(beego.NewNamespace(
		base.RootURL+base.DirectorBaseURI,
		beego.NSRouter("/node/", &controller.NodeRoot{
			RootController: base.RootController{
				TemplateImpl: new(controller.NodeRoot),
			},
		}),
	))

	// Enable CORS.
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	// Start.
	beego.Run()
}
