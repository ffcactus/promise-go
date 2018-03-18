package context

import (
	taskSDK "promise/sdk/task"
	wsSDK "promise/sdk/ws"
	serverClient "promise/server/client"
	"promise/server/db"
	"promise/server/object/constvalue"
	serverM "promise/server/object/model"
	taskDto "promise/task/object/dto"
	taskModel "promise/task/object/model"

	log "github.com/sirupsen/logrus"
)

// ServerContextInterface Server context interface.
type ServerContextInterface interface {
	ErrorHandlerInterface
	CredentialHandlerInterface
	TaskHandlerInterface
	serverClient.ServerClientInterface
	db.ServerDBInterface
	ServerIndexInterface
}

// ServerContext Server context.
type ServerContext struct {
	ErrorHandler
	CredentialHandler
	ServerClient serverClient.ServerClientInterface
	db.ServerDBImplement
	ServerIndex
	Server *serverM.Server
}

// CreateServerContext Create server context by server.
func CreateServerContext(server *serverM.Server) *ServerContext {
	var context ServerContext
	context.ServerClient = serverClient.GetServerClient(server)
	context.Server = server
	context.ServerIndex = *CreateServerIndex()
	return &context
}

// UpdateServer Update date the server in the context.
// Many methods need the server model, but it's hard to keep the server in the context
// always updated. The context won't update the server since the context created. Users
// can update the server manually.
func (c *ServerContext) UpdateServer() {
	if c.Server == nil {
		log.Warn("Update server in the context failed, server = nil.")
	}
	c.Server = c.ServerDBImplement.GetServerFull(c.Server.ID)
}

// DispatchServerUpdate Dispatch the server updated event.
func (c *ServerContext) DispatchServerUpdate() {
	c.UpdateServer()
	wsSDK.DispatchServerUpdate(c.Server)
}

// DispatchServerDelete Dispatch the server deleted event.
func (c *ServerContext) DispatchServerDelete() {
	if c.Server == nil {
		log.Warn("Dispatch server in the context failed, server = nil.")
	}
	wsSDK.DispatchServerDelete(constvalue.ToServerURI(c.Server.ID))
}

// CreateTask Create task.
func (c *ServerContext) CreateTask(request *taskDto.PostTaskRequest) {
	log.WithFields(log.Fields{"id": c.Server.ID, "task": request.Name}).Debug("Create server task.")
	taskResp, message, err := taskSDK.CreateTask(request)
	if err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "task": request.Name, "err": err}).Warn("Create server task failed.")
		return
	}
	if message != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "task": request.Name, "message": message[0].ID}).Warn("Create server task failed.")
		log.Warn("Failed to create task, error message return, server ID =", c.Server.ID, "task =", request.Name)
		c.AppendMessages(message)
	}
	log.WithFields(log.Fields{"id": c.Server.ID, "taskName": taskResp.Name, "taskID": taskResp.ID}).Info("Create server task.")
	c.Server.CurrentTask = taskResp.ID
	c.SetServerTask(c.Server.ID, taskResp.ID)
	wsSDK.DispatchServerUpdate(c.Server)
}

// UpdateStepExecutionState Update the step's execution state.
func (c *ServerContext) UpdateStepExecutionState(stepName string, state taskModel.ExecutionState) {
	log.WithFields(log.Fields{"id": c.Server.ID, "task": c.Server.CurrentTask, "step": stepName, "state": state}).Debug("Update step execution state.")
	_, message, err := taskSDK.SetStepExecutionState(c.Server.CurrentTask, stepName, state)
	if err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "task": c.Server.CurrentTask, "step": stepName, "state": state, "err": err}).Warn("Update task step execution state failed.")
		return
	}
	if message != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "task": c.Server.CurrentTask, "step": stepName, "state": state, "message": message[0].ID}).Warn("Update task step execution state failed.")
		c.AppendMessages(message)
	}
}

// UpdateStepExecutionResultState Update the step's execution result state.
func (c *ServerContext) UpdateStepExecutionResultState(stepName string, state taskModel.ExecutionResultState) {
	log.WithFields(log.Fields{"id": c.Server.ID, "task": c.Server.CurrentTask, "step": stepName, "state": state}).Debug("Update step execution result state.")
	_, message, err := taskSDK.SetStepExecutionResultState(c.Server.CurrentTask, stepName, state)
	if err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "task": c.Server.CurrentTask, "step": stepName, "state": state, "err": err}).Warn("Update task step execution result state failed.")
		return
	}
	if message != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "task": c.Server.CurrentTask, "step": stepName, "state": state, "message": message[0].ID}).Warn("Update task step execution result state failed.")
		c.AppendMessages(message)
	}
}

