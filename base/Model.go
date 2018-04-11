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

// CollectionMemberModelInterface is the interface a Member should have
type CollectionMemberModelInterface interface {
	GetID() string
	GetCategory() string
}

// CollectionMemberModel is the collection member used in Promise project.
type CollectionMemberModel struct {
	ID       string
	Category string
}

// GetID return the ID.
func (m *CollectionMemberModel) GetID() string {
	return m.ID
}

// GetCategory return the category.
func (m *CollectionMemberModel) GetCategory() string {
	return m.Category
}

// CollectionModelTemplateInterface is the interface that a concrete implement should have.
type CollectionModelTemplateInterface interface {
	NewCollectionMemberModel()
}

// CollectionModel is a collection of model.
type CollectionModel struct {
	Start   int64
	Count   int64
	Total   int64
	Members []interface{}
}

// SetStart set the start.
func (m *CollectionModel) SetStart(v int64) {
	m.Start = v
}

// SetCount set the count.
func (m *CollectionModel) SetCount(v int64) {
	m.Count = v
}

// SetTotal set the total.
func (m *CollectionModel) SetTotal(v int64) {
	m.Total = v
}

// SubModel is the sub models in a model.
// For example the phone numbers of a person.
type SubModel struct {
	ID uint64
}
