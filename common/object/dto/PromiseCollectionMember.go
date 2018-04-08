package dto

import (
	"promise/common/object/model"
	"promise/common/object/constvalue"
)

// PromiseCollectionMember is the Promise collection member in DTO.
type PromiseCollectionMember struct {
	ID        string    `json:"ID"`
	URI       string    `json:"URI"`
	Category  string    `json:"Category"`
}

// Load will load data from model.
func (dto *PromiseCollectionMember) Load(m *model.PromiseCollectionMember) {
	dto.ID = m.ID
	dto.URI = constvalue.CategoryToURI(m.Category, m.ID)
	dto.Category = m.Category
}