package strategy

import (
	"promise/server/context"
	"promise/server/object/constValue"

	log "github.com/sirupsen/logrus"
)

// RackServerRefreshStrategy is the strategy of rack server refresh.
type RackServerRefreshStrategy struct {
	ServerStrategy
}

// Execute will execute all the steps.
func (s *RackServerRefreshStrategy) Execute(c *context.RefreshServerContext) error {
	// defer s.IndexServer(&c.ServerContext)
	// defer s.SetServerState(&c.ServerContext, ServerStateReady)
	// defer s.SetServerHealth(&c.ServerContext, ServerHealthOK)
	log.WithFields(log.Fields{"id": c.Server.ID}).Info("Refresh server.")
	// Lock server.
	if err := s.LockServer(&c.ServerContext); err != nil {
		return err
	}

	c.CreateTask(CreateRefreshTaskRequest(c.Server))

	// Chassis.Power
	if err := s.StepWarper(ServerRefreshTaskStepNamePower, c, s.RefreshPower); err != nil {
		return err
	}
	// Chassis.Thermal
	if err := s.StepWarper(ServerRefreshTaskStepNameThermal, c, s.RefreshThermal); err != nil {
		return err
	}
	// Chassis.OemHuaweiBoards
	if err := s.StepWarper(ServerRefreshTaskStepNameBoards, c, s.RefreshOemHuaweiBoards); err != nil {
		return err
	}
	// Chassis.NetworkAdapters
	if err := s.StepWarper(ServerRefreshTaskStepNameNetworkAdapters, c, s.RefreshNetworkAdapters); err != nil {
		return err
	}
	// Chassis.Drives
	if err := s.StepWarper(ServerRefreshTaskStepNameDrives, c, s.RefreshDrives); err != nil {
		return err
	}
	// Chassis.PCIeDevices
	if err := s.StepWarper(ServerRefreshTaskStepNamePCIeDevices, c, s.RefreshPCIeDevices); err != nil {
		return err
	}
	// ComputerSystem.Processors
	if err := s.StepWarper(ServerRefreshTaskStepNameProcessors, c, s.RefreshProcessors); err != nil {
		return err
	}
	// ComputerSystem.Memory
	if err := s.StepWarper(ServerRefreshTaskStepNameMemory, c, s.RefreshMemory); err != nil {
		return err
	}
	// ComputerSystem.EthernetInterfaces
	if err := s.StepWarper(ServerRefreshTaskStepNameEthernetInterfaces, c, s.RefreshEthernetInterfaces); err != nil {
		return err
	}
	// ComputerSystem.NetworkInterfaces
	if err := s.StepWarper(ServerRefreshTaskStepNameNetworkInterfaces, c, s.RefreshNetworkInterfaces); err != nil {
		return err
	}
	// ComputerSystem.Storages
	if err := s.StepWarper(ServerRefreshTaskStepNameStorages, c, s.RefreshStorages); err != nil {
		return err
	}
	s.SetServerHealth(&c.ServerContext, constValue.ServerHealthOK)
	s.SetServerState(&c.ServerContext, constValue.ServerStateReady)
	s.IndexServer(&c.ServerContext)
	log.WithFields(log.Fields{"id": c.Server.ID}).Info("Refresh server done.")
	return nil
}

// Stepfunc is the func point.
type Stepfunc func(c *context.RefreshServerContext) error

// StepWarper is a warper.
func (s *RackServerRefreshStrategy) StepWarper(stepName string, c *context.RefreshServerContext, stepfunc Stepfunc) error {
	c.SetTaskStepRunning(stepName)
	err := stepfunc(c)
	if err != nil {
		c.SetTaskStepError(stepName)
	} else {
		c.SetTaskStepFinished(stepName)
	}
	return err
}

