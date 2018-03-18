package controller

import (
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
	commonDto "promise/common/object/dto"
	commonM "promise/common/object/model"
	"promise/server/object/constvalue"
	"promise/server/object/message"
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
	case constvalue.ServerActionRefresh:
		if resp, messages := service.RefreshServer(id); messages != nil {
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.Output.SetStatus(messages[0].StatusCode)
			log.WithFields(log.Fields{"action": action, "id": id, "message": messages[0].ID}).Info("Cast action on server failed.")
		} else {
			c.Data["json"] = &resp
			c.Ctx.Output.SetStatus(http.StatusAccepted)
		}
	default:
		messages := []commonM.Message{}
		messages = append(messages, message.NewServerParameterError())
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
	}
	c.ServeJSON()
}
