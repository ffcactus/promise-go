package strategy

import (
	"promise/base"
	"promise/enclosure/context"
	"promise/enclosure/object/constvalue"
	"promise/enclosure/object/model"
	taskDTO "promise/task/object/dto"
	"time"
)

// Refresh is the refresh strategy.
type Refresh struct {
	sub []Strategy
}

// NewRefresh creates the refresh strategy.
func NewRefresh(ctx *context.RefreshContext) MainStrategy {
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
	if base.ContainsString(ctx.Request.Targets, model.RefreshAppliance) {
		ret.Add(NewRefreshAppliance())
	}
	if base.ContainsString(ctx.Request.Targets, model.RefreshPower) {
		ret.Add(NewRefreshPower())
	}
	if base.ContainsString(ctx.Request.Targets, model.RefreshFan) {
		ret.Add(NewRefreshFan())
	}
	return &ret
}

// Add the sub strategys.
func (s *Refresh) Add(sub Strategy) {
	s.sub = append(s.sub, sub)
}

// Name returns the name of the strategy.
func (Refresh) Name() string {
	return "Refresh Enclosure"
}

// MessageID returns the message ID of the strategy.
func (Refresh) MessageID() string {
	return constvalue.RefreshTaskID
}

// Description returns the description of the strategy.
func (Refresh) Description() string {
	return "Refresh enclosure's settings and component."
}

// ExpectedExecutionMs returns the expected execution time in ms of the strategy.
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
	// execute each of the sub strategy
	for _, v := range s.sub {
		time.Sleep(time.Duration(5000) * time.Millisecond)
		v.Execute(ctx)
	}
}
