package controller

import (
	"promise/base"
	"promise/task/object/dto"
	"promise/task/service"
)

var (
	// StudentService is the service used in Student controller.
	taskService = &base.Service{
		TemplateImpl: new(service.TaskService),
	}
)

// TaskRootController is the root controller for task.
type TaskRootController struct {
}

// GetResourceName returns the name this controller handle of.
func (c *TaskRootController) GetResourceName() string {
	return "task"
}

// NewRequest creates a new request DTO.
func (c *TaskRootController) NewRequest() base.RequestInterface {
	request := new(dto.PostTaskRequest)
	request.TemplateImpl = request
	return request
}

// NewResponse creates a new response DTO.
func (c *TaskRootController) NewResponse() base.ResponseInterface {
	response := new(dto.GetTaskResponse)
	response.TemplateImpl = response
	return response
}

// GetService returns the service.
func (c *TaskRootController) GetService() base.ServiceInterface {
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
