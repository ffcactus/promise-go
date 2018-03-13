package controller

import (
	"net/http"
	"promise/ws/service"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

// RootController is the root controller.
type RootController struct {
	beego.Controller
}

// Get Handles GET requests for WebSocketController.
func (c *RootController) Get() {
	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Warn("Cannot setup WebSocket connection, error =", err)
		return
	}
	service.AddListener(ws)
}
