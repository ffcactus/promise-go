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
		// base.NSRouter("/node/action/:action", &base.ActionController{
		// 	TempalteImpl: new(controller.NodeRootAction),
		// }),
		// beego.NSRouter("node/:id", &base.IDController{
		// 	TemplateImpl: new(controller.NodeID),
		// }),	
	))
	// Add namespace for Service.
	// beego.AddNamespace(beego.NewNamespace(
	// 	base.RootURL+base.DirectorBaseURI,
	// 	beego.NSRouter("/service/", &base.RootController{
	// 		TemplateImpl: new(controller.NodeRoot),
	// 	}),
	// 	beego.NSRouter("service/:id", &base.IDController{
	// 		TemplateImpl: new(controller.NodeID),
	// 	}),
	// ))
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