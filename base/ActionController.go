package base

import (
	"encoding/json"
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// ActionInfo includes all the information that is required by controller to
// handler action.
type ActionInfo struct {
	Name    string
	Request UpdateActionRequestInterface
	Service ServiceInterface
}

// ActionControllerTemplateInterface is the interface that a concrete controller must implement.
type ActionControllerTemplateInterface interface {
	GetResourceName() string
	GetActionInfo() []ActionInfo
}

// ActionController is the controller to handle actions.
type ActionController struct {
	TemplateImpl ActionControllerTemplateInterface
	beego.Controller
}

// Post is the default method to handle POST method.
func (c *ActionController) Post() {
	var (
		messages   []Message
		action     = c.Ctx.Input.Param(":action")
		id         = c.Ctx.Input.Param(":id")
		actionInfo = c.TemplateImpl.GetActionInfo()
		service    ServiceInterface
		request    UpdateActionRequestInterface
	)

	log.WithFields(log.Fields{
		"resource": c.TemplateImpl.GetResourceName(),
		"action":   action,
		"id":       id,
	}).Info("Perform action.")

	// Find the matching ActionInfo.s
	for _, v := range actionInfo {
		if strings.ToLower(action) == strings.ToLower(v.Name) {
			service = v.Service
			request = v.Request
		}
	}
	if service == nil {
		messages = append(messages, NewMessageInvalidRequest())
		log.WithFields(log.Fields{
			"resource": c.TemplateImpl.GetResourceName(),
			"action":   action,
			"id":       id,
			"message":  messages[0].ID,
		}).Warn("Perform action failed, unknown action.")
		c.Data["json"] = &messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	// Unmarshal the request.
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, request); err != nil {
		messages = append(messages, NewMessageInvalidRequest())
		log.WithFields(log.Fields{
			"resource": c.TemplateImpl.GetResourceName(),
			"action":   action,
			"id":       id,
			"error":    err,
			"message":  messages[0].ID,
		}).Warn("Post resource failed, bad request.")
		c.Data["json"] = &messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	// Validate the request.
	if message := request.IsValid(); message != nil {
		messages = append(messages, *message)
		log.WithFields(log.Fields{
			"resource": c.TemplateImpl.GetResourceName(),
			"action":   action,
			"id":       id,
			"message":  messages[0].ID,
		}).Warn("Post resource failed, request validation failed.")
		c.Data["json"] = &messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}

	response, messages := service.Perform(id, request)
	if messages != nil {
		log.WithFields(log.Fields{
			"resource": c.TemplateImpl.GetResourceName(),
			"action":   action,
			"id":       id,
			"message":  messages[0].ID,
		}).Warn("Post resource failed.")
		c.Data["json"] = &messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	log.WithFields(log.Fields{
		"resource": c.TemplateImpl.GetResourceName(),
		"action":   action,
		"id":       id,
	}).Info("Perform action done.")
	c.Data["json"] = response
	c.Ctx.Output.SetStatus(http.StatusAccepted)
	c.ServeJSON()
}
