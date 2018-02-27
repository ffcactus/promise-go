package db

import (
	"promise/task/object/model"
)

// TaskDBInterface DB interface.
type TaskDBInterface interface {
	PostTask(task *model.Task) (*model.Task, error)
	GetTask(id string) *model.Task
	GetTaskCollection(start int, count int) (*model.TaskCollection, error)
	UpdateTask(id string, task *model.Task) (*model.Task, error)
}
