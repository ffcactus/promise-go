package service

import (
	"promise/base"
	"promise/sdk/event"
	"promise/student/db"
	"promise/student/object/dto"
)

var (
	// TODO should I put it inside the struct?
	studentDB = &base.DB{
		TemplateImpl: new(db.StudentDB),
	}
)

// StudentService is the service for student.
type StudentService struct {
	EventService event.Service
}

// GetCategory returns the category of this service.
func (s *StudentService) GetCategory() string {
	return base.CategoryStudent
}

// NewResponse creates a new response DTO.
func (s *StudentService) NewResponse() base.ResponseInterface {
	return new(dto.GetStudentResponse)
}

// GetDB returns the DB implementation.
func (s *StudentService) GetDB() base.DBInterface {
	return studentDB
}

// GetEventService returns the event service implementation.
func (s *StudentService) GetEventService() base.EventServiceInterface {
	return s.EventService
}
