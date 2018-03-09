package controller

import (
	commonDto "promise/common/object/dto"
	commonM "promise/common/object/model"
	"promise/server/object/message"
	"promise/server/service"
	"promise/server/util"
	"strings"

	"github.com/astaxie/beego"
)

// ServerActionController Server action controller
type ServerActionController struct {
	beego.Controller
}

// Post Post method.
func (c *ServerActionController) Post() {
	action := c.Ctx.Input.Param(":action")
	id := c.Ctx.Input.Param(":id")
	beego.Trace("Post(), action = ", action, ", server ID = ", action, id)
	switch strings.ToLower(action) {
	case util.ServerActionRefresh:
		if resp, messages := service.RefreshServer(id); messages != nil {
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.Output.SetStatus(messages[0].StatusCode)
		} else {
			c.Data["json"] = &resp
		}
	default:
		messages := []commonM.Message{}
		messages = append(messages, message.NewServerParameterError())
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
	}
	c.ServeJSON()
}
