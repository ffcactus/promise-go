package db

import (
	"promise/task/object/model"
	"promise/task/object/dto"
)

// TaskDBInterface DB interface.
type TaskDBInterface interface {
	PostTask(task *model.Task) (*model.Task, error)
	GetTask(id string) *model.Task
	GetTaskCollection(start int64, count int64, filter string) (*model.TaskCollection, error)
	UpdateTask(id string, updateRequest *dto.UpdateTaskRequest) (bool, *model.Task, bool, error)
	UpdateTaskStep(id string, updateRequest *dto.UpdateTaskStepRequest) (bool, *model.Task, bool, error)
}
