package main

import (
	commonDB "promise/common/db"
	"promise/common/app"
	"promise/task/controller"
	"promise/task/object/entity"

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
}

func main() {
	app.ReadConfig("TaskApp")	
	beego.SetLevel(beego.LevelDebug)
	initDB()
	ns := beego.NewNamespace(
		"/promise/v1/task",
		beego.NSRouter("/", &controller.RootController{}),
		beego.NSRouter("/:id", &controller.TaskController{}),
		beego.NSRouter("/:id/action/:action", &controller.TaskActionController{}),
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
