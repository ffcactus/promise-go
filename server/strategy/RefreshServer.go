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

// RefreshServer is the strategy for refresh server.
type RefreshServer interface {
	RefreshProcessors(c *context.RefreshServer, server *model.Server) error
	RefreshMemory(c *context.RefreshServer, server *model.Server) error
	RefreshEthernetInterfaces(c *context.RefreshServer, server *model.Server) error
	RefreshNetworkInterfaces(c *context.RefreshServer, server *model.Server) error
	RefreshStorages(c *context.RefreshServer, server *model.Server) error
	RefreshPower(c *context.RefreshServer, server *model.Server) error
	RefreshThermal(c *context.RefreshServer, server *model.Server) error
	RefreshBoards(c *context.RefreshServer, server *model.Server) error
	RefreshNetworkAdapters(c *context.RefreshServer, server *model.Server) error
	RefreshDrives(c *context.RefreshServer, server *model.Server) error
	RefreshPCIeDevices(c *context.RefreshServer, server *model.Server) error
	Execute(c *context.RefreshServer, server *model.Server) (string, []base.ErrorResponse)
}

// CreateRefreshServerStrategy creates the strategy based on server.
func CreateRefreshServerStrategy(server *model.Server) RefreshServer {
	if server.Vender == "HP" {
		return new(hp.Refresh)
	}
	if server.Vender == "Dell" {
		return new(dell.Refresh)
	}
	if server.Vender == "Huawei" {
		return new(huawei.RefreshRackServer)
	}
	if server.Vender == "Mock" {
		return new(huawei.RefreshMockServer)
	}
	log.WithFields(log.Fields{"id": server.ID, "vender": server.Vender, "type": server.Type}).Warn("Strategy find refresh strategy instance failed.")
	return nil
}
