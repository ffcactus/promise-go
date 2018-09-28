package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/context"
	"promise/server/object/constvalue"
	"promise/server/object/model"
)

// RefreshServer is the strategy for refresh server.
type RefreshServer interface {
	RefreshProcessors(c *context.RefreshServer, server *model.Server) error
	RefreshMemory(c *context.RefreshServer, server *model.Server) error
	RefreshEthernetInterfaces(c *context.RefreshServer, server *model.Server) error
	RefreshNetworkInterfaces(c *context.RefreshServer, server *model.Server) error
	RefreshStorages(c *context.RefreshServer, server *model.Server) error
	RefreshPower(c *context.RefreshServer, server *model.Server) error
	RefreshThermal(c *context.RefreshServer, server *model.Server) error
	RefreshOemHuaweiBoards(c *context.RefreshServer, server *model.Server) error
	RefreshNetworkAdapters(c *context.RefreshServer, server *model.Server) error
	RefreshDrives(c *context.RefreshServer, server *model.Server) error
	RefreshPCIeDevices(c *context.RefreshServer, server *model.Server) error
	Execute(c *context.RefreshServer, server *model.Server) (*string, []base.ErrorResponse)
}

// CreateRefreshServerStrategy creates the strategy based on server.
func CreateRefreshServerStrategy(server *model.Server) RefreshServer {
	switch server.Type {
	case constvalue.RackType:
		return new(RefreshRackServer)
	case constvalue.MockType:
		return new(RefreshMockServer)
	default:
		log.WithFields(log.Fields{"id": server.ID, "type": server.Type}).Warn("Can not find refresh server strategy.")
		return nil
	}
}
