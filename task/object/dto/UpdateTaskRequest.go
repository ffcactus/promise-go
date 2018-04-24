package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/task/object/model"
)

// UpdateTaskRequest UpdateTaskRequest that only includes the changable properties.
// Note: Update sub task not use this request.
type UpdateTaskRequest struct {
	Description         *string                       `json:"Description"`
	ExecutionState      *model.ExecutionState         `json:"ExecutionState"`
	ExpectedExecutionMs *uint64                       `json:"ExpectedExecutionMs"`
	Percentage          *uint32                       `json:"Percentage"`
	ExecutionResult     *UpdateExecutionResultRequest `json:"ExecutionResult"`
}

// NewInstance creates a new instance.
func (dto *UpdateTaskRequest) NewInstance() base.RequestInterface {
	return new(UpdateTaskRequest)
}

// IsValid return if the request is valid.
func (dto *UpdateTaskRequest) IsValid() *base.Message {
	message := base.NewMessageUnknownPropertyValue()
	if dto.Percentage != nil && *dto.Percentage > 100 {
		return message
	}
	if dto.ExecutionState != nil && !model.IsValidExecutionState(*dto.ExecutionState) {
		return message
	}
	if dto.ExecutionResult != nil && dto.ExecutionResult.State != nil && !model.IsValidExecutionResultState(*dto.ExecutionResult.State) {
		return message
	}
	return nil
}

// DebugInfo return the name for debug.
func (dto *UpdateTaskRequest) DebugInfo() string {
	return ""
}

// UpdateModel Update the model.
func (dto *UpdateTaskRequest) UpdateModel(i base.ModelInterface) error {
	m, ok := i.(*model.Task)
	if !ok {
		log.Error("UpdateTaskRequest.UpdateModel() convert interface failed.")
		return base.ErrorDataConvert
	}
	if dto.Description != nil {
		m.Description = dto.Description
	}
	if dto.ExecutionState != nil {
		m.ExecutionState = *dto.ExecutionState
	}
	if dto.ExpectedExecutionMs != nil {
		m.ExpectedExecutionMs = *dto.ExpectedExecutionMs
	}
	if dto.Percentage != nil {
		m.Percentage = *dto.Percentage
	}
	if dto.ExecutionResult != nil {
		dto.ExecutionResult.UpdateModel(&m.ExecutionResult)
	}
	return nil
}
