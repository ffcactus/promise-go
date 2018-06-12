package base

import (
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"os"
)

// Init will init the app.
func Init(appName string) {
	err := beego.LoadAppConfig("ini", "/opt/promise/conf/app.conf")
	if err != nil {
		panic(err)
	}
	port, err := beego.AppConfig.Int("Port")
	if err != nil {
		panic(err)
	}
	beego.BConfig.Listen.HTTPPort = port
	log.SetFormatter(&LogTextFormatter{App: appName, ForceColors: true, DisableSorting: false})
	log.SetLevel(log.InfoLevel)
	file, err := os.OpenFile("/tmp/promise.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}
