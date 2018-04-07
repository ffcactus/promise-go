package controller

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	commonController "promise/common/controller"
	commonDto "promise/common/object/dto"
	commonMessage "promise/common/object/message"
	"promise/task/object/dto"
	"promise/task/service"
)

// TaskRootController is the root controller for task.
type TaskRootController struct {
	commonController.PromiseRootController
}

// Post Post a new task.
func (c *TaskRootController) Post() {
	var (
		request  dto.PostTaskRequest
		response dto.GetTaskResponse
		messages []commonMessage.Message
	)

	if message, err := c.PromiseRootController.Post(&request); message != nil {
		messages = append(messages, *message)
		log.WithFields(log.Fields{
			"error":   err,
			"message": messages[0].ID}).
			Warn("Post task failed, bad request.")
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}

	log.WithFields(log.Fields{"name": request.Name}).Info("Post IPv4 pool.")

	// Create the context for this operation.
	if task, messages := service.PostTask(&request); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Post task failed.")
	} else {
		response.Load(task)
		c.Data["json"] = &response
		c.Ctx.Output.SetStatus(http.StatusCreated)
		log.WithFields(log.Fields{"name": response.Name, "ID": response.ID}).Info("Post IPv4 pool done.")
	}
	c.ServeJSON()
}

// Get Get task collection.
func (c *TaskRootController) Get() {
	var (
		messages []commonMessage.Message
		response dto.GetTaskCollectionResponse
	)

	start, count, filter, message, err := c.PromiseRootController.Get()
	if message != nil {
		messages = append(messages, *message)
		log.WithFields(log.Fields{
			"error":   err,
			"message": messages[0].ID}).
			Warn("Get task collection failed, bad request.")
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}

	if collection, messages := service.GetTaskCollection(start, count, filter); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.ResponseWriter.WriteHeader(messages[0].StatusCode)
	} else {
		response.Load(collection)
		c.Data["json"] = &response
	}

	c.ServeJSON()
}
