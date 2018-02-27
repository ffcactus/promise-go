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
func (this *TaskActionController) Post() {
	var messages []Message
	action := this.Ctx.Input.Param(":action")
	id := this.Ctx.Input.Param(":id")
	beego.Trace("Post() start, action = ", action, ", id = ", id)
	switch strings.ToLower(action) {
	case ActionUpdate:
		updateRequest := new(UpdateTaskRequest)
		if err := json.Unmarshal(this.Ctx.Input.RequestBody, updateRequest); err != nil {
			beego.Warning("Unmarshal() failed, action = ", action, ", id = ", id, " error = ", err)
			messages = []Message{m.NewMessageTaskBadRequest()}
			this.Data["json"] = commonDto.MessagesToDto(messages)
			this.Ctx.ResponseWriter.WriteHeader(messages[0].StatusCode)
		} else {
			if resp, messages := service.UpdateTask(id, updateRequest); messages != nil {
				this.Data["json"] = commonDto.MessagesToDto(messages)
				this.Ctx.ResponseWriter.WriteHeader(messages[0].StatusCode)
			} else {
				this.Data["json"] = &resp
			}
		}
	case ActionUpdateTaskStep:
		updateTaskStepRequest := new(UpdateTaskStepRequest)
		if err := json.Unmarshal(this.Ctx.Input.RequestBody, updateTaskStepRequest); err != nil {
			beego.Warning("Unmarshal() failed, action = ", action, ", id = ", id, " error = ", err)
			messages = []Message{m.NewMessageTaskBadRequest()}
			this.Data["json"] = commonDto.MessagesToDto(messages)
			this.Ctx.ResponseWriter.WriteHeader(messages[0].StatusCode)
		} else {
			if resp, messages := service.UpdateTaskStep(id, updateTaskStepRequest); messages != nil {
				this.Data["json"] = commonDto.MessagesToDto(messages)
				this.Ctx.ResponseWriter.WriteHeader(messages[0].StatusCode)
			} else {
				this.Data["json"] = &resp
			}
		}
	default:
		beego.Info("Unknown task action ", action)
		messages := []Message{}
		messages = append(messages, m.NewMessageTaskBadRequest())
		this.Data["json"] = commonDto.MessagesToDto(messages)
		this.Ctx.ResponseWriter.WriteHeader((messages)[0].StatusCode)
	}
	this.ServeJSON()
}
