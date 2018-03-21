package controller

import (
	"encoding/json"
	commonDto "promise/common/object/dto"
	commonMessage "promise/common/object/message"
	"promise/task/object/dto"
	"promise/task/object/message"
	"promise/task/service"
	"strconv"

	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
)

// RootController is the root controller for task.
type RootController struct {
	beego.Controller
}

// Post Post a new task.
func (c *RootController) Post() {
	log.Debug("Post task.")
	request := new(dto.PostTaskRequest)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, request); err != nil {
		log.Warn("Unmarshal() failed, error = ", err)
	}
	// Create the context for this operation.
	if task, messages := service.PostTask(request); messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
	} else {
		resp := new(dto.PostTaskResponse)
		resp.Load(task)
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Get Get task collection.
func (c *RootController) Get() {
	var (
		start, count       string = c.GetString("start"), c.GetString("count")
		startInt, countInt int    = 0, -1
		parameterError            = false
	)
	log.Debug("Get task collection, start = ", start, ", count = ", count)
	if start != "" {
		_startInt, err := strconv.Atoi(start)
		if err != nil || _startInt < 0 {
			log.Warn("Get(), invalid 'start' parameter, error = ", err)
			parameterError = true
		} else {
			startInt = _startInt
		}
	}
	if count != "" {
		_countInt, err := strconv.Atoi(count)
		// -1 means all.
		if err != nil || _countInt < -1 {
			log.Warn("Get() 'count' parameter error = %s\n", err)
			parameterError = true
		} else {
			countInt = _countInt
		}
	}

	if parameterError {
		messages := []commonMessage.Message{}
		messages = append(messages, message.NewMessageTaskBadRequest())
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.ResponseWriter.WriteHeader((messages)[0].StatusCode)
	} else {
		if serverCollection, messages := service.GetTaskCollection(startInt, countInt); messages != nil {
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.ResponseWriter.WriteHeader(messages[0].StatusCode)
		} else {
			resp := new(dto.GetTaskCollectionResponse)
			resp.Load(serverCollection)
			c.Data["json"] = resp
		}
	}
	c.ServeJSON()
}
