package model

import (
	"promise/common/object/message"
	"time"
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

// ExecutionResult Used by Task and it's TaskStep
type ExecutionResult struct {
	State   ExecutionResultState
	Message *message.Message
}

// TaskStep The TaskStep represents each planned steps in a task.
type TaskStep struct {
	MessageID           *string
	Name                string
	Description         string
	ExpectedExecutionMs int
	ExecutionState      ExecutionState
	ExecutionResult     ExecutionResult
}

// Task Task object.
type Task struct {
	ID                  string
	URI                 string
	MessageID           *string
	Name                string
	ParentTask          *string
	Description         string
	ExecutionState      ExecutionState
	CreatedByName       string
	CreatedByURI        string
	TargetName          string
	TargetURI           string
	ExpectedExecutionMs int
	Percentage          int
	CreatedAt           time.Time
	UpdatedAt           time.Time
	CurrentStep         string
	TaskSteps           []TaskStep
	SubTasks            []Task
	ExecutionResult     ExecutionResult
}
