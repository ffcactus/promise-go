package strategy

import (
	log "github.com/sirupsen/logrus"
	"math/rand"
	"promise/base"
	"promise/server/context"
	"promise/server/object/constvalue"
	"promise/server/object/errorResp"
	"promise/server/object/model"
)

// RefreshRackServer is the strategy of rack server refresh.
type RefreshRackServer struct {
	Base
	ServerEvent
	ServerTask
}

var executor = make(chan bool, 5)

// Execute will execute all the steps.
func (s *RefreshRackServer) Execute(c *context.RefreshServer, server *model.Server) (*string, []base.ErrorResponse) {
	select {
	case executor <- true:
		return s.execute(c, server)
	default:
		return nil, []base.ErrorResponse{*base.NewErrorResponseBusy()}
	}
}

func (s *RefreshRackServer) execute(c *context.RefreshServer, server *model.Server) (*string, []base.ErrorResponse) {
	log.WithFields(log.Fields{"id": server.ID}).Info("Refresh server.")
	// Lock server.
	if errorResp := s.LockServer(&c.Base, server); errorResp != nil {
		<-executor
		return nil, []base.ErrorResponse{*errorResp}
	}

	taskID, err := s.CreateRefreshServerTask(&c.Base, server)
	if err != nil {
		<-executor
		return nil, []base.ErrorResponse{*errorResp.NewErrorResponseServerRefreshTaskFailed()}
	}

	go s._execute(taskID, c, server)
	return &taskID, nil
}

func (s *RefreshRackServer) _execute(taskID string, c *context.RefreshServer, server *model.Server) {
	// Chassis.Power
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNamePower, c, server, s.RefreshPower); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": ServerRefreshTaskStepNamePower,
		}).Info("Strategy refresh server step failed.")
	}
	// Chassis.Thermal
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNameThermal, c, server, s.RefreshThermal); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": ServerRefreshTaskStepNameThermal,
		}).Info("Strategy refresh server step failed.")
	}
	// Chassis.OemHuaweiBoards
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNameBoards, c, server, s.RefreshOemHuaweiBoards); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": ServerRefreshTaskStepNameBoards,
		}).Info("Strategy refresh server step failed.")
	}
	// Chassis.NetworkAdapters
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNameNetworkAdapters, c, server, s.RefreshNetworkAdapters); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": ServerRefreshTaskStepNameNetworkAdapters,
		}).Info("Strategy refresh server step failed.")
	}
	// Chassis.Drives
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNameDrives, c, server, s.RefreshDrives); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": ServerRefreshTaskStepNameDrives,
		}).Info("Strategy refresh server step failed.")
	}
	// Chassis.PCIeDevices
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNamePCIeDevices, c, server, s.RefreshPCIeDevices); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": ServerRefreshTaskStepNamePCIeDevices,
		}).Info("Strategy refresh server step failed.")
	}
	// ComputerSystem.Processors
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNameProcessors, c, server, s.RefreshProcessors); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": ServerRefreshTaskStepNameProcessors,
		}).Info("Strategy refresh server step failed.")
	}
	// ComputerSystem.Memory
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNameMemory, c, server, s.RefreshMemory); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": ServerRefreshTaskStepNameMemory,
		}).Info("Strategy refresh server step failed.")
	}
	// ComputerSystem.EthernetInterfaces
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNameEthernetInterfaces, c, server, s.RefreshEthernetInterfaces); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": ServerRefreshTaskStepNameEthernetInterfaces,
		}).Info("Strategy refresh server step failed.")
	}
	// ComputerSystem.NetworkInterfaces
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNameNetworkInterfaces, c, server, s.RefreshNetworkInterfaces); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": ServerRefreshTaskStepNameNetworkInterfaces,
		}).Info("Strategy refresh server step failed.")
	}
	// ComputerSystem.Storages
	if err := s.StepWarper(taskID, ServerRefreshTaskStepNameStorages, c, server, s.RefreshStorages); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": ServerRefreshTaskStepNameStorages,
		}).Info("Strategy refresh server step failed.")
	}
	// s.SetServerHealth(&c.Base, server, constvalue.ServerHealthOK)
	s.SetServerHealth(&c.Base, server, randHealth())
	s.SetServerState(&c.Base, server, constvalue.ServerStateReady)
	log.WithFields(log.Fields{
		"id": server.ID,
	}).Info("Refresh server done.")
	<-executor
}

var health = []string{
	"OK",
	"Warning",
	"Critical",
}

func randHealth() string {
	return health[rand.Intn(len(health))]
}

// Stepfunc is the func point.
type Stepfunc func(c *context.RefreshServer, server *model.Server) error

// StepWarper is a warper.
func (s *RefreshRackServer) StepWarper(
	id string,
	stepName string,
	c *context.RefreshServer,
	server *model.Server,
	stepfunc Stepfunc) error {
	s.SetTaskStepRunning(&c.Base, id, stepName, server)
	err := stepfunc(c, server)
	if err != nil {
		s.SetTaskStepError(&c.Base, id, stepName, server)
	} else {
		s.SetTaskStepFinished(&c.Base, id, stepName, server)
	}
	return err
}

