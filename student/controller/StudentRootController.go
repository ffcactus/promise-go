package controller

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/student/object/dto"
	"promise/student/service"
)

var (
	// StudentService is the concrete service in student project.
	StudentService = &base.Service{
		Interface: new(service.StudentService),
	}
)

// StudentRootController is the root controller of the student.
type StudentRootController struct {
}

// GetResourceName return the name of the resource this controller deal with.
func (c *StudentRootController) GetResourceName() string {
	return "student"
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
	log.Info("--- Get service in StudentController.")
	return StudentService
}

// PostCallback is the callback that should call service.
// func (c *StudentRootController) PostCallback(request base.RequestInterface) (base.ModelInterface, []base.MessageInterface) {
// 	return request.ToModel(), nil
// }
