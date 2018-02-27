package app

import (
	"github.com/astaxie/beego"
)
const (
	// ProtocolScheme is the protocol scheme used by client.
	ProtocolScheme = "http://"
	// Host is the host used by client.
	Host = "localhost"
	// RootURL is the root URI for all the service.
	RootURL = "/promise/v1"
)

// ReadConfig will read config for this app.
func ReadConfig(appName string) {
	err := beego.LoadAppConfig("ini", "/opt/promise/conf/promise.conf")
	if err != nil {
		panic(err)
	}
	port, err := beego.AppConfig.Int(appName+"Port")
	if err != nil {
		panic(err)
	}
	beego.BConfig.Listen.HTTPPort = port
}