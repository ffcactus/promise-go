package controller

import (
	"encoding/json"
	"net/http"
	commonDto "promise/common/object/dto"
	commonM "promise/common/object/model"
	"promise/server/object/dto"
	"promise/server/object/message"
	"promise/server/service"
	"strconv"

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
	beego.Info("POST server group", response.Name)
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
		c.Ctx.Output.SetStatus(http.StatusCreated)
	}
	c.ServeJSON()
}

// Get will return server group collection.
func (c *ServerGroupRootController) Get() {
	var (
		start, count       string = c.GetString("start"), c.GetString("count")
		startInt, countInt int    = 0, -1
		parameterError     bool
	)
	beego.Trace("Get server group collection, start = ", start, ", count = ", count)
	if start != "" {
		_startInt, err := strconv.Atoi(start)
		if err != nil || _startInt < 0 {
			beego.Warning("Get(), invalid 'start' parameter, error = ", err)
			parameterError = true
		} else {
			startInt = _startInt
		}
	}
	if count != "" {
		_countInt, err := strconv.Atoi(count)
		// -1 means all.
		if err != nil || _countInt < -1 {
			beego.Warning("Get() 'count' parameter error = %s\n", err)
			parameterError = true
		} else {
			countInt = _countInt
		}
	}

	if parameterError {
		messages := []commonM.Message{}
		messages = append(messages, message.NewServerParameterError())
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
	} else {
		if serverCollection, messages := service.GetServerGroupCollection(startInt, countInt); messages != nil {
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.Output.SetStatus(messages[0].StatusCode)
		} else {
			resp := new(dto.GetServerGroupCollectionResponse)
			resp.Load(serverCollection)
			c.Data["json"] = resp
			c.Ctx.Output.SetStatus(http.StatusOK)
		}
	}
	c.ServeJSON()
}
