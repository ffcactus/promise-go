package message

import (
	"promise/base"
)

const (
	// MessageTaskNoStep is message ID.
	MessageTaskNoStep = "Task.Message.NoStep"
)

// NewMessageTaskNoStep returns a new message.
func NewMessageTaskNoStep() *base.Message {
	ret := base.NewMessage()
	ret.ID = MessageTaskNoStep
	ret.Severity = base.SeverityNormal
	ret.Description = "No task steps included.."
	ret.Supports = []base.Support{}
	return ret
}
