package entity

import (
	commonEntity "promise/common/object/entity"
	"promise/common/util"
	"promise/task/object/model"
)

// Task Task object.
type Task struct {
	commonEntity.PromiseEntity
	ParentTask          *string `gorm:"column:ParentTask"`
	MessageID           *string `gorm:"column:MessageID"`
	Name                string `gorm:"column:Name"`
	Description         *string `gorm:"column:Description"`
	ExecutionState      model.ExecutionState `gorm:"column:ExecutionState"`
	CreatedByName       string `gorm:"column:CreatedByName"`
	CreatedByURI        string `gorm:"column:CreatedByURI"`
	TargetName          string `gorm:"column:TargetName"`
	TargetURI           string `gorm:"column:TargetURI"`
	ExpectedExecutionMs uint64 `gorm:"column:ExpectedExecutionMs"`
	Percentage          int `gorm:"column:Percentage"`
	CurrentStep         string `gorm:"column:CurrentStep"`
	// Use string to represent.
	TaskSteps           string  `gorm:"column:TaskSteps"`
	SubTasks            []Task `gorm:"ForeignKey:ParentTask;column:SubTasks"`
	// Use string to represent.
	ExecutionResult     string  `gorm:"column:ExecutionResult"`
}

// TableName will set the table name.
func (Task) TableName() string {
	return "Task"
}

// ToModel will create a new model from entity.
func (e *Task) ToModel() *model.Task {
	m := new(model.Task)
	m.PromiseModel = e.PromiseEntity.ToModel()
	m.MessageID = e.MessageID
	m.Name = e.Name
	m.Description = e.Description
	m.ExecutionState = e.ExecutionState
	m.CreatedByName = e.CreatedByName
	m.CreatedByURI = e.CreatedByURI
	m.TargetName = e.TargetName
	m.TargetURI = e.TargetURI
	m.ExpectedExecutionMs = e.ExpectedExecutionMs
	m.Percentage = e.Percentage
	m.CreatedAt = e.CreatedAt
	m.UpdatedAt = e.UpdatedAt
	m.CurrentStep = e.CurrentStep
	util.StringToStruct(e.ExecutionResult, &m.ExecutionResult)
	util.StringToStruct(e.TaskSteps, &m.TaskSteps)
	for _, v := range e.SubTasks {
		vv := v.ToModel()
		m.SubTasks = append(m.SubTasks, *vv)
	}
	return m
}

// Load will load data from model.
func (e *Task) Load(m *model.Task) {
	e.PromiseEntity.Load(m.PromiseModel)
	e.MessageID = m.MessageID
	e.Name = m.Name
	e.Description = m.Description
	e.ExecutionState = m.ExecutionState
	e.CreatedByName = m.CreatedByName
	e.CreatedByURI = m.CreatedByURI
	e.TargetName = m.TargetName
	e.TargetURI = m.TargetURI
	e.ExpectedExecutionMs = m.ExpectedExecutionMs
	e.Percentage = m.Percentage
	e.CreatedAt = m.CreatedAt
	e.UpdatedAt = m.UpdatedAt
	e.CurrentStep = m.CurrentStep
	e.TaskSteps = util.StructToString(m.TaskSteps)
	e.ExecutionResult = util.StructToString(m.ExecutionResult)
}