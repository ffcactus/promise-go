package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/task/object/model"
	commonConstError "promise/common/object/consterror"
	commonDTO "promise/common/object/dto"
)

// GetTaskResponse Post task response DTO.
type GetTaskResponse struct {
	commonDTO.PromiseResponse
	Name                string                `json:"Name"`
	ParentTask          *string               `json:"ParentTask"`
	Description         *string               `json:"Description,omitempty"`
	ExecutionState      model.ExecutionState  `json:"ExecutionState"`
	CreatedByName       string                `json:"CreatedByName"`
	CreatedByURI        string                `json:"CreatedByURI"`
	TargetName          string                `json:"TargetName"`
	TargetURI           string                `json:"TargetURI"`
	ExpectedExecutionMs uint64                `json:"ExpectedExecutionMs"`
	Percentage          int                   `json:"Percentage"`
	CurrentStep         string                `json:"CurrentStep"`
	TaskSteps           []GetTaskStepResponse `json:"TaskSteps"`
	ExecutionResult     ExecutionResult       `json:"ExecutionResult"`
	SubTasks            []GetTaskResponse     `json:"SubTasks"`
}

// Load Load from model.
func (dto *GetTaskResponse) Load(data interface{}) error {
	m, ok := data.(*model.Task)
	if !ok {
		log.Warn("GetTaskResponse load data from model failed.")
		return commonConstError.ErrorDataConvert
	}
	dto.PromiseResponse.Load(&m.PromiseModel)
	dto.Name = m.Name
	dto.ParentTask = m.ParentTask
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
	dto.SubTasks = make([]GetTaskResponse, 0)
	for i := range m.SubTasks {
		sub := new(GetTaskResponse)
		sub.Load(&m.SubTasks[i])
		dto.SubTasks = append(dto.SubTasks, *sub)
	}
	return nil
}
