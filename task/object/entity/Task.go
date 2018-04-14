package entity

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/task/object/model"
)

// Task is the entity of the task.
type Task struct {
	base.Entity
	MessageID           *string              `gorm:"column:MessageID"`
	Name                string               `gorm:"column:Name"`
	Description         *string              `gorm:"column:Description"`
	ExecutionState      model.ExecutionState `gorm:"column:ExecutionState"`
	CreatedByName       string               `gorm:"column:CreatedByName"`
	CreatedByURI        string               `gorm:"column:CreatedByURI"`
	TargetName          string               `gorm:"column:TargetName"`
	TargetURI           string               `gorm:"column:TargetURI"`
	ExpectedExecutionMs uint64               `gorm:"column:ExpectedExecutionMs"`
	Percentage          int                  `gorm:"column:Percentage"`
	CurrentStep         string               `gorm:"column:CurrentStep"`
	// Use string to represent.
	TaskSteps string `gorm:"column:TaskSteps"`
	// Use string to represent.
	ExecutionResult string `gorm:"column:ExecutionResult"`
}

// TableName will set the table name.
func (Task) TableName() string {
	return "Task"
}

// GetID return the ID.
func (e *Task) GetID() string {
	return e.ID
}

// SetID set the ID.
func (e *Task) SetID(id string) {
	e.ID = id
}

// GetDebugName return the debug name of this entity.
func (e *Task) GetDebugName() string {
	return e.Name
}

// GetPropertyNameForDuplicationCheck return the property name used for duplication check.
func (e *Task) GetPropertyNameForDuplicationCheck() string {
	return ""
}

// GetPreload return the property names that need to be preload.
func (e *Task) GetPreload() []string {
	return nil
}

// GetAssociation return all the assocations that need to delete when deleting a resource.
func (e *Task) GetAssociation() []interface{} {
	return []interface{}{}
}

// GetTables returns the tables to delete when you want delete all the resources.
func (e *Task) GetTables() []interface{} {
	return []interface{}{Task{}}
}

// GetFilterNameList return all the property name that can be used in filter.
func (e *Task) GetFilterNameList() []string {
	return []string{"Name", "ExecutionState", "TargetName", "TargetURI", "CreatedByName", "CreatedByURI", "Percentage", "CurrentStep", "ExecutionResult"}
}

// Load will load data from model.
func (e *Task) Load(i base.ModelInterface) error {
	m, ok := i.(*model.Task)
	if !ok {
		log.Error("entity.Task.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	base.EntityLoad(&e.Entity, &m.Model)
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
	e.CurrentStep = m.CurrentStep
	e.TaskSteps = base.StructToString(m.TaskSteps)
	e.ExecutionResult = base.StructToString(m.ExecutionResult)
	return nil
}

// ToModel converts the entity to model.
func (e *Task) ToModel() base.ModelInterface {
	m := new(model.Task)
	base.EntityToModel(&e.Entity, &m.Model)
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
	m.CurrentStep = e.CurrentStep
	base.StringToStruct(e.ExecutionResult, &m.ExecutionResult)
	base.StringToStruct(e.TaskSteps, &m.TaskSteps)
	return m
}

// ToCollectionMember convert the entity to member.
func (e *Task) ToCollectionMember() base.CollectionMemberModelInterface {
	m := new(model.TaskCollectionMember)
	base.EntityToMember(&e.Entity, &m.CollectionMemberModel)
	m.Name = e.Name
	m.Description = e.Description
	m.CurrentStep = e.CurrentStep
	m.ExecutionState = e.ExecutionState
	m.Percentage = e.Percentage
	base.StringToStruct(e.ExecutionResult, &m.ExecutionResult)
	return m
}
