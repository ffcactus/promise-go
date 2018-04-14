package dto

import (
	"promise/base"
	"promise/student/object/model"
)

// PostStudentRequest is the DTO to post a student.
type PostStudentRequest struct {
	base.Request
	Name   string  `json:"Name"`
	Age    int     `json:"Age"`
	Phones []Phone `json:"Phones"`
}

// IsValid return if the request is valid. 
func (dto *PostStudentRequest) IsValid() *base.Message {
	return nil
}

// GetDebugName return the name for debug.
func (dto *PostStudentRequest) GetDebugName() string {
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
