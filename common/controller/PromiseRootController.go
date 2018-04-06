package controller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"promise/common/object/consterror"
	"promise/common/object/dto"
	"promise/common/object/message"
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
