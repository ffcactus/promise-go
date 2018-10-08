package context

import (
	beegoCtx "github.com/astaxie/beego/context"
	"promise/base"
)

// Base containes the basic context info
type Base struct {
	*beegoCtx.Context
	ID string
}

// SendResponse sents the response to client.
func (c Base) SendResponse(resposne base.ResponseInterface, task string, errorResps []base.ErrorResponse) {
	if errorResps != nil {
		c.Context.Output.SetStatus(errorResps[0].StatusCode)
		c.Context.Output.JSON(&errorResps, true, false)
	} else {
		if task != "" {
			c.Context.Output.Header("Location", task)
		}
		c.Context.Output.SetStatus(errorResps[0].StatusCode)
		c.Context.Output.JSON(resposne, true, false)
	}
}
