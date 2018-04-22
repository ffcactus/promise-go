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
		request  = c.TemplateImpl.Request()
		response = c.TemplateImpl.Response()
		messages []Message
	)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, request); err != nil {
		messages = append(messages, NewMessageInvalidRequest())
		log.WithFields(log.Fields{
			"resource": c.TemplateImpl.ResourceName(),
			"error":    err,
			"message":  messages[0].ID,
		}).Warn("Post resource failed, bad request.")
		c.Data["json"] = &messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}

	if message := request.IsValid(); message != nil {
		messages = append(messages, *message)
		log.WithFields(log.Fields{
			"resource": c.TemplateImpl.ResourceName(),
			"message":  messages[0].ID,
		}).Warn("Post resource failed, request validation failed.")
		c.Data["json"] = &messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	log.WithFields(log.Fields{
		"resource": c.TemplateImpl.ResourceName(),
		"request":  request.DebugInfo(),
	}).Info("Post resource.")
	model, messages := c.TemplateImpl.Service().Create(request)
	if messages != nil {
		log.WithFields(log.Fields{
			"message": messages[0].ID,
		}).Warn("Post resource failed, POST callback return message.")
		c.Data["json"] = messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	response.Load(model)
	log.WithFields(log.Fields{
		"resource": c.TemplateImpl.ResourceName(),
		"request":  request.DebugInfo(),
		"ID":       response.GetID(),
	}).Info("Post resource done.")
	c.Data["json"] = response
	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.ServeJSON()
}

// Get is the default implementation for GET method.
func (c *RootController) Get() {
	var (
		start, count, filter string = c.GetString("start"), c.GetString("count"), c.GetString("$filter")
		startInt, countInt   int64  = 0, -1
		parameterError       bool
	)
	log.WithFields(log.Fields{
		"resource": c.TemplateImpl.ResourceName(),
		"start":    start,
		"count":    count,
	}).Info("Get resource collection.")
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
		messages := []Message{NewMessageInvalidRequest()}
		log.WithFields(log.Fields{
			"resource": c.TemplateImpl.ResourceName(),
			"message":  messages[0].ID,
		}).Warn("Get resource collection failed, parameter error.")
		c.Data["json"] = &messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	collection, messages := c.TemplateImpl.Service().GetCollection(startInt, countInt, filter)
	if messages != nil {
		log.WithFields(log.Fields{
			"resource": c.TemplateImpl.ResourceName(),
			"message":  messages[0].ID,
		}).Warn("Get resource collection failed.")
		c.Data["json"] = &messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	response, err := c.TemplateImpl.ConvertCollectionModel(collection)
	if err != nil {
		messages := []Message{NewMessageTransactionError()}
		log.WithFields(log.Fields{
			"resource": c.TemplateImpl.ResourceName(),
			"message":  messages[0].ID,
			"error":    err,
		}).Warn("Convert resource collection response failed.")
		c.Data["json"] = &messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
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
	messages := c.TemplateImpl.Service().DeleteCollection()
	if messages != nil {
		c.Data["json"] = &messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	c.Ctx.Output.SetStatus(http.StatusAccepted)
	log.WithFields(log.Fields{
		"resource": c.TemplateImpl.ResourceName(),
	}).Info("DELETE resource collection done.")
	c.ServeJSON()
}

func (c *RootController) isValidFilter(filter string) bool {
	return true
}
