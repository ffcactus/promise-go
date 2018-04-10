package base

import (
	"encoding/json"
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// RootControllerInterface is the interface that a root controller should provide.
type RootControllerInterface interface {
	GetService() ServiceInterface
	NewRequest() RequestInterface
	NewResponse() ResponseInterface
	// PostCallback(request RequestInterface) (ModelInterface, []MessageInterface)
}

// RootController is the root controller in Promise.
type RootController struct {
	TemplateImpl RootControllerInterface
	beego.Controller
}

// Post is the default implementation for POST method.
func (c *RootController) Post() {
	var (
		request  = c.TemplateImpl.NewRequest()
		response = c.TemplateImpl.NewResponse()
		messages []MessageInterface
	)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, request); err != nil {
		messages = append(messages, NewMessageInvalidRequest())
		log.WithFields(log.Fields{
			"request": request.GetDebugName(),
			"error":   err,
			"message": messages[0].GetID(),
		}).Warn("Post resource failed, bad request.")
		c.Data["json"] = messages
		c.Ctx.Output.SetStatus(messages[0].GetStatusCode())
		c.ServeJSON()
		return
	}

	log.WithFields(log.Fields{
		"request": request.GetDebugName(),
	}).Info("Post resource.")
	model, messages := c.TemplateImpl.GetService().Post(request)
	if messages != nil {
		log.WithFields(log.Fields{
			"message": messages[0].GetID(),
		}).Warn("Post resource failed, POST callback return message.")
		c.Data["json"] = messages
		c.Ctx.Output.SetStatus(messages[0].GetStatusCode())
		c.ServeJSON()
		return
	}
	response.Load(model)
	log.WithFields(log.Fields{
		"request": request.GetDebugName(),
		"ID":      response.GetID(),
	}).Info("Post resource done.")
	c.Data["json"] = response
	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.ServeJSON()
}
