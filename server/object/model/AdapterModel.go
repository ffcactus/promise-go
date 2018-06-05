package model

import (
	"promise/base"
)

// AdapterModel is the model.
type AdapterModel struct {
	base.Model
	Name string
}

// DebugInfo return the debug name the model.
func (m *AdapterModel) DebugInfo() string {
	return m.Name
}

// ValueForDuplicationCheck return the value for duplication check.
func (m *AdapterModel) ValueForDuplicationCheck() string {
	return ""
}

// AdapterModelCollectionMember is the member in collection.
type AdapterModelCollectionMember struct {
	base.CollectionMemberModel
	Name string
}

// AdapterModelCollection is the model of collection.
type AdapterModelCollection struct {
	base.CollectionModel
}

// NewModelMember return a new ModelMember
func (m *AdapterModelCollection) NewModelMember() interface{} {
	return new(AdapterModelCollectionMember)
}
