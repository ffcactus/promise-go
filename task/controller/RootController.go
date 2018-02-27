package controller

import (
	"encoding/json"
	commonDto "promise/common/object/dto"
	. "promise/common/object/model"
	. "promise/task/object/dto"
	"promise/task/object/model"
	"promise/task/service"
	"strconv"

	"github.com/astaxie/beego"
)

type RootController struct {
	beego.Controller
}

// Post Post a new task.
func (this *RootController) Post() {
	beego.Trace("Post task.")
	request := new(PostTaskRequest)
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, request); err != nil {
		beego.Warning("Unmarshal() failed, error = ", err)
	}
	// Create the context for this operation.
	if task, messages := service.PostTask(request); messages != nil {
		this.Data["json"] = commonDto.MessagesToDto(messages)
		this.Ctx.Output.SetStatus(messages[0].StatusCode)
	} else {
		resp := PostTaskResponse{}
		resp.Load(task)
		this.Data["json"] = &resp
	}
	this.ServeJSON()
}

// Get Get task collection.
func (this *RootController) Get() {
	var (
		start, count       string = this.GetString("start"), this.GetString("count")
		startInt, countInt int    = 0, -1
		parameterError     bool   = false
	)
	beego.Trace("Get task collection, start = ", start, ", count = ", count)
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
		messages := []Message{}
		messages = append(messages, model.NewMessageTaskBadRequest())
		this.Data["json"] = commonDto.MessagesToDto(messages)
		this.Ctx.ResponseWriter.WriteHeader((messages)[0].StatusCode)
	} else {
		if serverCollection, messages := service.GetTaskCollection(startInt, countInt); messages != nil {
			this.Data["json"] = commonDto.MessagesToDto(messages)
			this.Ctx.ResponseWriter.WriteHeader(messages[0].StatusCode)
		} else {
			resp := GetTaskCollectionResponse{}
			resp.Load(serverCollection)
			this.Data["json"] = &resp
		}
	}
	this.ServeJSON()
}
