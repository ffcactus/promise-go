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

// GroupRootController The root controller
type GroupRootController struct {
	beego.Controller
}

// Post a new server group.
func (c *GroupRootController) Post() {
	var (
		request  dto.PostGroupRequest
		response dto.PostGroupResponse
	)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		log.WithFields(log.Fields{"err": err}).Info("Post server group failed, unable to get request.")
		messages := []commonM.Message{}
		messages = append(messages, message.NewServerParameterError())
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	log.WithFields(log.Fields{"name": request.Name}).Info("Post server group.")
	// Create the context for this operation.
	serverGroup, messages := service.PostGroup(&request)
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Info("Post server group failed.")
	} else {
		response.Load(serverGroup)
		c.Data["json"] = &response
		c.Ctx.Output.SetStatus(http.StatusCreated)
		log.WithFields(log.Fields{"name": response.Name, "ID": response.ID}).Info("Post server group done.")
	}
	c.ServeJSON()
}

// Get will return server group collection.
func (c *GroupRootController) Get() {
	var (
		start, count       string = c.GetString("start"), c.GetString("count")
		startInt, countInt int    = 0, -1
		parameterError     bool
	)
	log.WithFields(log.Fields{"start": start, "count": count}).Debug("Get server group collection.")
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
		log.Warn("Get server group collection failed, parameter error.")
	} else {
		if serverCollection, messages := service.GetGroupCollection(startInt, countInt); messages != nil {
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.Output.SetStatus(messages[0].StatusCode)
			log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Get server group collection failed")
		} else {
			resp := new(dto.GetGroupCollectionResponse)
			resp.Load(serverCollection)
			c.Data["json"] = resp
			c.Ctx.Output.SetStatus(http.StatusOK)
		}
	}
	c.ServeJSON()
}

// Delete will delete all the group except default "all" server group.
func (c *GroupRootController) Delete() {
	messages := service.DeleteGroupCollection()
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
	}
	c.Ctx.Output.SetStatus(http.StatusAccepted)
	log.Info("DELETE all server group")
	c.ServeJSON()
}
