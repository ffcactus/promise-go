package model

import (
	"promise/base"
)

// Student is the model used in student project.
type Student struct {
	base.Model
	Name   string
	Age    int
	Phones []Phone
}

// GetDebugName return the debug name the model.
func (m *Student) GetDebugName() string {
	return m.Name
}

// GetValueForDuplicationCheck return the value for duplication check.
func (m *Student) GetValueForDuplicationCheck() string {
	return m.Name
}

// StudentCollectionMember is the member in student collection.
type StudentCollectionMember struct {
	base.CollectionMemberModel
	Name string
}

// StudentCollection is the collection of student.
type StudentCollection struct {
	base.CollectionModel
}

// NewModelMember return a new ModelMember
func (m *StudentCollection) NewModelMember() interface{} {
	return new(StudentCollectionMember)
}