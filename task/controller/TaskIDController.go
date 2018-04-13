package controller

import (
	"promise/base"
	"promise/task/object/dto"
)

// TaskIDController Task controller
type TaskIDController struct {
}

// GetResourceName returns the name this controller handle of.
func (c *TaskIDController) GetResourceName() string {
	return "student"
}

// NewResponse creates a new response DTO.
func (c *TaskIDController) NewResponse() base.ResponseInterface {
	response := new(dto.GetTaskResponse)
	response.TemplateImpl = response
	return response
}

// GetService returns the service.
func (c *TaskIDController) GetService() base.ServiceInterface {
	return taskService
}
