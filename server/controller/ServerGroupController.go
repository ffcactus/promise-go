package controller

import (
	commonDto "promise/common/object/dto"
	dto "promise/server/object/dto"
	"promise/server/service"

	"github.com/astaxie/beego"
)

// ServerGroupController Server controller.
type ServerGroupController struct {
	beego.Controller
}

// Get will return the server group by ID.
func (c *ServerGroupController) Get() {
	var resp dto.GetServerGroupResponse
	beego.Trace("GET server group ID = ", c.Ctx.Input.Param(":id"))
	if sg, messages := service.GetServerGroup(c.Ctx.Input.Param(":id")); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
	} else {
		resp.Load(sg)
		c.Data["json"] = &resp
	}
	c.ServeJSON()
}
