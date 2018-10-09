package strategy

import (
	"promise/enclosure/context"
	taskDTO "promise/task/object/dto"
)

// Strategy defines the strategy.
type Strategy interface {
	MessageID() string
	Name() string
	Description() string
	ExpectedExecutionMs() uint64
	Execute(c *context.Base)
}

// MainStrategy is an kind of strategy that corresponding to a task.
type MainStrategy interface {
	Strategy
	Task() *taskDTO.PostTaskRequest
}
