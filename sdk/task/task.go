package task

import (
	"net/http"
	"promise/base"
	"promise/task/object/dto"
	"promise/task/object/model"
)

var (
	// TaskServerRoot The root of the service.
	TaskServerRoot = base.ProtocolScheme + base.Host + base.RootURL + base.TaskBaseURI + "/"
)

// CreateTask Create the task.
func CreateTask(task *dto.PostTaskRequest) (*dto.GetTaskResponse, []base.Message, error) {
	respDto := new(dto.GetTaskResponse)
	messages, err := base.REST(
		http.MethodPost,
		TaskServerRoot,
		*task, respDto,
		[]int{http.StatusCreated})
	return respDto, messages, err
}

// UpdateTask Update the task.
func UpdateTask(taskID string, task *dto.UpdateTaskRequest) (*dto.GetTaskResponse, []base.Message, error) {
	respDto := new(dto.GetTaskResponse)
	messages, err := base.REST(
		http.MethodPost,
		TaskServerRoot+taskID+"/action/updatetask",
		*task, respDto,
		[]int{http.StatusAccepted})
	return respDto, messages, err
}

// UpdateStep The task percentage will be set according to the c steop.
func UpdateStep(taskID string, step *dto.UpdateTaskStepRequest) (*dto.GetTaskResponse, []base.Message, error) {
	respDto := new(dto.GetTaskResponse)
	messages, err := base.REST(
		http.MethodPost,
		TaskServerRoot+taskID+"/action/updatetaskstep",
		*step, respDto,
		[]int{http.StatusAccepted})
	return respDto, messages, err
}

// SetStepExecutionState Set step execution state.
func SetStepExecutionState(taskID string, name string, state model.ExecutionState) (*dto.GetTaskResponse, []base.Message, error) {
	request := new(dto.UpdateTaskStepRequest)
	request.Name = name
	request.ExecutionState = &state
	respDto := new(dto.GetTaskResponse)
	messages, err := base.REST(
		http.MethodPost,
		TaskServerRoot+taskID+"/action/updatetaskstep",
		request, respDto,
		[]int{http.StatusAccepted})
	return respDto, messages, err
}

// SetStepExecutionResultState Set step execution result state.
func SetStepExecutionResultState(taskID string, name string, state model.ExecutionResultState) (*dto.GetTaskResponse, []base.Message, error) {
	request := new(dto.UpdateTaskStepRequest)
	request.Name = name
	request.ExecutionResult = new(dto.UpdateExecutionResultRequest)
	request.ExecutionResult.State = &state
	respDto := new(dto.GetTaskResponse)
	messages, err := base.REST(
		http.MethodPost,
		TaskServerRoot+taskID+"/action/updatetaskstep",
		request, respDto,
		[]int{http.StatusAccepted})
	return respDto, messages, err
}
