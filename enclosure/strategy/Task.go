package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/enclosure/context"
	taskSDK "promise/sdk/task"
	taskDTO "promise/task/object/dto"
	taskModel "promise/task/object/model"
)

// StepStart will update the task process to step specified by name, and set the step state to running.
func StepStart(c *context.Base, name string) {
	// TODO should use service error.
	_, errorResp, err := taskSDK.SetStepExecutionState(c.TaskID, name, taskModel.ExecutionStateRunning)
	if errorResp != nil || err != nil {
		log.WithFields(log.Fields{
			"resource":  c.ID,
			"task":      c.TaskID,
			"step":      name,
			"to":        taskModel.ExecutionStateRunning,
			"error":     err,
			"errorResp": errorResp[0].ID,
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

func doStepLog(c *context.Base, name, state string, err error, errorResp []base.ErrorResponse) {
	log.WithFields(log.Fields{
		"resource":  c.ID,
		"task":      c.TaskID,
		"step":      name,
		"to":        state,
		"error":     err,
		"errorResp": errorResp[0].ID,
	}).Warn("Set task step failed.")
}

// StepFinish Set the step to finish.
func StepFinish(c *context.Base, name string) {
	state := taskModel.ExecutionResultStateFinished
	req := newUpdateTaskStepRequest(name, state)
	if _, errorResp, err := taskSDK.UpdateStep(c.TaskID, req); err != nil || errorResp != nil {
		doStepLog(c, name, state, err, errorResp)
	}
}

// StepWarning Set the step to warning.
func StepWarning(c *context.Base, name string) {
	state := taskModel.ExecutionResultStateWarning
	req := newUpdateTaskStepRequest(name, state)
	if _, errorResp, err := taskSDK.UpdateStep(c.TaskID, req); err != nil || errorResp != nil {
		doStepLog(c, name, state, err, errorResp)
	}
}

// StepError Set the step to error.
func StepError(c *context.Base, name string) {
	state := taskModel.ExecutionResultStateError
	req := newUpdateTaskStepRequest(name, state)
	if _, errorResp, err := taskSDK.UpdateStep(c.TaskID, req); err != nil || errorResp != nil {
		doStepLog(c, name, state, err, errorResp)
	}
}
