package dto

import (
	"promise/base"
	"promise/task/object/model"
)

// PostTaskRequest Post task request DTO.
type PostTaskRequest struct {
	base.Request
	MessageID     *string               `json:"MessageID"`
	Name          string                `json:"Name"`
	Description   *string               `json:"Description"`
	CreatedByName string                `json:"CreatedByName"`
	CreatedByURI  string                `json:"CreatedByURI"`
	TargetName    string                `json:"TargetName"`
	TargetURI     string                `json:"TargetURI"`
	TaskSteps     []PostTaskStepRequest `json:"TaskSteps"`
}

// IsValid return if the request is valid. 
func (dto *PostTaskRequest) IsValid() *base.Message {
	if len(dto.TaskSteps) == 0 {
		message := base.NewMessageTaskNoStep()
		return &message
	}
	return nil
}

// GetDebugName return the name for debug.
func (dto *PostTaskRequest) GetDebugName() string {
	return dto.Name
}

// ToModel convert the DTO to model.
func (dto PostTaskRequest) ToModel() base.ModelInterface {
	m := model.Task{}
	m.Category = base.CategoryTask
	m.Name = dto.Name
	m.Description = dto.Description
	m.ExecutionState = model.ExecutionStateReady
	m.CreatedByName = dto.CreatedByName
	m.CreatedByURI = dto.CreatedByURI
	m.TargetName = dto.TargetName
	m.TargetURI = dto.TargetURI
	m.Percentage = 0
	m.ExecutionResult.State = model.ExecutionResultStateUnknown
	m.ExpectedExecutionMs = 0
	for i := range dto.TaskSteps {
		m.TaskSteps = append(m.TaskSteps, *dto.TaskSteps[i].ToModel())
		// The task execution time equals to the sum of every steps'.
		m.ExpectedExecutionMs += dto.TaskSteps[i].ExpectedExecutionMs
	}
	return &m
}
