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
	var id = c.Ctx.Input.Param(":id")
	log.WithFields(log.Fields{"id": id}).Debug("Get server group.")
	if sg, messages := service.GetServerGroup(id); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"id": id, "message": messages[0].ID}).Debug("Get server group failed.")
	} else {
		resp.Load(sg)
		c.Data["json"] = &resp
		c.Ctx.Output.SetStatus(http.StatusOK)
	}
	c.ServeJSON()
}

// Delete will delete the server group by ID.
func (c *ServerGroupController) Delete() {
	var id = c.Ctx.Input.Param(":id")
	log.WithFields(log.Fields{"id": id}).Debug("Delete server group.")
	if messages := service.DeleteServerGroup(id); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"id": id, "message": messages[0].ID}).Debug("Delete server group failed.")
	}
	c.Ctx.Output.SetStatus(http.StatusAccepted)
	
	c.ServeJSON()
}
