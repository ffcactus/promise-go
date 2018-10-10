package context

import (
	beegoCtx "github.com/astaxie/beego/context"
	"net/http"
	"promise/base"
	"promise/enclosure/client/enclosure"
	"promise/enclosure/db"
)

// Base containes the basic context info
type Base struct {
	*beegoCtx.Context
	Client    enclosure.Client
	DB        *db.Enclosure
	Enclosure base.ModelInterface
	TaskURL   string
	ID        string
}

// SendResponse sents the response to client.
func (c Base) SendResponse(resposne base.ResponseInterface, taskURL string, errorResps []base.ErrorResponse) {
	if errorResps != nil {
		c.Context.Output.SetStatus(errorResps[0].StatusCode)
		c.Context.Output.JSON(&errorResps, true, false)
	} else {
		if taskURL != "" {
			c.Context.Output.Header("Location", taskURL)
		}
		c.Context.Output.SetStatus(http.StatusAccepted)
		c.Context.Output.JSON(resposne, true, false)
	}
}
