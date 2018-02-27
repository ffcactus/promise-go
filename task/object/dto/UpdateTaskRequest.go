package dto

import (
	"promise/task/object/model"
	"time"
)

// UpdateTaskRequest UpdateTaskRequest that only includes the changable properties.
// Note: Update sub task not use o request.
type UpdateTaskRequest struct {
	Description         *string               `json:"Description"`
	ExecutionState      *model.ExecutionState `json:"ExecutionState"`
	ExpectedExecutionMs *int                  `json:"ExpectedExecutionMs"`
	Percentage          *int                  `json:"Percentage"`
	ExecutionResult     *ExecutionResult      `json:"ExecutionResult"`
}

// UpdateModel Update the model.
func (o *UpdateTaskRequest) UpdateModel(current *model.Task) {
	if o.Description != nil {
		current.Description = *o.Description
	}
	if o.ExecutionState != nil {
		current.ExecutionState = *o.ExecutionState
	}
	if o.ExpectedExecutionMs != nil {
		current.ExpectedExecutionMs = *o.ExpectedExecutionMs
	}
	if o.Percentage != nil {
		current.Percentage = *o.Percentage
	}
	if o.ExecutionResult != nil {
		current.ExecutionResult.State = (*o.ExecutionResult).State
		current.ExecutionResult.Message = (*o.ExecutionResult).Message.Model()
	}
	current.UpdatedAt = time.Now()
}
