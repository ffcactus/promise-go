package dto

import (
	"promise/common/category"
	commonDTO "promise/common/object/dto"
	commonMessage "promise/common/object/message"
	"promise/task/object/message"
	"promise/task/object/model"
	"time"
)

// PostTaskRequest Post task request DTO.
type PostTaskRequest struct {
	commonDTO.PromiseRequest
	MessageID     *string               `json:"MessageID"`
	Name          string                `json:"Name"`
	Description   *string               `json:"Description"`
	CreatedByName string                `json:"CreatedByName"`
	CreatedByURI  string                `json:"CreatedByURI"`
	TargetName    string                `json:"TargetName"`
	TargetURI     string                `json:"TargetURI"`
	TaskSteps     []PostTaskStepRequest `json:"TaskSteps"`
}

// Validate the request.
func (dto *PostTaskRequest) Validate() *commonMessage.Message {
	if len(dto.TaskSteps) == 0 {
		m := message.NewMessageTaskBadRequest()
		return &m
	}
	return nil
}

// ToModel Convert to model.
func (dto PostTaskRequest) ToModel() *model.Task {
	m := model.Task{}
	m.Category = category.Task
	m.Name = dto.Name
	m.Description = dto.Description
	m.ExecutionState = model.ExecutionStateReady
	m.CreatedByName = dto.CreatedByName
	m.CreatedByURI = dto.CreatedByURI
	m.TargetName = dto.TargetName
	m.TargetURI = dto.TargetURI
	m.Percentage = 0
	m.CreatedAt = time.Now()
	m.UpdatedAt = m.CreatedAt
	m.ExecutionResult.State = model.ExecutionResultStateUnknown
	m.ExpectedExecutionMs = 0
	for i := range dto.TaskSteps {
		m.TaskSteps = append(m.TaskSteps, *dto.TaskSteps[i].ToModel())
		// The task execution time equals to the sum of every steps'.
		m.ExpectedExecutionMs += dto.TaskSteps[i].ExpectedExecutionMs
	}
	return &m
}
