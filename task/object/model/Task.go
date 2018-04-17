package model

import (
	"promise/base"
)

// ExecutionState The execution state.
type ExecutionState = string

var (
	// ExecutionStateReady Execution state enum.
	ExecutionStateReady ExecutionState = "Ready"
	// ExecutionStateRunning Execution state enum.
	ExecutionStateRunning ExecutionState = "Running"
	// ExecutionStateSuspended Execution state enum.
	ExecutionStateSuspended ExecutionState = "Suspended"
	// ExecutionStateTerminated Execution state enum.
	ExecutionStateTerminated ExecutionState = "Terminated"
)

// ExecutionResultState Execution result state type.
type ExecutionResultState = string

var (
	// ExecutionResultStateFinished Execution result state enum.
	ExecutionResultStateFinished ExecutionResultState = "Finished"
	// ExecutionResultStateWarning Execution result state enum.
	ExecutionResultStateWarning ExecutionResultState = "Warning"
	// ExecutionResultStateError Execution result state enum.
	ExecutionResultStateError ExecutionResultState = "Error"
	// ExecutionResultStateAbort Execution result state enum.
	ExecutionResultStateAbort ExecutionResultState = "Abort"
	// ExecutionResultStateUnknown Execution result state enum.
	ExecutionResultStateUnknown ExecutionResultState = "Unknown"
)

// IsValidExecutionState check if the state is a valid execution state.
func IsValidExecutionState(state string) bool {
	switch state {
	case ExecutionStateReady,
		ExecutionStateRunning,
		ExecutionStateSuspended,
		ExecutionStateTerminated:
		return true
	default:
		return false
	}
}

// IsValidExecutionResultState check if the state is a valid execution result state.
func IsValidExecutionResultState(state string) bool {
	switch state {
	case ExecutionResultStateFinished,
		ExecutionResultStateWarning,
		ExecutionResultStateError,
		ExecutionResultStateAbort,
		ExecutionResultStateUnknown:
		return true
	default:
		return false
	}
}

// ExecutionResult Used by Task and it's TaskStep
type ExecutionResult struct {
	State   ExecutionResultState
	Message *base.Message
}

// TaskStep The TaskStep represents each planned steps in a task.
type TaskStep struct {
	MessageID           *string
	Name                string
	Description         *string
	ExpectedExecutionMs uint64
	ExecutionState      ExecutionState
	ExecutionResult     ExecutionResult
}

// Task Task object.
type Task struct {
	base.Model
	MessageID           *string
	Name                string
	ParentTask          *string
	Description         *string
	ExecutionState      ExecutionState
	CreatedByName       string
	CreatedByURI        string
	TargetName          string
	TargetURI           string
	ExpectedExecutionMs uint64
	Percentage          uint32
	CurrentStep         string
	TaskSteps           []TaskStep
	ExecutionResult     ExecutionResult
}

// GetDebugName return the debug name the model.
func (m *Task) GetDebugName() string {
	return m.Name
}

// GetValueForDuplicationCheck return the value for duplication check.
func (m *Task) GetValueForDuplicationCheck() string {
	return m.Name
}

// TaskCollectionMember is the member in student collection.
type TaskCollectionMember struct {
	base.CollectionMemberModel
	Name            string
	Description     *string
	ExecutionState  ExecutionState
	Percentage      uint32
	CurrentStep     string
	ExecutionResult ExecutionResult
}

// TaskCollection is the collection of student.
type TaskCollection struct {
	base.CollectionModel
}

// NewModelMember return a new ModelMember
func (m *TaskCollection) NewModelMember() interface{} {
	return new(TaskCollectionMember)
}
