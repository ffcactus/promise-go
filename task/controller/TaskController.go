package controller

import (
	commonDto "promise/common/object/dto"
	. "promise/task/object/dto"
	"promise/task/service"

	"github.com/astaxie/beego"
)

// TaskController Task controller
type TaskController struct {
	beego.Controller
}

// Get Get server by ID.
func (this *TaskController) Get() {
	beego.Trace("Get() start, ID = ", this.Ctx.Input.Param(":id"))
	if task, messages := service.GetTask(this.Ctx.Input.Param(":id")); messages != nil {
		this.Data["json"] = commonDto.MessagesToDto(messages)
		this.Ctx.ResponseWriter.WriteHeader(messages[0].StatusCode)
	} else {
		resp := new(PostTaskResponse)
		resp.Load(task)
		this.Data["json"] = resp
	}
	this.ServeJSON()
}
