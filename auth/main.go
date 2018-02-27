package main

import (
	"promise/auth/controller"
	"promise/auth/object/entity"
	"promise/auth/service"
	commonDB "promise/common/db"
	"promise/common/app"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func initDB() {
	commonDB.InitConnection()
	if recreateDB, _ := beego.AppConfig.Bool("recreate_db"); recreateDB {
		// Remove tables.
		if commonDB.RemoveTables(entity.Tables) {
			beego.Info("Remove all tables in DB done.")
		} else {
			beego.Warning("Failed to remove all tables in DB.")
		}
		// Create tables.
		if !commonDB.CreateTables(entity.Tables) {
			panic("DB Initialization failed.")
		} else {
			beego.Info("DB schema created.")
		}
	}
	service.CreateDefaultAdmin()
}

func main() {
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
