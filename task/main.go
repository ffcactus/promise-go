package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/task/controller"
	"promise/task/object/entity"
)

func main() {
	base.Init("TaskApp")
	initDB()
	ns := beego.NewNamespace(
		base.RootURL+base.TaskBaseURI,
		beego.NSRouter("/", &base.RootController{
			TemplateImpl: new(controller.TaskRootController),
		}),
		beego.NSRouter("/:id", &base.IDController{
			TemplateImpl: new(controller.TaskIDController),
		}),
		// beego.NSRouter("/:id/action/:action", &controller.TaskActionController{}),
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
}
