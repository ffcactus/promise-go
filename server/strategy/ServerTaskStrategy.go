package strategy

import (
	"fmt"
	taskDTO "promise/task/object/dto"
	taskModel "promise/task/object/model"
	commonDTO "promise/common/object/dto"
	"promise/server/context"
	"promise/common/object/constValue"
	"promise/server/object/model"
	taskSDK "promise/sdk/task"
	log "github.com/sirupsen/logrus"
)


func init() {
	refreshTaskTotalTime = 0
	for i := range ServerTaskRefreshStepLIST {
		refreshTaskTotalTime += ServerTaskRefreshStepLIST[i].ExpectedExecutionMs
	}
}

var (
	refreshTaskTotalTime int

	// ServerTaskRefresh is an ID.
	ServerTaskRefresh = "Server.Task.Refresh"

	// ServerRefreshTaskStepIDPower is an ID.
	ServerRefreshTaskStepIDPower = "Server.Refresh.Power"
	// ServerRefreshTaskStepIDThermal is an ID.
	ServerRefreshTaskStepIDThermal = "Server.Refresh.Thermal"
	// ServerRefreshTaskStepIDBoards is an ID.
	ServerRefreshTaskStepIDBoards = "Server.Refresh.OemHuaweiBoards"
	// ServerRefreshTaskStepIDNetworkAdapters is an ID.
	ServerRefreshTaskStepIDNetworkAdapters = "Server.Refresh.NetworkAdapters"
	// ServerRefreshTaskStepIDDrives is an ID.
	ServerRefreshTaskStepIDDrives = "Server.Refresh.Drives"
	// ServerRefreshTaskStepIDPCIeDevices is an ID.
	ServerRefreshTaskStepIDPCIeDevices = "Server.Refresh.PCIeDevices"
	// ServerRefreshTaskStepIDProcessors is an ID.
	ServerRefreshTaskStepIDProcessors = "Server.Refresh.Processors"
	// ServerRefreshTaskStepIDMemory is an ID.
	ServerRefreshTaskStepIDMemory = "Server.Refresh.Memory"
	// ServerRefreshTaskStepIDEthernetInterfaces is an ID.
	ServerRefreshTaskStepIDEthernetInterfaces = "Server.Refresh.EthernetInterfaces"
	// ServerRefreshTaskStepIDNetworkInterfaces is an ID.
	ServerRefreshTaskStepIDNetworkInterfaces = "Server.Refresh.NetworkInterfaces"
	// ServerRefreshTaskStepIDStorages is an ID.
	ServerRefreshTaskStepIDStorages = "Server.Refresh.Storages"

	// ServerRefreshTaskStepNamePower is task step name.
	ServerRefreshTaskStepNamePower = "Power"
	// ServerRefreshTaskStepNameThermal is task step name.
	ServerRefreshTaskStepNameThermal = "Thermal"
	// ServerRefreshTaskStepNameBoards is task step name.
	ServerRefreshTaskStepNameBoards = "OemHuaweiBoards"
	// ServerRefreshTaskStepNameNetworkAdapters is task step name.
	ServerRefreshTaskStepNameNetworkAdapters = "NetworkAdapters"
	// ServerRefreshTaskStepNameDrives is task step name.
	ServerRefreshTaskStepNameDrives = "Drives"
	// ServerRefreshTaskStepNamePCIeDevices is task step name.
	ServerRefreshTaskStepNamePCIeDevices = "PCIeDevices"
	// ServerRefreshTaskStepNameProcessors is task step name.
	ServerRefreshTaskStepNameProcessors = "Processors"
	// ServerRefreshTaskStepNameMemory is task step name.
	ServerRefreshTaskStepNameMemory = "Memory"
	// ServerRefreshTaskStepNameEthernetInterfaces is task step name.
	ServerRefreshTaskStepNameEthernetInterfaces = "EthernetInterfaces"
	// ServerRefreshTaskStepNameNetworkInterfaces is task step name.
	ServerRefreshTaskStepNameNetworkInterfaces = "NetworkInterfaces"
	// ServerRefreshTaskStepNameStorages is task step name.
	ServerRefreshTaskStepNameStorages = "Storages"

	// ServerTaskRefreshStepPower is a task step.
	ServerTaskRefreshStepPower = taskDTO.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDPower,
		Name:                ServerRefreshTaskStepNamePower,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepThermal is a task step.
	ServerTaskRefreshStepThermal = taskDTO.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDThermal,
		Name:                ServerRefreshTaskStepNameThermal,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepBoards is a task step.
	ServerTaskRefreshStepBoards = taskDTO.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDBoards,
		Name:                ServerRefreshTaskStepNameBoards,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepNetworkAdapters is a task step.
	ServerTaskRefreshStepNetworkAdapters = taskDTO.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDNetworkAdapters,
		Name:                ServerRefreshTaskStepNameNetworkAdapters,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepDrives is a task step.
	ServerTaskRefreshStepDrives = taskDTO.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDDrives,
		Name:                ServerRefreshTaskStepNameDrives,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepPCIeDevices is a task step.
	ServerTaskRefreshStepPCIeDevices = taskDTO.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDPCIeDevices,
		Name:                ServerRefreshTaskStepNamePCIeDevices,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepProcessors is a task step.
	ServerTaskRefreshStepProcessors = taskDTO.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDProcessors,
		Name:                ServerRefreshTaskStepNameProcessors,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepMemory is a task step.
	ServerTaskRefreshStepMemory = taskDTO.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDMemory,
		Name:                ServerRefreshTaskStepNameMemory,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepEthernetInterfaces is a task step.
	ServerTaskRefreshStepEthernetInterfaces = taskDTO.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDEthernetInterfaces,
		Name:                ServerRefreshTaskStepNameEthernetInterfaces,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepNetworkInterfaces is a task step.
	ServerTaskRefreshStepNetworkInterfaces = taskDTO.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDNetworkInterfaces,
		Name:                ServerRefreshTaskStepNameNetworkInterfaces,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepStorages is a task step.
	ServerTaskRefreshStepStorages = taskDTO.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDStorages,
		Name:                ServerRefreshTaskStepNameStorages,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepLIST is a task step list.
	ServerTaskRefreshStepLIST = []taskDTO.PostTaskStepRequest{
		ServerTaskRefreshStepPower,
		ServerTaskRefreshStepThermal,
		ServerTaskRefreshStepBoards,
		ServerTaskRefreshStepNetworkAdapters,
		ServerTaskRefreshStepDrives,
		ServerTaskRefreshStepPCIeDevices,
		ServerTaskRefreshStepProcessors,
		ServerTaskRefreshStepMemory,
		ServerTaskRefreshStepEthernetInterfaces,
		ServerTaskRefreshStepNetworkInterfaces,
		ServerTaskRefreshStepStorages,
	}
)

