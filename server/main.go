package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/controller"
	"promise/server/object/entity"
	"promise/server/service"
)

func initDB() {
	if err := base.InitConnection(); err != nil {
		log.Error("Init DB failed, App exit.")
		panic("Init DB failed, App exit.")
	}
	// if recreateDB, _ := beego.AppConfig.Bool("recreate_db"); recreateDB {
	// Remove tables.
	if base.RemoveTables(entity.Tables) {
		log.Info("Remove all tables in DB done.")
	} else {
		log.Warn("Failed to remove all tables in DB.")
	}
	// Create tables.
	if !base.CreateTables(entity.Tables) {
		panic("DB Initialization failed.")
	} else {
		log.Info("DB schema created.")
	}
	// }
	service.LoadModel()
	service.CreateDefaultServerGroup()
}

func main() {
	base.Init("ServerApp")
	initDB()
	bgRefresh := service.Refresh{}
	bgRefresh.StartBackgroundRefresh()

	serverNS := beego.NewNamespace(
		base.RootURL+base.ServerBaseURI,
		beego.NSRouter("/", &base.RootController{
			TemplateImpl: new(controller.ServerRoot),
		}),
		beego.NSRouter("/action/:action", &base.ActionController{
			TemplateImpl: new(controller.ServerRootAction),
		}),
		beego.NSRouter("/:id", &base.IDController{
			TemplateImpl: new(controller.ServerID),
		}),
		beego.NSRouter("/:id/action/:action", &base.ActionController{
			TemplateImpl: new(controller.ServerAction),
		}),
	)
	beego.AddNamespace(serverNS)

	serverGroupNS := beego.NewNamespace(
		base.RootURL+base.ServerGroupBaseURI,
		beego.NSRouter("/", &base.RootController{
			TemplateImpl: new(controller.ServerGroupRoot),
		}),
		beego.NSRouter("/:id", &base.IDController{
			TemplateImpl: new(controller.ServerGroupID),
		}),
	)
	beego.AddNamespace(serverGroupNS)

	serverServerGroupNS := beego.NewNamespace(
		base.RootURL+base.ServerServerGroupBaseURI,
		beego.NSRouter("/", &base.RootController{
			TemplateImpl: new(controller.ServerServerGroupRoot),
		}),
		beego.NSRouter("/:id", &base.IDController{
			TemplateImpl: new(controller.ServerServerGroupID),
		}),
	)
	beego.AddNamespace(serverServerGroupNS)

	adapterConfigNS := beego.NewNamespace(
		base.RootURL+base.AdapterConfigBaseURI,
		beego.NSRouter("/", &base.RootController{
			TemplateImpl: new(controller.AdapterConfigRoot),
		}),
		beego.NSRouter("/:id", &base.IDController{
			TemplateImpl: new(controller.AdapterConfigID),
		}),
	)
	beego.AddNamespace(adapterConfigNS)

	adapterModelNS := beego.NewNamespace(
		base.RootURL+base.AdapterModelBaseURI,
		beego.NSRouter("/", &base.RootController{
			TemplateImpl: new(controller.AdapterModelRoot),
		}),
		beego.NSRouter("/:id", &base.IDController{
			TemplateImpl: new(controller.AdapterModelID),
		}),
	)
	beego.AddNamespace(adapterModelNS)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.Run()
}
