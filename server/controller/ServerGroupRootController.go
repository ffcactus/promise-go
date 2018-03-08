package controller

import (
	commonDto "promise/common/object/dto"
	//	commonM "promise/common/object/model"
	dto "promise/server/object/dto"
	//	m "promise/server/object/model"

	"encoding/json"
	"promise/server/service"
	//	"strconv"

	"github.com/astaxie/beego"
)

// ServerGroupRootController The root controller
type ServerGroupRootController struct {
	beego.Controller
}

// Post a new server group.
func (c *ServerGroupRootController) Post() {
	var (
		request  dto.PostServerGroupRequest
		response dto.PostServerGroupResponse
	)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		beego.Warning("error: ", err)
	}
	// Create the context for this operation.
	serverGroup, messages := service.PostServerGroup(&request)
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		beego.Info("POST server group failed", messages[0].ID)
		// TODO Why the header not includes application/json?
	} else {
		response.Load(serverGroup)
		c.Data["json"] = &response
		beego.Info("POST server group", response.Name)
	}
	c.ServeJSON()

}
