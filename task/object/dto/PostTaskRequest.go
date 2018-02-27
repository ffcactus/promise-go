package dto

import (
	"promise/task/object/model"
	"time"
)

// PostTaskRequest Post task request DTO.
type PostTaskRequest struct {
	MessageID     *string               `json:"MessageID"`
	Name          string                `json:"Name"`
	Description   string                `json:"Description"`
	CreatedByName string                `json:"CreatedByName"`
	CreatedByURI  string                `json:"CreatedByURI"`
	TargetName    string                `json:"TargetName"`
	TargetURI     string                `json:"TargetURI"`
	TaskSteps     []PostTaskStepRequest `json:"TaskSteps"`
}

// ToModel Convert to model.
func (o PostTaskRequest) ToModel() *model.Task {
	m := new(model.Task)
	m.Name = o.Name
	m.Description = o.Description
	m.ExecutionState = model.ExecutionStateReady
	m.CreatedByName = o.CreatedByName
	m.CreatedByURI = o.CreatedByURI
	m.TargetName = o.TargetName
	m.TargetURI = o.TargetURI
	m.Percentage = 0
	m.CreatedAt = time.Now()
	m.UpdatedAt = m.CreatedAt
	m.ExecutionResult.State = model.ExecutionResultStateUnknown
	m.ExpectedExecutionMs = 0
	for i := range o.TaskSteps {
		m.TaskSteps = append(m.TaskSteps, *o.TaskSteps[i].ToModel())
		// The task execution time equals to the sum of every steps'.
		m.ExpectedExecutionMs += o.TaskSteps[i].ExpectedExecutionMs
	}
	return m
}
