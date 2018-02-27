package dto

import (
	"promise/task/object/model"
)

// PostTaskStepRequest Post task step request DTO.
type PostTaskStepRequest struct {
	MessageID           *string `json:"MessageID"`
	Name                string  `json:"Name"`
	Description         string  `json:"Description"`
	ExpectedExecutionMs int     `json:"ExpectedExecutionMs"`
}

// ToModel Convert to model.
func (o *PostTaskStepRequest) ToModel() *model.TaskStep {
	var m model.TaskStep
	m.MessageID = o.MessageID
	m.Name = o.Name
	m.Description = o.Description
	m.ExpectedExecutionMs = o.ExpectedExecutionMs
	m.ExecutionState = model.ExecutionStateReady
	m.ExecutionResult.State = model.ExecutionResultStateUnknown
	return &m
}

// PostTaskStepResponse Post task step response DTO.
type PostTaskStepResponse struct {
	MessageID           *string              `json:"MessageID"`
	Name                string               `json:"Name"`
	Description         string               `json:"Description"`
	ExpectedExecutionMs int                  `json:"ExpectedExecutionMs"`
	ExecutionState      model.ExecutionState `json:"ExecutionState"`
	ExecutionResult     ExecutionResult      `json:"ExecutionResult"`
}

// Load Load from model.
func (o *PostTaskStepResponse) Load(m *model.TaskStep) {
	o.MessageID = m.MessageID
	o.Name = m.Name
	o.Description = m.Description
	o.ExpectedExecutionMs = m.ExpectedExecutionMs
	o.ExecutionState = m.ExecutionState
	o.ExecutionResult.Load(&m.ExecutionResult)
}

// UpdateTaskStepRequest Update task step request DTO.
type UpdateTaskStepRequest struct {
	Name            string                        `json:"Name"`
	ExecutionState  *model.ExecutionState         `json:"ExecutionState"`
	ExecutionResult *UpdateExecutionResultRequest `json:"ExecutionResult"`
}

// UpdateModel Update the model.
func (o *UpdateTaskStepRequest) UpdateModel(current *model.Task) {
	for i := range current.TaskSteps {
		if o.Name == current.TaskSteps[i].Name {
			if o.ExecutionState != nil {
				current.TaskSteps[i].ExecutionState = *o.ExecutionState
			}
			if o.ExecutionResult != nil {
				o.ExecutionResult.UpdateModel(&current.TaskSteps[i].ExecutionResult)
			}
		}
	}
}
