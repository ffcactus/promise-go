package base

import (
	"promise/apps"
	"time"
)

// ResponseTemplateInterface is the inteface that a concrete Request should have.
type ResponseTemplateInterface interface {
	GetDebugName() string
	Load(ModelInterface) error
	GetID() string
	GetURI() string
	GetCategory() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

// ResponseInterface is the inteface that a Request should have.
type ResponseInterface interface {
	GetDebugName() string
	Load(ModelInterface) error
	GetID() string
	GetURI() string
	GetCategory() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

// Response is the response DTO used in Promise project.
type Response struct {
	TemplateImpl ResponseTemplateInterface `json:"-"`
	ID           string                    `json:"ID"`
	URI          string                    `json:"URI"`
	Category     string                    `json:"Category"`
	CreatedAt    time.Time                 `json:"CreatedAt"`
	UpdatedAt    time.Time                 `json:"UpdatedAt"`
}

// GetDebugName return the name for debug.
func (dto *Response) GetDebugName() string {
	return dto.TemplateImpl.GetDebugName()
}

// GetID returns ID.
func (dto *Response) GetID() string {
	return dto.ID
}

// GetURI returns URI.
func (dto *Response) GetURI() string {
	return dto.URI
}

// GetCategory returns Category.
func (dto *Response) GetCategory() string {
	return dto.Category
}

// GetCreatedAt returns CreatedAt.
func (dto *Response) GetCreatedAt() time.Time {
	return dto.CreatedAt
}

// GetUpdatedAt returns UpdatedAt.
func (dto *Response) GetUpdatedAt() time.Time {
	return dto.UpdatedAt
}

// Load data from model.
func (dto *Response) Load(m ModelInterface) error {
	return dto.TemplateImpl.Load(m)
}

// ResponseLoad load model to DTO.
func ResponseLoad(dto *Response, m *Model) {
	dto.ID = m.ID
	dto.URI = apps.CategoryToURI(m.Category, m.ID)
	dto.Category = m.Category
	dto.CreatedAt = m.CreatedAt
	dto.UpdatedAt = m.UpdatedAt
}
