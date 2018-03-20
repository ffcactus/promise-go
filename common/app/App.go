package app

import (
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	// ProtocolScheme is the protocol scheme used by client.
	ProtocolScheme = "http://"
	// Host is the host used by client.
	Host = "localhost"
	// RootURL is the root URI for all the service.
	RootURL = "/promise/v1"
)

// Init will init the app.
func Init(appName string) {
	err := beego.LoadAppConfig("ini", "/opt/promise/conf/promise.conf")
	if err != nil {
		panic(err)
	}
	port, err := beego.AppConfig.Int(appName + "Port")
	if err != nil {
		panic(err)
	}
	beego.BConfig.Listen.HTTPPort = port
	log.SetFormatter(&PromiseTextFormatter{App: appName, ForceColors: true})
	log.SetLevel(log.InfoLevel)
	file, err := os.OpenFile("/tmp/promise.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}
