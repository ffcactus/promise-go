package dto

import (
	"promise/base"
	"promise/apps"
	"promise/student/object/model"
)

// PostStudentRequest is the DTO to post a student.
type PostStudentRequest struct {
	base.Request
	Name string
	Age int
}

// GetDebugName return the name for debug.
func (dto *PostStudentRequest) GetDebugName() string {
	return dto.Name
}

// ToModel convert the DTO to model.
func (dto *PostStudentRequest) ToModel() base.ModelInterface {
	var m model.Student
	m.Category = apps.CategoryStudent
	return &m	
}