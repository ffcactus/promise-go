package controller

import (
	commonDto "promise/common/object/dto"
	dto "promise/server/object/dto"
	"promise/server/service"

	"github.com/astaxie/beego"
)

// ServerController Server controller.
type ServerController struct {
	beego.Controller
}

// Get Get server by ID.
func (c *ServerController) Get() {
	beego.Trace("Get(), server ID = ", c.Ctx.Input.Param(":id"))
	if server, messages := service.GetServer(c.Ctx.Input.Param(":id")); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.ResponseWriter.WriteHeader(messages[0].StatusCode)
	} else {
		resp := dto.GetServerResponse{}
		resp.Load(server)
		c.Data["json"] = &resp
	}
	c.ServeJSON()
}
