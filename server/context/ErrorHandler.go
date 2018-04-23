package context

import (
	"promise/base"
)

// ErrorHandlerInterface The interface of error handler.
type ErrorHandlerInterface interface {
	AppendErrorMessage(message base.Message)
	ErrorMessages() *[]base.Message
}

// ErrorHandler The implementation of error handler.
type ErrorHandler struct {
	messages []base.Message
}

// AppendMessage Append message.
func (c *Base) AppendMessage(message base.Message) {
	c.messages = append(c.messages, message)
}

// AppendMessages Append messages.
func (c *Base) AppendMessages(messages []base.Message) {
	for i := range messages {
		c.messages = append(c.messages, messages[i])
	}
}

// Messages Get messages.
func (c *Base) Messages() []base.Message {
	return c.messages
}
