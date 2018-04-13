package controller

import (
	"encoding/json"
	commonDto "promise/common/object/dto"
	commonMessage "promise/common/object/message"
	"promise/task/object/dto"
	"promise/task/object/message"
	"promise/task/service"
	"strings"

	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
)

var (
	// ActionUpdate Action enum.
	ActionUpdate = "update"
	// ActionUpdateTaskStep Action enum.
	ActionUpdateTaskStep = "updatetaskstep"
)

// TaskActionController Task action controller
type TaskActionController struct {
	beego.Controller
}

// Post POST method.
func (c *TaskActionController) Post() {
	var messages []commonMessage.Message
	action := c.Ctx.Input.Param(":action")
	id := c.Ctx.Input.Param(":id")
	log.Debug("Post() start, action = ", action, ", id = ", id)
	switch strings.ToLower(action) {
	case ActionUpdate:
		updateRequest := new(dto.UpdateTaskRequest)
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, updateRequest); err != nil {
			log.Warn("Unmarshal() failed, action = ", action, ", id = ", id, " error = ", err)
			messages = []commonMessage.Message{message.NewMessageTaskBadRequest()}
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.ResponseWriter.WriteHeader(messages[0].StatusCode)
		} else {
			if resp, messages := service.UpdateTask(id, updateRequest); messages != nil {
				c.Data["json"] = commonDto.MessagesToDto(messages)
				c.Ctx.ResponseWriter.WriteHeader(messages[0].StatusCode)
			} else {
				c.Data["json"] = &resp
			}
		}
	case ActionUpdateTaskStep:
		updateTaskStepRequest := new(dto.UpdateTaskStepRequest)
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, updateTaskStepRequest); err != nil {
			log.Warn("Unmarshal() failed, action = ", action, ", id = ", id, " error = ", err)
			messages = []commonMessage.Message{message.NewMessageTaskBadRequest()}
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.ResponseWriter.WriteHeader(messages[0].StatusCode)
		} else {
			if resp, messages := service.UpdateTaskStep(id, updateTaskStepRequest); messages != nil {
				c.Data["json"] = commonDto.MessagesToDto(messages)
				c.Ctx.ResponseWriter.WriteHeader(messages[0].StatusCode)
			} else {
				c.Data["json"] = &resp
			}
		}
	default:
		log.Info("Unknown task action ", action)
		messages := []commonMessage.Message{}
		messages = append(messages, message.NewMessageTaskBadRequest())
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.ResponseWriter.WriteHeader((messages)[0].StatusCode)
	}
	c.ServeJSON()
}
