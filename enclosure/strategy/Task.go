package strategy

import (
	"promise/enclosure/context"
	"prrmise/base"
	taskSDK "promise/sdk/task"
	taskDTO "promise/task/object/dto"
	taskModel "promise/task/object/model"	
)

// Task can be used as a base struct to support task operation.
type Task struct {

}

// StepStart will update the task process to step specified by name, and set the step state to running.
func (s Task) StepStart(c *context.Base, name string) {
	// TODO should use service error.
	_, errorResp, err := taskSDK.SetStepExecutionState(c.ID, name, taskModel.ExecutionStateRunning)
	if errorResp != nil {
		log.WithFields(log.Fields{
			"resource":    c.ID,
			"task":      c.TaskURL
			"step":      name
			"state":     state,
			"errorResp": errorResp[0].ID,
		}).Warn("Update task step execution state failed.")
	} else {
		log.WithFields(logFields {
			"resource": c.ID,
			"task": c.TaskURL,
			"step": name,
			"state": state,
		}).Debug("Update task step state done.")		
	}
}

// StepSuccess Set the step to finish.
func (s *ServerTask) StepSuccess(c *context.Base, step string) {
	request := new(taskDTO.UpdateTaskStepRequest)
	request.Name = stepName
	request.ExecutionState = &taskModel.ExecutionStateTerminated
	request.ExecutionResult = new(taskDTO.UpdateExecutionResultRequest)
	request.ExecutionResult.State = &taskModel.ExecutionResultStateFinished
	_, errorResp, err := taskSDK.UpdateStep(id, request)
	s.logUpdateStepResult(c, id, stepName, server, err, errorResp)
}