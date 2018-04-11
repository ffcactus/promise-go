package controller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
	commonConstError "promise/common/object/consterror"
	commonDto "promise/common/object/dto"
	commonMessage "promise/common/object/message"
	"promise/pool/object/constvalue"
	"promise/pool/object/dto"
	"promise/pool/object/message"
	"promise/pool/service"
	"strings"
)

// IPv4ActionController IPv4 pool action controller
type IPv4ActionController struct {
	beego.Controller
}

// Post will do the action.
func (c *IPv4ActionController) Post() {
	action := c.Ctx.Input.Param(":action")
	id := c.Ctx.Input.Param(":id")
	switch strings.ToLower(action) {
	case constvalue.IPv4PoolActionAllocate:
		var request dto.AllocateIPv4Request
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
			log.WithFields(log.Fields{
				"action": action,
				"id":     id,
				"error":  err}).
				Warn("Allocate from IPv4 pool failed, unable to unmarshal request.")
			messages := []commonMessage.Message{}
			messages = append(messages, commonMessage.NewInvalidRequest())
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.Output.SetStatus(messages[0].StatusCode)
			c.ServeJSON()
			return
		}
		key := ""
		if request.Key != nil {
			key = *request.Key
		}
		log.WithFields(log.Fields{
			"action": action,
			"id":     id,
			"key":    key}).
			Info("Allocate IP from pool.")

		if address, pool, messages := service.AllocateIPv4Address(id, key); messages != nil {
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.Output.SetStatus(messages[0].StatusCode)
			log.WithFields(log.Fields{
				"action":  action,
				"id":      id,
				"message": messages[0].ID}).
				Warn("Allocate IPv4 failed.")
		} else {
			resp := dto.AllocateIPv4Response{}
			resp.Address = address
			if err := resp.Pool.Load(pool); err != nil {
				messages := []commonMessage.Message{}
				messages = append(messages, commonMessage.NewInternalError())
				log.WithFields(log.Fields{
					"action":  action,
					"id":      id,
					"address": address,
					"message": messages[0].ID}).
					Warn("Allocate IPv4 failed, failed to encode response, address is allocated.")
				c.Data["json"] = commonDto.MessagesToDto(messages)
				c.Ctx.Output.SetStatus(messages[0].StatusCode)
				c.ServeJSON()
				return
			}
			c.Data["json"] = &resp
			c.Ctx.Output.SetStatus(http.StatusAccepted)
		}
	case constvalue.IPv4PoolActionFree:
		var request dto.FreeIPv4Request
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
			log.WithFields(log.Fields{
				"action": action,
				"id":     id,
				"error":  err}).
				Warn("Free IPv4 failed, unable to unmarshal request.")
			messages := []commonMessage.Message{}
			messages = append(messages, commonMessage.NewInvalidRequest())
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.Output.SetStatus(messages[0].StatusCode)
			c.ServeJSON()
			return
		}
		if err := request.Validate(); err != nil {
			messages := []commonMessage.Message{}
			switch err.Error() {
			case commonConstError.ErrorDataConvert.Error():
				messages = append(messages, message.NewIPv4FormatError())
				break
			default:
				messages = append(messages, commonMessage.NewInvalidRequest())
				break
			}
			log.WithFields(log.Fields{
				"message": messages[0].ID}).
				Warn("Free IPv4 failed, invalid request.")
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.Output.SetStatus(messages[0].StatusCode)
			c.ServeJSON()
			return
		}
		log.WithFields(log.Fields{
			"action":  action,
			"id":      id,
			"address": request.Address}).
			Info("Free IPv4 to pool.")
		if pool, messages := service.FreeIPv4Address(id, request.Address); messages != nil {
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.Output.SetStatus(messages[0].StatusCode)
			log.WithFields(log.Fields{
				"action":  action,
				"id":      id,
				"message": messages[0].ID}).
				Warn("Free IPv4 failed.")
		} else {
			resp := dto.GetIPv4PoolResponse{}
			if err := resp.Load(pool); err != nil {
				messages := []commonMessage.Message{}
				messages = append(messages, commonMessage.NewInternalError())
				log.WithFields(log.Fields{
					"action":  action,
					"id":      id,
					"address": request.Address,
					"message": messages[0].ID}).
					Warn("Free IPv4 failed, failed to encode response, address is allocated.")
				c.Data["json"] = commonDto.MessagesToDto(messages)
				c.Ctx.Output.SetStatus(messages[0].StatusCode)
				c.ServeJSON()
				return
			}
			c.Data["json"] = &resp
			c.Ctx.Output.SetStatus(http.StatusAccepted)
		}
	default:
		messages := []commonMessage.Message{}
		messages = append(messages, commonMessage.NewInvalidRequest())
		log.WithFields(log.Fields{
			"action":  action,
			"id":      id,
			"message": messages[0].ID}).
			Info("Perform action on IPv4 pool failed.")
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
	}
	c.ServeJSON()
}
