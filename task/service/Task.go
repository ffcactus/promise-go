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

// Task is the concrete service.
type Task struct {
}

// Category returns the category of this service.
func (s *Task) Category() string {
	return base.CategoryTask
}

// Response creates a new response DTO.
func (s *Task) Response() base.GetResponseInterface {
	return new(dto.GetTaskResponse)
}

// DB returns the DB implementation.
func (s *Task) DB() base.DBInterface {
	return taskDB
}

// EventService returns the event service implementation.
func (s *Task) EventService() base.EventServiceInterface {
	return eventService
}
