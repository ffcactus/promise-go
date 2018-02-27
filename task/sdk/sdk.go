package sdk

import (
	"net/http"
	commonDto "promise/common/object/dto"
	"promise/common/util/localClient"
	"promise/task/object/dto"
	"promise/task/object/model"
)

var (
	// TaskServerRoot The root of the service.
	TaskServerRoot = "/promise/v1/task"
)

// TaskClient Task client.
type TaskClient struct {
	localClient *localClient.LocalClient
}

// CreateTaskClient Create the task client.
func CreateTaskClient() *TaskClient {
	client := new(TaskClient)
	client.localClient = localClient.Instance()
	return client
}

// CreateTask Create the task.
func (c *TaskClient) CreateTask(task *dto.PostTaskRequest) (*dto.PostTaskResponse, []commonDto.Message, error) {
	respDto := new(dto.PostTaskResponse)
	messages, err := c.localClient.Rest(
		http.MethodPost,
		TaskServerRoot,
		*task, respDto,
		[]int{http.StatusOK})
	return respDto, messages, err
}

// UpdateTask Update the task.
func (c *TaskClient) UpdateTask(uri string, task *dto.UpdateTaskRequest) (*dto.PostTaskResponse, []commonDto.Message, error) {
	respDto := new(dto.PostTaskResponse)
	messages, err := c.localClient.Rest(
		http.MethodPost,
		uri+"/action/update",
		*task, respDto,
		[]int{http.StatusOK})
	return respDto, messages, err
}

// UpdateStep The task percentage will be set according to the c steop.
func (c *TaskClient) UpdateStep(uri string, step *dto.UpdateTaskStepRequest) (*dto.PostTaskResponse, []commonDto.Message, error) {
	respDto := new(dto.PostTaskResponse)
	messages, err := c.localClient.Rest(
		http.MethodPost,
		uri+"/action/updateTaskStep",
		*step, respDto,
		[]int{http.StatusOK})
	return respDto, messages, err
}

// SetStepExecutionState Set step execution state.
func (c *TaskClient) SetStepExecutionState(uri string, name string, state model.ExecutionState) (*dto.PostTaskResponse, []commonDto.Message, error) {
	request := new(dto.UpdateTaskStepRequest)
	request.Name = name
	request.ExecutionState = &state
	respDto := new(dto.PostTaskResponse)
	messages, err := c.localClient.Rest(
		http.MethodPost,
		uri+"/action/updateTaskStep",
		request, respDto,
		[]int{http.StatusOK})
	return respDto, messages, err
}

// SetStepExecutionResultState Set step execution result state.
func (c *TaskClient) SetStepExecutionResultState(uri string, name string, state model.ExecutionResultState) (*dto.PostTaskResponse, []commonDto.Message, error) {
	request := new(dto.UpdateTaskStepRequest)
	request.Name = name
	request.ExecutionResult = new(dto.UpdateExecutionResultRequest)
	request.ExecutionResult.State = &state
	respDto := new(dto.PostTaskResponse)
	messages, err := c.localClient.Rest(
		http.MethodPost,
		uri+"/action/updateTaskStep",
		request, respDto,
		[]int{http.StatusOK})
	return respDto, messages, err
}
