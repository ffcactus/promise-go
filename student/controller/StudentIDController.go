package controller

import (
	"promise/base"
	"promise/student/object/dto"
)

// StudentIDController is the ID controller of the student.
type StudentIDController struct {
}

// ResourceName returns the name this controller handle of.
func (c *StudentIDController) ResourceName() string {
	return "student"
}

// Response creates a new response DTO.
func (c *StudentIDController) Response() base.GetResponseInterface {
	return new(dto.GetStudentResponse)
}

// Service returns the service.
func (c *StudentIDController) Service() base.CRUDServiceInterface {
	return StudentService
}
