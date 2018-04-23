package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/server/context"
	"promise/server/object/constvalue"
	"promise/server/object/model"
)

// DiscoverServerStrategy is the interface of post server strategy.
type DiscoverServerStrategy interface {
	CreateManagementAccount(c *context.DiscoverServer, server *model.Server) error
	Claim(c *context.DiscoverServer, server *model.Server) error
	Execute(c *context.DiscoverServer, server *model.Server) (*model.Server, error)
}

// CreateDiscoverServerStrategy will create the post server strategy based on the server type.
func CreateDiscoverServerStrategy(server *model.Server) DiscoverServerStrategy {
	switch server.Type {
	case constvalue.RackType:
		return new(DiscoverRackServerStrategy)
	case constvalue.MockType:
		return new(DiscoverMockServerStrategy)
	default:
		log.WithFields(log.Fields{"hostname": server.Hostname, "type": server.Type}).Warn("Can not find post server strategy.")
		return nil
	}
}
