package controller

import (
	"promise/base"
	"promise/task/object/dto"
	"promise/task/service"
)

var (
	updateTaskStep = base.ActionInfo{
		Name: "UpdateTaskStep",
		Request: &dto.UpdateTaskStepRequest{
			ActionRequest: base.ActionRequest{
				TemplateImpl: new(dto.UpdateTaskStepRequest),
			},
		},
		Service: &base.ActionService{
			TemplateImpl: new(service.UpdateTaskStep),
		},
	}
	updateTask = base.ActionInfo{
		Name: "UpdateTask",
		Request: &dto.UpdateTaskRequest{
			ActionRequest: base.ActionRequest{
				TemplateImpl: new(dto.UpdateTaskRequest),
			},
		},
		Service: &base.ActionService{
			TemplateImpl: new(service.UpdateTask),
		},
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
