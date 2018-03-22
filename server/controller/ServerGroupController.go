package controller

import (
	"net/http"
	commonDto "promise/common/object/dto"
	dto "promise/server/object/dto"
	"promise/server/service"

	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
)

// ServerGroupController is the servergroup controller.
type ServerGroupController struct {
	beego.Controller
}

// Get will return the group by ID.
func (c *ServerGroupController) Get() {
	var resp dto.GetServerGroupResponse
	var id = c.Ctx.Input.Param(":id")
	log.WithFields(log.Fields{"id": id}).Debug("Get servergroup.")
	if sg, messages := service.GetServerGroup(id); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"id": id, "message": messages[0].ID}).Warn("Get servergroup failed.")
	} else {
		resp.Load(sg)
		c.Data["json"] = &resp
		c.Ctx.Output.SetStatus(http.StatusOK)
	}
	c.ServeJSON()
}

// Delete will delete the servergroup by ID.
func (c *ServerGroupController) Delete() {
	var id = c.Ctx.Input.Param(":id")
	if messages := service.DeleteServerGroup(id); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"id": id, "message": messages[0].ID}).Warn("Delete servergroup failed.")
	} else {
		c.Ctx.Output.SetStatus(http.StatusAccepted)
		log.WithFields(log.Fields{"id": id}).Info("Delete servergroup done.")
	}
	c.ServeJSON()
}
