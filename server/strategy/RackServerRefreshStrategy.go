package strategy

import (
	"promise/server/context"
	"promise/server/util"

	"github.com/astaxie/beego"
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

	beego.Info("Refresh server start", "Server ID =", c.Server.ID)
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
	s.SetServerHealth(&c.ServerContext, util.ServerHealthOK)
	s.SetServerState(&c.ServerContext, util.ServerStateReady)
	s.IndexServer(&c.ServerContext)
	beego.Info("Refresh server done", "Server ID =", c.Server.ID)
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
	processors, err := c.ServerClient.GetProcessors(*c.Server.OriginURIs.System)
	if err != nil {
		beego.Warning("GetProcessors() failed, error =", err)
		return err
	}
	if err := c.ServerDBImplement.UpdateProcessors(c.Server.ID, processors); err != nil {
		beego.Warning("UpdateProcessors() failed, error =", err)
		return err
	}
	c.DispatchServerUpdate()
	beego.Trace("RefreshProcessors() done, server ID =", c.Server.ID)
	return nil
}

// RefreshMemory refresh the memory
func (s *RackServerRefreshStrategy) RefreshMemory(c *context.RefreshServerContext) error {
	memory, err := c.ServerClient.GetMemory(*c.Server.OriginURIs.System)
	if err != nil {
		beego.Warning("GetMemory() failed, error =", err)
		return err
	}
	if err := c.ServerDBImplement.UpdateMemory(c.Server.ID, memory); err != nil {
		beego.Warning("UpdateMemory() failed, error =", err)
	}
	c.DispatchServerUpdate()
	beego.Trace("RefreshMemory() done, server ID =", c.Server.ID)
	return nil
}

// RefreshEthernetInterfaces refresh the ethernet interfaces.
func (s *RackServerRefreshStrategy) RefreshEthernetInterfaces(c *context.RefreshServerContext) error {
	ethernet, err := c.ServerClient.GetEthernetInterfaces(*c.Server.OriginURIs.System)
	if err != nil {
		beego.Warning("GetEthernetInterfaces() failed, error =", err)
		return err
	}
	if err := c.ServerDBImplement.UpdateEthernetInterfaces(c.Server.ID, ethernet); err != nil {
		beego.Warning("UpdateEthernetInterfaces() failed, error =", err)
	}
	c.DispatchServerUpdate()
	beego.Trace("RefreshEthernetInterfaces() done, server ID =", c.Server.ID)
	return nil
}

// RefreshNetworkInterfaces refresh the network interfaces.
func (s *RackServerRefreshStrategy) RefreshNetworkInterfaces(c *context.RefreshServerContext) error {
	networks, err := c.ServerClient.GetNetworkInterfaces(*c.Server.OriginURIs.System)
	if err != nil {
		beego.Warning("GetNetworkInterfaces() failed, error =", err)
		return err
	}
	if err := c.ServerDBImplement.UpdateNetworkInterfaces(c.Server.ID, networks); err != nil {
		beego.Warning("UpdateNetworkInterfaces() failed, error =", err)
	}
	c.DispatchServerUpdate()
	beego.Trace("RefreshNetworkInterfaces() done, server ID =", c.Server.ID)
	return nil
}

// RefreshStorages refresh the storages.
func (s *RackServerRefreshStrategy) RefreshStorages(c *context.RefreshServerContext) error {
	storages, err := c.ServerClient.GetStorages(*c.Server.OriginURIs.System)
	if err != nil {
		beego.Warning("GetStorages() failed, error =", err)
		return err
	}
	if err := c.ServerDBImplement.UpdateStorages(c.Server.ID, storages); err != nil {
		beego.Warning("UpdateStorages() failed, error =", err)
	}
	c.DispatchServerUpdate()
	beego.Trace("RefreshStorages() done, server ID =", c.Server.ID)
	return nil
}