// SetTaskStepRunning Set the task to running.
func (c *ServerContext) SetTaskStepRunning(stepName string) {
	log.WithFields(log.Fields{"id": c.Server.ID, "task": c.Server.CurrentTask, "step": stepName}).Debug("Set task step to running.")
	_, message, err := taskSDK.SetStepExecutionState(c.Server.CurrentTask, stepName, taskModel.ExecutionStateRunning)
	if err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "task": c.Server.CurrentTask, "step": stepName, "err": err}).Warn("Set task step to running failed.")
		return
	}
	if message != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "task": c.Server.CurrentTask, "step": stepName, "message": message[0].ID}).Warn("Set task step to running failed.")
		c.AppendMessages(message)
	}
}

// SetTaskStepFinished Set the task to finished.
func (c *ServerContext) SetTaskStepFinished(stepName string) {
	log.WithFields(log.Fields{"id": c.Server.ID, "task": c.Server.CurrentTask, "step": stepName}).Debug("Set task step to finished.")
	request := new(taskDto.UpdateTaskStepRequest)
	request.Name = stepName
	request.ExecutionState = &taskModel.ExecutionStateTerminated
	request.ExecutionResult = new(taskDto.UpdateExecutionResultRequest)
	request.ExecutionResult.State = &taskModel.ExecutionResultStateFinished
	_, message, err := taskSDK.UpdateStep(c.Server.CurrentTask, request)
	if err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "task": c.Server.CurrentTask, "step": stepName, "err": err}).Warn("Set task step to finished failed.")
		return
	}
	if message != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "task": c.Server.CurrentTask, "step": stepName, "message": message[0].ID}).Warn("Set task step to finished failed.")
		c.AppendMessages(message)
	}
}

// SetTaskStepWarning Set the task step to warning.
func (c *ServerContext) SetTaskStepWarning(stepName string) {
	log.WithFields(log.Fields{"id": c.Server.ID, "task": c.Server.CurrentTask, "step": stepName}).Debug("Set task step to warning.")
	request := new(taskDto.UpdateTaskStepRequest)
	request.Name = stepName
	request.ExecutionState = &taskModel.ExecutionStateTerminated
	request.ExecutionResult = new(taskDto.UpdateExecutionResultRequest)
	request.ExecutionResult.State = &taskModel.ExecutionResultStateWarning
	_, message, err := taskSDK.UpdateStep(c.Server.CurrentTask, request)
	if err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "task": c.Server.CurrentTask, "step": stepName, "err": err}).Warn("Set task step to warning failed.")
		return
	}
	if message != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "task": c.Server.CurrentTask, "step": stepName, "message": message[0].ID}).Warn("Set task step to warning failed.")
		c.AppendMessages(message)
	}
}

// SetTaskStepError Set the task step to error.
func (c *ServerContext) SetTaskStepError(stepName string) {
	log.WithFields(log.Fields{"id": c.Server.ID, "task": c.Server.CurrentTask, "step": stepName}).Debug("Set task step to error.")
	request := new(taskDto.UpdateTaskStepRequest)
	request.Name = stepName
	request.ExecutionState = &taskModel.ExecutionStateTerminated
	request.ExecutionResult = new(taskDto.UpdateExecutionResultRequest)
	request.ExecutionResult.State = &taskModel.ExecutionResultStateError
	log.Debug("SetTaskStepError(), task ID =", c.Server.CurrentTask, "step name =", stepName)
	_, message, err := taskSDK.UpdateStep(c.Server.CurrentTask, request)
	if err != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "task": c.Server.CurrentTask, "step": stepName, "err": err}).Warn("Set task step to error failed.")
		return
	}
	if message != nil {
		log.WithFields(log.Fields{"id": c.Server.ID, "task": c.Server.CurrentTask, "step": stepName, "message": message[0].ID}).Warn("Set task step to error failed.")
		c.AppendMessages(message)
	}
}
