package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/pool/controller"
	"promise/pool/object/entity"
)

func initDB() {
	if err := base.InitConnection("pool"); err != nil {
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
}

func main() {
	base.Init("pool")
	base.InitMQService()
	defer base.StopMQService()
	initDB()
	ipv4 := beego.NewNamespace(
		base.RootURL+base.IDPoolBaseURI,
		beego.NSRouter("/ipv4", &base.RootController{
			TemplateImpl: new(controller.IPv4Root),
		}),
		beego.NSRouter("/ipv4/:id", &base.IDController{
			TemplateImpl: new(controller.IPv4ID),
		}),
		beego.NSRouter("/ipv4/:id/action/:action", &base.ActionController{
			TemplateImpl: new(controller.IPv4Action),
		}),
	)
	beego.AddNamespace(ipv4)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.Run()
}