// ServerTaskStrategy is the server task strategy implementation.
type ServerTaskStrategy struct {

}

// createRefreshTaskRequest will return a refresh server task request.
func createRefreshTaskRequest(server *model.Server) *taskDTO.PostTaskRequest {
	var request taskDTO.PostTaskRequest

	request.MessageID = &ServerTaskRefresh
	request.Name = "Refresh Server"
	request.Description = "Refresh server resources and re-configure it."
	request.CreatedByName = "CreatedByName???"
	request.CreatedByURI = "CreatedByURI???"
	request.TargetName = server.Name
	request.TargetURI = constValue.ToServerURI(server.ID)
	request.TaskSteps = ServerTaskRefreshStepLIST
	return &request
}

// createTask creates the task.
func (s *ServerTaskStrategy) createTask (request *taskDTO.PostTaskRequest, server *model.Server) (string, error) {
	taskResp, message, err := taskSDK.CreateTask(request)
	if err != nil {
		log.WithFields(log.Fields{
			"server": server.ID, 
			"name": request.Name, 
			"err": err}).
			Warn("Create server task failed.")
		return "", err
	}
	if message != nil {
		log.WithFields(log.Fields{
			"server": server.ID, 
			"name": request.Name, 
			"message": message[0].ID}).
			Warn("Create server task failed.")
		return "", fmt.Errorf("create task failed")
	}
	log.WithFields(log.Fields{
		"server": server.ID, 
		"name": taskResp.Name, 
		"task": taskResp.ID}).
		Info("Create server task.")
	return taskResp.ID, nil
}

// CreateRefreshServerTask will create a server refresh task.
func (s *ServerTaskStrategy) CreateRefreshServerTask(c *context.ServerContext, server *model.Server) (string, error) {
	return s.createTask(createRefreshTaskRequest(server), server)
}

// UpdateStepExecutionState Update the step's execution state.
func (s *ServerTaskStrategy) UpdateStepExecutionState(id string, stepName string, state taskModel.ExecutionState, server *model.Server) {
	_, message, err := taskSDK.SetStepExecutionState(id, stepName, state)
	if err != nil {
		log.WithFields(log.Fields{
			"server": server.ID, 
			"task": id, 
			"step": stepName, 
			"state": state, 
			"error": err}).Warn("Update task step execution state failed.")
	}
	if message != nil {
		log.WithFields(log.Fields{
			"server": server.ID, 
			"task": id, 
			"step": stepName, 
			"state": state, 
			"message": message[0].ID}).
			Warn("Update task step execution state failed.")
	}
}

