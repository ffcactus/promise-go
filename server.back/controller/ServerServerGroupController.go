package controller

import (
	"net/http"
	commonDto "promise/common/object/dto"
	dto "promise/server/object/dto"
	"promise/server/service"

	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
)

// ServerServerGroupController is the server-servergroup controller.
type ServerServerGroupController struct {
	beego.Controller
}

// Get will return the server-servergroup association by ID.
func (c *ServerServerGroupController) Get() {
	var resp dto.GetServerServerGroupResponse
	var id = c.Ctx.Input.Param(":id")
	log.WithFields(log.Fields{"id": id}).Debug("Get server-servergroup.")
	if sg, messages := service.GetServerServerGroup(id); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"id": id, "message": messages[0].ID}).Warn("Get server-servergroup failed.")
	} else {
		resp.Load(sg)
		c.Data["json"] = &resp
		c.Ctx.Output.SetStatus(http.StatusOK)
	}
	c.ServeJSON()
}

// Delete will delete the server-servergroup by ID.
func (c *ServerServerGroupController) Delete() {
	var id = c.Ctx.Input.Param(":id")
	log.WithFields(log.Fields{"id": id}).Debug("Delete server-servergroup.")
	if messages := service.DeleteServerServerGroup(id); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"id": id, "message": messages[0].ID}).Warn("Delete server-servergroup failed.")
	} else {
		c.Ctx.Output.SetStatus(http.StatusAccepted)
	}
	c.ServeJSON()
}
