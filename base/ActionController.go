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
		errorResps []ErrorResponse
		action     = c.Ctx.Input.Param(":action")
		id         = c.Ctx.Input.Param(":id")
		actionInfo = c.TemplateImpl.ActionInfo()

		request            RequestInterface
		updateRequest      UpdateRequestInterface
		actionRequest      ActionRequestInterface
		asychActionRequest AsychActionRequestInterface

		service            ServiceInterface
		updateService      CRUDServiceInterface
		actionService      ActionServiceInterface
		asychActionService AsychActionServiceInterface

		response   ResponseInterface
		taskURI    string
		actionType string
	)

	// Find the matching ActionInfo.s
	for _, v := range actionInfo {
		if strings.ToLower(action) == strings.ToLower(v.Name) {
			service = v.Service
			actionType = v.Type
			// We need create a new instance here.
			if v.Request == nil {
				request = nil
			} else {
				request = v.Request.NewInstance()
			}
		}
	}
	if service == nil {
		errorResps = append(errorResps, *NewErrorResponseInvalidRequest())
		log.WithFields(log.Fields{
			"resource":  c.TemplateImpl.ResourceName(),
			"action":    action,
			"type":      actionType,
			"id":        id,
			"errorResp": errorResps[0].ID,
		}).Warn("ActionController perform action failed, unknown action.")
		c.Data["json"] = &errorResps
		c.Ctx.Output.SetStatus(errorResps[0].StatusCode)
		c.ServeJSON()
		return
	}
	if request != nil {
		// Unmarshal the request.
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, request); err != nil {
			errorResps = append(errorResps, *NewErrorResponseInvalidRequest())
			log.WithFields(log.Fields{
				"resource":  c.TemplateImpl.ResourceName(),
				"action":    action,
				"type":      actionType,
				"id":        id,
				"error":     err,
				"errorResp": errorResps[0].ID,
			}).Warn("ActionController perform action failed, bad request.")
			c.Data["json"] = &errorResps
			c.Ctx.Output.SetStatus(errorResps[0].StatusCode)
			c.ServeJSON()
			return
		}
		// Validate the request.
		if errorResp := request.IsValid(); errorResp != nil {
			errorResps = append(errorResps, *errorResp)
			log.WithFields(log.Fields{
				"resource":  c.TemplateImpl.ResourceName(),
				"action":    action,
				"type":      actionType,
				"id":        id,
				"errorResp": errorResps[0].ID,
			}).Warn("ActionController perform action failed, request validation failed.")
			c.Data["json"] = &errorResps
			c.Ctx.Output.SetStatus(errorResps[0].StatusCode)
			c.ServeJSON()
			return
		}
	}

	log.WithFields(log.Fields{
		"resource": c.TemplateImpl.ResourceName(),
		"action":   action,
		"request":  request,
		"id":       id,
	}).Info("ActionController perform action.")
	// Now the request is correct, select the right runtine by action type.
	ok := true
	switch actionType {
	case ActionTypeUpdate:
		updateRequest, updateService, ok = c.convertToUpdate(request, service)
		if ok {
			response, errorResps = updateService.Update(id, updateRequest)
		}
	case ActionTypeSych:
		actionRequest, actionService, ok = c.convertToAction(request, service)
		if ok {
			response, errorResps = actionService.Perform(id, actionRequest)
		}
	case ActionTypeAsych:
		// TODO can we return the result in action?
		asychActionRequest, asychActionService, ok = c.convertToAsychAction(request, service)
		if ok {
			response, taskURI, errorResps = asychActionService.PerformAsych(c.Ctx, id, asychActionRequest)
		}
	default:
		log.WithFields(log.Fields{
			"resource": c.TemplateImpl.ResourceName(),
			"action":   action,
			"type":     actionType,
			"id":       id,
		}).Warn("ActionController perform action failed, Unknown action type.")
		ok = false
	}
	if !ok {
		errorResps = []ErrorResponse{*NewErrorResponseInternalError()}
		log.WithFields(log.Fields{
			"resource":  c.TemplateImpl.ResourceName(),
			"action":    action,
			"type":      actionType,
			"id":        id,
			"errorResp": errorResps[0].ID,
		}).Warn("ActionController perform action failed, convert request and service failed.")
		c.Data["json"] = &errorResps
		c.Ctx.Output.SetStatus(errorResps[0].StatusCode)
		c.ServeJSON()
		return
	}
	if errorResps != nil {
		log.WithFields(log.Fields{
			"resource":  c.TemplateImpl.ResourceName(),
			"action":    action,
			"type":      actionType,
			"id":        id,
			"errorResp": errorResps[0].ID,
		}).Warn("ActionController perform action failed.")
		c.Data["json"] = &errorResps
		c.Ctx.Output.SetStatus(errorResps[0].StatusCode)
		c.ServeJSON()
		return
	}
	log.WithFields(log.Fields{
		"resource": c.TemplateImpl.ResourceName(),
		"action":   action,
		"type":     actionType,
		"id":       id,
		"response": response,
		"task":     taskURI,
	}).Info("Perform action done.")
	if taskURI != "" {
		c.Ctx.Output.Header("Location", taskURI)
	}
	if response != nil {
		c.Data["json"] = response
		c.Ctx.Output.SetStatus(http.StatusAccepted)
		c.ServeJSON()
	}
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
		log.Info("--- convert service failed")
		return nil, nil, ok
	}
	if request == nil {
		return nil, asychActionService, true
	}
	asychActionRequest, ok := request.(AsychActionRequestInterface)
	if !ok {
		return nil, nil, ok
	}
	return asychActionRequest, asychActionService, true
}