// RefreshProcessors refresh the processor of the server.
func (s *RackServerRefreshStrategy) RefreshProcessors(c *context.RefreshServerContext) error {
	var component = "Processor"
	processors, err := c.ServerClient.GetProcessors(*c.Server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.ServerDBImplement.UpdateProcessors(c.Server.ID, processors); err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Update server component failed.")
		return err
	}
	c.DispatchServerUpdate()
	log.WithFields(log.Fields{"id": c.Server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshMemory refresh the memory
func (s *RackServerRefreshStrategy) RefreshMemory(c *context.RefreshServerContext) error {
	var component = "Memory"
	memory, err := c.ServerClient.GetMemory(*c.Server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.ServerDBImplement.UpdateMemory(c.Server.ID, memory); err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	c.DispatchServerUpdate()
	log.WithFields(log.Fields{"id": c.Server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshEthernetInterfaces refresh the ethernet interfaces.
func (s *RackServerRefreshStrategy) RefreshEthernetInterfaces(c *context.RefreshServerContext) error {
	var component = "EthernetInterfaces"
	ethernet, err := c.ServerClient.GetEthernetInterfaces(*c.Server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.ServerDBImplement.UpdateEthernetInterfaces(c.Server.ID, ethernet); err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	c.DispatchServerUpdate()
	log.WithFields(log.Fields{"id": c.Server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshNetworkInterfaces refresh the network interfaces.
func (s *RackServerRefreshStrategy) RefreshNetworkInterfaces(c *context.RefreshServerContext) error {
	var component = "NetworkInterfaces"
	networks, err := c.ServerClient.GetNetworkInterfaces(*c.Server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.ServerDBImplement.UpdateNetworkInterfaces(c.Server.ID, networks); err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	c.DispatchServerUpdate()
	log.WithFields(log.Fields{"id": c.Server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshStorages refresh the storages.
func (s *RackServerRefreshStrategy) RefreshStorages(c *context.RefreshServerContext) error {
	var component = "Storages"
	storages, err := c.ServerClient.GetStorages(*c.Server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.ServerDBImplement.UpdateStorages(c.Server.ID, storages); err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	c.DispatchServerUpdate()
	log.WithFields(log.Fields{"id": c.Server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshPower refresh the power.
func (s *RackServerRefreshStrategy) RefreshPower(c *context.RefreshServerContext) error {
	var component = "Power"
	power, err := c.ServerClient.GetPower(*c.Server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.ServerDBImplement.UpdatePower(c.Server.ID, power); err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	c.DispatchServerUpdate()
	log.WithFields(log.Fields{"id": c.Server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshThermal refresh the thermal.
func (s *RackServerRefreshStrategy) RefreshThermal(c *context.RefreshServerContext) error {
	var component = "Thermal"
	thermal, err := c.ServerClient.GetThermal(*c.Server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.ServerDBImplement.UpdateThermal(c.Server.ID, thermal); err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	c.DispatchServerUpdate()
	log.WithFields(log.Fields{"id": c.Server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshOemHuaweiBoards refresh the OEM boards.
func (s *RackServerRefreshStrategy) RefreshOemHuaweiBoards(c *context.RefreshServerContext) error {
	var component = "OemHuaweiBoards"
	boards, err := c.ServerClient.GetOemHuaweiBoards(*c.Server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.ServerDBImplement.UpdateOemHuaweiBoards(c.Server.ID, boards); err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	c.DispatchServerUpdate()
	log.WithFields(log.Fields{"id": c.Server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshNetworkAdapters refresh the network adapters.
func (s *RackServerRefreshStrategy) RefreshNetworkAdapters(c *context.RefreshServerContext) error {
	var component = "NetworkAdapters"
	networkAdapters, err := c.ServerClient.GetNetworkAdapters(*c.Server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.ServerDBImplement.UpdateNetworkAdapters(c.Server.ID, networkAdapters); err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	c.DispatchServerUpdate()
	log.WithFields(log.Fields{"id": c.Server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshDrives refresh the drives.
func (s *RackServerRefreshStrategy) RefreshDrives(c *context.RefreshServerContext) error {
	var component = "Drives"
	drives, err := c.ServerClient.GetDrives(*c.Server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.ServerDBImplement.UpdateDrives(c.Server.ID, drives); err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	c.DispatchServerUpdate()
	log.WithFields(log.Fields{"id": c.Server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}

// RefreshPCIeDevices refresh the PCIe devices.
func (s *RackServerRefreshStrategy) RefreshPCIeDevices(c *context.RefreshServerContext) error {
	var component = "PCIeDevices"
	pcieDevice, err := c.ServerClient.GetPCIeDevices(*c.Server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Get server component failed.")
		return err
	}
	if err := c.ServerDBImplement.UpdatePCIeDevices(c.Server.ID, pcieDevice); err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "component": component, "err": err}).Warn("Update server component failed.")
	}
	c.DispatchServerUpdate()
	log.WithFields(log.Fields{"id": c.Server.ID, "component": component}).Info("Refresh server component done.")
	return nil
}
