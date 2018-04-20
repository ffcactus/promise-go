package context

import (
	m "promise/server/object/model"
	taskDto "promise/task/object/dto"
	taskModel "promise/task/object/model"
	// "github.com/astaxie/beego"
)

// TaskHandlerInterface Task handler interface.
type TaskHandlerInterface interface {
	CreateTask(server *m.Server, request *taskDto.PostTaskRequest)
	UpdateStepExecutionState(server *m.Server, state taskModel.ExecutionState)
	UpdateStepExecutionResultState(server *m.Server, state taskModel.ExecutionResultState)
}
