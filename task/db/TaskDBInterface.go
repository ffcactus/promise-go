package db

import (
	"promise/task/object/model"
)

// TaskDBInterface DB interface.
type TaskDBInterface interface {
	PostTask(task *model.Task) (*model.Task, error)
	GetTask(id string) *model.Task
	GetTaskCollection(start int64, count int64, filter string) (*model.TaskCollection, error)
	UpdateTask(id string, task *model.Task) (*model.Task, error)
}
