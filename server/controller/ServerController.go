package controller

import (
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
	commonDto "promise/common/object/dto"
	dto "promise/server/object/dto"
	"promise/server/service"
)

// ServerController Server controller.
type ServerController struct {
	beego.Controller
}

// Get Get server by ID.
func (c *ServerController) Get() {
	log.Debug("Get(), server ID = ", c.Ctx.Input.Param(":id"))
	if server, messages := service.GetServer(c.Ctx.Input.Param(":id")); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
	} else {
		resp := dto.GetServerResponse{}
		resp.Load(server)
		c.Data["json"] = &resp
		c.Ctx.Output.SetStatus(http.StatusOK)
	}
	c.ServeJSON()
}
