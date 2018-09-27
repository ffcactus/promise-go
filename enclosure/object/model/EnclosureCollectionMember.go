package model

import (
	"promise/base"
)

// EnclosureCollectionMember is the member in collection.
type EnclosureCollectionMember struct {
	base.CollectionMemberModel
	Name        string
	Description string
	Type        string
	State       string
	Health      string
}

// EnclosureCollection is the model of collection.
type EnclosureCollection struct {
	base.CollectionModel
}

// NewModelMember return a new ModelMember
func (m *EnclosureCollection) NewModelMember() interface{} {
	return new(EnclosureCollectionMember)
}
