package controller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
	commonDto "promise/common/object/dto"
	commonM "promise/common/object/model"
	"promise/server/object/dto"
	"promise/server/object/message"
	"promise/server/service"
	"strconv"
)

// ServerServerGroupRootController The root controller
type ServerServerGroupRootController struct {
	beego.Controller
}

// Post a new server-servergroup.
func (c *ServerServerGroupRootController) Post() {
	var (
		request  dto.PostServerServerGroupRequest
		response dto.PostServerServerGroupResponse
	)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		log.WithFields(log.Fields{"err": err}).Info("Post server-servergroup failed, unable to get request.")
		messages := []commonM.Message{}
		messages = append(messages, message.NewServerParameterError())
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	log.WithFields(log.Fields{"serverID": request.ServerID, "serverGroupID": request.ServerGroupID}).Info("Post server-servergroup.")

	serverServerGroup, messages := service.PostServerServerGroup(&request)
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Info("Post server-servergroup failed.")
	} else {
		response.Load(serverServerGroup)
		c.Data["json"] = &response
		c.Ctx.Output.SetStatus(http.StatusCreated)
		log.WithFields(log.Fields{"name": response.Name, "ID": response.ID}).Info("Post server-group done.")
	}
	c.ServeJSON()
}

// Get will return server-group collection.
func (c *ServerServerGroupRootController) Get() {
	var (
		start, count       string = c.GetString("start"), c.GetString("count")
		startInt, countInt int    = 0, -1
		parameterError     bool
	)
	log.WithFields(log.Fields{"start": start, "count": count}).Debug("Get server-group collection.")
	if start != "" {
		_startInt, err := strconv.Atoi(start)
		if err != nil || _startInt < 0 {
			parameterError = true
		} else {
			startInt = _startInt
		}
	}
	if count != "" {
		_countInt, err := strconv.Atoi(count)
		// -1 means all.
		if err != nil || _countInt < -1 {
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
		log.Warn("Get server-group collection failed, parameter error.")
	} else {
		if collection, messages := service.GetServerServerGroupCollection(startInt, countInt); messages != nil {
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.Output.SetStatus(messages[0].StatusCode)
			log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Get server-group collection failed")
		} else {
			resp := new(dto.GetServerServerGroupCollectionResponse)
			resp.Load(collection)
			c.Data["json"] = resp
			c.Ctx.Output.SetStatus(http.StatusOK)
		}
	}
	c.ServeJSON()
}

// Delete will delete all the server-servergroup.
func (c *ServerServerGroupRootController) Delete() {
	messages := service.DeleteServerServerGroupCollection()
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
	}
	c.Ctx.Output.SetStatus(http.StatusAccepted)
	log.Info("DELETE all server-group")
	c.ServeJSON()
}