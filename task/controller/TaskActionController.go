package controller

import (
	"promise/base"
	"promise/task/object/dto"
)

var (
	updateTaskStep = base.ActionInfo{
		Name:    "UpdateTaskStep",
		Type:    base.ActionTypeUpdate,
		Request: new(dto.UpdateTaskStepRequest),
		Service: taskService,
	}
	updateTask = base.ActionInfo{
		Name:    "UpdateTask",
		Type:    base.ActionTypeUpdate,
		Request: new(dto.UpdateTaskRequest),
		Service: taskService,
	}

	actionInfo = []base.ActionInfo{updateTaskStep, updateTask}
)

// TaskActionController is implements ActionControllerTemplateInterface.
type TaskActionController struct {
}

// ResourceName returns the name this controller handle of.
func (c *TaskActionController) ResourceName() string {
	return "task"
}

// ActionInfo returns the name this controller handle of.
func (c *TaskActionController) ActionInfo() []base.ActionInfo {
	return actionInfo
}
