package controller

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	commonController "promise/common/controller"
	commonDto "promise/common/object/dto"
	commonMessage "promise/common/object/message"
	"promise/server/object/dto"
	"promise/server/service"
)

// ServerRootController The root controller
type ServerRootController struct {
	commonController.PromiseRootController
}

// Post Post a new server.
func (c *ServerRootController) Post() {
	var (
		request  dto.PostServerRequest
		response dto.GetServerResponse
		messages []commonMessage.Message
	)
	if message, err := c.PromiseRootController.Post(&request); message != nil {
		messages = append(messages, *message)
		log.WithFields(log.Fields{
			"error":   err,
			"message": messages[0].ID}).
			Warn("Post server failed, bad request.")
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}

	log.WithFields(log.Fields{"hostname": request.Hostname}).Info("Post server start.")
	// Create the context for this operation.
	server, messages := service.PostServer(&request)
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Post server failed.")
	} else {
		response.Load(server)
		c.Data["json"] = &response
		c.Ctx.Output.SetStatus(http.StatusCreated)
		log.WithFields(log.Fields{"name": request.Hostname, "ID": response.ID}).Info("Post server done.")
	}
	c.ServeJSON()
}

// Get Get server collection.
func (c *ServerRootController) Get() {
	var (
		messages []commonMessage.Message
		response dto.GetServerCollectionResponse
	)

	start, count, filter, message, err := c.PromiseRootController.Get()
	if message != nil {
		messages = append(messages, *message)
		log.WithFields(log.Fields{
			"error":   err,
			"message": messages[0].ID}).
			Warn("Get server collection failed, bad request.")
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}

	if collection, messages := service.GetServerCollection(start, count, filter); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Get server collection failed")
	} else {
		response.Load(collection)
		c.Data["json"] = &response
		c.Ctx.Output.SetStatus(http.StatusOK)
	}

	c.ServeJSON()
}

// Delete will delete all servers.
func (c *ServerRootController) Delete() {
	messages := service.DeleteServerCollection()
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Delete server collection failed")
	} else {
		c.Ctx.Output.SetStatus(http.StatusAccepted)
		log.Info("DELETE all servers done.")
	}
	c.ServeJSON()
}
