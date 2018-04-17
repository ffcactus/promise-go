package service

import (
	"promise/base"
	"promise/sdk/event"
	"promise/task/db"
	"promise/task/object/dto"
)

var (
	taskDB = &db.TaskDB{
		DB: base.DB{
			TemplateImpl: new(db.TaskDB),
		},
	}

	eventService event.Service
)

// Task is the service for task.
type Task struct {
}

// GetCategory returns the category of this service.
func (s *Task) GetCategory() string {
	return base.CategoryTask
}

// NewResponse creates a new response DTO.
func (s *Task) NewResponse() base.ResponseInterface {
	return new(dto.GetTaskResponse)
}

// GetDB returns the DB implementation.
func (s *Task) GetDB() base.DBInterface {
	return taskDB
}

// GetEventService returns the event service implementation.
func (s *Task) GetEventService() base.EventServiceInterface {
	return eventService
}
