package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/task/object/model"
)

// PostTaskStepRequest Post task step request DTO.
type PostTaskStepRequest struct {
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
	Name            string                        `json:"Name"`
	ExecutionState  *model.ExecutionState         `json:"ExecutionState"`
	ExecutionResult *UpdateExecutionResultRequest `json:"ExecutionResult"`
}

// NewInstance creates a new instance.
func (dto *UpdateTaskStepRequest) NewInstance() base.RequestInterface {
	return new(UpdateTaskStepRequest)
}

// IsValid return if the request is valid.
func (dto *UpdateTaskStepRequest) IsValid() *base.Message {
	message := base.NewMessageUnknownPropertyValue()
	if dto.ExecutionState != nil && !model.IsValidExecutionState(*dto.ExecutionState) {
		return message
	}
	if dto.ExecutionResult != nil && dto.ExecutionResult.State != nil && !model.IsValidExecutionResultState(*dto.ExecutionResult.State) {
		return message
	}
	return nil
}

// String return the name for debug.
func (dto UpdateTaskStepRequest) String() string {
	return dto.Name
}

// UpdateModel Update the model.
func (dto *UpdateTaskStepRequest) UpdateModel(i base.ModelInterface) error {
	m, ok := i.(*model.Task)
	if !ok {
		log.Error("UpdateTaskStepRequest.UpdateModel() convert interface failed.")
		return base.ErrorDataConvert
	}
	currentTime := uint64(0)
	foundStep := false
	for i := range m.TaskSteps {
		currentTime += m.TaskSteps[i].ExpectedExecutionMs
		if m.TaskSteps[i].Name == dto.Name {
			// Found the step, and update the current time.
			foundStep = true

			if dto.ExecutionState != nil {
				// whenever the use update a task step we think the current step is this one.
				m.CurrentStep = m.TaskSteps[i].Name

				m.TaskSteps[i].ExecutionState = *dto.ExecutionState
				// if the user set the step to not yet finished,
				// we need adjust the percentage, and set the current step accrodingly.
				switch *dto.ExecutionState {
				case model.ExecutionStateReady,
					model.ExecutionStateRunning,
					model.ExecutionStateSuspended:
					currentTime -= m.TaskSteps[i].ExpectedExecutionMs
				}
			}
			if dto.ExecutionResult != nil {
				dto.ExecutionResult.UpdateModel(&m.TaskSteps[i].ExecutionResult)
			}

			percentageF := (float32)(currentTime) / (float32)(m.ExpectedExecutionMs)
			m.Percentage = (uint32)((percentageF * 100) + 0.5)
			if m.Percentage > 100 {
				m.Percentage = 100
			}
			break
		}
	}
	if !foundStep {
		return base.ErrorUnknownPropertyValue
	}
	return nil
}
