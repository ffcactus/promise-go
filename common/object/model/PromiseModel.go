package model

import (
	"time"
)

// PromiseModelInterface is the interface of PromsieModel.
type PromiseModelInterface interface {
	GetID() string
	GetCategory() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

// PromiseModel is the common model in Promise.
type PromiseModel struct {
	ID        string
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GetID return the ID.
func (m *PromiseModel) GetID() string {
	return m.ID
}

// GetCategory return the category.
func (m *PromiseModel) GetCategory() string {
	return m.Category
}

// GetCreatedAt return the created at.
func (m *PromiseModel) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// GetUpdatedAt return the updated at.
func (m *PromiseModel) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}
