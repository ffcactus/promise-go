package controller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"promise/auth/object/dto"
	"promise/auth/service"
)

// Login Login controller
type Login struct {
	beego.Controller
}

// Post Post Login.
func (c *Login) Post() {
	log.Debug("POST Login request.")
	request := new(dto.PostLoginRequest)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, request); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warn("Login failed, unable to unmarshal request.")
	}

	if session, messages := service.Login(request); messages != nil {
		c.Data["json"] = &messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{
			"message": messages[0].ID,
		}).Warn("Login failed.")
	} else {
		resp := new(dto.PostLoginResponse)
		resp.Load(session)
		c.Data["json"] = &resp
		log.WithFields(log.Fields{
			"username": request.Name,
		}).Info("Login done.")
	}
	c.ServeJSON()
}
