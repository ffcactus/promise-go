package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/task/object/model"
)

// TaskCollectionMember is the DTO used in collection response.
type TaskCollectionMember struct {
	base.CollectionMemberResponse
	Name            string               `json:"Name"`
	Description     *string              `json:"Description,omitempty"`
	ExecutionState  model.ExecutionState `json:"ExecutionState"`
	CurrentStep     string               `json:"CurrentStep"`
	Percentage      uint32               `json:"Percentage"`
	ExecutionResult ExecutionResult      `json:"ExecutionResult"`
}

// Load will load info from model.
func (dto *TaskCollectionMember) Load(i base.CollectionMemberModelInterface) error {
	m, ok := i.(*model.TaskCollectionMember)
	if !ok {
		log.Error("TaskCollectionMember.Load() failed, convert data failed.")
		return base.ErrorDataConvert
	}
	dto.CollectionMemberResponse.Load(&m.CollectionMemberModel)
	dto.Name = m.Name
	dto.Description = m.Description
	dto.ExecutionState = m.ExecutionState
	dto.CurrentStep = m.CurrentStep
	dto.Percentage = m.Percentage
	dto.ExecutionResult.Load(&m.ExecutionResult)
	return nil
}

// GetTaskCollectionResponse Get task collection response DTO.
type GetTaskCollectionResponse struct {
	base.CollectionResponse
}

// Load Load from model.
func (dto *GetTaskCollectionResponse) Load(m *base.CollectionModel) error {
	dto.Start = m.Start
	dto.Count = m.Count
	dto.Total = m.Total
	for _, v := range m.Members {
		member := TaskCollectionMember{}
		if err := member.Load(v); err != nil {
			return err
		}
		dto.Members = append(dto.Members, &member)
	}
	return nil
}
