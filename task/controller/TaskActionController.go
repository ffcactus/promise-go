package controller

import (
	"encoding/json"
	commonDto "promise/common/object/dto"
	. "promise/common/object/model"
	. "promise/task/object/dto"
	m "promise/task/object/model"
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
	var messages []Message
	action := c.Ctx.Input.Param(":action")
	id := c.Ctx.Input.Param(":id")
	log.Debug("Post() start, action = ", action, ", id = ", id)
	switch strings.ToLower(action) {
	case ActionUpdate:
		updateRequest := new(UpdateTaskRequest)
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, updateRequest); err != nil {
			log.Warn("Unmarshal() failed, action = ", action, ", id = ", id, " error = ", err)
			messages = []Message{m.NewMessageTaskBadRequest()}
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
		updateTaskStepRequest := new(UpdateTaskStepRequest)
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, updateTaskStepRequest); err != nil {
			log.Warn("Unmarshal() failed, action = ", action, ", id = ", id, " error = ", err)
			messages = []Message{m.NewMessageTaskBadRequest()}
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
		messages := []Message{}
		messages = append(messages, m.NewMessageTaskBadRequest())
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.ResponseWriter.WriteHeader((messages)[0].StatusCode)
	}
	c.ServeJSON()
}
