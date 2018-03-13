package controller

import (
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
)

// RootController is the root controller.
type RootController struct {
	beego.Controller
}

// Get Handles GET requests for WebSocketController.
func (c *RootController) Get() {
	log.Debug("POST Login request.")
	c.ServeJSON()
}
