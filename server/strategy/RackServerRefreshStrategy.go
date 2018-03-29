package strategy

import (
	"promise/server/context"
	"promise/server/object/constValue"
	"promise/server/object/model"
	log "github.com/sirupsen/logrus"
)

// RackServerRefreshStrategy is the strategy of rack server refresh.
type RackServerRefreshStrategy struct {
	ServerStrategy
	ServerEventStrategy
	ServerTaskStrategy
}

// Execute will execute all the steps.
func (s *RackServerRefreshStrategy) Execute(c *context.RefreshServerContext, server *model.Server) error {
	// defer s.IndexServer(&c.ServerContext)
	// defer s.SetServerState(&c.ServerContext, ServerStateReady)
	// defer s.SetServerHealth(&c.ServerContext, ServerHealthOK)
	log.WithFields(log.Fields{"id": server.ID}).Info("Refresh server.")
	// Lock server.
	if err := s.LockServer(&c.ServerContext, server); err != nil {
		return err
	}
	
	taskID, err := s.CreateRefreshServerTask(&c.ServerContext, server)
	if err != nil {
		return err
	}

	// Chassis.Power
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNamePower, c, server, s.RefreshPower); err != nil {
		return err
	}
	// Chassis.Thermal
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNameThermal, c, server, s.RefreshThermal); err != nil {
		return err
	}
	// Chassis.OemHuaweiBoards
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNameBoards, c, server, s.RefreshOemHuaweiBoards); err != nil {
		return err
	}
	// Chassis.NetworkAdapters
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNameNetworkAdapters, c, server, s.RefreshNetworkAdapters); err != nil {
		return err
	}
	// Chassis.Drives
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNameDrives, c, server, s.RefreshDrives); err != nil {
		return err
	}
	// Chassis.PCIeDevices
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNamePCIeDevices, c, server, s.RefreshPCIeDevices); err != nil {
		return err
	}
	// ComputerSystem.Processors
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNameProcessors, c, server, s.RefreshProcessors); err != nil {
		return err
	}
	// ComputerSystem.Memory
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNameMemory, c, server, s.RefreshMemory); err != nil {
		return err
	}
	// ComputerSystem.EthernetInterfaces
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNameEthernetInterfaces, c, server, s.RefreshEthernetInterfaces); err != nil {
		return err
	}
	// ComputerSystem.NetworkInterfaces
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNameNetworkInterfaces, c, server, s.RefreshNetworkInterfaces); err != nil {
		return err
	}
	// ComputerSystem.Storages
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNameStorages, c, server, s.RefreshStorages); err != nil {
		return err
	}
	s.SetServerHealth(&c.ServerContext, server, constValue.ServerHealthOK)
	s.SetServerState(&c.ServerContext, server, constValue.ServerStateReady)
	log.WithFields(log.Fields{"id": server.ID}).Info("Refresh server done.")
	return nil
}

// Stepfunc is the func point.
type Stepfunc func(c *context.RefreshServerContext, server *model.Server) error

// StepWarper is a warper.
func (s *RackServerRefreshStrategy) StepWarper(
	id string, 
	stepName string, 
	c *context.RefreshServerContext, 
	server *model.Server, 
	stepfunc Stepfunc) error {
	s.SetTaskStepRunning(&c.ServerContext, id, stepName, server)
	err := stepfunc(c, server)
	if err != nil {
		s.SetTaskStepError(&c.ServerContext, id, stepName, server)
	} else {
		s.SetTaskStepFinished(&c.ServerContext, id, stepName, server)
	}
	return err
}

// RefreshProcessors refresh the processor of the server.
func (s *RackServerRefreshStrategy) RefreshProcessors(c *context.RefreshServerContext, server *model.Server) error {
	var component = "Processor"
	processors, err := c.ServerClient.GetProcessors(*server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.DB.UpdateProcessors(server.ID, processors); err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Update server component failed.")
		return err
	}
	s.DispatchServerUpdate(&c.ServerContext, server)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshMemory refresh the memory
func (s *RackServerRefreshStrategy) RefreshMemory(c *context.RefreshServerContext, server *model.Server) error{
	var component = "Memory"
	memory, err := c.ServerClient.GetMemory(*server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.DB.UpdateMemory(server.ID, memory); err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.ServerContext, server)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshEthernetInterfaces refresh the ethernet interfaces.
func (s *RackServerRefreshStrategy) RefreshEthernetInterfaces(c *context.RefreshServerContext, server *model.Server) error{
	var component = "EthernetInterfaces"
	ethernet, err := c.ServerClient.GetEthernetInterfaces(*server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.DB.UpdateEthernetInterfaces(server.ID, ethernet); err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.ServerContext, server)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshNetworkInterfaces refresh the network interfaces.
func (s *RackServerRefreshStrategy) RefreshNetworkInterfaces(c *context.RefreshServerContext, server *model.Server) error{
	var component = "NetworkInterfaces"
	networks, err := c.ServerClient.GetNetworkInterfaces(*server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.DB.UpdateNetworkInterfaces(server.ID, networks); err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.ServerContext, server)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshStorages refresh the storages.
func (s *RackServerRefreshStrategy) RefreshStorages(c *context.RefreshServerContext, server *model.Server) error{
	var component = "Storages"
	storages, err := c.ServerClient.GetStorages(*server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.DB.UpdateStorages(server.ID, storages); err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.ServerContext, server)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshPower refresh the power.
func (s *RackServerRefreshStrategy) RefreshPower(c *context.RefreshServerContext, server *model.Server) error{
	var component = "Power"
	power, err := c.ServerClient.GetPower(*server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.DB.UpdatePower(server.ID, power); err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.ServerContext, server)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshThermal refresh the thermal.
func (s *RackServerRefreshStrategy) RefreshThermal(c *context.RefreshServerContext, server *model.Server) error{
	var component = "Thermal"
	thermal, err := c.ServerClient.GetThermal(*server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.DB.UpdateThermal(server.ID, thermal); err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.ServerContext, server)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshOemHuaweiBoards refresh the OEM boards.
func (s *RackServerRefreshStrategy) RefreshOemHuaweiBoards(c *context.RefreshServerContext, server *model.Server) error{
	var component = "OemHuaweiBoards"
	boards, err := c.ServerClient.GetOemHuaweiBoards(*server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.DB.UpdateOemHuaweiBoards(server.ID, boards); err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.ServerContext, server)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshNetworkAdapters refresh the network adapters.
func (s *RackServerRefreshStrategy) RefreshNetworkAdapters(c *context.RefreshServerContext, server *model.Server) error{
	var component = "NetworkAdapters"
	networkAdapters, err := c.ServerClient.GetNetworkAdapters(*server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.DB.UpdateNetworkAdapters(server.ID, networkAdapters); err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.ServerContext, server)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshDrives refresh the drives.
func (s *RackServerRefreshStrategy) RefreshDrives(c *context.RefreshServerContext, server *model.Server) error{
	var component = "Drives"
	drives, err := c.ServerClient.GetDrives(*server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.DB.UpdateDrives(server.ID, drives); err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.ServerContext, server)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshPCIeDevices refresh the PCIe devices.
func (s *RackServerRefreshStrategy) RefreshPCIeDevices(c *context.RefreshServerContext, server *model.Server) error{
	var component = "PCIeDevices"
	pcieDevice, err := c.ServerClient.GetPCIeDevices(*server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.DB.UpdatePCIeDevices(server.ID, pcieDevice); err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.ServerContext, server)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}
