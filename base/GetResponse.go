package base

import (
	"time"
)

// GetResponseInterface is the interface that GetResponse have.
type GetResponseInterface interface {
	ResponseInterface
	GetID() string
	GetCategory() string
	Load(current ModelInterface) error
}

// GetResponse is the DTO of a get resource response.
type GetResponse struct {
	ID        string    `json:"ID"`
	URI       string    `json:"URI"`
	Category  string    `json:"Category"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

// Load the data from Model to GetResponse.
func (dto *GetResponse) Load(m *Model) {
	dto.ID = m.ID
	dto.URI = CategoryToURI(m.Category, m.ID)
	dto.Category = m.Category
	dto.CreatedAt = m.CreatedAt
	dto.UpdatedAt = m.UpdatedAt
}

// GetID returns ID.
func (dto *GetResponse) GetID() string {
	return dto.ID
}

// GetCategory returns Category.
func (dto *GetResponse) GetCategory() string {
	return dto.Category
}
