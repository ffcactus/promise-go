package controller

import (
	"encoding/json"
	"net/http"
	commonDto "promise/common/object/dto"
	commomMessage "promise/common/object/message"
	"promise/server/object/dto"
	"promise/server/object/message"
	"promise/server/service"
	"strconv"

	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
)

// ServerRootController The root controller
type ServerRootController struct {
	beego.Controller
}

// Post Post a new server.
func (c *ServerRootController) Post() {
	request := new(dto.PostServerRequest)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, request); err != nil {
		log.WithFields(log.Fields{"err": err}).Info("Post server failed, unable to get request.")
		messages := []commomMessage.Message{}
		messages = append(messages, message.NewServerParameterError())
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	log.WithFields(log.Fields{"hostname": request.Hostname}).Info("Post server.")
	// Create the context for this operation.
	server, messages := service.PostServer(request)
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Info("Post server failed.")
	} else {
		resp := dto.PostServerResponse{}
		resp.Load(server)
		c.Data["json"] = &resp
		c.Ctx.Output.SetStatus(http.StatusCreated)
		log.WithFields(log.Fields{"name": request.Hostname, "ID": resp.ID}).Info("Post server done.")
	}
	c.ServeJSON()
}

// Get Get server collection.
func (c *ServerRootController) Get() {
	var (
		start, count       string = c.GetString("start"), c.GetString("count")
		startInt, countInt int    = 0, -1
		parameterError     bool
	)
	log.WithFields(log.Fields{"start": start, "count": count}).Debug("Get server collection.")
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
		messages := []commomMessage.Message{}
		messages = append(messages, message.NewServerParameterError())
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.Warn("Get server collection failed, parameter error.")
	} else {
		if serverCollection, messages := service.GetServerCollection(startInt, countInt); messages != nil {
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.Output.SetStatus(messages[0].StatusCode)
			log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Get server collection failed")
		} else {
			resp := new(dto.GetServerCollectionResponse)
			resp.Load(serverCollection)
			c.Data["json"] = resp
			c.Ctx.Output.SetStatus(http.StatusOK)
		}
	}
	c.ServeJSON()
}

// Delete will delete all servers.
func (c *ServerRootController) Delete() {
	messages := service.DeleteServerCollection()
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
	} else {
		c.Ctx.Output.SetStatus(http.StatusAccepted)
		log.Info("DELETE all servers.")
	}
	c.ServeJSON()
}
