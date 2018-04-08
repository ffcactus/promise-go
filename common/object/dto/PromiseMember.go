package dto

import (
	"promise/common/object/model"
	"promise/common/object/constvalue"
)

// PromiseMember is the Promise collection member in DTO.
type PromiseMember struct {
	ID        string    `json:"ID"`
	URI       string    `json:"URI"`
	Category  string    `json:"Category"`
}

// Load will load data from model.
func (dto *PromiseMember) Load(m *model.PromiseMember) {
	dto.ID = m.ID
	dto.URI = constvalue.CategoryToURI(m.Category, m.ID)
	dto.Category = m.Category
}