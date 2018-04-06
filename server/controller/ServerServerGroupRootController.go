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

// ServerServerGroupRootController The root controller
type ServerServerGroupRootController struct {
	commonController.PromiseRootController
}

// Post a new server-servergroup.
func (c *ServerServerGroupRootController) Post() {
	var (
		request  dto.PostServerServerGroupRequest
		response dto.GetServerServerGroupResponse
		messages []commonMessage.Message
	)

	if message, err := c.PromiseRootController.Post(&request); message != nil {
		messages = append(messages, *message)
		log.WithFields(log.Fields{
			"error":   err,
			"message": messages[0].ID}).
			Warn("Post server-servergroup failed, bad request.")
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}

	log.WithFields(log.Fields{"serverID": request.ServerID, "serverGroupID": request.ServerGroupID}).Info("Post server-servergroup.")

	serverServerGroup, messages := service.PostServerServerGroup(&request)
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Post server-servergroup failed.")
	} else {
		response.Load(serverServerGroup)
		c.Data["json"] = &response
		c.Ctx.Output.SetStatus(http.StatusCreated)
		log.WithFields(log.Fields{"ID": response.ID}).Info("Post server-servergroup done.")
	}
	c.ServeJSON()
}

// Get will return server-servergroup collection.
func (c *ServerServerGroupRootController) Get() {
	var (
		messages []commonMessage.Message
		response dto.GetServerServerGroupCollectionResponse
	)

	start, count, filter, message, err := c.PromiseRootController.Get()
	if message != nil {
		messages = append(messages, *message)
		log.WithFields(log.Fields{
			"error":   err,
			"message": messages[0].ID}).
			Warn("Get server-servergroup collection failed, bad request.")
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}

	if collection, messages := service.GetServerServerGroupCollection(start, count, filter); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Get server-servergroup collection failed.")
	} else {
		response.Load(collection)
		c.Data["json"] = &response
		c.Ctx.Output.SetStatus(http.StatusOK)
	}

	c.ServeJSON()
}

// Delete will delete all the server-servergroup.
func (c *ServerServerGroupRootController) Delete() {
	messages := service.DeleteServerServerGroupCollection()
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Delete server-servergroup collection failed.")
	} else {
		c.Ctx.Output.SetStatus(http.StatusAccepted)
		log.Info("DELETE all server-servergroups.")
	}
	c.ServeJSON()
}

func (c *ServerServerGroupRootController) isValidFilter(filter string) bool {
	return true
}
