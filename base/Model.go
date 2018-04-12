package base

import (
	"time"
)

// ModelInterface is the interface of Model.
type ModelInterface interface {
	GetValueForDuplicationCheck() string
	GetDebugName() string
}

// Model is the model object used in Promise project.
type Model struct {
	ID        string
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CollectionMemberModelInterface is the interface a Member should have
type CollectionMemberModelInterface interface {

}

// CollectionMemberModel is the collection member used in Promise project.
type CollectionMemberModel struct {
	ID       string
	Category string
}

// CollectionModel is a collection of model.
type CollectionModel struct {
	Start   int64
	Count   int64
	Total   int64
	Members []interface{}
}

// SubModel is the sub models in a model.
// For example the phone numbers of a person.
type SubModel struct {
	ID uint64
}
