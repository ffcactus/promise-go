package controller

import (
	"promise/base"
	"promise/student/object/dto"
	"promise/student/service"
)

var (
	// StudentService is the concrete service in student project.
	StudentService = &base.CRUDService{
		TemplateImpl: new(service.StudentCRUDService),
	}
)

// StudentRootController is the root controller of the student.
type StudentRootController struct {
}

// ResourceName returns the name this controller handle of.
func (c *StudentRootController) ResourceName() string {
	return "student"
}

// Request creates a new request DTO.
func (c *StudentRootController) Request() base.PostRequestInterface {
	return new(dto.PostStudentRequest)
}

// Response creates a new response DTO.
func (c *StudentRootController) Response() base.GetResponseInterface {
	return new(dto.GetStudentResponse)
}

// Service returns the service.
func (c *StudentRootController) Service() base.CRUDServiceInterface {
	return StudentService
}

// ConvertCollectionModel convert data to concrete DTO.
func (c *StudentRootController) ConvertCollectionModel(m *base.CollectionModel) (interface{}, error) {
	ret := new(dto.GetStudentCollectionResponse)
	if err := ret.Load(m); err != nil {
		return nil, err
	}
	return ret, nil
}
