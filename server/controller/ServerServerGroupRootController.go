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
		log.WithFields(log.Fields{"name": response.Name, "ID": response.ID}).Info("Post server-servergroup done.")
	}
	c.ServeJSON()
}

// Get will return server-servergroup collection.
func (c *ServerServerGroupRootController) Get() {
	var (
		start, count, filter string = c.GetString("start"), c.GetString("count"), c.GetString("$filter")
		startInt, countInt   int    = 0, -1
		parameterError       bool
	)
	log.WithFields(log.Fields{"start": start, "count": count}).Debug("Get server-servergroup collection.")
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

	if !c.isValidFilter(filter) {
		parameterError = true
	}

	if parameterError {
		messages := []commonM.Message{}
		messages = append(messages, message.NewServerParameterError())
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.Warn("Get server-servergroup collection failed, parameter error.")
	} else {
		if collection, messages := service.GetServerServerGroupCollection(startInt, countInt, filter); messages != nil {
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.Output.SetStatus(messages[0].StatusCode)
			log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Get server-servergroup collection failed.")
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
		log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Delete server-servergroup collection failed.")
	} else {
		c.Ctx.Output.SetStatus(http.StatusAccepted)
		log.Info("DELETE all server-servergroups.")
	}
	c.ServeJSON()
}

func (c *ServerServerGroupRootController) isValidFilter(filter string) bool {
	return true
}
