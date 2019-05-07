package huawei

import (
	log "github.com/sirupsen/logrus"
	"math/rand"
	"promise/base"
	"promise/server/context"
	"promise/server/object/constvalue"
	"promise/server/object/errorResp"
	"promise/server/object/model"
	"promise/server/strategy/common"
)

// RefreshRackServer is the strategy of rack server refresh.
type RefreshRackServer struct {
	common.Common
}

var executor = make(chan bool, 5)

// Execute will execute all the steps.
func (s *RefreshRackServer) Execute(c *context.RefreshServer, server *model.Server) (string, []base.ErrorResponse) {
	select {
	case executor <- true:
		return s.execute(c, server)
	default:
		return "", []base.ErrorResponse{*base.NewErrorResponseBusy()}
	}
}

func (s *RefreshRackServer) execute(c *context.RefreshServer, server *model.Server) (string, []base.ErrorResponse) {
	log.WithFields(log.Fields{"id": server.ID}).Info("Strategy refresh server start.")
	// Lock server.
	if errorResp := s.LockServer(&c.Base, server); errorResp != nil {
		<-executor
		return "", []base.ErrorResponse{*errorResp}
	}

	taskID, err := s.CreateRefreshServerTask(&c.Base, server)
	if err != nil {
		<-executor
		return "", []base.ErrorResponse{*errorResp.NewErrorResponseServerRefreshTaskFailed()}
	}

	go s._execute(taskID, c, server)
	return taskID, nil
}

func (s *RefreshRackServer) _execute(taskID string, c *context.RefreshServer, server *model.Server) {
	// Chassis.Power
	if err := s.StepWarper(taskID, common.ServerRefreshTaskStepNamePower, c, server, s.RefreshPower); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": common.ServerRefreshTaskStepNamePower,
		}).Warn("Strategy refresh server step failed.")
	}
	// Chassis.Thermal
	if err := s.StepWarper(taskID, common.ServerRefreshTaskStepNameThermal, c, server, s.RefreshThermal); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": common.ServerRefreshTaskStepNameThermal,
		}).Warn("Strategy refresh server step failed.")
	}
	// Chassis.Boards
	if err := s.StepWarper(taskID, common.ServerRefreshTaskStepNameBoards, c, server, s.RefreshBoards); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": common.ServerRefreshTaskStepNameBoards,
		}).Warn("Strategy refresh server step failed.")
	}
	// Chassis.NetworkAdapters
	if err := s.StepWarper(taskID, common.ServerRefreshTaskStepNameNetworkAdapters, c, server, s.RefreshNetworkAdapters); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": common.ServerRefreshTaskStepNameNetworkAdapters,
		}).Warn("Strategy refresh server step failed.")
	}
	// Chassis.Drives
	if err := s.StepWarper(taskID, common.ServerRefreshTaskStepNameDrives, c, server, s.RefreshDrives); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": common.ServerRefreshTaskStepNameDrives,
		}).Warn("Strategy refresh server step failed.")
	}
	// Chassis.PCIeDevices
	if err := s.StepWarper(taskID, common.ServerRefreshTaskStepNamePCIeDevices, c, server, s.RefreshPCIeDevices); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": common.ServerRefreshTaskStepNamePCIeDevices,
		}).Warn("Strategy refresh server step failed.")
	}
	// ComputerSystem.Processors
	if err := s.StepWarper(taskID, common.ServerRefreshTaskStepNameProcessors, c, server, s.RefreshProcessors); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": common.ServerRefreshTaskStepNameProcessors,
		}).Warn("Strategy refresh server step failed.")
	}
	// ComputerSystem.Memory
	if err := s.StepWarper(taskID, common.ServerRefreshTaskStepNameMemory, c, server, s.RefreshMemory); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": common.ServerRefreshTaskStepNameMemory,
		}).Warn("Strategy refresh server step failed.")
	}
	// ComputerSystem.EthernetInterfaces
	if err := s.StepWarper(taskID, common.ServerRefreshTaskStepNameEthernetInterfaces, c, server, s.RefreshEthernetInterfaces); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": common.ServerRefreshTaskStepNameEthernetInterfaces,
		}).Warn("Strategy refresh server step failed.")
	}
	// ComputerSystem.NetworkInterfaces
	if err := s.StepWarper(taskID, common.ServerRefreshTaskStepNameNetworkInterfaces, c, server, s.RefreshNetworkInterfaces); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": common.ServerRefreshTaskStepNameNetworkInterfaces,
		}).Warn("Strategy refresh server step failed.")
	}
	// ComputerSystem.Storages
	if err := s.StepWarper(taskID, common.ServerRefreshTaskStepNameStorages, c, server, s.RefreshStorages); err != nil {
		log.WithFields(log.Fields{
			"id":   server.ID,
			"step": common.ServerRefreshTaskStepNameStorages,
		}).Warn("Strategy refresh server step failed.")
	}
	// s.SetServerHealth(&c.Base, server, constvalue.ServerHealthOK)
	s.SetServerHealth(&c.Base, server, randHealth())
	s.SetState(&c.Base, server, constvalue.ServerStateReady)
	log.WithFields(log.Fields{
		"id": server.ID,
	}).Info("Strategy refresh server done.")
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

