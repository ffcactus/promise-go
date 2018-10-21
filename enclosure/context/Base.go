package context

import (
	"fmt"
	beegoCtx "github.com/astaxie/beego/context"
	log "github.com/sirupsen/logrus"
	"net/http"
	"promise/base"
	"promise/enclosure/client/enclosure"
	"promise/enclosure/db"
	"promise/enclosure/object/model"
)

// Base containes the basic context info
type Base struct {
	*beegoCtx.Context
	Client    enclosure.Client
	DB        *db.Enclosure
	Enclosure *model.Enclosure
	TaskID    string
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

// String return the debug info.
func (c Base) String() string {
	return fmt.Sprintf("(ID = %s)", c.ID)
}

// UpdateEnclosure will update the enclosure reference in context.
func (c *Base) UpdateEnclosure(i base.ModelInterface) {
	m, ok := i.(*model.Enclosure)
	if !ok {
		log.Error("Context update enclosure failed, convert to model failed.")
		return
	}
	c.Enclosure = m
}
