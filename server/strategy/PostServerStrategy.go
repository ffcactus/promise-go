package strategy

import (
	"promise/server/context"
	"promise/server/object/model"
	"github.com/astaxie/beego"
)

// PostServerStrategy is the interface of post server strategy.
type PostServerStrategy interface {
	CreateManagementAccount(c *context.PostServerContext) error
	Claim(c *context.PostServerContext) error
	Execute(c *context.PostServerContext) error
}

// CreatePostServerStrategy will create the post server strategy based on the server type.
func CreatePostServerStrategy(server *model.Server) PostServerStrategy {
	switch server.Type {
	case model.RackType:
		return new(RackServerPostStrategy)
	case model.MockType:
		return new(MockServerPostStrategy)
	default:
		beego.Warning("CreatePostServerStrategy() failed, server type = ", server.Type)
		return nil
	}
}
