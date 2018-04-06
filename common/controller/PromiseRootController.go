package controller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"promise/common/object/consterror"
	"promise/common/object/dto"
	"promise/common/object/message"
	"strconv"
)

// PromiseRootController is the common root controller in Promise project.
type PromiseRootController struct {
	beego.Controller
}

// Post will do basic post operation, it will check the request.
func (c *PromiseRootController) Post(requestP dto.PromiseRequestInterface) (*message.Message, error) {
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, requestP); err != nil {
		message := message.NewInvalidRequest()
		return &message, err
	}

	if message := requestP.Validate(); message != nil {
		return message, consterror.ErrorDataConvert
	}
	return nil, nil
}

// Get will do basic get operation.
func (c *PromiseRootController) Get() (int64, int64, string, *message.Message, error) {
	var (
		start, count, filter string = c.GetString("start"), c.GetString("count"), c.GetString("$filter")
		startInt, countInt   int64  = 0, -1
		parameterError       bool
	)

	if start != "" {
		_startInt, err := strconv.ParseInt(start, 10, 64)
		if err != nil || _startInt < 0 {
			parameterError = true
		} else {
			startInt = _startInt
		}
	}
	if count != "" {
		_countInt, err := strconv.ParseInt(count, 10, 64)
		// -1 means all.
		if err != nil || _countInt < -1 {
			parameterError = true
		} else {
			countInt = _countInt
		}
	}
	if parameterError {
		m := message.NewInvalidRequest()
		return startInt, countInt, filter, &m, consterror.ErrorInvalidURLParameter
	}
	return startInt, countInt, filter, nil, nil
}
