package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	log "github.com/sirupsen/logrus"
	"promise/common/app"
	commonDB "promise/common/db"
	"promise/common/object/constValue"
	"promise/pool/controller"
	"promise/pool/object/entity"
)

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
}

func main() {
	app.Init("IDPoolApp")
	initDB()
	ipNS := beego.NewNamespace(
		app.RootURL+constValue.TaskBaseURI,
		beego.NSRouter("/ipv4", &controller.IPv4RootController{}),
		beego.NSRouter("/ipv4/:id", &controller.IPv4Controller{}),
	)
	beego.AddNamespace(ipNS)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.Run()
}
