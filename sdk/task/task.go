package task

import (
	"promise/base"
	"promise/task/object/dto"
	"promise/task/object/model"
)

var (
	// TaskServerRoot The root of the service.
	TaskServerRoot = base.RootURL + base.TaskBaseURI + "/"
	client         base.Client
)

func init() {
	client.Protocol = "http"
	client.Addresses = []string{"task"}
	client.Username = "Username"
	client.Password = "Password"
	client.CurrentAddress = "task"
}

// CreateTask Create the task.
func CreateTask(request *dto.PostTaskRequest) (*dto.GetTaskResponse, base.ClientError) {
	response := new(dto.GetTaskResponse)
	err := client.Post(TaskServerRoot, request, response)
	return response, err
}

// UpdateTask will update the task.
func UpdateTask(taskID string, request *dto.UpdateTaskRequest) (*dto.GetTaskResponse, base.ClientError) {
	response := new(dto.GetTaskResponse)
	err := client.Post(TaskServerRoot+taskID+"/action/updatetask", request, response)
	return response, err
}

// UpdateStep will update the task to step and update it's percentage.
func UpdateStep(taskID string, request *dto.UpdateTaskStepRequest) (*dto.GetTaskResponse, base.ClientError) {
	response := new(dto.GetTaskResponse)
	err := client.Post(TaskServerRoot+taskID+"/action/updatetaskstep", request, response)
	return response, err
}

// SetStepExecutionState Set step execution state.
func SetStepExecutionState(taskID string, name string, state model.ExecutionState) (*dto.GetTaskResponse, base.ClientError) {
	request := new(dto.UpdateTaskStepRequest)
	request.Name = name
	request.ExecutionState = &state
	response := new(dto.GetTaskResponse)
	err := client.Post(TaskServerRoot+taskID+"/action/updatetaskstep", request, response)
	return response, err
}

// SetStepExecutionResultState Set step execution result state.
func SetStepExecutionResultState(taskID string, name string, state model.ExecutionResultState) (*dto.GetTaskResponse, base.ClientError) {
	request := new(dto.UpdateTaskStepRequest)
	request.Name = name
	request.ExecutionResult = new(dto.UpdateExecutionResultRequest)
	request.ExecutionResult.State = &state
	response := new(dto.GetTaskResponse)
	err := client.Post(TaskServerRoot+taskID+"/action/updatetaskstep", request, response)
	return response, err
}
