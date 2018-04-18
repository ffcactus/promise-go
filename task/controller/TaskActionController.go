package controller

import (
	"promise/base"
	"promise/task/object/dto"
)

var (
	updateTaskStep = base.ActionInfo{
		Name: "UpdateTaskStep",
		Request: &dto.UpdateTaskStepRequest{
			UpdateActionRequest: base.UpdateActionRequest{
				TemplateImpl: new(dto.UpdateTaskStepRequest),
			},
		},
		Service: taskService,
	}
	updateTask = base.ActionInfo{
		Name: "UpdateTask",
		Request: &dto.UpdateTaskRequest{
			UpdateActionRequest: base.UpdateActionRequest{
				TemplateImpl: new(dto.UpdateTaskRequest),
			},
		},
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
