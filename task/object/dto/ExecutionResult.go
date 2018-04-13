package dto

import (
	"promise/base"
	"promise/task/object/model"
)

// ExecutionResult Execution result DTO.
type ExecutionResult struct {
	State   model.ExecutionResultState `json:"State"`
	Message *base.Message              `json:"Message"`
}

// Load Load from model.
func (o *ExecutionResult) Load(m *model.ExecutionResult) {
	o.State = m.State
	o.Message = m.Message
}

// UpdateExecutionResultRequest Update execution result step request DTO.
type UpdateExecutionResultRequest struct {
	State   *model.ExecutionResultState `json:"State"`
	Message *base.Message               `json:"Message"`
}

// UpdateModel Update from model.
func (o *UpdateExecutionResultRequest) UpdateModel(m *model.ExecutionResult) {
	if o.State != nil {
		m.State = *o.State
	}
	if o.Message != nil {
		m.Message = o.Message
	}
}
