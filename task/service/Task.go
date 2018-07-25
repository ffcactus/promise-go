package service

import (
	"promise/base"
	"promise/task/db"
	"promise/task/object/dto"
)

var (
	taskDB = &db.Task{
		DB: base.DB{
			TemplateImpl: new(db.Task),
		},
	}
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
