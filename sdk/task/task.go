package task

import (
	"net/http"
	"promise/common/app"
	"promise/common/app/rest"
	commonDto "promise/common/object/dto"
	"promise/task/object/dto"
	"promise/task/object/model"
	"promise/common/object/constValue"
)

var (
	// TaskServerRoot The root of the service.
	TaskServerRoot = app.ProtocolScheme + app.Host + app.RootURL + constValue.TaskBaseURI
)

// CreateTask Create the task.
func CreateTask(task *dto.PostTaskRequest) (*dto.PostTaskResponse, []commonDto.Message, error) {
	respDto := new(dto.PostTaskResponse)
	messages, err := rest.Do(
		http.MethodPost,
		TaskServerRoot,
		*task, respDto,
		[]int{http.StatusOK})
	return respDto, messages, err
}

// UpdateTask Update the task.
func UpdateTask(taskID string, task *dto.UpdateTaskRequest) (*dto.PostTaskResponse, []commonDto.Message, error) {
	respDto := new(dto.PostTaskResponse)
	messages, err := rest.Do(
		http.MethodPost,
		TaskServerRoot+taskID+"/action/update",
		*task, respDto,
		[]int{http.StatusOK})
	return respDto, messages, err
}

// UpdateStep The task percentage will be set according to the c steop.
func UpdateStep(taskID string, step *dto.UpdateTaskStepRequest) (*dto.PostTaskResponse, []commonDto.Message, error) {
	respDto := new(dto.PostTaskResponse)
	messages, err := rest.Do(
		http.MethodPost,
		TaskServerRoot+taskID+"/action/updateTaskStep",
		*step, respDto,
		[]int{http.StatusOK})
	return respDto, messages, err
}

// SetStepExecutionState Set step execution state.
func SetStepExecutionState(taskID string, name string, state model.ExecutionState) (*dto.PostTaskResponse, []commonDto.Message, error) {
	request := new(dto.UpdateTaskStepRequest)
	request.Name = name
	request.ExecutionState = &state
	respDto := new(dto.PostTaskResponse)
	messages, err := rest.Do(
		http.MethodPost,
		TaskServerRoot+taskID+"/action/updateTaskStep",
		request, respDto,
		[]int{http.StatusOK})
	return respDto, messages, err
}

// SetStepExecutionResultState Set step execution result state.
func SetStepExecutionResultState(taskID string, name string, state model.ExecutionResultState) (*dto.PostTaskResponse, []commonDto.Message, error) {
	request := new(dto.UpdateTaskStepRequest)
	request.Name = name
	request.ExecutionResult = new(dto.UpdateExecutionResultRequest)
	request.ExecutionResult.State = &state
	respDto := new(dto.PostTaskResponse)
	messages, err := rest.Do(
		http.MethodPost,
		TaskServerRoot+taskID+"/action/updateTaskStep",
		request, respDto,
		[]int{http.StatusOK})
	return respDto, messages, err
}
