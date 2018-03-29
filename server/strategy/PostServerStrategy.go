package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/server/context"
	"promise/server/object/constValue"
	"promise/server/object/model"
)

// PostServerStrategy is the interface of post server strategy.
type PostServerStrategy interface {
	CreateManagementAccount(c *context.PostServerContext, server *model.Server) error
	Claim(c *context.PostServerContext, server *model.Server) error
	Execute(c *context.PostServerContext, server *model.Server) error
}

// CreatePostServerStrategy will create the post server strategy based on the server type.
func CreatePostServerStrategy(server *model.Server) PostServerStrategy {
	switch server.Type {
	case constValue.RackType:
		return new(RackServerPostStrategy)
	case constValue.MockType:
		return new(MockServerPostStrategy)
	default:
		log.WithFields(log.Fields{"hostname": server.Hostname, "type": server.Type}).Warn("Can not find post server strategy.")
		return nil
	}
}
