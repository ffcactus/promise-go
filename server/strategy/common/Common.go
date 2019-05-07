package common

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/context"
	"promise/server/object/model"
)

// Common contains basic method for other strategies.
type Common struct {
	ServerEvent
	ServerTask
}

// LockServer Lock the server.
func (s *Common) LockServer(c *context.Base, server *model.Server) *base.ErrorResponse {
	success, lockedServer := c.DB.GetAndLock(server.ID)
	if lockedServer == nil {
		log.WithFields(log.Fields{"id": server.ID}).Info("Can not get and lock server, server not exist.")
		return base.NewErrorResponseNotExist()
	}
	if !success {
		log.WithFields(log.Fields{"id": server.ID, "state": server.State}).Info("Can not get and lock server.")
		return base.NewErrorResponseErrorState()
	}
	s.DispatchServerUpdate(c, server)
	return nil
}

// SetState Set server state.
func (s *Common) SetState(c *context.Base, server *model.Server, state string) error {
	updatedServer, err := c.DB.SetState(server.ID, state)
	if err != nil {
		return base.ErrorTransaction
	}
	s.DispatchServerUpdate(c, updatedServer)
	return nil
}

// SetServerHealth Set server health.
func (s *Common) SetServerHealth(c *context.Base, server *model.Server, health string) error {
	updatedServer, err := c.DB.SetServerHealth(server.ID, health)
	if err != nil {
		return base.ErrorTransaction
	}
	s.DispatchServerUpdate(c, updatedServer)
	return nil
}

// RefreshProcessors refresh the processor of the server.
func (s *Common) RefreshProcessors(c *context.RefreshServer, server *model.Server) error {
	var component = "Processor"
	processors, err := c.ServerClient.GetProcessors(server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdateProcessors(server.ID, processors)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy update server component failed.")
		return err
	}
	_updatedServer := updatedServer.(*model.Server)
	log.WithFields(log.Fields{"id": server.ID, "component": component, "count": len(_updatedServer.ComputerSystem.Processors)}).Info("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshMemory refresh the memory
func (s *Common) RefreshMemory(c *context.RefreshServer, server *model.Server) error {
	var component = "Memory"
	memory, err := c.ServerClient.GetMemory(server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdateMemory(server.ID, memory)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy update server component failed.")
	}
	_updatedServer := updatedServer.(*model.Server)
	log.WithFields(log.Fields{"id": server.ID, "component": component, "count": len(_updatedServer.ComputerSystem.Memory)}).Info("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshEthernetInterfaces refresh the ethernet interfaces.
func (s *Common) RefreshEthernetInterfaces(c *context.RefreshServer, server *model.Server) error {
	var component = "EthernetInterfaces"
	ethernet, err := c.ServerClient.GetEthernetInterfaces(server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdateEthernetInterfaces(server.ID, ethernet)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy update server component failed.")
	}
	_updatedServer := updatedServer.(*model.Server)
	log.WithFields(log.Fields{"id": server.ID, "component": component, "count": len(_updatedServer.ComputerSystem.EthernetInterfaces)}).Info("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshNetworkInterfaces refresh the network interfaces.
func (s *Common) RefreshNetworkInterfaces(c *context.RefreshServer, server *model.Server) error {
	var component = "NetworkInterfaces"
	networks, err := c.ServerClient.GetNetworkInterfaces(server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdateNetworkInterfaces(server.ID, networks)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy update server component failed.")
	}
	_updatedServer := updatedServer.(*model.Server)
	log.WithFields(log.Fields{"id": server.ID, "component": component, "count": len(_updatedServer.ComputerSystem.NetworkInterfaces)}).Info("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshStorages refresh the storages.
func (s *Common) RefreshStorages(c *context.RefreshServer, server *model.Server) error {
	var component = "Storages"
	storages, err := c.ServerClient.GetStorages(server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdateStorages(server.ID, storages)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy update server component failed.")
	}
	_updatedServer := updatedServer.(*model.Server)
	log.WithFields(log.Fields{"id": server.ID, "component": component, "count": len(_updatedServer.ComputerSystem.Storages)}).Info("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshPower refresh the power.
func (s *Common) RefreshPower(c *context.RefreshServer, server *model.Server) error {
	var component = "Power"
	power, err := c.ServerClient.GetPower(server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdatePower(server.ID, power)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy update server component failed.")
	}
	_updatedServer := updatedServer.(*model.Server)
	log.WithFields(log.Fields{"id": server.ID, "component": component, "power": _updatedServer.Chassis.Power}).Info("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshThermal refresh the thermal.
func (s *Common) RefreshThermal(c *context.RefreshServer, server *model.Server) error {
	var component = "Thermal"
	thermal, err := c.ServerClient.GetThermal(server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdateThermal(server.ID, thermal)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy update server component failed.")
	}
	_updatedServer := updatedServer.(*model.Server)
	log.WithFields(log.Fields{"id": server.ID, "component": component, "thermal": _updatedServer.Chassis.Thermal}).Info("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshBoards refresh the boards.
func (s *Common) RefreshBoards(c *context.RefreshServer, server *model.Server) error {
	var component = "Boards"
	boards, err := c.ServerClient.GetBoards(server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdateBoards(server.ID, boards)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy update server component failed.")
	}
	_updatedServer := updatedServer.(*model.Server)
	log.WithFields(log.Fields{"id": server.ID, "component": component, "count": len(_updatedServer.Chassis.Boards)}).Info("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshNetworkAdapters refresh the network adapters.
func (s *Common) RefreshNetworkAdapters(c *context.RefreshServer, server *model.Server) error {
	var component = "NetworkAdapters"
	networkAdapters, err := c.ServerClient.GetNetworkAdapters(server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdateNetworkAdapters(server.ID, networkAdapters)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy update server component failed.")
	}
	_updatedServer := updatedServer.(*model.Server)
	log.WithFields(log.Fields{"id": server.ID, "component": component, "count": len(_updatedServer.Chassis.NetworkAdapters)}).Info("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshDrives refresh the drives.
func (s *Common) RefreshDrives(c *context.RefreshServer, server *model.Server) error {
	var component = "Drives"
	drives, err := c.ServerClient.GetDrives(server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdateDrives(server.ID, drives)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy update server component failed.")
	}
	_updatedServer := updatedServer.(*model.Server)
	log.WithFields(log.Fields{"id": server.ID, "component": component, "count": len(_updatedServer.Chassis.Drives)}).Info("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshPCIeDevices refresh the PCIe devices.
func (s *Common) RefreshPCIeDevices(c *context.RefreshServer, server *model.Server) error {
	var component = "PCIeDevices"
	pcieDevice, err := c.ServerClient.GetPCIeDevices(server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdatePCIeDevices(server.ID, pcieDevice)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Strategy update server component failed.")
	}
	_updatedServer := updatedServer.(*model.Server)
	log.WithFields(log.Fields{"id": server.ID, "component": component, "count": len(_updatedServer.Chassis.PCIeDevices)}).Info("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}
