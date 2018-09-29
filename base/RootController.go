package base

import (
	"encoding/json"
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// RootControllerTemplateInterface is the interface that concrete RootController should have.
type RootControllerTemplateInterface interface {
	ResourceName() string
	Service() CRUDServiceInterface
	Request() PostRequestInterface
	Response() GetResponseInterface
	ConvertCollectionModel(*CollectionModel) (interface{}, error)
}

// RootController is the controller that handle request on resource's root URL
// For example, the request to /rest/v1/student
type RootController struct {
	TemplateImpl RootControllerTemplateInterface
	beego.Controller
}

// Post is the default implementation for POST method.
func (c *RootController) Post() {
	var (
		request    = c.TemplateImpl.Request()
		response   = c.TemplateImpl.Response()
		errorResps []ErrorResponse
	)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, request); err != nil {
		errorResps = append(errorResps, *NewErrorResponseInvalidRequest())
		log.WithFields(log.Fields{
			"resource":  c.TemplateImpl.ResourceName(),
			"error":     err,
			"errorResp": errorResps[0].ID,
		}).Warn("RootController post resource failed, bad request.")
		c.Data["json"] = &errorResps
		c.Ctx.Output.SetStatus(errorResps[0].StatusCode)
		c.ServeJSON()
		return
	}

	if errorResp := request.IsValid(); errorResp != nil {
		errorResps = append(errorResps, *errorResp)
		log.WithFields(log.Fields{
			"resource":  c.TemplateImpl.ResourceName(),
			"errorResp": errorResps[0].ID,
		}).Warn("RootController post resource failed, request validation failed.")
		c.Data["json"] = &errorResps
		c.Ctx.Output.SetStatus(errorResps[0].StatusCode)
		c.ServeJSON()
		return
	}
	log.WithFields(log.Fields{
		"resource": c.TemplateImpl.ResourceName(),
		"request":  request,
	}).Debug("RootController post resource.")
	model, errorResps := c.TemplateImpl.Service().Create(request)
	if errorResps != nil {
		log.WithFields(log.Fields{
			"errorResp": errorResps[0].ID,
		}).Warn("RootController post resource failed, POST callback return errorResp.")
		c.Data["json"] = errorResps
		c.Ctx.Output.SetStatus(errorResps[0].StatusCode)
		c.ServeJSON()
		return
	}
	response.Load(model)
	log.WithFields(log.Fields{
		"resource": c.TemplateImpl.ResourceName(),
		"request":  request,
		"ID":       response.GetID(),
	}).Info("Post resource done.")
	c.Data["json"] = response
	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.ServeJSON()
}

// Get is the default implementation for GET method.
func (c *RootController) Get() {
	var (
		start, count, filter string = c.GetString("start"), c.GetString("count"), c.GetString("filter")
		startInt, countInt   int64  = 0, -1
		parameterError       bool
	)
	log.WithFields(log.Fields{
		"resource": c.TemplateImpl.ResourceName(),
		"start":    start,
		"count":    count,
	}).Debug("RootController get resource collection.")
	if start != "" {
		_startInt, err := strconv.ParseInt(start, 10, 64)
		if err != nil || _startInt < 0 {
			parameterError = true
		} else {
			startInt = _startInt
		}
	}
	if count != "" {
		_countInt, err := strconv.ParseInt(count, 10, 64)
		// -1 means all.
		if err != nil || _countInt < -1 {
			parameterError = true
		} else {
			countInt = _countInt
		}
	}

	if !c.isValidFilter(filter) {
		parameterError = true
	}

	if parameterError {
		errorResps := []ErrorResponse{*NewErrorResponseInvalidRequest()}
		log.WithFields(log.Fields{
			"resource":  c.TemplateImpl.ResourceName(),
			"errorResp": errorResps[0].ID,
		}).Warn("RootController get resource collection failed, parameter error.")
		c.Data["json"] = &errorResps
		c.Ctx.Output.SetStatus(errorResps[0].StatusCode)
		c.ServeJSON()
		return
	}
	collection, errorResps := c.TemplateImpl.Service().GetCollection(startInt, countInt, filter)
	if errorResps != nil {
		log.WithFields(log.Fields{
			"resource":  c.TemplateImpl.ResourceName(),
			"errorResp": errorResps[0].ID,
		}).Warn("RootController get resource collection failed.")
		c.Data["json"] = &errorResps
		c.Ctx.Output.SetStatus(errorResps[0].StatusCode)
		c.ServeJSON()
		return
	}
	response, err := c.TemplateImpl.ConvertCollectionModel(collection)
	if err != nil {
		errorResps := []ErrorResponse{*NewErrorResponseTransactionError()}
		log.WithFields(log.Fields{
			"resource":  c.TemplateImpl.ResourceName(),
			"errorResp": errorResps[0].ID,
			"error":     err,
		}).Warn("RootController convert resource collection response failed.")
		c.Data["json"] = &errorResps
		c.Ctx.Output.SetStatus(errorResps[0].StatusCode)
		c.ServeJSON()
		return
	}
	log.WithFields(log.Fields{
		"start":    startInt,
		"count":    countInt,
		"filter":   filter,
		"resource": c.TemplateImpl.ResourceName(),
	}).Info("Get resource collection done.")
	c.Data["json"] = response
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.ServeJSON()
}

// Delete is the default DELETE method on root controller.
func (c *RootController) Delete() {
	errorResps := c.TemplateImpl.Service().DeleteCollection()
	if errorResps != nil {
		c.Data["json"] = &errorResps
		c.Ctx.Output.SetStatus(errorResps[0].StatusCode)
		c.ServeJSON()
		return
	}
	c.Ctx.Output.SetStatus(http.StatusAccepted)
	log.WithFields(log.Fields{
		"resource": c.TemplateImpl.ResourceName(),
	}).Info("Delete resource collection done.")
	c.ServeJSON()
}

func (c *RootController) isValidFilter(filter string) bool {
	return true
}
