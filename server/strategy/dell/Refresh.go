package dell

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/context"
	"promise/server/object/constvalue"
	"promise/server/object/errorResp"
	"promise/server/object/model"
	"promise/server/strategy/common"
)

// Refresh strategy for HP servers.
type Refresh struct {
	common.Common
}

var executor = make(chan bool, 5)

// Execute will execute all the steps.
func (s *Refresh) Execute(c *context.RefreshServer, server *model.Server) (string, []base.ErrorResponse) {
	select {
	case executor <- true:
		return s.execute(c, server)
	default:
		return "", []base.ErrorResponse{*base.NewErrorResponseBusy()}
	}
}

func (s *Refresh) execute(c *context.RefreshServer, server *model.Server) (string, []base.ErrorResponse) {
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

func (s *Refresh) _execute(taskID string, c *context.RefreshServer, server *model.Server) {
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
	// if err := s.StepWarper(taskID, common.ServerRefreshTaskStepNameBoards, c, server, s.RefreshBoards); err != nil {
	// 	log.WithFields(log.Fields{
	// 		"id":   server.ID,
	// 		"step": common.ServerRefreshTaskStepNameBoards,
	// 	}).Warn("Strategy refresh server step failed.")
	// }

	// Chassis.NetworkAdapters
	// if err := s.StepWarper(taskID, common.ServerRefreshTaskStepNameNetworkAdapters, c, server, s.RefreshNetworkAdapters); err != nil {
	// 	log.WithFields(log.Fields{
	// 		"id":   server.ID,
	// 		"step": common.ServerRefreshTaskStepNameNetworkAdapters,
	// 	}).Warn("Strategy refresh server step failed.")
	// }

	// Chassis.Drives
	// if err := s.StepWarper(taskID, common.ServerRefreshTaskStepNameDrives, c, server, s.RefreshDrives); err != nil {
	// 	log.WithFields(log.Fields{
	// 		"id":   server.ID,
	// 		"step": common.ServerRefreshTaskStepNameDrives,
	// 	}).Warn("Strategy refresh server step failed.")
	// }

	// Chassis.PCIeDevices
	// if err := s.StepWarper(taskID, common.ServerRefreshTaskStepNamePCIeDevices, c, server, s.RefreshPCIeDevices); err != nil {
	// 	log.WithFields(log.Fields{
	// 		"id":   server.ID,
	// 		"step": common.ServerRefreshTaskStepNamePCIeDevices,
	// 	}).Warn("Strategy refresh server step failed.")
	// }

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
	// if err := s.StepWarper(taskID, common.ServerRefreshTaskStepNameEthernetInterfaces, c, server, s.RefreshEthernetInterfaces); err != nil {
	// 	log.WithFields(log.Fields{
	// 		"id":   server.ID,
	// 		"step": common.ServerRefreshTaskStepNameEthernetInterfaces,
	// 	}).Warn("Strategy refresh server step failed.")
	// }

	// ComputerSystem.NetworkInterfaces
	// if err := s.StepWarper(taskID, common.ServerRefreshTaskStepNameNetworkInterfaces, c, server, s.RefreshNetworkInterfaces); err != nil {
	// 	log.WithFields(log.Fields{
	// 		"id":   server.ID,
	// 		"step": common.ServerRefreshTaskStepNameNetworkInterfaces,
	// 	}).Warn("Strategy refresh server step failed.")
	// }

	// ComputerSystem.Storages
	// if err := s.StepWarper(taskID, common.ServerRefreshTaskStepNameStorages, c, server, s.RefreshStorages); err != nil {
	// 	log.WithFields(log.Fields{
	// 		"id":   server.ID,
	// 		"step": common.ServerRefreshTaskStepNameStorages,
	// 	}).Warn("Strategy refresh server step failed.")
	// }

	s.SetServerHealth(&c.Base, server, constvalue.ServerHealthOK)

	s.SetState(&c.Base, server, constvalue.ServerStateReady)
	log.WithFields(log.Fields{
		"id": server.ID,
	}).Info("Strategy refresh server done.")
	<-executor
}

// Stepfunc is the func point.
type Stepfunc func(c *context.RefreshServer, server *model.Server) error

// StepWarper is a warper.
func (s *Refresh) StepWarper(
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
func (s *Refresh) RefreshProcessors(c *context.RefreshServer, server *model.Server) error {
	var component = "Processor"
	processors, err := c.ServerClient.GetProcessors(server.OriginURIs.System)
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
func (s *Refresh) RefreshMemory(c *context.RefreshServer, server *model.Server) error {
	var component = "Memory"
	memory, err := c.ServerClient.GetMemory(server.OriginURIs.System)
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
func (s *Refresh) RefreshEthernetInterfaces(c *context.RefreshServer, server *model.Server) error {
	var component = "EthernetInterfaces"
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Debug("Refresh server component done.")
	return nil
}

// RefreshNetworkInterfaces refresh the network interfaces.
func (s *Refresh) RefreshNetworkInterfaces(c *context.RefreshServer, server *model.Server) error {
	var component = "NetworkInterfaces"
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Debug("Refresh server component done.")
	return nil
}

// RefreshStorages refresh the storages.
func (s *Refresh) RefreshStorages(c *context.RefreshServer, server *model.Server) error {
	var component = "Storages"
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Debug("Refresh server component done.")
	return nil
}

// RefreshPower refresh the power.
func (s *Refresh) RefreshPower(c *context.RefreshServer, server *model.Server) error {
	var component = "Power"
	power, err := c.ServerClient.GetPower(server.OriginURIs.Chassis)
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
func (s *Refresh) RefreshThermal(c *context.RefreshServer, server *model.Server) error {
	var component = "Thermal"
	thermal, err := c.ServerClient.GetThermal(server.OriginURIs.Chassis)
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

// RefreshBoards refresh the OEM boards.
func (s *Refresh) RefreshBoards(c *context.RefreshServer, server *model.Server) error {
	var component = "Boards"
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Debug("Refresh server component done.")
	return nil
}

// RefreshNetworkAdapters refresh the network adapters.
func (s *Refresh) RefreshNetworkAdapters(c *context.RefreshServer, server *model.Server) error {
	var component = "NetworkAdapters"
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Debug("Refresh server component done.")
	return nil
}

// RefreshDrives refresh the drives.
func (s *Refresh) RefreshDrives(c *context.RefreshServer, server *model.Server) error {
	var component = "Drives"
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Debug("Refresh server component done.")
	return nil
}

// RefreshPCIeDevices refresh the PCIe devices.
func (s *Refresh) RefreshPCIeDevices(c *context.RefreshServer, server *model.Server) error {
	var component = "PCIeDevices"
	log.WithFields(log.Fields{"id": server.ID, "component": component}).Debug("Refresh server component done.")
	return nil
}
*/
