package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/student/controller"
	"promise/student/object/entity"
)

func main() {
	base.Init("StudentApp")
	initDB()

	student := beego.NewNamespace(
		base.RootURL+base.StudentBaseURI,
		beego.NSRouter("/", &base.RootController{
			TemplateImpl: new(controller.StudentRootController),
		}),
		beego.NSRouter("/:id", &base.IDController{
			TemplateImpl: new(controller.StudentIDController),
		}),
	)
	beego.AddNamespace(student)
	// beego.Post("/", controller.Post)

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
