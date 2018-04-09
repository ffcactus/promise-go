package base

import (
	"time"
)

// ModelInterface is the interface of Model.
type ModelInterface interface {
	GetValueForDuplicationCheck() string
	GetDebugName() string
	GetID() string
	GetCategory() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

// Model is the model object used in Promise project.
type Model struct {
	ID        string
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GetID return the ID.
func (m *Model) GetID() string {
	return m.ID
}

// GetCategory return the category.
func (m *Model) GetCategory() string {
	return m.Category
}

// GetCreatedAt return the created at.
func (m *Model) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// GetUpdatedAt return the updated at.
func (m *Model) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}
