package main

import (
	"promise/auth/controller"
	"promise/auth/object/entity"
	"promise/auth/service"
	"promise/common/app"
	commonDB "promise/common/db"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	log "github.com/sirupsen/logrus"
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
	service.CreateDefaultAdmin()
}

func main() {
	app.InitLog()
	app.ReadConfig("AuthApp")
	beego.SetLevel(beego.LevelDebug)
	initDB()
	ns := beego.NewNamespace(
		app.RootURL+"/auth",
		beego.NSRouter("/login", &controller.LoginController{}),
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
