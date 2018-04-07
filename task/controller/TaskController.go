package controller

import (
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
	commonDto "promise/common/object/dto"
	"promise/task/object/dto"
	"promise/task/service"
)

// TaskController Task controller
type TaskController struct {
	beego.Controller
}

// Get Get server by ID.
func (c *TaskController) Get() {
	var (
		id       = c.Ctx.Input.Param(":id")
		response dto.GetTaskResponse
	)
	log.WithFields(log.Fields{"id": id}).Debug("Get task.")
	if task, messages := service.GetTask(c.Ctx.Input.Param(":id")); messages != nil {
		log.WithFields(log.Fields{"id": id, "message": messages[0].ID}).Warn("Get task failed.")
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.ResponseWriter.WriteHeader(messages[0].StatusCode)
		c.Ctx.Output.SetStatus(http.StatusOK)
	} else {
		response.Load(task)
		c.Data["json"] = &response
		c.Ctx.Output.SetStatus(http.StatusOK)
	}
	c.ServeJSON()
}
