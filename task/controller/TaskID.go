package controller

import (
	"promise/base"
	"promise/task/object/dto"
)

// TaskID Task controller
type TaskID struct {
}

// ResourceName returns the name this controller handle of.
func (c *TaskID) ResourceName() string {
	return "task"
}

// Response creates a new response DTO.
func (c *TaskID) Response() base.GetResponseInterface {
	return new(dto.GetTaskResponse)
}

// Service returns the service.
func (c *TaskID) Service() base.CRUDServiceInterface {
	return taskService
}
