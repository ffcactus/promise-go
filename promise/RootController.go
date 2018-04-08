package promise

import (
	"encoding/json"
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
)

// RootControllerInterface is the interface that a root controller should provide.
type RootControllerInterface interface {
	GetResourceName() string
	Post()
	newRequest() RequestInterface
	newResponse() ResponseInterface
	newModel()
}

// RootController is the root controller in Promise.
type RootController struct {
	beego.Controller
}

func (c *RootController) newRequest() RequestInterface {
	return new(Request)
}

func (c *RootController) newResponse() ResponseInterface {
	return new(Response)
}

// GetResourceName return the resource name.
func (c *RootController) GetResourceName() string {
	return "root"
}

// Post is the default implementation for POST method.
func (c *RootController) Post() {
	var (
		request  = c.newRequest()
		response = c.newResponse()
		messages []MessageInterface
	)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, request); err != nil {
		messages = append(messages, NewMessageInvalidRequest())
		log.WithFields(log.Fields{
			"resource": c.GetResourceName(),
			"error":    err,
			"message":  messages[0].GetID()}).
			Warn("Post IPv4 pool failed, bad request.")
		c.Data["json"] = messages
		c.Ctx.Output.SetStatus(messages[0].GetStatusCode())
		c.ServeJSON()
	}

	log.WithFields(log.Fields{
		"requestName": request.GetDebugName(),
	}).Info("Post resource.")
}
