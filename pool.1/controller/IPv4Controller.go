package controller

import (
	"net/http"
	commonDto "promise/common/object/dto"
	dto "promise/pool/object/dto"
	"promise/pool/service"

	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
)

// IPv4Controller is the IPv4 pool controller.
type IPv4Controller struct {
	beego.Controller
}

// Get will return the IPv4 pool by ID.
func (c *IPv4Controller) Get() {
	var resp dto.GetIPv4PoolResponse
	var id = c.Ctx.Input.Param(":id")
	log.WithFields(log.Fields{"id": id}).Debug("Get IPv4 pool.")
	if sg, messages := service.GetIPv4Pool(id); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"id": id, "message": messages[0].ID}).Warn("Get IPv4 pool failed.")
	} else {
		resp.Load(sg)
		c.Data["json"] = &resp
		c.Ctx.Output.SetStatus(http.StatusOK)
	}
	c.ServeJSON()
}

// Delete will delete the IPv4 pool by ID.
func (c *IPv4Controller) Delete() {
	var id = c.Ctx.Input.Param(":id")
	if messages := service.DeleteIPv4Pool(id); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"id": id, "message": messages[0].ID}).Warn("Delete IPv4 pool failed.")
	} else {
		c.Ctx.Output.SetStatus(http.StatusAccepted)
		log.WithFields(log.Fields{"id": id}).Info("Delete IPv4 pool done.")
	}
	c.ServeJSON()
}
