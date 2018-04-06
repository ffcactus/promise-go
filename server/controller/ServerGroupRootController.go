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

// ServerGroupRootController The root controller
type ServerGroupRootController struct {
	commonController.PromiseRootController
}

// Post a new servergroup.
func (c *ServerGroupRootController) Post() {
	var (
		request  dto.PostServerGroupRequest
		response dto.GetServerGroupResponse
		messages []commonMessage.Message
	)

	if message, err := c.PromiseRootController.Post(&request); message != nil {
		messages = append(messages, *message)
		log.WithFields(log.Fields{
			"error":   err,
			"message": messages[0].ID}).
			Warn("Post servergroup failed, bad request.")
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}

	log.WithFields(log.Fields{"name": request.Name}).Info("Post servergroup.")

	serverGroup, messages := service.PostServerGroup(&request)
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Post servergroup failed.")
	} else {
		response.Load(serverGroup)
		c.Data["json"] = &response
		c.Ctx.Output.SetStatus(http.StatusCreated)
		log.WithFields(log.Fields{"name": response.Name, "ID": response.ID}).Info("Post servergroup done.")
	}
	c.ServeJSON()
}

// Get will return servergroup collection.
func (c *ServerGroupRootController) Get() {
	var (
		messages []commonMessage.Message
		response dto.GetServerGroupCollectionResponse
	)

	start, count, filter, message, err := c.PromiseRootController.Get()
	if message != nil {
		messages = append(messages, *message)
		log.WithFields(log.Fields{
			"error":   err,
			"message": messages[0].ID}).
			Warn("Get servergroup collection failed, bad request.")
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}

	if collection, messages := service.GetServerGroupCollection(start, count, filter); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Get servergroup collection failed.")
	} else {
		response.Load(collection)
		c.Data["json"] = response
		c.Ctx.Output.SetStatus(http.StatusOK)
	}

	c.ServeJSON()
}

// Delete will delete all the group except default "all" group.
func (c *ServerGroupRootController) Delete() {
	messages := service.DeleteServerGroupCollection()
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Delete servergroup collection failed.")
	} else {
		c.Ctx.Output.SetStatus(http.StatusAccepted)
		log.Info("DELETE all servergroups.")
	}
	c.ServeJSON()
}

func (c *ServerGroupRootController) isValidFilter(filter string) bool {
	return true
}
