package service

import (
	"promise/base"
	"promise/sdk/event"
	"promise/student/db"
	"promise/student/object/dto"
)

var (
	StudentDB = &base.DB{
		TemplateImpl: new(db.StudentDB),
	}
)

// StudentService is the service for student.
type StudentService struct {
	EventService event.Service
}

// NewResponse creates a new response DTO.
func (s *StudentService) NewResponse() base.ResponseInterface {
	return new(dto.GetStudentResponse)
}

// GetDB returns the DB implementation.
func (s *StudentService) GetDB() base.DBInterface {
	return StudentDB
}

// GetEventService returns the event service implementation.
func (s *StudentService) GetEventService() base.EventServiceInterface {
	return s.EventService
}
