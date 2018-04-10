package dto

import (
	commonDTO "promise/common/object/dto"
	"promise/task/object/consterror"
	"promise/task/object/model"
)

// PostTaskStepRequest Post task step request DTO.
type PostTaskStepRequest struct {
	commonDTO.PromiseRequest
	MessageID           *string `json:"MessageID"`
	Name                string  `json:"Name"`
	Description         *string `json:"Description"`
	ExpectedExecutionMs uint64  `json:"ExpectedExecutionMs"`
}

// ToModel Convert to model.
func (dto *PostTaskStepRequest) ToModel() *model.TaskStep {
	var m model.TaskStep
	m.MessageID = dto.MessageID
	m.Name = dto.Name
	m.Description = dto.Description
	m.ExpectedExecutionMs = dto.ExpectedExecutionMs
	m.ExecutionState = model.ExecutionStateReady
	m.ExecutionResult.State = model.ExecutionResultStateUnknown
	return &m
}

// GetTaskStepResponse Post task step response DTO.
type GetTaskStepResponse struct {
	MessageID           *string              `json:"MessageID"`
	Name                string               `json:"Name"`
	Description         *string              `json:"Description,omitempty"`
	ExpectedExecutionMs uint64               `json:"ExpectedExecutionMs"`
	ExecutionState      model.ExecutionState `json:"ExecutionState"`
	ExecutionResult     ExecutionResult      `json:"ExecutionResult"`
}

// Load Load from model.
func (dto *GetTaskStepResponse) Load(m *model.TaskStep) {
	dto.MessageID = m.MessageID
	dto.Name = m.Name
	dto.Description = m.Description
	dto.ExpectedExecutionMs = m.ExpectedExecutionMs
	dto.ExecutionState = m.ExecutionState
	dto.ExecutionResult.Load(&m.ExecutionResult)
}

// UpdateTaskStepRequest Update task step request DTO.
type UpdateTaskStepRequest struct {
	commonDTO.PromiseRequest
	Name            string                        `json:"Name"`
	ExecutionState  *model.ExecutionState         `json:"ExecutionState"`
	ExecutionResult *UpdateExecutionResultRequest `json:"ExecutionResult"`
}

// Validate the request.
func (dto *UpdateTaskStepRequest) Validate() error {
	if dto.ExecutionState != nil {
		if *dto.ExecutionState == model.ExecutionStateReady {
			return nil
		}
		if *dto.ExecutionState == model.ExecutionStateRunning {
			return nil
		}
		if *dto.ExecutionState == model.ExecutionStateSuspended {
			return nil
		}
		if *dto.ExecutionState == model.ExecutionStateTerminated {
			return nil
		}
		return consterror.ErrorUnknownExecutionState
	}
	if dto.ExecutionResult != nil && dto.ExecutionResult.State != nil {
		if *dto.ExecutionResult.State == model.ExecutionResultStateFinished {
			return nil
		}
		if *dto.ExecutionResult.State == model.ExecutionResultStateWarning {
			return nil
		}
		if *dto.ExecutionResult.State == model.ExecutionResultStateError {
			return nil
		}
		if *dto.ExecutionResult.State == model.ExecutionResultStateAbort {
			return nil
		}
		if *dto.ExecutionResult.State == model.ExecutionResultStateUnknown {
			return nil
		}
		return consterror.ErrorUnknownExecutionResult
	}
	return nil
}

// UpdateModel Update the model.
func (dto *UpdateTaskStepRequest) _updateModel(current *model.Task) {
	for i := range current.TaskSteps {
		if dto.Name == current.TaskSteps[i].Name {
			if dto.ExecutionState != nil {
				current.TaskSteps[i].ExecutionState = *dto.ExecutionState
			}
			if dto.ExecutionResult != nil {
				dto.ExecutionResult.UpdateModel(&current.TaskSteps[i].ExecutionResult)
			}
		}
	}
}

// UpdateModel Update the model.
func (dto *UpdateTaskStepRequest) UpdateModel(m *model.Task) {
	currentTime := uint64(0)
	for i := range m.TaskSteps {
		currentTime += m.TaskSteps[i].ExpectedExecutionMs
		if m.TaskSteps[i].Name == dto.Name {
			// Found the step, and update the current time.
			switch m.TaskSteps[i].ExecutionState {
			case model.ExecutionStateTerminated:
			case model.ExecutionStateRunning:
			case model.ExecutionStateSuspended:
				currentTime -= m.TaskSteps[i].ExpectedExecutionMs
			default:
			}
			if dto.ExecutionState != nil {
				m.TaskSteps[i].ExecutionState = *dto.ExecutionState
				if *dto.ExecutionState == model.ExecutionStateRunning {
					m.CurrentStep = m.TaskSteps[i].Name
				}
			}
			if dto.ExecutionResult != nil {
				dto.ExecutionResult.UpdateModel(&m.TaskSteps[i].ExecutionResult)
			}
			percentageF := (float32)(currentTime) / (float32)(m.ExpectedExecutionMs)
			m.Percentage = (int)((percentageF * 100) + 0.5)
			if m.Percentage > 100 {
				m.Percentage = 100
			}
		}
	}
}
