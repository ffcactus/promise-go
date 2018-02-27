package controller

import (
	"github.com/astaxie/beego"
)

// RootController is the root controller.
type RootController struct {
	beego.Controller
}

// Get Handles GET requests for WebSocketController.
func (c *RootController) Get() {
	beego.Trace("POST Login request.")
	c.ServeJSON()
}
