package controller

import (
	"promise/base"
	"promise/task/object/dto"
)

// NewUpdateTaskStepRequest returns a new request.
func NewUpdateTaskStepRequest() base.UpdateActionRequestInterface {
	return &dto.UpdateTaskStepRequest{
		UpdateActionRequest: base.UpdateActionRequest{
			TemplateImpl: new(dto.UpdateTaskStepRequest),
		},
	}
}

// NewUpdateTaskRequest returns a new request.
func NewUpdateTaskRequest() base.UpdateActionRequestInterface {
	return &dto.UpdateTaskRequest{
		UpdateActionRequest: base.UpdateActionRequest{
			TemplateImpl: new(dto.UpdateTaskRequest),
		},
	}
}

var (
	updateTaskStep = base.ActionInfo{
		Name:    "UpdateTaskStep",
		Request: NewUpdateTaskStepRequest,
		Service: taskService,
	}
	updateTask = base.ActionInfo{
		Name:    "UpdateTask",
		Request: NewUpdateTaskRequest,
		Service: taskService,
	}

	actionInfo = []base.ActionInfo{updateTaskStep, updateTask}
)

// TaskActionController is implements ActionControllerTemplateInterface.
type TaskActionController struct {
}

// GetResourceName returns the name this controller handle of.
func (c *TaskActionController) GetResourceName() string {
	return "task"
}

// GetActionInfo returns the name this controller handle of.
func (c *TaskActionController) GetActionInfo() []base.ActionInfo {
	return actionInfo
}
