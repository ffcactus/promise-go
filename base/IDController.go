package base

import (
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// IDControllerTemplateInterface is the interface that a concrete ID controller should provide.
type IDControllerTemplateInterface interface {
	GetResourceName() string
	GetService() ServiceInterface
	NewResponse() ResponseInterface
}

// IDController is the controller that handle request on a specific resource.
// For example, the request to /rest/v1/student/0001
type IDController struct {
	TemplateImpl IDControllerTemplateInterface
	beego.Controller
}

// Get is the default GET method handler.
func (c *IDController) Get() {
	var (
		id       = c.Ctx.Input.Param(":id")
		response = c.TemplateImpl.NewResponse()
	)
	log.WithFields(log.Fields{
		"resource": c.TemplateImpl.GetResourceName(),
		"id":       id,
	}).Debug("Get resource.")
	model, messages := c.TemplateImpl.GetService().Get(id)
	if messages != nil {
		log.WithFields(log.Fields{
			"resource": c.TemplateImpl.GetResourceName(),
			"id":       id,
			"message":  messages[0].ID,
		}).Warn("Get resource failed.")
		c.Data["json"] = &messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	response.Load(model)
	c.Data["json"] = response
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.ServeJSON()
}

// Delete is the default DELETE method handler.
func (c *IDController) Delete() {
	var (
		id = c.Ctx.Input.Param(":id")
	)
	if messages := c.TemplateImpl.GetService().Delete(id); messages != nil {
		log.WithFields(log.Fields{
			"resource": c.TemplateImpl.GetResourceName(),
			"id":       id,
			"message":  messages[0].ID,
		}).Warn("Delete resource failed.")
		c.Data["json"] = &messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	log.WithFields(log.Fields{
		"resource": c.TemplateImpl.GetResourceName(),
		"id":       id,
	}).Info("Delete resource done.")
	c.Ctx.Output.SetStatus(http.StatusAccepted)
	c.ServeJSON()
}