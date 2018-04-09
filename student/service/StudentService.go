package service

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/sdk/event"
	"promise/student/db"
	"promise/student/object/dto"
)

// StudentService is the service for student.
type StudentService struct {
	// base.Service
	DB           db.StudentDB
	EventService event.Service
}

// NewResponse creates a new response DTO.
func (s *StudentService) NewResponse() base.ResponseInterface {
	return new(dto.GetStudentResponse)
}

// GetDB returns the DB implementation.
func (s *StudentService) GetDB() base.DBInterface {
	log.Info("--- GetDB() in StudentService.")
	return &s.DB
}

// GetEventService returns the event service implementation.
func (s *StudentService) GetEventService() base.EventServiceInterface {
	return s.EventService
}
