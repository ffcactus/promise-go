package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/context"
	"promise/server/object/constvalue"
	"promise/server/object/model"
)

// DiscoverServer is the interface of post server strategy.
type DiscoverServer interface {
	CreateManagementAccount(c *context.DiscoverServer, server *model.Server) error
	Claim(c *context.DiscoverServer, server *model.Server) error
	Execute(c *context.DiscoverServer, server *model.Server) (base.ModelInterface, error)
}

// CreateDiscoverServerStrategy will create the post server strategy based on the server type.
func CreateDiscoverServerStrategy(server *model.Server) DiscoverServer {
	switch server.Type {
	case constvalue.RackType:
		return new(DiscoverRackServer)
	case constvalue.MockType:
		return new(DiscoverMockServer)
	default:
		log.WithFields(log.Fields{"hostname": server.Hostname, "type": server.Type}).Warn("Can not find post server strategy.")
		return nil
	}
}
