package controller

import (
	commonDto "promise/common/object/dto"
	commonM "promise/common/object/model"
	dto "promise/server/object/dto"
	m "promise/server/object/model"

	"encoding/json"
	"promise/server/service"
	"strconv"

	"github.com/astaxie/beego"
)

// RootController The root controller
type RootController struct {
	beego.Controller
}

// Post Post a new server.
func (c *RootController) Post() {
	request := new(dto.PostServerRequest)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, request); err != nil {
		beego.Warning("error: ", err)
	}
	beego.Info("Post() start, address = " + request.Address)
	// Create the context for this operation.
	server, messages := service.PostServer(request)
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		// TODO Why the header not includes application/json?
	} else {
		resp := dto.PostServerResponse{}
		resp.Load(server)
		c.Data["json"] = &resp
	}
	c.ServeJSON()
	beego.Info("Post() done, server ID = ", server.ID)
}

// Get Get server collection.
func (c *RootController) Get() {
	var (
		start, count       string = c.GetString("start"), c.GetString("count")
		startInt, countInt int    = 0, -1
		parameterError     bool
	)
	beego.Trace("Get server collection, start = ", start, ", count = ", count)
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
		messages = append(messages, m.NewServerParameterError())
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.ResponseWriter.WriteHeader((messages)[0].StatusCode)
	} else {
		if serverCollection, messages := service.GetServerCollection(startInt, countInt); messages != nil {
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.ResponseWriter.WriteHeader(messages[0].StatusCode)
		} else {
			resp := new(dto.GetServerCollectionResponse)
			resp.Load(serverCollection)
			c.Data["json"] = resp
		}
	}
	c.ServeJSON()
}