// UpdateStepExecutionResultState Update the step's execution result state.
func (s *ServerTaskStrategy) UpdateStepExecutionResultState(c *context.ServerContext, id string, stepName string, state taskModel.ExecutionResultState, server *model.Server) {
	_, message, err := taskSDK.SetStepExecutionResultState(id, stepName, state)
	if err != nil {
		log.WithFields(log.Fields{
			"server": server.ID, 
			"task": id, 
			"step": stepName, 
			"state": state, 
			"error": err}).
			Warn("Update task step execution result state failed.")
	}
	if message != nil {
		log.WithFields(log.Fields{
			"server": server.ID, 
			"task": id, 
			"step": stepName, 
			"state": state, 
			"message": message[0].ID}).
			Warn("Update task step execution result state failed.")
	}
}

// SetTaskStepRunning Set the task to running.
func (s *ServerTaskStrategy) SetTaskStepRunning(c *context.ServerContext, id string, stepName string, server *model.Server) {
	log.WithFields(log.Fields{
		"server": server.ID, 
		"task": id, 
		"step": stepName}).
		Debug("Set task step to running.")
	_, message, err := taskSDK.SetStepExecutionState(id, stepName, taskModel.ExecutionStateRunning)
	if err != nil {
		log.WithFields(log.Fields{
			"server": server.ID, 
			"task": id, 
			"step": stepName, 
			"error": err}).
			Warn("Set task step to running failed.")
	}
	if message != nil {
		log.WithFields(log.Fields{
			"server": server.ID, 
			"task": id, 
			"step": stepName, 
			"message": message[0].ID}).
			Warn("Set task step to running failed.")
	}
}

func (s *ServerTaskStrategy) logUpdateStepResult(c *context.ServerContext, id string, stepName string, server *model.Server, err error, message []commonDTO.Message) {
	if err != nil {
		log.WithFields(log.Fields{
			"server": server.ID, 
			"task": id, 
			"step": stepName, 
			"error": err}).
			Warn("Set task step to finished failed.")
	}
	if message != nil {
		log.WithFields(log.Fields{
			"server": server.ID, 
			"task": id, 
			"step": stepName, 
			"message": message[0].ID}).
			Warn("Set task step to finished failed.")
	}
}

// SetTaskStepFinished Set the task to finished.
func (s *ServerTaskStrategy) SetTaskStepFinished(c *context.ServerContext, id string, stepName string, server *model.Server) {
	request := new(taskDTO.UpdateTaskStepRequest)
	request.Name = stepName
	request.ExecutionState = &taskModel.ExecutionStateTerminated
	request.ExecutionResult = new(taskDTO.UpdateExecutionResultRequest)
	request.ExecutionResult.State = &taskModel.ExecutionResultStateFinished
	_, message, err := taskSDK.UpdateStep(id, request)
	s.logUpdateStepResult(c, id, stepName, server, err, message)
}

// SetTaskStepWarning Set the task step to warning.
func (s *ServerTaskStrategy) SetTaskStepWarning(c *context.ServerContext, id string, stepName string, server *model.Server) {
	request := new(taskDTO.UpdateTaskStepRequest)
	request.Name = stepName
	request.ExecutionState = &taskModel.ExecutionStateTerminated
	request.ExecutionResult = new(taskDTO.UpdateExecutionResultRequest)
	request.ExecutionResult.State = &taskModel.ExecutionResultStateWarning
	_, message, err := taskSDK.UpdateStep(id, request)
	s.logUpdateStepResult(c, id, stepName, server, err, message)
}

// SetTaskStepError Set the task step to error.
func (s *ServerTaskStrategy) SetTaskStepError(c *context.ServerContext, id string, stepName string, server *model.Server) {
	log.WithFields(log.Fields{"server": server.ID, "task": id, "step": stepName}).Debug("Set task step to error.")
	request := new(taskDTO.UpdateTaskStepRequest)
	request.Name = stepName
	request.ExecutionState = &taskModel.ExecutionStateTerminated
	request.ExecutionResult = new(taskDTO.UpdateExecutionResultRequest)
	request.ExecutionResult.State = &taskModel.ExecutionResultStateError
	_, message, err := taskSDK.UpdateStep(id, request)
	s.logUpdateStepResult(c, id, stepName, server, err, message)
}
