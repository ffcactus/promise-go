package action

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/enclosure/context"
	"promise/enclosure/db"
	"promise/enclosure/object/constvalue"
	"promise/enclosure/object/model"
	taskSDK "promise/sdk/task"
	taskDTO "promise/task/object/dto"
)

// Refresh is the refresh action.
type Refresh struct {
	sub []Action
}

// NewRefreshAction creates the refresh action.
func NewRefreshAction(ctx *context.RefreshContext) Action {
	ret := Refresh{}

	if base.ContainsString(ctx.Request.Targets, model.RefreshManager) {
		ret.Add(NewRefreshManager())
	}
	if base.ContainsString(ctx.Request.Targets, model.RefreshServer) {
		ret.Add(NewRefreshServer())
	}
	if base.ContainsString(ctx.Request.Targets, model.RefreshSwitch) {
		ret.Add(NewRefreshSwitch())
	}

	return &ret
}

// Add the sub actions.
func (s *Refresh) Add(sub Action) {
	s.sub = append(s.sub, sub)
}

// Name returns the name of the action.
func (Refresh) Name() string {
	return "Refresh Enclosure"
}

// MessageID returns the message ID of the action.
func (Refresh) MessageID() string {
	return constvalue.RefreshTaskID
}

// Description returns the description of the action.
func (Refresh) Description() string {
	return "Refresh enclosure's settings and component."
}

// ExpectedExecutionMs returns the expected execution time in ms of the action.
func (Refresh) ExpectedExecutionMs() uint64 {
	return 0
}

// Task returns the post task request.
func (s *Refresh) Task() *taskDTO.PostTaskRequest {
	dto := taskDTO.PostTaskRequest{}
	dto.MessageID = s.MessageID()
	dto.Name = s.Name()
	dto.Description = s.Description()

	for _, v := range s.sub {
		step := taskDTO.PostTaskStepRequest{}
		step.MessageID = v.MessageID()
		step.Name = v.Name()
		step.Description = v.Description()
		step.ExpectedExecutionMs = v.ExpectedExecutionMs()
		dto.TaskSteps = append(dto.TaskSteps, step)
	}
	return &dto
}

// Execute implements the interface of Action.
func (s *Refresh) Execute(ctx *context.Base) {
	var (
		enclosure base.ModelInterface
		dbImpl    db.Enclosure
	)

	dbImpl.TemplateImpl = new(db.Enclosure)
	// defer to unlock the enclosure.
	defer func() {
		// TODO we should check if we need turn to unmanaged.
		dbImpl.SetState(ctx.ID, model.StateReady, model.StateReasonAuto)
	}()

	// lock the enclosure.
	enclosure, err := dbImpl.GetAndLock(ctx.ID)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("Action refresh failed, get and lock enclosure failed, db operation failed.")
		// TODO set internal error.
		return
	}
	if enclosure == nil {
		log.WithFields(log.Fields{"error": err}).Warn("Action refresh failed, enclosure does not exist.")
		return
	}
	// create the task.
	createTaskRequest := s.Task()
	// TODO we should use client error here.
	createTaskResponse, errorResponse, err := taskSDK.CreateTask(createTaskRequest)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("Action refresh failed, create task failed.")
	}
	if errorResponse != nil && len(errorResponse) > 0 {
		log.WithFields(log.Fields{"error": errorResponse[0]}).Warn("Action refresh failed, create task failed.")
	}
	log.WithFields(log.Fields{"task": createTaskResponse.GetID()}).Info("Action refresh, create task.")
	// return the current work and continue.

	// execute each of the sub action
	for _, v := range s.sub {
		v.Execute(ctx)
	}
}
