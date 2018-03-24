package controller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
	commonDto "promise/common/object/dto"
	commonMessage "promise/common/object/message"
	"promise/ws/object/dto"
	"promise/ws/service"
)

// WsSenderRootController is the root controller.
type WsSenderRootController struct {
	beego.Controller
}

// Post handles POST requests.
func (c *WsSenderRootController) Post() {
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
