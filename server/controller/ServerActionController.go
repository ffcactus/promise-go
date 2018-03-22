package controller

import (
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
	commonDto "promise/common/object/dto"
	commomMessage "promise/common/object/message"
	"promise/server/object/constValue"
	"promise/server/service"
	"strings"
)

// ServerActionController Server action controller
type ServerActionController struct {
	beego.Controller
}

// Post Post method.
func (c *ServerActionController) Post() {
	action := c.Ctx.Input.Param(":action")
	id := c.Ctx.Input.Param(":id")
	log.WithFields(log.Fields{"action": action, "id": id}).Info("Cast action on server.")
	switch strings.ToLower(action) {
	case constValue.ServerActionRefresh:
		if resp, messages := service.RefreshServer(id); messages != nil {
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.Output.SetStatus(messages[0].StatusCode)
			log.WithFields(log.Fields{"action": action, "id": id, "message": messages[0].ID}).Warn("Cast action on server failed.")
		} else {
			c.Data["json"] = &resp
			c.Ctx.Output.SetStatus(http.StatusAccepted)
		}
	default:
		messages := []commomMessage.Message{}
		messages = append(messages, commomMessage.NewInvalidRequest())
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
	}
	c.ServeJSON()
}
