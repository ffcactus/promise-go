package entity

import (
	"promise/task/object/model"
	"time"
)

// Task Task object.
type Task struct {
	ID                  string `gorm:"primary_key"`
	ParentTask          *string
	MessageID           *string
	Name                string
	Description         string
	ExecutionState      model.ExecutionState
	CreatedByName       string
	CreatedByURI        string
	TargetName          string
	TargetURI           string
	ExpectedExecutionMs int
	Percentage          int
	CreatedAt           time.Time
	UpdatedAt           time.Time
	CurrentStep         string
	TaskSteps           string // Use string to represent.
	SubTasks            []Task `gorm:"ForeignKey:ParentTask"`
	ExecutionResult     string // Use string to represent.
}
