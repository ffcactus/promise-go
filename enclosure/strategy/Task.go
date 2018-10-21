package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/enclosure/context"
	taskSDK "promise/sdk/task"
	taskDTO "promise/task/object/dto"
	taskModel "promise/task/object/model"
)

// StepStart will update the task process to step specified by name, and set the step state to running.
func StepStart(c *context.Base, name string) {
	// TODO should use service error.
	_, err := taskSDK.SetStepExecutionState(c.TaskID, name, taskModel.ExecutionStateRunning)
	if err != nil {
		log.WithFields(log.Fields{
			"resource": c.ID,
			"task":     c.TaskID,
			"step":     name,
			"to":       taskModel.ExecutionStateRunning,
			"error":    err,
		}).Warn("Update task step execution state failed.")
	} else {
		log.WithFields(log.Fields{
			"resource": c.ID,
			"task":     c.TaskID,
			"step":     name,
		}).Debug("Update task step state done.")
	}
}

func newUpdateTaskStepRequest(name, state string) *taskDTO.UpdateTaskStepRequest {
	request := taskDTO.UpdateTaskStepRequest{}
	request.Name = name
	request.ExecutionState = &taskModel.ExecutionStateTerminated
	request.ExecutionResult = new(taskDTO.UpdateExecutionResultRequest)
	request.ExecutionResult.State = &state
	return &request
}

func doStepLog(c *context.Base, name, state string, err error) {
	log.WithFields(log.Fields{
		"resource": c.ID,
		"task":     c.TaskID,
		"step":     name,
		"to":       state,
		"error":    err,
	}).Warn("Set task step failed.")
}

// StepFinish Set the step to finish.
func StepFinish(c *context.Base, name string) {
	state := taskModel.ExecutionResultStateFinished
	req := newUpdateTaskStepRequest(name, state)
	if _, err := taskSDK.UpdateStep(c.TaskID, req); err != nil {
		doStepLog(c, name, state, err)
	}
}

// StepWarning Set the step to warning.
func StepWarning(c *context.Base, name string) {
	state := taskModel.ExecutionResultStateWarning
	req := newUpdateTaskStepRequest(name, state)
	if _, err := taskSDK.UpdateStep(c.TaskID, req); err != nil {
		doStepLog(c, name, state, err)
	}
}

// StepError Set the step to error.
func StepError(c *context.Base, name string) {
	state := taskModel.ExecutionResultStateError
	req := newUpdateTaskStepRequest(name, state)
	if _, err := taskSDK.UpdateStep(c.TaskID, req); err != nil {
		doStepLog(c, name, state, err)
	}
}
