package dto

import (
	"promise/task/object/model"
)

// UpdateTaskRequest UpdateTaskRequest that only includes the changable properties.
// Note: Update sub task not use o request.
type UpdateTaskRequest struct {
	Description         *string               `json:"Description"`
	ExecutionState      *model.ExecutionState `json:"ExecutionState"`
	ExpectedExecutionMs *uint64               `json:"ExpectedExecutionMs"`
	Percentage          *int                  `json:"Percentage"`
	ExecutionResult     *ExecutionResult      `json:"ExecutionResult"`
}

// UpdateModel Update the model.
func (dto *UpdateTaskRequest) UpdateModel(current *model.Task) {
	if dto.Description != nil {
		current.Description = dto.Description
	}
	if dto.ExecutionState != nil {
		current.ExecutionState = *dto.ExecutionState
	}
	if dto.ExpectedExecutionMs != nil {
		current.ExpectedExecutionMs = *dto.ExpectedExecutionMs
	}
	if dto.Percentage != nil {
		current.Percentage = *dto.Percentage
	}
	if dto.ExecutionResult != nil {
		current.ExecutionResult.State = (*dto.ExecutionResult).State
		current.ExecutionResult.Message = (*dto.ExecutionResult).Message
	}
}
