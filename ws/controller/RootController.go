package controller

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"promise/ws/service"
)

// RootController is the root controller.
type RootController struct {
	beego.Controller
}

// Get Handles GET requests.
func (c *RootController) Get() {
	// Upgrade from http request to WebSocket.
	log.WithFields(log.Fields{
		"remote": c.Ctx.Request.RemoteAddr,
	}).Info("There is a websocket connection request.")
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024*64, 1024*64)

	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		log.WithFields(log.Fields{
			"remote": c.Ctx.Request.RemoteAddr,
		}).Warn("Not a websocket handshake.")
		return
	} else if err != nil {
		log.WithFields(log.Fields{
			"remote": c.Ctx.Request.RemoteAddr,
			"error":  err,
		}).Warn("Cannot setup websocket connection.")
		return
	}
	count := service.AddListener(ws)
	log.WithFields(log.Fields{
		"count":  count,
		"remote": c.Ctx.Request.RemoteAddr,
	}).Info("Websocket add a listener.")
}
