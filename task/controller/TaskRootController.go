package controller

import (
	"promise/base"
	"promise/task/object/dto"
	"promise/task/service"
)

var (
	// StudentService is the service used in Student controller.
	taskService = &base.CRUDService{
		TemplateImpl: new(service.Task),
	}
)

// TaskRootController is the root controller for task.
type TaskRootController struct {
}

// ResourceName returns the name this controller handle of.
func (c *TaskRootController) ResourceName() string {
	return "task"
}

// Request creates a new request DTO.
func (c *TaskRootController) Request() base.PostRequestInterface {
	return new(dto.PostTaskRequest)
}

// Response creates a new response DTO.
func (c *TaskRootController) Response() base.GetResponseInterface {
	return new(dto.GetTaskResponse)
}

// Service returns the service.
func (c *TaskRootController) Service() base.CRUDServiceInterface {
	return taskService
}

// ConvertCollectionModel convert data to concrete DTO.
func (c *TaskRootController) ConvertCollectionModel(m *base.CollectionModel) (interface{}, error) {
	ret := new(dto.GetTaskCollectionResponse)
	if err := ret.Load(m); err != nil {
		return nil, err
	}
	return ret, nil
}
