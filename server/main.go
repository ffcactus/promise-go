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
	base.InitConnection()
	if recreateDB, _ := beego.AppConfig.Bool("recreate_db"); recreateDB {
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
	}
	service.CreateDefaultServerGroup()
}

func main() {
	base.Init("ServerApp")
	initDB()
	// go service.FindServerStateAdded()
	serverNS := beego.NewNamespace(
		base.RootURL+base.ServerBaseURI,
		beego.NSRouter("/", &base.IDController{
			TemplateImpl: new(controller.ServerRootController),
		}),
		beego.NSRouter("/:id", &base.IDController{
			TemplateImpl: new(controller.ServerIDController),
		}),
		beego.NSRouter("/:id/action/:action", &base.IDController{
			TemplateImpl: new(controller.ServerActionController),
		}),
	)
	beego.AddNamespace(serverNS)

	serverGroupNS := beego.NewNamespace(
		base.RootURL+base.ServerGroupBaseURI,
		beego.NSRouter("/", &base.RootController{
			TemplateImpl: new(controller.ServerGroupRootController),
		}),
		beego.NSRouter("/:id", &base.IDController{
			TemplateImpl: new(controller.ServerGroupIDController),
		}),
	)
	beego.AddNamespace(serverGroupNS)

	serverServerGroupNS := beego.NewNamespace(
		base.RootURL+base.ServerServerGroupBaseURI,
		beego.NSRouter("/", &base.IDController{
			TemplateImpl: new(controller.ServerServerGroupRootController),
		}),
		beego.NSRouter("/:id", &base.IDController{
			TemplateImpl: new(controller.ServerServerGroupIDController),
		}),
	)
	beego.AddNamespace(serverServerGroupNS)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.Run()
}
