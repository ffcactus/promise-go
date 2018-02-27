package dto

import (
	"promise/task/object/model"
	"time"
)

// TaskMember Task member DTO.
type TaskMember struct {
	URI              string               `json:"URI"`
	Name             string               `json:"Name"`
	Description      string               `json:"Description"`
	CreatedAt        time.Time            `json:"CreatedAt"`
	UpdatedAt        time.Time            `json:"UpdatedAt"`
	CreatedByName    string               `json:"CreatedByName"`
	CreatedByPageURI string               `json:"CreatedByPageURI"`
	TargetName       string               `json:"TargetName"`
	TargetPageURI    string               `json:"TargetPageURI"`
	ExecutionState   model.ExecutionState `json:"ExecutionState"`
	CurrentStep      string               `json:"CurrentStep"`
	Percentage       int                  `json:"Percentage"`
	ExecutionResult  ExecutionResult      `json:"ExecutionResult"`
}

// GetTaskCollectionResponse Get task collection response DTO.
type GetTaskCollectionResponse struct {
	Start       int          `json:"Start"`
	Count       int          `json:"Count"`
	Total       int          `json:"Total"`
	Members     []TaskMember `json:"Members"`
	NextPageURI *string      `json:"NextPageURI,omitempty"`
	PrevPageURI *string      `json:"PrevPageURI,omitempty"`
}

// Load Load from model.
func (o *GetTaskCollectionResponse) Load(m *model.TaskCollection) {
	o.Start = m.Start
	o.Count = m.Count
	o.Total = m.Total
	for i := range m.Members {
		each := new(TaskMember)
		each.URI = m.Members[i].URI
		each.Name = m.Members[i].Name
		each.Description = m.Members[i].Description
		each.ExecutionState = m.Members[i].ExecutionState
		each.Percentage = m.Members[i].Percentage
		each.ExecutionResult.Load(&m.Members[i].ExecutionResult)
		o.Members = append(o.Members, *each)
	}
	if m.NextPageURI != "" {
		o.NextPageURI = &m.NextPageURI
	}
	if m.PrevPageURI != "" {
		o.PrevPageURI = &m.PrevPageURI
	}
}
