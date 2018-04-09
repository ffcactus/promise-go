package base

import (
	"time"
	"promise/apps"
	log "github.com/sirupsen/logrus"
)

// ResponseInterface is the inteface that a promise Request should have.
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
	ID        string    `json:"ID"`
	URI       string    `json:"URI"`
	Category  string    `json:"Category"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

// GetDebugName return the name for debug.
func (dto *Response) GetDebugName() string {
	log.Error("Using default Response.GetDebugName().")
	return "NotProvided"
}

// Load data from model.
func (dto *Response) Load(m ModelInterface) error {
	dto.ID = m.GetID()
	dto.URI = apps.CategoryToURI(m.GetCategory(), m.GetID())
	dto.Category = m.GetCategory()
	dto.CreatedAt = m.GetCreatedAt()
	dto.UpdatedAt = m.GetUpdatedAt()
	return nil
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