// RefreshPower refresh the power.
func (s *RackServerRefreshStrategy) RefreshPower(c *context.RefreshServerContext) error {
	power, err := c.ServerClient.GetPower(*c.Server.OriginURIs.Chassis)
	if err != nil {
		beego.Warning("GetPower() failed, error =", err)
		return err
	}
	if err := c.ServerDBImplement.UpdatePower(c.Server.ID, power); err != nil {
		beego.Warning("UpdatePower() failed, error =", err)
	}
	c.DispatchServerUpdate()
	beego.Trace("RefreshPower() done, server ID =", c.Server.ID)
	return nil
}

// RefreshThermal refresh the thermal.
func (s *RackServerRefreshStrategy) RefreshThermal(c *context.RefreshServerContext) error {
	thermal, err := c.ServerClient.GetThermal(*c.Server.OriginURIs.Chassis)
	if err != nil {
		beego.Warning("GetThermal() failed, error =", err)
		return err
	}
	if err := c.ServerDBImplement.UpdateThermal(c.Server.ID, thermal); err != nil {
		beego.Warning("UpdateThermal() failed, error =", err)
	}
	c.DispatchServerUpdate()
	beego.Trace("RefreshThermal() done, server ID =", c.Server.ID)
	return nil
}

// RefreshOemHuaweiBoards refresh the OEM boards.
func (s *RackServerRefreshStrategy) RefreshOemHuaweiBoards(c *context.RefreshServerContext) error {
	boards, err := c.ServerClient.GetOemHuaweiBoards(*c.Server.OriginURIs.Chassis)
	if err != nil {
		beego.Warning("GetOemHuaweiBoards() failed, error =", err)
		return err
	}
	if err := c.ServerDBImplement.UpdateOemHuaweiBoards(c.Server.ID, boards); err != nil {
		beego.Warning("UpdateOemHuaweiBoards() failed, error =", err)
	}
	c.DispatchServerUpdate()
	beego.Trace("RefreshOemHuaweiBoards() done, server ID =", c.Server.ID)
	return nil
}

// RefreshNetworkAdapters refresh the network adapters.
func (s *RackServerRefreshStrategy) RefreshNetworkAdapters(c *context.RefreshServerContext) error {
	networkAdapters, err := c.ServerClient.GetNetworkAdapters(*c.Server.OriginURIs.Chassis)
	if err != nil {
		beego.Warning("GetOemHuaweiBoards() failed, error =", err)
		return err
	}
	if err := c.ServerDBImplement.UpdateNetworkAdapters(c.Server.ID, networkAdapters); err != nil {
		beego.Warning("UpdateOemHuaweiBoards() failed, error =", err)
	}
	c.DispatchServerUpdate()
	beego.Trace("RefreshOemHuaweiBoards() done, server ID =", c.Server.ID)
	return nil
}

// RefreshDrives refresh the drives.
func (s *RackServerRefreshStrategy) RefreshDrives(c *context.RefreshServerContext) error {
	drives, err := c.ServerClient.GetDrives(*c.Server.OriginURIs.Chassis)
	if err != nil {
		beego.Warning("GetDrives() failed, error =", err)
		return err
	}
	if err := c.ServerDBImplement.UpdateDrives(c.Server.ID, drives); err != nil {
		beego.Warning("UpdateDrives() failed, error =", err)
	}
	c.DispatchServerUpdate()
	beego.Trace("RefreshDrives() done, server ID =", c.Server.ID)
	return nil
}

// RefreshPCIeDevices refresh the PCIe devices.
func (s *RackServerRefreshStrategy) RefreshPCIeDevices(c *context.RefreshServerContext) error {
	pcieDevice, err := c.ServerClient.GetPCIeDevices(*c.Server.OriginURIs.Chassis)
	if err != nil {
		beego.Warning("GetPCIeDevices() failed, error =", err)
		return err
	}
	if err := c.ServerDBImplement.UpdatePCIeDevices(c.Server.ID, pcieDevice); err != nil {
		beego.Warning("UpdatePCIeDevices() failed, error =", err)
	}
	c.DispatchServerUpdate()
	beego.Trace("RefreshPCIeDevices() done, server ID =", c.Server.ID)
	return nil
}
