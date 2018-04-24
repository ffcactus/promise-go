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

// TaskAction is implements ActionControllerTemplateInterface.
type TaskAction struct {
}

// ResourceName returns the name this controller handle of.
func (c *TaskAction) ResourceName() string {
	return "task"
}

// ActionInfo returns the name this controller handle of.
func (c *TaskAction) ActionInfo() []base.ActionInfo {
	return actionInfo
}
