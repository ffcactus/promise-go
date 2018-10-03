package action

import (
	"promise/enclosure/object/constvalue"
	taskDTO "promise/task/object/dto"
)

// Refresh is the refresh action.
type Refresh struct {
	sub []Action
}

// Add the sub actions.
func (s *Refresh) Add(sub *Action) {
	s.sub = append(s.sub, *sub)
}

// Task returns the post task request.
func (s *Refresh) Task() *taskDTO.PostTaskRequest {
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
