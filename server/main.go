package main

import (
	"promise/common/app"
	commonDB "promise/common/db"
	"promise/server/controller"
	"promise/server/object/entity"
	"promise/server/service"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	log "github.com/sirupsen/logrus"
)

func main() {
	app.Init("ServerApp")
	initDB()

	go service.FindServerStateAdded()

	serverNS := beego.NewNamespace(
		app.RootURL+"/server",
		beego.NSRouter("/", &controller.ServerRootController{}),
		beego.NSRouter("/:id", &controller.ServerController{}),
		beego.NSRouter("/:id/action/:action", &controller.ServerActionController{}),
	)
	beego.AddNamespace(serverNS)
	serverGroupNS := beego.NewNamespace(
		app.RootURL+"/servergroup",
		beego.NSRouter("/", &controller.GroupRootController{}),
		beego.NSRouter("/:id", &controller.GroupController{}),
		// beego.NSRouter("/:id/action/:action", &controller.GroupActionController{}),

	)
	beego.AddNamespace(serverGroupNS)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.Run()
}

func initDB() {
	commonDB.InitConnection()
	if recreateDB, _ := beego.AppConfig.Bool("recreate_db"); recreateDB {
		// Remove tables.
		if commonDB.RemoveTables(entity.Tables) {
			log.Info("Remove all tables in DB done.")
		} else {
			log.Warn("Failed to remove all tables in DB.")
		}
		// Create tables.
		if !commonDB.CreateTables(entity.Tables) {
			panic("DB Initialization failed.")
		} else {
			log.Info("DB schema created.")
		}
	}
	service.CreateDefaultGroup()
}
