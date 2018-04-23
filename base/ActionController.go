package base

import (
	"encoding/json"
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	// ActionTypeUpdate means the update action.
	ActionTypeUpdate = "Update"
	// ActionTypeSych means the sychronous action.
	ActionTypeSych = "Sychronous"
	// ActionTypeAsych means the asychronous action.
	ActionTypeAsych = "Asychronous"
)

// ActionInfo includes all the information that is required by controller to
// handler action.
type ActionInfo struct {
	Name    string           // The action name.
	Type    string           // The type of this action.
	Request RequestInterface // we need create a new one each time.
	Service ServiceInterface
}

// ActionControllerTemplateInterface is the interface that a concrete controller must implement.
type ActionControllerTemplateInterface interface {
	ResourceName() string
	ActionInfo() []ActionInfo
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
		actionInfo = c.TemplateImpl.ActionInfo()
		service    ServiceInterface
		request    RequestInterface
		response   ResponseInterface
		taskURI    *string
		actionType string
	)

	// Find the matching ActionInfo.s
	for _, v := range actionInfo {
		if strings.ToLower(action) == strings.ToLower(v.Name) {
			service = v.Service
			actionType = v.Type
			// We need create a new instance here.
			request = v.Request.NewInstance()
		}
	}
	if service == nil {
		messages = append(messages, NewMessageInvalidRequest())
		log.WithFields(log.Fields{
			"resource": c.TemplateImpl.ResourceName(),
			"action":   action,
			"type":     actionType,
			"id":       id,
			"message":  messages[0].ID,
		}).Warn("Perform action failed, unknown action.")
		c.Data["json"] = &messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}

	if request != nil {
		// Unmarshal the request.
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, request); err != nil {
			messages = append(messages, NewMessageInvalidRequest())
			log.WithFields(log.Fields{
				"resource": c.TemplateImpl.ResourceName(),
				"action":   action,
				"type":     actionType,
				"id":       id,
				"error":    err,
				"message":  messages[0].ID,
			}).Warn("Perform action failed, bad request.")
			c.Data["json"] = &messages
			c.Ctx.Output.SetStatus(messages[0].StatusCode)
			c.ServeJSON()
			return
		}
		// Validate the request.
		if message := request.IsValid(); message != nil {
			messages = append(messages, *message)
			log.WithFields(log.Fields{
				"resource": c.TemplateImpl.ResourceName(),
				"action":   action,
				"type":     actionType,
				"id":       id,
				"message":  messages[0].ID,
			}).Warn("Perform action failed, request validation failed.")
			c.Data["json"] = &messages
			c.Ctx.Output.SetStatus(messages[0].StatusCode)
			c.ServeJSON()
			return
		}
	}

	requestDebugInfo := ""
	if request != nil {
		requestDebugInfo = request.DebugInfo()
	}
	log.WithFields(log.Fields{
		"resource": c.TemplateImpl.ResourceName(),
		"action":   action,
		"request":  requestDebugInfo,
		"id":       id,
	}).Info("Perform action.")
	// Now the request is correct, select the right runtine by action type.
	ok := true
	switch actionType {
	case ActionTypeUpdate:
		updateRequest, updateService, ok := c.convertToUpdate(request, service)
		if ok {
			response, messages = updateService.Update(id, updateRequest)
		}
	case ActionTypeSych:
		actionRequest, actionService, ok := c.convertToAction(request, service)
		if ok {
			response, messages = actionService.Perform(id, actionRequest)
		}
	case ActionTypeAsych:
		asychActionRequest, asychActionService, ok := c.convertToAsychAction(request, service)
		if ok {
			response, taskURI, messages = asychActionService.PerformAsych(id, asychActionRequest)
		}
	default:
		log.WithFields(log.Fields{
			"resource": c.TemplateImpl.ResourceName(),
			"action":   action,
			"type":     actionType,
			"id":       id,
		}).Warn("Perform action failed, Unknown action type.")
		ok = false
	}
	if !ok {
		messages = []Message{NewMessageInternalError()}
		log.WithFields(log.Fields{
			"resource": c.TemplateImpl.ResourceName(),
			"action":   action,
			"type":     actionType,
			"id":       id,
			"message":  messages[0].ID,
		}).Warn("Perform action failed, convert request and service failed.")
		c.Data["json"] = &messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	if messages != nil {
		log.WithFields(log.Fields{
			"resource": c.TemplateImpl.ResourceName(),
			"action":   action,
			"type":     actionType,
			"id":       id,
			"message":  messages[0].ID,
		}).Warn("Perform action failed.")
		c.Data["json"] = &messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	log.WithFields(log.Fields{
		"resource": c.TemplateImpl.ResourceName(),
		"action":   action,
		"type":     actionType,
		"id":       id,
		"response": response.DebugInfo(),
		"task":     taskURI,
	}).Info("Perform action done.")
	if taskURI != nil {
		c.Ctx.Output.Header("Location", *taskURI)
	}
	c.Data["json"] = response
	c.Ctx.Output.SetStatus(http.StatusAccepted)
	c.ServeJSON()
}

func (c *ActionController) convertToUpdate(request RequestInterface, service ServiceInterface) (UpdateRequestInterface, CRUDServiceInterface, bool) {
	updateService, ok := service.(CRUDServiceInterface)
	if !ok {
		return nil, nil, ok
	}
	if request == nil {
		return nil, updateService, true
	}
	updateRequest, ok := request.(UpdateRequestInterface)
	if !ok {
		return nil, nil, ok
	}
	return updateRequest, updateService, true
}

func (c *ActionController) convertToAction(request RequestInterface, service ServiceInterface) (ActionRequestInterface, ActionServiceInterface, bool) {
	actionService, ok := service.(ActionServiceInterface)
	if !ok {
		return nil, nil, ok
	}
	if request == nil {
		return nil, actionService, true
	}
	actionRequest, ok := request.(ActionRequestInterface)
	if !ok {
		return nil, nil, ok
	}
	return actionRequest, actionService, true
}

func (c *ActionController) convertToAsychAction(request RequestInterface, service ServiceInterface) (AsychActionRequestInterface, AsychActionServiceInterface, bool) {
	asychActionService, ok := service.(AsychActionServiceInterface)
	if !ok {
		return nil, nil, ok
	}
	if request == nil {
		return nil, asychActionService, true
	}
	asychActionRequest, ok := request.(UpdateRequestInterface)
	if !ok {
		return nil, nil, ok
	}
	return asychActionRequest, asychActionService, true
}
