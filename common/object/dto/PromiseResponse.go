package dto

import (
	"promise/common/object/constValue"
	"promise/common/object/model"
	"time"
)

// PromiseResponseInterface is the APIs that a DTO should have.
type PromiseResponseInterface interface {
	Load(data interface{}) error
	GetID() string
	GetURI() string
	GetCategory() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

// PromiseResponse is the base type of a response in Promise.
type PromiseResponse struct {
	ID        string    `json:"ID"`
	URI       string    `json:"URI"`
	Category  string    `json:"Category"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

// GetID returns ID.
func (dto *PromiseResponse) GetID() string {
	return dto.ID
}

// GetURI returns URI.
func (dto *PromiseResponse) GetURI() string {
	return dto.URI
}

// GetCategory returns Category.
func (dto *PromiseResponse) GetCategory() string {
	return dto.Category
}

// GetCreatedAt returns CreatedAt.
func (dto *PromiseResponse) GetCreatedAt() time.Time {
	return dto.CreatedAt
}

// GetUpdatedAt returns UpdatedAt.
func (dto *PromiseResponse) GetUpdatedAt() time.Time {
	return dto.UpdatedAt
}

// Load data from model.
func (dto *PromiseResponse) Load(m *model.PromiseModel) {
	dto.ID = m.ID
	dto.URI = constValue.CategoryToURI(m.Category, m.ID)
	dto.Category = m.Category
	dto.CreatedAt = m.CreatedAt
	dto.UpdatedAt = m.UpdatedAt
}
