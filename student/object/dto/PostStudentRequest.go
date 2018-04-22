package dto

import (
	"promise/base"
	"promise/student/object/model"
)

// PostStudentRequest is the DTO to post a student.
type PostStudentRequest struct {
	Name   string  `json:"Name"`
	Age    int     `json:"Age"`
	Phones []Phone `json:"Phones"`
}

// NewInstance creates a new instance.
func (dto *PostStudentRequest) NewInstance() base.RequestInterface {
	return new(PostStudentRequest)
}

// IsValid return if the request is valid.
func (dto *PostStudentRequest) IsValid() *base.Message {
	return nil
}

// DebugInfo return the name for debug.
func (dto *PostStudentRequest) DebugInfo() string {
	return dto.Name
}

// ToModel convert the DTO to model.
func (dto *PostStudentRequest) ToModel() base.ModelInterface {
	var m model.Student
	m.Category = base.CategoryStudent
	m.Name = dto.Name
	m.Age = dto.Age
	for _, v := range dto.Phones {
		phone := v.ToModel()
		m.Phones = append(m.Phones, *phone)
	}
	return &m
}
