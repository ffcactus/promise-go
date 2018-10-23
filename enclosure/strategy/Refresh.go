package strategy

import (
	"promise/base"
	"promise/enclosure/context"
	"promise/enclosure/object/constvalue"
	"promise/enclosure/object/model"
	taskDTO "promise/task/object/dto"
)

// Refresh is the interface that a refresh operation should support.
type Refresh interface {
	Execute(c context.Refresh)
	Task() *taskDTO.PostTaskRequest
}

// TaskStepStrategy is an kind of strategy that corresponding to a task step.
type TaskStepStrategy interface {
	Execute(c context.Refresh)
	MessageID() string
	Name() string
	Description() string
	ExpectedExecutionMs() uint64
}

// RefreshImpl implements Refresh interface.
type RefreshImpl struct {
	sub []TaskStepStrategy
}

// NewRefresh creates the refresh strategy.
func NewRefresh(ctx context.Refresh) Refresh {
	ret := RefreshImpl{}

	if base.ContainsString(ctx.GetRequest().Targets, model.RefreshManager) {
		ret.Add(NewRefreshManager())
	}
	if base.ContainsString(ctx.GetRequest().Targets, model.RefreshServer) {
		ret.Add(NewRefreshServer())
	}
	if base.ContainsString(ctx.GetRequest().Targets, model.RefreshSwitch) {
		ret.Add(NewRefreshSwitch())
	}
	if base.ContainsString(ctx.GetRequest().Targets, model.RefreshAppliance) {
		ret.Add(NewRefreshAppliance())
	}
	if base.ContainsString(ctx.GetRequest().Targets, model.RefreshPower) {
		ret.Add(NewRefreshPower())
	}
	if base.ContainsString(ctx.GetRequest().Targets, model.RefreshFan) {
		ret.Add(NewRefreshFan())
	}
	ret.Add(&RefreshState{})
	return &ret
}

// Add the sub strategys.
func (s *RefreshImpl) Add(sub TaskStepStrategy) {
	s.sub = append(s.sub, sub)
}

// Task returns the post task request.
func (s *RefreshImpl) Task() *taskDTO.PostTaskRequest {
	dto := taskDTO.PostTaskRequest{}
	dto.MessageID = constvalue.RefreshTaskID
	dto.Name = "Refresh Enclosure"
	dto.Description = "Refresh enclosure's settings and component."
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

// Execute do the refresh work.
func (s *RefreshImpl) Execute(ctx context.Refresh) {
	// execute each of the sub strategy
	for _, v := range s.sub {
		v.Execute(ctx)
	}
}
