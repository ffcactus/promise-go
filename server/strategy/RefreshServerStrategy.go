package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/server/context"
	"promise/server/object/model"
)

// RefreshServerStrategy is the strategy for refresh server.
type RefreshServerStrategy interface {
	RefreshProcessors(c *context.RefreshServerContext) error
	RefreshMemory(c *context.RefreshServerContext) error
	RefreshEthernetInterfaces(c *context.RefreshServerContext) error
	RefreshNetworkInterfaces(c *context.RefreshServerContext) error
	RefreshStorages(c *context.RefreshServerContext) error
	RefreshPower(c *context.RefreshServerContext) error
	RefreshThermal(c *context.RefreshServerContext) error
	RefreshOemHuaweiBoards(c *context.RefreshServerContext) error
	RefreshNetworkAdapters(c *context.RefreshServerContext) error
	RefreshDrives(c *context.RefreshServerContext) error
	RefreshPCIeDevices(c *context.RefreshServerContext) error
	Execute(c *context.RefreshServerContext) error
}

// CreateRefreshServerStrategy creates the strategy based on server.
func CreateRefreshServerStrategy(server *model.Server) RefreshServerStrategy {
	switch server.Type {
	case model.RackType:
		return new(RackServerRefreshStrategy)
	case model.MockType:
		return new(MockServerRefreshStrategy)
	default:
		log.Warn("CreateRefreshServerStrategy() failed, server type =", server.Type)
		return nil
	}
}
