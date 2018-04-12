package controller

import (
	"promise/base"
	"promise/student/object/dto"
)

// StudentIDController is the ID controller of the student.
type StudentIDController struct {
}

// GetResourceName returns the name this controller handle of.
func (c *StudentIDController) GetResourceName() string {
	return "student"
}

// NewResponse creates a new response DTO.
func (c *StudentIDController) NewResponse() base.ResponseInterface {
	response := new(dto.GetStudentResponse)
	response.TemplateImpl = response
	return response
}

// GetService returns the service.
func (c *StudentIDController) GetService() base.ServiceInterface {
	return StudentService
}
