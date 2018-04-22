package controller

import (
	"promise/base"
	"promise/task/object/dto"
)

// TaskIDController Task controller
type TaskIDController struct {
}

// ResourceName returns the name this controller handle of.
func (c *TaskIDController) ResourceName() string {
	return "task"
}

// Response creates a new response DTO.
func (c *TaskIDController) Response() base.GetResponseInterface {
	return new(dto.GetTaskResponse)
}

// Service returns the service.
func (c *TaskIDController) Service() base.CRUDServiceInterface {
	return taskService
}
