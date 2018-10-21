package base

import (
	"fmt"
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

// Init will init the app.
func Init(appName string) {
	// err := beego.LoadAppConfig("ini", "/opt/promise/conf/app.conf")
	// if err != nil {
	// 	panic(err)
	// }
	// port, err := beego.AppConfig.Int("Port")
	// if err != nil {
	// 	panic(err)
	// }
	// beego.BConfig.Listen.HTTPPort = port
	beego.BConfig.Listen.HTTPPort = 80
	beego.BConfig.CopyRequestBody = true
	log.SetFormatter(&LogTextFormatter{App: fmt.Sprintf("%-12.12s", appName), ForceColors: true, DisableSorting: false})
	log.SetLevel(log.InfoLevel)
	file, err := os.OpenFile("/opt/promise/log/promise.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		mw := io.MultiWriter(os.Stdout, file)
		log.SetOutput(mw)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}