/*
// RefreshProcessors refresh the processor of the server.
func (s *RefreshRackServer) RefreshProcessors(c *context.RefreshServer, server *model.Server) error {
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
	log.WithFields(log.Fields{"id": server.ID, "component": component, "count": updatedServer.ComputerSystem.Processors.Count}).Debug("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshMemory refresh the memory
func (s *RefreshRackServer) RefreshMemory(c *context.RefreshServer, server *model.Server) error {
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
	log.WithFields(log.Fields{"id": server.ID, "component": component, "count": updatedServer.ComputerSystem.Memory.Count}).Debug("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshEthernetInterfaces refresh the ethernet interfaces.
func (s *RefreshRackServer) RefreshEthernetInterfaces(c *context.RefreshServer, server *model.Server) error {
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
	log.WithFields(log.Fields{"id": server.ID, "component": component, "count": updatedServer.ComputerSystem.EthernetInterfaces.Count}).Debug("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshNetworkInterfaces refresh the network interfaces.
func (s *RefreshRackServer) RefreshNetworkInterfaces(c *context.RefreshServer, server *model.Server) error {
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
	log.WithFields(log.Fields{"id": server.ID, "component": component, "count": updatedServer.ComputerSystem.NetworkInterfaces.Count}).Debug("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshStorages refresh the storages.
func (s *RefreshRackServer) RefreshStorages(c *context.RefreshServer, server *model.Server) error {
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
	log.WithFields(log.Fields{"id": server.ID, "component": component, "count": updatedServer.ComputerSystem.Storages.Count}).Debug("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshPower refresh the power.
func (s *RefreshRackServer) RefreshPower(c *context.RefreshServer, server *model.Server) error {
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
	log.WithFields(log.Fields{"id": server.ID, "component": component, "power": updatedServer.Chassis.Power}).Debug("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshThermal refresh the thermal.
func (s *RefreshRackServer) RefreshThermal(c *context.RefreshServer, server *model.Server) error {
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

	log.WithFields(log.Fields{"id": server.ID, "component": component, "thermal": updatedServer.Chassis.Thermal}).Debug("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshBoards refresh the boards.
func (s *RefreshRackServer) RefreshBoards(c *context.RefreshServer, server *model.Server) error {
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
	log.WithFields(log.Fields{"id": server.ID, "component": component, "count": updatedServer.Chassis.Boards.Count}).Debug("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshNetworkAdapters refresh the network adapters.
func (s *RefreshRackServer) RefreshNetworkAdapters(c *context.RefreshServer, server *model.Server) error {
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
	log.WithFields(log.Fields{"id": server.ID, "component": component, "count": updatedServer.Chassis.NetworkAdapters.Count}).Debug("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshDrives refresh the drives.
func (s *RefreshRackServer) RefreshDrives(c *context.RefreshServer, server *model.Server) error {
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
	log.WithFields(log.Fields{"id": server.ID, "component": component, "count": updatedServer.Chassis.Drives.Count}).Debug("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}

// RefreshPCIeDevices refresh the PCIe devices.
func (s *RefreshRackServer) RefreshPCIeDevices(c *context.RefreshServer, server *model.Server) error {
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
	log.WithFields(log.Fields{"id": server.ID, "component": component, "count": updatedServer.Chassis.PCIeDevices.Count}).Debug("Strategy refresh server component done.")
	s.DispatchServerUpdate(&c.Base, updatedServer)
	return nil
}
*/
