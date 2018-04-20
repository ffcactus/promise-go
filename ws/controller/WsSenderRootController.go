package controller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
	"promise/base"
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
		log.WithFields(log.Fields{
			"error": err,
		}).Warn("Post ws failed, unable to unmarshal request.")
		messages := []base.Message{}
		messages = append(messages, base.NewMessageInvalidRequest())
		c.Data["json"] = messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	service.DispatchEvent(&request)
	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.ServeJSON()
}
