package controller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	commonDto "promise/common/object/dto"
	commonMessage "promise/common/object/message"
	"promise/ws/object/dto"
	"promise/ws/service"
)

// RootController is the root controller.
type RootController struct {
	beego.Controller
}

// Get Handles GET requests.
func (c *RootController) Get() {
	// Upgrade from http request to WebSocket.
	log.WithFields(log.Fields{"remote": c.Ctx.Request.RemoteAddr}).Info("There is a websocket connection request.")
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		log.WithFields(log.Fields{"remote": c.Ctx.Request.RemoteAddr}).Warn("Not a websocket handshake.")
		return
	} else if err != nil {
		log.WithFields(log.Fields{"remote": c.Ctx.Request.RemoteAddr, "error": err}).Warn("Cannot setup websocket connection.")
		return
	}
	service.AddListener(ws)
	log.WithFields(log.Fields{"remote": c.Ctx.Request.RemoteAddr}).Info("Websocket add a listener.")
}

// Post handles POST requests.
func (c *RootController) Post() {
	var request dto.PostEventRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("Post ws failed, unable to unmarshal request.")
		messages := []commonMessage.Message{}
		messages = append(messages, commonMessage.NewInvalidRequest())
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	service.DispatchEvent(&request)
	c.Ctx.Output.SetStatus(http.StatusCreated)
	log.WithFields(log.Fields{
		"category": request.Category,
		"type":     request.Type,
		"resource": request.ResourceID}).Info("Post ws message done.")
	c.ServeJSON()
}
