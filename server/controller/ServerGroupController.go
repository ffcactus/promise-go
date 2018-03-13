package controller

import (
	"net/http"
	commonDto "promise/common/object/dto"
	dto "promise/server/object/dto"
	"promise/server/service"

	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
)

// ServerGroupController Server controller.
type ServerGroupController struct {
	beego.Controller
}

// Get will return the server group by ID.
func (c *ServerGroupController) Get() {
	var resp dto.GetServerGroupResponse
	log.Debug("GET server group ID = ", c.Ctx.Input.Param(":id"))
	if sg, messages := service.GetServerGroup(c.Ctx.Input.Param(":id")); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
	} else {
		resp.Load(sg)
		c.Data["json"] = &resp
		c.Ctx.Output.SetStatus(http.StatusOK)
	}
	c.ServeJSON()
}

// Delete will delete the server group by ID.
func (c *ServerGroupController) Delete() {
	if messages := service.DeleteServerGroup(c.Ctx.Input.Param(":id")); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
	}
	c.Ctx.Output.SetStatus(http.StatusAccepted)
	log.Info("DELETE server group ", c.Ctx.Input.Param(":id"))
	c.ServeJSON()
}
