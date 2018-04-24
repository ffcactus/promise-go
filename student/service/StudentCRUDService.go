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
		TemplateImpl: new(db.Student),
	}
)

// StudentCRUDService is the service for student.
type StudentCRUDService struct {
	eventService event.Service
}

// Category returns the category of this service.
func (s *StudentCRUDService) Category() string {
	return base.CategoryStudent
}

// Response creates a new response DTO.
func (s *StudentCRUDService) Response() base.GetResponseInterface {
	return new(dto.GetStudentResponse)
}

// DB returns the DB implementation.
func (s *StudentCRUDService) DB() base.DBInterface {
	return studentDB
}

// EventService returns the event service implementation.
func (s *StudentCRUDService) EventService() base.EventServiceInterface {
	return s.eventService
}
