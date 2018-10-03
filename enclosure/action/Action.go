package action

import (
	"promise/base"
	"promise/enclosure/context"
	taskDTO "promise/task/object/dto"
)

// Action defines the strategy.
type Action interface {
	MessageID() string
	Name() string
	Description() string
	ExpectedExecutionMs() uint64
	Add(sub Action) Action
	Task() *taskDTO.PostTaskRequest
	Execute(c *context.Base)
}

// Enclosure includes some basic opertion that an action needs.
type Enclosure struct {
}

// Lock the enclosure.
func (s Enclosure) Lock(c *context.Base) *base.ErrorResponse {
	return nil
}
