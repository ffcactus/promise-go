package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	log "github.com/sirupsen/logrus"
	"os"
	"promise/base"
	"promise/enclosure/controller"
	"promise/enclosure/object/entity"
)

func recreateDB() bool {
	args := os.Args[1:]
	for _, v := range args {
		if v == "recreatedb" {
			return true
		}
	}
	return false
}

func initDB() {
	if err := base.InitConnection("enclosure"); err != nil {
		log.Error("Init DB failed, App exit.")
		panic("Init DB failed, App exit.")
	}
	if recreateDB() {
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

func main() {
	base.Init("enclosure")
	base.InitMQService()
	defer base.StopMQService()
	initDB()

	enclosureNS := beego.NewNamespace(
		base.RootURL+base.EnclosureBaseURI,
		beego.NSRouter("/", &base.RootController{
			TemplateImpl: new(controller.EnclosureRoot),
		}),
		beego.NSRouter("/action/:action", &base.ActionController{
			TemplateImpl: new(controller.EnclosureRootAction),
		}),
		beego.NSRouter("/:id", &base.IDController{
			TemplateImpl: new(controller.EnclosureID),
		}),
		beego.NSRouter("/:id/action/:action", &base.ActionController{
			TemplateImpl: new(controller.EnclosureAction),
		}),
	)
	beego.AddNamespace(enclosureNS)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.Run()
}
