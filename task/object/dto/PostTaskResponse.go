package dto

import (
	"promise/task/object/model"
	"time"
)

// PostTaskResponse Post task response DTO.
type PostTaskResponse struct {
	ID                  string                 `json:"ID"`
	URI                 string                 `json:"URI"`
	Name                string                 `json:"Name"`
	ParentTask          *string                `json:"ParentTask"`
	Description         string                 `json:"Description"`
	ExecutionState      model.ExecutionState   `json:"ExecutionState"`
	CreatedByName       string                 `json:"CreatedByName"`
	CreatedByURI        string                 `json:"CreatedByURI"`
	TargetName          string                 `json:"TargetName"`
	TargetURI           string                 `json:"TargetURI"`
	ExpectedExecutionMs int                    `json:"ExpectedExecutionMs"`
	Percentage          int                    `json:"Percentage"`
	CreatedAt           time.Time              `json:"CreatedAt"`
	UpdatedAt           time.Time              `json:"UpdatedAt"`
	CurrentStep         string                 `json:"CurrentStep"`
	TaskSteps           []PostTaskStepResponse `json:"TaskSteps"`
	ExecutionResult     ExecutionResult        `json:"ExecutionResult"`
	SubTasks            []PostTaskResponse     `json:"SubTasks"`
}

// Load Load from model.
func (o *PostTaskResponse) Load(m *model.Task) {
	o.ID = m.ID
	o.URI = m.URI
	o.Name = m.Name
	o.ParentTask = m.ParentTask
	o.Description = m.Description
	o.ExecutionState = m.ExecutionState
	o.CreatedByName = m.CreatedByName
	o.CreatedByURI = m.CreatedByURI
	o.TargetName = m.TargetName
	o.TargetURI = m.TargetURI
	o.ExpectedExecutionMs = m.ExpectedExecutionMs
	o.Percentage = m.Percentage
	o.CreatedAt = m.CreatedAt
	o.UpdatedAt = m.UpdatedAt
	o.CurrentStep = m.CurrentStep
	o.TaskSteps = make([]PostTaskStepResponse, 0)
	for i := range m.TaskSteps {
		step := new(PostTaskStepResponse)
		step.Load(&m.TaskSteps[i])
		o.TaskSteps = append(o.TaskSteps, *step)
	}
	o.ExecutionResult.Load(&m.ExecutionResult)
	o.SubTasks = make([]PostTaskResponse, 0)
	for i := range m.SubTasks {
		sub := new(PostTaskResponse)
		sub.Load(&m.SubTasks[i])
		o.SubTasks = append(o.SubTasks, *sub)
	}
}
