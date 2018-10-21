package base

import (
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// IDControllerTemplateInterface is the interface that a concrete ID controller should provide.
type IDControllerTemplateInterface interface {
	ResourceName() string
	Service() CRUDServiceInterface
	Response() GetResponseInterface
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
		response = c.TemplateImpl.Response()
	)
	log.WithFields(log.Fields{
		"resource": c.TemplateImpl.ResourceName(),
		"id":       id,
	}).Debug("IDController get resource.")
	model, errorResps := c.TemplateImpl.Service().Get(id)
	if errorResps != nil {
		log.WithFields(log.Fields{
			"resource":  c.TemplateImpl.ResourceName(),
			"id":        id,
			"errorResp": errorResps[0].ID,
		}).Warn("IDController get resource failed.")
		c.Data["json"] = &errorResps
		c.Ctx.Output.SetStatus(errorResps[0].StatusCode)
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
	if errorResps := c.TemplateImpl.Service().Delete(id); errorResps != nil {
		log.WithFields(log.Fields{
			"resource":  c.TemplateImpl.ResourceName(),
			"id":        id,
			"errorResp": errorResps[0].ID,
		}).Warn("IDController delete resource failed.")
		c.Data["json"] = &errorResps
		c.Ctx.Output.SetStatus(errorResps[0].StatusCode)
		c.ServeJSON()
		return
	}
	log.WithFields(log.Fields{
		"resource": c.TemplateImpl.ResourceName(),
		"id":       id,
	}).Info("Delete resource done.")
	c.Ctx.Output.SetStatus(http.StatusAccepted)
	c.ServeJSON()
}
