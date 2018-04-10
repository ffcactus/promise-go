package controller

import (
	"promise/base"
	"promise/student/object/dto"
	"promise/student/service"
)

var (
	// StudentService is the concrete service in student project.
	StudentService = &base.Service{
		TemplateImpl: new(service.StudentService),
	}
)

// StudentRootController is the root controller of the student.
type StudentRootController struct {
}

// NewRequest creates a new request DTO.
func (c *StudentRootController) NewRequest() base.RequestInterface {
	return new(dto.PostStudentRequest)
}

// NewResponse creates a new response DTO.
func (c *StudentRootController) NewResponse() base.ResponseInterface {
	return new(dto.GetStudentResponse)
}

// GetService returns the service.
func (c *StudentRootController) GetService() base.ServiceInterface {
	return StudentService
}
