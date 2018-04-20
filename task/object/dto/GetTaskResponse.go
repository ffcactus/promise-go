package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/task/object/model"
)

// GetTaskResponse Post task response DTO.
type GetTaskResponse struct {
	base.Response
	MessageID           *string               `json:"MessageID,omitempty"`
	Name                string                `json:"Name"`
	Description         *string               `json:"Description,omitempty"`
	ExecutionState      model.ExecutionState  `json:"ExecutionState"`
	CreatedByName       string                `json:"CreatedByName"`
	CreatedByURI        string                `json:"CreatedByURI"`
	TargetName          string                `json:"TargetName"`
	TargetURI           string                `json:"TargetURI"`
	ExpectedExecutionMs uint64                `json:"ExpectedExecutionMs"`
	Percentage          uint32                `json:"Percentage"`
	CurrentStep         string                `json:"CurrentStep"`
	TaskSteps           []GetTaskStepResponse `json:"TaskSteps"`
	ExecutionResult     ExecutionResult       `json:"ExecutionResult"`
}

// GetDebugName return the name for debug.
func (dto *GetTaskResponse) GetDebugName() string {
	return dto.Name
}

// Load will load info from mode
func (dto *GetTaskResponse) Load(data base.ModelInterface) error {
	m, ok := data.(*model.Task)
	if !ok {
		log.Error("GetTaskResponse.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	base.ResponseLoad(&dto.Response, &m.Model)
	dto.MessageID = m.MessageID
	dto.Name = m.Name
	dto.Description = m.Description
	dto.ExecutionState = m.ExecutionState
	dto.CreatedByName = m.CreatedByName
	dto.CreatedByURI = m.CreatedByURI
	dto.TargetName = m.TargetName
	dto.TargetURI = m.TargetURI
	dto.ExpectedExecutionMs = m.ExpectedExecutionMs
	dto.Percentage = m.Percentage
	dto.CurrentStep = m.CurrentStep
	dto.TaskSteps = make([]GetTaskStepResponse, 0)
	for i := range m.TaskSteps {
		step := new(GetTaskStepResponse)
		step.Load(&m.TaskSteps[i])
		dto.TaskSteps = append(dto.TaskSteps, *step)
	}
	dto.ExecutionResult.Load(&m.ExecutionResult)
	return nil
}
