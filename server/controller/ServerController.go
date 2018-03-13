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
	var id = c.Ctx.Input.Param(":id")
	log.WithFields(log.Fields{"id": id}).Debug("Get server.")
	if server, messages := service.GetServer(id); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"id": id, "message": messages[0].ID}).Info("Get server failed.")
	} else {
		resp := dto.GetServerResponse{}
		resp.Load(server)
		c.Data["json"] = &resp
		c.Ctx.Output.SetStatus(http.StatusOK)
	}
	c.ServeJSON()
}
