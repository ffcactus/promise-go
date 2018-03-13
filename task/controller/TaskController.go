package controller

import (
	commonDto "promise/common/object/dto"
	. "promise/task/object/dto"
	"promise/task/service"

	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
)

// TaskController Task controller
type TaskController struct {
	beego.Controller
}

// Get Get server by ID.
func (c *TaskController) Get() {
	log.Debug("Get() start, ID = ", c.Ctx.Input.Param(":id"))
	if task, messages := service.GetTask(c.Ctx.Input.Param(":id")); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.ResponseWriter.WriteHeader(messages[0].StatusCode)
	} else {
		resp := new(PostTaskResponse)
		resp.Load(task)
		c.Data["json"] = resp
	}
	c.ServeJSON()
}
