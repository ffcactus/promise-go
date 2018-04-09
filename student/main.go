package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	log "github.com/sirupsen/logrus"
	"promise/apps"
	"promise/base"
	"promise/student/controller"
	"promise/student/object/entity"
	"promise/student/service"
)

func main() {
	apps.Init("StudentApp")
	initDB()

	studentRootController := new(controller.StudentRootController)
	studentRootController.Service = base.Service{
		Interface: new(service.StudentService),
	}
	student := beego.NewNamespace(
		apps.RootURL+apps.StudentBaseURI,
		beego.NSRouter("/", &base.RootController{
			Interface: studentRootController,
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
	apps.InitConnection()
	if recreateDB, _ := beego.AppConfig.Bool("recreate_db"); recreateDB {
		// Remove tables.
		if apps.RemoveTables(entity.Tables) {
			log.Info("Remove all tables in DB done.")
		} else {
			log.Warn("Failed to remove all tables in DB.")
		}
		// Create tables.
		if !apps.CreateTables(entity.Tables) {
			panic("DB Initialization failed.")
		} else {
			log.Info("DB schema created.")
		}
	}
}
