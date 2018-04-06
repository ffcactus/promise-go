package controller

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	commonController "promise/common/controller"
	commonDto "promise/common/object/dto"
	commonMessage "promise/common/object/message"
	"promise/pool/object/dto"
	"promise/pool/service"
)

// IPv4RootController is the ipv4 pool controller.
type IPv4RootController struct {
	commonController.PromiseRootController
}

// Post a new IPv4 range.
func (c *IPv4RootController) Post() {
	var (
		request  dto.PostIPv4PoolRequest
		response dto.GetIPv4PoolResponse
		messages []commonMessage.Message
	)

	if message, err := c.PromiseRootController.Post(&request); message != nil {
		messages = append(messages, *message)
		log.WithFields(log.Fields{
			"error":   err,
			"message": messages[0].ID}).
			Warn("Post IPv4 pool failed, bad request.")
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}

	log.WithFields(log.Fields{"name": request.Name}).Info("Post IPv4 pool.")

	ipv4Pool, messages := service.PostIPv4Pool(&request)
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Post IPv4 pool failed.")
	} else {
		response.Load(ipv4Pool)
		c.Data["json"] = &response
		c.Ctx.Output.SetStatus(http.StatusCreated)
		log.WithFields(log.Fields{"name": response.Name, "ID": response.ID}).Info("Post IPv4 pool done.")
	}
	c.ServeJSON()
}

// Get will return IPv4 pool collection.
func (c *IPv4RootController) Get() {
	var (
		messages             []commonMessage.Message
	)

	start, count, filter, message, err := c.PromiseRootController.Get()
	if message != nil {
		messages = append(messages, *message)
		log.WithFields(log.Fields{
			"error":   err,
			"message": messages[0].ID}).
			Warn("Get IPv4 pool collection failed, bad request.")
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}

	if collection, messages := service.GetIPv4PoolCollection(start, count, filter); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Get IPv4 pool collection failed.")
	} else {
		resp := new(dto.GetIPv4PoolCollectionResponse)
		resp.Load(collection)
		c.Data["json"] = resp
		c.Ctx.Output.SetStatus(http.StatusOK)
	}

	c.ServeJSON()
}

// Delete will delete all the IPv4 pool.
func (c *IPv4RootController) Delete() {
	messages := service.DeleteIPv4PoolCollection()
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Delete IPv4 pool collection failed.")
	} else {
		c.Ctx.Output.SetStatus(http.StatusAccepted)
		log.Info("DELETE all IPv4 pool.")
	}
	c.ServeJSON()
}

func (c *IPv4RootController) isValidFilter(filter string) bool {
	return true
}
