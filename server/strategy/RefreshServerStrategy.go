package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/server/context"
	"promise/server/object/constValue"
	"promise/server/object/model"
)

// RefreshServerStrategy is the strategy for refresh server.
type RefreshServerStrategy interface {
	RefreshProcessors(c *context.RefreshServerContext, server *model.Server) error
	RefreshMemory(c *context.RefreshServerContext, server *model.Server) error
	RefreshEthernetInterfaces(c *context.RefreshServerContext, server *model.Server) error
	RefreshNetworkInterfaces(c *context.RefreshServerContext, server *model.Server) error
	RefreshStorages(c *context.RefreshServerContext, server *model.Server) error
	RefreshPower(c *context.RefreshServerContext, server *model.Server) error
	RefreshThermal(c *context.RefreshServerContext, server *model.Server) error
	RefreshOemHuaweiBoards(c *context.RefreshServerContext, server *model.Server) error
	RefreshNetworkAdapters(c *context.RefreshServerContext, server *model.Server) error
	RefreshDrives(c *context.RefreshServerContext, server *model.Server) error
	RefreshPCIeDevices(c *context.RefreshServerContext, server *model.Server) error
	Execute(c *context.RefreshServerContext, server *model.Server) error
}

// CreateRefreshServerStrategy creates the strategy based on server.
func CreateRefreshServerStrategy(server *model.Server) RefreshServerStrategy {
	switch server.Type {
	case constValue.RackType:
		return new(RackServerRefreshStrategy)
	case constValue.MockType:
		return new(MockServerRefreshStrategy)
	default:
		log.WithFields(log.Fields{"id": server.ID, "type": server.Type}).Warn("Can not find refresh server strategy.")
		return nil
	}
}
