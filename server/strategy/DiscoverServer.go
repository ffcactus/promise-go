package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/context"
	"promise/server/object/model"
	"promise/server/strategy/dell"
	"promise/server/strategy/hp"
	"promise/server/strategy/huawei"
)

// DiscoverServer is the interface of post server strategy.
type DiscoverServer interface {
	CreateManagementAccount(c *context.DiscoverServer, server *model.Server) error
	Claim(c *context.DiscoverServer, server *model.Server) error
	Execute(c *context.DiscoverServer, server *model.Server) (base.ModelInterface, error)
}

// CreateDiscoverServerStrategy will create the post server strategy based on the server type.
func CreateDiscoverServerStrategy(server *model.Server) DiscoverServer {
	if server.Vender == "HP" {
		return new(hp.Discover)
	}
	if server.Vender == "Dell" {
		return new(dell.Discover)
	}
	if server.Vender == "Huawei" {
		return new(huawei.DiscoverRackServer)
	}
	if server.Vender == "Mock" {
		return new(huawei.DiscoverMockServer)
	}
	log.WithFields(log.Fields{"hostname": server.Hostname, "vender": server.Vender, "type": server.Type}).Warn("Strategy find discover strategy instance failed.")
	return nil
}
