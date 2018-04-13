package service

import (
	"promise/base"
	"promise/sdk/event"
	"promise/task/db"
	"promise/task/object/dto"
)

var (
	taskDB = &base.DB{
		TemplateImpl: new(db.TaskDB),
	}
)

// TaskService is the service for student.
type TaskService struct {
	EventService event.Service
}

// GetCategory returns the category of this service.
func (s *TaskService) GetCategory() string {
	return base.CategoryTask
}

// NewResponse creates a new response DTO.
func (s *TaskService) NewResponse() base.ResponseInterface {
	return new(dto.GetTaskResponse)
}

// GetDB returns the DB implementation.
func (s *TaskService) GetDB() base.DBInterface {
	return taskDB
}

// GetEventService returns the event service implementation.
func (s *TaskService) GetEventService() base.EventServiceInterface {
	return s.EventService
}
