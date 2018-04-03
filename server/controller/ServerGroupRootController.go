package controller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
	commonDto "promise/common/object/dto"
	commonMessage "promise/common/object/message"
	"promise/server/object/dto"
	"promise/server/service"
	"strconv"
)

// ServerGroupRootController The root controller
type ServerGroupRootController struct {
	beego.Controller
}

// Post a new servergroup.
func (c *ServerGroupRootController) Post() {
	var (
		request  dto.PostServerGroupRequest
		response dto.GetServerGroupResponse
	)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("Post servergroup failed, unable to unmarshal request.")
		messages := []commonMessage.Message{}
		messages = append(messages, commonMessage.NewInvalidRequest())
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	log.WithFields(log.Fields{"name": request.Name}).Info("Post servergroup.")

	serverGroup, messages := service.PostServerGroup(&request)
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Post servergroup failed.")
	} else {
		response.Load(serverGroup)
		c.Data["json"] = &response
		c.Ctx.Output.SetStatus(http.StatusCreated)
		log.WithFields(log.Fields{"name": response.Name, "ID": response.ID}).Info("Post servergroup done.")
	}
	c.ServeJSON()
}

// Get will return servergroup collection.
func (c *ServerGroupRootController) Get() {
	var (
		start, count, filter string = c.GetString("start"), c.GetString("count"), c.GetString("$filter")
		startInt, countInt   int    = 0, -1
		parameterError       bool
	)
	log.WithFields(log.Fields{"start": start, "count": count}).Debug("Get servergroup collection.")
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
		messages := []commonMessage.Message{}
		messages = append(messages, commonMessage.NewInvalidRequest())
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.Warn("Get servergroup collection failed, parameter error.")
	} else {
		if collection, messages := service.GetServerGroupCollection(startInt, countInt, filter); messages != nil {
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.Output.SetStatus(messages[0].StatusCode)
			log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Get servergroup collection failed.")
		} else {
			resp := new(dto.GetServerGroupCollectionResponse)
			resp.Load(collection)
			c.Data["json"] = resp
			c.Ctx.Output.SetStatus(http.StatusOK)
		}
	}
	c.ServeJSON()
}

// Delete will delete all the group except default "all" group.
func (c *ServerGroupRootController) Delete() {
	messages := service.DeleteServerGroupCollection()
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Delete servergroup collection failed.")
	} else {
		c.Ctx.Output.SetStatus(http.StatusAccepted)
		log.Info("DELETE all servergroups.")
	}
	c.ServeJSON()
}

func (c *ServerGroupRootController) isValidFilter(filter string) bool {
	return true
}