// RefreshProcessors refresh the processor of the server.
func (s *RefreshRackServer) RefreshProcessors(c *context.RefreshServer, server *model.Server) error {
	var component = "Processor"
	processors, err := c.ServerClient.GetProcessors(*server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdateProcessors(server.ID, processors)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Update server component failed.")
		return err
	}
	s.DispatchServerUpdate(&c.Base, updatedServer)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Debug("Refresh server component done.")
	return nil
}

// RefreshMemory refresh the memory
func (s *RefreshRackServer) RefreshMemory(c *context.RefreshServer, server *model.Server) error {
	var component = "Memory"
	memory, err := c.ServerClient.GetMemory(*server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdateMemory(server.ID, memory)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.Base, updatedServer)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Debug("Refresh server component done.")
	return nil
}

// RefreshEthernetInterfaces refresh the ethernet interfaces.
func (s *RefreshRackServer) RefreshEthernetInterfaces(c *context.RefreshServer, server *model.Server) error {
	var component = "EthernetInterfaces"
	ethernet, err := c.ServerClient.GetEthernetInterfaces(*server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdateEthernetInterfaces(server.ID, ethernet)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.Base, updatedServer)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Debug("Refresh server component done.")
	return nil
}

// RefreshNetworkInterfaces refresh the network interfaces.
func (s *RefreshRackServer) RefreshNetworkInterfaces(c *context.RefreshServer, server *model.Server) error {
	var component = "NetworkInterfaces"
	networks, err := c.ServerClient.GetNetworkInterfaces(*server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdateNetworkInterfaces(server.ID, networks)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.Base, updatedServer)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Debug("Refresh server component done.")
	return nil
}

// RefreshStorages refresh the storages.
func (s *RefreshRackServer) RefreshStorages(c *context.RefreshServer, server *model.Server) error {
	var component = "Storages"
	storages, err := c.ServerClient.GetStorages(*server.OriginURIs.System)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdateStorages(server.ID, storages)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.Base, updatedServer)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Debug("Refresh server component done.")
	return nil
}

// RefreshPower refresh the power.
func (s *RefreshRackServer) RefreshPower(c *context.RefreshServer, server *model.Server) error {
	var component = "Power"
	power, err := c.ServerClient.GetPower(*server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdatePower(server.ID, power)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.Base, updatedServer)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Debug("Refresh server component done.")
	return nil
}

// RefreshThermal refresh the thermal.
func (s *RefreshRackServer) RefreshThermal(c *context.RefreshServer, server *model.Server) error {
	var component = "Thermal"
	thermal, err := c.ServerClient.GetThermal(*server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdateThermal(server.ID, thermal)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.Base, updatedServer)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Debug("Refresh server component done.")
	return nil
}

// RefreshOemHuaweiBoards refresh the OEM boards.
func (s *RefreshRackServer) RefreshOemHuaweiBoards(c *context.RefreshServer, server *model.Server) error {
	var component = "OemHuaweiBoards"
	boards, err := c.ServerClient.GetOemHuaweiBoards(*server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdateOemHuaweiBoards(server.ID, boards)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.Base, updatedServer)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Debug("Refresh server component done.")
	return nil
}

// RefreshNetworkAdapters refresh the network adapters.
func (s *RefreshRackServer) RefreshNetworkAdapters(c *context.RefreshServer, server *model.Server) error {
	var component = "NetworkAdapters"
	networkAdapters, err := c.ServerClient.GetNetworkAdapters(*server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdateNetworkAdapters(server.ID, networkAdapters)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.Base, updatedServer)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Debug("Refresh server component done.")
	return nil
}

// RefreshDrives refresh the drives.
func (s *RefreshRackServer) RefreshDrives(c *context.RefreshServer, server *model.Server) error {
	var component = "Drives"
	drives, err := c.ServerClient.GetDrives(*server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdateDrives(server.ID, drives)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.Base, updatedServer)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Debug("Refresh server component done.")
	return nil
}

// RefreshPCIeDevices refresh the PCIe devices.
func (s *RefreshRackServer) RefreshPCIeDevices(c *context.RefreshServer, server *model.Server) error {
	var component = "PCIeDevices"
	pcieDevice, err := c.ServerClient.GetPCIeDevices(*server.OriginURIs.Chassis)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Get server component failed.")
		return err
	}
	updatedServer, err := c.DB.UpdatePCIeDevices(server.ID, pcieDevice)
	if err != nil {
		log.WithFields(log.Fields{"id": server.ID, "component": component, "error": err}).Warn("Update server component failed.")
	}
	s.DispatchServerUpdate(&c.Base, updatedServer)
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Debug("Refresh server component done.")
	return nil
}
