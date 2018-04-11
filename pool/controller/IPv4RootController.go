package controller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
	commonDto "promise/common/object/dto"
	commonMessage "promise/common/object/message"
	commonConstError "promise/common/object/consterror"
	"promise/pool/object/consterror"
	"promise/pool/object/message"
	"promise/pool/object/dto"
	"promise/pool/service"
	"strconv"
)

// IPv4RootController is the ipv4 pool controller.
type IPv4RootController struct {
	beego.Controller
}

// Post a new IPv4 range.
func (c *IPv4RootController) Post() {
	var (
		request  dto.PostIPv4PoolRequest
		response dto.GetIPv4PoolResponse
	)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		messages := []commonMessage.Message{}
		messages = append(messages, commonMessage.NewInvalidRequest())
		log.WithFields(log.Fields{
			"error":   err,
			"message": messages[0].ID}).
			Warn("Post IPv4 pool failed, unable to unmarshal request.")

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
			break;
		case consterror.ErrorRangeEndAddress.Error():
			messages = append(messages, message.NewIPv4RangeEndAddressError())
			break;
		case consterror.ErrorRangeSize.Error():
			messages = append(messages, message.NewIPv4RangeSizeError())
			break;
		case consterror.ErrorRangeCount.Error():
			messages = append(messages, message.NewIPv4RangeCountError())
		default:
			messages = append(messages, commonMessage.NewInvalidRequest())
			break;
		}
		log.WithFields(log.Fields{
			"message": messages[0].ID}).
			Warn("Post IPv4 pool failed, invalid request.")
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
	log.WithFields(log.Fields{"name": request.Name}).Info("Post IPv4 pool.")

	ipv4Pool, messages := service.PostIPv4Pool(&request)
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Post IPv4 pool failed.")
	} else {
		response.Load(ipv4Pool)
		c.Data["json"] = &response
		c.Ctx.Output.SetStatus(http.StatusCreated)
		log.WithFields(log.Fields{"name": response.Name, "ID": response.ID}).Info("Post IPv4 pool done.")
	}
	c.ServeJSON()
}

// Get will return IPv4 pool collection.
func (c *IPv4RootController) Get() {
	var (
		start, count, filter string = c.GetString("start"), c.GetString("count"), c.GetString("$filter")
		startInt, countInt   int    = 0, -1
		parameterError       bool
	)
	log.WithFields(log.Fields{"start": start, "count": count}).Debug("Get IPv4 pool collection.")
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
		log.Warn("Get IPv4 pool collection failed, parameter error.")
	} else {
		if collection, messages := service.GetIPv4PoolCollection(startInt, countInt, filter); messages != nil {
			c.Data["json"] = commonDto.MessagesToDto(messages)
			c.Ctx.Output.SetStatus(messages[0].StatusCode)
			log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Get IPv4 pool collection failed.")
		} else {
			resp := new(dto.GetIPv4PoolCollectionResponse)
			resp.Load(collection)
			c.Data["json"] = resp
			c.Ctx.Output.SetStatus(http.StatusOK)
		}
	}
	c.ServeJSON()
}

// Delete will delete all the IPv4 pool.
func (c *IPv4RootController) Delete() {
	messages := service.DeleteIPv4PoolCollection()
	if messages != nil {
		c.Data["json"] = commonDto.MessagesToDto(messages)
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		log.WithFields(log.Fields{"message": messages[0].ID}).Warn("Delete IPv4 pool collection failed.")
	} else {
		c.Ctx.Output.SetStatus(http.StatusAccepted)
		log.Info("DELETE all IPv4 pool.")
	}
	c.ServeJSON()
}

func (c *IPv4RootController) isValidFilter(filter string) bool {
	return true
}
