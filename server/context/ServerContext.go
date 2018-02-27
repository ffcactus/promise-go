package context

import (
	taskSDK "promise/sdk/task"
	wsSDK "promise/sdk/ws"
	serverClient "promise/server/client"
	"promise/server/db"
	serverM "promise/server/object/model"
	taskDto "promise/task/object/dto"
	taskModel "promise/task/object/model"

	"github.com/astaxie/beego"
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
		beego.Warning("UpdateServer() failed, server is nil.")
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
		beego.Warning("DispatchServerDelete() failed, server is nil.")
	}
	wsSDK.DispatchServerDelete(c.Server.URI)
}

// CreateTask Create task.
func (c *ServerContext) CreateTask(request *taskDto.PostTaskRequest) {
	taskResp, message, err := taskSDK.CreateTask(request)
	if err != nil {
		beego.Warning("Failed to create task, server ID =", c.Server.ID, "task =", request.Name, "error =", err)
		return
	}
	if message != nil {
		beego.Warning("Failed to create task, error message return, server ID =", c.Server.ID, "task =", request.Name)
		c.AppendMessages(message)
	}
	beego.Debug("Create task for server", "server ID =", c.Server.ID, "Task ID =", taskResp.ID)
	c.Server.CurrentTask = taskResp.ID
	c.SetServerTask(c.Server.ID, taskResp.ID)
	wsSDK.DispatchServerUpdate(c.Server)
}

// UpdateStepExecutionState Update the step's execution state.
func (c *ServerContext) UpdateStepExecutionState(stepName string, state taskModel.ExecutionState) {
	_, message, err := taskSDK.SetStepExecutionState(c.Server.CurrentTask, stepName, state)
	if err != nil {
		beego.Warning("SetStepExecutionState() failed, server ID =", c.Server.ID, "task ID =", c.Server.CurrentTask, "step name =", stepName, "state =", state, "error =", err)
		return
	}
	if message != nil {
		beego.Warning("SetStepExecutionState(), error message return, server ID =", c.Server.ID, "task ID =", c.Server.CurrentTask, "step name =", stepName, "state =", state)
		c.AppendMessages(message)
	}
}

// UpdateStepExecutionResultState Update the step's execution result state.
func (c *ServerContext) UpdateStepExecutionResultState(stepName string, state taskModel.ExecutionResultState) {
	_, message, err := taskSDK.SetStepExecutionResultState(c.Server.CurrentTask, stepName, state)
	if err != nil {
		beego.Warning("SetStepExecutionResultState() failed, server ID =", c.Server.ID, "task ID =", c.Server.CurrentTask, "step name =", stepName, "state =", state, "error =", err)
		return
	}
	if message != nil {
		beego.Warning("SetStepExecutionResultState() failed, error message return, server ID =", c.Server.ID, "task ID =", c.Server.CurrentTask, "step name =", stepName, "state =", state)
		c.AppendMessages(message)
	}
}

// SetTaskStepRunning Set the task to running.
func (c *ServerContext) SetTaskStepRunning(stepName string) {
	beego.Trace("SetTaskStepRunning(), task ID =", c.Server.CurrentTask, "step name =", stepName)
	_, message, err := taskSDK.SetStepExecutionState(c.Server.CurrentTask, stepName, taskModel.ExecutionStateRunning)
	if err != nil {
		beego.Warning("SetTaskStepRunning() failed, server ID =", c.Server.ID, "task ID =", c.Server.CurrentTask, "step name =", stepName, "error =", err)
		return
	}
	if message != nil {
		beego.Warning("SetTaskStepRunning() failed, error message return, server ID =", c.Server.ID, "task ID =", c.Server.CurrentTask, "step name =", stepName)
		c.AppendMessages(message)
	}
}

// SetTaskStepFinished Set the task to finished.
func (c *ServerContext) SetTaskStepFinished(stepName string) {
	request := new(taskDto.UpdateTaskStepRequest)
	request.Name = stepName
	request.ExecutionState = &taskModel.ExecutionStateTerminated
	request.ExecutionResult = new(taskDto.UpdateExecutionResultRequest)
	request.ExecutionResult.State = &taskModel.ExecutionResultStateFinished
	beego.Trace("SetTaskStepFinished(), task ID =", c.Server.CurrentTask, "step name =", stepName)
	_, message, err := taskSDK.UpdateStep(c.Server.CurrentTask, request)
	if err != nil {
		beego.Warning("SetTaskStepFinished() failed, server ID =", c.Server.ID, "task ID =", c.Server.CurrentTask, "error =", err)
		return
	}
	if message != nil {
		beego.Warning("SetTaskStepRunning() failed, error message return, server ID =", c.Server.ID, "task ID =", c.Server.CurrentTask, "step name =", stepName)
		c.AppendMessages(message)
	}
}

// SetTaskStepWarning Set the task step to warning.
func (c *ServerContext) SetTaskStepWarning(stepName string) {
	request := new(taskDto.UpdateTaskStepRequest)
	request.Name = stepName
	request.ExecutionState = &taskModel.ExecutionStateTerminated
	request.ExecutionResult = new(taskDto.UpdateExecutionResultRequest)
	request.ExecutionResult.State = &taskModel.ExecutionResultStateWarning
	beego.Trace("SetTaskStepWarning(), task ID =", c.Server.CurrentTask, "step name =", stepName)
	_, message, err := taskSDK.UpdateStep(c.Server.CurrentTask, request)
	if err != nil {
		beego.Warning("SetTaskStepWarning() failed, server ID =", c.Server.ID, "task ID =", c.Server.CurrentTask, "error =", err)
		return
	}
	if message != nil {
		beego.Warning("SetTaskStepWarning() failed, error message return, server ID =", c.Server.ID, "task ID =", c.Server.CurrentTask, "step name =", stepName)
		c.AppendMessages(message)
	}
}

// SetTaskStepError Set the task step to error.
func (c *ServerContext) SetTaskStepError(stepName string) {
	request := new(taskDto.UpdateTaskStepRequest)
	request.Name = stepName
	request.ExecutionState = &taskModel.ExecutionStateTerminated
	request.ExecutionResult = new(taskDto.UpdateExecutionResultRequest)
	request.ExecutionResult.State = &taskModel.ExecutionResultStateError
	beego.Trace("SetTaskStepError(), task ID =", c.Server.CurrentTask, "step name =", stepName)
	_, message, err := taskSDK.UpdateStep(c.Server.CurrentTask, request)
	if err != nil {
		beego.Warning("SetTaskStepError() failed, server ID =", c.Server.ID, "task ID =", c.Server.CurrentTask, "error =", err)
		return
	}
	if message != nil {
		beego.Warning("SetTaskStepError() failed, error message return, server ID =", c.Server.ID, "task ID =", c.Server.CurrentTask, "step name =", stepName)
		c.AppendMessages(message)
	}
}
